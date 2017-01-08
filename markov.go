package markov

import (
	"errors"
	"math/rand"
	"strings"
	"sync"
	"time"

	mecab "github.com/shogo82148/go-mecab"
)

const (
	BOS = "\tBOS,,,,,,,,"
	EOS = "\tEOS,,,,,,,,"
)

type Markov struct {
	MeCab mecab.MeCab
	Mutex sync.Mutex
	Data  Data
}

func init() {
	rand.Seed(time.Now().UnixNano())
}

func New(order int, args map[string]string) (*Markov, error) {
	if order <= 0 {
		return &Markov{}, errors.New("invalid order.")
	}
	tagger, err := mecab.New(args)
	if err != nil {
		return &Markov{}, err
	}

	return &Markov{
		MeCab: tagger,
		Data: Data{
			MarkovArgs:       args,
			Order:            order,
			FirstPrefixField: firstPrefix(order),
			Chain:            map[string][]MorphemeString{},
		},
	}, nil
}

func (m *Markov) Destroy() {
	m.MeCab.Destroy()
	m.Data = Data{}
}

func (m Markov) Add(s string) error {
	parsedStr, err := m.MeCab.Parse(s)
	if err != nil {
		return err
	}
	nodes := MakePhraseString(parsedStr)
	prefix := m.Data.FirstPrefix()

	m.Mutex.Lock()
	defer m.Mutex.Unlock()
	for _, node := range nodes {
		key := prefix.String()
		m.Data.Chain[key] = append(m.Data.Chain[key], node)
		prefix.Shift(node)
	}
	key := prefix.String()
	m.Data.Chain[key] = append(m.Data.Chain[key], EOS)
	return nil
}

func MakePhraseString(s string) PhraseString {
	nodes := strings.Split(s, "\n")
	nodes = nodes[:len(nodes)-2]
	ans := make([]MorphemeString, len(nodes))
	for i, node := range nodes {
		tab := strings.Split(node, "\t")
		tab[1] = strings.Replace(tab[1], "*", "", -1)
		nodes[i] = strings.Join(tab, "\t")
		ans[i] = MorphemeString(nodes[i])
	}
	return ans
}

func (m *Markov) Generate(maxNodes int, isTerminal func(Morpheme) bool) Phrase {
	if len(m.Data.Chain) == 0 {
		return Phrase{}
	}
	ans := make(Phrase, 0, maxNodes)
	prefix := m.Data.FirstPrefix()
	for i := 0; i < maxNodes; i++ {
		candidate := m.Data.Chain[prefix.String()]
		nextNode := candidate[rand.Intn(len(candidate))]
		if nextNode == EOS {
			break
		}
		nextMorpheme := nextNode.Morpheme()
		ans = append(ans, nextMorpheme)
		if isTerminal(nextMorpheme) {
			break
		}
		prefix.Shift(nextNode)
	}
	return ans
}
