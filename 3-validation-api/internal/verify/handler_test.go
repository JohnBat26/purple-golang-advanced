package verify

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"demo/3-validation-api/configs"
	"demo/3-validation-api/internal/datastore"
)

type MockEmailSender struct {
	Called bool
}

func (m *MockEmailSender) SendEmail(to, subject, body string, config configs.Config) error {
	m.Called = true
	return nil
}

func TestSend(t *testing.T) {
	conf := configs.LoadConfig(".env.test")
	defer datastore.DropFile(conf.StoreFilename)

	mockSender := &MockEmailSender{}

	handler := &VerifyHandler{
		Config:      conf,
		EmailSender: mockSender,
		DataStore:   datastore.NewDataStore(conf.StoreFilename),
	}

	email := "test@example.com"

	doSend(t, email, handler, mockSender)
}

func TestVerify(t *testing.T) {
	conf := configs.LoadConfig(".env.test")
	defer datastore.DropFile(conf.StoreFilename)

	mockSender := &MockEmailSender{}

	handler := &VerifyHandler{
		Config:      conf,
		EmailSender: mockSender,
		DataStore:   datastore.NewDataStore(conf.StoreFilename),
	}
	email := "test@example.com"

	doSend(t, email, handler, mockSender)
	doVerify(t, email, handler)
}

func doSend(t *testing.T, email string, handler *VerifyHandler, mockSender *MockEmailSender) {
	t.Helper()

	reqBody := EmailRequest{
		Email: email,
	}

	jsonReq, _ := json.Marshal(reqBody)

	req := httptest.NewRequest(http.MethodPost, "/send", bytes.NewBuffer(jsonReq))
	req.Header.Set("Content-Type", "application/json")

	w := httptest.NewRecorder()
	handler.Send().ServeHTTP(w, req)

	if w.Code != http.StatusCreated {
		t.Errorf("Expected status code %d, got %d", http.StatusOK, w.Code)
	}

	expectedResponse := "test@example.com"
	if !strings.Contains(w.Body.String(), expectedResponse) {
		t.Errorf("Expected response %s, got %s", expectedResponse, w.Body.String())
	}

	if !mockSender.Called {
		t.Error("Expected SendEmail to be called")
	}

	item := handler.DataStore.FindByEmail(reqBody.Email)

	if item.Email != reqBody.Email {
		t.Errorf("Emails don't matg, got: %s, want: %s.", item.Email, reqBody.Email)
	}
}

func doVerify(t *testing.T, email string, handler *VerifyHandler) {
	t.Helper()

	hash := handler.DataStore.FindByEmail(email).Hash

	url := "/verify/{hash}"

	req := httptest.NewRequest(http.MethodGet, url, nil)
	req.Header.Set("Content-Type", "application/json")
	req.SetPathValue("hash", hash)

	w := httptest.NewRecorder()
	handler.Verify().ServeHTTP(w, req)

	if w.Code != http.StatusOK {
		t.Errorf("Expected status code %d, got %d", http.StatusOK, w.Code)
	}

	expectedResponse := "true"
	if !strings.Contains(w.Body.String(), expectedResponse) {
		t.Errorf("Expected response %s, got %s", expectedResponse, w.Body.String())
	}

	item := handler.DataStore.FindByEmail(email)

	if item != nil {
		t.Errorf("Items should be equals nil")
	}
}
