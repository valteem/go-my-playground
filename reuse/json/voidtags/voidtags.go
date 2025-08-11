package voidtags

type Header struct {
	Key    string
	Values []string
}

type Request struct {
	Paths   []string
	Headers []Header
}

type Peer struct {
	Principals []string
}

type Rule struct {
	Name    string
	Source  Peer
	Request Request
}

type RuleOption func(*Rule)

func RuleName(s string) RuleOption {
	return func(r *Rule) {
		r.Name = s
	}
}

func RuleSource(p Peer) RuleOption {
	return func(r *Rule) {
		r.Source = p
	}
}

func RuleRequest(req Request) RuleOption {
	return func(r *Rule) {
		r.Request = req
	}
}

func NewRule(opts ...RuleOption) Rule {
	r := Rule{}
	for _, opt := range opts {
		opt(&r)
	}
	return r
}
