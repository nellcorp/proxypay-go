package proxypaygo

import (
	"errors"
	"fmt"
	"time"

	"github.com/shopspring/decimal"
)

const (
	sandboxUrl            = "https://api.sandbox.proxypay.co.ao"
	productionUrl         = "https://api.proxypay.co.ao"
	paymentTransaction    = "payment"
	acceptResponsePayload = "application/vnd.proxypay.v2+json"
)

var (
	// errors
	ErrinvalidEnvironment = errors.New("invalid environment")
	ENV_URL               = map[string]string{
		"production":  productionUrl,
		"development": sandboxUrl,
	}
)

type (
	ProxyPay struct {
		Token       string
		Environment string
		baseURL     string
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
		baseURL:     ENV_URL[environment],
	}
	return
}

func (s *ProxyPay) IssuePaymentReference(amount decimal.Decimal, endDatetime time.Time) (referenceID int64, err error) {

	referenceID = generateNineDigitNumber()
	url := fmt.Sprintf("%s/references/%d", s.baseURL, referenceID)

	request := map[string]interface{}{
		"amount":       amount,
		"end_datetime": endDatetime,
	}
	_, _, err = httpPut(url,
		map[string]string{
			"Authorization": fmt.Sprintf("Token %s", s.Token),
			"Accept":        acceptResponsePayload,
		}, request)

	if err != nil {
		return
	}
	return
}

func (s *ProxyPay) UpdatePaymentReference(amount decimal.Decimal, referenceID int64, endDatetime time.Time) (err error) {
	url := fmt.Sprintf("%s/references/%d", s.baseURL, referenceID)

	request := map[string]interface{}{
		"amount":       amount,
		"end_datetime": endDatetime,
	}
	_, _, err = httpPut(url,
		map[string]string{
			"Authorization": fmt.Sprintf("Token %s", s.Token),
			"Accept":        acceptResponsePayload,
		}, request)

	if err != nil {
		return
	}
	return
}

func (s *ProxyPay) DeletePaymentReference(referenceID int64) (err error) {

	url := fmt.Sprintf("%s/references", s.baseURL)
	_, _, err = httpPost(url, map[string]string{
		"Authorization": fmt.Sprintf("Token %s", s.Token),
		"Accept":        acceptResponsePayload,
	}, nil)
	return
}

func (s *ProxyPay) AknowledgePayment(paymentID int64) (err error) {
	url := fmt.Sprintf("%s/payments/%d", sandboxUrl, paymentID)
	_, _, err = httpDelete(url, map[string]string{
		"Authorization": fmt.Sprintf("Token %s", s.Token),
		"Accept":        acceptResponsePayload,
	}, nil)
	return
}
