package api

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"sync"
)

type ApiCredentials struct {
	sync.Mutex
	Key    string `json:"app_key"`
	Secret string `json:"app_secret"`
	Token  string `json:"-"`
}

type ApiClient struct {
	Credentials *ApiCredentials
	TokenStore  ApiTokenStore
	Endpoint    string
}

type ApiTokenStore interface {
	ReadToken() string
	WriteToken(string)
}

func (c *ApiClient) Get(urlStr string, respTo interface{}) error {
	return c.Do("GET", urlStr, nil, respTo)
}

func (c *ApiClient) Post(urlStr string, body interface{}, respTo interface{}) error {
	return c.Do("POST", urlStr, body, respTo)
}

func (c *ApiClient) Do(method, urlStr string, body interface{}, respTo interface{}) error {
	return c.execute(method, urlStr, body, c.AccessToken(), respTo)
}

func (c *ApiClient) AccessToken() string {
	store := c.tokenStore()
	if token := store.ReadToken(); token != "" {
		return token
	}

	token := &struct {
		Token string `json:"token"`
	}{}
	if c.execute("POST", "/api/v1/token", c.Credentials, "", token) == nil {
		store.WriteToken(token.Token)
		return token.Token
	}
	return ""
}

func (c *ApiClient) httpClient() *http.Client {
	return http.DefaultClient
}

func cloneRequest(ireq *http.Request, deepCopy bool) *http.Request {
	req := new(http.Request)
	*req = *ireq // shallow clone
	if deepCopy {
		h, h2 := req.Header, make(http.Header, len(req.Header))
		for k, vv := range h {
			vv2 := make([]string, len(vv))
			copy(vv2, vv)
			h2[k] = vv2
		}
		req.Header = h2
	}
	return req
}

var defaultEndpoint, _ = url.Parse("https://open.c.163.com/")

func (c *ApiClient) endpoint() *url.URL {
	if c.Endpoint != "" {
		ret, err := url.Parse(c.Endpoint)
		if err != nil {
			log.Fatalln(err)
		}
		return ret
	}
	return defaultEndpoint
}

func (c *ApiClient) tokenStore() ApiTokenStore {
	if c.TokenStore != nil {
		return c.TokenStore
	}
	return c.Credentials
}

func (store *ApiCredentials) ReadToken() string {
	return store.Token
}
func (store *ApiCredentials) WriteToken(token string) {
	store.Lock()
	defer store.Unlock()
	store.Token = token
}

func (c *ApiClient) execute(method, urlStr string, body interface{}, accessToken string, respTo interface{}) error {
	req, buffer, err := httpRequest(c, method, urlStr, body, accessToken)
	log.Printf("[DEBUG][API] %s %s %s", method, urlStr, string(buffer))
	if err != nil {
		return &ApiError{Status: -1, Cause: err}
	}
	resp, err := c.httpClient().Do(req)
	if err != nil {
		return &ApiError{Cause: err}
	}
	result, buffer, err := httpResponseToResult(resp, respTo)
	log.Printf("[DEBUG][API] HTTP %d %s %s", resp.StatusCode, urlStr, string(buffer))
	if err != nil {
		return &ApiError{Status: resp.StatusCode, Cause: err}
	}
	return result
}

func httpRequest(c *ApiClient, method, urlStr string, body interface{}, accessToken string) (*http.Request, []byte, error) {
	bodyReader, buffer, err := io.Reader(nil), []byte(nil), error(nil)
	if body != nil {
		if buffer, err = json.Marshal(body); err == nil {
			bodyReader = bytes.NewReader(buffer)
		} else {
			return nil, buffer, err
		}
	}

	if reqUrl, err := url.Parse(urlStr); err == nil {
		urlStr = c.endpoint().ResolveReference(reqUrl).String()
	} else {
		return nil, buffer, err
	}

	req, err := http.NewRequest(method, urlStr, bodyReader)
	if err != nil {
		return nil, buffer, err
	}
	if accessToken != "" {
		req.Header.Set("Authorization", fmt.Sprintf("Token %s", accessToken))
	}
	if bodyReader != nil {
		req.Header.Set("Content-Type", "application/json")
	}
	return req, buffer, nil
}

func httpResponseToResult(resp *http.Response, respTo interface{}) (error, []byte, error) {
	buffer, err := ioutil.ReadAll(resp.Body)
	if err == nil {
		if resp.StatusCode >= 200 && resp.StatusCode < 300 {
			if respTo != nil {
				return nil, buffer, json.Unmarshal(buffer, respTo)
			}
			return nil, buffer, nil
		}
		result := &ApiError{Status: resp.StatusCode}
		return result, buffer, json.Unmarshal(buffer, result)
	}
	return nil, nil, err
}
