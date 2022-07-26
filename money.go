package money

var mon1 float64
var mon2 float64

type Money struct {
	rupees int
	paise  int
}

func NewMoney(rupees int, paise int) Money {
	if rupees < 0 || paise < 0 {
		panic("Rupees and paise should be positive ")
	}

	return Money{rupees, paise}
}
func (mon Money) AddMoney(newmoney Money) float64 {

	mon1 = float64(newmoney.rupees*100 + newmoney.paise)
	mon2 = float64(mon.rupees*100 + mon.paise)
	return (mon1 + mon2) / 100

}

func (mon Money) SubtractMoney(newmoney Money) float64 {

	mon1 = float64(newmoney.rupees*100 + newmoney.paise)
	mon2 = float64(mon.rupees*100 + mon.paise)
	return (mon1 - mon2) / 100

}

func (mon Money) Equals(mon2 Money) bool {
	if (mon.rupees == mon2.rupees) && (mon.paise == mon2.paise) {
		return true
	}
	return false
}
