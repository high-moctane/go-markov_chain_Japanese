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

func (p Phrase) Morae() ([]Mora, bool) {
	pron, ok := p.Pronounciation()
	if !ok {
		return []Mora{}, false
	}
	runes := []rune(pron + "*")
	ans := make([]Mora, 0, len(runes)-1)

	for i, end := 0, len(runes)-1; i < end; i++ {
		if mora, ok := katakana[string(runes[i:i+2])]; ok {
			ans = append(ans, mora)
			i++
		} else if mora, ok := katakana[string(runes[i])]; ok {
			ans = append(ans, mora)
		} else if runes[i] == 'ー' {
			newMora := Mora{"", ans[len(ans)-1].vowel}
			ans = append(ans, newMora)
		} else {
			return []Mora{}, false
		}
	}
	return ans, true
}
