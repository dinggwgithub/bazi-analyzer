package dizhi

type HaiPair struct {
	Dizhi1 Dizhi
	Dizhi2 Dizhi
}

var HaiPairs = []HaiPair{
	{Zi, Wei},
	{Chou, Wu},
	{Yin, Si},
	{Mao, Chen},
	{Shen, Hai},
	{You, Xu},
}

func FindHai(dizhis []Dizhi) []HaiPair {
	var results []HaiPair
	for _, pair := range HaiPairs {
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

func GetHai(d Dizhi) *Dizhi {
	for _, pair := range HaiPairs {
		if pair.Dizhi1 == d {
			return &pair.Dizhi2
		}
		if pair.Dizhi2 == d {
			return &pair.Dizhi1
		}
	}
	return nil
}

func IsHai(d1, d2 Dizhi) bool {
	for _, pair := range HaiPairs {
		if (pair.Dizhi1 == d1 && pair.Dizhi2 == d2) ||
			(pair.Dizhi1 == d2 && pair.Dizhi2 == d1) {
			return true
		}
	}
	return false
}
