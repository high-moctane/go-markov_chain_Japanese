package markov

type Phrase []Morpheme

func (p Phrase) OriginalForm() (ans string) {
	for _, v := range p {
		ans += v.OriginalForm
	}
	return
}

func (p Phrase) Pronounciation() (string, bool) {
	var ans string
	for _, v := range p {
		if v.Pronounciation == "" {
			return "", false
		}
		ans += v.Pronounciation
	}
	return ans, true
}
