package main

import (
	"fmt"
	"log"

	"github.com/davecgh/go-spew/spew"
	"github.com/teal-finance/rainbow"
	"github.com/teal-finance/rainbow/pkg/deribit"
	"github.com/teal-finance/rainbow/pkg/psyoptions"
	"github.com/teal-finance/rainbow/pkg/zerox"
)

func main() {

	all()
}

func tryOpyn() ([]rainbow.Options, error) {
	instruments := zerox.Instruments()
	//spew.Dump(instruments)
	/*orderBook, err := zerox.GetOrderBook(instruments, "Opyn")
	if err != nil {
		log.Println(err)
		return []rainbow.Options{}, err
	}*/

	//spew.Dump(orderBook[0])

	//log.Println(zerox.ConvertToSolidity(10.0, 8))

	orderBook, err := zerox.GetAggregatedOrderBook(instruments, "Opyn", 2.0)
	if err != nil {
		log.Println(err)
		return []rainbow.Options{}, err

	}

	return orderBook, err
}

func tryDeribit() ([]rainbow.Options, error) {
	instruments, err := deribit.Instruments("BTC")
	if err != nil {
		log.Println(err)
		return []rainbow.Options{}, err
	}

	log.Println(instruments[10])
	spew.Dump(instruments[10])

	orderBook, err := deribit.GetOrderBook(instruments[10:15], 5)
	if err != nil {
		log.Println(err)
		return []rainbow.Options{}, err
	}

	//spew.Dump(orderBook[0].Offers)
	return orderBook, nil
}

func tryPsyops() ([]rainbow.Options, error) {
	return psyoptions.InstrumentsFromAllMarkets()

}

func all() {
	options := []rainbow.Options{}
	//psy
	psy, err := psyoptions.InstrumentsFromAllMarkets()
	if err != nil {
		fmt.Println("psy error ", err)
		return
	}
	options = append(options, psy...)
	//opyn
	op, err := tryOpyn()
	if err != nil {
		fmt.Println("opyn error ", err)
		return
	}
	options = append(options, op...)

	der, err := tryDeribit()
	if err != nil {
		fmt.Println("der error ", err)
		return
	}
	options = append(options, der...)

	spew.Dump(options)
}
