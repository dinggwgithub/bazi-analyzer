package step2_rebuild

import (
	"bazi-analyzer/internal/engine/step1_scan"
	"bazi-analyzer/internal/model/base"
)

type RebuiltElement struct {
	OriginalName   string
	OriginalWuXing base.WuXing
	RebuiltWuXing  base.WuXing
	IsRebuilt      bool
	Reason         string
}

func (r *Step2RebuildResult) IsDiZhiAbsorbed(dz base.DiZhi) bool {
	return r.AbsorbedDiZhi[dz.Name]
}

func (r *Step2RebuildResult) IsTianGanAbsorbed(tg base.TianGan) bool {
	return r.AbsorbedTianGan[tg.Name]
}

func (r *Step2RebuildResult) MarkDiZhiAbsorbed(dz base.DiZhi) {
	r.AbsorbedDiZhi[dz.Name] = true
}

func (r *Step2RebuildResult) MarkTianGanAbsorbed(tg base.TianGan) {
	r.AbsorbedTianGan[tg.Name] = true
}

type Step2RebuildResult struct {
	OriginalChart   base.BaZiChart
	Step1Result     *step1_scan.Step1ScanResult
	RebuiltTianGan  []RebuiltElement
	RebuiltDiZhi    []RebuiltElement
	AbsorbedDiZhi   map[string]bool
	AbsorbedTianGan map[string]bool
	HasChanges      bool
}

func NewStep2RebuildResult(chart base.BaZiChart, step1Result *step1_scan.Step1ScanResult) *Step2RebuildResult {
	result := &Step2RebuildResult{
		OriginalChart:   chart,
		Step1Result:     step1Result,
		AbsorbedDiZhi:   make(map[string]bool),
		AbsorbedTianGan: make(map[string]bool),
		HasChanges:      false,
	}

	for _, tg := range chart.GetAllTianGan() {
		result.RebuiltTianGan = append(result.RebuiltTianGan, RebuiltElement{
			OriginalName:   tg.Name,
			OriginalWuXing: tg.Element,
			RebuiltWuXing:  tg.Element,
			IsRebuilt:      false,
		})
	}

	for _, dz := range chart.GetAllDiZhi() {
		result.RebuiltDiZhi = append(result.RebuiltDiZhi, RebuiltElement{
			OriginalName:   dz.Name,
			OriginalWuXing: dz.Element,
			RebuiltWuXing:  dz.Element,
			IsRebuilt:      false,
		})
	}

	return result
}

func RebuildBaZi(chart *base.BaZiChart, step1Result *step1_scan.Step1ScanResult) *Step2RebuildResult {
	result := NewStep2RebuildResult(*chart, step1Result)

	processSanHui(chart, result)
	processSanHe(chart, result)
	processTianGanWuHe(chart, result)
	processLiuHe(chart, result)

	return result
}
