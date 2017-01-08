package markov

import "strings"

type MorphemeString string

type Morpheme struct {
	OriginalForm         string
	PartOfSpeech         string
	PartOfSpeechSection1 string
	PartOfSpeechSection2 string
	PartOfSpeechSection3 string
	ConjugatedForm1      string
	ConjugatedForm2      string
	Inflection           string
	Reading              string
	Pronounciation       string
}

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
	return Morpheme{
		OriginalForm:         ans[0],
		PartOfSpeech:         ans[1],
		PartOfSpeechSection1: ans[2],
		PartOfSpeechSection2: ans[3],
		PartOfSpeechSection3: ans[4],
		ConjugatedForm1:      ans[5],
		ConjugatedForm2:      ans[6],
		Inflection:           ans[7],
		Reading:              ans[8],
		Pronounciation:       ans[9],
	}
}
