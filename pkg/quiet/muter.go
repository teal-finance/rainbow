// Copyright 2021 Teal.Finance/Rainbow contributors
// This file is part of Teal.Finance/Rainbow,
// a screener for DeFi options under the MIT License.
// SPDX-License-Identifier: MIT

// Package quiet inhibits alerts when too much.
// Similar wording: mute, limiter, reducer, quieter, mouth-closer, inhibitor.
package quiet

import (
	"fmt"
	"strconv"
	"time"

	"github.com/teal-finance/garcon/timex"
)

// Muter counts the successive alerts, and decrements its counter in absence of alerts.
// When the counter is over the Threshold, new incoming alerts are dropped.
// Once the counter returns to zero, or after NoAlertDuration,
// Muter resumes the alerting. Muter uses the Hysteresis principle:
// https://wikiless.org/wiki/Hysteresis
type Muter struct {
	// Threshold is the level disabling the alerts.
	// The counter must return to zero to consider the situation is back to normal.
	Threshold int

	// NoAlertDuration also allows to consider the verbosity is back to normal.
	NoAlertDuration time.Duration

	// RemindMuteState, if non-zero, allows to still send one alert to remind
	// the alerting is muted since a while.
	// Set value 100 to send this reminder every 100 dropped alerts.
	RemindMuteState int

	// counter is incremented/decremented depending on alerting activity.
	// counter represents the alerting verbosity.
	// counter is always positive.
	counter int

	// muted becomes true when Muter starts inhibiting the alerting,
	// (when counter > Threshold) and returns false when counter=0
	// or d>NoAlertDuration.
	muted bool

	// quietTime is the first call of successive Decrement()
	// without any Increment(). quietTime is used to
	// inform the time since no alert has been notified.
	quietTime time.Time

	// drops is the number of dropped alerts
	// (the unsent alerts to the Notifier).
	drops int
}

func (m *Muter) Increment() (bool, string) {
	m.counter++

	if m.muted {
		m.drops++
		if (m.drops % m.RemindMuteState) > 0 {
			return false, ""
		}
	}

	if m.muted {
		return true, "\n" + "⛔ Still muted, already dropped " + strconv.Itoa(m.drops) + " alerts"
	}

	if m.counter > m.Threshold {
		m.muted = true
		m.drops = 0
		m.quietTime = time.Time{}
		return true, "\n" + "⛔ Mute alerts"
	}

	return true, ""
}

// Decrement decrements the alert verbosity level (the counter)
// and switches to un-muted state when counter reaches zero.
func (m *Muter) Decrement() (bool, string) {
	if !m.muted {
		return false, "" // Already un-muted, do nothing
	}

	var sinceQuietTime time.Duration
	if m.quietTime.IsZero() {
		m.quietTime = time.Now()
	} else {
		sinceQuietTime = time.Since(m.quietTime)
	}

	m.counter--
	if (m.counter > 0) && (sinceQuietTime < m.NoAlertDuration) {
		return false, ""
	}

	m.muted = false
	m.counter = 0

	msg := "✅ No alerts "
	if sinceQuietTime > 0 {
		msg += fmt.Sprintf("since %s (%s ago)", m.quietTime.Format("15:04"), timex.DStr(sinceQuietTime))
	} else {
		msg += "now"
	}
	if m.drops > 1 {
		msg += fmt.Sprintf(", dropped %d alerts", m.drops)
	}
	return true, msg
}
