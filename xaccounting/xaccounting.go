package xaccounting

import (
	"errors"
	"strings"
	"math"
	"fmt"
	"strconv"
	"bytes"
)

type Xaccounting struct{
	symbol string
	precision int
	thousand string
	decimal string
	money float64
	format string
	nformat string
	roundon float64
}

func New() *Xaccounting {
	return NewFromFloat(0)
}

func NewFromFloat(y float64) *Xaccounting {
	return &Xaccounting{
		money: y,
		thousand: ",",
		decimal: ".",
		precision: 0,
		format: "%s%v",
		nformat: "%s(%v)",
		roundon: 0.5,
	}
}

func NewFromString(y string) *Xaccounting {
	f, err := strconv.ParseFloat(y, 64)
	if err != nil {
		f = 0
	}
	return &Xaccounting{
		money: f,
		thousand: ",",
		decimal: ".",
		precision: 0,
		format: "%s%v",
		nformat: "%s (%v)",
		roundon: 0.5,
	}
}


func (x *Xaccounting) Allocate(y ...int) *Xaccounting {
	return x
}

func (x *Xaccounting) Absolute(y Xaccounting) *Xaccounting {
	return x
}





func (x *Xaccounting) Negative() *Xaccounting {
	x.money = math.Abs(x.money) * (-1)
	return x
}

func (x *Xaccounting) Positive() *Xaccounting {
	x.money = math.Abs(x.money)
	return x
}

func (x *Xaccounting) Reset() *Xaccounting {
	x.thousand = ","
	x.precision = 0
	x.decimal = "."
	x.format = "%s%v"
	x.nformat="%s (%v)"
	return x
}

func (x *Xaccounting) Format(y string) (*Xaccounting,error) {
	var err error
	if strings.Contains(y,"%s") {
		if strings.Contains(y,"%v") {
			x.format = y
		} else {
			err = errors.New("The format string needs to contain '%v' to display the value.")
		}
	} else {
		err = errors.New("The format string needs to contain '%s' to display the symbol.")
	}
	return x,err
}

func (x *Xaccounting) NegativeFormat(y string) (*Xaccounting,error) {
	var err error
	if strings.Contains(y,"%s") {
		if strings.Contains(y,"%v") {
			x.nformat = y
		} else {
			err = errors.New("The format string needs to contain '%v' to display the value.")
		}
	} else {
		err = errors.New("The format string needs to contain '%s' to display the symbol.")
	}
	return x,err
}

func (x *Xaccounting) Symbol(y string) (*Xaccounting,error) {
	var err error
	if len(y) < 4 {
		x.symbol = y
	} else {
		err = errors.New("Only single string symbols are supported.")
	}
	return x,err
}

func (x *Xaccounting) Thousand(y string) (*Xaccounting,error) {
	var err error
	if len(y) == 1 {
		x.thousand = y
	} else {
		err = errors.New("Only single string thousand separators are supported.")
	}
	return x,err
}

func (x *Xaccounting) Decimal(y string) (*Xaccounting,error) {
	var err error
	if len(y) == 1 {
		x.decimal = y
	} else {
		err = errors.New("Only single string decimal separators are supported.")
	}
	return x,err
}

func (x *Xaccounting) Precision(y int) *Xaccounting {
	x.precision = y
	return x
}

func (x *Xaccounting) Money(y float64) (string,error) {
	x.money = y
	round := fmt.Sprintf(fmt.Sprintf("%%%s%df", x.decimal,x.precision), y)
	splits := strings.Split(round,x.decimal)
	round = thousands(splits[0],x.thousand)

	if len(splits) == 2 {
		round += x.decimal + splits[1]
	}

	if y < 0 {
		round = fmt.Sprintf(x.nformat,x.symbol,round)
	} else {
		round = fmt.Sprintf(x.format,x.symbol,round)
	}

	return round,nil
}

func (x *Xaccounting) String() string {
	a,b:=x.Money(x.money)
	if b == nil {
		return a
	}
	return ""
}

func (x *Xaccounting) IsZero() bool {
	if x.money == 0 {
		return true
	}
	return false
}

func (x *Xaccounting) IsPositive() bool {
	return !x.IsNegative()
}

func (x *Xaccounting) IsNegative() bool {
	if x.money < 0 {
		return true
	}
	return false
}

func (x *Xaccounting) Tax(taxpercent float64) string {
	tmp,_ := New().Symbol(x.symbol)
	tmp.precision = x.precision
	if !x.IsZero() {
		a,e := tmp.Money(x.getTax(taxpercent))
		if e == nil {
			return a
		}
	}
	return ""
}

func (x *Xaccounting) getTax(taxpercent float64) float64 {
	a := x.money * (taxpercent /100)
	return a
}

func (x *Xaccounting) PreTax(taxpercent float64) string {
	tax := x.getTax(taxpercent)
	if tax > 0 {
		tmp,_ := New().Symbol(x.symbol)
		tmp.precision = x.precision
		a,e := tmp.Money(x.money - x.getTax(taxpercent))
		if e == nil {
			return a
		}
	}
	return ""
}

func (x *Xaccounting) Split(number int) string {
	if number < 101 {
		tmp := New()
		tmp.precision = x.precision
		tmp.symbol = x.symbol
		res,e := tmp.Money(x.round(x.money/float64(number)))
		if e == nil {
			return res
		}
	}
	return ""
}

func (x *Xaccounting) Add(y *Xaccounting) *Xaccounting {
	x.money += y.money
	return x
}

func (x *Xaccounting) Sub(y *Xaccounting) *Xaccounting {
	x.money -= y.money
	return x
}

func (x *Xaccounting) Divide(y int) *Xaccounting {
	x.money = x.money / float64(y)
	return x
}

func (x *Xaccounting) Multiply(y int) *Xaccounting {
	x.money = x.money * float64(y)
	return x
}

func (x *Xaccounting) Equals(y *Xaccounting) bool {
	if x.money == y.money {
		return true
	}
	return false
}

func (x *Xaccounting) GreaterThan(y *Xaccounting) bool {
	if x.money > y.money {
		return true
	}
	return false
}

func (x *Xaccounting) GreaterThanEqual(y *Xaccounting) bool {
	if x.money >= y.money {
		return true
	}
	return false
}

func (x *Xaccounting) LessThan(y *Xaccounting) bool {
	if x.money < y.money {
		return true
	}
	return false
}

func (x *Xaccounting) LessThaEqual(y *Xaccounting) bool {
	if x.money <= y.money {
		return true
	}
	return false
}

func (x *Xaccounting) round(val float64) (newVal float64) {
	var round float64
	pow := math.Pow(10, float64(x.precision))
	digit := pow * val
	_, div := math.Modf(digit)
	if div >= x.roundon {
		round = math.Ceil(digit)
	} else {
		round = math.Floor(digit)
	}
	newVal = round / pow
	return
}

func thousands(num string, sep string) string {
	ret := ""
	num = reverse(num)
	tmp := splitSub(num,3)
	for i:=0;i<len(tmp);i++{
		ret += tmp[i]
		if i < len(tmp)-1 {
			ret += sep
		}
	}
	ret = reverse(ret)
	return ret
}

func splitSub(s string, n int) []string {
	sub := ""
	subs := []string{}

	runes := bytes.Runes([]byte(s))
	l := len(runes)
	for i, r := range runes {
		sub = sub + string(r)
		if (i + 1) % n == 0 {
			subs = append(subs, sub)
			sub = ""
		} else if (i + 1) == l {
			subs = append(subs, sub)
		}
	}

	return subs
}

func reverse(s string) (result string) {
	for _,v := range s {
		result = string(v) + result
	}
	return
}