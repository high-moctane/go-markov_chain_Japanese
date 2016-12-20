package markov

import (
	"errors"
	"math/rand"
	"reflect"
	"strings"
	"time"

	mecab "github.com/shogo82148/go-mecab"
)

const (
	BOS = "\tBOS,,,,,,,,"
	EOS = "\tEOS,,,,,,,,"
)

type Markov struct {
	order int
	mecab mecab.MeCab
	chain map[string][]string
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
		order: order,
		mecab: tagger,
		chain: map[string][]string{},
	}, nil
}

func (m Markov) initialPrefix() []string {
	prefix := make([]string, m.order)
	for i, _ := range prefix {
		prefix[i] = BOS
	}
	return prefix
}

func shiftPrefix(prefix []string, node string) {
	copy(prefix, prefix[1:])
	prefix[len(prefix)-1] = node

}

func (m Markov) Add(s string) error {
	parsedStr, err := m.mecab.Parse(s)
	if err != nil {
		return err
	}
	nodes := strToStrNodes(parsedStr)
	prefix := m.initialPrefix()
	for _, node := range nodes {
		key := strings.Join(prefix, "\n")
		m.chain[key] = append(m.chain[key], node)
		shiftPrefix(prefix, node)
	}
	key := strings.Join(prefix, "\n")
	m.chain[key] = append(m.chain[key], EOS)
	return nil
}

func strToStrNodes(s string) []string {
	ans := strings.Split(s, "\n")
	ans = ans[:len(ans)-2]
	return ans
}

func (m Markov) Generate(maxNodes int) [][]string {
	ans := make([][]string, 0, maxNodes)
	prefix := m.initialPrefix()
	for i := 0; i < maxNodes; i++ {
		candidate := m.chain[strings.Join(prefix, "\n")]
		nextNode := candidate[rand.Intn(len(candidate))]
		if reflect.DeepEqual(nextNode, EOS) {
			break
		}
		ans = append(ans, strNodeToSlice(nextNode))
		shiftPrefix(prefix, nextNode)
	}
	return ans
}

func strNodeToSlice(s string) []string {
	var ans = make([]string, 10)
	tab := strings.Split(s, "\t")
	ans[0] = tab[0]
	for i, v := range strings.Split(tab[1], ",") {
		if v == "*" {
			ans[i+1] = ""
		} else {
			ans[i+1] = v
		}
	}
	return ans
}
