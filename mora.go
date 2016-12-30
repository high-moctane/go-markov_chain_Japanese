package markov

var (
	katakana = map[string]Mora{
		"ア": Mora{"", "a"}, "イ": Mora{"", "i"}, "ウ": Mora{"", "u"}, "エ": Mora{"", "e"}, "オ": Mora{"", "o"},
		"カ": Mora{"k", "a"}, "キ": Mora{"k", "i"}, "ク": Mora{"k", "u"}, "ケ": Mora{"k", "e"}, "コ": Mora{"k", "o"},
		"サ": Mora{"s", "a"}, "シ": Mora{"sh", "i"}, "ス": Mora{"s", "u"}, "セ": Mora{"s", "e"}, "ソ": Mora{"s", "o"},
		"タ": Mora{"t", "a"}, "チ": Mora{"ch", "i"}, "ツ": Mora{"ts", "u"}, "テ": Mora{"t", "e"}, "ト": Mora{"t", "o"},
		"ナ": Mora{"n", "a"}, "ニ": Mora{"n", "i"}, "ヌ": Mora{"n", "u"}, "ネ": Mora{"n", "e"}, "ノ": Mora{"n", "o"},
		"ハ": Mora{"h", "a"}, "ヒ": Mora{"h", "i"}, "フ": Mora{"f", "u"}, "ヘ": Mora{"h", "e"}, "ホ": Mora{"h", "o"},
		"マ": Mora{"m", "a"}, "ミ": Mora{"m", "i"}, "ム": Mora{"m", "u"}, "メ": Mora{"m", "e"}, "モ": Mora{"m", "o"},
		"ヤ": Mora{"y", "a"}, "ユ": Mora{"y", "u"}, "ヨ": Mora{"y", "o"},
		"ラ": Mora{"r", "a"}, "リ": Mora{"r", "i"}, "ル": Mora{"r", "u"}, "レ": Mora{"r", "e"}, "ロ": Mora{"r", "o"},
		"ワ": Mora{"w", "a"}, "ヲ": Mora{"", "o"}, "ン": Mora{"*n", "*n"},
		"ガ": Mora{"g", "a"}, "ギ": Mora{"g", "i"}, "グ": Mora{"g", "u"}, "ゲ": Mora{"g", "e"}, "ゴ": Mora{"g", "o"},
		"ザ": Mora{"z", "a"}, "ジ": Mora{"j", "i"}, "ズ": Mora{"z", "u"}, "ゼ": Mora{"z", "e"}, "ゾ": Mora{"z", "o"},
		"ダ": Mora{"d", "a"}, "ヂ": Mora{"j", "i"}, "ヅ": Mora{"z", "u"}, "デ": Mora{"d", "e"}, "ド": Mora{"d", "o"},
		"バ": Mora{"b", "a"}, "ビ": Mora{"b", "i"}, "ブ": Mora{"b", "u"}, "ベ": Mora{"b", "e"}, "ボ": Mora{"b", "o"},
		"パ": Mora{"p", "a"}, "ピ": Mora{"p", "i"}, "プ": Mora{"p", "u"}, "ペ": Mora{"p", "e"}, "ポ": Mora{"p", "o"},
		"キャ": Mora{"ky", "a"}, "キュ": Mora{"ky", "u"}, "キョ": Mora{"ky", "o"},
		"シャ": Mora{"sh", "a"}, "シュ": Mora{"sh", "u"}, "ショ": Mora{"sh", "o"},
		"チャ": Mora{"ch", "a"}, "チュ": Mora{"ch", "u"}, "チョ": Mora{"ch", "o"},
		"ニャ": Mora{"ny", "a"}, "ニュ": Mora{"ny", "u"}, "ニョ": Mora{"ny", "o"},
		"ヒャ": Mora{"hy", "a"}, "ヒュ": Mora{"hy", "u"}, "ヒョ": Mora{"hy", "o"},
		"ミャ": Mora{"my", "a"}, "ミュ": Mora{"my", "u"}, "ミョ": Mora{"my", "o"},
		"リャ": Mora{"ry", "a"}, "リュ": Mora{"ry", "u"}, "リョ": Mora{"ry", "o"},
		"ギャ": Mora{"gy", "a"}, "ギュ": Mora{"gy", "u"}, "ギョ": Mora{"gy", "o"},
		"ジャ": Mora{"j", "a"}, "ジュ": Mora{"j", "u"}, "ジョ": Mora{"j", "o"},
		"ビャ": Mora{"by", "a"}, "ビュ": Mora{"by", "u"}, "ビョ": Mora{"by", "o"},
		"ピャ": Mora{"py", "a"}, "ピュ": Mora{"py", "u"}, "ピョ": Mora{"py", "o"},
		"ファ": Mora{"f", "a"}, "フィ": Mora{"f", "i"}, "フェ": Mora{"f", "e"}, "フォ": Mora{"f", "o"},
		"フュ": Mora{"fy", "u"},
		"ウィ": Mora{"w", "i"}, "ウェ": Mora{"w", "e"}, "ウォ": Mora{"w", "o"},
		"ヴァ": Mora{"v", "a"}, "ヴィ": Mora{"v", "i"}, "ヴェ": Mora{"v", "e"}, "ヴォ": Mora{"v", "o"},
		"ツァ": Mora{"ts", "a"}, "ツィ": Mora{"ts", "i"}, "ツェ": Mora{"ts", "e"}, "ツォ": Mora{"ts", "o"},
		"チェ": Mora{"ch", "e"}, "シェ": Mora{"sh", "e"}, "ジェ": Mora{"j", "e"},
		"ティ": Mora{"t", "i"}, "ディ": Mora{"d", "i"},
		"デュ": Mora{"d", "u"}, "トゥ": Mora{"t", "u"},
		"ッ": Mora{"*xtu", "*xtu"},
	}
)

type Mora struct {
	consonant string
	vowel     string
}

type MoraWeight struct {
	consonant float64
	vowel     float64
}
