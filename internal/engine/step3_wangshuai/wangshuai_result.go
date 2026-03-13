package step3_wangshuai

import (
	"bazi-analyzer/internal/engine/step2_rebuild"
	"bazi-analyzer/internal/model/base"
)

type WangShuaiType string

const (
	ShenWang  WangShuaiType = "身旺"
	ShenRuo   WangShuaiType = "身弱"
	CongQiang WangShuaiType = "从强"
	CongRuo   WangShuaiType = "从弱"
)

type Step3WangShuaiResult struct {
	DayMaster       base.TianGan
	DayMasterWuXing base.WuXing
	IsDeLing        bool
	IsDeDi          bool
	IsDeShi         bool
	ShengFuCount    int
	KeXieHaoCount   int
	WangShuaiType   WangShuaiType
	Reason          string
}

func AnalyzeWangShuai(chart *base.BaZiChart, step2Result *step2_rebuild.Step2RebuildResult) *Step3WangShuaiResult {
	dayMaster := chart.GetDayMaster()
	dayMasterWuXing := step2_rebuild.GetRebuiltTianGanWuXing(step2Result, dayMaster)

	result := &Step3WangShuaiResult{
		DayMaster:       dayMaster,
		DayMasterWuXing: dayMasterWuXing,
	}

	analyzeDeLing(chart, step2Result, result)
	analyzeDeDi(chart, step2Result, result)
	analyzeDeShi(chart, step2Result, result)
	analyzeWangShuaiType(result)

	return result
}

func analyzeDeLing(chart *base.BaZiChart, step2Result *step2_rebuild.Step2RebuildResult, result *Step3WangShuaiResult) {
	monthLing := chart.GetMonthLing()
	monthLingWuXing := step2_rebuild.GetRebuiltDiZhiWuXing(step2Result, monthLing)

	if monthLingWuXing == result.DayMasterWuXing || base.Sheng(monthLingWuXing, result.DayMasterWuXing) {
		result.IsDeLing = true
	}
}

func analyzeDeDi(chart *base.BaZiChart, step2Result *step2_rebuild.Step2RebuildResult, result *Step3WangShuaiResult) {
	for _, dz := range chart.GetAllDiZhi() {
		dzWuXing := step2_rebuild.GetRebuiltDiZhiWuXing(step2Result, dz)
		if dzWuXing == result.DayMasterWuXing {
			result.IsDeDi = true
			return
		}
	}
}

func analyzeDeShi(chart *base.BaZiChart, step2Result *step2_rebuild.Step2RebuildResult, result *Step3WangShuaiResult) {
	shengFu := 0
	keXieHao := 0

	for _, tg := range chart.GetAllTianGan() {
		if step2Result.IsTianGanAbsorbed(tg) && tg.Name != result.DayMaster.Name {
			continue
		}
		tgWuXing := step2_rebuild.GetRebuiltTianGanWuXing(step2Result, tg)
		if tgWuXing == result.DayMasterWuXing || base.Sheng(tgWuXing, result.DayMasterWuXing) {
			shengFu++
		} else {
			keXieHao++
		}
	}

	for _, dz := range chart.GetAllDiZhi() {
		if step2Result.IsDiZhiAbsorbed(dz) {
			continue
		}
		dzWuXing := step2_rebuild.GetRebuiltDiZhiWuXing(step2Result, dz)
		if dzWuXing == result.DayMasterWuXing || base.Sheng(dzWuXing, result.DayMasterWuXing) {
			shengFu++
		} else {
			keXieHao++
		}
	}

	result.ShengFuCount = shengFu
	result.KeXieHaoCount = keXieHao

	if shengFu > 0 {
		result.IsDeShi = true
	}
}

func analyzeWangShuaiType(result *Step3WangShuaiResult) {
	if isCongGe(result) {
		externalShengFu := result.ShengFuCount
		if externalShengFu > 0 {
			externalShengFu--
		}
		if externalShengFu == 0 {
			result.WangShuaiType = CongRuo
			result.Reason = "日主无任何生扶，全局气势专一，构成从格"
		} else {
			result.WangShuaiType = CongQiang
			result.Reason = "日主气势极旺，全局顺其气势，构成从强格"
		}
		return
	}

	deLingDeDiDeShi := 0
	if result.IsDeLing {
		deLingDeDiDeShi++
	}
	if result.IsDeDi {
		deLingDeDiDeShi++
	}
	if result.IsDeShi {
		deLingDeDiDeShi++
	}

	if result.ShengFuCount > result.KeXieHaoCount || deLingDeDiDeShi >= 2 {
		result.WangShuaiType = ShenWang
		result.Reason = "生扶力量大于克泄耗，或得令得地得势满足两项以上"
	} else {
		result.WangShuaiType = ShenRuo
		result.Reason = "克泄耗力量大于生扶"
	}
}

func isCongGe(result *Step3WangShuaiResult) bool {
	externalShengFu := result.ShengFuCount
	if externalShengFu > 0 {
		externalShengFu--
	}

	if !result.IsDeLing && !result.IsDeDi && externalShengFu <= 1 {
		if result.KeXieHaoCount >= 3 && result.KeXieHaoCount >= (externalShengFu+1)*3 {
			return true
		}
	}
	return false
}
