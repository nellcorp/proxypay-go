package proxypaygo

import (
	"errors"
	"fmt"
	"strconv"
	"time"

	"github.com/shopspring/decimal"
)

const (
	sandboxUrl            = "some url"
	productionUrl         = "some url"
	paymentTransaction    = "payment"
	acceptResponsePayload = "application/vnd.proxypay.v2+json"
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

func (s *ProxyPay) IssuePaymentReference(amount decimal.Decimal, endDatetime time.Time) (referenceID int64, err error) {

	referenceID, err = s.GenerateReferenceID()
	if err != nil {
		return
	}
	url := fmt.Sprintf("%s/references/%d", sandboxUrl, referenceID)

	request := IsusePaymentReferenceParams{
		Amount:      amount,
		EndDateTime: endDatetime,
	}
	_, _, err = httpPost(url,
		map[string]string{
			"Authorization": fmt.Sprintf("Token %s", s.Token),
			"Accept":        acceptResponsePayload,
		}, request)

	if err != nil {
		return
	}
	return
}

func (s *ProxyPay) GenerateReferenceID() (referenceID int64, err error) {
	url := fmt.Sprintf("%s/reference_ids", sandboxUrl)
	data, _, err := httpPost(url, map[string]string{
		"Authorization": fmt.Sprintf("Token %s", s.Token),
		"Accept":        acceptResponsePayload,
	}, nil)
	if err != nil {
		return
	}

	rf, err := strconv.Atoi(string(data))
	if err != nil {
		panic(err)
	}
	referenceID = int64(rf)
	return
}
func (s *ProxyPay) DeletePaymentReference(referenceID string) (err error) {

	url := fmt.Sprintf("%s/refereces", sandboxUrl)
	_, _, err = httpPost(url, map[string]string{
		"Authorization": fmt.Sprintf("Token %s", s.Token),
		"Accept":        acceptResponsePayload,
	}, nil)
	return
}

/*func (s *ProxyPay) AknowledgePayment(paymentID string) (err error) {
	url := fmt.Sprintf("%s/payments/%s", sandboxUrl, paymentID)
	_, _, err = httpPost(url, map[string]string{
		"Authorization": fmt.Sprintf("Token %s", s.Token),
		"Accept":        acceptResponsePayload,
	}, nil)
	return
}*/
