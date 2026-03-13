package base

type HiddenStem struct {
	TianGan TianGan
	Power   string
}

type DiZhi struct {
	Name        string
	Element     WuXing
	HiddenStems []HiddenStem
}

func (d DiZhi) String() string {
	return d.Name
}

var (
	Zi = DiZhi{
		Name:    "子",
		Element: Water,
		HiddenStems: []HiddenStem{
			{TianGan: Gui, Power: "本气"},
		},
	}
	Chou = DiZhi{
		Name:    "丑",
		Element: Earth,
		HiddenStems: []HiddenStem{
			{TianGan: Ji, Power: "本气"},
			{TianGan: Gui, Power: "中气"},
			{TianGan: Xin, Power: "余气"},
		},
	}
	Yin = DiZhi{
		Name:    "寅",
		Element: Wood,
		HiddenStems: []HiddenStem{
			{TianGan: Jia, Power: "本气"},
			{TianGan: Bing, Power: "中气"},
			{TianGan: Wu, Power: "余气"},
		},
	}
	Mao = DiZhi{
		Name:    "卯",
		Element: Wood,
		HiddenStems: []HiddenStem{
			{TianGan: Yi, Power: "本气"},
		},
	}
	Chen = DiZhi{
		Name:    "辰",
		Element: Earth,
		HiddenStems: []HiddenStem{
			{TianGan: Wu, Power: "本气"},
			{TianGan: Yi, Power: "中气"},
			{TianGan: Gui, Power: "余气"},
		},
	}
	Si = DiZhi{
		Name:    "巳",
		Element: Fire,
		HiddenStems: []HiddenStem{
			{TianGan: Bing, Power: "本气"},
			{TianGan: Geng, Power: "中气"},
			{TianGan: Wu, Power: "余气"},
		},
	}
	WuDiZhi = DiZhi{
		Name:    "午",
		Element: Fire,
		HiddenStems: []HiddenStem{
			{TianGan: Ding, Power: "本气"},
			{TianGan: Ji, Power: "中气"},
		},
	}
	Wei = DiZhi{
		Name:    "未",
		Element: Earth,
		HiddenStems: []HiddenStem{
			{TianGan: Ji, Power: "本气"},
			{TianGan: Ding, Power: "中气"},
			{TianGan: Yi, Power: "余气"},
		},
	}
	Shen = DiZhi{
		Name:    "申",
		Element: Metal,
		HiddenStems: []HiddenStem{
			{TianGan: Geng, Power: "本气"},
			{TianGan: Ren, Power: "中气"},
			{TianGan: Wu, Power: "余气"},
		},
	}
	You = DiZhi{
		Name:    "酉",
		Element: Metal,
		HiddenStems: []HiddenStem{
			{TianGan: Xin, Power: "本气"},
		},
	}
	Xu = DiZhi{
		Name:    "戌",
		Element: Earth,
		HiddenStems: []HiddenStem{
			{TianGan: Wu, Power: "本气"},
			{TianGan: Xin, Power: "中气"},
			{TianGan: Ding, Power: "余气"},
		},
	}
	Hai = DiZhi{
		Name:    "亥",
		Element: Water,
		HiddenStems: []HiddenStem{
			{TianGan: Ren, Power: "本气"},
			{TianGan: Jia, Power: "中气"},
		},
	}
)

var DiZhiMap = map[string]DiZhi{
	"子": Zi,
	"丑": Chou,
	"寅": Yin,
	"卯": Mao,
	"辰": Chen,
	"巳": Si,
	"午": WuDiZhi,
	"未": Wei,
	"申": Shen,
	"酉": You,
	"戌": Xu,
	"亥": Hai,
}

func GetDiZhi(name string) (DiZhi, bool) {
	dz, ok := DiZhiMap[name]
	return dz, ok
}

func GetAllDiZhi() []DiZhi {
	return []DiZhi{Zi, Chou, Yin, Mao, Chen, Si, WuDiZhi, Wei, Shen, You, Xu, Hai}
}

func (d DiZhi) GetBenQi() TianGan {
	if len(d.HiddenStems) > 0 {
		return d.HiddenStems[0].TianGan
	}
	return TianGan{}
}
