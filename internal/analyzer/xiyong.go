package analyzer

import (
	"bazi-analyzer/internal/model"
	"bazi-analyzer/pkg/wuxing"
)

type XiyongAnalyzer struct{}

func NewXiyongAnalyzer() *XiyongAnalyzer {
	return &XiyongAnalyzer{}
}

func (a *XiyongAnalyzer) Analyze(wangshuaiResult *model.WangshuaiResult, bazi *model.Bazi) *model.XiyongResult {
	riZhuWuxing := bazi.GetRizhuTiangan().Wuxing()
	
	result := model.GetXiyong(wangshuaiResult.Type, riZhuWuxing)
	
	return &result
}

func (a *XiyongAnalyzer) GetXishenWuxing(wangshuaiType model.WangshuaiType, riZhuWuxing wuxing.Wuxing) []wuxing.Wuxing {
	result := model.GetXiyong(wangshuaiType, riZhuWuxing)
	return result.Xishen
}

func (a *XiyongAnalyzer) GetJishenWuxing(wangshuaiType model.WangshuaiType, riZhuWuxing wuxing.Wuxing) []wuxing.Wuxing {
	result := model.GetXiyong(wangshuaiType, riZhuWuxing)
	return result.Jishen
}

func (a *XiyongAnalyzer) IsXishen(wx wuxing.Wuxing, xiyongResult *model.XiyongResult) bool {
	for _, x := range xiyongResult.Xishen {
		if x == wx {
			return true
		}
	}
	return false
}

func (a *XiyongAnalyzer) IsJishen(wx wuxing.Wuxing, xiyongResult *model.XiyongResult) bool {
	for _, j := range xiyongResult.Jishen {
		if j == wx {
			return true
		}
	}
	return false
}
