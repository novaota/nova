package rest

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/http/cookiejar"

	"github.com/gorilla/sessions"
	"nova/shared"
)

var ContentType string = "application/json"

func NewClient(certSettings shared.CertificateSettings, endpoint string) *Client {
	cookieJar, _ := cookiejar.New(nil)
	return &Client{
		CertificateSettings: certSettings,
		Endpoint: endpoint,
		client: &http.Client{
			Jar: cookieJar,
		},
	}
}

type Client struct {
	shared.CertificateSettings
	Endpoint string
	cookieStore sessions.CookieStore
	client *http.Client
}

func (c *Client) Get(out interface{}, urlFormat string, a ...interface{}) error {
	return c.do("GET", nil, out, urlFormat, a)
}

func (c *Client) Post(body interface{}, out interface{}, urlFormat string, a ...interface{}) error{
	return c.do("POST", body, out, urlFormat, a)
}

func (c *Client) Put(body interface{}, out interface{}, urlFormat string, a ...interface{}) error {
	return c.do("POST", body, out, urlFormat, a)
}

func (c *Client) Delete(out interface{}, urlFormat string, a ...interface{}) error {
	return c.do("DELETE", nil, out, urlFormat, a)
}

func (c *Client) do(method string, body interface{}, out interface{}, urlFormat string, a []interface{}) error {
	var urlAction string

	if len(a) == 0 || (len(a) == 1 && a[0] == nil) {
		urlAction = urlFormat
	} else {
		urlAction = fmt.Sprintf(urlFormat, a)
	}

	url := c.Endpoint + urlAction
	bodyReader, err := c.createBodyReader(body)

	if err != nil {
		return err
	}

	req, err := http.NewRequest(method, url, bodyReader)

	if err != nil {
		return err
	}

	resp, err := c.client.Do(req)

	if err != nil {
		return err
	}

	data := c.readResponse(resp)

	if resp.StatusCode != http.StatusOK {
		return NewRestError(resp.StatusCode)
	}

	if out != nil {
		err := json.Unmarshal(data, out)
		if err != nil {
			return err
		}
	}

	return nil
}

func (c *Client) createBodyReader(in interface{}) (io.Reader, error) {
	if in == nil {
		return nil, nil
	}

	jsonBody, err := json.Marshal(in)

	if err != nil {
		return nil, err
	}

	return bytes.NewReader(jsonBody), nil
}

func (c *Client) readResponse(resp *http.Response) []byte {
	buf := new(bytes.Buffer)
	buf.ReadFrom(resp.Body)
	return buf.Bytes()
}

