package core

// TianGanInfo 天干详细信息
type TianGanInfo struct {
	Name    string  // 名称
	WuXing  WuXing  // 五行属性
	YinYang YinYang // 阴阳属性
	Index   int     // 索引
}

// TianGanTable 天干信息表
var TianGanTable = []TianGanInfo{
	{Name: "甲", WuXing: Wood, YinYang: Yang, Index: 0},  // 甲
	{Name: "乙", WuXing: Wood, YinYang: Yin, Index: 1},   // 乙
	{Name: "丙", WuXing: Fire, YinYang: Yang, Index: 2},  // 丙
	{Name: "丁", WuXing: Fire, YinYang: Yin, Index: 3},   // 丁
	{Name: "戊", WuXing: Earth, YinYang: Yang, Index: 4}, // 戊
	{Name: "己", WuXing: Earth, YinYang: Yin, Index: 5},  // 己
	{Name: "庚", WuXing: Metal, YinYang: Yang, Index: 6}, // 庚
	{Name: "辛", WuXing: Metal, YinYang: Yin, Index: 7},  // 辛
	{Name: "壬", WuXing: Water, YinYang: Yang, Index: 8}, // 壬
	{Name: "癸", WuXing: Water, YinYang: Yin, Index: 9},  // 癸
}

// GetTianGanInfo 获取天干信息
func GetTianGanInfo(gan TianGan) TianGanInfo {
	if gan < 0 || gan > 9 {
		return TianGanInfo{}
	}
	return TianGanTable[gan]
}

// GetTianGanWuXing 获取天干五行
func GetTianGanWuXing(gan TianGan) WuXing {
	return GetTianGanInfo(gan).WuXing
}

// GetTianGanYinYang 获取天干阴阳
func GetTianGanYinYang(gan TianGan) YinYang {
	return GetTianGanInfo(gan).YinYang
}

// IsTianGanYang 判断天干是否为阳
func IsTianGanYang(gan TianGan) bool {
	return GetTianGanYinYang(gan) == Yang
}

// IsTianGanYin 判断天干是否为阴
func IsTianGanYin(gan TianGan) bool {
	return GetTianGanYinYang(gan) == Yin
}

// GetTianGanByName 根据名称获取天干
func GetTianGanByName(name string) (TianGan, bool) {
	for i, info := range TianGanTable {
		if info.Name == name {
			return TianGan(i), true
		}
	}
	return -1, false
}

// GetTianGanIndex 获取天干索引
func GetTianGanIndex(gan TianGan) int {
	return GetTianGanInfo(gan).Index
}

// IsSameWuXing 判断两个天干是否同五行
func IsSameWuXing(gan1, gan2 TianGan) bool {
	return GetTianGanWuXing(gan1) == GetTianGanWuXing(gan2)
}

// IsSameYinYang 判断两个天干是否同阴阳
func IsSameYinYang(gan1, gan2 TianGan) bool {
	return GetTianGanYinYang(gan1) == GetTianGanYinYang(gan2)
}

// IsBiJie 判断两个天干是否为比劫关系（同五行）
func IsBiJie(gan1, gan2 TianGan) bool {
	return IsSameWuXing(gan1, gan2)
}

// TianGanSheng 获取天干所生之干（按五行生克）
// 返回与所生五行相同的天干列表
func TianGanSheng(gan TianGan) []TianGan {
	wx := GetTianGanWuXing(gan)
	shengWx := GetSheng(wx)

	var result []TianGan
	for i, info := range TianGanTable {
		if info.WuXing == shengWx {
			result = append(result, TianGan(i))
		}
	}
	return result
}

// TianGanKe 获取天干所克之干（按五行生克）
// 返回与所克五行相同的天干列表
func TianGanKe(gan TianGan) []TianGan {
	wx := GetTianGanWuXing(gan)
	keWx := GetKe(wx)

	var result []TianGan
	for i, info := range TianGanTable {
		if info.WuXing == keWx {
			result = append(result, TianGan(i))
		}
	}
	return result
}

// KeTianGan 获取克该天干之干（按五行生克）
// 返回能克此天干的五行对应的天干列表
func KeTianGan(gan TianGan) []TianGan {
	wx := GetTianGanWuXing(gan)
	// 克我者为：我的五行被什么五行克
	// 金克木，木克土，土克水，水克火，火克金
	var keWx WuXing
	switch wx {
	case Wood:
		keWx = Metal
	case Fire:
		keWx = Water
	case Earth:
		keWx = Wood
	case Metal:
		keWx = Fire
	case Water:
		keWx = Earth
	}

	var result []TianGan
	for i, info := range TianGanTable {
		if info.WuXing == keWx {
			result = append(result, TianGan(i))
		}
	}
	return result
}

// ShengTianGan 获取生该天干之干（按五行生克）
// 返回能生此天干的五行对应的天干列表
func ShengTianGan(gan TianGan) []TianGan {
	wx := GetTianGanWuXing(gan)
	// 生我者为：什么五行生我的五行
	// 木生火，火生土，土生金，金生水，水生木
	var shengWx WuXing
	switch wx {
	case Wood:
		shengWx = Water
	case Fire:
		shengWx = Wood
	case Earth:
		shengWx = Fire
	case Metal:
		shengWx = Earth
	case Water:
		shengWx = Metal
	}

	var result []TianGan
	for i, info := range TianGanTable {
		if info.WuXing == shengWx {
			result = append(result, TianGan(i))
		}
	}
	return result
}

// GetAllTianGan 获取所有天干
func GetAllTianGan() []TianGan {
	return []TianGan{Jia, Yi, Bing, Ding, Wu, Ji, Geng, Xin, Ren, Gui}
}
