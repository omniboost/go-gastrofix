package gastrofix_test

import (
	"encoding/json"
	"log"
	"testing"
)

func TestBusinessPeriodsGet(t *testing.T) {
	req := client.NewBusinessPeriodsGetRequest()
	resp, err := req.Do()
	if err != nil {
		t.Error(err)
	}

	b, _ := json.MarshalIndent(resp, "", "  ")
	log.Println(string(b))
}
