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

// Credentials are used to authenticate a hexon client
type Credentials struct {
	Username string
	Password string
}

// Validate can be called to ensure the instance of the hexon credentials is valid.
// In the event of invalid credentials, this method will return an error
func (hc *Credentials) Validate() error {
	errs := make([]error, 0, 3)

	if hc.Username == "" {
		errs = append(errs, ErrorEmptyUserName)
	}
	if hc.Password == "" {
		errs = append(errs, ErrorEmptyPassword)
	}

	if len(errs) > 0 {
		return concatErrors("The provided credentials are invalid", errs...)
	}

	return nil
}

// LinkType represents a link returned by the hexon api
type LinkType struct {
	Rel    string `json:"rel"`
	Href   string `json:"href"`
	Method string `json:"method"`
}

// APIResponse is the response returned from a request to hexon
type APIResponse struct {
	Errors   []string   `json:"errors"`
	Warnings []string   `json:"warnings"`
	Debugs   []string   `json:"debugs"`
	Result   []string   `json:"result"`
	Links    []LinkType `json:"_links"`
}

// Client is the struct used to interface with the Hexon Api. It is intended to be intialized using the
// NewClient factory constructor
type Client struct {
	Credentials
	http http.Client
}

// CreateVehicle takes a vehicle and creates it in hexon,
// returning the api response and an error if one occurred.
func (c *Client) CreateVehicle(vehicle Vehicle) (*APIResponse, error) {
	req, err := c.req(http.MethodPost, "vehicle/", payloadify(vehicle))
	if err != nil {
		return nil, err
	}
	return c.send(req)
}

// PublishVehicle instructs the hexon api to publish an existing vehicle
// identified by Vin to the provided channel identified by sitecode.
// It will return the api response and an error if one occurred.
func (c *Client) PublishVehicle(vin, sitecode string) (*APIResponse, error) {
	msg, err := makePublishMessage(vin, sitecode)
	if err != nil {
		return nil, err
	}
	req, err := c.req(http.MethodPost, "vehiclesiteselections/", msg)
	if err != nil {
		return nil, err
	}
	return c.send(req)
}

// UpdateVehicle will send an update to hexon for the provided vehicle.
// It will return the api response and an error if one occurred.
func (c *Client) UpdateVehicle(vehicle Vehicle) (*APIResponse, error) {
	req, err := c.req(http.MethodPut, "vehicle/"+vehicle.Vin, payloadify(vehicle))
	if err != nil {
		return nil, err
	}
	return c.send(req)
}

func (c *Client) req(method, path string, payload interface{}) (*http.Request, error) {
	json, err := json.Marshal(payload)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest(method, fmt.Sprintf("%s/%s", basePath, path), bytes.NewBuffer(json))
	if err != nil {
		return nil, err
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Accept-Language", "en_GB")
	req.Header.Set("Accept", "application/json")
	req.SetBasicAuth(c.Username, c.Password)
	return req, nil
}

func (c *Client) send(req *http.Request) (*APIResponse, error) {
	res, err := c.http.Do(req)
	if err != nil {
		return nil, err
	}
	var apiRes APIResponse
	err = json.NewDecoder(res.Body).Decode(&apiRes)
	if err != nil {
		return nil, err
	}
	if res.StatusCode < 200 || res.StatusCode > 299 {
		err = ErrorHexonAPI
	}
	res.Body.Close()
	return &apiRes, err
}

// NewClient a factory constructor for initializing a hexon client
func NewClient(credentials Credentials) *Client {
	return &Client{
		Credentials: credentials,
		http: http.Client{
			Timeout: time.Second * 5,
		},
	}
}
