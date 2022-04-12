package main

type Arb struct {
	Type         string `json:"type"` // CALL / PUT
	BuyQ         float64
	BuyPx        float64
	BuyProvider  string
	SellQ        float64
	SellPx       float64
	SellProvider string
	SellM        float64
	ROI          float64 //percentage of benef
}

type Arbs []Arb

func buylowsellhigh(blocks Blocks) Arbs {
	arbs := Arbs{}
	for _, block := range blocks {
		if len(block.Options) == 1 {
			continue
		}

		for i := 0; i < len(block.Options)-1; i++ {
			opt := block.Options[i]
			//this also test local arbs, if in one place the bid > ask
			for _, o := range block.Options[i:] {
				if len(opt.Ask) != 0 && len(o.Bid) != 0 && opt.Ask[0].Price < o.Bid[0].Price {
					arbs = append(arbs, Arb{
						Type:         o.Type,
						BuyPx:        opt.Ask[0].Price,
						BuyQ:         opt.Ask[0].Size,
						BuyProvider:  opt.Provider,
						SellPx:       o.Bid[0].Price,
						SellQ:        o.Bid[0].Size,
						SellProvider: o.Provider,
						ROI:          100 * (o.Bid[0].Price - opt.Ask[0].Price) / block.Strike,
					})
				} else if len(opt.Bid) != 0 && len(o.Ask) != 0 && opt.Bid[0].Price > o.Ask[0].Price {
					arbs = append(arbs, Arb{
						Type:        o.Type,
						BuyPx:       o.Ask[0].Price,
						BuyQ:        o.Ask[0].Size,
						BuyProvider: o.Provider,

						SellPx:       opt.Bid[0].Price,
						SellQ:        opt.Bid[0].Size,
						SellProvider: opt.Provider,
						ROI:          100 * (opt.Bid[0].Price - o.Ask[0].Price) / block.Strike,
					})
				}

			}
		}
	}

	return arbs

}
