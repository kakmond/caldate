package caldate

import (
	"fmt"
	"math"
	"strconv"
	"time"
)

type Date struct {
	Date, Month, Year int
}

func UnitWeek(targetDate int) string {
	totalWeek := targetDate / 7
	totalDay := targetDate % 7
	if totalDay > 0 {
		resultHaveDay := strconv.Itoa(totalWeek) + " weeks and " + strconv.Itoa(totalDay) + " days"
		return resultHaveDay
	}
	resultNotHaveDay := strconv.Itoa(totalWeek) + " weeks"
	return resultNotHaveDay
}

func ResultDay(startDate, endDate Date) int {
	startTime := time.Date(startDate.Year, time.Month(startDate.Month), startDate.Date, 0, 0, 0, 0, time.UTC)
	endTime := time.Date(endDate.Year, time.Month(endDate.Month), endDate.Date, 0, 0, 0, 0, time.UTC)
	diff := endTime.Sub(startTime)
	//kkk := time
	return int(diff.Hours()/24) + 1
}

func ConvertToSecond(days int) uint64 {
	return uint64(days * 86400)
}

func ConvertToMin(second uint64) uint64 {
	return second / 60
}

type DetailStruct struct {
	Year, Month, Day int
}

func ResultDetail(startDate, endDate Date) Date {
	yearDiff := endDate.Year - startDate.Year
	monthDiff := endDate.Month - startDate.Month
	dayDiff := endDate.Date - startDate.Date + 1
	return Date{Year: yearDiff, Month: monthDiff, Date: dayDiff}
}

func (a Date) ResultDetailSameYear(endDate Date) bool {
	return a.Year == endDate.Year
}

func diff(a, b time.Time) (year, month, day int) {
	if a.Location() != b.Location() {
		b = b.In(a.Location())
	}
	if a.After(b) {
		a, b = b, a
	}
	y1, M1, d1 := a.Date()
	y2, M2, d2 := b.Date()

	year = int(y2 - y1)
	month = int(M2 - M1)
	day = int(d2 - d1)

	if day < 0 {
		// days in month:
		t := time.Date(y1, M1, 32, 0, 0, 0, 0, time.UTC)
		day += 32 - t.Day()
		month--
	}
	if month < 0 {
		month += 12
		year--
	}
	return
}

func FormatDateConverter(date Date) string {
	dateTime := time.Date(date.Year, time.Month(date.Month), date.Date, 0, 0, 0, 0, time.UTC)
	return fmt.Sprintf("%s, %d %s %d", dateTime.Weekday().String(),
		dateTime.Day(), dateTime.Month().String(), dateTime.Year())
}

func CalPercent(days int) float64 {
	percentile := math.Round(((float64(days)*100)/365)*100) / 100
	return float64(percentile)
}
func toi(s string) int {
	res, _ := strconv.Atoi(s)
	return res
}

func NewDate(date, month, year string) Date {
	return Date{
		Date:  toi(date),
		Month: toi(month),
		Year:  toi(year),
	}
}
