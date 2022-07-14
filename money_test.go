package money

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestMoney(t *testing.T) {
	t.Run("Initialize the Money with rupees and paise", func(t *testing.T) {
		assert.NotPanics(t, func() {
			NewMoney(4, 56)
		})
	})

	t.Run("new money should return the total amount", func(t *testing.T) {
		assert.IsType(t, Money{}, NewMoney(4, 5))
	})
	t.Run("the rupees paise values  must be positive", func(t *testing.T) {
		assert.Panics(t, func() {
			NewMoney(-1, 4)
		})
	})

}

func TestTotalamount(t *testing.T) {
	t.Run("should return 5 if ruppes is 5 and paise is 0", func(t *testing.T) {
		amount := NewMoney(5, 0).TotalAmount()
		assert.Equal(t, 5.0, amount)
	})

	t.Run("should return 0.90 if ruppes is 0 and paise is 90", func(t *testing.T) {
		amount := NewMoney(0, 90).TotalAmount()
		assert.Equal(t, 0.9, amount)
	})

	t.Run("should return 1.90 if ruppes is 1 and paise is 90", func(t *testing.T) {
		amount := NewMoney(1, 90).TotalAmount()
		assert.Equal(t, 1.9, amount)
	})

	t.Run("should return 2.90 if ruppes is 1 and paise is 190", func(t *testing.T) {
		amount := NewMoney(2, 90).TotalAmount()
		assert.Equal(t, 2.9, amount)
	})
}
