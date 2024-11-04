package model

type User struct {
	ID        	int    `json:"id"`
	Username  	string `json:"username"`
	DisplayName string `json:"displayName"`
}
