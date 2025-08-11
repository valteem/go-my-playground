package voidtags

type header struct {
	Key    string
	Values []string
}

type request struct {
	Paths   []string
	Headers []header
}

type peer struct {
	Principals []string
}

type rule struct {
	Name    string
	Source  peer
	Request request
}
