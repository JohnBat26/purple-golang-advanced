package verify

type VerifyResponse struct {
	Email    string `json:"email"`
	Verified bool   `json:"verified"`
}

type SendRequest struct {
	Email string `json:"email" validate:"required,email"`
}
