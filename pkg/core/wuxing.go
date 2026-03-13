package core

// WuXingRelation 五行关系类型
type WuXingRelation int

const (
	Sheng WuXingRelation = iota // 生
	Ke                          // 克
	Tong                        // 同（比和）
)

func (wr WuXingRelation) String() string {
	switch wr {
	case Sheng:
		return "生"
	case Ke:
		return "克"
	case Tong:
		return "同"
	default:
		return "未知"
	}
}

// WuXingMatrix 五行生克关系矩阵
// 行表示"我"，列表示"对方"
// 例如: WuXingMatrix[Wood][Fire] = Sheng 表示木生火
var WuXingMatrix = [5][5]WuXingRelation{
	// 木    火    土    金    水
	{Tong, Sheng, Ke, Tong, Tong}, // 木
	{Tong, Tong, Sheng, Ke, Tong}, // 火
	{Tong, Tong, Tong, Sheng, Ke}, // 土
	{Ke, Tong, Tong, Tong, Sheng}, // 金
	{Sheng, Ke, Tong, Tong, Tong}, // 水
}

// GetRelation 获取两个五行之间的关系
func GetRelation(from, to WuXing) WuXingRelation {
	if from < 0 || from > 4 || to < 0 || to > 4 {
		return Tong
	}
	return WuXingMatrix[from][to]
}

// IsSheng 判断 from 是否生 to
func IsSheng(from, to WuXing) bool {
	return GetRelation(from, to) == Sheng
}

// IsKe 判断 from 是否克 to
func IsKe(from, to WuXing) bool {
	return GetRelation(from, to) == Ke
}

// IsTong 判断两个五行是否相同
func IsTong(from, to WuXing) bool {
	return from == to
}

// GetSheng 获取生我的五行
// 木生火，火生土，土生金，金生水，水生木
func GetShengWo(wx WuXing) WuXing {
	switch wx {
	case Wood:
		return Water
	case Fire:
		return Wood
	case Earth:
		return Fire
	case Metal:
		return Earth
	case Water:
		return Metal
	default:
		return wx
	}
}

// GetWoSheng 获取我生的五行
// 木生火，火生土，土生金，金生水，水生木
func GetWoSheng(wx WuXing) WuXing {
	switch wx {
	case Wood:
		return Fire
	case Fire:
		return Earth
	case Earth:
		return Metal
	case Metal:
		return Water
	case Water:
		return Wood
	default:
		return wx
	}
}

// GetKeWo 获取克我的五行
// 金克木，木克土，土克水，水克火，火克金
func GetKeWo(wx WuXing) WuXing {
	switch wx {
	case Wood:
		return Metal
	case Fire:
		return Water
	case Earth:
		return Wood
	case Metal:
		return Fire
	case Water:
		return Earth
	default:
		return wx
	}
}

// GetWoKe 获取我克的五行
// 金克木，木克土，土克水，水克火，火克金
func GetWoKe(wx WuXing) WuXing {
	switch wx {
	case Wood:
		return Earth
	case Fire:
		return Metal
	case Earth:
		return Water
	case Metal:
		return Wood
	case Water:
		return Fire
	default:
		return wx
	}
}

// GetSheng 获取所生五行（同GetWoSheng）
func GetSheng(wx WuXing) WuXing {
	return GetWoSheng(wx)
}

// GetKe 获取所克五行（同GetWoKe）
func GetKe(wx WuXing) WuXing {
	return GetWoKe(wx)
}

// IsShengWo 判断某五行是否生我
func IsShengWo(wx, other WuXing) bool {
	return IsSheng(other, wx)
}

// IsWoSheng 判断是否我生某五行
func IsWoSheng(wx, other WuXing) bool {
	return IsSheng(wx, other)
}

// IsKeWo 判断某五行是否克我
func IsKeWo(wx, other WuXing) bool {
	return IsKe(other, wx)
}

// IsWoKe 判断是否我克某五行
func IsWoKe(wx, other WuXing) bool {
	return IsKe(wx, other)
}

// GetXiangShengChain 获取相生链
// 返回五行相生顺序: 木->火->土->金->水->木
func GetXiangShengChain() []WuXing {
	return []WuXing{Wood, Fire, Earth, Metal, Water}
}

// GetXiangKeChain 获取相克链
// 返回五行相克顺序: 木->土->水->火->金->木
func GetXiangKeChain() []WuXing {
	return []WuXing{Wood, Earth, Water, Fire, Metal}
}

// GetAllWuXing 获取所有五行
func GetAllWuXing() []WuXing {
	return []WuXing{Wood, Fire, Earth, Metal, Water}
}

// WuXingScore 五行分数计算（用于旺衰分析）
type WuXingScore struct {
	WuXing WuXing
	Score  int
}

// CalculateWuXingScore 计算八字中各五行的分数
// 用于分析五行力量的分布
func CalculateWuXingScore(bazi *Bazi) []WuXingScore {
	scores := make(map[WuXing]int)

	// 天干五行分数（权重：天干 = 2分）
	scores[GetTianGanWuXing(bazi.Year.Gan)] += 2
	scores[GetTianGanWuXing(bazi.Month.Gan)] += 2
	scores[GetTianGanWuXing(bazi.Day.Gan)] += 2
	scores[GetTianGanWuXing(bazi.Hour.Gan)] += 2

	// 地支五行分数（权重：地支本气 = 3分，中气 = 1.5分，余气 = 0.5分）
	for _, zhi := range []DiZhi{bazi.Year.Zhi, bazi.Month.Zhi, bazi.Day.Zhi, bazi.Hour.Zhi} {
		info := GetDiZhiInfo(zhi)
		scores[GetTianGanWuXing(info.BenQi)] += 3
		if info.ZhongQi >= 0 {
			scores[GetTianGanWuXing(info.ZhongQi)] += 1
		}
		if info.YuQi >= 0 {
			scores[GetTianGanWuXing(info.YuQi)] += 1
		}
	}

	// 转换为切片
	var result []WuXingScore
	for wx, score := range scores {
		result = append(result, WuXingScore{WuXing: wx, Score: score})
	}

	return result
}

// GetStrongestWuXing 获取八字中最强的五行
func GetStrongestWuXing(bazi *Bazi) WuXing {
	scores := CalculateWuXingScore(bazi)
	var strongest WuXing
	maxScore := -1

	for _, ws := range scores {
		if ws.Score > maxScore {
			maxScore = ws.Score
			strongest = ws.WuXing
		}
	}

	return strongest
}

// GetWeakestWuXing 获取八字中最弱的五行
func GetWeakestWuXing(bazi *Bazi) WuXing {
	scores := CalculateWuXingScore(bazi)
	var weakest WuXing
	minScore := 1000

	for _, ws := range scores {
		if ws.Score < minScore {
			minScore = ws.Score
			weakest = ws.WuXing
		}
	}

	return weakest
}
