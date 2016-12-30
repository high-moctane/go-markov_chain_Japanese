package markov

import "strings"

const (
	OriginalForm = iota
	PartOfSpeech
	PartOfSpeechSection1
	PartOfSpeechSection2
	PartOfSpeechSection3
	ConjugatedForm1
	ConjugatedForm2
	Inflection
	Reading
	Pronounciation
)

type MorphemeString string
type Morpheme []string

func NewMorphemeString(s string) MorphemeString {
	tab := strings.Split(s, "\t")
	tab[1] = strings.Replace(tab[1], "*", "", -1)
	ans := strings.Join(tab, "\t")
	commas := strings.Count(ans, ",")
	return MorphemeString(ans + strings.Repeat(",", 8-commas))
}

func (m MorphemeString) Morpheme() Morpheme {
	return NewMorpheme(m)
}

func NewMorpheme(m MorphemeString) Morpheme {
	ans := make([]string, 10)
	tab := strings.Split(string(m), "\t")
	ans[0] = tab[0]
	copy(ans[1:], strings.Split(tab[1], ","))
	return ans
}
