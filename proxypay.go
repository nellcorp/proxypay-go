package proxypaygo

import (
	"errors"
	"fmt"
	"time"

	"github.com/shopspring/decimal"
)

const (
	sandboxUrl         = "some url"
	productionUrl      = "some url"
	paymentTransaction = "payment"
)

var (
	// errors
	ErrinvalidEnvironment = errors.New("invalid environment")
)

type (
	ProxyPay struct {
		Token       string
		Environment string
	}

	Payment struct {
		ID                    int64
		TransactionId         int64
		TerminalType          string
		TerminalTransactionId int64
		TerminalLocation      any
		TerminalId            int64
		ReferenceID           int64
		ProductID             string
		PeriodID              int64
		PeriodEndDateTime     time.Time
		ParameterID           string
		Fee                   decimal.Decimal
		EntityID              int64
		PeriodStartDatetime   time.Time
		DateTime              int64
		CustomFields          any
		Amount                decimal.Decimal
	}

	IsusePaymentReferenceParams struct {
		Amount      decimal.Decimal
		EndDateTime time.Time
	}
)

func NewProxyPay(token string, environment string) (proxyPay *ProxyPay, err error) {

	if environment != "development" && environment != "production" {
		err = ErrinvalidEnvironment
		return
	}
	proxyPay = &ProxyPay{
		Token:       token,
		Environment: environment,
	}
	return
}

func (s *ProxyPay) IssuePaymentReference(params IsusePaymentReferenceParams) (transaction Payment, err error) {

	id, err := s.GenerateReferenceID()
	if err != nil {
		return
	}
	url := fmt.Sprintf("%s", sandboxUrl)
	response, responseHeaders, err := httpPost(url,
		map[string]string{
			"Authorization": fmt.Sprintf("Token %s", s.Token),
		}, map[string]interface{}{
			"amount":       params.Amount,
			"end_datetime": params.EndDateTime,
		})

	fmt.Println(response, responseHeaders)

	responseHeaders.Get("status")

	return
}

func (s *ProxyPay) GenerateReferenceID() (referenceID string, err error) {
	return
}
func (s *ProxyPay) DeletePaymentReference() (err error) {
	return
}

func (s *ProxyPay) GetPayment() (payment Payment, err error) {
	return
}

func (s *ProxyPay) AknowledgePayment() (err error) {
	return
}
