package markov

type Data struct {
	MarkovArgs       map[string]string
	Order            int
	FirstPrefixField Prefix
	Chain            map[string][]MorphemeString
}

func (d *Data) FirstPrefix() Prefix {
	ans := make(Prefix, len(d.FirstPrefixField))
	copy(ans, d.FirstPrefixField)
	return ans
}
