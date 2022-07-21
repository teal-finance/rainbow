// Copyright 2021 Teal.Finance/Rainbow contributors
// This file is part of Teal.Finance/Rainbow,
// a screener for DeFi options under the MIT License.
// SPDX-License-Identifier: MIT

package provider

import (
	"fmt"
	"log"
	"time"

	"github.com/teal-finance/notifier"
	"github.com/teal-finance/rainbow/pkg/quiet"
	"github.com/teal-finance/rainbow/pkg/rainbow"
)

const (
	muteLevel            = 4         // above => mute the alerting
	backToNormalDuration = time.Hour // after this time => resumes alerting
	remindMuteState      = 200       // one notification every 100
)

type alerter struct {
	provider rainbow.Provider
	muter    quiet.Muter
	notifier notifier.Notifier

	// prefix is inserted in all alerts, and can be in Markdown format.
	prefix string
}

func newAlerter(namespace string, p rainbow.Provider, n notifier.Notifier) *alerter {
	return &alerter{
		provider: p, // "**" = bold in markdown ; trailing space = separator
		prefix:   namespace + ".**" + p.Name() + "** ",
		notifier: n,
		muter: quiet.Muter{
			Threshold:       muteLevel,
			NoAlertDuration: backToNormalDuration,
			RemindMuteState: remindMuteState,
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
			log.Printf("ERR Alerter %s: %s", a.Name(), e)
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
		ok, msg = a.muter.Increment()
	} else {
		// no alert => decrement alerting verbosity
		ok, msg = a.muter.Decrement()
	}

	if !ok {
		return nil // muted: do not notify
	}

	switch {
	case notifyError:
		msg = fmt.Sprint(":alert: API error: ", err) + msg
	case notifyEmpty:
		msg = ":question: no options" + msg
	}

	return a.notifier.Notify(a.prefix + msg)
}
