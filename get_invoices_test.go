package fortnox_test

import (
	"encoding/json"
	"log"
	"testing"
	"time"

	fortnox "github.com/omniboost/go-gastrofix"
)

func TestGetInvoices(t *testing.T) {
	req := client.NewGetInvoicesRequest()
	req.QueryParams().FromDate = fortnox.Date{time.Date(2022, 12, 1, 0, 0, 0, 0, time.Local)}
	req.QueryParams().ToDate = fortnox.Date{time.Date(2023, 3, 5, 0, 0, 0, 0, time.Local)}
	resp, err := req.Do()
	if err != nil {
		t.Error(err)
	}

	b, _ := json.MarshalIndent(resp, "", "  ")
	log.Println(string(b))
}
