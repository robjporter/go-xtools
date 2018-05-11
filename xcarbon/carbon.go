package xcarbon

import (
	"math"
	"reflect"
	"strconv"
	"strings"
	"time"
)

const (
	DATE_TIME_LAYOUT = "2006-01-02 15:04:05"
)

// form @github.com/jinzhu/now and more
var TimeFormats = []string{
	DATE_TIME_LAYOUT,
	"1/2/2006", "1/2/2006 15:4:5", "2006-1-2 15:4:5", "2006-1-2 15:4", "2006-1-2", "1-2", "15:4:5", "15:4",
	"2013-02-03", "15:4:5 Jan 2, 2006 MST", "15:04:05", "2013/02/03",
	"2013-02-03 19:54:00 PST",
	time.RFC822, time.RubyDate, time.RFC822Z, time.RFC3339,
}

type Carbon struct {
	time.Time
}

// return Carbon with time Now
func Now() *Carbon {
	c := &Carbon{time.Now()}
	return c
}

// return Carbon with time Now
func CreateFromTime(timeObject time.Time) *Carbon {
	c := &Carbon{timeObject}
	return c
}

// create Carbon with year, month, day, hour, minute, second, nanosecond, int,  TZ string
// Create(2001, 12, 01, 15, 25, 55, 0, "UTC") = 2001-12-01 13:25:55 +0000 UTC
func Create(year, month, day int, args ...interface{}) *Carbon {
	var tz string
	Month := time.Month(month)
	hour := time.Now().Hour()
	minute := time.Now().Minute()
	second := time.Now().Second()
	nanosecond := time.Now().Nanosecond()
	location := time.Now().Location()
	for key, val := range args {
		if key == 0 {
			hour = val.(int)
		}
		if key == 1 {
			minute = val.(int)
		}
		if key == 2 {
			second = val.(int)
		}
		if key == 3 {
			nanosecond = val.(int)
		}
		if key == 4 {
			tz = val.(string)
		}
	}
	c := &Carbon{time.Date(year, Month, day, hour, minute, second, nanosecond, location)}
	if tz != "" {
		c.SetTZ(tz)
	}
	return c
}

// create Carbon from time string, return in UTC, if tz not specified
// CreateFrom("2015-01-25 15:4:5") = 2015-01-25 15:04:55 +0000 UTC
func CreateFrom(stringDate string) (*Carbon, error) {
	var carbon = Now()
	var err error
	var tm time.Time
	for _, format := range TimeFormats {
		tm, err = time.Parse(format, stringDate)
		if err == nil {
			carbon.Time = tm
			break
		}
	}

	return carbon, err
}

// custom Unmarshal func form many times formats
func (t *Carbon) UnmarshalJSON(buf []byte) error {
	tt, err := CreateFrom(strings.Trim(string(buf), `"`))
	if err != nil {
		return err
	}
	t.Time = tt.Time
	return nil
}

// set tx location, ex. "UTC"
func (c *Carbon) SetTZ(tz string) *Carbon {
	location, _ := time.LoadLocation(tz)
	if location != nil {
		c.Time = c.Time.In(location)
	}
	return c
}

// set time, ex. "23.59.29"
func (c *Carbon) SetTime(hours, minutes, seconds int) *Carbon {
	c.Time = time.Date(
		c.Year(),
		c.Month(),
		c.Day(),
		hours,
		minutes,
		seconds,
		c.Nanosecond(),
		c.Location(),
	)

	return c
}

func (c *Carbon) SubDay() *Carbon {
	return c.SubDays(1)
}

func (c *Carbon) SubDays(days int) *Carbon {
	c.Time = c.Time.AddDate(0, 0, -days)
	return c
}

func (c *Carbon) SubMonth() *Carbon {
	return c.SubMonths(1)
}

func (c *Carbon) SubMonths(months int) *Carbon {
	c.Time = c.Time.AddDate(0, -months, 0)
	return c
}

func (c *Carbon) SubYear() *Carbon {
	return c.SubYears(1)
}

func (c *Carbon) SubYears(years int) *Carbon {
	c.Time = c.Time.AddDate(-years, 0, 0)
	return c
}

func (c *Carbon) AddDay() *Carbon {
	return c.AddDays(1)
}

func (c *Carbon) AddDays(days int) *Carbon {
	c.Time = c.Time.AddDate(0, 0, days)
	return c
}

func (c *Carbon) AddMonth() *Carbon {
	return c.AddMonths(1)
}

func (c *Carbon) AddMonths(months int) *Carbon {
	c.Time = c.Time.AddDate(0, months, 0)
	return c
}

func (c *Carbon) AddYear() *Carbon {
	return c.AddYears(1)
}

func (c *Carbon) AddYears(years int) *Carbon {
	c.Time = c.Time.AddDate(years, 0, 0)
	return c
}

func (c *Carbon) DiffInSeconds(from *Carbon) int {
	return round(c.Sub(from.Time).Seconds())
}

func (c *Carbon) DiffInMinutes(from *Carbon) int {
	return round(c.Sub(from.Time).Minutes())
}

func (c *Carbon) DiffInHours(from *Carbon) int {
	return round(c.Sub(from.Time).Hours())
}

func (c *Carbon) DiffInDays(from *Carbon) int {
	return c.DiffInHours(from) / 24
}

// Determines if the instance is equal to another
func (c *Carbon) Eq(another *Carbon) bool {
	return c.Equal(another.Time)
}

// Determines if the instance is greater (after) than another
func (c *Carbon) Gt(another *Carbon) bool {
	return c.After(another.Time)
}

// Determines if the instance is less (Before) than another
func (c *Carbon) Lt(another *Carbon) bool {
	return c.Before(another.Time)
}

// Determines if the instance is greater than before and less than after
func (c *Carbon) Between(before, after *Carbon) bool {
	return c.After(before.Time) && c.Before(after.Time)
}

func (c *Carbon) StartOfHour() *Carbon {
	c.Time = c.Truncate(time.Hour)
	return c
}

func (c *Carbon) EndOfHour() *Carbon {
	c.StartOfHour()
	c.Time = c.Add(time.Hour - time.Second)
	return c
}

func (c *Carbon) StartOfDay() *Carbon {
	c.Time = c.StartOfHour().Add(-time.Hour * time.Duration(c.Hour()))
	return c
}

func (c *Carbon) EndOfDay() *Carbon {
	c.Time = c.StartOfDay().Add(time.Hour*time.Duration(24) - time.Second)
	return c
}

func (c *Carbon) StartOfWeek(firstDayOfWeekIsMonday ...bool) *Carbon {
	firstDay := time.Monday
	corrFirstDay := 1
	if len(firstDayOfWeekIsMonday) > 0 {
		if !firstDayOfWeekIsMonday[0] {
			firstDay = time.Sunday
			corrFirstDay = 0
		}
	}
	c.StartOfDay()
	if c.Weekday() != firstDay {
		c.Time = c.Add(-time.Hour * 24 * time.Duration(-corrFirstDay+int(c.Weekday())))
	}
	return c
}

func (c *Carbon) EndOfWeek() *Carbon {
	c.Time = c.StartOfWeek().Add(time.Hour*time.Duration(24*7) - time.Second)
	return c
}

func (c *Carbon) StartOfMonth() *Carbon {
	year := c.Year()
	Month := c.Month()
	location := time.Now().Location()
	c = &Carbon{time.Date(year, Month, 1, 0, 0, 0, 0, location)}

	return c
}

func (c *Carbon) EndOfMonth() *Carbon {
	c.Time = c.StartOfMonth().AddDate(0, 1, 0).Add(-time.Second)
	return c
}

func (c *Carbon) StartOfYear() *Carbon {
	year := c.Year()
	location := time.Now().Location()
	c = &Carbon{time.Date(year, time.Month(1), 1, 0, 0, 0, 0, location)}

	return c
}

func (c *Carbon) EndOfYear() *Carbon {
	c.Time = c.StartOfYear().AddDate(1, 0, 0).Add(-time.Second)
	return c
}

func (c *Carbon) PreviousMonth() *Carbon {
	c.Time = c.StartOfMonth().Add(-time.Second)
	return c
}

func (c *Carbon) NextMonth() *Carbon {
	c = c.AddMonth()
	return c
}

func (c *Carbon) PreviousMonthLastDay() *Carbon {
	c = c.SubMonth().EndOfMonth()
	return c
}

func (c *Carbon) PreviousMonthStartDay() *Carbon {
	c = &Carbon{time.Date(c.StartOfMonth().SubMonth().Year(), c.StartOfMonth().SubMonth().Month(), 1, 0, 0, 0, 0, c.Location())}
	return c
}

func (c *Carbon) MonthName() string {
	return c.Month().String()
}

func (c *Carbon) DayNumber() int {
	return c.Day()
}
func (c *Carbon) MonthNumber() int {
	return int(c.Month())
}
func (c *Carbon) YearNumber() int {
	return c.Year()
}

// return string with DateTime format "2006-01-25 15:04:05"
func (c *Carbon) ToDateTimeString() string {
	return c.Format(DATE_TIME_LAYOUT)
}

// CUSTOM

func round(f float64) int {
	if math.Abs(f) < 0.5 {
		return 0
	}
	return int(f + math.Copysign(0.5, f))
}

func (c *Carbon) ToTimeStamp() int64 {
	return c.Unix()
}

func (c *Carbon) Quarter(FYStartMonth ...time.Month) int {
	var startFY time.Month

	if FYStartMonth == nil {
		startFY = time.January
	} else {
		startFY = FYStartMonth[0]
	}

	currentMonth := c.MonthNumber()

	if exists, _ := inArray(currentMonth, getQuarter(1, startFY)); exists {
		return 1
	}
	startFY += 3
	if exists, _ := inArray(currentMonth, getQuarter(2, startFY)); exists {
		return 2
	}
	startFY += 3
	if exists, _ := inArray(currentMonth, getQuarter(3, startFY)); exists {
		return 3
	}
	startFY += 3
	if exists, _ := inArray(currentMonth, getQuarter(4, startFY)); exists {
		return 4
	}

	return 0
}

func getQuarter(quarter, monthName time.Month) []int {
	month := int(monthName)
	var res []int

	for i := 0; i < 3; i++ {
		if month > 12 {
			month -= 12
		}
		res = append(res, month)
		month++
	}
	return res
}

func inArray(val interface{}, array interface{}) (exists bool, index int) {
	exists = false
	index = -1

	switch reflect.TypeOf(array).Kind() {
	case reflect.Slice:
		s := reflect.ValueOf(array)

		for i := 0; i < s.Len(); i++ {
			if reflect.DeepEqual(val, s.Index(i).Interface()) == true {
				index = i
				exists = true
				return
			}
		}
	}
	return
}

func (c *Carbon) Tomorrow() *Carbon {
	return c.AddDay()
}

func (c *Carbon) Yesterday() *Carbon {
	return c.SubDay()
}

func (c *Carbon) Ordinal() string {
	return strconv.Itoa(c.DayNumber()) + c.OrdinalOnly()
}

func (c *Carbon) OrdinalOnly() string {
	suffix := "th"
	switch c.DayNumber() % 10 {
	case 1:
		if c.DayNumber()%100 != 11 {
			suffix = "st"
		}
	case 2:
		if c.DayNumber()%100 != 12 {
			suffix = "nd"
		}
	case 3:
		if c.DayNumber()%100 != 13 {
			suffix = "rd"
		}
	}

	return suffix
}

func (c *Carbon) DaysInMonth() int {
	days := 31
	switch c.Month() {
	case time.April, time.June, time.September, time.November:
		days = 30
		break
	case time.February:
		days = 28
		if c.IsLeapYear() {
			days = 29
		}
		break
	}

	return days
}

func (c *Carbon) IsLeapYear() bool {
	year := c.Year()
	return year%4 == 0 && (year%100 != 0 || year%400 == 0)
}

func (c *Carbon) DaysLeftInWeek() int {
	if int(c.Weekday()) == 0 {
		return 0
	}
	return 7 - int(c.Weekday())
}

func (c *Carbon) DaysLeftInWorkWeek() int {
	if c.Weekday() == time.Saturday || c.Weekday() == time.Sunday {
		return 5
	}
	return 5 - int(c.Weekday())
}

func (c *Carbon) DaysLeftInMonth() int {
	return c.DaysInMonth() - c.Day()
}

func (c *Carbon) DaysLeftInYear() int {
	if c.IsLeapYear() {
		return 366 - c.YearDay()
	} else {
		return 365 - c.YearDay()
	}
}

func (c *Carbon) DaysToHours(days int) int {
	return days * 24
}

func (c *Carbon) IsWeekday() bool {
	return !c.IsWeekend()
}

func (c *Carbon) IsWeekend() bool {
	if c.Weekday() == 0 || c.Weekday() == 6 {
		return true
	}
	return false
}

func (c *Carbon) TaxYear() string {
	if c.Month() < 4 && c.Day() < 6 {
		return strconv.Itoa(c.Year()-1) + "-" + strconv.Itoa(c.Year())
	} else if c.Month() > 3 && c.Day() > 5 {
		return strconv.Itoa(c.Year()) + "-" + strconv.Itoa(c.Year()+1)
	}
	return ""
}

func (c *Carbon) AddDecade() *Carbon {
	return c.AddYears(10)
}

func (c *Carbon) AddDecades(decades int) *Carbon {
	return c.AddYears(decades * 10)
}

func (c *Carbon) SubDecade() *Carbon {
	return c.SubYears(10)
}

func (c *Carbon) SubDecades(decades int) *Carbon {
	return c.SubYears(decades * 10)
}

func (c *Carbon) NextLeapYear() int {
	year := c.Year()
	for i := 0; i < 6; i++ {
		if internalIsLeapYear(year) {
			return year
		} else {
			year++
		}
	}

	return 0
}

func internalIsLeapYear(year int) bool {
	return year%4 == 0 && (year%100 != 0 || year%400 == 0)
}

func (c *Carbon) DaysToSpring() int {
	// 20th March
	year := c.Year()
	month := 3
	day := 20
	if c.Month() > 3 || c.Month() == 3 && c.Day() > 20 {
		year++
	}
	return c.getDaysDifference(day, month, year)
}

func (c *Carbon) DaysToSummer() int {
	// 21st June
	year := c.Year()
	month := 6
	day := 21
	if c.Month() > 6 || c.Month() == 6 && c.Day() > 21 {
		year++
	}
	return c.getDaysDifference(day, month, year)
}

func (c *Carbon) DaysToAutumn() int {
	// 22nd September
	year := c.Year()
	month := 9
	day := 22
	if c.Month() > 9 || c.Month() == 9 && c.Day() > 22 {
		year++
	}
	return c.getDaysDifference(day, month, year)
}

func (c *Carbon) DaysToWinter() int {
	// 21st December
	// Year before leap year = 22nd
	leapyear := c.NextLeapYear()
	leapyear--
	year := c.Year()
	month := 12
	day := 21
	if c.Month() == 12 && c.Day() > day {
		year++
	}
	if leapyear == year {
		day = 22
	}
	return c.getDaysDifference(day, month, year)
}

func (c *Carbon) DaysToChristmas() int {
	year := c.Year()
	month := 12
	day := 25
	return c.getDaysDifference(day, month, year)
}

func (c *Carbon) getDaysDifference(day, month, year int) int {
	tmp := Create(year, month, day)
	return tmp.DiffInDays(c)
}

func (c *Carbon) IsSpring() bool {
	start := Create(c.Year(), 3, 20, 00, 00, 00)
	end := Create(c.Year(), 6, 20, 23, 59, 59)

	if c.Between(start, end) {
		return true
	}
	return false
}

func (c *Carbon) IsSummer() bool {
	// 21st June
	start := Create(c.Year(), 6, 21, 00, 00, 00)
	end := Create(c.Year(), 9, 21, 23, 59, 59)

	if c.Between(start, end) {
		return true
	}
	return false
}

func (c *Carbon) IsAutumn() bool {
	// 22nd September
	start := Create(c.Year(), 9, 22, 00, 00, 00)
	end := Create(c.Year(), 9, 21, 23, 59, 59)

	if c.Between(start, end) {
		return true
	}
	return false
}

func (c *Carbon) IsWinter() bool {
	// 21st December
	// Year before leap year = 22nd
	start1 := Create(c.Year(), 1, 1, 00, 00, 00)
	end1 := Create(c.Year(), 3, 19, 23, 59, 59)
	start2 := Create(c.Year(), 12, 21, 00, 00, 00)
	end2 := Create(c.Year(), 12, 31, 23, 59, 59)

	if c.Between(start1, end1) || c.Between(start2, end2) {
		return true
	}
	return false
}

func (c *Carbon) Season() string {
	if c.IsSpring() {
		return "Spring"
	} else if c.IsSummer() {
		return "Summer"
	} else if c.IsAutumn() {
		return "Autumn"
	} else {
		return "Winter"
	}
}

func (c *Carbon) Copy() *Carbon {
	t2 := *c
	return &t2
}

func (c *Carbon) IsFuture(t2 *Carbon) bool {
	return t2.Gt(c)
}

func (c *Carbon) IsPast(t2 *Carbon) bool {
	return t2.Lt(c)
}

func (c *Carbon) Week() int {
	_, b := c.ISOWeek()
	return b
}

func (c *Carbon) PreviousDay(day time.Weekday) *Carbon {
	for i := 0; i < 8; i++ {
		c.SubDay()
		if c.Weekday() == day {
			return c
		}
	}

	return c
}

func (c *Carbon) NextDay(day time.Weekday) *Carbon {
	for i := 0; i < 8; i++ {
		c.AddDay()
		if c.Weekday() == day {
			return c
		}
	}

	return c
}

func (c *Carbon) FormatIt(format int) string {
	return c.Format(getFormat(format))
}

func getFormat(format int) string {
	switch format {
	case 1:
		return "02/01/06 03:04:05 PM Jan"
	case 2:
		return "02/01/2006 03:04:05 PM Jan"
	case 3:
		return "02/Jan/2006 03:04:05 PM"
	case 4:
		return "02/Jan/2006 15:04:05"
	case 5:
		return "02/01/06 03:04:05 PM Mon Jan"
	case 6:
		return "02/01/06 03:04:05 PM Monay January"
	case 7:
		return "02/01/06 03:04:05 PM Jan"
	case 8:
		return "2/1/06 3:4:05 PM"
	case 9:
		return "_2/1/06 3:4:05 PM"
	case 10:
		return "02/01/06 03:04:05 PM"
	case 11:
		return "02/01/2006 03:04:05 PM"
	case 12:
		return "02/01/2006 03:04:05.000 PM"
	case 13:
		return "02/01/2006 03:04:05.000000 PM"
	case 14:
		return "02/01/2006 03:04:05.000000000 PM"
	case 15:
		return "02/01/2006 15:04:05 MST"
	case 16:
		return "02/01/2006 15:04:05 Z7"
	case 17:
		return "02/01/2006 15:04:05 Z07"
	case 18:
		return "02/01/2006 15:04:05 Z0700"
	case 19:
		return "02/01/2006 15:04:05 Z07:00"
	case 20:
		return "02/01/2006 15:04:05 -07:00"
	case 21:
		return time.RFC822
	case 22:
		return time.RFC1123
	case 23:
		return time.RFC1123Z
	case 24:
		return time.RFC3339
	case 25:
		return time.RFC3339Nano
	case 26:
		return time.RFC822Z
	case 27:
		return time.RFC850
	}
	return "02/01/06 03:04:05 PM Jan"
}

func (c *Carbon) IsMonday() bool {
	return c.IsDay(time.Monday)
}

func (c *Carbon) IsTuesday() bool {
	return c.IsDay(time.Tuesday)
}

func (c *Carbon) IsWednesday() bool {
	return c.IsDay(time.Wednesday)
}

func (c *Carbon) IsThursday() bool {
	return c.IsDay(time.Thursday)
}

func (c *Carbon) IsFriday() bool {
	return c.IsDay(time.Friday)
}

func (c *Carbon) IsSaturday() bool {
	return c.IsDay(time.Saturday)
}

func (c *Carbon) IsSunday() bool {
	return c.IsDay(time.Sunday)
}

func (c *Carbon) IsDay(day time.Weekday) bool {
	if c.Weekday() == day {
		return true
	}
	return false
}

func (c *Carbon) AddHour() *Carbon {
	return c.AddHours(1)
}

func (c *Carbon) AddHours(hours int) *Carbon {
	c.Time = c.Time.Add(time.Duration(hours) * time.Hour)
	return c
}

func (c *Carbon) SubHour() *Carbon {
	return c.SubHours(1)
}

func (c *Carbon) SubHours(hours int) *Carbon {
	c.Time = c.Time.Add(time.Duration(-hours) * time.Hour)
	return c
}

func (c *Carbon) AddMinute() *Carbon {
	return c.AddMinutes(1)
}

func (c *Carbon) AddMinutes(minutes int) *Carbon {
	c.Time = c.Time.Add(time.Duration(minutes) * time.Minute)
	return c
}

func (c *Carbon) SubMinute() *Carbon {
	return c.SubMinutes(1)
}

func (c *Carbon) SubMinutes(minutes int) *Carbon {
	c.Time = c.Time.Add(time.Duration(-minutes) * time.Minute)
	return c
}

func (c *Carbon) AddSecond() *Carbon {
	return c.AddSeconds(1)
}

func (c *Carbon) AddSeconds(seconds int) *Carbon {
	c.Time = c.Time.Add(time.Duration(seconds) * time.Second)
	return c
}

func (c *Carbon) SubSecond() *Carbon {
	return c.SubSeconds(1)
}

func (c *Carbon) SubSeconds(seconds int) *Carbon {
	c.Time = c.Time.Add(time.Duration(-seconds) * time.Second)
	return c
}
