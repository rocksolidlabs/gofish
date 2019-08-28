//
// SPDX-License-Identifier: BSD-3-Clause
//

package gofish

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"net/http"

	"strings"

	"github.com/rocksolidlabs/gofish/common"
)

// ApiClient represents a connection to a Redfish/Swordfish enabled service
// or device.
type ApiClient struct {
	// Endpoint is the URL of the *fish service
	Endpoint string

	// Token is the session token to be used for all requests issued
	Token string

	// httpClient is for direct http actions
	httpClient *http.Client
}

// APIClient creates a new client connection to a Redfish service.
func APIClient(endpoint string, httpClient *http.Client) (c *ApiClient, err error) {
	if !strings.HasPrefix(endpoint, "http") {
		return c, fmt.Errorf("endpoint must starts with http or https")
	}
	client := &ApiClient{Endpoint: endpoint}
	if httpClient != nil {
		client.httpClient = httpClient
	} else {
		client.httpClient = &http.Client{}
	}
	return client, err
}

// Get performs a GET request against the Redfish service.
func (c *ApiClient) Get(relativePath string) (*http.Response, error) {
	return c.do(relativePath, http.MethodGet, nil, http.StatusOK)
}

// Post performs a Post request against the Redfish service.
func (c *ApiClient) Post(relativePath string, payload []byte) (*http.Response, error) {
	return c.do(relativePath, http.MethodPost, payload, http.StatusOK, http.StatusCreated, http.StatusAccepted, http.StatusNoContent)
}

// Put makes a PUT call.
func (c *ApiClient) Put(relativePath string, payload []byte) (*http.Response, error) {
	return c.do(relativePath, http.MethodPut, payload, http.StatusOK, http.StatusCreated, http.StatusAccepted, http.StatusNoContent)
}

// Patch makes a PATCH call.
func (c *ApiClient) Patch(relativePath string, payload []byte) (*http.Response, error) {
	return c.do(relativePath, http.MethodPatch, payload, http.StatusOK, http.StatusCreated, http.StatusAccepted, http.StatusNoContent)
}

// Delete performs a Delete request against the Redfish service.
func (c *ApiClient) Delete(relativePath string) (*http.Response, error) {
	return c.do(relativePath, http.MethodDelete, nil, http.StatusOK, http.StatusAccepted, http.StatusNoContent)
}

func (c *ApiClient) do(relativePath, method string, payload []byte, statuses ...int) (*http.Response, error) {
	if relativePath == "" {
		relativePath = common.DefaultServiceRoot
	}

	var req *http.Request
	var err error
	endpoint := fmt.Sprintf("%s%s", c.Endpoint, relativePath)
	if payload == nil {
		req, err = http.NewRequest(method, endpoint, nil)
	} else {
		req, err = http.NewRequest(method, endpoint, bytes.NewBuffer(payload))
		req.Header.Set("Content-Type", "application/json")
	}
	if err != nil {
		return nil, err
	}

	req.Header.Set("User-Agent", "gofish/1.0.0")
	req.Header.Set("Accept", "application/json")
	if c.Token != "" {
		req.Header.Set("X-Auth-Token", c.Token)
	}
	req.Close = true

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return nil, err
	}

	if !checkStatus(resp.StatusCode, statuses...) {
		payload, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			return nil, err
		}
		defer resp.Body.Close()
		return nil, fmt.Errorf("%d: %s", resp.StatusCode, string(payload))
	}

	return resp, err
}

func checkStatus(status int, statuses ...int) bool {
	for _, s := range statuses {
		if status == s {
			return true
		}
	}
	return false
}
