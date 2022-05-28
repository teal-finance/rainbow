package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strconv"
	"sync"
	"time"

	"github.com/teal-finance/rainbow/pkg/provider"
	"github.com/teal-finance/rainbow/pkg/rainbow"
)

func main() {
	now := time.Now()
	providers := provider.AllProviders()
	var wg = sync.WaitGroup{}
	log.Print("number providers " + fmt.Sprint(len(providers)))
	for _, p := range providers {
		wg.Add(1)
		go loop(p, &wg)
	}

	wg.Wait()
	log.Print(time.Now().Sub(now))

}

func loop(p rainbow.Provider, wg *sync.WaitGroup) {

	defer wg.Done()

	log.Print(p.Name() + " start")
	options, err := p.Options()
	if err != nil {
		log.Print("WARN fetching data from ", p, " : ", err)
	}
	log.Printf("Fetched %v=%v", p.Name(), len(options))
	if len(options) == 0 {
		return
	}
	for _, o := range options {

		save(o)
	}

}

func save(option rainbow.Option) {

	newpath := filepath.Join("./data/", option.Provider)
	err := os.MkdirAll(newpath, os.ModePerm)
	if err != nil {
		log.Print("weird error at file cration: ", err)
	}

	path := newpath + "/" + option.Name + ".csv"
	f, err := os.OpenFile(path, os.O_RDWR|os.O_CREATE|os.O_EXCL, 0666)
	if err == nil {
		log.Print(f.Name() + " created")
		defer f.Close()

		writer := csv.NewWriter(f)
		header := []string{"timestamp", "strike", "bidIV",
			"bid_Price0", "bid_size0", "bid_Price1", "bid_size1", "bid_Price2", "bid_size2", "bid_Price3", "bid_size3", "bid_Price4", "bid_size4",
			"ask_Price0", "ask_size0", "ask_Price1", "ask_size1", "ask_Price2", "ask_size2", "ask_Price3", "ask_size3", "ask_Price4", "ask_size4",
			"askIV", "delta", "gamma", "vega", "theta", "rho", "marketIV"}
		writer.Write(header)
		writer.Flush()
	}
	//log.Print(err)
	//log.Print(path + " exist")
	f, err = os.OpenFile(path, os.O_RDWR|os.O_APPEND, 0666)
	if err != nil {
		log.Print("error opening file that exists", err)
	}
	writer := csv.NewWriter(f)

	optionsToCsv(option, writer)

	//check if file exists

	//if no create file
	// and make csv header

	//if yes, make row
	//append to file
}

func optionsToCsv(option rainbow.Option, writer *csv.Writer) {
	t := time.Now().Unix()
	line := []string{strconv.FormatInt(t, 10), fmt.Sprintf("%.2f", option.Strike), fmt.Sprintf("%.2f", option.BidIV)}
	bids := ordersToCSV(option.Bid)
	asks := ordersToCSV(option.Ask)
	line = append(line, bids...)
	line = append(line, asks...)
	line = append(line,
		fmt.Sprintf("%.4f", option.AskIV),
		fmt.Sprintf("%.4f", option.Greeks.Delta),
		fmt.Sprintf("%.4f", option.Greeks.Gamma),
		fmt.Sprintf("%.4f", option.Greeks.Vega),
		fmt.Sprintf("%.4f", option.Greeks.Theta),
		fmt.Sprintf("%.4f", option.Greeks.Rho),
		fmt.Sprintf("%.4f", option.MarketIV))

	writer.Write(line)

	writer.Flush()
}
func fill(orders []rainbow.Order) []rainbow.Order {
	if len(orders) < 5 {
		return append(orders, make([]rainbow.Order, 5-len(orders))...)
	}
	return orders[0:5]
}

func ordersToCSV(orders []rainbow.Order) []string {
	o := fill(orders)
	return []string{fmt.Sprintf("%f", o[0].Price), fmt.Sprintf("%f", o[0].Size),
		fmt.Sprintf("%f", o[1].Price), fmt.Sprintf("%f", o[1].Size),
		fmt.Sprintf("%f", o[2].Price), fmt.Sprintf("%f", o[2].Size),
		fmt.Sprintf("%f", o[3].Price), fmt.Sprintf("%f", o[3].Size),
		fmt.Sprintf("%f", o[4].Price), fmt.Sprintf("%f", o[4].Size),
	}
}
