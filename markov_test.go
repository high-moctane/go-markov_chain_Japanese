package markov

import (
	"math/rand"
	"reflect"
	"testing"

	"github.com/high-moctane/go-mecab_slice"
)

func TestAdd(t *testing.T) {
	var ma Markov
	var d DataMap
	var me mecabs.MeCabS
	var input0, input1 string
	var e map[string][]mecabs.MorphemeString

	input0 = "こんにちは世界"
	input1 = "こんにちは宇宙"
	e = map[string][]mecabs.MorphemeString{
		"\tBOS,,,,,,,,": {
			"こんにちは\t感動詞,,,,,,こんにちは,コンニチハ,コンニチワ",
			"こんにちは\t感動詞,,,,,,こんにちは,コンニチハ,コンニチワ",
		},
		"こんにちは\t感動詞,,,,,,こんにちは,コンニチハ,コンニチワ": {
			"世界\t名詞,一般,,,,,世界,セカイ,セカイ",
			"宇宙\t名詞,一般,,,,,宇宙,ウチュウ,ウチュー",
		},
		"世界\t名詞,一般,,,,,世界,セカイ,セカイ":   {"\tEOS,,,,,,,,"},
		"宇宙\t名詞,一般,,,,,宇宙,ウチュウ,ウチュー": {"\tEOS,,,,,,,,"},
	}
	me, err := mecabs.New(map[string]string{})
	if err != nil {
		t.Fatalf("unexpected error")
	}
	ma = New(&me, &d)
	d = NewDataMap(1)
	err = ma.Add(input0)
	if err != nil {
		t.Fatalf("unexpected error")
	}
	err = ma.Add(input1)
	if err != nil {
		t.Fatalf("unexpected error")
	}
	if !reflect.DeepEqual(d.Chain, e) {
		t.Errorf("expected\n%v\nbut\n%v", e, d.Chain)
	}
}

func TestGenerate(t *testing.T) {
	var ma Markov
	var d DataMap
	var me mecabs.MeCabS
	var input0, input1 string
	var e, ans mecabs.Phrase
	rand.Seed(0)

	input0 = "こんにちは世界"
	input1 = "こんにちは宇宙"
	e = mecabs.Phrase{
		{"こんにちは", "感動詞", "", "", "", "", "", "こんにちは", "コンニチハ", "コンニチワ"},
		{"世界", "名詞", "一般", "", "", "", "", "世界", "セカイ", "セカイ"},
	}
	me, err := mecabs.New(map[string]string{})
	if err != nil {
		t.Fatalf("unexpected error")
	}
	ma = New(&me, &d)
	d = NewDataMap(1)
	err = ma.Add(input0)
	if err != nil {
		t.Fatalf("unexpected error")
	}
	err = ma.Add(input1)
	if err != nil {
		t.Fatalf("unexpected error")
	}
	ans = ma.Generate(5)
	if !reflect.DeepEqual(ans, e) {
		t.Errorf("expected\n%v\nbut\n%v", e, ans)
	}
}
