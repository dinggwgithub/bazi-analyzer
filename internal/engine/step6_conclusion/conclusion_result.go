package step6_conclusion

import (
	"bazi-analyzer/internal/engine/step1_scan"
	"bazi-analyzer/internal/engine/step2_rebuild"
	"bazi-analyzer/internal/engine/step3_wangshuai"
	"bazi-analyzer/internal/engine/step4_xiji"
	"bazi-analyzer/internal/engine/step5_analyze"
	"bazi-analyzer/internal/model/base"
)

type GeJuLevel string

const (
	GeJuUpper   GeJuLevel = "上格"
	GeJuMiddle  GeJuLevel = "中格"
	GeJuLower   GeJuLevel = "下格"
	GeJuSpecial GeJuLevel = "特殊格局"
)

type Step6ConclusionResult struct {
	GeJuType        string
	GeJuLevel       GeJuLevel
	GeJuDescription string
	SuiYunAnalysis  []string
	OverallSummary  string
	KeyAdvice       []string
}

func AnalyzeConclusion(chart *base.BaZiChart, step1Result *step1_scan.Step1ScanResult, step2Result *step2_rebuild.Step2RebuildResult, step3Result *step3_wangshuai.Step3WangShuaiResult, step4Result *step4_xiji.Step4XiJiResult, step5Result *step5_analyze.Step5AnalyzeResult) *Step6ConclusionResult {
	result := &Step6ConclusionResult{}

	analyzeGeJuType(chart, step1Result, step2Result, step3Result, step4Result, step5Result, result)
	analyzeGeJuLevel(chart, step1Result, step2Result, step3Result, step4Result, step5Result, result)
	analyzeSuiYun(chart, step1Result, step2Result, step3Result, step4Result, step5Result, result)
	generateOverallSummary(chart, step1Result, step2Result, step3Result, step4Result, step5Result, result)
	generateKeyAdvice(chart, step1Result, step2Result, step3Result, step4Result, step5Result, result)

	return result
}

func analyzeGeJuType(chart *base.BaZiChart, step1Result *step1_scan.Step1ScanResult, step2Result *step2_rebuild.Step2RebuildResult, step3Result *step3_wangshuai.Step3WangShuaiResult, step4Result *step4_xiji.Step4XiJiResult, step5Result *step5_analyze.Step5AnalyzeResult, result *Step6ConclusionResult) {
	dayMaster := chart.GetDayMaster()

	if step3Result.WangShuaiType == step3_wangshuai.CongRuo || step3Result.WangShuaiType == step3_wangshuai.CongQiang {
		result.GeJuType = "从格"
		if len(step2Result.RebuiltDiZhi) > 0 {
			heJuType := getHeJuType(step2Result)
			result.GeJuType = "从" + heJuType + "格"
		} else {
			dominantWuXing := getDominantWuXing(step2Result, step4Result)
			if dominantWuXing != "" {
				shiShen := base.GetShiShen(dayMaster.Element, dominantWuXing, dayMaster.IsYang)
				result.GeJuType = "从" + string(shiShen) + "格"
			}
		}
		return
	}

	if len(step2Result.RebuiltDiZhi) > 0 {
		heJuType := getHeJuType(step2Result)
		result.GeJuType = heJuType + "局成格"
		return
	}

	if step3Result.WangShuaiType == step3_wangshuai.ShenWang {
		result.GeJuType = "身旺格"
	} else {
		result.GeJuType = "身弱格"
	}
}

func getHeJuType(step2Result *step2_rebuild.Step2RebuildResult) string {
	for _, dz := range step2Result.RebuiltDiZhi {
		if dz.RebuiltWuXing == base.Wood {
			return "木"
		} else if dz.RebuiltWuXing == base.Fire {
			return "火"
		} else if dz.RebuiltWuXing == base.Earth {
			return "土"
		} else if dz.RebuiltWuXing == base.Metal {
			return "金"
		} else if dz.RebuiltWuXing == base.Water {
			return "水"
		}
	}
	return ""
}

func getDominantWuXing(step2Result *step2_rebuild.Step2RebuildResult, step4Result *step4_xiji.Step4XiJiResult) base.WuXing {
	for _, fx := range step4Result.FavorableWuXing {
		return fx
	}
	return ""
}

func analyzeGeJuLevel(chart *base.BaZiChart, step1Result *step1_scan.Step1ScanResult, step2Result *step2_rebuild.Step2RebuildResult, step3Result *step3_wangshuai.Step3WangShuaiResult, step4Result *step4_xiji.Step4XiJiResult, step5Result *step5_analyze.Step5AnalyzeResult, result *Step6ConclusionResult) {
	score := 0

	if len(step2Result.RebuiltDiZhi) > 0 {
		score += 3
	}
	if step3Result.WangShuaiType == step3_wangshuai.CongRuo || step3Result.WangShuaiType == step3_wangshuai.CongQiang {
		score += 2
	}
	if len(step4Result.FavorableWuXing) > 0 {
		score += 1
	}

	if result.GeJuType == "从杀格" {
		result.GeJuLevel = GeJuSpecial
		result.GeJuDescription = "全局气势专一，从杀成格，属于特殊格局，层次较高"
	} else if score >= 4 {
		result.GeJuLevel = GeJuUpper
		result.GeJuDescription = "格局清纯，五行流通有情，层次较高"
	} else if score >= 2 {
		result.GeJuLevel = GeJuMiddle
		result.GeJuDescription = "格局基本平衡，有得有失，层次中等"
	} else {
		result.GeJuLevel = GeJuLower
		result.GeJuDescription = "五行驳杂，矛盾较多，层次较低"
	}
}

func analyzeSuiYun(chart *base.BaZiChart, step1Result *step1_scan.Step1ScanResult, step2Result *step2_rebuild.Step2RebuildResult, step3Result *step3_wangshuai.Step3WangShuaiResult, step4Result *step4_xiji.Step4XiJiResult, step5Result *step5_analyze.Step5AnalyzeResult, result *Step6ConclusionResult) {
	dayMaster := chart.GetDayMaster()

	if len(step2Result.RebuiltDiZhi) > 0 {
		heJuWuXing := getHeJuWuXing(step2Result)
		result.SuiYunAnalysis = append(result.SuiYunAnalysis, "原局地支合局成立，岁运遇"+heJuWuXing.String()+"旺地则引动合局力量")
	}

	if len(step1Result.TianGanSiChong) > 0 {
		for _, chong := range step1Result.TianGanSiChong {
			result.SuiYunAnalysis = append(result.SuiYunAnalysis, "原局有"+chong.Elements[0].Name+chong.Elements[1].Name+"四冲，岁运引动则主变动")
		}
	}

	if len(step4Result.FavorableWuXing) > 0 {
		favStr := ""
		for _, wx := range step4Result.FavorableWuXing {
			favStr += wx.String() + "、"
		}
		if len(favStr) > 0 {
			favStr = favStr[:len(favStr)-1]
		}
		result.SuiYunAnalysis = append(result.SuiYunAnalysis, "岁运走"+favStr+"运为吉，主得助力")
	}

	if len(step4Result.UnfavorableWuXing) > 0 {
		unfavStr := ""
		for _, wx := range step4Result.UnfavorableWuXing {
			unfavStr += wx.String() + "、"
		}
		if len(unfavStr) > 0 {
			unfavStr = unfavStr[:len(unfavStr)-1]
		}
		result.SuiYunAnalysis = append(result.SuiYunAnalysis, "岁运走"+unfavStr+"运为凶，主压力与不顺")
	}

	if step3Result.WangShuaiType == step3_wangshuai.CongRuo {
		result.SuiYunAnalysis = append(result.SuiYunAnalysis, dayMaster.Element.String()+"日主极弱从势，最忌岁运逢"+dayMaster.Element.String()+"通根得地")
	}
}

func getHeJuWuXing(step2Result *step2_rebuild.Step2RebuildResult) base.WuXing {
	for _, dz := range step2Result.RebuiltDiZhi {
		return dz.RebuiltWuXing
	}
	return ""
}

func generateOverallSummary(chart *base.BaZiChart, step1Result *step1_scan.Step1ScanResult, step2Result *step2_rebuild.Step2RebuildResult, step3Result *step3_wangshuai.Step3WangShuaiResult, step4Result *step4_xiji.Step4XiJiResult, step5Result *step5_analyze.Step5AnalyzeResult, result *Step6ConclusionResult) {
	dayMaster := chart.GetDayMaster()
	summary := dayMaster.Element.String() + "(" + dayMaster.Name + ")日主"

	if step3Result.WangShuaiType == step3_wangshuai.CongRuo {
		summary += "极弱从势，"
	} else if step3Result.WangShuaiType == step3_wangshuai.ShenWang {
		summary += "身旺，"
	} else {
		summary += "身弱，"
	}

	if len(step2Result.RebuiltDiZhi) > 0 {
		heJuType := getHeJuType(step2Result)
		summary += "地支" + heJuType + "局合化成功，气势专一。"
	} else {
		summary += "需平衡五行流通。"
	}

	favStr := ""
	for _, wx := range step4Result.FavorableWuXing {
		favStr += wx.String() + "、"
	}
	if len(favStr) > 0 {
		favStr = favStr[:len(favStr)-1]
		summary += "喜用" + favStr
	}

	result.OverallSummary = summary
}

func generateKeyAdvice(chart *base.BaZiChart, step1Result *step1_scan.Step1ScanResult, step2Result *step2_rebuild.Step2RebuildResult, step3Result *step3_wangshuai.Step3WangShuaiResult, step4Result *step4_xiji.Step4XiJiResult, step5Result *step5_analyze.Step5AnalyzeResult, result *Step6ConclusionResult) {
	dayMaster := chart.GetDayMaster()

	if result.GeJuType == "从杀格" {
		result.KeyAdvice = append(result.KeyAdvice, "从杀成格，宜顺势而为，忌逆势反抗")
		result.KeyAdvice = append(result.KeyAdvice, "宜从事与"+getHeJuType(step2Result)+"相关行业")
	}

	if len(step1Result.TianGanSiChong) > 0 {
		result.KeyAdvice = append(result.KeyAdvice, "命带四冲，宜注意情绪管理，避免冲动决策")
	}

	if step3Result.WangShuaiType == step3_wangshuai.ShenRuo {
		result.KeyAdvice = append(result.KeyAdvice, dayMaster.Element.String()+"日主身弱，宜增强自身能量，注意身体健康")
	}

	favStr := ""
	for _, wx := range step4Result.FavorableWuXing {
		favStr += wx.String() + "、"
	}
	if len(favStr) > 0 {
		favStr = favStr[:len(favStr)-1]
		result.KeyAdvice = append(result.KeyAdvice, "方位、颜色、职业宜取"+favStr+"为用")
	}
}
