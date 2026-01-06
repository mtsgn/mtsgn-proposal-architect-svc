package easemob

import (
	"boilerplate-api/pkg/config"
	"fmt"
)

type Config = config.EasemobConfig

type Client struct {
	config Config
}

func NewClient(cfg Config) *Client {
	return &Client{
		config: cfg,
	}
}

type EasemobUserRegistrationRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type EasemobUserRegistrationResponse struct {
	Entities []struct {
		UUID string `json:"uuid"`
	} `json:"entities"`
}

func (c *Client) RegisterUser(req EasemobUserRegistrationRequest) (*EasemobUserRegistrationResponse, error) {
	// Stub implementation - replace with actual Easemob API call
	// For now, return a mock response
	return &EasemobUserRegistrationResponse{
		Entities: []struct {
			UUID string `json:"uuid"`
		}{
			{UUID: fmt.Sprintf("mock-uuid-%s", req.Username)},
		},
	}, nil
}
