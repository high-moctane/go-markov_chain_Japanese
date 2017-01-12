package markov

import (
	"math/rand"
	"reflect"
	"testing"

	"github.com/high-moctane/go-mecab_slice"
)

func TestGenFirstPrefix(t *testing.T) {
	var e, ans prefix

	e = prefix{mecabs.BOMS}
	ans = genFirstPrefix(1)
	if !reflect.DeepEqual(e, ans) {
		t.Errorf("expected\n%v\nbut\n%v", e, ans)
	}

	e = prefix{mecabs.BOMS, mecabs.BOMS}
	ans = genFirstPrefix(2)
	if !reflect.DeepEqual(e, ans) {
		t.Errorf("expected\n%v\nbut\n%v", e, ans)
	}
}

func TestAdd_DataMap(t *testing.T) {
	var d DataMap
	var input0 mecabs.PhraseString
	var input1 mecabs.PhraseString
	var e map[string][]mecabs.MorphemeString

	d = NewDataMap(1)
	input0 = mecabs.PhraseString{
		"こんにちは\t感動詞,,,,,,こんにちは,コンニチハ,コンニチワ",
		"世界\t名詞,一般,,,,,世界,セカイ,セカイ",
	}
	input1 = mecabs.PhraseString{
		"こんにちは\t感動詞,,,,,,こんにちは,コンニチハ,コンニチワ",
		"宇宙\t名詞,一般,,,,,宇宙,ウチュウ,ウチュー",
	}
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
	d.Add(input0)
	d.Add(input1)
	if !reflect.DeepEqual(d.Chain, e) {
		t.Errorf("expected\n%v\nbut\n%v", e, d.Chain)
	}

	d = NewDataMap(2)
	input0 = mecabs.PhraseString{
		"こんにちは\t感動詞,,,,,,こんにちは,コンニチハ,コンニチワ",
		"世界\t名詞,一般,,,,,世界,セカイ,セカイ",
	}
	input1 = mecabs.PhraseString{
		"こんにちは\t感動詞,,,,,,こんにちは,コンニチハ,コンニチワ",
		"宇宙\t名詞,一般,,,,,宇宙,ウチュウ,ウチュー",
	}
	e = map[string][]mecabs.MorphemeString{
		"\tBOS,,,,,,,,\n\tBOS,,,,,,,,": {
			"こんにちは\t感動詞,,,,,,こんにちは,コンニチハ,コンニチワ",
			"こんにちは\t感動詞,,,,,,こんにちは,コンニチハ,コンニチワ",
		},
		"\tBOS,,,,,,,,\nこんにちは\t感動詞,,,,,,こんにちは,コンニチハ,コンニチワ": {
			"世界\t名詞,一般,,,,,世界,セカイ,セカイ",
			"宇宙\t名詞,一般,,,,,宇宙,ウチュウ,ウチュー",
		},
		"こんにちは\t感動詞,,,,,,こんにちは,コンニチハ,コンニチワ\n世界\t名詞,一般,,,,,世界,セカイ,セカイ": {
			"\tEOS,,,,,,,,",
		},
		"こんにちは\t感動詞,,,,,,こんにちは,コンニチハ,コンニチワ\n宇宙\t名詞,一般,,,,,宇宙,ウチュウ,ウチュー": {
			"\tEOS,,,,,,,,",
		},
	}
	d.Add(input0)
	d.Add(input1)
	if !reflect.DeepEqual(d.Chain, e) {
		t.Errorf("expected\n%v\nbut\n%v", e, d.Chain)
	}
}

func TestGenerate_DataMap(t *testing.T) {
	rand.Seed(0)
	var d DataMap
	var e, ans mecabs.PhraseString
	input0 := mecabs.PhraseString{
		"こんにちは\t感動詞,,,,,,こんにちは,コンニチハ,コンニチワ",
		"世界\t名詞,一般,,,,,世界,セカイ,セカイ",
	}
	input1 := mecabs.PhraseString{
		"こんにちは\t感動詞,,,,,,こんにちは,コンニチハ,コンニチワ",
		"宇宙\t名詞,一般,,,,,宇宙,ウチュウ,ウチュー",
	}

	d = NewDataMap(1)
	e = mecabs.PhraseString{
		"こんにちは\t感動詞,,,,,,こんにちは,コンニチハ,コンニチワ",
		"世界\t名詞,一般,,,,,世界,セカイ,セカイ",
	}
	d.Add(input0)
	d.Add(input1)
	ans = d.Generate(5)
	if !reflect.DeepEqual(e, ans) {
		t.Errorf("expected\n%v\nbut\n%v", e, ans)
	}

	d = NewDataMap(1)
	e = mecabs.PhraseString{
		"こんにちは\t感動詞,,,,,,こんにちは,コンニチハ,コンニチワ",
	}
	d.Add(input0)
	d.Add(input1)
	ans = d.Generate(1)
	if !reflect.DeepEqual(e, ans) {
		t.Errorf("expected\n%v\nbut\n%v", e, ans)
	}

	d = NewDataMap(2)
	e = mecabs.PhraseString{
		"こんにちは\t感動詞,,,,,,こんにちは,コンニチハ,コンニチワ",
		"世界\t名詞,一般,,,,,世界,セカイ,セカイ",
	}
	d.Add(input0)
	d.Add(input1)
	ans = d.Generate(5)
	if !reflect.DeepEqual(e, ans) {
		t.Errorf("expected\n%v\nbut\n%v", e, ans)
	}
}
