package rainbow

import (
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
		day  time.Time
		hour int
		want []time.Time
	}{
		"ok": {
			day:  time.Date(2020, time.January, 1, 0, 0, 0, 0, time.UTC),
			hour: 8,
			want: []time.Time{
				time.Date(2020, time.January, 1, 8, 0, 0, 0, time.UTC),
				time.Date(2020, time.January, 2, 8, 0, 0, 0, time.UTC),
				time.Date(2020, time.January, 10, 8, 0, 0, 0, time.UTC),
				time.Date(2020, time.January, 17, 8, 0, 0, 0, time.UTC),
				time.Date(2020, time.January, 24, 8, 0, 0, 0, time.UTC),
				time.Date(2020, time.March, 27, 8, 0, 0, 0, time.UTC),
				time.Date(2020, time.June, 26, 8, 0, 0, 0, time.UTC),
				time.Date(2020, time.September, 25, 8, 0, 0, 0, time.UTC),
				time.Date(2020, time.December, 25, 8, 0, 0, 0, time.UTC),
			},
		},
	}
	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			got := Expiries(tt.day, tt.hour)
			assert.ElementsMatch(t, tt.want, got)
		})
	}
}

func TestNextNFridays(t *testing.T) {
	t.Parallel()

	cases := []struct {
		name string
		t    time.Time
		n    int
		want []time.Time
	}{{
		name: "one-month",
		t:    time.Date(2020, time.January, 1, 0, 0, 0, 0, time.UTC),
		n:    2,
		want: []time.Time{
			time.Date(2020, time.January, 10, 0, 0, 0, 0, time.UTC),
			time.Date(2020, time.January, 17, 0, 0, 0, 0, time.UTC),
		},
	}, {
		name: "two-months",
		t:    time.Date(2021, time.January, 20, 0, 0, 0, 0, time.UTC),
		n:    3,
		want: []time.Time{
			time.Date(2021, time.January, 29, 0, 0, 0, 0, time.UTC),
			time.Date(2021, time.February, 5, 0, 0, 0, 0, time.UTC),
			time.Date(2021, time.February, 12, 0, 0, 0, 0, time.UTC),
		},
	}}

	for _, c := range cases {
		c := c // required for parallel test

		t.Run(c.name, func(t *testing.T) {
			t.Parallel()
			got := NextNFridays(c.t, c.n)
			assert.ElementsMatch(t, c.want, got)
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
