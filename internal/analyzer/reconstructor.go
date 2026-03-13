package analyzer

import (
	"bazi-analyzer/internal/model"
	"bazi-analyzer/pkg/dizhi"
	"bazi-analyzer/pkg/tiangan"
	"bazi-analyzer/pkg/wuxing"
)

type PatternReconstructor struct{}

func NewPatternReconstructor() *PatternReconstructor {
	return &PatternReconstructor{}
}

type ReconstructResult struct {
	OriginalBazi      *model.Bazi
	ReconstructedBazi *model.Bazi
	Transformations   []Transformation
	MajorRelation     *model.DizhiRelation
}

type Transformation struct {
	Type        string
	Original    string
	Transformed string
	Reason      string
}

func (r *PatternReconstructor) Reconstruct(bazi *model.Bazi, scanResult *ScanResult) *ReconstructResult {
	result := &ReconstructResult{
		OriginalBazi: bazi,
	}
	
	reconstructed := *bazi
	
	majorRelation := r.findMajorRelation(scanResult)
	if majorRelation != nil {
		result.MajorRelation = majorRelation
		
		if majorRelation.Type == model.RelationSanhui || majorRelation.Type == model.RelationSanhe {
			transformations := r.applyDizhiHehua(&reconstructed, majorRelation)
			result.Transformations = transformations
		}
	}
	
	for i, wuhe := range scanResult.TianganRelations {
		if wuhe.Type == model.RelationWuhe {
			scanResult.TianganRelations[i].Success = r.checkTianganWuhua(bazi, &wuhe)
		}
	}
	
	result.ReconstructedBazi = &reconstructed
	
	return result
}

func (r *PatternReconstructor) findMajorRelation(scanResult *ScanResult) *model.DizhiRelation {
	for _, rel := range scanResult.DizhiRelations {
		if rel.Type == model.RelationSanhui {
			return &rel
		}
	}
	
	for _, rel := range scanResult.DizhiRelations {
		if rel.Type == model.RelationSanhe {
			return &rel
		}
	}
	
	return nil
}

func (r *PatternReconstructor) applyDizhiHehua(bazi *model.Bazi, relation *model.DizhiRelation) []Transformation {
	var transformations []Transformation
	
	for _, d := range relation.Dizhis {
		originalWuxing := d.Wuxing()
		if originalWuxing != relation.ResultWuxing {
			transformations = append(transformations, Transformation{
				Type:        "地支合化",
				Original:    d.String() + "(" + originalWuxing.String() + ")",
				Transformed: d.String() + "(" + relation.ResultWuxing.String() + ")",
				Reason:      relation.Type.String() + "成立",
			})
		}
	}
	
	return transformations
}

func (r *PatternReconstructor) checkTianganWuhua(bazi *model.Bazi, wuhe *model.TianganRelation) bool {
	monthWuxing := bazi.GetMonthWuxing()
	
	if monthWuxing == wuhe.ResultWuxing {
		return true
	}
	
	tiangans := bazi.GetTiangans()
	for _, t := range tiangans {
		if t.Wuxing() == wuhe.ResultWuxing {
			return true
		}
	}
	
	return false
}

func (r *PatternReconstructor) GetTransformedWuxing(d dizhi.Dizhi, relation *model.DizhiRelation) wuxing.Wuxing {
	if relation == nil {
		return d.Wuxing()
	}
	
	for _, rd := range relation.Dizhis {
		if rd == d {
			return relation.ResultWuxing
		}
	}
	
	return d.Wuxing()
}

func (r *PatternReconstructor) IsDizhiInHehua(d dizhi.Dizhi, relation *model.DizhiRelation) bool {
	if relation == nil {
		return false
	}
	
	for _, rd := range relation.Dizhis {
		if rd == d {
			return true
		}
	}
	return false
}

func (r *PatternReconstructor) AnalyzeGanzhiRelation(pillar model.Pillar) string {
	tg := pillar.Tiangan
	dz := dizhi.Dizhi(pillar.Dizhi.Original)
	tgWuxing := tg.Wuxing()
	dzWuxing := dz.Wuxing()
	
	if tgWuxing == dzWuxing {
		return "通根"
	}
	
	if tgWuxing.Sheng() == dzWuxing {
		return "泄气"
	}
	
	if dzWuxing.Sheng() == tgWuxing {
		return "得益"
	}
	
	if tgWuxing.Ke() == dzWuxing {
		return "盖头"
	}
	
	if dzWuxing.Ke() == tgWuxing {
		return "截脚"
	}
	
	return "普通"
}

func (r *PatternReconstructor) GetTonggenStrength(tg tiangan.Tiangan, dz dizhi.Dizhi) int {
	tgWuxing := tg.Wuxing()
	
	canggan := dz.GetCanggan()
	for i, c := range canggan {
		cgWuxing := getWuxingFromTianganName(c.TianganName)
		if cgWuxing == tgWuxing {
			switch i {
			case 0:
				return 3
			case 1:
				return 2
			case 2:
				return 1
			}
		}
	}
	return 0
}

func getWuxingFromTianganName(name string) wuxing.Wuxing {
	m := map[string]wuxing.Wuxing{
		"甲": wuxing.Mu, "乙": wuxing.Mu,
		"丙": wuxing.Huo, "丁": wuxing.Huo,
		"戊": wuxing.Tu, "己": wuxing.Tu,
		"庚": wuxing.Jin, "辛": wuxing.Jin,
		"壬": wuxing.Shui, "癸": wuxing.Shui,
	}
	return m[name]
}
