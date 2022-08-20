package main

import (
	"encoding/json"
	"fmt"
)

const CFEndpoint = "https://api.cloudflare.com/client/v4"

type CFError struct {
	Code    int64  `json:"code"`
	Message string `json:"message"`
}

type CFResultList[T any] struct {
	Success  bool      `json:"success"`
	Errors   []CFError `json:"errors"`
	Messages []string  `json:"messages"`
	Results  []T       `json:"result"`
}
type CFResult[T any] struct {
	Success  bool      `json:"success"`
	Errors   []CFError `json:"errors"`
	Messages []string  `json:"messages"`
	Results  T         `json:"result"`
}

type CFKVNS struct {
	Id                  string `json:"id"`
	Title               string `json:"title"`
	SupportsUrlEncoding bool   `json:"supports_url_encoding"`
}

func (cfg *Config) ListKVNS() []CFKVNS {
	url := fmt.Sprintf("%s/accounts/%s/storage/kv/namespaces", CFEndpoint, cfg.CFAccount)
	hb := NewHeaderBuilder().ContentTypeJson().BearerToken(cfg.CFWorkerApiToken)
	resp, err := Get(url, hb)
	if err != nil {
		fmt.Printf("get kv ns error: %s", err)
	}
	defer resp.Body.Close()
	r := CFResultList[CFKVNS]{}
	err = json.NewDecoder(resp.Body).Decode(&r)

	if err != nil {
		fmt.Printf("decode json failed %s", err)
	}
	if r.Success {

		return r.Results
	} else {
		fmt.Println(r.Errors)
		return nil
	}
}

func (cfg *Config) CreateKVNS(name string) *CFKVNS {
	data := map[string]string{
		"title": name,
	}
	url := fmt.Sprintf("%s/accounts/%s/storage/kv/namespaces", CFEndpoint, cfg.CFAccount)
	data2, err := json.Marshal(data)
	if err != nil {
		fmt.Printf("json seailize failed : %s", err)
		return nil
	}
	hb := NewHeaderBuilder().ContentTypeJson().BearerToken(cfg.CFWorkerApiToken)
	resp, err := Post(url, hb, data2)
	if err != nil {
		fmt.Printf("get kv ns error: %s", err)
	}
	defer resp.Body.Close()
	r := CFResult[CFKVNS]{}
	err = json.NewDecoder(resp.Body).Decode(&r)

	if err != nil {
		fmt.Printf("decode json failed %s", err)
	}
	if r.Success {

		return &r.Results
	} else {
		fmt.Println(r.Errors)
		return nil
	}
}
