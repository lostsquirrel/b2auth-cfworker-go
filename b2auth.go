package main

import (
	"encoding/json"
	"fmt"
)

const b2ApiUrl = "https://api.backblazeb2.com/b2api/v2"

type B2AuthAccount struct {
	ApiUrl string `json:"apiUrl"`
	Token  string `json:"authorizationToken"`
}

type B2DownloadAuth struct {
	BucketId           string `json:"bucketId"`
	FileNamePrefix     string `json:"fileNamePrefix"`
	AuthorizationToken string `json:"authorizationToken"`
}

func (cfg *Config) AuthAccount() *B2AuthAccount {
	url := fmt.Sprintf("%s/b2_authorize_account", b2ApiUrl)
	hb := NewHeaderBuilder().ContentTypeJson().BasicAuth(cfg.B2AppKeyId, cfg.B2AppKey)
	resp, err := Get(url, hb)
	if err != nil {
		fmt.Printf("get kv ns error: %s", err)
		return nil
	}
	defer resp.Body.Close()
	r := B2AuthAccount{}
	err = json.NewDecoder(resp.Body).Decode(&r)

	if err != nil {
		fmt.Printf("decode json failed %s", err)
		return nil
	}
	return &r
}

func (cfg *Config) AuthToken(apiUrl, token string) string {
	url := fmt.Sprintf("%s/b2api/v2/b2_get_download_authorization", apiUrl)
	data := map[string]interface{}{
		"bucketId":               cfg.B2BucketId,
		"fileNamePrefix":         "",
		"validDurationInSeconds": cfg.B2TokenTTL,
	}
	hb := NewHeaderBuilder().ContentTypeJson().Auth(token)
	data2, err := json.Marshal(data)
	if err != nil {
		fmt.Printf("json seailize failed : %s", err)
		return ""
	}
	resp, err := Post(url, hb, data2)
	if err != nil {
		fmt.Printf("get kv ns error: %s", err)
		return ""
	}
	defer resp.Body.Close()
	r := B2DownloadAuth{}
	err = json.NewDecoder(resp.Body).Decode(&r)

	if err != nil {
		fmt.Printf("decode json failed %s", err)
	}
	return r.AuthorizationToken
}
