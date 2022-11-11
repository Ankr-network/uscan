package types

type Pager struct {
	Offset int64 `query:"offset"`
	Limit  int64 `query:"limit"`
}

func (p *Pager) Complete() {
	if p.Offset < 0 {
		p.Offset = 0
	}
	if p.Limit <= 0 {
		p.Limit = 10
	}
}

type SearchFilter struct {
	Type    int    `query:"type"` // 1 all filter 2 address
	Keyword string `query:"keyword"`
}
