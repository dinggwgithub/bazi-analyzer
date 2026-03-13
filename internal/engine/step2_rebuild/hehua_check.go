package step2_rebuild

import (
	"bazi-analyzer/internal/model/base"
)

func checkDeLing(chart *base.BaZiChart, huaShen base.WuXing) bool {
	monthLing := chart.GetMonthLing()
	return monthLing.Element == huaShen
}

func checkTouGan(chart *base.BaZiChart, huaShen base.WuXing) bool {
	for _, tg := range chart.GetAllTianGan() {
		if tg.Element == huaShen {
			return true
		}
	}
	return false
}

func CheckHeHuaSuccess(chart *base.BaZiChart, huaShen base.WuXing) bool {
	if huaShen == "" {
		return false
	}

	deLing := checkDeLing(chart, huaShen)
	touGan := checkTouGan(chart, huaShen)

	return (deLing || touGan)
}

func GetRebuiltTianGanWuXing(result *Step2RebuildResult, tg base.TianGan) base.WuXing {
	if result.IsTianGanAbsorbed(tg) {
		for _, re := range result.RebuiltTianGan {
			if re.OriginalName == tg.Name {
				return re.RebuiltWuXing
			}
		}
	}
	return tg.Element
}

func GetRebuiltDiZhiWuXing(result *Step2RebuildResult, dz base.DiZhi) base.WuXing {
	if result.IsDiZhiAbsorbed(dz) {
		for _, re := range result.RebuiltDiZhi {
			if re.OriginalName == dz.Name {
				return re.RebuiltWuXing
			}
		}
	}
	return dz.Element
}
