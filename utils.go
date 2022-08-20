package main

import (
	"bytes"
	"encoding/base64"
	"fmt"
	"net/http"
)

type HeaderBuilder struct {
	headers map[string]string
}

func NewHeaderBuilder() *HeaderBuilder {
	return &HeaderBuilder{
		headers: make(map[string]string),
	}
}
func (hb *HeaderBuilder) Add(key, value string) *HeaderBuilder {
	hb.headers[key] = value
	return hb
}

func (hb *HeaderBuilder) ContentTypeJson() *HeaderBuilder {
	hb.Add("Content-Type", "application/json")
	return hb
}

func (hb *HeaderBuilder) BearerToken(token string) *HeaderBuilder {
	hb.Add("Authorization", fmt.Sprintf("Bearer %s", token))
	return hb
}

func (hb *HeaderBuilder) BasicAuth(username, password string) *HeaderBuilder {
	account := fmt.Sprintf("%s:%s", username, password)
	token := base64.StdEncoding.EncodeToString([]byte(account))
	hb.Add("Authorization", fmt.Sprintf("Basic %s", token))
	return hb
}

func (hb *HeaderBuilder) Auth(auth string) *HeaderBuilder {
	hb.Add("Authorization", auth)
	return hb
}

func (hb *HeaderBuilder) req(req *http.Request) {
	for key, value := range hb.headers {
		req.Header.Add(key, value)
	}
}
func Get(url string, hb *HeaderBuilder) (*http.Response, error) {
	client := &http.Client{}
	req, _ := http.NewRequest("GET", url, nil)
	hb.req(req)
	return client.Do(req)
}

func Post(url string, hb *HeaderBuilder, data []byte) (*http.Response, error) {
	client := &http.Client{}
	req, _ := http.NewRequest("GET", url, bytes.NewBuffer(data))
	hb.req(req)
	return client.Do(req)
}
