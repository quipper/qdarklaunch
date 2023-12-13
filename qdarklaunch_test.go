package qdarklaunch_test

import (
	"bytes"
	"errors"
	"fmt"
	"github.com/quipper/qdarklaunch"
	"github.com/quipper/qdarklaunch/test/mocks"
	"github.com/stretchr/testify/assert"
	"io"
	"net/http"
	"testing"
)

type Tests struct {
	name              string
	darklaunchValue   bool
	darklaunchVersion string
	darklaunchName    string
	darklauncUserId   string
	response          *qdarklaunch.DarklaunchResponse
	expectedError     error
}

func init() {
	qdarklaunch.Client = &mocks.MockClient{}
}

func TestGetDarklaunch(t *testing.T) {
	tests := []Tests{
		{
			name:              "darklaunch is false",
			darklaunchValue:   false,
			darklaunchVersion: "v1",
			darklaunchName:    "enable-testing",
			darklauncUserId:   "12345",
			response: &qdarklaunch.DarklaunchResponse{
				Result: false,
				Error:  nil,
			},
			expectedError: errors.New(""),
		},
		{
			name:              "darklaunch is true",
			darklaunchValue:   true,
			darklaunchVersion: "v1",
			darklaunchName:    "enable-testing",
			darklauncUserId:   "12345",
			response: &qdarklaunch.DarklaunchResponse{
				Result: true,
				Error:  nil,
			},
			expectedError: errors.New(""),
		},
		{
			name:              "error missing darklaunch version",
			darklaunchValue:   false,
			darklaunchVersion: "",
			darklaunchName:    "enable-testing",
			darklauncUserId:   "12345",
			response: &qdarklaunch.DarklaunchResponse{
				Result: false,
				Error:  nil,
			},
			expectedError: errors.New("darklaunch version cannot be blank"),
		},
		{
			name:              "error missing darklaunch name",
			darklaunchValue:   false,
			darklaunchVersion: "v1",
			darklaunchName:    "",
			darklauncUserId:   "12345",
			response: &qdarklaunch.DarklaunchResponse{
				Result: false,
				Error:  nil,
			},
			expectedError: errors.New("darklaunch name cannot be blank"),
		},
		{
			name:              "error missing darklaunch name and version",
			darklaunchValue:   false,
			darklaunchVersion: "",
			darklaunchName:    "",
			darklauncUserId:   "12345",
			response: &qdarklaunch.DarklaunchResponse{
				Result: false,
				Error:  nil,
			},
			expectedError: errors.New("darklaunch name and version cannot be empty"),
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			// build response JSON
			s := fmt.Sprintf(string(`{"result": %v}`), test.darklaunchValue)
			// create a new reader with that JSON
			r := io.NopCloser(bytes.NewReader([]byte(s)))
			mocks.GetDoFunc = func(*http.Request) (*http.Response, error) {
				return &http.Response{
					StatusCode: 200,
					Body:       r,
				}, nil
			}

			resp, err := qdarklaunch.GetDarklaunch(test.darklaunchVersion, test.darklaunchName, test.darklauncUserId)
			if err != nil {
				assert.Equal(t, err, test.expectedError, "the value should be equal")
			}
			assert.Equal(t, resp, test.response.Result, "the value should be equal")
		})
	}
}
