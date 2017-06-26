package api

import (
	"log"
	"os"
	"testing"
)

func TestNamespaces(t *testing.T) {
	api := &ApiClient{
		Credentials: &ApiCredentials{
			Key:    os.ExpandEnv("${APP_KEY}"),
			Secret: os.ExpandEnv("${APP_SECRET}"),
		},
	}
	if namespaces, err := api.ListNamespaces(); err != nil {
		log.Printf("Error %T: %v", err, err)
	} else {
		for index, ns := range namespaces {
			log.Println(index, "=", ns)
		}
	}
}
