package main

import (
	"net/http"
	"time"
)

type Client struct {
	httpClient http.Client
}

func NewPokeClient(timeout time.Duration) Client {
	return Client{
		httpClient: http.Client{
			Timeout: timeout,
		},
	}
}
