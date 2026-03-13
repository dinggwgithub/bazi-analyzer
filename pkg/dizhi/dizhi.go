package dizhi

import "bazi-analyzer/pkg/wuxing"

type Dizhi int

const (
	Zi Dizhi = iota
	Chou
	Yin
	Mao
	Chen
	Si
	Wu
	Wei
	Shen
	You
	Xu
	Hai
)

type YinyangType int

const (
	Yang YinyangType = iota
	YinType
)

func (d Dizhi) String() string {
	names := []string{"子", "丑", "寅", "卯", "辰", "巳", "午", "未", "申", "酉", "戌", "亥"}
	if d >= 0 && int(d) < len(names) {
		return names[d]
	}
	return "未知"
}

func (d Dizhi) Yinyang() YinyangType {
	if d%2 == 0 {
		return Yang
	}
	return YinType
}

func (d Dizhi) Wuxing() wuxing.Wuxing {
	wuxingMap := map[Dizhi]wuxing.Wuxing{
		Zi:   wuxing.Shui,
		Chou: wuxing.Tu,
		Yin:  wuxing.Mu,
		Mao:  wuxing.Mu,
		Chen: wuxing.Tu,
		Si:   wuxing.Huo,
		Wu:   wuxing.Huo,
		Wei:  wuxing.Tu,
		Shen: wuxing.Jin,
		You:  wuxing.Jin,
		Xu:   wuxing.Tu,
		Hai:  wuxing.Shui,
	}
	return wuxingMap[d]
}

func ParseDizhi(s string) Dizhi {
	dizhiMap := map[string]Dizhi{
		"子": Zi, "丑": Chou, "寅": Yin, "卯": Mao, "辰": Chen, "巳": Si,
		"午": Wu, "未": Wei, "申": Shen, "酉": You, "戌": Xu, "亥": Hai,
	}
	return dizhiMap[s]
}

func AllDizhi() []Dizhi {
	return []Dizhi{Zi, Chou, Yin, Mao, Chen, Si, Wu, Wei, Shen, You, Xu, Hai}
}
