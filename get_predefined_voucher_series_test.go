package fortnox_test

import (
	"encoding/json"
	"log"
	"testing"
)

func TestGetPredefinedVoucherSeries(t *testing.T) {
	req := client.NewGetPredefinedVoucherSeriesRequest()
	resp, err := req.Do()
	if err != nil {
		t.Error(err)
	}

	b, _ := json.MarshalIndent(resp, "", "  ")
	log.Println(string(b))
}
