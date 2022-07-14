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

func (mon Money) TotalAmount() float64 {
	if mon.paise < 100 {
		var temp float64 = ((float64)(mon.paise) / 100)
		return (mon.rupees + temp)
	} else if mon.paise > 100 {
		var rup float64 = (float64)(mon.paise / 100)
		var pai float64 = (float64)(mon.paise % 100)
		return ((mon.rupees + rup) + (pai / 100))
	}
	return 0
}
