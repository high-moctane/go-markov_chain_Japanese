package markov

import (
	"math/rand"
	"reflect"
	"strings"
	"testing"

	"github.com/k0kubun/pp"
)

var sampleInput = strings.Split(`ある日の暮方の事である。一人の下人が、羅生門の下で雨やみを待っていた。広い門の下には、この男のほかに誰もいない。ただ、所々丹塗の剥げた、大きな円柱に、蟋蟀が一匹とまっている。羅生門が、朱雀大路にある以上は、この男のほかにも、雨やみをする市女笠や揉烏帽子が、もう二三人はありそうなものである。それが、この男のほかには誰もいない。何故かと云うと、この二三年、京都には、地震とか辻風とか火事とか饑饉とか云う災がつづいて起った。そこで洛中のさびれ方は一通りではない。旧記によると、仏像や仏具を打砕いて、その丹がついたり、金銀の箔がついたりした木を、路ばたにつみ重ねて、薪の料に売っていたと云う事である。洛中がその始末であるから、羅生門の修理などは、元より誰も捨てて顧る者がなかった。するとその荒れ果てたのをよい事にして、狐狸が棲すむ。盗人が棲む。とうとうしまいには、引取り手のない死人を、この門へ持って来て、棄てて行くと云う習慣さえ出来た。そこで、日の目が見えなくなると、誰でも気味を悪るがって、この門の近所へは足ぶみをしない事になってしまったのである。その代りまた鴉がどこからか、たくさん集って来た。昼間見ると、その鴉が何羽となく輪を描いて、高い鴟尾のまわりを啼きながら、飛びまわっている。ことに門の上の空が、夕焼けであかくなる時には、それが胡麻をまいたようにはっきり見えた。鴉は、勿論、門の上にある死人の肉を、啄みに来るのである。――もっとも今日は、刻限が遅いせいか、一羽も見えない。ただ、所々、崩れかかった、そうしてその崩れ目に長い草のはえた石段の上に、鴉の糞が、点々と白くこびりついているのが見える。下人は七段ある石段の一番上の段に、洗いざらした紺の襖の尻を据えて、右の頬に出来た、大きな面皰を気にしながら、ぼんやり、雨のふるのを眺めていた。`, "。")

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
	expected := map[string][]MorphemeString{
		"\tBOS,,,,,,,,\n\tBOS,,,,,,,,": []MorphemeString{"こんにちは	感動詞,,,,,,こんにちは,コンニチハ,コンニチワ", "こんにちは	感動詞,,,,,,こんにちは,コンニチハ,コンニチワ"},
		"\tBOS,,,,,,,,\nこんにちは	感動詞,,,,,,こんにちは,コンニチハ,コンニチワ": []MorphemeString{"世界	名詞,一般,,,,,世界,セカイ,セカイ", "宇宙	名詞,一般,,,,,宇宙,ウチュウ,ウチュー"},
		"こんにちは	感動詞,,,,,,こんにちは,コンニチハ,コンニチワ\n世界	名詞,一般,,,,,世界,セカイ,セカイ": []MorphemeString{"\tEOS,,,,,,,,"},
		"こんにちは	感動詞,,,,,,こんにちは,コンニチハ,コンニチワ\n宇宙	名詞,一般,,,,,宇宙,ウチュウ,ウチュー": []MorphemeString{"\tEOS,,,,,,,,"},
	}
	if !reflect.DeepEqual(markov.Data.Chain, expected) {
		t.Errorf("expected\n%sv, but\n%s.", pp.Sprint(expected), pp.Sprint(markov))
	}
	for k, v := range markov.Data.Chain {
		if !reflect.DeepEqual(expected[k], v) {
			t.Errorf("expected\n%v, but\n%v.", expected, markov.Data.Chain)
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

	everFalse := func(Morpheme) bool {
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
		ans = append(ans, v.OriginalForm)
	}
	expected := "こんにちは世界"
	if strings.Join(ans, "") != expected {
		t.Errorf("expected %v, but %v", expected, ans)
		return
	}

	endCondition := func(m Morpheme) bool {
		if m.OriginalForm == "こんにちは" {
			return true
		}
		return false
	}

	ans = []string{}
	for _, v := range markov.Generate(100, endCondition) {
		ans = append(ans, v.OriginalForm)
	}
	expected = "こんにちは"
	if strings.Join(ans, "") != expected {
		t.Errorf("expected %v, but %v", expected, ans)
		return
	}
}

func BenchmarkAdd(b *testing.B) {
	markov, err := New(2, map[string]string{})
	if err != nil {
		b.Errorf("unexpected error: %v", err)
		return
	}
	defer markov.Destroy()

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		for _, v := range sampleInput {
			markov.Add(v)
		}
	}
}

func BenchmarkGenerate(b *testing.B) {
	rand.Seed(0)

	everFalse := func(Morpheme) bool {
		return false
	}

	markov, err := New(2, map[string]string{})
	if err != nil {
		b.Errorf("unexpected error: %v", err)
		return
	}
	defer markov.Destroy()

	for _, v := range sampleInput {
		markov.Add(v)
	}

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		markov.Generate(1000, everFalse)
	}
}
