package domain_test

import (
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/vinigracindo/tdd-golang/ports/domain"
)

func TestClient_IsValid(t *testing.T) {
	client := domain.Client{}

	t.Run("IsValid test", func(t *testing.T) {
		client.Name = "Foo"
		client.Email = "foo@bar.com"
		client.Phone = "9999999999"

		_, err := client.IsValid()
		require.Nil(t, err)
	})

	t.Run("Invalid email", func(t *testing.T) {
		client.Email = "foobar"
		_, err := client.IsValid()
		require.Error(t, err)
	})

	t.Run("Invalid phone", func(t *testing.T) {
		client.Email = "foo@bar.com"
		client.Phone = "foorbar"
		_, err := client.IsValid()
		require.Error(t, err)
	})
}
