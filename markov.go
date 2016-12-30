package markov

import (
	"errors"
	"math/rand"
	"strings"
	"time"

	mecab "github.com/shogo82148/go-mecab"
)

const (
	BOS = "\tBOS,,,,,,,,"
	EOS = "\tEOS,,,,,,,,"
)

type Markov struct {
	mecab mecab.MeCab
	data  Data
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
		mecab: tagger,
		data: Data{
			MarkovArgs:  args,
			Order:       order,
			firstPrefix: firstPrefix(order),
			Chain:       map[string][]MorphemeString{},
		},
	}, nil
}

func (m *Markov) Destroy() {
	m.mecab.Destroy()
	m.data = Data{}
}

func (m Markov) Add(s string) error {
	parsedStr, err := m.mecab.Parse(s)
	if err != nil {
		return err
	}
	nodes := makeMorphemes(parsedStr)
	prefix := m.data.FirstPrefix()
	for _, node := range nodes {
		key := prefix.String()
		m.data.Chain[key] = append(m.data.Chain[key], node)
		prefix.Shift(node)
	}
	key := prefix.String()
	m.data.Chain[key] = append(m.data.Chain[key], EOS)
	return nil
}

func makeMorphemes(s string) []MorphemeString {
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

func (m *Markov) Generate(maxNodes int, isTerminal func([]string) bool) [][]string {
	if len(m.data.Chain) == 0 {
		return [][]string{}
	}
	ans := make([][]string, 0, maxNodes)
	prefix := m.data.FirstPrefix()
	for i := 0; i < maxNodes; i++ {
		candidate := m.data.Chain[prefix.String()]
		nextNode := candidate[rand.Intn(len(candidate))]
		if nextNode == EOS {
			break
		}
		nextNodeSlice := nextNode.Morpheme()
		ans = append(ans, nextNodeSlice)
		if isTerminal(nextNodeSlice) {
			break
		}
		prefix.Shift(nextNode)
	}
	return ans
}
