package analyzer

import (
	"bazi-analyzer/internal/model"
	"bazi-analyzer/pkg/dizhi"
	"bazi-analyzer/pkg/tiangan"
	"bazi-analyzer/pkg/wuxing"
)

type RelationScanner struct{}

func NewRelationScanner() *RelationScanner {
	return &RelationScanner{}
}

type ScanResult struct {
	TianganRelations []model.TianganRelation
	DizhiRelations   []model.DizhiRelation
}

func (s *RelationScanner) Scan(bazi *model.Bazi) *ScanResult {
	result := &ScanResult{}
	
	result.TianganRelations = s.scanTianganRelations(bazi)
	result.DizhiRelations = s.scanDizhiRelations(bazi)
	
	return result
}

func (s *RelationScanner) scanTianganRelations(bazi *model.Bazi) []model.TianganRelation {
	var relations []model.TianganRelation
	tiangans := bazi.GetTiangans()
	positions := []model.Position{
		model.PositionNian, model.PositionYue, model.PositionRi, model.PositionShi,
	}
	
	for i := 0; i < len(tiangans); i++ {
		for j := i + 1; j < len(tiangans); j++ {
			t1, t2 := tiangans[i], tiangans[j]
			pos1, pos2 := positions[i], positions[j]
			
			if wuhe := tiangan.GetWuhe(t1, t2); wuhe != nil {
				relations = append(relations, model.TianganRelation{
					Type:         model.RelationWuhe,
					Tiangan1:     t1,
					Tiangan2:     t2,
					ResultWuxing: wuhe.HuaWuxing,
					Success:      false,
					Position1:    pos1,
					Position2:    pos2,
				})
			}
			
			if tiangan.IsSichong(t1, t2) {
				relations = append(relations, model.TianganRelation{
					Type:      model.RelationSichong,
					Tiangan1:  t1,
					Tiangan2:  t2,
					Position1: pos1,
					Position2: pos2,
				})
			}
		}
	}
	
	return relations
}

func (s *RelationScanner) scanDizhiRelations(bazi *model.Bazi) []model.DizhiRelation {
	var relations []model.DizhiRelation
	dizhis := bazi.GetDizhis()
	positions := []model.Position{
		model.PositionNian, model.PositionYue, model.PositionRi, model.PositionShi,
	}
	
	if sanhui := dizhi.FindSanhui(dizhis); sanhui != nil {
		relations = append(relations, model.DizhiRelation{
			Type:         model.RelationSanhui,
			Dizhis:       sanhui.Dizhis,
			ResultWuxing: sanhui.HuaWuxing,
			Success:      true,
			Positions:    s.findPositions(dizhis, sanhui.Dizhis, positions),
		})
		return relations
	}
	
	if sanhe := dizhi.FindSanhe(dizhis); sanhe != nil {
		relations = append(relations, model.DizhiRelation{
			Type:         model.RelationSanhe,
			Dizhis:       []dizhi.Dizhi{sanhe.Changsheng, sanhe.Diwang, sanhe.Mu},
			ResultWuxing: sanhe.HuaWuxing,
			Success:      true,
			Positions:    s.findPositions(dizhis, []dizhi.Dizhi{sanhe.Changsheng, sanhe.Diwang, sanhe.Mu}, positions),
		})
		return relations
	}
	
	liuheList := dizhi.FindLiuhe(dizhis)
	for _, liuhe := range liuheList {
		relations = append(relations, model.DizhiRelation{
			Type:         model.RelationLiuhe,
			Dizhis:       []dizhi.Dizhi{liuhe.Dizhi1, liuhe.Dizhi2},
			ResultWuxing: liuhe.HuaWuxing,
			Success:      false,
			Positions:    s.findPositions(dizhis, []dizhi.Dizhi{liuhe.Dizhi1, liuhe.Dizhi2}, positions),
		})
	}
	
	chongList := dizhi.FindChong(dizhis)
	for _, chong := range chongList {
		relations = append(relations, model.DizhiRelation{
			Type:      model.RelationChong,
			Dizhis:    []dizhi.Dizhi{chong.Dizhi1, chong.Dizhi2},
			Positions: s.findPositions(dizhis, []dizhi.Dizhi{chong.Dizhi1, chong.Dizhi2}, positions),
		})
	}
	
	xingList := dizhi.FindXing(dizhis)
	for _, xing := range xingList {
		relations = append(relations, model.DizhiRelation{
			Type:      model.RelationXing,
			Dizhis:    xing.Dizhis,
			Positions: s.findPositions(dizhis, xing.Dizhis, positions),
		})
	}
	
	haiList := dizhi.FindHai(dizhis)
	for _, hai := range haiList {
		relations = append(relations, model.DizhiRelation{
			Type:      model.RelationHai,
			Dizhis:    []dizhi.Dizhi{hai.Dizhi1, hai.Dizhi2},
			Positions: s.findPositions(dizhis, []dizhi.Dizhi{hai.Dizhi1, hai.Dizhi2}, positions),
		})
	}
	
	poList := dizhi.FindPo(dizhis)
	for _, po := range poList {
		relations = append(relations, model.DizhiRelation{
			Type:      model.RelationPo,
			Dizhis:    []dizhi.Dizhi{po.Dizhi1, po.Dizhi2},
			Positions: s.findPositions(dizhis, []dizhi.Dizhi{po.Dizhi1, po.Dizhi2}, positions),
		})
	}
	
	return relations
}

func (s *RelationScanner) findPositions(allDizhis []dizhi.Dizhi, targetDizhis []dizhi.Dizhi, positions []model.Position) []model.Position {
	var result []model.Position
	for _, td := range targetDizhis {
		for i, ad := range allDizhis {
			if ad == td {
				result = append(result, positions[i])
				break
			}
		}
	}
	return result
}

func (s *RelationScanner) CheckWuhua(bazi *model.Bazi, wuhe *model.TianganRelation) bool {
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

func (s *RelationScanner) CheckDizhiHehua(bazi *model.Bazi, heRelation *model.DizhiRelation) bool {
	monthWuxing := bazi.GetMonthWuxing()
	
	if monthWuxing == heRelation.ResultWuxing {
		return true
	}
	
	tiangans := bazi.GetTiangans()
	for _, t := range tiangans {
		if t.Wuxing() == heRelation.ResultWuxing {
			return true
		}
	}
	
	return false
}

func (s *RelationScanner) GetHighestPriorityRelation(relations []model.DizhiRelation) *model.DizhiRelation {
	if len(relations) == 0 {
		return nil
	}
	
	highest := &relations[0]
	for i := 1; i < len(relations); i++ {
		if relations[i].Type.Priority() < highest.Type.Priority() {
			highest = &relations[i]
		}
	}
	return highest
}

func (s *RelationScanner) HasSanhuiOrSanhe(relations []model.DizhiRelation) bool {
	for _, r := range relations {
		if r.Type == model.RelationSanhui || r.Type == model.RelationSanhe {
			return true
		}
	}
	return false
}

func (s *RelationScanner) GetSanhuiOrSanhe(relations []model.DizhiRelation) *model.DizhiRelation {
	for _, r := range relations {
		if r.Type == model.RelationSanhui || r.Type == model.RelationSanhe {
			return &r
		}
	}
	return nil
}

func (s *RelationScanner) FilterRelationsByPriority(relations []model.DizhiRelation, priority int) []model.DizhiRelation {
	var result []model.DizhiRelation
	for _, r := range relations {
		if r.Type.Priority() <= priority {
			result = append(result, r)
		}
	}
	return result
}

func (s *RelationScanner) GetDizhisInRelations(relations []model.DizhiRelation) map[dizhi.Dizhi]bool {
	used := make(map[dizhi.Dizhi]bool)
	for _, r := range relations {
		for _, d := range r.Dizhis {
			used[d] = true
		}
	}
	return used
}

func (s *RelationScanner) GetWuxingFromRelation(relType model.RelationType) wuxing.Wuxing {
	return wuxing.Tu
}
