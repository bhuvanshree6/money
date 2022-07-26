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

func TestAddMoney(t *testing.T) {
	t.Run("should return 11.0 if 6 rupees 0 paise added to 5 rupees 0 paise", func(t *testing.T) {

		money2 := NewMoney(6, 0)
		assert.Equal(t, 11.0, NewMoney(5, 0).AddMoney(money2))
	})

	t.Run("should return 5.9 if 0 rupees 90 paise added to 5 rupees 0 paise", func(t *testing.T) {
		money2 := NewMoney(0, 90)
		assert.Equal(t, 5.9, NewMoney(5, 0).AddMoney(money2))
	})

	t.Run("should return 6.9 if 1 rupees 90 paise added to 5 rupees 0 paise", func(t *testing.T) {
		money2 := NewMoney(1, 90)
		assert.Equal(t, 6.9, NewMoney(5, 0).AddMoney(money2))
	})

	t.Run("should return 6.9 if 1 rupees 90 paise added to 5 rupees 0 paise", func(t *testing.T) {
		money2 := NewMoney(1, 90)
		assert.Equal(t, 2.8, NewMoney(0, 90).AddMoney(money2))
	})

	t.Run("should return 8.8 if 2 rupees 90 paise added to 5 rupees 90 paise", func(t *testing.T) {
		money2 := NewMoney(2, 90)
		assert.Equal(t, 8.8, NewMoney(5, 90).AddMoney(money2))
	})

	t.Run("should return 10.8 if 2 rupees 190 paise added to 5 rupees 190 paise", func(t *testing.T) {
		money2 := NewMoney(2, 190)
		assert.Equal(t, 10.8, NewMoney(5, 190).AddMoney(money2))
	})

	t.Run("should return 8.8 if 1 rupees 90 paise added to 5 rupees 190 paise", func(t *testing.T) {
		money2 := NewMoney(1, 90)
		assert.Equal(t, 8.8, NewMoney(5, 190).AddMoney(money2))
	})

	t.Run("should return 1.8 if 0 rupees 90 paise added to 0 rupees 90 paise", func(t *testing.T) {
		money2 := NewMoney(0, 90)
		assert.Equal(t, 1.8, NewMoney(0, 90).AddMoney(money2))
	})

	t.Run("should return 5.0 if 5 rupees 0 paise added to 0 rupees 0 paise", func(t *testing.T) {
		money2 := NewMoney(5, 0)
		assert.Equal(t, 5.0, NewMoney(0, 0).AddMoney(money2))
	})

}
func TestSubtractMoney(t *testing.T) {
	t.Run("should return 11.0 if 6 rupees 0 paise subtracted from 5 rupees 0 paise", func(t *testing.T) {

		money2 := NewMoney(6, 0)
		assert.Equal(t, 1.0, NewMoney(5, 0).SubtractMoney(money2))
	})

	t.Run("should return 0.5 if 0 rupees 40 paise subtracted from 0 rupees 90 paise", func(t *testing.T) {

		money2 := NewMoney(0, 90)
		assert.Equal(t, 0.5, NewMoney(0, 40).SubtractMoney(money2))
	})

	t.Run("should return 6.5 if 10 rupees 90 paise subtracted from 4 rupees 40 paise", func(t *testing.T) {

		money2 := NewMoney(10, 90)
		assert.Equal(t, 6.5, NewMoney(4, 40).SubtractMoney(money2))
	})

	t.Run("should return 6.6 if 10 rupees 190 paise subtracted from 4 rupees 140 paise", func(t *testing.T) {

		money2 := NewMoney(10, 190)
		assert.Equal(t, 6.6, NewMoney(4, 130).SubtractMoney(money2))
	})

	t.Run("should return -6.5 if 4 rupees 140 paise subtracted 10 rupees 190 paise", func(t *testing.T) {

		money2 := NewMoney(4, 140)
		assert.Equal(t, -6.5, NewMoney(10, 190).SubtractMoney(money2))
	})

}

func TestMoney_Equals(t *testing.T) {
	t.Run("should return true when comparing two equal money entities", func(t *testing.T) {
		assert.Equal(t, true, NewMoney(1, 0).Equals(NewMoney(1, 0)))
	})

	t.Run("should return true when comparing two equal money entities", func(t *testing.T) {
		assert.Equal(t, true, NewMoney(0, 20).Equals(NewMoney(0, 20)))
	})

	t.Run("should return true when comparing two equal money entities", func(t *testing.T) {
		assert.Equal(t, true, NewMoney(1, 20).Equals(NewMoney(1, 20)))
	})
	t.Run("should return true when comparing two unequal money entities", func(t *testing.T) {
		assert.Equal(t, false, NewMoney(1, 0).Equals(NewMoney(2, 0)))
	})
	t.Run("should return true when comparing two unequal money entities", func(t *testing.T) {
		assert.Equal(t, false, NewMoney(0, 20).Equals(NewMoney(0, 40)))
	})
	t.Run("should return true when comparing two unequal money entities", func(t *testing.T) {
		assert.Equal(t, false, NewMoney(1, 20).Equals(NewMoney(2, 40)))
	})
}
