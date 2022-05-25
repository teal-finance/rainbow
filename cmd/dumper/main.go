package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"path/filepath"
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
	log.Print("back to the main")

	wg.Wait()
	fmt.Println(time.Now().Sub(now))

}

func loop(p rainbow.Provider, wg *sync.WaitGroup) {

	defer wg.Done()

	log.Print(p.Name() + " start")
	o, err := p.Options()
	if err != nil {
		log.Print("WARN fetching data from ", p, " : ", err)
	}
	log.Printf("Fetched %v=%v", p.Name(), len(o))
	save(o)
	log.Print(p.Name() + " end")

}

func save(options []rainbow.Option) {
	if len(options) == 0 {
		return
	}
	newpath := filepath.Join(".", options[0].Provider)
	log.Print(newpath)
	err := os.MkdirAll(newpath, os.ModePerm)
	if err != nil {
		log.Print("weird error at file cration: ", err)
	}

	path := newpath + "/" + options[0].Name + ".csv"
	f, err := os.OpenFile(path, os.O_RDWR|os.O_CREATE|os.O_EXCL, 0666)
	if err == nil {
		log.Print(f.Name() + " created")
		defer f.Close()

		writer := csv.NewWriter(f)
		header := []string{"BlockNumber", "TxNumber", "Min total tx", "Max total tx",
			"50th percentile of propagation(ms)", "75th percentile of propagation(ms)", "95th percentile of propagation(ms)"}
		writer.Write(header)

		optionsToCsv(options[0], writer)

		return
	}
	log.Print(err)
	log.Print(path + " exist")
	f, err = os.OpenFile(path, os.O_RDWR|os.O_APPEND, 0666)
	if err != nil {
		log.Print(err)
	}
	writer := csv.NewWriter(f)

	optionsToCsv(options[0], writer)

	//check if file exists

	//if no create file
	// and make csv header

	//if yes, make row
	//append to file
}

func optionsToCsv(option rainbow.Option, writer *csv.Writer) {
	line := []string{"line", "line", "line"}
	writer.Write(line)
	writer.Flush()
}
