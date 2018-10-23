package hexon

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

const (
	basePath string = "https://api.hexon.nl/spi/api/v2/rest"
)

type HexonCredentials struct {
	siteCode string
	username string
	password string
}

type Client struct {
	HexonCredentials
	http http.Client
}

// CreateVehicle takes a vehicle and creates it in hexon, returning an error if the operation fails
func (c *Client) CreateVehicle(vehicle Vehicle) error {
	req, err := c.req(http.MethodPost, "vehicle/", payloadify(vehicle))
	if err != nil {
		return err
	}

	resp, err := c.http.Do(req)
	if err != nil {
		return err
	}
	return nil
}

func (c *Client) PublishVehicle(vin string) error {

}

func (c *Client) UpdateVehicle(vehicle Vehicle) error {

}

func (c *Client) req(method, path string, vehicle hexonvehicle) (*http.Request, error) {
	json, err := json.Marshal(vehicle)
	if err != nil {
		return nil, err
	}
	req, err := http.NewRequest(method, fmt.Sprintf("%s/%s", basePath, path), bytes.NewBuffer(json))
	if err != nil {
		return nil, err
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Accept", "application/json")
	req.SetBasicAuth(c.username, c.password)
	return req, nil
}

func NewClient(credentials HexonCredentials) *Client {
	return &Client{
		HexonCredentials: credentials,
		http: http.Client{
			Timeout: time.Second * 5,
		},
	}
}
