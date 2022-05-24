package main

import (
	"log"
	"time"

	"github.com/jasonmerecki/gopriceoptions"
	"github.com/spewerspew/spew"
	"github.com/teal-finance/rainbow/pkg/provider/deribit"
	"github.com/teal-finance/rainbow/pkg/rainbow"
)

func main() {
	price, _ := deribit.GetIndexPrice("SOL")
	log.Print("Price $", price)
	p := new(deribit.Provider)
	options, err := p.Options()
	if err != nil {
		log.Fatalf("option fetch ", err)
	}
	spew.Dump(options[0:2])
	o := options[1]
	//Call true Put false
	oType := o.Type == "CALL"

	t, err := time.Parse("2006-01-02 15:04:05", o.Expiry)
	if err != nil {
		log.Fatalf("time parse ", err)
	}

	log.Print("IV0: ", o.BidIV)
	iv := rainbow.IV(o.Type, o.Bid[0].Price, price, o.Strike, o.Expiry)
	log.Print("IV1: ", iv*100)
	log.Print("IV2: ", o.AskIV)
	iv = rainbow.IV(o.Type, o.Ask[0].Price, price, o.Strike, o.Expiry)
	log.Print("IV3: ", iv*100)

	d := t.Sub(time.Now().UTC()).Seconds() / (86400 * 365)
	s := price
	k := o.Strike
	time := d
	v := 123.23 / 100 //o.AskIV / 100
	r := 0.0
	q := 0.0
	delta := gopriceoptions.BSDelta(oType, s, k, time, v, r, q)
	gamma := gopriceoptions.BSGamma(s, k, time, v, r, q)
	vega := gopriceoptions.BSVega(s, k, time, v, r, q)
	theta := gopriceoptions.BSTheta(oType, s, k, time, v, r, q)
	rho := gopriceoptions.BSRho(oType, s, k, time, v, r, q)
	log.Print("vega ", vega)
	log.Print("theta ", theta)
	log.Print("rho ", rho)
	log.Print("gamma ", gamma)
	log.Print("delta ", delta)

}
