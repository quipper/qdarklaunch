package qdarklaunch

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
)

type DarklaunchResponse struct {
	Result bool
	Error  string
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

func GetDarklaunch(version string, name string) (bool, error) {
	isValid, errorMessage := validateParams(name, version)
	if !isValid {
		return false, errors.New(errorMessage)
	}

	darklaunchUrl := fmt.Sprintf("http://api/%v/darklaunch/%v", version, name)
	resp, err := http.Get(darklaunchUrl)
	if err != nil {
		return false, errors.New(err.Error())
	}

	// Read the response body
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return false, errors.New(err.Error())
	}

	var response DarklaunchResponse
	if err := json.Unmarshal([]byte(body), &response); err != nil {
		return false, errors.New(response.Error)
	}

	return response.Result, nil
}
