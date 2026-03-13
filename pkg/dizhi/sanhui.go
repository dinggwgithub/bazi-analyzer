package dizhi

import "bazi-analyzer/pkg/wuxing"

type SanhuiGroup struct {
	Dizhis   []Dizhi
	HuaWuxing wuxing.Wuxing
	Name     string
}

var SanhuiGroups = []SanhuiGroup{
	{Dizhis: []Dizhi{Yin, Mao, Chen}, HuaWuxing: wuxing.Mu, Name: "东方木局"},
	{Dizhis: []Dizhi{Si, Wu, Wei}, HuaWuxing: wuxing.Huo, Name: "南方火局"},
	{Dizhis: []Dizhi{Shen, You, Xu}, HuaWuxing: wuxing.Jin, Name: "西方金局"},
	{Dizhis: []Dizhi{Hai, Zi, Chou}, HuaWuxing: wuxing.Shui, Name: "北方水局"},
}

func FindSanhui(dizhis []Dizhi) *SanhuiGroup {
	for _, group := range SanhuiGroups {
		count := 0
		for _, d := range dizhis {
			for _, gd := range group.Dizhis {
				if d == gd {
					count++
					break
				}
			}
		}
		if count == 3 {
			return &group
		}
	}
	return nil
}

func HasSanhuiElement(d Dizhi) bool {
	for _, group := range SanhuiGroups {
		for _, gd := range group.Dizhis {
			if d == gd {
				return true
			}
		}
	}
	return false
}

func GetSanhuiGroup(d Dizhi) *SanhuiGroup {
	for _, group := range SanhuiGroups {
		for _, gd := range group.Dizhis {
			if d == gd {
				return &group
			}
		}
	}
	return nil
}
