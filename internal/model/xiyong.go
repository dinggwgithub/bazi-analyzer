package model

import "bazi-analyzer/pkg/wuxing"

type XiyongResult struct {
	WangshuaiType WangshuaiType
	Xishen       []wuxing.Wuxing
	Yongshen     wuxing.Wuxing
	Jishen       []wuxing.Wuxing
	XishenDesc   string
	JishenDesc   string
}

func GetXiyong(wangshuaiType WangshuaiType, riZhuWuxing wuxing.Wuxing) XiyongResult {
	result := XiyongResult{
		WangshuaiType: wangshuaiType,
	}
	
	switch wangshuaiType {
	case WangshuaiShenruo:
		result.Xishen = []wuxing.Wuxing{
			riZhuWuxing,
			riZhuWuxing.BeiSheng(),
		}
		result.Jishen = []wuxing.Wuxing{
			riZhuWuxing.Ke(),
			riZhuWuxing.Sheng(),
			riZhuWuxing.BeiKe(),
		}
		result.Yongshen = riZhuWuxing.BeiSheng()
		result.XishenDesc = "身弱喜生扶，喜印星、比劫"
		result.JishenDesc = "身弱忌克泄耗，忌官杀、食伤、财星"
		
	case WangshuaiShenwang:
		result.Xishen = []wuxing.Wuxing{
			riZhuWuxing.Ke(),
			riZhuWuxing.Sheng(),
			riZhuWuxing.BeiKe(),
		}
		result.Jishen = []wuxing.Wuxing{
			riZhuWuxing,
			riZhuWuxing.BeiSheng(),
		}
		result.Yongshen = riZhuWuxing.Ke()
		result.XishenDesc = "身旺喜克泄耗，喜官杀、食伤、财星"
		result.JishenDesc = "身旺忌生扶，忌印星、比劫"
		
	case WangshuaiCongqiang:
		result.Xishen = []wuxing.Wuxing{
			riZhuWuxing,
			riZhuWuxing.BeiSheng(),
		}
		result.Jishen = []wuxing.Wuxing{
			riZhuWuxing.Ke(),
			riZhuWuxing.Sheng(),
			riZhuWuxing.BeiKe(),
		}
		result.Yongshen = riZhuWuxing
		result.XishenDesc = "从强顺其气势，喜印星、比劫"
		result.JishenDesc = "从强忌逆势，忌官杀、食伤、财星"
		
	case WangshuaiCongruo:
		result.Xishen = []wuxing.Wuxing{
			riZhuWuxing.Ke(),
			riZhuWuxing.Sheng(),
			riZhuWuxing.BeiKe(),
		}
		result.Jishen = []wuxing.Wuxing{
			riZhuWuxing,
			riZhuWuxing.BeiSheng(),
		}
		result.Yongshen = riZhuWuxing.Ke()
		result.XishenDesc = "从弱顺其气势，喜官杀、食伤、财星"
		result.JishenDesc = "从弱忌生扶，忌印星、比劫"
	}
	
	return result
}
