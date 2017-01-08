package markov

import (
	"reflect"
	"testing"

	mecab "github.com/shogo82148/go-mecab"
)

func TestPhrase(t *testing.T) {
	mecab, _ := mecab.New(map[string]string{})
	defer mecab.Destroy()
	var phraseString PhraseString
	var parsed string
	var expected Phrase

	parsed, _ = mecab.Parse("こんにちは世界")
	phraseString = MakePhraseString(parsed)
	expected = Phrase{
		{
			OriginalForm:         "こんにちは",
			PartOfSpeech:         "感動詞",
			PartOfSpeechSection1: "",
			PartOfSpeechSection2: "",
			PartOfSpeechSection3: "",
			ConjugatedForm1:      "",
			ConjugatedForm2:      "",
			Inflection:           "こんにちは",
			Reading:              "コンニチハ",
			Pronounciation:       "コンニチワ",
		},
		{
			OriginalForm:         "世界",
			PartOfSpeech:         "名詞",
			PartOfSpeechSection1: "一般",
			PartOfSpeechSection2: "",
			PartOfSpeechSection3: "",
			ConjugatedForm1:      "",
			ConjugatedForm2:      "",
			Inflection:           "世界",
			Reading:              "セカイ",
			Pronounciation:       "セカイ",
		},
	}
	if !reflect.DeepEqual(phraseString.Phrase(), expected) {
		t.Errorf("expected %v, but %v", expected, phraseString.Phrase())
	}
}
