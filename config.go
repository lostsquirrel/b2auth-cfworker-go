package main

import (
	"fmt"
	"os"
	"reflect"
)

type Config struct {
	CFWorkerApiToken string `env:"CF_WORKER_API_TOKEN"`
	CFAccount        string `env:"CF_ACCOUNT"`
	B2BucketId       string `env:"B2_BUCKET_ID"`
	B2AppKey         string `env:"B2_APP_KEY"`
	B2AppKeyId       string `env:"B2_APP_KEY_ID"`
	B2TokenTTL       int    `env:"B2_TOKEN_TTL"`
}

func BuildFromEnv() *Config {
	c := Config{}
	x := reflect.TypeOf(c)
	fields := reflect.VisibleFields(x)
	values := reflect.ValueOf(&c)
	for _, field := range fields {
		envKey := field.Tag.Get("env")
		envValue, ok := os.LookupEnv(envKey)
		if ok {
			rValue := reflect.Indirect(values).FieldByName(field.Name)
			fmt.Println(rValue)
			rValue.SetString(envValue)
		}
	}
	return &c
}
