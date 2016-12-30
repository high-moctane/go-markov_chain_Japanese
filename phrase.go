package markov

type Phrase []Morpheme

func (p Phrase) OriginalForm() (ans string) {
	for _, v := range p {
		ans += v.OriginalForm
	}
	return
}

func (p Phrase) Pronounciation() (ans string) {
	for _, v := range p {
		ans += v.Pronounciation
	}
	return
}
