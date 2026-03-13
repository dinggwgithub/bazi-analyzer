package dizhi

type PoPair struct {
	Dizhi1 Dizhi
	Dizhi2 Dizhi
}

var PoPairs = []PoPair{
	{Zi, You},
	{Yin, Hai},
	{Chen, Chou},
	{Wu, Mao},
	{Shen, Si},
	{Xu, Wei},
}

func FindPo(dizhis []Dizhi) []PoPair {
	var results []PoPair
	for _, pair := range PoPairs {
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

func GetPo(d Dizhi) *Dizhi {
	for _, pair := range PoPairs {
		if pair.Dizhi1 == d {
			return &pair.Dizhi2
		}
		if pair.Dizhi2 == d {
			return &pair.Dizhi1
		}
	}
	return nil
}

func IsPo(d1, d2 Dizhi) bool {
	for _, pair := range PoPairs {
		if (pair.Dizhi1 == d1 && pair.Dizhi2 == d2) ||
			(pair.Dizhi1 == d2 && pair.Dizhi2 == d1) {
			return true
		}
	}
	return false
}
