package main

import (
	"log"
	"os"

	//"github.com/davecgh/go-spew/spew"

	"github.com/gookit/color"
	"github.com/jedib0t/go-pretty/v6/table"
	"github.com/teal-finance/rainbow"
	"github.com/teal-finance/rainbow/pkg/deribit"
	"github.com/teal-finance/rainbow/pkg/psyoptions"
	"github.com/teal-finance/rainbow/pkg/zerox"
)

func main() {

	//spew.Dump(all())
	options, err := all() //tryDeribit() //tryPsyops() //tryOpyn() //all()
	if err != nil {
		log.Println(err)
		return
	}
	PrintOptions(options)
	//spew.Dump(options)

}
func PrintOptions(options []rainbow.Options) {
	green := color.FgGreen.Render
	red := color.FgRed.Render

	t := newTable(fmt.Sprintf("\t\t 29OCT21 CeDeFi options: (%v)", len(options)))
	t.AppendHeader(table.Row{"Provider", "Asset", "Type", "Bids size", "Bids Price", "Strike", "Asks Price", "Asks size", "Instrument"})
	for _, option := range options {
		t.AppendRows([]table.Row{{prov(option.Provider), option.Asset, option.Type,
			option.Offers[0].Quantity, green(option.Offers[0].Price), option.Strike,
			red(option.Offers[len(option.Offers)-1].Price), option.Offers[len(option.Offers)-1].Quantity, option.Name}})

	}

	t.SortBy([]table.SortBy{
		{Name: "Strike", Mode: table.AscNumeric},
		{Name: "Type", Mode: table.Dsc},
	})
	t.Render()
}
func prov(p string) string {
	magenta := color.FgMagenta.Render
	green := color.FgGreen.Render
	blue := color.FgCyan.Render
	s := ""
	if p == "Opyn" {
		s = blue(p)
	} else if p == "Deribit" {
		s = green(p)

	} else if p == "PsyOptions" {
		s = magenta(p)
	}
	return s
}
func newTable(title string) table.Writer {
	t := table.NewWriter()
	t.SetOutputMirror(os.Stdout)
	t.SetStyle(table.StyleLight)
	t.SetTitle(title)
	return t
}

func tryOpyn() ([]rainbow.Options, error) {
	instruments := zerox.Instruments()
	//spew.Dump(instruments)
	/*orderBook, err := zerox.GetOrderBook(instruments, "Opyn")
	if err != nil {
		log.Println(err)
		return []rainbow.Options{}, err
	}*/

	// spew.Dump(orderBook[0])

	// log.Println(zerox.ConvertToSolidity(10.0, 8))

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

	// log.Println(instruments[10])
	// spew.Dump(instruments[10])

	orderBookBTC, err := deribit.GetOrderBook(instruments, 5)
	if err != nil {
		log.Println(err)
		return []rainbow.Options{}, err
	}

	instruments, err = deribit.Instruments("ETH")
	if err != nil {
		log.Println(err)
		return []rainbow.Options{}, err
	}

	// log.Println(instruments[10])
	// spew.Dump(instruments[10])

	orderBookETH, err := deribit.GetOrderBook(instruments, 5)
	if err != nil {
		log.Println(err)
		return []rainbow.Options{}, err
	}

	// spew.Dump(orderBook[0].Offers)
	return append(orderBookBTC, orderBookETH...), nil
}

func tryPsyops() ([]rainbow.Options, error) {
	return psyoptions.InstrumentsFromAllMarkets()
}

func all() ([]rainbow.Options, error) {
	options := []rainbow.Options{}

	// psy
	psy, err := psyoptions.InstrumentsFromAllMarkets()
	if err != nil {
		log.Println("psy error ", err)
		return []rainbow.Options{}, err
	}

	options = append(options, psy...)

	// opyn
	op, err := tryOpyn()
	if err != nil {
		log.Println("opyn error ", err)
		return []rainbow.Options{}, err
	}

	options = append(options, op...)

	der, err := tryDeribit()
	if err != nil {
		log.Println("der error ", err)
		return []rainbow.Options{}, err
	}

	return append(options, der...), nil
}
