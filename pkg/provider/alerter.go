// Copyright 2021 Teal.Finance/Rainbow contributors
// This file is part of Teal.Finance/Rainbow,
// a screener for DeFi options under the MIT License.
// SPDX-License-Identifier: MIT

package provider

import (
	"fmt"
	"strconv"
	"time"

	"github.com/teal-finance/garcon/notifier"
	"github.com/teal-finance/garcon/timex"
	"github.com/teal-finance/rainbow/pkg/rainbow"
)

const (
	// threshold is the verbosity level disabling the alerting.
	threshold = 4

	// noAlertDuration is the time after the verbosity is considered back to normal.
	noAlertDuration = time.Hour

	// remindMuteState allows to still send one alert to remind the alerting is muted since a while.
	// Set value 100 to send this reminder every 100 dropped alerts.
	// Set to zero to disable this feature.
	remindMuteState = 200
)

type alerter struct {
	provider rainbow.Provider
	muter    muter
	notifier notifier.Notifier

	// prefix (in Markdown format) is inserted in all notifications.
	prefix string
}

func newAlerter(namespace string, p rainbow.Provider, n notifier.Notifier) *alerter {
	return &alerter{
		provider: p, // "**" = bold in markdown ; trailing space = separator
		prefix:   namespace + ".**" + p.Name() + "** ",
		notifier: n,
		muter: muter{
			counter:   0,
			muted:     false,
			quietTime: time.Time{},
			dropped:   0,
		},
	}
}

// Name is part of the Provider interface.
func (a *alerter) Name() string {
	return a.provider.Name()
}

// Options is part of the Provider interface.
func (a *alerter) Options() ([]rainbow.Option, error) {
	options, err := a.provider.Options()

	go func() {
		e := a.vet(options, err)
		if e != nil {
			log.Errorf("Alerter %s: %s", a.Name(), e)
		}
	}()

	return options, err
}

func (a *alerter) vet(options []rainbow.Option, err error) error {
	notifyError := (err != nil)
	notifyEmpty := (len(options) == 0)

	var ok bool
	var msg string
	if notifyError || notifyEmpty {
		// => increment alerting verbosity
		ok, msg = a.muter.increment()
	} else {
		// no alert => decrement alerting verbosity
		ok, msg = a.muter.decrement()
	}

	if !ok {
		return nil // muted: do not notify
	}

	switch {
	case notifyError:
		msg = ":alert: API error: " + err.Error() + msg
	case notifyEmpty:
		msg = ":question: no options" + msg
	}

	return a.notifier.Notify(a.prefix + msg)
}

// muter inhibits alerting when too verbose.
// muter counts the successive alerts, and decrements its counter in absence of alerts.
// When the counter is over the Threshold, new incoming alerts are dropped.
// Once the counter returns to zero, or after NoAlertDuration,
// muter resumes the alerting. This is the Hysteresis principle:
// https://wikiless.org/wiki/Hysteresis
type muter struct {
	// counter is incremented/decremented depending on alerting activity, but is never negative.
	counter int

	// muted becomes true when Muter starts inhibiting the alerting,
	muted bool

	// quietTime is the first call of Decrement() after Increment().
	// quietTime is used to inform the time since the situation is back to normal.
	quietTime time.Time

	// dropped is the number of dropped alerts (the unsent alerts to the Notifier).
	dropped int
}

// increment increments the internal counter and returns false when in muted state.
// Every remindMuteState calls, increment also returns a message to remind the muted state.
func (m *muter) increment() (ok bool, msg string) {
	m.counter++

	if m.muted {
		m.dropped++
		if (remindMuteState == 0) || (m.dropped%remindMuteState) > 0 {
			return false, ""
		}
	}

	if m.muted {
		return true, "\n" + "⛔ Still muted, already dropped " + strconv.Itoa(m.dropped) + " alerts"
	}

	if m.counter > threshold {
		m.muted = true
		m.dropped = 0
		m.quietTime = time.Time{}
		return true, "\n" + "⛔ Mute alerts"
	}

	return true, ""
}

// decrement decrements the alert verbosity level (the counter)
// and switches to un-muted state when counter reaches zero.
func (m *muter) decrement() (ok bool, msg string) {
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
	if (m.counter > 0) && (sinceQuietTime < noAlertDuration) {
		return false, ""
	}

	m.muted = false
	m.counter = 0

	msg = "✅ No alerts "
	if sinceQuietTime > 0 {
		msg += fmt.Sprintf("since %s (%s ago)", m.quietTime.Format("15:04"), timex.DStr(sinceQuietTime))
	} else {
		msg += "now"
	}
	if m.dropped > 1 {
		msg += fmt.Sprintf(", dropped %d alerts", m.dropped)
	}
	return true, msg
}
