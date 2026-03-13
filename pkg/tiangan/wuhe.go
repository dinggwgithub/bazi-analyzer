package tiangan

import "bazi-analyzer/pkg/wuxing"

type WuheResult struct {
	Tiangan1  Tiangan
	Tiangan2  Tiangan
	HuaWuxing wuxing.Wuxing
}

var WuhePairs = []WuheResult{
	{Jia, Ji, wuxing.Tu},
	{Yi, Geng, wuxing.Jin},
	{Bing, Xin, wuxing.Shui},
	{Ding, Ren, wuxing.Mu},
	{Wu, Gui, wuxing.Huo},
}

func GetWuhe(t1, t2 Tiangan) *WuheResult {
	for _, pair := range WuhePairs {
		if (pair.Tiangan1 == t1 && pair.Tiangan2 == t2) ||
			(pair.Tiangan1 == t2 && pair.Tiangan2 == t1) {
			return &pair
		}
	}
	return nil
}

func IsWuhe(t1, t2 Tiangan) bool {
	return GetWuhe(t1, t2) != nil
}

func CanWuhua(t1, t2 Tiangan, monthWuxing wuxing.Wuxing, hasTougan bool) bool {
	result := GetWuhe(t1, t2)
	if result == nil {
		return false
	}
	if monthWuxing == result.HuaWuxing {
		return true
	}
	if hasTougan {
		return true
	}
	return false
}
