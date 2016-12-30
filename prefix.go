package markov

import "strings"

type Prefix []MorphemeString

func firstPrefix(order int) Prefix {
	prefix := make([]MorphemeString, order)
	for i, _ := range prefix {
		prefix[i] = BOS
	}
	return prefix
}

func (p Prefix) Shift(m MorphemeString) {
	copy(p, p[1:])
	p[len(p)-1] = m
}

func (p Prefix) String() string {
	slice := make([]string, len(p))
	for i, v := range p {
		slice[i] = string(v)
	}
	return strings.Join(slice, "\n")
}
