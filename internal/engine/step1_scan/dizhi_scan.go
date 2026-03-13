package step1_scan

import (
	"bazi-analyzer/internal/model/base"
	"bazi-analyzer/internal/model/relation"
)

func scanDiZhiRelations(chart *base.BaZiChart, result *Step1ScanResult) {
	diZhiList := chart.GetAllDiZhi()

	scanSanHui(diZhiList, result)
	scanSanHe(diZhiList, result)
	scanLiuHe(diZhiList, result)
}

func scanSanHui(diZhiList []base.DiZhi, result *Step1ScanResult) {
	for _, sanhui := range relation.SanHuiList {
		count := 0
		for _, dz := range sanhui.Elements {
			if relation.ContainsDiZhi(diZhiList, dz) {
				count++
			}
		}
		if count == 3 {
			result.DiZhiSanHui = append(result.DiZhiSanHui, sanhui)
		}
	}
}

func scanSanHe(diZhiList []base.DiZhi, result *Step1ScanResult) {
	for _, sanhe := range relation.SanHeList {
		count := 0
		for _, dz := range sanhe.Elements {
			if relation.ContainsDiZhi(diZhiList, dz) {
				count++
			}
		}
		if count == 3 {
			result.DiZhiSanHe = append(result.DiZhiSanHe, sanhe)
		}
	}
}

func scanLiuHe(diZhiList []base.DiZhi, result *Step1ScanResult) {
	for i := 0; i < len(diZhiList); i++ {
		for j := i + 1; j < len(diZhiList); j++ {
			if liuhe, ok := relation.GetDiZhiLiuHe(diZhiList[i], diZhiList[j]); ok {
				result.DiZhiLiuHe = append(result.DiZhiLiuHe, liuhe)
			}
		}
	}
}
