package money

var rup float64
var pai float64

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

func (mon Money) AddMoney() float64 {
	if mon.paise < 100 {
		var temp float64 = (float64)(mon.paise) / 100
		return mon.rupees + temp
	} else if mon.paise > 100 {
		rup = (float64)(mon.paise / 100)
		pai = (float64)(mon.paise % 100)
		return (mon.rupees + rup) + (pai / 100)
	}
	return 5.0
}

func (mon Money) Equals(mon2 Money) bool {
	if (mon.rupees == mon2.rupees) && (mon.paise == mon2.paise) {
		return true
	}
	return false
}
