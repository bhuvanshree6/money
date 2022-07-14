package money

type Money struct {
	rupees float64
	paise  int
}

func NewMoney(rupees float64, paise int) Money {
	if rupees < 0 || paise < 0 {
		panic("Rupees and paise should be positive ")
	}
	return Money{rupees, paise}
}
