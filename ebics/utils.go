package ebics

import (
	"math/rand"
	"time"
	"github.com/rickar/cal"
)

// const lettersAndNumbers = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ1234567890"
const numbers = "0123456789"
const letters = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
const ssigns = "+-?:().,'"

func RestrictedIdentificationSEPA1() string {
	b := make([]byte, 35)
	selected := letters + numbers + ssigns
	for i := range b {
		b[i] = selected[rand.Int63()%int64(len(selected))]
	}
	return string(b)
}

// PossibleFillingTime calculates the next possible filling time at ecb
func PossibleFillingTime(t time.Time) time.Time {
	// create new calendar
	c := cal.NewCalendar()

	// add german holidays
	cal.AddGermanHolidays(c)

	// add ecb target2 holidays
	cal.AddEcbHolidays(c)

	// time of day before 7:00 (before ecb working time)
	if t.Hour() < 7 {
		t = time.Date(t.Year(), t.Month(), t.Day(), 7, 0, 0, 0, t.Location())
	}

	// time of day after 18:00 (after ecb working time)
	if t.Hour() >= 18 {
		t = time.Date(t.Year(), t.Month(), t.Day(), 7, 0, 0, 0, t.Location())
		t = t.Add(time.Hour * 24)
	}

	// skip no german or target2 working days
	for !c.IsWorkday(t) {
		t = time.Date(t.Year(), t.Month(), t.Day(), 7, 0, 0, 0, t.Location())
		t = t.Add(time.Hour * 24)
	}

	return t
}

// NextTarget2Workday calculates the next target2 workday
func NextTarget2Workday(t time.Time) time.Time {
	// create new calendar
	c := cal.NewCalendar()

	// add ecb target2 holidays
	cal.AddEcbHolidays(c)

	// get next day 7:00
	t = t.Add(time.Hour * 24)
	t = time.Date(t.Year(), t.Month(), t.Day(), 7, 0, 0, 0, t.Location())

	// skip no target2 working days
	for !c.IsWorkday(t) {
		t = t.Add(time.Hour * 24)
	}

	return t
}

// EarliestRequestedCollectionDate calculates earliest possible requested collection date
func EarliestRequestedCollectionDate(t time.Time) time.Time {
	// get next possible filling time at ecn
	t = PossibleFillingTime(t)

	if t.Hour() < 10 || t.Hour() == 10 && t.Minute() < 30 {
		// filling time before 10:30 -> one target2 workday
		t = NextTarget2Workday(t)

	} else if t.Hour() == 10 && t.Minute() >= 30 ||
		t.Hour() > 10 && t.Hour() < 18 || t.Hour() > 10 && t.Hour() == 18 && t.Minute() < 30 {
		// filling time before 18:30 -> two target2 workdays
		t = NextTarget2Workday(NextTarget2Workday(t))
	}

	return t
}
