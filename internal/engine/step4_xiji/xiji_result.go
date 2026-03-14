package step4_xiji

import (
	"bazi-analyzer/internal/engine/step2_rebuild"
	"bazi-analyzer/internal/engine/step3_wangshuai"
	"bazi-analyzer/internal/model/base"
	"strings"
)

type BingYuanType string

const (
	BingYuanTiaoHou  BingYuanType = "调候失衡"
	BingYuanTongGuan BingYuanType = "五行相战"
	BingYuanFuYi     BingYuanType = "扶抑失衡"
	BingYuanTeShu    BingYuanType = "特殊格局"
)

type Step4XiJiResult struct {
	DayMaster          base.TianGan
	DayMasterWuXing    base.WuXing
	WangShuaiType      step3_wangshuai.WangShuaiType
	HeXinBingYuan      string
	BingYuanType       BingYuanType
	YongShen           []base.WuXing
	YongShenShiShen    []string
	JiShen             []base.WuXing
	JiShenShiShen      []string
	XiShen             []base.WuXing
	XiShenShiShen      []string
	FavorableWuXing    []base.WuXing
	UnfavorableWuXing  []base.WuXing
	FavorableShiShen   []string
	UnfavorableShiShen []string
	Reason             string
	TiaoHouAnalysis    string
	TongGuanAnalysis   string
	FuYiAnalysis       string
	GeJuAnalysis       string
}

type bingYuanInfo struct {
	Type     BingYuanType
	Desc     string
	Severity int
}

func AnalyzeXiJi(chart *base.BaZiChart, step3Result *step3_wangshuai.Step3WangShuaiResult, step2Result *step2_rebuild.Step2RebuildResult) *Step4XiJiResult {
	result := &Step4XiJiResult{
		DayMaster:       step3Result.DayMaster,
		DayMasterWuXing: step3Result.DayMasterWuXing,
		WangShuaiType:   step3Result.WangShuaiType,
	}

	if step3Result.WangShuaiType == step3_wangshuai.CongQiang || step3Result.WangShuaiType == step3_wangshuai.CongRuo {
		analyzeSpecialGeJu(result, chart, step2Result)
		return result
	}

	wxCount := countWuXing(chart, step2Result)
	hasXiangZhan := analyzeTongGuan(result, wxCount)
	tiaoHouResult := analyzeTiaoHou(result, chart, step2Result)

	var byInfo *bingYuanInfo
	if tiaoHouResult != nil && tiaoHouResult.Severity >= 8 {
		byInfo = tiaoHouResult
	} else if hasXiangZhan {
		byInfo = &bingYuanInfo{Type: BingYuanTongGuan, Desc: result.TongGuanAnalysis, Severity: 7}
	} else {
		analyzeFuYi(result, step3Result)
		byInfo = &bingYuanInfo{Type: BingYuanFuYi, Desc: result.FuYiAnalysis, Severity: 6}
	}

	result.HeXinBingYuan = byInfo.Desc
	result.BingYuanType = byInfo.Type

	determineYaoShen(result, byInfo, wxCount, step3Result)

	normalizeResult(result)

	return result
}

func analyzeSpecialGeJu(result *Step4XiJiResult, chart *base.BaZiChart, step2Result *step2_rebuild.Step2RebuildResult) {
	result.BingYuanType = BingYuanTeShu
	if result.WangShuaiType == step3_wangshuai.CongQiang {
		result.HeXinBingYuan = "日主气势极旺，全局顺其气势，构成从强格，宜顺势不宜逆"
		result.YongShen = []base.WuXing{result.DayMasterWuXing, getWuXingShengWo(result.DayMasterWuXing)}
		result.YongShenShiShen = getShiShenByWuXing(result.YongShen, result.DayMasterWuXing)
		result.JiShen = getKeXieHaoWuXing(result.DayMasterWuXing)
		result.JiShenShiShen = getShiShenByWuXing(result.JiShen, result.DayMasterWuXing)
		result.Reason = "从强格病在气势过旺，药在顺势，喜比劫印星，忌克泄耗"
	} else if result.WangShuaiType == step3_wangshuai.CongRuo {
		wxCount := countWuXing(chart, step2Result)
		var mainWuXing base.WuXing
		maxCount := 0
		for wx, cnt := range wxCount {
			if cnt > maxCount && wx != result.DayMasterWuXing && wx != getWuXingShengWo(result.DayMasterWuXing) {
				maxCount = cnt
				mainWuXing = wx
			}
		}
		result.HeXinBingYuan = "日主无依，全局气势专一，构成从弱格，宜从势不宜扶"
		result.YongShen = []base.WuXing{mainWuXing}
		if sheng := base.WuXingSheng[mainWuXing]; sheng != "" {
			result.YongShen = append(result.YongShen, sheng)
		}
		result.YongShenShiShen = getShiShenByWuXing(result.YongShen, result.DayMasterWuXing)
		result.JiShen = []base.WuXing{result.DayMasterWuXing, getWuXingShengWo(result.DayMasterWuXing)}
		result.JiShenShiShen = getShiShenByWuXing(result.JiShen, result.DayMasterWuXing)
		result.Reason = "从弱格病在日主无气，药在从势，喜克泄耗，忌生扶"
	}
	result.GeJuAnalysis = "特殊格局：" + string(result.WangShuaiType)
	normalizeResult(result)
}

func analyzeTiaoHou(result *Step4XiJiResult, chart *base.BaZiChart, step2Result *step2_rebuild.Step2RebuildResult) *bingYuanInfo {
	month := chart.GetMonthLing().Name
	wxCount := countWuXing(chart, step2Result)

	var seasonDesc string
	var isCold, isHot bool

	switch month {
	case "子", "丑", "亥":
		seasonDesc = "生于寒冬，金水旺相，火土衰弱，气候过寒"
		isCold = true
	case "寅", "卯", "辰":
		seasonDesc = "生于春季，木旺土虚，气候温和"
	case "巳", "午", "未":
		seasonDesc = "生于盛夏，火旺土燥，金水受伤，气候过热"
		isHot = true
	case "申", "酉", "戌":
		seasonDesc = "生于秋季，金旺木衰，气候凉爽"
	}

	fireCount := wxCount[base.Fire]
	waterCount := wxCount[base.Water]
	woodCount := wxCount[base.Wood]

	var analysis strings.Builder
	analysis.WriteString(seasonDesc)

	if isCold {
		if fireCount == 0 {
			analysis.WriteString("，全局无火调候，寒金冻水，生机受阻")
			result.TiaoHouAnalysis = analysis.String()
			return &bingYuanInfo{Type: BingYuanTiaoHou, Desc: result.TiaoHouAnalysis, Severity: 10}
		} else if fireCount == 1 {
			analysis.WriteString("，火微不足以暖局，仍有偏寒之病")
			result.TiaoHouAnalysis = analysis.String()
			return &bingYuanInfo{Type: BingYuanTiaoHou, Desc: result.TiaoHouAnalysis, Severity: 7}
		}
	} else if isHot {
		if waterCount == 0 {
			if woodCount > 0 {
				analysis.WriteString("，全局无水调候，木火通明，燥气过重")
			} else {
				analysis.WriteString("，全局无水调候，火炎土燥，万物焦枯")
			}
			result.TiaoHouAnalysis = analysis.String()
			return &bingYuanInfo{Type: BingYuanTiaoHou, Desc: result.TiaoHouAnalysis, Severity: 10}
		} else if waterCount == 1 {
			analysis.WriteString("，水微不足以解热，仍有偏燥之病")
			result.TiaoHouAnalysis = analysis.String()
			return &bingYuanInfo{Type: BingYuanTiaoHou, Desc: result.TiaoHouAnalysis, Severity: 7}
		}
	}

	analysis.WriteString("，调候基本得宜")
	result.TiaoHouAnalysis = analysis.String()
	return nil
}

func analyzeTongGuan(result *Step4XiJiResult, wxCount map[base.WuXing]int) bool {
	keys := []base.WuXing{base.Wood, base.Fire, base.Earth, base.Metal, base.Water}
	var strongWx []base.WuXing
	for _, wx := range keys {
		if wxCount[wx] >= 3 {
			strongWx = append(strongWx, wx)
		}
	}

	if len(strongWx) >= 2 {
		for i, wx1 := range strongWx {
			for _, wx2 := range strongWx[i+1:] {
				if base.Ke(wx1, wx2) || base.Ke(wx2, wx1) {
					var keZhe, beiKe base.WuXing
					if base.Ke(wx1, wx2) {
						keZhe, beiKe = wx1, wx2
					} else {
						keZhe, beiKe = wx2, wx1
					}
					tongGuanShen := findTongGuan(keZhe, beiKe)
					if wxCount[tongGuanShen] == 0 {
						result.TongGuanAnalysis = string(keZhe) + "旺克" + string(beiKe) + "，五行相战，无" + string(tongGuanShen) + "通关"
						return true
					}
				}
			}
		}
	}

	hasKeZhan := false
	var sb strings.Builder
	for i, wx1 := range keys {
		for j := i + 1; j < len(keys); j++ {
			wx2 := keys[j]
			if wxCount[wx1] >= 2 && wxCount[wx2] >= 2 {
				if base.Ke(wx1, wx2) || base.Ke(wx2, wx1) {
					var keZhe, beiKe base.WuXing
					if base.Ke(wx1, wx2) {
						keZhe, beiKe = wx1, wx2
					} else {
						keZhe, beiKe = wx2, wx1
					}
					tongGuanShen := findTongGuan(keZhe, beiKe)
					if wxCount[tongGuanShen] == 0 {
						if !hasKeZhan {
							sb.WriteString("命局")
							hasKeZhan = true
						} else {
							sb.WriteString("；")
						}
						sb.WriteString(string(keZhe) + "与" + string(beiKe) + "相战，无" + string(tongGuanShen) + "通关")
					}
				}
			}
		}
	}

	if hasKeZhan {
		result.TongGuanAnalysis = sb.String()
		return true
	}

	result.TongGuanAnalysis = "五行流通，无严重相战"
	return false
}

func findTongGuan(keZhe, beiKe base.WuXing) base.WuXing {
	if base.WuXingSheng[keZhe] == beiKe {
		return ""
	}
	for wx, sheng := range base.WuXingSheng {
		if sheng == beiKe && base.WuXingSheng[keZhe] == wx {
			return wx
		}
	}
	return ""
}

func analyzeFuYi(result *Step4XiJiResult, step3Result *step3_wangshuai.Step3WangShuaiResult) {
	if step3Result.WangShuaiType == step3_wangshuai.ShenWang {
		result.FuYiAnalysis = "日主身旺，过旺则需克泄耗"
	} else {
		result.FuYiAnalysis = "日主身弱，过弱则需生扶"
	}
}

func determineYaoShen(result *Step4XiJiResult, byInfo *bingYuanInfo, wxCount map[base.WuXing]int, step3Result *step3_wangshuai.Step3WangShuaiResult) {
	dm := result.DayMasterWuXing

	switch byInfo.Type {
	case BingYuanTiaoHou:
		if strings.Contains(byInfo.Desc, "过寒") || strings.Contains(byInfo.Desc, "偏寒") {
			result.YongShen = []base.WuXing{base.Fire}
			if wxCount[base.Wood] == 0 {
				result.YongShen = append(result.YongShen, base.Wood)
			}
			result.JiShen = []base.WuXing{base.Water, base.Metal}
		} else if strings.Contains(byInfo.Desc, "过热") || strings.Contains(byInfo.Desc, "偏燥") {
			result.YongShen = []base.WuXing{base.Water}
			if wxCount[base.Metal] == 0 {
				result.YongShen = append(result.YongShen, base.Metal)
			}
			result.JiShen = []base.WuXing{base.Fire, base.Wood}
		}
		result.Reason = "调候为急，以" + joinWuXing(result.YongShen) + "暖局/解热为用"

	case BingYuanTongGuan:
		for wx := range base.WuXingSheng {
			if strings.Contains(byInfo.Desc, "无"+string(wx)+"通关") {
				result.YongShen = []base.WuXing{wx}
				result.JiShen = extractKeZheFromDesc(byInfo.Desc)
				result.Reason = "通关为要，以" + string(wx) + "调和" + joinWuXing(result.JiShen) + "之争"
				break
			}
		}
		if len(result.YongShen) == 0 {
			parts := strings.Split(byInfo.Desc, "无")
			if len(parts) > 1 {
				tg := strings.Split(parts[1], "通关")[0]
				result.YongShen = []base.WuXing{base.WuXing(tg)}
			}
		}

	case BingYuanFuYi:
		if step3Result.WangShuaiType == step3_wangshuai.ShenWang {
			result.YongShen = getKeXieHaoWuXing(dm)
			result.JiShen = []base.WuXing{getWuXingShengWo(dm), dm}
			result.Reason = "身旺喜克泄耗，以" + joinWuXing(result.YongShen) + "为用"
		} else {
			result.YongShen = []base.WuXing{getWuXingShengWo(dm), dm}
			result.JiShen = getKeXieHaoWuXing(dm)
			result.Reason = "身弱喜生扶，以" + joinWuXing(result.YongShen) + "为用"
		}
	}

	result.YongShenShiShen = getShiShenByWuXing(result.YongShen, dm)
	result.JiShenShiShen = getShiShenByWuXing(result.JiShen, dm)

	result.XiShen = findXiShen(result.YongShen, result.JiShen, wxCount)
	result.XiShenShiShen = getShiShenByWuXing(result.XiShen, dm)
}

func extractKeZheFromDesc(desc string) []base.WuXing {
	var result []base.WuXing
	allWx := []string{"木", "火", "土", "金", "水"}
	for _, wx := range allWx {
		if strings.Contains(desc, wx+"旺") || strings.Contains(desc, wx+"与") {
			result = append(result, base.WuXing(wx))
		}
	}
	return result
}

func findXiShen(yongShen, jiShen []base.WuXing, wxCount map[base.WuXing]int) []base.WuXing {
	isYongShen := make(map[base.WuXing]bool)
	isJiShen := make(map[base.WuXing]bool)
	for _, wx := range yongShen {
		isYongShen[wx] = true
	}
	for _, wx := range jiShen {
		isJiShen[wx] = true
	}

	var xiShen []base.WuXing
	for _, wx := range yongShen {
		shengZhe := getWuXingShengWo(wx)
		if !isYongShen[shengZhe] && !isJiShen[shengZhe] && shengZhe != "" {
			xiShen = append(xiShen, shengZhe)
		}
	}

	allWx := []base.WuXing{base.Wood, base.Fire, base.Earth, base.Metal, base.Water}
	for _, wx := range allWx {
		if !isYongShen[wx] && !isJiShen[wx] {
			already := false
			for _, x := range xiShen {
				if x == wx {
					already = true
					break
				}
			}
			if !already {
				xiShen = append(xiShen, wx)
			}
		}
	}

	return uniqueWuXing(xiShen)
}

func countWuXing(chart *base.BaZiChart, step2Result *step2_rebuild.Step2RebuildResult) map[base.WuXing]int {
	count := make(map[base.WuXing]int)
	for _, tg := range chart.GetAllTianGan() {
		if step2Result.IsTianGanAbsorbed(tg) {
			continue
		}
		wx := step2_rebuild.GetRebuiltTianGanWuXing(step2Result, tg)
		count[wx]++
	}
	for _, dz := range chart.GetAllDiZhi() {
		if step2Result.IsDiZhiAbsorbed(dz) {
			continue
		}
		wx := step2_rebuild.GetRebuiltDiZhiWuXing(step2Result, dz)
		count[wx]++
	}
	return count
}

func getWuXingShengWo(dm base.WuXing) base.WuXing {
	for k, v := range base.WuXingSheng {
		if v == dm {
			return k
		}
	}
	return ""
}

func getWuXingKeWo(dm base.WuXing) base.WuXing {
	for k, v := range base.WuXingKe {
		if v == dm {
			return k
		}
	}
	return ""
}

func getKeXieHaoWuXing(dm base.WuXing) []base.WuXing {
	return []base.WuXing{
		getWuXingKeWo(dm),
		base.WuXingSheng[dm],
		base.WuXingKe[dm],
	}
}

func getShiShenByWuXing(wxList []base.WuXing, dm base.WuXing) []string {
	var result []string
	for _, wx := range wxList {
		result = append(result, getShiShen(wx, dm))
	}
	return uniqueStrings(result)
}

func getShiShen(wx, dm base.WuXing) string {
	if wx == dm {
		return "比劫"
	}
	if base.WuXingSheng[wx] == dm {
		return "印星"
	}
	if base.WuXingSheng[dm] == wx {
		return "食伤"
	}
	if base.Ke(wx, dm) {
		return "官杀"
	}
	if base.Ke(dm, wx) {
		return "财星"
	}
	return ""
}

func uniqueWuXing(wx []base.WuXing) []base.WuXing {
	seen := make(map[base.WuXing]bool)
	var result []base.WuXing
	for _, w := range wx {
		if w != "" && !seen[w] {
			seen[w] = true
			result = append(result, w)
		}
	}
	return result
}

func uniqueStrings(s []string) []string {
	seen := make(map[string]bool)
	var result []string
	for _, v := range s {
		if v != "" && !seen[v] {
			seen[v] = true
			result = append(result, v)
		}
	}
	return result
}

func joinWuXing(wx []base.WuXing) string {
	strs := make([]string, len(wx))
	for i, w := range wx {
		strs[i] = string(w)
	}
	return strings.Join(strs, "、")
}

func normalizeResult(result *Step4XiJiResult) {
	result.FavorableWuXing = uniqueWuXing(append(result.YongShen, result.XiShen...))
	result.UnfavorableWuXing = uniqueWuXing(result.JiShen)
	result.FavorableShiShen = uniqueStrings(append(result.YongShenShiShen, result.XiShenShiShen...))
	result.UnfavorableShiShen = uniqueStrings(result.JiShenShiShen)
}
