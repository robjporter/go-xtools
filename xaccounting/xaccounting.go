package xaccounting

import (
	"errors"
	"strings"
	"math"
	"fmt"
	"strconv"
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
		nformat: "%s (%v)",
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

func (x *Xaccounting) Precision(y int) (*Xaccounting,error) {
	x.precision = y
	return x,nil
}

func (x *Xaccounting) Money(y float64) (string,error) {
	x.money = y
	round := fmt.Sprintf(fmt.Sprintf("%%.%df", x.precision), y)



	
	return "",nil
}

func (x *Xaccounting) Tax(taxpercent float64) int64 {
	return 0
}


func (x *Xaccounting) Split(number int) *Xaccounting {
	return x
}

func (x *Xaccounting) Add(y Xaccounting) *Xaccounting {
	return x
}

func (x *Xaccounting) Subtract(y Xaccounting) *Xaccounting {
	return x
}

func (x *Xaccounting) Divide(y int) *Xaccounting {
	return x
}

func (x *Xaccounting) Multiply(y int) *Xaccounting {
	return x
}

func (x *Xaccounting) Absolute(y Xaccounting) *Xaccounting {
	return x
}

func (x *Xaccounting) Negative() *Xaccounting {
	return x
}

func (x *Xaccounting) Allocate(y ...int) *Xaccounting {
	return x
}

func (x *Xaccounting) String() string {
	return ""
}

func (x *Xaccounting) IsZero() bool {
	return true
}

func (x *Xaccounting) IsPositive() bool {
	return true
}

func (x *Xaccounting) IsNegative() bool {
	return true
}

func (x *Xaccounting) Equals(y Xaccounting) bool {
	return true
}

func (x *Xaccounting) GreaterThan(y Xaccounting) bool {
	return true
}

func (x *Xaccounting) GreaterThanEqual(y Xaccounting) bool {
	return true
}

func (x *Xaccounting) LessThan(y Xaccounting) bool {
	return true
}

func (x *Xaccounting) LessThaEqual(y Xaccounting) bool {
	return true
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