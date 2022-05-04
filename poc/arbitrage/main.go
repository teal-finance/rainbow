package main

import (
	"bufio"
	"log"
	"os"

	"github.com/davecgh/go-spew/spew"

	"github.com/teal-finance/rainbow/pkg/provider"
	"github.com/teal-finance/rainbow/pkg/rainbow"
	"github.com/teal-finance/rainbow/pkg/rainbow/storage/dbram"
)

func main() {
	service := rainbow.NewService(provider.AllProviders(), dbram.NewDB())
	service.FetchOptionsFromProviders()

	options, err := service.Options(rainbow.StoreArgs{})
	if err != nil {
		panic(err)
	}
	fileB, err := os.OpenFile("raw.txt", os.O_CREATE|os.O_WRONLY, 0644)

	if err != nil {
		log.Fatalf("failed creating file: %s", err)
	}

	datawriterB := bufio.NewWriter(fileB)

	spew.Fdump(datawriterB, options)
	datawriterB.Flush()
	fileB.Close()
	b := buildCallPut(options)
	file, err := os.OpenFile("dump.txt", os.O_CREATE|os.O_WRONLY, 0644)

	if err != nil {
		log.Fatalf("failed creating file: %s", err)
	}

	datawriter := bufio.NewWriter(file)

	spew.Fdump(datawriter, b)
	datawriter.Flush()
	file.Close()
	arbs := buylowsellhigh(b, 0.25)
	fileA, err := os.OpenFile("arbs.txt", os.O_CREATE|os.O_WRONLY, 0644)

	if err != nil {
		log.Fatalf("failed creating file: %s", err)
	}

	datawriterA := bufio.NewWriter(fileA)
	spew.Fdump(datawriterA, arbs)
	datawriterA.Flush()
	fileA.Close()

}
