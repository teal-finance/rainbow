// Copyright 2021 Teal.Finance/Rainbow contributors
// This file is part of Teal.Finance/Rainbow,
// a screener for DeFi options under the MIT License.
// SPDX-License-Identifier: MIT

// Package quiet inhibits alerts when too much.
// Similar wording: mute, limiter, reducer, quieter, mouth-closer, inhibitor.
package quiet

import (
	"fmt"
	"log"
	"time"

	"github.com/teal-finance/garcon/timex"
	"github.com/teal-finance/notifier"
)

// Muter drops alerts when it is forwarding too much.
// Muter counts the successive alerts, and decrements its counter in absence of alerts.
// When the counter is over the Threshold, new incoming alerts are dropped.
// Once the counter returns to zero, or after NoAlertDuration,
// Muter resumes the alerting. Muter uses the Hysteresis principle:
// https://wikiless.org/wiki/Hysteresis
// Muter implements the Notifier interface.
type Muter struct {
	// Prefix is inserted in all alerts.
	// Thus Prefix can be in Markdown format.
	Prefix string

	Notifier notifier.Notifier

	// Threshold is the level disabling the alerts.
	// The counter must return to zero
	// to consider the situation is back to normal.
	Threshold int

	// NoAlertDuration also allows to
	// consider the verbosity is back to normal.
	NoAlertDuration time.Duration

	// RemindMuteState, if non-zero, allows to
	// still send a alerts to remind
	// the alerts are muted since a while.
	// Set value 100 to send this reminder
	// every 100 dropped alerts.
	RemindMuteState int

	// counter is incremented for each alerts
	// (when Notify() is called), and decremented
	// when NotifyLowVerbosity() is called.
	// counter is always positive.
	// The value counter=0 indicates the verbosity is low
	// thus Muter enable again the alerts.
	counter int

	// muted becomes true when Muter starts inhibiting the alerting,
	// (when counter > Threshold) and returns false when counter=0.
	muted bool

	// quietTime is the first call to NotifyLowVerbosity()
	// after Notify() had been called. quietTime is used to
	// inform the time since Notify() has not been called.
	quietTime time.Time

	// drops is the number of dropped alerts
	// (the unsent alerts to Notifier).
	drops int
}

func (m *Muter) Notify(msg string) error {
	m.counter++

	if m.muted {
		m.drops++
		if (m.drops % m.RemindMuteState) > 0 {
			return nil
		}
	}

	note := ""
	if m.muted {
		note = fmt.Sprintf("\n"+"⛔ Still muted, "+
			"already dropped %d alerts", m.drops)
	} else if m.counter > m.Threshold {
		m.muted = true
		m.drops = 0
		m.quietTime = time.Time{}
		note = "\n" + "⛔ Mute alerts"
	}

	return m.notify(msg + note)
}

// NoAlert decrements the alert verbosity level (the counter)
// and switches to un-muted state when counter reaches zero.
func (m *Muter) NoAlert() error {
	if !m.muted {
		return nil // Already un-muted, do nothing
	}

	var sinceQuietTime time.Duration
	if m.quietTime.IsZero() {
		m.quietTime = time.Now()
	} else {
		sinceQuietTime = time.Since(m.quietTime)
	}

	m.counter--
	if (m.counter > 0) && (sinceQuietTime < m.NoAlertDuration) {
		return nil
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
	return m.notify(msg)
}

func (m *Muter) notify(msg string) error {
	return m.Notifier.Notify(m.Prefix + " " + msg)
}
