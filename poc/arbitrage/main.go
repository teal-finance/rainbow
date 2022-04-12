package main

import (
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
	b := buildCallPut(options)
	spew.Dump(b)

}
