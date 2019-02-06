package shared

import (
	"bytes"
	"errors"
	"io"
	"net/http"
	"os"
)

//Whitelist urls
var Whitelist []string = []string{"https://www.almesberger.com"}

var DefaultHTTPUtil *HTTPUtil = &HTTPUtil{}

type HTTPUtil struct {
	AllowUnsecureConnections bool
}

func (util *HTTPUtil) Get(url string) (string, error) {
	buf, err := util.GetResponseBuffer(url)

	if err != nil {
		return "", err
	}
	return buf.String(), nil
}

func (util *HTTPUtil) GetBytes(url string)  ([]byte, error){
	buf, err := util.GetResponseBuffer(url)

	if err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}

func (util *HTTPUtil) GetResponseBuffer(url string) (*bytes.Buffer, error) {
	if !util.isConnectionGranted(url) {
		return nil, errors.New("Connection could not be granted")
	}

	resp, err := http.Get(url)

	if err != nil {
		return nil, err
	}

	buf := new(bytes.Buffer)
	buf.ReadFrom(resp.Body)
	return buf, nil
}

func (util *HTTPUtil) Download(url string, destination string) error {
	if !util.isConnectionGranted(url) {
		return errors.New("Connection could not be granted")
	}

	// Create the file
	out, err := os.Create(destination)
	if err != nil {
		return err
	}
	defer out.Close()

	// Get the data
	resp, err := http.Get(url)
	if err != nil {
		return err
	}

	defer resp.Body.Close()

	// Write the body to file
	_, err = io.Copy(out, resp.Body)
	if err != nil {
		return err
	}

	return nil
}

func (util *HTTPUtil) isConnectionGranted(url string) bool {
	if util.AllowUnsecureConnections {
		return true
	}

	//TODO: Further checking
	return true
}
