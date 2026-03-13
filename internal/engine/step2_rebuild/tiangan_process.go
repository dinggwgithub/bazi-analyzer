package step2_rebuild

import (
	"bazi-analyzer/internal/model/base"
)

func processTianGanWuHe(chart *base.BaZiChart, result *Step2RebuildResult) {
	for _, wuhe := range result.Step1Result.TianGanWuHe {
		allAvailable := true
		for _, tg := range wuhe.Elements {
			if result.IsTianGanAbsorbed(tg) {
				allAvailable = false
				break
			}
		}

		if !allAvailable {
			continue
		}

		if !CheckHeHuaSuccess(chart, wuhe.Result) {
			continue
		}

		for _, tg := range wuhe.Elements {
			result.MarkTianGanAbsorbed(tg)
			for i := range result.RebuiltTianGan {
				if result.RebuiltTianGan[i].OriginalName == tg.Name {
					result.RebuiltTianGan[i].IsRebuilt = true
					result.RebuiltTianGan[i].RebuiltWuXing = wuhe.Result
					result.RebuiltTianGan[i].Reason = "天干五合合化" + string(wuhe.Result)
				}
			}
		}
		result.HasChanges = true
	}
}
