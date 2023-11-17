package pet

import (
	"time"
)

type Pet struct {
	Name      string `json:"name"`
	FedAt     int64  `json:"fed_at"`
	AmountFed string `json:"amount_fed"`
	IsInside  bool   `json:"is_inside"`
	InsideAt  int64  `json:"inside_at"`
}

func (p *Pet) UpdateInside() {
	if p.IsInside {
		p.InsideAt = time.Now().Unix()
		p.IsInside = false
	} else {
		p.InsideAt = time.Now().Unix()
		p.IsInside = true
	}
}

func (p *Pet) Fed(amount string) {
	p.FedAt = time.Now().Unix()
	p.AmountFed = amount
}
