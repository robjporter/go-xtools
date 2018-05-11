package main

import (
	"fmt"
	"strconv"
	"time"

	"../xcarbon"
)

func main() {
	fmt.Println("NOW: ", xcarbon.Now())
	fmt.Println("== SUB ==============================================================")
	fmt.Println("NOW SUB SECOND: ", xcarbon.Now().SubSecond())
	fmt.Println("NOW SUB 2 SECONDS: ", xcarbon.Now().SubSeconds(2))
	fmt.Println("NOW SUB MINUTE: ", xcarbon.Now().SubMinute())
	fmt.Println("NOW SUB 2 MINUTES: ", xcarbon.Now().SubMinutes(2))
	fmt.Println("NOW SUB HOUR: ", xcarbon.Now().SubHour())
	fmt.Println("NOW SUB 2 HOURS: ", xcarbon.Now().SubHours(2))
	fmt.Println("NOW SUB DAY: ", xcarbon.Now().SubDay())
	fmt.Println("NOW SUB 2 DAYS: ", xcarbon.Now().SubDays(2))
	fmt.Println("NOW SUB MONTH: ", xcarbon.Now().SubMonth())
	fmt.Println("NOW SUB 2 MONTHS: ", xcarbon.Now().SubMonths(2))
	fmt.Println("NOW SUB YEAR: ", xcarbon.Now().SubYear())
	fmt.Println("NOW SUB 2 YEARS: ", xcarbon.Now().SubYears(2))
	fmt.Println("== ADD ==============================================================")
	fmt.Println("NOW ADD SECOND: ", xcarbon.Now().AddSecond())
	fmt.Println("NOW ADD 2 SECONDS: ", xcarbon.Now().AddSeconds(2))
	fmt.Println("NOW ADD MINUTE: ", xcarbon.Now().AddMinute())
	fmt.Println("NOW ADD 2 MINUTES: ", xcarbon.Now().AddMinutes(2))
	fmt.Println("NOW ADD HOUR: ", xcarbon.Now().AddHour())
	fmt.Println("NOW ADD 2 HOURS: ", xcarbon.Now().AddHours(2))
	fmt.Println("NOW ADD DAY: ", xcarbon.Now().AddDay())
	fmt.Println("NOW ADD 2 DAYS: ", xcarbon.Now().AddDays(2))
	fmt.Println("NOW ADD MONTH: ", xcarbon.Now().AddMonth())
	fmt.Println("NOW ADD 2 MONTHS: ", xcarbon.Now().AddMonths(2))
	fmt.Println("NOW ADD YEAR: ", xcarbon.Now().AddYear())
	fmt.Println("NOW ADD 2 YEARS: ", xcarbon.Now().AddYears(2))
	fmt.Println("== START ==============================================================")
	fmt.Println("START OF HOUR: ", xcarbon.Now().StartOfHour())
	fmt.Println("START OF DAY: ", xcarbon.Now().StartOfDay())
	fmt.Println("START OF WEEK: ", xcarbon.Now().StartOfWeek())
	fmt.Println("START OF MONTH: ", xcarbon.Now().StartOfMonth())
	fmt.Println("START OF YEAR: ", xcarbon.Now().StartOfYear())
	fmt.Println("== END ==============================================================")
	fmt.Println("END OF HOUR: ", xcarbon.Now().EndOfHour())
	fmt.Println("END OF DAY: ", xcarbon.Now().EndOfDay())
	fmt.Println("END OF WEEK: ", xcarbon.Now().EndOfWeek())
	fmt.Println("END OF MONTH: ", xcarbon.Now().EndOfMonth())
	fmt.Println("END OF YEAR: ", xcarbon.Now().EndOfYear())
	fmt.Println("== PREVIOUS ==============================================================")
	fmt.Println("PREVIOUS MONTH: ", xcarbon.Now().PreviousMonth())
	fmt.Println("PREVIOUS MONTH LAST DAY: ", xcarbon.Now().PreviousMonthLastDay())
	fmt.Println("PREVIOUS MONTH START DAY: ", xcarbon.Now().PreviousMonthStartDay())
	fmt.Println("== INFO ==============================================================")
	fmt.Println("MONTH NAME: ", xcarbon.Now().MonthName())
	fmt.Println("DAY NUMBER: ", xcarbon.Now().DayNumber())
	fmt.Println("MONTH NUMBER: ", xcarbon.Now().MonthNumber())
	fmt.Println("YEAR NUMBER: ", xcarbon.Now().YearNumber())
	fmt.Println("DATE STRING: ", xcarbon.Now().ToDateTimeString())
	fmt.Println("== EXTRAS ==============================================================")
	fmt.Println("TIMESTAMP: ", xcarbon.Now().ToTimeStamp())
	fmt.Println("QUARTER: ", xcarbon.Now().Quarter())
	fmt.Println("TOMORROW: ", xcarbon.Now().Tomorrow())
	fmt.Println("YESTERDAY: ", xcarbon.Now().Yesterday())
	fmt.Println("ORDINAL: ", xcarbon.Now().Ordinal())
	fmt.Println("DAYS IN MONTH: ", xcarbon.Now().DaysInMonth())
	fmt.Println("IS LEAP YEAR: ", xcarbon.Now().IsLeapYear())
	fmt.Println("DAYS LEFT IN WEEK: ", xcarbon.Now().DaysLeftInWeek())
	fmt.Println("DAYS LEFT IN WORK WEEK: ", xcarbon.Now().DaysLeftInWorkWeek())
	fmt.Println("DAYS LEFT IN MONTH: ", xcarbon.Now().DaysLeftInMonth())
	fmt.Println("DAYS LEFT IN YEAR: ", xcarbon.Now().DaysLeftInYear())
	fmt.Println("DAYS TO HOURS: ", xcarbon.Now().DaysToHours(4))
	fmt.Println("IS WEEKDAY: ", xcarbon.Now().IsWeekday())
	fmt.Println("IS WEEKEND: ", xcarbon.Now().IsWeekend())
	fmt.Println("TAX YEAR: ", xcarbon.Now().TaxYear())
	fmt.Println("ADD DECADE: ", xcarbon.Now().AddDecade())
	fmt.Println("ADD DECADES: ", xcarbon.Now().AddDecades(2))
	fmt.Println("SUB DECADE: ", xcarbon.Now().SubDecade())
	fmt.Println("SUB DECADES: ", xcarbon.Now().SubDecades(2))
	fmt.Println("NEXT LEAPYEAR: ", xcarbon.Now().NextLeapYear())
	fmt.Println("DAYS TO SPRING: ", xcarbon.Now().DaysToSpring())
	fmt.Println("DAYS TO SUMMER: ", xcarbon.Now().DaysToSummer())
	fmt.Println("DAYS TO AUTUMN: ", xcarbon.Now().DaysToAutumn())
	fmt.Println("DAYS TO WINTER: ", xcarbon.Now().DaysToWinter())
	fmt.Println("DAYS TO CHRISTMAS: ", xcarbon.Now().DaysToChristmas())
	fmt.Println("IS SPRING: ", xcarbon.Now().IsSpring())
	fmt.Println("IS SUMMER: ", xcarbon.Now().IsSummer())
	fmt.Println("IS AUTUMN: ", xcarbon.Now().IsAutumn())
	fmt.Println("IS WINTER: ", xcarbon.Now().IsWinter())
	fmt.Println("SEASON: ", xcarbon.Now().Season())
	fmt.Println("WEEK: ", xcarbon.Now().Week())
	fmt.Println("PREVIOUS DAY: ", xcarbon.Now().PreviousDay(time.Monday))
	fmt.Println("NEXT DAY: ", xcarbon.Now().NextDay(time.Monday))
	fmt.Println("IS MONDAY: ", xcarbon.Now().IsMonday())
	fmt.Println("IS TUESDAY: ", xcarbon.Now().IsTuesday())
	fmt.Println("IS WEDNESDAY: ", xcarbon.Now().IsWednesday())
	fmt.Println("IS THURSDAY: ", xcarbon.Now().IsThursday())
	fmt.Println("IS FRIDAY: ", xcarbon.Now().IsFriday())
	fmt.Println("IS SATURDAY: ", xcarbon.Now().IsSaturday())
	fmt.Println("IS SUNDAY: ", xcarbon.Now().IsSunday())
	fmt.Println("== FORMATS ==============================================================")
	for i := 0; i < 30; i++ {
		fmt.Println("FORMAT "+strconv.Itoa(i)+": ", xcarbon.Now().FormatIt(i))
	}

}
