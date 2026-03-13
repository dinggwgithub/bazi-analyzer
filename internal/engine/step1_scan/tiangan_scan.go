package step1_scan

import (
	"bazi-analyzer/internal/model/base"
	"bazi-analyzer/internal/model/relation"
)

func scanTianGanRelations(chart *base.BaZiChart, result *Step1ScanResult) {
	tianGanList := chart.GetAllTianGan()

	for i := 0; i < len(tianGanList); i++ {
		for j := i + 1; j < len(tianGanList); j++ {
			tg1 := tianGanList[i]
			tg2 := tianGanList[j]

			if wuhe, ok := relation.GetTianGanWuHe(tg1, tg2); ok {
				result.TianGanWuHe = append(result.TianGanWuHe, wuhe)
			}

			if sichong, ok := relation.GetTianGanSiChong(tg1, tg2); ok {
				result.TianGanSiChong = append(result.TianGanSiChong, sichong)
			}
		}
	}
}
