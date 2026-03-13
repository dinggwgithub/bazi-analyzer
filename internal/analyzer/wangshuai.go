package analyzer

import (
	"bazi-analyzer/internal/model"
	"bazi-analyzer/pkg/dizhi"
	"bazi-analyzer/pkg/tiangan"
	"bazi-analyzer/pkg/wuxing"
)

type WangshuaiAnalyzer struct{}

func NewWangshuaiAnalyzer() *WangshuaiAnalyzer {
	return &WangshuaiAnalyzer{}
}

func (a *WangshuaiAnalyzer) Analyze(bazi *model.Bazi, reconstructResult *ReconstructResult) *model.WangshuaiResult {
	riZhu := bazi.GetRizhuTiangan()
	riZhuWuxing := riZhu.Wuxing()
	
	deling := a.analyzeDeling(bazi, riZhuWuxing, reconstructResult)
	didi := a.analyzeDidi(bazi, riZhu, riZhuWuxing, reconstructResult)
	deshi := a.analyzeDeshi(bazi, riZhu, riZhuWuxing)
	kexieha := a.analyzeKexieha(bazi, riZhu, riZhuWuxing)
	
	totalScore := deling.Score + didi.Score + deshi.Score
	
	wangshuaiType := a.determineWangshuaiType(totalScore, kexieha.TotalScore, deling.IsDeling)
	
	return &model.WangshuaiResult{
		Type:        wangshuaiType,
		DelingScore: deling.Score,
		DidiScore:   didi.Score,
		DeshiScore:  deshi.Score,
		TotalScore:  totalScore,
		Deling:      deling,
		Didi:        didi,
		Deshi:       deshi,
		Kexieha:     kexieha,
	}
}

func (a *WangshuaiAnalyzer) analyzeDeling(bazi *model.Bazi, riZhuWuxing wuxing.Wuxing, reconstructResult *ReconstructResult) model.DelingAnalysis {
	var monthWuxing wuxing.Wuxing
	
	if reconstructResult != nil && reconstructResult.MajorRelation != nil {
		monthDizhi := bazi.GetYuezhi()
		if reconstructResult.MajorRelation.Type == model.RelationSanhui || 
		   reconstructResult.MajorRelation.Type == model.RelationSanhe {
			for _, d := range reconstructResult.MajorRelation.Dizhis {
				if d == monthDizhi {
					monthWuxing = reconstructResult.MajorRelation.ResultWuxing
					break
				}
			}
		}
	}
	
	if monthWuxing == 0 {
		monthWuxing = bazi.GetMonthWuxing()
	}
	
	relation := riZhuWuxing.Relation(monthWuxing)
	
	score := 0
	isDeling := false
	
	if riZhuWuxing == monthWuxing {
		score = 5
		isDeling = true
	} else if riZhuWuxing.BeiSheng() == monthWuxing {
		score = 4
		isDeling = true
	} else if riZhuWuxing.Ke() == monthWuxing {
		score = -2
	} else if riZhuWuxing.BeiKe() == monthWuxing {
		score = -3
	} else if riZhuWuxing.Sheng() == monthWuxing {
		score = -1
	}
	
	relationStr := "无特殊关系"
	switch relation {
	case wuxing.RelationSame:
		relationStr = "同五行（当令）"
	case wuxing.RelationBeiSheng:
		relationStr = "被月令所生（印星当令）"
	case wuxing.RelationKe:
		relationStr = "克月令（耗气）"
	case wuxing.RelationBeiKe:
		relationStr = "被月令所克（受制）"
	case wuxing.RelationSheng:
		relationStr = "生月令（泄气）"
	}
	
	return model.DelingAnalysis{
		IsDeling:    isDeling,
		MonthWuxing: monthWuxing,
		RiZhuWuxing: riZhuWuxing,
		Relation:    relationStr,
		Score:       score,
	}
}

func (a *WangshuaiAnalyzer) analyzeDidi(bazi *model.Bazi, riZhu tiangan.Tiangan, riZhuWuxing wuxing.Wuxing, reconstructResult *ReconstructResult) model.DidiAnalysis {
	pillars := bazi.GetPillars()
	rootDizhis := []string{}
	rootType := ""
	hasRoot := false
	score := 0
	
	for _, pillar := range pillars {
		if pillar.Position == model.PositionRi {
			continue
		}
		
		dz := dizhi.Dizhi(pillar.Dizhi.Original)
		
		var dzWuxing wuxing.Wuxing
		if reconstructResult != nil && reconstructResult.MajorRelation != nil {
			dzWuxing = getTransformedWuxing(dz, reconstructResult.MajorRelation)
		} else {
			dzWuxing = dz.Wuxing()
		}
		
		if dzWuxing == riZhuWuxing {
			hasRoot = true
			rootDizhis = append(rootDizhis, dz.String())
			
			benqi := dz.GetBenqi()
			if benqi != nil {
				benqiWuxing := getWuxingFromTianganName(benqi.TianganName)
				if benqiWuxing == riZhuWuxing {
					rootType = "本气根"
					score += 3
				}
			}
		}
		
		if dzWuxing == riZhuWuxing.BeiSheng() {
			hasRoot = true
			rootDizhis = append(rootDizhis, dz.String()+"(印)")
			score += 2
		}
		
		for _, c := range dz.GetCanggan() {
			cgWuxing := getWuxingFromTianganName(c.TianganName)
			if cgWuxing == riZhuWuxing {
				hasRoot = true
				if c.Strength == 1 {
					rootDizhis = append(rootDizhis, dz.String()+"(本气)")
					score += 2
				} else if c.Strength == 2 {
					rootDizhis = append(rootDizhis, dz.String()+"(中气)")
					score += 1
				}
			}
		}
	}
	
	riDz := dizhi.Dizhi(bazi.RiPillar.Dizhi.Original)
	riDzWuxing := riDz.Wuxing()
	if riDzWuxing == riZhuWuxing {
		hasRoot = true
		rootDizhis = append(rootDizhis, riDz.String()+"(坐支)")
		rootType = "坐根"
		score += 3
	}
	
	return model.DidiAnalysis{
		HasRoot:    hasRoot,
		RootType:   rootType,
		RootDizhis: rootDizhis,
		Score:      score,
	}
}

func (a *WangshuaiAnalyzer) analyzeDeshi(bazi *model.Bazi, riZhu tiangan.Tiangan, riZhuWuxing wuxing.Wuxing) model.DeshiAnalysis {
	tiangans := bazi.GetTiangans()
	
	bijieCount := 0
	yinCount := 0
	score := 0
	
	for i, t := range tiangans {
		if i == 2 {
			continue
		}
		
		tWuxing := t.Wuxing()
		isSameYinyang := t.Yinyang() == riZhu.Yinyang()
		
		shishen := model.GetShishen(riZhuWuxing, tWuxing, isSameYinyang)
		
		if shishen.IsBiJie() {
			bijieCount++
			score += 2
		}
		
		if shishen.IsYin() {
			yinCount++
			score += 1
		}
	}
	
	return model.DeshiAnalysis{
		BijieCount: bijieCount,
		YinCount:   yinCount,
		Score:      score,
	}
}

func (a *WangshuaiAnalyzer) analyzeKexieha(bazi *model.Bazi, riZhu tiangan.Tiangan, riZhuWuxing wuxing.Wuxing) model.KexiehaAnalysis {
	tiangans := bazi.GetTiangans()
	dizhis := bazi.GetDizhis()
	
	guanShaScore := 0
	shiShangScore := 0
	caiScore := 0
	
	for i, t := range tiangans {
		if i == 2 {
			continue
		}
		
		tWuxing := t.Wuxing()
		isSameYinyang := t.Yinyang() == riZhu.Yinyang()
		
		shishen := model.GetShishen(riZhuWuxing, tWuxing, isSameYinyang)
		
		if shishen.IsGuanSha() {
			guanShaScore += 2
		}
		if shishen.IsShiShang() {
			shiShangScore += 1
		}
		if shishen.IsCai() {
			caiScore += 1
		}
	}
	
	for _, d := range dizhis {
		dWuxing := d.Wuxing()
		
		if dWuxing == riZhuWuxing.BeiKe() {
			guanShaScore += 2
		}
		if dWuxing == riZhuWuxing.Sheng() {
			shiShangScore += 1
		}
		if dWuxing == riZhuWuxing.Ke() {
			caiScore += 1
		}
	}
	
	return model.KexiehaAnalysis{
		GuanShaScore:  guanShaScore,
		ShiShangScore: shiShangScore,
		CaiScore:      caiScore,
		TotalScore:    guanShaScore + shiShangScore + caiScore,
	}
}

func (a *WangshuaiAnalyzer) determineWangshuaiType(totalScore int, kexiehaScore int, isDeling bool) model.WangshuaiType {
	adjustedScore := totalScore - kexiehaScore/2
	
	if adjustedScore >= 8 {
		return model.WangshuaiShenwang
	} else if adjustedScore <= -3 {
		if isDeling {
			return model.WangshuaiShenruo
		}
		return model.WangshuaiCongruo
	} else if adjustedScore >= 6 && !isDeling {
		return model.WangshuaiCongqiang
	}
	
	return model.WangshuaiShenruo
}

func getTransformedWuxing(d dizhi.Dizhi, relation *model.DizhiRelation) wuxing.Wuxing {
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
