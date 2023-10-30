package qdarklaunch

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"time"
)

type DarklaunchResponse struct {
	Result bool
	Error  error
}

// HTTPClient interface
type HTTPClient interface {
	Do(req *http.Request) (*http.Response, error)
}

var (
	Client HTTPClient
)

func init() {
	Client = &http.Client{
		Timeout: 60 * time.Second,
	}
}

func validateParams(name string, version string) (bool, string) {
	if name == "" && version == "" {
		return false, "darklaunch name and version cannot be empty"
	} else if name == "" {
		return false, "darklaunch name cannot be blank"
	} else if version == "" {
		return false, "darklaunch version cannot be blank"
	} else {
		return true, ""
	}
}

func GetDarklaunch(version string, name string, userId string) (bool, error) {
	isValid, errorMessage := validateParams(name, version)
	if !isValid {
		return false, errors.New(errorMessage)
	}

	darklaunchUrl := fmt.Sprintf("http://api/%v/darklaunch/%v?user[id]=%v", version, name, userId)
	resp, err := httpGetDarklaunch(darklaunchUrl)
	return resp, err
}

var httpGetDarklaunch = func(url string) (bool, error) {
	req, _ := http.NewRequest(http.MethodGet, url, nil)
	resp, err := Client.Do(req)
	if err != nil {
		return false, errors.New(err.Error())
	}

	body, err := io.ReadAll(resp.Body)

	if err != nil {
		return false, errors.New(err.Error())
	}

	var response DarklaunchResponse
	json.Unmarshal([]byte(body), &response)
	if response.Error != nil {
		return false, response.Error
	}

	return response.Result, nil
}
