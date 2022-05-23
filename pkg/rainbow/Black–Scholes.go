package rainbow

import (
	"log"
	"time"

	"github.com/jasonmerecki/gopriceoptions"
)

// Optimized from https://github.com/jasonmerecki/gopriceoptions/blob/master/blacklike.go

func IV(optionType string, premium float64, underlyingPrice float64, strike float64, expiry string) float64 {

	//Call true Put false
	callType := optionType == "CALL"

	t, err := time.Parse("2006-01-02 15:04:05", o.Expiry)
	if err != nil {
		log.Fatalf("time parse ", err)
	}
	timeToExpiration := t.Sub(time.Now().UTC()).Seconds() / (86400 * 365)
	// startAnchorVolatility at 0 because their code starts it at 0.5 . I don't know if that is OK!
	// riskFreeInterest is at 0 for all providers AFAICT
	// dividend is at 0 but check for staking assets if anyone bother to put the exact rate
	return gopriceoptions.BSImpliedVol(callType, premium, underlyingPrice, strike, timeToExpiration, 0, 0, 0)
}

/*
func computeIV(options[]Option)[]Option{
	for _,o:= range options{
		log.Print(o.AskIV)
		log.Print(IV(o.Type,o.Ask[0],))
	}
}*/
/*
func ComputeGreeks()TheGreeks{
	var greeks TheGreeks
	greeks.Delta=
	greeks.Rho=
	greeks.Theta=
	greeks.Vega=

	return greeks

}*/
