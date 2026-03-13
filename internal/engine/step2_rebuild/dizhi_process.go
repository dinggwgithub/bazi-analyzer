package step2_rebuild

import (
	"bazi-analyzer/internal/model/base"
)

func processSanHui(chart *base.BaZiChart, result *Step2RebuildResult) {
	for _, sanhui := range result.Step1Result.DiZhiSanHui {
		allAvailable := true
		for _, dz := range sanhui.Elements {
			if result.IsDiZhiAbsorbed(dz) {
				allAvailable = false
				break
			}
		}

		if !allAvailable {
			continue
		}

		if !CheckHeHuaSuccess(chart, sanhui.Result) {
			continue
		}

		for _, dz := range sanhui.Elements {
			result.MarkDiZhiAbsorbed(dz)
			for i := range result.RebuiltDiZhi {
				if result.RebuiltDiZhi[i].OriginalName == dz.Name {
					result.RebuiltDiZhi[i].IsRebuilt = true
					result.RebuiltDiZhi[i].RebuiltWuXing = sanhui.Result
					result.RebuiltDiZhi[i].Reason = "三会局合化" + string(sanhui.Result)
				}
			}
		}
		result.HasChanges = true
	}
}

func processSanHe(chart *base.BaZiChart, result *Step2RebuildResult) {
	for _, sanhe := range result.Step1Result.DiZhiSanHe {
		allAvailable := true
		for _, dz := range sanhe.Elements {
			if result.IsDiZhiAbsorbed(dz) {
				allAvailable = false
				break
			}
		}

		if !allAvailable {
			continue
		}

		if !CheckHeHuaSuccess(chart, sanhe.Result) {
			continue
		}

		for _, dz := range sanhe.Elements {
			result.MarkDiZhiAbsorbed(dz)
			for i := range result.RebuiltDiZhi {
				if result.RebuiltDiZhi[i].OriginalName == dz.Name {
					result.RebuiltDiZhi[i].IsRebuilt = true
					result.RebuiltDiZhi[i].RebuiltWuXing = sanhe.Result
					result.RebuiltDiZhi[i].Reason = "三合局合化" + string(sanhe.Result)
				}
			}
		}
		result.HasChanges = true
	}
}

func processLiuHe(chart *base.BaZiChart, result *Step2RebuildResult) {
	for _, liuhe := range result.Step1Result.DiZhiLiuHe {
		allAvailable := true
		for _, dz := range liuhe.Elements {
			if result.IsDiZhiAbsorbed(dz) {
				allAvailable = false
				break
			}
		}

		if !allAvailable {
			continue
		}

		if !CheckHeHuaSuccess(chart, liuhe.Result) {
			continue
		}

		for _, dz := range liuhe.Elements {
			result.MarkDiZhiAbsorbed(dz)
			for i := range result.RebuiltDiZhi {
				if result.RebuiltDiZhi[i].OriginalName == dz.Name {
					result.RebuiltDiZhi[i].IsRebuilt = true
					result.RebuiltDiZhi[i].RebuiltWuXing = liuhe.Result
					result.RebuiltDiZhi[i].Reason = "六合合化" + string(liuhe.Result)
				}
			}
		}
		result.HasChanges = true
	}
}
