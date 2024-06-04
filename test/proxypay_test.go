package test

import (
	"os"
	"testing"
	"time"

	proxypaygo "github.com/nellcorp/proxypay-go"
	"github.com/shopspring/decimal"
	"github.com/subosito/gotenv"
)

func init() {
	gotenv.Load("../test.env")
}
func TestProxyPay(t *testing.T) {

	token := os.Getenv("SANDBOX_TOKEN")
	environment := "development"

	proxyPay, _ := proxypaygo.NewProxyPay(token, environment)

	_, err := proxyPay.IssuePaymentReference(decimal.NewFromFloat(23.2), time.Now().AddDate(0, 0, 2))

	if err != nil {
		t.Errorf(" IssuePaymentReference failed with error %v", err)
	}
}
