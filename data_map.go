package markov

import (
	"math/rand"
	"sync"
	"time"

	"github.com/high-moctane/go-mecab_slice"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

type DataMap struct {
	Order       int
	FirstPrefix prefix
	Chain       map[string][]mecabs.MorphemeString
	mutex       *sync.Mutex
}

type prefix mecabs.PhraseString

func (p prefix) shift(ms mecabs.MorphemeString) {
	copy(p, p[1:])
	p[len(p)-1] = ms
}

func (p prefix) string() (ans string) {
	for i, ms := range p {
		if i > 0 {
			ans += "\n"
		}
		ans += string(ms)
	}
	return
}

func NewDataMap(order int) DataMap {
	return DataMap{
		Order:       order,
		FirstPrefix: genFirstPrefix(order),
		Chain:       map[string][]mecabs.MorphemeString{},
		mutex:       new(sync.Mutex),
	}
}

func genFirstPrefix(order int) prefix {
	p := make(prefix, order)
	for i, _ := range p {
		p[i] = mecabs.BOMS
	}
	return p
}

func (d *DataMap) firstPrefix() prefix {
	ans := make(prefix, len(d.FirstPrefix))
	copy(ans, d.FirstPrefix)
	return ans
}

func (d *DataMap) Add(ps mecabs.PhraseString) {
	var ms mecabs.MorphemeString
	prefix := d.firstPrefix()
	for _, ms = range ps {
		d.mutex.Lock()
		d.Chain[prefix.string()] = append(d.Chain[prefix.string()], ms)
		d.mutex.Unlock()
		prefix.shift(ms)
	}
	d.mutex.Lock()
	d.Chain[prefix.string()] = append(d.Chain[prefix.string()], mecabs.EOMS)
	d.mutex.Unlock()
}

func (d *DataMap) Generate(length int) mecabs.PhraseString {
again:
	ans := make(mecabs.PhraseString, 0, length)
	prefix := d.firstPrefix()
	for i := 0; i < length; i++ {
		candidate, ok := d.Chain[prefix.string()]
		if !ok {
			/*
				NOTE:
				本来ならcandidateが存在しないということはないが、
				go-rapbotでのバグを防ぐために仕方なく導入した
			*/
			goto again
		}
		next := candidate[rand.Intn(len(candidate))]
		if next == mecabs.EOMS {
			return ans
		}
		ans = append(ans, next)
		prefix.shift(next)
	}
	return ans
}
