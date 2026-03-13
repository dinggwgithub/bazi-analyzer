package step4_xiji

import (
	"bazi-analyzer/internal/engine/step3_wangshuai"
	"bazi-analyzer/internal/model/base"
)

type Step4XiJiResult struct {
	DayMaster          base.TianGan
	DayMasterWuXing    base.WuXing
	WangShuaiType      step3_wangshuai.WangShuaiType
	FavorableWuXing    []base.WuXing
	UnfavorableWuXing  []base.WuXing
	FavorableShiShen   []string
	UnfavorableShiShen []string
	Reason             string
}

func AnalyzeXiJi(chart *base.BaZiChart, step3Result *step3_wangshuai.Step3WangShuaiResult) *Step4XiJiResult {
	result := &Step4XiJiResult{
		DayMaster:       step3Result.DayMaster,
		DayMasterWuXing: step3Result.DayMasterWuXing,
		WangShuaiType:   step3Result.WangShuaiType,
	}

	switch step3Result.WangShuaiType {
	case step3_wangshuai.ShenRuo:
		analyzeShenRuo(result)
	case step3_wangshuai.ShenWang:
		analyzeShenWang(result)
	case step3_wangshuai.CongQiang:
		analyzeCongQiang(result)
	case step3_wangshuai.CongRuo:
		analyzeCongRuo(result)
	}

	return result
}

func analyzeShenRuo(result *Step4XiJiResult) {
	result.FavorableWuXing = []base.WuXing{
		getWuXingShengWo(result.DayMasterWuXing),
		result.DayMasterWuXing,
	}
	result.FavorableShiShen = []string{"印星", "比劫"}

	result.UnfavorableWuXing = getKeXieHaoWuXing(result.DayMasterWuXing)
	result.UnfavorableShiShen = []string{"官杀", "食伤", "财星"}

	result.Reason = "身弱喜生扶，忌克泄耗"
}

func analyzeShenWang(result *Step4XiJiResult) {
	result.FavorableWuXing = getKeXieHaoWuXing(result.DayMasterWuXing)
	result.FavorableShiShen = []string{"官杀", "食伤", "财星"}

	result.UnfavorableWuXing = []base.WuXing{
		getWuXingShengWo(result.DayMasterWuXing),
		result.DayMasterWuXing,
	}
	result.UnfavorableShiShen = []string{"印星", "比劫"}

	result.Reason = "身旺喜克泄耗，忌生扶"
}

func analyzeCongQiang(result *Step4XiJiResult) {
	result.FavorableWuXing = []base.WuXing{
		result.DayMasterWuXing,
		getWuXingShengWo(result.DayMasterWuXing),
	}
	result.FavorableShiShen = []string{"比劫", "印星"}

	result.UnfavorableWuXing = getKeXieHaoWuXing(result.DayMasterWuXing)
	result.UnfavorableShiShen = []string{"官杀", "食伤", "财星"}

	result.Reason = "从强格顺其强势，喜比劫印星，忌克泄耗"
}

func analyzeCongRuo(result *Step4XiJiResult) {
	result.FavorableWuXing = getKeXieHaoWuXing(result.DayMasterWuXing)
	result.FavorableShiShen = []string{"官杀", "食伤", "财星"}

	result.UnfavorableWuXing = []base.WuXing{
		result.DayMasterWuXing,
		getWuXingShengWo(result.DayMasterWuXing),
	}
	result.UnfavorableShiShen = []string{"比劫", "印星"}

	result.Reason = "从弱格顺其弱势，喜克泄耗，忌生扶"
}

func getWuXingShengWo(dm base.WuXing) base.WuXing {
	for k, v := range base.WuXingSheng {
		if v == dm {
			return k
		}
	}
	return ""
}

func getWuXingKeWo(dm base.WuXing) base.WuXing {
	for k, v := range base.WuXingKe {
		if v == dm {
			return k
		}
	}
	return ""
}

func getKeXieHaoWuXing(dm base.WuXing) []base.WuXing {
	return []base.WuXing{
		getWuXingKeWo(dm),
		base.WuXingSheng[dm],
		base.WuXingKe[dm],
	}
}
