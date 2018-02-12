package test

import (
	"testing"
	"time"
	"github.com/fino-digital/ebics-lib/ebics"
	"fmt"
)

func TestNextTarget2Workday2018(t *testing.T) {
	// good friday - 1
	goodFriday := time.Date(2018, 3, 29, 7, 0, 0, 0, time.UTC)
	goodFridayNextWorkday := time.Date(2018, 4, 3, 7, 0, 0, 0, time.UTC)
	if ebics.NextTarget2Workday(goodFriday) != goodFridayNextWorkday {
		fmt.Println(ebics.NextTarget2Workday(goodFriday))
		fmt.Println(goodFridayNextWorkday)
		t.Fatal("Good friday don't match")
	}

	// easter monday - 1
	easterMonday := time.Date(2018, 4, 1, 7, 0, 0, 0, time.UTC)
	easterMondayNextWorkday := time.Date(2018, 4, 3, 7, 0, 0, 0, time.UTC)
	if ebics.NextTarget2Workday(easterMonday) != easterMondayNextWorkday {
		t.Fatal("Easter monday don't match")
	}

	// new years day - 1
	newYearsDay := time.Date(2017, 12, 31, 7, 0, 0, 0, time.UTC)
	newYearsDayNextWorkday := time.Date(2018, 1, 2, 7, 0, 0, 0, time.UTC)
	if ebics.NextTarget2Workday(newYearsDay) != newYearsDayNextWorkday {
		t.Fatal("New years day don't match")
	}

	// labour day - 1
	labourDay := time.Date(2018, 4, 30, 7, 0, 0, 0, time.UTC)
	labourDayNextWorkday := time.Date(2018, 5, 2, 7, 0, 0, 0, time.UTC)
	if ebics.NextTarget2Workday(labourDay) != labourDayNextWorkday {
		t.Fatal("Labour day don't match")
	}

	// chrismas day - 1
	christmasDay := time.Date(2018, 12, 24, 7, 0, 0, 0, time.UTC)
	christmasDayNextWorkday := time.Date(2018, 12, 27, 7, 0, 0, 0, time.UTC)
	if ebics.NextTarget2Workday(christmasDay) != christmasDayNextWorkday {
		t.Fatal("Christmas day don't match")
	}

	// chrismas holiday - 1
	christmasHoliday := time.Date(2018, 12, 25, 7, 0, 0, 0, time.UTC)
	christmasHolidayNextWorkday := time.Date(2018, 12, 27, 7, 0, 0, 0, time.UTC)
	if ebics.NextTarget2Workday(christmasHoliday) != christmasHolidayNextWorkday {
		t.Fatal("Christmas holiday don't match")
	}

	// saturday - 1 = friday
	saturday := time.Date(2018, 2, 9, 7, 0, 0, 0, time.UTC)
	saturdayNextWorkday := time.Date(2018, 2, 12, 7, 0, 0, 0, time.UTC)
	if ebics.NextTarget2Workday(saturday) != saturdayNextWorkday {
		t.Fatal("Saturday don't match")
	}

	// sunday - 1 = saturday
	sunday := time.Date(2018, 2, 9, 7, 0, 0, 0, time.UTC)
	sundayNextWorkday := time.Date(2018, 2, 12, 7, 0, 0, 0, time.UTC)
	if ebics.NextTarget2Workday(sunday) != sundayNextWorkday {
		t.Fatal("Sunday don't match")
	}

	// workday
	workday := time.Date(2018, 2, 8, 7, 0, 0, 0, time.UTC)
	workdayNextWorkday := time.Date(2018, 2, 9, 7, 0, 0, 0, time.UTC)
	if ebics.NextTarget2Workday(workday) != workdayNextWorkday {
		t.Fatal("Workday don't match")
	}
}

func TestNextTarget2Workday2019(t *testing.T) {
	// good friday - 1
	goodFriday := time.Date(2019, 4, 18, 7, 0, 0, 0, time.UTC)
	goodFridayNextWorkday := time.Date(2019, 4, 23, 7, 0, 0, 0, time.UTC)
	if ebics.NextTarget2Workday(goodFriday) != goodFridayNextWorkday {
		t.Fatal("Good friday don't match")
	}

	// easter monday - 1
	easterMonday := time.Date(2018, 4, 21, 7, 0, 0, 0, time.UTC)
	easterMondayNextWorkday := time.Date(2018, 4, 23, 7, 0, 0, 0, time.UTC)
	if ebics.NextTarget2Workday(easterMonday) != easterMondayNextWorkday {
		t.Fatal("Easter monday day don't match")
	}

	// new years day - 1
	newYearsDay := time.Date(2018, 12, 31, 7, 0, 0, 0, time.UTC)
	newYearsDayNextWorkday := time.Date(2019, 1, 2, 7, 0, 0, 0, time.UTC)
	if ebics.NextTarget2Workday(newYearsDay) != newYearsDayNextWorkday {
		t.Fatal("New years day don't match")
	}

	// labour day - 1
	labourDay := time.Date(2019, 4, 30, 7, 0, 0, 0, time.UTC)
	labourDayNextWorkday := time.Date(2019, 5, 2, 7, 0, 0, 0, time.UTC)
	if ebics.NextTarget2Workday(labourDay) != labourDayNextWorkday {
		t.Fatal("Labour day don't match")
	}

	// chrismas day - 1
	christmasDay := time.Date(2019, 12, 24, 7, 0, 0, 0, time.UTC)
	christmasDayNextWorkday := time.Date(2019, 12, 27, 7, 0, 0, 0, time.UTC)
	if ebics.NextTarget2Workday(christmasDay) != christmasDayNextWorkday {
		t.Fatal("Christmas day don't match")
	}

	// chrismas holiday - 1
	christmasHoliday := time.Date(2019, 12, 25, 7, 0, 0, 0, time.UTC)
	christmasHolidayNextWorkday := time.Date(2019, 12, 27, 7, 0, 0, 0, time.UTC)
	if ebics.NextTarget2Workday(christmasHoliday) != christmasHolidayNextWorkday {
		t.Fatal("Christmas holiday don't match")
	}

	// saturday - 1 = friday
	saturday := time.Date(2019, 2, 8, 7, 0, 0, 0, time.UTC)
	saturdayNextWorkday := time.Date(2019, 2, 11, 7, 0, 0, 0, time.UTC)
	if ebics.NextTarget2Workday(saturday) != saturdayNextWorkday {
		t.Fatal("Saturday don't match")
	}

	// sunday - 1 = saturday
	sunday := time.Date(2019, 2, 9, 7, 0, 0, 0, time.UTC)
	sundayNextWorkday := time.Date(2019, 2, 11, 7, 0, 0, 0, time.UTC)
	if ebics.NextTarget2Workday(sunday) != sundayNextWorkday {
		t.Fatal("Sunday don't match")
	}

	// workday
	workday := time.Date(2019, 2, 7, 7, 0, 0, 0, time.UTC)
	workdayNextWorkday := time.Date(2019, 2, 8, 7, 0, 0, 0, time.UTC)
	if ebics.NextTarget2Workday(workday) != workdayNextWorkday {
		t.Fatal("Workday don't match")
	}
}
