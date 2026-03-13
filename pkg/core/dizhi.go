package core

// DiZhiInfo 地支详细信息
type DiZhiInfo struct {
	Name       string  // 名称
	WuXing     WuXing  // 五行属性
	YinYang    YinYang // 阴阳属性
	Index      int     // 索引
	BenQi      TianGan // 本气（最强）
	ZhongQi    TianGan // 中气（次之）
	YuQi       TianGan // 余气（最弱）
	HasZhongQi bool    // 是否有中气
	HasYuQi    bool    // 是否有余气
}

// DiZhiTable 地支信息表
var DiZhiTable = []DiZhiInfo{
	{Name: "子", WuXing: Water, YinYang: Yang, Index: 0, BenQi: Gui},                                                            // 子
	{Name: "丑", WuXing: Earth, YinYang: Yin, Index: 1, BenQi: Ji, ZhongQi: Gui, YuQi: Xin, HasZhongQi: true, HasYuQi: true},    // 丑
	{Name: "寅", WuXing: Wood, YinYang: Yang, Index: 2, BenQi: Jia, ZhongQi: Bing, YuQi: Wu, HasZhongQi: true, HasYuQi: true},   // 寅 (YinZhi)
	{Name: "卯", WuXing: Wood, YinYang: Yin, Index: 3, BenQi: Yi},                                                               // 卯
	{Name: "辰", WuXing: Earth, YinYang: Yang, Index: 4, BenQi: Wu, ZhongQi: Yi, YuQi: Gui, HasZhongQi: true, HasYuQi: true},    // 辰
	{Name: "巳", WuXing: Fire, YinYang: Yin, Index: 5, BenQi: Bing, ZhongQi: Wu, YuQi: Geng, HasZhongQi: true, HasYuQi: true},   // 巳
	{Name: "午", WuXing: Fire, YinYang: Yang, Index: 6, BenQi: Ding, ZhongQi: Ji, HasZhongQi: true},                             // 午
	{Name: "未", WuXing: Earth, YinYang: Yin, Index: 7, BenQi: Ji, ZhongQi: Ding, YuQi: Yi, HasZhongQi: true, HasYuQi: true},    // 未
	{Name: "申", WuXing: Metal, YinYang: Yang, Index: 8, BenQi: Geng, ZhongQi: Ren, YuQi: Wu, HasZhongQi: true, HasYuQi: true},  // 申
	{Name: "酉", WuXing: Metal, YinYang: Yin, Index: 9, BenQi: Xin},                                                             // 酉
	{Name: "戌", WuXing: Earth, YinYang: Yang, Index: 10, BenQi: Wu, ZhongQi: Xin, YuQi: Ding, HasZhongQi: true, HasYuQi: true}, // 戌
	{Name: "亥", WuXing: Water, YinYang: Yin, Index: 11, BenQi: Ren, ZhongQi: Jia, HasZhongQi: true},                            // 亥
}

// GetDiZhiInfo 获取地支信息
func GetDiZhiInfo(zhi DiZhi) DiZhiInfo {
	if zhi < 0 || zhi > 11 {
		return DiZhiInfo{}
	}
	return DiZhiTable[zhi]
}

// GetDiZhiWuXing 获取地支五行
func GetDiZhiWuXing(zhi DiZhi) WuXing {
	return GetDiZhiInfo(zhi).WuXing
}

// GetDiZhiYinYang 获取地支阴阳
func GetDiZhiYinYang(zhi DiZhi) YinYang {
	return GetDiZhiInfo(zhi).YinYang
}

// IsDiZhiYang 判断地支是否为阳
func IsDiZhiYang(zhi DiZhi) bool {
	return GetDiZhiYinYang(zhi) == Yang
}

// IsDiZhiYin 判断地支是否为阴
func IsDiZhiYin(zhi DiZhi) bool {
	return GetDiZhiYinYang(zhi) == Yin
}

// GetDiZhiByName 根据名称获取地支
func GetDiZhiByName(name string) (DiZhi, bool) {
	for i, info := range DiZhiTable {
		if info.Name == name {
			return DiZhi(i), true
		}
	}
	return -1, false
}

// GetDiZhiIndex 获取地支索引
func GetDiZhiIndex(zhi DiZhi) int {
	return GetDiZhiInfo(zhi).Index
}

// GetCangGan 获取地支藏干（本气、中气、余气）
func GetCangGan(zhi DiZhi) []TianGan {
	info := GetDiZhiInfo(zhi)
	var result []TianGan
	result = append(result, info.BenQi)
	if info.HasZhongQi {
		result = append(result, info.ZhongQi)
	}
	if info.HasYuQi {
		result = append(result, info.YuQi)
	}
	return result
}

// GetBenQi 获取地支本气
func GetBenQi(zhi DiZhi) TianGan {
	return GetDiZhiInfo(zhi).BenQi
}

// GetZhongQi 获取地支中气
func GetZhongQi(zhi DiZhi) TianGan {
	return GetDiZhiInfo(zhi).ZhongQi
}

// GetYuQi 获取地支余气
func GetYuQi(zhi DiZhi) TianGan {
	return GetDiZhiInfo(zhi).YuQi
}

// HasCangGan 判断地支是否包含某天干
func HasCangGan(zhi DiZhi, gan TianGan) bool {
	cangGan := GetCangGan(zhi)
	for _, g := range cangGan {
		if g == gan {
			return true
		}
	}
	return false
}

// GetAllDiZhi 获取所有地支
func GetAllDiZhi() []DiZhi {
	return []DiZhi{Zi, Chou, YinZhi, Mao, Chen, Si, WuZhi, Wei, Shen, You, Xu, Hai}
}

// DiZhiSheng 获取地支所生之支（按五行生克）
// 返回与所生五行相同的地支列表
func DiZhiSheng(zhi DiZhi) []DiZhi {
	wx := GetDiZhiWuXing(zhi)
	shengWx := GetSheng(wx)

	var result []DiZhi
	for i, info := range DiZhiTable {
		if info.WuXing == shengWx {
			result = append(result, DiZhi(i))
		}
	}
	return result
}

// DiZhiKe 获取地支所克之支（按五行生克）
// 返回与所克五行相同的地支列表
func DiZhiKe(zhi DiZhi) []DiZhi {
	wx := GetDiZhiWuXing(zhi)
	keWx := GetKe(wx)

	var result []DiZhi
	for i, info := range DiZhiTable {
		if info.WuXing == keWx {
			result = append(result, DiZhi(i))
		}
	}
	return result
}

// IsChangSheng 判断地支是否为某五行的长生位
// 木长生于亥，火长生于寅，金长生于巳，水长生于申
func IsChangSheng(zhi DiZhi, wx WuXing) bool {
	switch wx {
	case Wood:
		return zhi == Hai
	case Fire:
		return zhi == YinZhi
	case Metal:
		return zhi == Si
	case Water:
		return zhi == Shen
	default:
		return false
	}
}

// IsDiWang 判断地支是否为某五行的帝旺位
// 木旺于卯，火旺于午，金旺于酉，水旺于子
func IsDiWang(zhi DiZhi, wx WuXing) bool {
	switch wx {
	case Wood:
		return zhi == Mao
	case Fire:
		return zhi == WuZhi
	case Metal:
		return zhi == You
	case Water:
		return zhi == Zi
	default:
		return false
	}
}

// IsMuKu 判断地支是否为某五行的墓库位
// 木库在未，火库在戌，金库在丑，水库在辰
func IsMuKu(zhi DiZhi, wx WuXing) bool {
	switch wx {
	case Wood:
		return zhi == Wei
	case Fire:
		return zhi == Xu
	case Metal:
		return zhi == Chou
	case Water:
		return zhi == Chen
	default:
		return false
	}
}

// IsTongXingGen 判断地支是否为某五行的根（藏干中有该五行）
func IsTongXingGen(zhi DiZhi, wx WuXing) bool {
	cangGan := GetCangGan(zhi)
	for _, gan := range cangGan {
		if GetTianGanWuXing(gan) == wx {
			return true
		}
	}
	return false
}

// GetGenQiLevel 获取地支对某五行的根气级别
// 0: 无根, 1: 余气, 2: 中气, 3: 本气, 4: 禄旺, 5: 长生
func GetGenQiLevel(zhi DiZhi, wx WuXing) int {
	info := GetDiZhiInfo(zhi)

	// 检查本气
	if GetTianGanWuXing(info.BenQi) == wx {
		// 检查是否为禄旺位
		if IsDiWang(zhi, wx) {
			return 4
		}
		return 3
	}

	// 检查中气
	if info.HasZhongQi && GetTianGanWuXing(info.ZhongQi) == wx {
		return 2
	}

	// 检查余气
	if info.HasYuQi && GetTianGanWuXing(info.YuQi) == wx {
		return 1
	}

	// 检查长生位
	if IsChangSheng(zhi, wx) {
		return 5
	}

	return 0
}
