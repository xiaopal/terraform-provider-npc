package api

import (
	"log"
	"os"
	"testing"
)

func Test(t *testing.T) {
	api := &ApiClient{
		Credentials: &ApiCredentials{
			Key:    os.ExpandEnv("${APP_KEY}"),
			Secret: os.ExpandEnv("${APP_SECRET}"),
		},
	}
	log.Println("AccessToken = ", api.AccessToken())
	if err := api.Get("/api/", nil); err != nil {
		log.Println(err)
	} else {
		t.Fail()
	}
	if err := api.Get("/api/v1/namespaces", nil); err != nil {
		log.Println(err)
	} else {
		log.Println("GET /api/v1/namespaces OK")
	}
}
