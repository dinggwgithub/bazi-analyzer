package base

type ShiShen string

const (
	BiJian   ShiShen = "比肩"
	JieCai   ShiShen = "劫财"
	ZhengYin ShiShen = "正印"
	PianYin  ShiShen = "偏印"
	ZhengCai ShiShen = "正财"
	PianCai  ShiShen = "偏财"
	ZhengGuan ShiShen = "正官"
	PianGuan  ShiShen = "偏官"
	ShangGuan ShiShen = "伤官"
	ShiShenGod  ShiShen = "食神"
)

func GetShiShen(dayMasterWuXing WuXing, targetWuXing WuXing, isYang bool) ShiShen {
	if dayMasterWuXing == targetWuXing {
		if isYang {
			return BiJian
		}
		return JieCai
	}

	if Sheng(targetWuXing, dayMasterWuXing) {
		if isYang {
			return ZhengYin
		}
		return PianYin
	}

	if Sheng(dayMasterWuXing, targetWuXing) {
		if isYang {
			return ShangGuan
		}
		return ShiShenGod
	}

	if Ke(dayMasterWuXing, targetWuXing) {
		if isYang {
			return ZhengGuan
		}
		return PianGuan
	}

	if Ke(targetWuXing, dayMasterWuXing) {
		if isYang {
			return ZhengCai
		}
		return PianCai
	}

	return ""
}
