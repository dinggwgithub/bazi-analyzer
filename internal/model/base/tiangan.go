package base

type TianGan struct {
	Name    string
	Element WuXing
	IsYang  bool
}

func (t TianGan) String() string {
	return t.Name
}

var (
	Jia = TianGan{Name: "甲", Element: Wood, IsYang: true}
	Yi  = TianGan{Name: "乙", Element: Wood, IsYang: false}
	Bing = TianGan{Name: "丙", Element: Fire, IsYang: true}
	Ding = TianGan{Name: "丁", Element: Fire, IsYang: false}
	Wu  = TianGan{Name: "戊", Element: Earth, IsYang: true}
	Ji  = TianGan{Name: "己", Element: Earth, IsYang: false}
	Geng = TianGan{Name: "庚", Element: Metal, IsYang: true}
	Xin = TianGan{Name: "辛", Element: Metal, IsYang: false}
	Ren = TianGan{Name: "壬", Element: Water, IsYang: true}
	Gui = TianGan{Name: "癸", Element: Water, IsYang: false}
)

var TianGanMap = map[string]TianGan{
	"甲": Jia,
	"乙": Yi,
	"丙": Bing,
	"丁": Ding,
	"戊": Wu,
	"己": Ji,
	"庚": Geng,
	"辛": Xin,
	"壬": Ren,
	"癸": Gui,
}

func GetTianGan(name string) (TianGan, bool) {
	tg, ok := TianGanMap[name]
	return tg, ok
}

func GetAllTianGan() []TianGan {
	return []TianGan{Jia, Yi, Bing, Ding, Wu, Ji, Geng, Xin, Ren, Gui}
}
