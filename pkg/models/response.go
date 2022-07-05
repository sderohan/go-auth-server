package models

type AuthResponse struct {
	Success    bool        `json:"success"`
	StatusCode uint        `json:"statusCode"`
	Data       interface{} `json:"data"`
}

type TokenResponse struct {
	Username string `json:"username"`
	Role     string `json:"role"`
	Token    string `json:"token"`
}
