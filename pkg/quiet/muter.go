// Copyright 2021 Teal.Finance/Rainbow contributors
// This file is part of Teal.Finance/Rainbow,
// a screener for DeFi options under the MIT License.
// SPDX-License-Identifier: MIT

// Package quiet temporary inhibits notifications.
// Similar wording: mute, limiter, quieter, mouth-closer, inhibitor.
package quiet

import (
	"fmt"
	"log"
	"time"

	"github.com/teal-finance/garcon/timex"
	"github.com/teal-finance/notifier"
)

// Muter pauses notification when
// the rate becomes too high and
// waits for a lower level to resume notifications.
// This is the Hysteresis principle:
// https://wikiless.org/wiki/Hysteresis
// Muter implements the Notifier interface.
type Muter struct {
	// Prefix is inserted in all notifications.
	// Thus Prefix can be in Markdown format.
	Prefix string

	Notifier notifier.Notifier

	// Threshold is the level disabling the notifications.
	// The counter must return to zero
	// to consider the situation is back to normal.
	Threshold int

	// BackToNormalDuration also allows to
	// consider the situation is back to normal.
	BackToNormalDuration time.Duration

	// RemindMuteState, if non-zero, allows to
	// still send a notification to remind
	// the notifications are muted since a while.
	// Set value 100 to send this reminder
	// every 100 dropped notifications.
	RemindMuteState int

	counter   int
	muted     bool
	quietTime time.Time
	drops     int // number of dropped notifications
}

func (m *Muter) Notify(msg string) error {
	m.counter++

	if m.muted {
		m.drops++
		if (m.drops % m.RemindMuteState) > 0 {
			log.Printf("DBG: Muter %s drops notification count=%d d=%s",
				m.Prefix, m.counter, timex.DStr(time.Since(m.quietTime)))
			return nil
		}
	}

	note := ""
	if m.muted {
		note = fmt.Sprintf("\n"+"⛔ Still muted, "+
			"already dropped %d notifications", m.drops)
	} else if m.counter > m.Threshold {
		m.muted = true
		m.drops = 0
		m.quietTime = time.Time{}
		note = "\n" + "⛔ Mute notifications"
	}

	return m.notify(msg + note)
}

func (m *Muter) NotifyIfBackToNormal() error {
	if m.counter == 0 {
		return nil // Already OK, do nothing
	}

	var d time.Duration
	if m.quietTime.IsZero() {
		m.quietTime = time.Now()
	} else {
		d = time.Since(m.quietTime)
	}

	m.counter--
	if (m.counter > 0) && (d < m.BackToNormalDuration) {
		log.Printf("DBG Muter %s: Not yet back to normal count=%d d=%s",
			m.Prefix, m.counter, timex.DStr(time.Since(m.quietTime)))
		return nil
	}

	m.muted = false
	m.counter = 0

	msg := "✅ Back to normal"
	if d > 0 {
		msg += fmt.Sprintf(" since %s (%s ago)", m.quietTime.Format("15:04"), timex.DStr(d))
	}
	msg += fmt.Sprintf(", dropped %d notifications", m.drops)
	return m.notify(msg)
}

func (m *Muter) notify(msg string) error {
	return m.Notifier.Notify(m.Prefix + " " + msg)
}
