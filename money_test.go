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

}
