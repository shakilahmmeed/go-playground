// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package model

type Bike struct {
	ID    string `json:"id"`
	Model string `json:"model"`
	Owner string `json:"owner"`
}

type Login struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type NewBike struct {
	ID    string `json:"id"`
	Model string `json:"model"`
	Owner string `json:"owner"`
}

type NewLink struct {
	Title   string `json:"title"`
	Address string `json:"address"`
}

type NewUser struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type RefreshTokenInput struct {
	Token string `json:"token"`
}

type User struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}
