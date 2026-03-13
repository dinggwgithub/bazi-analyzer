package dizhi

import "bazi-analyzer/pkg/wuxing"

type LiuhePair struct {
	Dizhi1    Dizhi
	Dizhi2    Dizhi
	HuaWuxing wuxing.Wuxing
}

var LiuhePairs = []LiuhePair{
	{Zi, Chou, wuxing.Tu},
	{Yin, Hai, wuxing.Mu},
	{Mao, Xu, wuxing.Huo},
	{Chen, You, wuxing.Jin},
	{Si, Shen, wuxing.Shui},
	{Wu, Wei, wuxing.Huo},
}

func FindLiuhe(dizhis []Dizhi) []LiuhePair {
	var results []LiuhePair
	for _, pair := range LiuhePairs {
		hasD1 := false
		hasD2 := false
		for _, d := range dizhis {
			if d == pair.Dizhi1 {
				hasD1 = true
			}
			if d == pair.Dizhi2 {
				hasD2 = true
			}
		}
		if hasD1 && hasD2 {
			results = append(results, pair)
		}
	}
	return results
}

func GetLiuhe(d Dizhi) *LiuhePair {
	for _, pair := range LiuhePairs {
		if pair.Dizhi1 == d || pair.Dizhi2 == d {
			return &pair
		}
	}
	return nil
}

func IsLiuhe(d1, d2 Dizhi) bool {
	for _, pair := range LiuhePairs {
		if (pair.Dizhi1 == d1 && pair.Dizhi2 == d2) ||
			(pair.Dizhi1 == d2 && pair.Dizhi2 == d1) {
			return true
		}
	}
	return false
}
