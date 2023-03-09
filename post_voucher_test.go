package fortnox_test

import (
	"encoding/json"
	"log"
	"testing"
	"time"

	fortnox "github.com/omniboost/go-gastrofix"
)

func TestPostVoucher(t *testing.T) {
	req := client.NewPostVoucherRequest()
	voucher := req.RequestBody().Voucher
	voucher.Description = "Test"
	voucher.TransactionDate = fortnox.Date{time.Now()}
	rows := fortnox.VoucherRows{
		{
			Account:                1249,
			Credit:                 12.0,
			TransactionInformation: "Row 1",
		},
		{
			Account:                1250,
			Debit:                  12.0,
			TransactionInformation: "Row 2",
		},
	}
	voucher.VoucherRows = append(voucher.VoucherRows, rows...)
	req.RequestBody().Voucher = voucher
	resp, err := req.Do()
	if err != nil {
		t.Error(err)
	}

	b, _ := json.MarshalIndent(resp, "", "  ")
	log.Println(string(b))
}
