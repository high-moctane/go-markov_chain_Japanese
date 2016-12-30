package markov

type Data struct {
	MarkovArgs  map[string]string
	Order       int
	firstPrefix Prefix
	Chain       map[string][]MorphemeString
}

func (d *Data) FirstPrefix() Prefix {
	ans := make(Prefix, len(d.firstPrefix))
	copy(ans, d.firstPrefix)
	return ans
}
