package awardwallet

import "net/http"

//Client Cliente de la API
type Client struct {
	apikey string       // API Key
	client *http.Client // Cliente HTTP para realizar consultas
}

//NewClient Constructor Client
func NewClient(apiKey string) *Client {
	return &Client{
		apikey: apiKey,
		client: &http.Client{},
	}
}
