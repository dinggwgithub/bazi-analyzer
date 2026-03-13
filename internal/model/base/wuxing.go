package base

const (
	Wood  = "木"
	Fire  = "火"
	Earth = "土"
	Metal = "金"
	Water = "水"
)

type WuXing string

func (w WuXing) String() string {
	return string(w)
}

var WuXingSheng = map[WuXing]WuXing{
	Wood:  Fire,
	Fire:  Earth,
	Earth: Metal,
	Metal: Water,
	Water: Wood,
}

var WuXingKe = map[WuXing]WuXing{
	Wood:  Earth,
	Earth: Water,
	Water: Fire,
	Fire:  Metal,
	Metal: Wood,
}

func Sheng(from, to WuXing) bool {
	return WuXingSheng[from] == to
}

func Ke(from, to WuXing) bool {
	return WuXingKe[from] == to
}
