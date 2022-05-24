package main

import (
	"bytes"

	"github.com/davecgh/go-spew/spew"
	"github.com/ethereum/go-ethereum/common"
	"github.com/teal-finance/rainbow/pkg/rainbow"
)

func main() {
	e := rainbow.Expiries(8)
	spew.Dump(e)

	l := common.HexToHash("0x4c554e4100000000000000000000000000000000000000000000000000000000")
	b := l.Bytes()
	b = bytes.Trim(b, "\x00")
	spew.Dump(string(b))

	/*d := e[0]
	d = d.Add(24 * time.Hour)
	fmt.Println(d)
	d = d.Add(time.Duration(-d.Weekday()+time.Friday) * 24 * time.Hour)
	fmt.Println(d.Weekday())
	fmt.Println(d)*/

	//fmt.Println((d.Weekday() - 7 + time.Friday) % 7)
	//fmt.Println((d.Weekday() + 7 - time.Friday) % 7)
}
