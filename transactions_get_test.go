package gastrofix_test

import (
	"encoding/json"
	"log"
	"testing"
)

func TestTransactionsGet(t *testing.T) {
	req := client.NewTransactionsGetRequest()
	req.PathParams().PeriodID = 17
	resp, err := req.Do()
	if err != nil {
		log.Fatal(err)
		t.Error(err)
	}

	b, _ := json.MarshalIndent(resp, "", "  ")
	log.Println(string(b))
}
