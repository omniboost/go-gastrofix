package fortnox_test

import (
	"log"
	"net/url"
	"os"
	"testing"

	fortnox "github.com/omniboost/go-fortnox"
	"golang.org/x/oauth2"
)

var (
	client    *fortnox.Client
	companyID int
)

func TestMain(m *testing.M) {
	baseURLString := os.Getenv("BASE_URL")

	// old authenticatio
	clientSecret := os.Getenv("CLIENT_SECRET")
	accessToken := os.Getenv("ACCESS_TOKEN")

	// new oauth2 authentication
	clientID := os.Getenv("CLIENT_ID")
	clientSecret = os.Getenv("CLIENT_SECRET")
	refreshToken := os.Getenv("REFRESH_TOKEN")
	tokenURL := os.Getenv("TOKEN_URL")

	debug := os.Getenv("DEBUG")

	oauthConfig := fortnox.NewOauth2Config()
	oauthConfig.ClientID = clientID
	oauthConfig.ClientSecret = clientSecret

	if clientID != "" && clientSecret != "" && refreshToken != "" {
		// setup oauth2 client

		// set alternative token url
		if tokenURL != "" {
			oauthConfig.Endpoint.TokenURL = tokenURL
		}

		token := &oauth2.Token{
			RefreshToken: refreshToken,
		}

		// get http client with automatic oauth logic
		httpClient := oauthConfig.Client(oauth2.NoContext, token)
		client = fortnox.NewClient(httpClient)
	} else {
		// setup old auth client
		client = fortnox.NewClient(nil, clientSecret, accessToken)
		client.SetClientSecret(clientSecret)
		client.SetAccessToken(accessToken)
	}

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
