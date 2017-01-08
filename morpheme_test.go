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
