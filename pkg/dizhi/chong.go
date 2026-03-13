package dizhi

type ChongPair struct {
	Dizhi1 Dizhi
	Dizhi2 Dizhi
}

var ChongPairs = []ChongPair{
	{Zi, Wu},
	{Chou, Wei},
	{Yin, Shen},
	{Mao, You},
	{Chen, Xu},
	{Si, Hai},
}

func FindChong(dizhis []Dizhi) []ChongPair {
	var results []ChongPair
	for _, pair := range ChongPairs {
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

func GetChong(d Dizhi) *Dizhi {
	for _, pair := range ChongPairs {
		if pair.Dizhi1 == d {
			return &pair.Dizhi2
		}
		if pair.Dizhi2 == d {
			return &pair.Dizhi1
		}
	}
	return nil
}

func IsChong(d1, d2 Dizhi) bool {
	for _, pair := range ChongPairs {
		if (pair.Dizhi1 == d1 && pair.Dizhi2 == d2) ||
			(pair.Dizhi1 == d2 && pair.Dizhi2 == d1) {
			return true
		}
	}
	return false
}
