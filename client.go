// Package revolut provides an SDK to access the Revolut for Business API.
package revolut

import (
	"bytes"
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
	"strings"
	"time"
)

// Client is the core structure for Revolut API access.
type Client struct {
	http.Client
	baseURL string
	//ErrorCode is the last HTTP error code.
	ErrorCode int
	// Agent should be customised per app.
	Agent string
	// bearer is the authentication header string, generated from the API key.
	bearer string
}

// ErrorResponse from JSON endpoints.
type ErrorResponse struct {
	// Message is a clarification of what went wrong.
	Message string `json:"message"`
	// Code is an internal error code.
	Code int `json:"code"`
}

const (
	defaultAgent  = "Revolut unofficial Go SDK"
	urlSandbox    = "https://sandbox-b2b.revolut.com/api/1.0/"
	urlProduction = "https://b2b.revolut.com/api/1.0/"
)

const (
	//
	// API Endpoints
	//
	epAccounts       = "accounts"
	epAccountDetails = "bank-details"
	epCounterparties = "counterparties"
	epCounterparty   = "counterparty"
	epTransfer       = "transfer"
	epPay            = "pay"
	epTransaction    = "transaction"
	epTransactions   = "transactions"
	epWebhook        = "webhook"

	//
	// Webhook events
	//

	// EventCreated is for TransactionCreated hooks.
	EventCreated = "TransactionCreated"
	// EventStateChange is for TransactionStateChanged hooks.
	EventStateChange = "TransactionStateChanged"
)

// NewClient creates a new Revolut client with some reasonable HTTP request defaults.
func NewClient(key string) (*Client, error) {
	c := Client{}
	c.Timeout = time.Second * 5
	c.Agent = defaultAgent
	c.Transport = &http.Transport{
		MaxIdleConns:        50,
		MaxIdleConnsPerHost: 50,
	}
	return &c, c.SetAPI(key)
}

// SetAPI sets the API key and type to use (sandbox or production).
// Sandbox keys start with "sand_" and production keys start with "prod_".
func (c *Client) SetAPI(key string) error {
	switch key[0:5] {
	case "sand_":
		c.baseURL = urlSandbox
	case "prod_":
		c.baseURL = urlProduction
	default:
		return errors.New(ErrKeyFormat)
	}

	c.bearer = "Bearer " + key
	return nil
}

// GetJSON builds the full endpoint path and gets the raw JSON.
func (c *Client) GetJSON(path string) ([]byte, int, error) {
	var url strings.Builder
	url.WriteString(c.baseURL)
	url.WriteString(path)
	req, err := http.NewRequest("GET", url.String(), nil)
	if err != nil {
		return nil, 0, err
	}

	c.setHeader(req)
	response, err := c.Do(req)
	defer response.Body.Close()
	if err != nil {
		return nil, response.StatusCode, err
	}

	contents, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return nil, response.StatusCode, err
	}

	return contents, response.StatusCode, nil
}

// PostJSON builds the full endpoint path and posts the provided data, returning the JSON response.
func (c *Client) PostJSON(path string, data interface{}) ([]byte, int, error) {
	msg, err := json.Marshal(data)
	if err != nil {
		return nil, 0, err
	}

	url := strings.Join([]string{c.baseURL, path}, "")
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(msg))
	if err != nil {
		return nil, 0, err
	}

	c.setHeader(req)
	response, err := c.Do(req)
	defer response.Body.Close()
	if err != nil {
		_, _ = ioutil.ReadAll(response.Body)
		return nil, response.StatusCode, err
	}

	contents, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return nil, response.StatusCode, err
	}

	return contents, response.StatusCode, nil
}

// Delete sends a delete command to an endpoint. The URL is the data and the HTTP response code is the only result.
func (c *Client) Delete(path string) ([]byte, int, error) {
	url := strings.Join([]string{c.baseURL, path}, "")
	req, err := http.NewRequest("DELETE", url, nil)
	if err != nil {
		return nil, 0, err
	}

	c.setHeader(req)
	response, err := c.Do(req)
	defer response.Body.Close()
	if err != nil {
		return nil, response.StatusCode, err
	}

	contents, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return nil, response.StatusCode, err
	}

	return contents, response.StatusCode, nil
}

// setHeader helper function.
func (c *Client) setHeader(req *http.Request) {
	req.Header.Set("Authorization", c.bearer)
}

// jsonError extracts the message from a JSON error response
func jsonError(data []byte) error {
	var resp ErrorResponse
	err := json.Unmarshal(data, &resp)
	if err != nil {
		return err
	}

	return errors.New(resp.Message)
}
