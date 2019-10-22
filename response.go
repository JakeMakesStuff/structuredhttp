package structuredhttp

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"math"
	"net/http"
	"strconv"
)

// Response defines the higher level HTTP response.
type Response struct {
	RawResponse *http.Response
}

// Bytes gets the response as bytes.
func (r *Response) Bytes() ([]byte, error) {
	return ioutil.ReadAll(r.RawResponse.Body)
}

// JSON returns the result as a interface which can be converted how the user wishes.
func (r *Response) JSON() (interface{}, error) {
	b, err := r.Bytes()
	if err != nil {
		return nil, err
	}
	var BasicInterface interface{}
	err = json.Unmarshal(b, &BasicInterface)
	if err != nil {
		return nil, err
	}
	return BasicInterface, nil
}

// RaiseForStatus throws a error if the request is a 4XX/5XX.
func (r *Response) RaiseForStatus() error {
	FirstDigitStatus := math.Floor(float64(r.RawResponse.StatusCode) / 100)
	if FirstDigitStatus == 4 || FirstDigitStatus == 5 {
		return errors.New("returned the status " + strconv.Itoa(r.RawResponse.StatusCode))
	}
	return nil
}

// Text returns the status as a string.
func (r *Response) Text() (string, error) {
	b, err := r.Bytes()
	if err != nil {
		return "", err
	}
	return string(b), nil
}
