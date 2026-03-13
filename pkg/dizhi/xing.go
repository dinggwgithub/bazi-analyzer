package dizhi

type XingType int

const (
	XingWuli XingType = iota
	XingShishi
	XingWuen
	XingZi
)

type XingGroup struct {
	Dizhis []Dizhi
	Type   XingType
	Name   string
}

var XingGroups = []XingGroup{
	{Dizhis: []Dizhi{Zi, Mao}, Type: XingWuli, Name: "无礼之刑"},
	{Dizhis: []Dizhi{Yin, Si, Shen}, Type: XingShishi, Name: "恃势之刑"},
	{Dizhis: []Dizhi{Chou, Wei, Xu}, Type: XingWuen, Name: "无恩之刑"},
}

var ZixingDizhi = []Dizhi{Chen, Wu, You, Hai}

func FindXing(dizhis []Dizhi) []XingGroup {
	var results []XingGroup
	for _, group := range XingGroups {
		count := 0
		for _, d := range dizhis {
			for _, gd := range group.Dizhis {
				if d == gd {
					count++
					break
				}
			}
		}
		if count >= 2 {
			results = append(results, group)
		}
	}
	for _, d := range dizhis {
		for _, zd := range ZixingDizhi {
			if d == zd {
				count := 0
				for _, dd := range dizhis {
					if dd == zd {
						count++
					}
				}
				if count >= 2 {
					results = append(results, XingGroup{
						Dizhis: []Dizhi{zd, zd},
						Type:   XingZi,
						Name:   "自刑",
					})
				}
			}
		}
	}
	return results
}

func IsXing(d1, d2 Dizhi) bool {
	for _, group := range XingGroups {
		hasD1 := false
		hasD2 := false
		for _, d := range group.Dizhis {
			if d == d1 {
				hasD1 = true
			}
			if d == d2 {
				hasD2 = true
			}
		}
		if hasD1 && hasD2 {
			return true
		}
	}
	return false
}
