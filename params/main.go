package params

import (
	"net/url"
	"strings"
)

type Params struct {
	params url.Values
}

func Build() *Params {
	return &Params{
		params: url.Values{},
	}
}
func (p *Params) URLSafe() string {
	return strings.Replace(p.params.Encode(), "+", "%20", -1)
}

func (p *Params) Parameter(parameter string) {
	p.params[parameter] = []string{}
}

func (p *Params) Is(parameter string, value ...string) {
	p.params[parameter] = value
}
