package pet

import (
	"time"
)

type Pet struct {
	FedAt     int64  `json:"fed_at"`
	AmountFed string `json:"amount_fed"`
	IsInside  bool   `json:"is_inside"`
}

func (p *Pet) UpdateInside() {
	if p.IsInside {
		p.IsInside = false
	} else {
		p.IsInside = true
	}
}

func (p *Pet) Fed(amount string) {
	p.FedAt = time.Now().Unix()
	p.AmountFed = amount
}
