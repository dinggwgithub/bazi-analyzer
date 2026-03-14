package step4_xiji

import (
	"bazi-analyzer/internal/engine/step2_rebuild"
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

	BingYuanDiagnosis *BingYuanDiagnosis
	YongShenResult    *YongShenResult

	CoreBingYuan string
	YongShen     string
	YongShenType string
	YongShenDesc string
	JiShen       []string
	XiShen       []string
	AnalysisPath []string
}

func AnalyzeXiJi(chart *base.BaZiChart, step3Result *step3_wangshuai.Step3WangShuaiResult) *Step4XiJiResult {
	return AnalyzeXiJiWithRebuild(chart, nil, step3Result)
}

func AnalyzeXiJiWithRebuild(chart *base.BaZiChart, step2Result *step2_rebuild.Step2RebuildResult, step3Result *step3_wangshuai.Step3WangShuaiResult) *Step4XiJiResult {
	result := &Step4XiJiResult{
		DayMaster:       step3Result.DayMaster,
		DayMasterWuXing: step3Result.DayMasterWuXing,
		WangShuaiType:   step3Result.WangShuaiType,
		JiShen:          make([]string, 0),
		XiShen:          make([]string, 0),
		AnalysisPath:    make([]string, 0),
	}

	if step2Result == nil {
		step2Result = step2_rebuild.NewStep2RebuildResult(*chart, nil)
	}

	diagnosis := DiagnoseBingYuan(chart, step2Result, step3Result)
	result.BingYuanDiagnosis = diagnosis

	yongShenResult := DeriveYongShen(chart, step2Result, step3Result, diagnosis)
	result.YongShenResult = yongShenResult

	populateStructuredOutput(result, diagnosis, yongShenResult)

	populateLegacyFields(result, step3Result)

	return result
}

func populateStructuredOutput(result *Step4XiJiResult, diagnosis *BingYuanDiagnosis, yongShenResult *YongShenResult) {
	if diagnosis.CoreBingYuan != nil {
		result.CoreBingYuan = diagnosis.CoreBingYuan.Description
		result.AnalysisPath = append(result.AnalysisPath, "【核心病源】"+diagnosis.CoreBingYuan.Description)
	}

	if yongShenResult.YongShen != nil {
		result.YongShen = yongShenResult.YongShen.WuXing.String() + "(" + yongShenResult.YongShen.ShiShen + ")"
		result.YongShenType = string(yongShenResult.YongShen.Type)
		result.YongShenDesc = yongShenResult.YongShen.Description
		result.AnalysisPath = append(result.AnalysisPath, "【用神】"+result.YongShen)
		result.AnalysisPath = append(result.AnalysisPath, "【用神类型】"+result.YongShenType)
		result.AnalysisPath = append(result.AnalysisPath, "【用神说明】"+result.YongShenDesc)
	}

	for _, xs := range yongShenResult.XiShen {
		xiStr := xs.WuXing.String() + "(" + xs.ShiShen + ")"
		result.XiShen = append(result.XiShen, xiStr)
		result.AnalysisPath = append(result.AnalysisPath, "【喜神】"+xiStr+" - "+xs.Reason)
	}

	for _, js := range yongShenResult.JiShen {
		jiStr := js.WuXing.String() + "(" + js.ShiShen + ")"
		result.JiShen = append(result.JiShen, jiStr)
		result.AnalysisPath = append(result.AnalysisPath, "【忌神】"+jiStr+" - "+js.Reason)
	}

	for _, sub := range diagnosis.SubBingYuan {
		result.AnalysisPath = append(result.AnalysisPath, "【次级病源】"+sub.Description)
	}

	result.AnalysisPath = append(result.AnalysisPath, yongShenResult.Derivation...)
}

func populateLegacyFields(result *Step4XiJiResult, step3Result *step3_wangshuai.Step3WangShuaiResult) {
	if result.YongShenResult == nil || result.YongShenResult.YongShen == nil {
		switch result.WangShuaiType {
		case step3_wangshuai.ShenRuo:
			analyzeShenRuo(result)
		case step3_wangshuai.ShenWang:
			analyzeShenWang(result)
		case step3_wangshuai.CongQiang:
			analyzeCongQiang(result)
		case step3_wangshuai.CongRuo:
			analyzeCongRuo(result)
		}
		return
	}

	yongShen := result.YongShenResult.YongShen
	result.FavorableWuXing = []base.WuXing{yongShen.WuXing}
	result.FavorableShiShen = []string{yongShen.ShiShen}

	for _, xs := range result.YongShenResult.XiShen {
		result.FavorableWuXing = append(result.FavorableWuXing, xs.WuXing)
		result.FavorableShiShen = append(result.FavorableShiShen, xs.ShiShen)
	}

	for _, js := range result.YongShenResult.JiShen {
		result.UnfavorableWuXing = append(result.UnfavorableWuXing, js.WuXing)
		result.UnfavorableShiShen = append(result.UnfavorableShiShen, js.ShiShen)
	}

	result.Reason = generateReasonFromBingYuan(result)
}

func generateReasonFromBingYuan(result *Step4XiJiResult) string {
	if result.BingYuanDiagnosis == nil || result.BingYuanDiagnosis.CoreBingYuan == nil {
		return "按旺衰法推断"
	}

	core := result.BingYuanDiagnosis.CoreBingYuan
	reason := "病药理论："

	switch core.Type {
	case BingYuanTiaoHou:
		reason += "命局寒暖失衡，" + core.Description + "，以" + result.YongShenResult.YongShen.WuXing.String() + "调候为用"
	case BingYuanTongGuan:
		reason += "五行相战，" + core.Description + "，以" + result.YongShenResult.YongShen.WuXing.String() + "通关为用"
	case BingYuanFuYi:
		reason += core.Description + "，以" + result.YongShenResult.YongShen.WuXing.String() + "扶抑为用"
	case BingYuanBingYao:
		reason += core.Description + "，以" + result.YongShenResult.YongShen.WuXing.String() + "为药，对症下药"
	case BingYuanGeJu:
		reason += core.Description + "，以" + result.YongShenResult.YongShen.WuXing.String() + "顺势为用"
	}

	return reason
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

func getKeXieHaoWuXing(dm base.WuXing) []base.WuXing {
	return []base.WuXing{
		getWuXingKeWo(dm),
		base.WuXingSheng[dm],
		base.WuXingKe[dm],
	}
}
