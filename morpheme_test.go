package markov

import (
	"reflect"
	"testing"
)

func TestNewMorphemeString(t *testing.T) {
	var input string

	input = "こんにちは	感動詞,,,,,,こんにちは,コンニチハ,コンニチワ"
	if NewMorphemeString(input) != "こんにちは	感動詞,,,,,,こんにちは,コンニチハ,コンニチワ" {
		t.Errorf("parse failed: %v", input)
	}

	input = "こんにちは	感動詞,*,*,*,*,*,こんにちは"
	if NewMorphemeString(input) != "こんにちは	感動詞,,,,,,こんにちは,," {
		t.Errorf("parse failed: %v", input)
	}
}

func TestNewMorpheme(t *testing.T) {
	var input MorphemeString
	var expect Morpheme

	input = NewMorphemeString("こんにちは	感動詞,,,,,,こんにちは,コンニチハ,コンニチワ")
	expect = Morpheme(Morpheme{"こんにちは", "感動詞", "", "", "", "", "", "こんにちは", "コンニチハ", "コンニチワ"})
	if !reflect.DeepEqual(NewMorpheme(input), expect) {
		t.Errorf("parse failed: %v", input)
	}

	input = NewMorphemeString("こんにちは	感動詞,,,,,,こんにちは")
	expect = Morpheme(Morpheme{"こんにちは", "感動詞", "", "", "", "", "", "こんにちは", "", ""})
	if !reflect.DeepEqual(NewMorpheme(input), expect) {
		t.Errorf("parse failed: %v", input)
	}
}

func TestMora(t *testing.T) {
	var input Morpheme
	var expected []Mora

	input = NewMorpheme(NewMorphemeString("	,,,,,,,,コンニチワ"))
	expected = []Mora{{"k", "o"}, {"*n", "*n"}, {"n", "i"}, {"ch", "i"}, {"w", "a"}}
	if !reflect.DeepEqual(input.Mora(), expected) {
		t.Errorf("expected %v, but %v", expected, input.Mora())
	}

	input = NewMorpheme(NewMorphemeString("	,,,,,,,,バール"))
	expected = []Mora{{"b", "a"}, {"", "a"}, {"r", "u"}}
	if !reflect.DeepEqual(input.Mora(), expected) {
		t.Errorf("expected %v, but %v", expected, input.Mora())
	}

	input = NewMorpheme(NewMorphemeString("	,,,,,,,,チェッカー"))
	expected = []Mora{{"ch", "e"}, {"*xtu", "*xtu"}, {"k", "a"}, {"", "a"}}
	if !reflect.DeepEqual(input.Mora(), expected) {
		t.Errorf("expected %v, but %v", expected, input.Mora())
	}

	input = NewMorpheme(NewMorphemeString("	,,,,,,,,シューズ"))
	expected = []Mora{{"sh", "u"}, {"", "u"}, {"z", "u"}}
	if !reflect.DeepEqual(input.Mora(), expected) {
		t.Errorf("expected %v, but %v", expected, input.Mora())
	}
}
