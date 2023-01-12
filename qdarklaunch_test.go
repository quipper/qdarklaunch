package qdarklaunch

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type Tests struct {
	name            string
	darklaunchValue bool
	response        *DarklaunchResponse
	expectedError   error
}

func TestGetDarklaunch(t *testing.T) {
	tests := []Tests{
		{
			name:            "darklaunch-is-true",
			darklaunchValue: true,
			response: &DarklaunchResponse{
				Result: true,
				Error:  "",
			},
			expectedError: nil,
		},
		{
			name:            "darklaunch-is-false",
			darklaunchValue: false,
			response: &DarklaunchResponse{
				Result: false,
				Error:  "",
			},
			expectedError: nil,
		},
	}

	// store actual httpGetDarklaunch function into _httpGetDarklaunch variable
	_httpGetDarklaunch := httpGetDarklaunch
	// reset it back at the end of testing using 'defer' function
	defer func() {
		httpGetDarklaunch = _httpGetDarklaunch
	}()

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			httpGetDarklaunch = func(string) (bool, error) {
				return test.darklaunchValue, nil
			}

			resp, _ := GetDarklaunch("v1", "darklaunch-name")
			assert.Equal(t, resp, test.darklaunchValue, "the value should be equal")
		})
	}
}
