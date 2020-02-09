package admin

// LoginRequest ...
type LoginRequest struct {
	Phone    string `json:"phone"`
	Password string `json:"password"`
}

// Admin ...
type Admin struct {
	ID          string `json:"user_id"`
	AccessToken string `json:"access_token"`
}
