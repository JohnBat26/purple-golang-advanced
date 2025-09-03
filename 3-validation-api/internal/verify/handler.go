// Package verify for verify email
package verify

import (
	"crypto/rand"
	"crypto/sha256"
	"encoding/base64"
	"fmt"
	"log"
	"net/http"
	"net/smtp"
	"time"

	"demo/3-validation-api/configs"
	"demo/3-validation-api/internal/datastore"
	"demo/3-validation-api/pkg/email"
	"demo/3-validation-api/pkg/req"
	"demo/3-validation-api/pkg/res"
)

type VerifyHandler struct {
	*configs.Config
	EmailSender email.EmailSender
	DataStore   *datastore.DataStore
}

type VerifyHandlerDeps struct {
	*configs.Config
	emailSender *email.EmailSender
	dataStore   *datastore.DataStore
}

type EmailRequest struct {
	Email string `json:"email"`
}

func NewVerifyHandler(router *http.ServeMux, deps VerifyHandlerDeps) {
	smtpMailSender := email.SMTPEmailSender{
		SMTPAddr: deps.Address + ":" + deps.Port,
		Auth:     smtp.PlainAuth("", deps.Config.Email, deps.Config.Password, deps.Config.Address),
	}

	dataStore := datastore.NewDataStore(deps.Config.StoreFilename)

	handler := &VerifyHandler{
		Config:      deps.Config,
		EmailSender: &smtpMailSender,
		DataStore:   dataStore,
	}

	router.HandleFunc("POST /send", handler.Send())
	router.HandleFunc("GET /verify/{hash}", handler.Verify())
}

func (handler *VerifyHandler) Send() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		body, err := req.HandleBody[SendRequest](&w, r)
		if err != nil {
			return
		}

		hash, _, _ := generateEmailVerificationHash(body.Email)
		verifyBody := `
        <html>
        <body>
            <h3>Здравствуйте, ` + handler.Email + `!</h3>
            <p>Для подтверждения вашего email, пожалуйста перейдите по ссылке:
            <a href="http://localhost:8081/verify/` + hash + `"> Подверждение email </a>.
        </body>
        </html>
    `

		item := datastore.Item{Email: body.Email, Hash: hash}

		if handler.DataStore.FindByEmail(body.Email) != nil {
			handler.DataStore.RemoveItem(item)
		}

		handler.DataStore.AddItem(item)

		err = handler.EmailSender.SendEmail(body.Email, "Подтверждение адреса email", verifyBody, *handler.Config)
		if err != nil {
			log.Println("Ошибка отправки: ", err)
			res.JSON(w, err, http.StatusInternalServerError)
			return
		}

		data := VerifyResponse{
			Email: body.Email,
		}

		res.JSON(w, data, http.StatusCreated)
	}
}

func (handler *VerifyHandler) Verify() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		hash := r.PathValue("hash")

		item := handler.DataStore.FindByHash(hash)

		var data VerifyResponse

		if item != nil {
			data = VerifyResponse{
				Email:    item.Email,
				Verified: true,
			}
			handler.DataStore.RemoveItem(*item)
		} else {
			data = VerifyResponse{
				Email:    "",
				Verified: false,
			}
		}

		res.JSON(w, data, http.StatusOK)
	}
}

func generateEmailVerificationHash(email string) (string, string, error) {
	// Генерация уникального токена
	token := make([]byte, 32)
	if _, err := rand.Read(token); err != nil {
		return "", "", err
	}

	// Добавление временной метки и email
	data := fmt.Sprintf("%s:%s:%d", email, base64.URLEncoding.EncodeToString(token), time.Now().UnixNano())

	// Создание хеша
	hash := sha256.Sum256([]byte(data))
	verificationHash := base64.URLEncoding.EncodeToString(hash[:])

	return verificationHash, base64.URLEncoding.EncodeToString(token), nil
}
