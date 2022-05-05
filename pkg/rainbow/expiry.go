package rainbow

import (
	"fmt"
	"time"
)

// return today
// tomorrow
// next n fridays

func Expiries(hour int) []time.Time {
	t := time.Now()
	expiries := make([]time.Time, 0, 5)
	//Deribit is 8:00 UTC
	// Delta is 12:00 UTC
	today := time.Date(t.Year(), t.Month(), t.Day(), hour, 0, 0, 0, time.UTC)
	tomorrow := today.Add(24 * time.Hour)
	expiries = append(expiries, today, tomorrow)
	//we start looking for Fridays on the day after tomorrow ;-)
	expiries = append(expiries, NextNFridays(today, 8)...)
	expiries = append(expiries, LastFridayOfEachQuarter(hour)...)
	return expiries

}

// NextNFridays returns the next n Fridays
// t included if it's a Friday
func NextNFridays(t time.Time, fridays int) []time.Time {
	expiries := make([]time.Time, fridays)
	// adding "7" if not it counts the friday to the current week, not the next one
	d := t.Add(time.Duration(7-t.Weekday()+time.Friday) * 24 * time.Hour)
	fmt.Println(t)
	fmt.Println(d)
	for i := 0; i < fridays; i++ {
		expiries[i] = d.Add(time.Duration(i) * 7 * 24 * time.Hour)
	}
	return expiries
}

// function inspired by
// https://www.rosettacode.org/wiki/Last_Friday_of_each_month#Go
func LastFridayOfEachQuarter(hour int) []time.Time {
	now := time.Now()
	y := now.Year()
	expiries := make([]time.Time, 0, 4)

	i := 0
	for m := now.Month(); i <= 12; m++ {
		if m%3 == 0 {
			d := time.Date(y, m+1, 1, hour, 0, 0, 0, time.UTC).Add(-24 * time.Hour)
			d = d.Add(-time.Duration((d.Weekday()+7-time.Friday)%3) * 24 * time.Hour)
			//fmt.Println(d.Format("2006-01-02"))
			expiries = append(expiries, d)
		}
		i++
	}
	return expiries
}
