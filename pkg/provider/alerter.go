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
	remindMuteState      = 100       // one notification every 100
)

type alerter struct {
	provider rainbow.Provider
	notifier quiet.Muter
}

func newAlerter(namespace string, p rainbow.Provider, n notifier.Notifier) *alerter {
	prefix := namespace + ".**" + p.Name() + "**" // Markdown format
	return &alerter{
		provider: p,
		notifier: quiet.Muter{
			Prefix:          prefix,
			Notifier:        n,
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
	if err != nil {
		return a.notifier.Notify(fmt.Sprint(":alert: API error: ", err))
	}

	if len(options) == 0 {
		return a.notifier.Notify(":question: no options")
	}

	return a.notifier.NotifyLowVerbosity()
}
