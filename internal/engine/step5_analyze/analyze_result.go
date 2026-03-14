package step5_analyze

import (
	"bazi-analyzer/internal/engine/step1_scan"
	"bazi-analyzer/internal/engine/step2_rebuild"
	"bazi-analyzer/internal/engine/step4_xiji"
	"bazi-analyzer/internal/model/base"
)

type TianGanInteraction struct {
	From        base.TianGan
	To          base.TianGan
	Type        string
	Strength    string
	Description string
}

type DiZhiInteraction struct {
	From        base.DiZhi
	To          base.DiZhi
	Type        string
	Description string
}

type Step5AnalyzeResult struct {
	TianGanInteractions []TianGanInteraction
	DiZhiInteractions   []DiZhiInteraction
	KeyPoints           []string
}

func AnalyzeInteractions(chart *base.BaZiChart, step1Result *step1_scan.Step1ScanResult, step2Result *step2_rebuild.Step2RebuildResult, step4Result *step4_xiji.Step4XiJiResult) *Step5AnalyzeResult {
	result := &Step5AnalyzeResult{}

	analyzeTianGan(chart, step1Result, step2Result, step4Result, result)
	analyzeDiZhi(chart, step2Result, result)
	generateKeyPoints(chart, step1Result, step2Result, step4Result, result)

	return result
}

func analyzeTianGan(chart *base.BaZiChart, step1Result *step1_scan.Step1ScanResult, step2Result *step2_rebuild.Step2RebuildResult, step4Result *step4_xiji.Step4XiJiResult, result *Step5AnalyzeResult) {
	dayMaster := chart.GetDayMaster()

	for _, chong := range step1Result.TianGanSiChong {
		isFavorable := isXiJiFavorable(step4Result, chong.Elements[0].Element, chong.Elements[1].Element)
		strength := "强"
		if isFavorable {
			strength = "吉"
		} else {
			strength = "凶"
		}
		result.TianGanInteractions = append(result.TianGanInteractions, TianGanInteraction{
			From:        chong.Elements[0],
			To:          chong.Elements[1],
			Type:        "四冲",
			Strength:    strength,
			Description: chong.Elements[0].Name + chong.Elements[1].Name + "四冲",
		})
	}

	tianGanList := chart.GetAllTianGan()
	for i := 0; i < len(tianGanList); i++ {
		tg1 := tianGanList[i]
		if step2Result.IsTianGanAbsorbed(tg1) && tg1.Name != dayMaster.Name {
			continue
		}
		for j := i + 1; j < len(tianGanList); j++ {
			tg2 := tianGanList[j]
			if step2Result.IsTianGanAbsorbed(tg2) && tg2.Name != dayMaster.Name {
				continue
			}
			if hasWuHeOrSiChong(step1Result, tg1, tg2) {
				continue
			}
			wx1 := step2_rebuild.GetRebuiltTianGanWuXing(step2Result, tg1)
			wx2 := step2_rebuild.GetRebuiltTianGanWuXing(step2Result, tg2)
			if base.Sheng(wx1, wx2) {
				result.TianGanInteractions = append(result.TianGanInteractions, TianGanInteraction{
					From:        tg1,
					To:          tg2,
					Type:        "生",
					Strength:    "中",
					Description: tg1.Name + wx1.String() + "生" + tg2.Name + wx2.String(),
				})
			} else if base.Ke(wx1, wx2) {
				isFavorable := isXiJiFavorable(step4Result, wx1, wx2)
				strength := "中"
				if isFavorable {
					strength = "吉"
				} else {
					strength = "凶"
				}
				result.TianGanInteractions = append(result.TianGanInteractions, TianGanInteraction{
					From:        tg1,
					To:          tg2,
					Type:        "克",
					Strength:    strength,
					Description: tg1.Name + wx1.String() + "克" + tg2.Name + wx2.String(),
				})
			}
		}
	}
}

func analyzeDiZhi(chart *base.BaZiChart, step2Result *step2_rebuild.Step2RebuildResult, result *Step5AnalyzeResult) {
	if len(step2Result.RebuiltDiZhi) > 0 {
		result.DiZhiInteractions = append(result.DiZhiInteractions, DiZhiInteraction{
			Type:        "合局",
			Description: "地支合局成立，内部力量融合为一体",
		})
	}
}

func generateKeyPoints(chart *base.BaZiChart, step1Result *step1_scan.Step1ScanResult, step2Result *step2_rebuild.Step2RebuildResult, step4Result *step4_xiji.Step4XiJiResult, result *Step5AnalyzeResult) {
	if len(step2Result.RebuiltDiZhi) > 0 {
		result.KeyPoints = append(result.KeyPoints, "地支合局为全局核心，主导命局气势")
	}
	if len(step1Result.TianGanSiChong) > 0 {
		result.KeyPoints = append(result.KeyPoints, "天干四冲带来变动与压力")
	}
}

func hasWuHeOrSiChong(step1Result *step1_scan.Step1ScanResult, tg1, tg2 base.TianGan) bool {
	for _, wuhe := range step1Result.TianGanWuHe {
		if (wuhe.Elements[0].Name == tg1.Name && wuhe.Elements[1].Name == tg2.Name) ||
			(wuhe.Elements[0].Name == tg2.Name && wuhe.Elements[1].Name == tg1.Name) {
			return true
		}
	}
	for _, sichong := range step1Result.TianGanSiChong {
		if (sichong.Elements[0].Name == tg1.Name && sichong.Elements[1].Name == tg2.Name) ||
			(sichong.Elements[0].Name == tg2.Name && sichong.Elements[1].Name == tg1.Name) {
			return true
		}
	}
	return false
}

func isXiJiFavorable(step4Result *step4_xiji.Step4XiJiResult, wx1, wx2 base.WuXing) bool {
	for _, fx := range step4Result.YongShenWuXing {
		if wx1 == fx || wx2 == fx {
			return true
		}
	}
	return false
}
