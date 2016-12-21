package markov

import (
	"math/rand"
	"reflect"
	"strings"
	"testing"
)

func TestInitialPrefix(t *testing.T) {
	markov, err := New(2, map[string]string{})
	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}
	defer markov.Destroy()

	prefix := markov.initialPrefix()
	expected := []string{"\tBOS,,,,,,,,", "\tBOS,,,,,,,,"}
	if !reflect.DeepEqual(prefix, expected) {
		t.Errorf("expected %v, but %v.", expected, prefix)
	}
}

func TestAdd(t *testing.T) {
	markov, err := New(2, map[string]string{})
	if err != nil {
		t.Errorf("unexpected error: %v", err)
		return
	}
	defer markov.Destroy()

	err = markov.Add("こんにちは世界")
	if err != nil {
		t.Errorf("unexpected error: %v", err)
		return
	}
	err = markov.Add("こんにちは宇宙")
	if err != nil {
		t.Errorf("unexpected error: %v", err)
		return
	}
	expected := map[string][]string{
		"\tBOS,,,,,,,,\n\tBOS,,,,,,,,": []string{"こんにちは	感動詞,,,,,,こんにちは,コンニチハ,コンニチワ", "こんにちは	感動詞,,,,,,こんにちは,コンニチハ,コンニチワ"},
		"\tBOS,,,,,,,,\nこんにちは	感動詞,,,,,,こんにちは,コンニチハ,コンニチワ": []string{"世界	名詞,一般,,,,,世界,セカイ,セカイ", "宇宙	名詞,一般,,,,,宇宙,ウチュウ,ウチュー"},
		"こんにちは	感動詞,,,,,,こんにちは,コンニチハ,コンニチワ\n世界	名詞,一般,,,,,世界,セカイ,セカイ": []string{"\tEOS,,,,,,,,"},
		"こんにちは	感動詞,,,,,,こんにちは,コンニチハ,コンニチワ\n宇宙	名詞,一般,,,,,宇宙,ウチュウ,ウチュー": []string{"\tEOS,,,,,,,,"},
	}
	if !reflect.DeepEqual(markov.chain, expected) {
		t.Errorf("expected\n%v, but\n%v.", expected, markov)
	}
	for k, v := range markov.chain {
		if !reflect.DeepEqual(expected[k], v) {
			t.Errorf("expected\n%v, but\n%v.", expected, markov.chain)
			return
		}
	}
}

func TestGenerate(t *testing.T) {
	rand.Seed(0)

	markov, err := New(2, map[string]string{})
	if err != nil {
		t.Errorf("unexpected error: %v", err)
		return
	}
	defer markov.Destroy()

	everFalse := func([]string) bool {
		return false
	}

	markov.Generate(100, everFalse)

	err = markov.Add("こんにちは世界")
	if err != nil {
		t.Errorf("unexpected error: %v", err)
		return
	}
	err = markov.Add("こんにちは宇宙")
	if err != nil {
		t.Errorf("unexpected error: %v", err)
		return
	}

	var ans []string
	for _, v := range markov.Generate(100, everFalse) {
		ans = append(ans, v[0])
	}
	expected := "こんにちは世界"
	if strings.Join(ans, "") != expected {
		t.Errorf("expected %v, but %v", expected, ans)
		return
	}

	endCondition := func(s []string) bool {
		if s[0] == "こんにちは" {
			return true
		}
		return false
	}

	ans = []string{}
	for _, v := range markov.Generate(100, endCondition) {
		ans = append(ans, v[0])
	}
	expected = "こんにちは"
	if strings.Join(ans, "") != expected {
		t.Errorf("expected %v, but %v", expected, ans)
		return
	}
}
