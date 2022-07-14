package money

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestMoney(t *testing.T) {
	t.Run("Initialize the Money with rupees and paise", func(t *testing.T) {
		assert.NotPanics(t, func() {
			NewMoney(4, 5)
		})
	})

	t.Run("new money should return the total amount", func(t *testing.T) {
		assert.IsType(t, Money{}, NewMoney(4, 5))
	})

}
