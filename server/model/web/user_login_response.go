package web

type UserLoginJWT struct {
	AccessToken string 
	ID string 
	Username string 
}

type UserLoginResponse struct {
	ID string `json:"id"`
	Username string `json:"username"`
}