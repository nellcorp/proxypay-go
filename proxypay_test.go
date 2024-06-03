package proxypaygo

import (
	"errors"
	"testing"
)

func TestProxyPay(t *testing.T) {

	t.Run("baseURL should be equal sandboxURL if environemnt equal development", func(t *testing.T) {

		token := "testtoken"
		environment := "development"

		proxyPay, _ := NewProxyPay(token, environment)

		if proxyPay.baseURL != sandboxUrl {
			t.Errorf("baseURL should be equal sandboxURL")
		}
	})

	t.Run("baseURL should be equal productionURL if environemnt equal production", func(t *testing.T) {

		token := "testtoken"
		environment := "production"

		proxyPay, _ := NewProxyPay(token, environment)

		if proxyPay.baseURL != productionUrl {
			t.Errorf("baseURL should be equal sandboxURL")
		}
	})

	t.Run("NewProxyPay should return an invalid environment error if environment is not equal to either production or development", func(t *testing.T) {
		token := "testtoken"
		invalid_environment := "some_invalid_environment"

		_, err := NewProxyPay(token, invalid_environment)

		if errors.Is(err, ErrinvalidEnvironment) {
			t.Errorf("baseURL should be equal sandboxURL")
		}
	})
}
