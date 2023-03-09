package gastrofix_test

import (
	"log"
	"net/url"
	"os"
	"testing"

	gastrofix "github.com/omniboost/go-gastrofix"
)

var (
	client    *gastrofix.Client
	companyID int
)

func TestMain(m *testing.M) {
	baseURLString := os.Getenv("BASE_URL")

	// old authenticatio
	businessID := os.Getenv("BUSINESS_ID")
	token := os.Getenv("TOKEN")
	debug := os.Getenv("DEBUG")

	// setup old auth client
	client = gastrofix.NewClient(nil)
	client.SetBusinessID(businessID)
	client.SetToken(token)

	if debug != "" {
		client.SetDebug(true)
	}

	if baseURLString != "" {
		baseURL, err := url.Parse(baseURLString)
		if err != nil {
			log.Fatal(err)
		}
		client.SetBaseURL(*baseURL)
	}

	client.SetDisallowUnknownFields(true)
	m.Run()
}
