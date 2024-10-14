package model

type Data struct {
	ID          string `json:"id"`
	Name        string `json:"name"`
	Age         int    `json:"age"`
	IsConnected bool   `json:"isconnected"`
}
