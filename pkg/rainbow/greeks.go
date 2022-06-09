package rainbow

import (
	"log"
	"time"

	"github.com/jasonmerecki/gopriceoptions"
)

// Optimized from https://github.com/jasonmerecki/gopriceoptions/blob/master/blacklike.go

// inspiration/references:
// Black Scholes Model Python: https://www.codearmo.com/python-tutorial/options-trading-black-scholes-model

//BSIV: IV computation from Black Scholes
func BSIV(optionType string, premium float64, underlyingPrice float64, strike float64, expiry string) float64 {

	//Call true Put false
	callType := optionType == "CALL"

	t, err := time.Parse("2006-01-02 15:04:05", expiry)
	if err != nil {
		log.Fatalf("time parse ", err)
	}
	timeToExpiration := t.Sub(time.Now().UTC()).Seconds() / (86400 * 365)
	// startAnchorVolatility at 0 because their code starts it at 0.5 . I don't know if that is OK!
	// riskFreeInterest is at 0 for all providers AFAICT
	// dividend is at 0 but check for staking assets if anyone bother to put the exact rate
	return gopriceoptions.BSImpliedVol(callType, premium, underlyingPrice, strike, timeToExpiration, 0, 0, 0)
}

//BSGreeks: Greeks computation from Black Scholes
// this assumes you have or computed IV first, using BSIV for example
func BSGreeks(optionType string, premium float64, underlyingPrice float64, strike float64,
	expiry string, volatility float64) TheGreeks {
	var greeks TheGreeks

	//Call true Put false
	callType := optionType == "CALL"

	t, err := time.Parse("2006-01-02 15:04:05", expiry)
	if err != nil {
		log.Fatalf("time parse ", err)
	}
	timeToExpiration := t.Sub(time.Now().UTC()).Seconds() / (86400 * 365)
	greeks.Delta = gopriceoptions.BSDelta(callType, underlyingPrice, strike, timeToExpiration, volatility, 0, 0)
	greeks.Rho = gopriceoptions.BSRho(callType, underlyingPrice, strike, timeToExpiration, volatility, 0, 0)
	greeks.Theta = gopriceoptions.BSTheta(callType, underlyingPrice, strike, timeToExpiration, volatility, 0, 0)
	greeks.Gamma = gopriceoptions.BSGamma(underlyingPrice, strike, timeToExpiration, volatility, 0, 0)
	greeks.Vega = gopriceoptions.BSVega(underlyingPrice, strike, timeToExpiration, volatility, 0, 0)
	return greeks

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



}*/

// TODO Black Implementation to compute the greeks from Deribit and in case we need them for somtheing else
// see source here: https://fr.wikipedia.org/wiki/Mod%C3%A8le_de_Black
