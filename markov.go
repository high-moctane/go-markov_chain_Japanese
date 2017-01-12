package markov

import "github.com/high-moctane/go-mecab_slice"

type Data interface {
	Add(mecabs.PhraseString)
	Generate(int) mecabs.PhraseString
}

type Markov struct {
	MeCabS *mecabs.MeCabS
	Data   Data
}

func New(m *mecabs.MeCabS, d Data) Markov {
	return Markov{
		MeCabS: m,
		Data:   d,
	}
}

func (m *Markov) Add(s string) error {
	ps, err := m.MeCabS.NewPhraseString(s)
	if err != nil {
		return err
	}
	m.Data.Add(ps)
	return nil
}

func (m *Markov) Generate(length int) mecabs.Phrase {
	return m.Data.Generate(length).Phrase()
}
