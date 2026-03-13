package tiangan

import "bazi-analyzer/pkg/wuxing"

type Tiangan int

const (
	Jia Tiangan = iota
	Yi
	Bing
	Ding
	Wu
	Ji
	Geng
	Xin
	Ren
	Gui
)

type Yinyang int

const (
	Yang Yinyang = iota
	Yin
)

func (t Tiangan) String() string {
	names := []string{"甲", "乙", "丙", "丁", "戊", "己", "庚", "辛", "壬", "癸"}
	if t >= 0 && int(t) < len(names) {
		return names[t]
	}
	return "未知"
}

func (t Tiangan) Yinyang() Yinyang {
	if t%2 == 0 {
		return Yang
	}
	return Yin
}

func (t Tiangan) Wuxing() wuxing.Wuxing {
	wuxingMap := map[Tiangan]wuxing.Wuxing{
		Jia:  wuxing.Mu,
		Yi:   wuxing.Mu,
		Bing: wuxing.Huo,
		Ding: wuxing.Huo,
		Wu:   wuxing.Tu,
		Ji:   wuxing.Tu,
		Geng: wuxing.Jin,
		Xin:  wuxing.Jin,
		Ren:  wuxing.Shui,
		Gui:  wuxing.Shui,
	}
	return wuxingMap[t]
}

func ParseTiangan(s string) Tiangan {
	tianganMap := map[string]Tiangan{
		"甲": Jia, "乙": Yi, "丙": Bing, "丁": Ding, "戊": Wu,
		"己": Ji, "庚": Geng, "辛": Xin, "壬": Ren, "癸": Gui,
	}
	return tianganMap[s]
}

func AllTiangan() []Tiangan {
	return []Tiangan{Jia, Yi, Bing, Ding, Wu, Ji, Geng, Xin, Ren, Gui}
}
