package x

import "flag"

const (
	DefaultAddr      = "http://localhost"
	DefaultPort      = 8090
	DefaultProviders = "ALL" //"deribit,delta,thalex,lyra,synquote,aevo,rysk,sdx,thales"
)

// I have no idea what was supposed to be the purpose of this var
var Note = flag.Bool("note", false, "Show additional note ...")
