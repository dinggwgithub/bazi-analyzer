package tiangan

type SichongPair struct {
	Tiangan1 Tiangan
	Tiangan2 Tiangan
}

var SichongPairs = []SichongPair{
	{Jia, Geng},
	{Yi, Xin},
	{Bing, Ren},
	{Ding, Gui},
}

func IsSichong(t1, t2 Tiangan) bool {
	for _, pair := range SichongPairs {
		if (pair.Tiangan1 == t1 && pair.Tiangan2 == t2) ||
			(pair.Tiangan1 == t2 && pair.Tiangan2 == t1) {
			return true
		}
	}
	return false
}

func GetSichong(t Tiangan) *Tiangan {
	for _, pair := range SichongPairs {
		if pair.Tiangan1 == t {
			return &pair.Tiangan2
		}
		if pair.Tiangan2 == t {
			return &pair.Tiangan1
		}
	}
	return nil
}
