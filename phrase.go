package markov

type Phrase []Morpheme
type PhraseString []MorphemeString

func (p PhraseString) Phrase() Phrase {
	ans := make(Phrase, 0, len(p))
	for _, v := range p {
		ans = append(ans, v.Morpheme())
	}
	return ans
}

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
