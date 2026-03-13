package dizhi

import "bazi-analyzer/pkg/wuxing"

type SanheGroup struct {
	Changsheng Dizhi
	Diwang     Dizhi
	Mu         Dizhi
	HuaWuxing  wuxing.Wuxing
	Name       string
}

var SanheGroups = []SanheGroup{
	{Changsheng: Shen, Diwang: Zi, Mu: Chen, HuaWuxing: wuxing.Shui, Name: "水局"},
	{Changsheng: Hai, Diwang: Mao, Mu: Wei, HuaWuxing: wuxing.Mu, Name: "木局"},
	{Changsheng: Yin, Diwang: Wu, Mu: Xu, HuaWuxing: wuxing.Huo, Name: "火局"},
	{Changsheng: Si, Diwang: You, Mu: Chou, HuaWuxing: wuxing.Jin, Name: "金局"},
}

func FindSanhe(dizhis []Dizhi) *SanheGroup {
	for _, group := range SanheGroups {
		hasChangsheng := false
		hasDiwang := false
		hasMu := false
		for _, d := range dizhis {
			if d == group.Changsheng {
				hasChangsheng = true
			}
			if d == group.Diwang {
				hasDiwang = true
			}
			if d == group.Mu {
				hasMu = true
			}
		}
		if hasChangsheng && hasDiwang && hasMu {
			return &group
		}
	}
	return nil
}

func FindSanhePartial(dizhis []Dizhi) []*SanheGroup {
	var results []*SanheGroup
	for _, group := range SanheGroups {
		count := 0
		for _, d := range dizhis {
			if d == group.Changsheng || d == group.Diwang || d == group.Mu {
				count++
			}
		}
		if count >= 2 {
			results = append(results, &group)
		}
	}
	return results
}

func IsSanheElement(d Dizhi) bool {
	for _, group := range SanheGroups {
		if d == group.Changsheng || d == group.Diwang || d == group.Mu {
			return true
		}
	}
	return false
}

func GetSanheGroup(d Dizhi) *SanheGroup {
	for _, group := range SanheGroups {
		if d == group.Changsheng || d == group.Diwang || d == group.Mu {
			return &group
		}
	}
	return nil
}
