package rainbow

import (
	"fmt"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestIsExpiryAvailable(t *testing.T) {
	tests := map[string]struct {
		expiries []time.Time
		expiry   time.Time
		want     bool
	}{
		"expiry-in-list": {
			expiries: []time.Time{
				time.Date(2020, time.January, 1, 0, 0, 0, 0, time.UTC),
				time.Date(2020, time.January, 2, 0, 0, 0, 0, time.UTC),
			},
			expiry: time.Date(2020, time.January, 1, 0, 0, 0, 0, time.UTC),
			want:   true,
		},
		"expiry-not-in-list": {
			expiries: []time.Time{},
			expiry:   time.Date(2020, time.January, 1, 0, 0, 0, 0, time.UTC),
			want:     false,
		},
	}
	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			got := IsExpiryAvailable(tt.expiries, tt.expiry)
			assert.Equal(t, tt.want, got)
		})
	}
}

func TestExpiries(t *testing.T) {
	tests := map[string]struct {
		hour int
		want []time.Time
	}{
		// Problem here because we are testing with the current time.

		// Or you change your code, with the current date as an argument (will use time.Now() for call in our code
		//and wathever we want to test).
	}
	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			got := Expiries(tt.hour)
			for _, expiry := range got {
				fmt.Println(expiry)
			}
			assert.ElementsMatch(t, tt.want, got)
		})
	}
}

func TestTodayTomorrow(t *testing.T) {
	now := time.Now()
	tests := map[string]struct {
		hour         int
		wantToday    time.Time
		wantTomorrow time.Time
	}{
		"hour-8": {
			hour:         8,
			wantToday:    time.Date(now.Year(), now.Month(), now.Day(), 8, 0, 0, 0, time.UTC),
			wantTomorrow: time.Date(now.Year(), now.Month(), now.Day()+1, 8, 0, 0, 0, time.UTC),
		},
	}
	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			today, tomorrow := TodayTomorrow(tt.hour)
			assert.Equal(t, tt.wantToday, today)
			assert.Equal(t, tt.wantTomorrow, tomorrow)
		})
	}
}

func TestNextNFridays(t *testing.T) {
	tests := map[string]struct {
		t    time.Time
		n    int
		want []time.Time
	}{
		"one-month": {
			t: time.Date(2020, time.January, 1, 0, 0, 0, 0, time.UTC),
			n: 2,
			want: []time.Time{
				time.Date(2020, time.January, 10, 0, 0, 0, 0, time.UTC),
				time.Date(2020, time.January, 17, 0, 0, 0, 0, time.UTC),
			},
		},
		"two-months": {
			t: time.Date(2021, time.January, 20, 0, 0, 0, 0, time.UTC),
			n: 3,
			want: []time.Time{
				time.Date(2021, time.January, 29, 0, 0, 0, 0, time.UTC),
				time.Date(2021, time.February, 05, 0, 0, 0, 0, time.UTC),
				time.Date(2021, time.February, 12, 0, 0, 0, 0, time.UTC),
			},
		},
	}
	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			got := NextNFridays(tt.t, tt.n)
			assert.ElementsMatch(t, tt.want, got)
		})
	}
}

func TestLastFridayOfEachQuarter(t *testing.T) {
	tests := map[string]struct {
		t    time.Time
		hour int
		want []time.Time
	}{
		"ok": {
			t:    time.Date(2020, time.January, 1, 0, 0, 0, 0, time.UTC),
			hour: 8,
			want: []time.Time{
				time.Date(2020, time.March, 27, 8, 0, 0, 0, time.UTC),
				time.Date(2020, time.June, 26, 8, 0, 0, 0, time.UTC),
				time.Date(2020, time.September, 25, 8, 0, 0, 0, time.UTC),
				time.Date(2020, time.December, 25, 8, 0, 0, 0, time.UTC),
			},
		},
	}
	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			got := LastFridayOfEachQuarter(tt.t, tt.hour)
			assert.ElementsMatch(t, tt.want, got)
		})
	}
}
