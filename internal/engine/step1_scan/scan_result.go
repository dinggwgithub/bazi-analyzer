package step1_scan

import (
	"bazi-analyzer/internal/model/base"
	"bazi-analyzer/internal/model/relation"
)

type Step1ScanResult struct {
	TianGanWuHe   []relation.TianGanRelation
	TianGanSiChong []relation.TianGanRelation
	DiZhiSanHui   []relation.DiZhiRelation
	DiZhiSanHe    []relation.DiZhiRelation
	DiZhiLiuHe    []relation.DiZhiRelation
}

func NewStep1ScanResult() *Step1ScanResult {
	return &Step1ScanResult{
		TianGanWuHe:   make([]relation.TianGanRelation, 0),
		TianGanSiChong: make([]relation.TianGanRelation, 0),
		DiZhiSanHui:   make([]relation.DiZhiRelation, 0),
		DiZhiSanHe:    make([]relation.DiZhiRelation, 0),
		DiZhiLiuHe:    make([]relation.DiZhiRelation, 0),
	}
}

func ScanBaZi(chart *base.BaZiChart) *Step1ScanResult {
	result := NewStep1ScanResult()

	scanTianGanRelations(chart, result)
	scanDiZhiRelations(chart, result)

	return result
}
