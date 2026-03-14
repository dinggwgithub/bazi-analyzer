package step4_xiji

import (
	"bazi-analyzer/internal/engine/step2_rebuild"
	"bazi-analyzer/internal/engine/step3_wangshuai"
	"bazi-analyzer/internal/model/base"
	"fmt"
	"sort"
)

// BingYaoType 病源类型
type BingYaoType string

const (
	BingYao_TiaoHou  BingYaoType = "调候病" // 调候病：过寒过热
	BingYao_TongGuan BingYaoType = "通关病" // 通关病：五行相战
	BingYao_FuYi     BingYaoType = "扶抑病" // 扶抑病：日主过旺过弱
	BingYao_GeJu     BingYaoType = "格局病" // 格局病：特殊格局缺陷
	BingYao_ShenRu   BingYaoType = "神煞病" // 神煞病：凶神无制
)

// BingYaoAnalysis 病药分析结果
type BingYaoAnalysis struct {
	BingType   BingYaoType   // 病源类型
	BingDesc   string        // 病源描述
	Severity   int           // 严重程度 1-5
	YaoWuXing  []base.WuXing // 药（用神）- 五行
	YaoShiShen []string      // 药（用神）- 十神
	JiWuXing   []base.WuXing // 忌神 - 五行
	JiShiShen  []string      // 忌神 - 十神
	Reason     string        // 判断依据
}

// Step4XiJiResult 喜忌用神分析结果
type Step4XiJiResult struct {
	DayMaster       base.TianGan                  // 日主
	DayMasterWuXing base.WuXing                   // 日主五行
	WangShuaiType   step3_wangshuai.WangShuaiType // 旺衰类型

	// 多维度诊断结果
	TiaoHouAnalysis  *BingYaoAnalysis // 调候诊断
	TongGuanAnalysis *BingYaoAnalysis // 通关诊断
	FuYiAnalysis     *BingYaoAnalysis // 扶抑诊断
	GeJuAnalysis     *BingYaoAnalysis // 格局诊断
	BingYaoAnalysis  *BingYaoAnalysis // 病药诊断（综合）

	// 核心病源（最主要矛盾）
	CoreBingYao *BingYaoAnalysis

	// 最终喜忌结论
	YongShenWuXing  []base.WuXing // 用神（药）- 五行
	YongShenShiShen []string      // 用神（药）- 十神
	JiShenWuXing    []base.WuXing // 忌神 - 五行
	JiShenShiShen   []string      // 忌神 - 十神
	XiShenWuXing    []base.WuXing // 喜神 - 五行
	XiShenShiShen   []string      // 喜神 - 十神

	// 解释说明
	CoreBingDesc   string // 核心病源描述
	YongShenReason string // 用神判定依据
	Summary        string // 综合结论
}

// AnalyzeXiJi 喜忌用神分析主入口
func AnalyzeXiJi(chart *base.BaZiChart, step2Result *step2_rebuild.Step2RebuildResult, step3Result *step3_wangshuai.Step3WangShuaiResult) *Step4XiJiResult {
	result := &Step4XiJiResult{
		DayMaster:       step3Result.DayMaster,
		DayMasterWuXing: step3Result.DayMasterWuXing,
		WangShuaiType:   step3Result.WangShuaiType,
	}

	// 1. 调候诊断
	result.TiaoHouAnalysis = analyzeTiaoHou(chart, step2Result, step3Result)

	// 2. 通关诊断
	result.TongGuanAnalysis = analyzeTongGuan(chart, step2Result, step3Result)

	// 3. 扶抑诊断
	result.FuYiAnalysis = analyzeFuYi(chart, step2Result, step3Result)

	// 4. 格局诊断
	result.GeJuAnalysis = analyzeGeJu(chart, step2Result, step3Result)

	// 5. 综合病药诊断，确定核心病源
	result.BingYaoAnalysis = analyzeBingYao(chart, step2Result, step3Result, result)

	// 6. 确定核心病源（取最严重的病）
	result.CoreBingYao = determineCoreBingYao(result)

	// 7. 综合判定喜忌用神
	determineFinalXiJi(result)

	// 8. 生成解释说明
	generateExplanation(result)

	return result
}

// analyzeTiaoHou 调候诊断：分析命局冷暖燥湿
func analyzeTiaoHou(chart *base.BaZiChart, step2Result *step2_rebuild.Step2RebuildResult, step3Result *step3_wangshuai.Step3WangShuaiResult) *BingYaoAnalysis {
	analysis := &BingYaoAnalysis{
		BingType: BingYao_TiaoHou,
	}

	// 计算命局寒暖程度
	fireCount := 0.0
	waterCount := 0.0
	winterMonths := map[string]bool{"子": true, "丑": true, "亥": true}
	summerMonths := map[string]bool{"午": true, "巳": true, "未": true}

	monthLing := chart.GetMonthLing()

	// 统计天干五行
	for _, tg := range chart.GetAllTianGan() {
		if step2Result.IsTianGanAbsorbed(tg) {
			continue
		}
		tgWuXing := step2_rebuild.GetRebuiltTianGanWuXing(step2Result, tg)
		if tgWuXing == base.Fire {
			fireCount++
		} else if tgWuXing == base.Water {
			waterCount++
		}
	}

	// 统计地支五行（含藏干）
	for _, dz := range chart.GetAllDiZhi() {
		if step2Result.IsDiZhiAbsorbed(dz) {
			continue
		}
		dzWuXing := step2_rebuild.GetRebuiltDiZhiWuXing(step2Result, dz)
		if dzWuXing == base.Fire {
			fireCount++
		} else if dzWuXing == base.Water {
			waterCount++
		}

		// 藏干也计入
		for _, hidden := range dz.HiddenStems {
			if hidden.TianGan.Element == base.Fire {
				fireCount += 0.5
			} else if hidden.TianGan.Element == base.Water {
				waterCount += 0.5
			}
		}
	}

	// 判断调候病
	isWinter := winterMonths[monthLing.Name]
	isSummer := summerMonths[monthLing.Name]

	if isWinter && fireCount < 1.5 {
		// 冬生无火，命局过寒
		analysis.BingDesc = "冬生无火，命局过寒"
		analysis.Severity = 4
		analysis.YaoWuXing = []base.WuXing{base.Fire}
		analysis.YaoShiShen = []string{"官杀", "印星"}
		analysis.JiWuXing = []base.WuXing{base.Water, base.Metal}
		analysis.JiShiShen = []string{"比劫", "食伤"}
		analysis.Reason = "冬月水旺火死，需火调候暖局"
	} else if isSummer && waterCount < 1.5 {
		// 夏生无水，命局过热
		analysis.BingDesc = "夏生无水，命局过热"
		analysis.Severity = 4
		analysis.YaoWuXing = []base.WuXing{base.Water}
		analysis.YaoShiShen = []string{"食伤", "财星"}
		analysis.JiWuXing = []base.WuXing{base.Fire, base.Wood}
		analysis.JiShiShen = []string{"印星", "官杀"}
		analysis.Reason = "夏月火旺水囚，需水调候润局"
	} else if fireCount > waterCount+3 {
		// 火多水少，偏燥
		analysis.BingDesc = "火多水少，命局偏燥"
		analysis.Severity = 3
		analysis.YaoWuXing = []base.WuXing{base.Water}
		analysis.YaoShiShen = []string{"食伤", "财星"}
		analysis.JiWuXing = []base.WuXing{base.Fire}
		analysis.JiShiShen = []string{"印星"}
		analysis.Reason = "火炎土燥，需水调候"
	} else if waterCount > fireCount+3 {
		// 水多火少，偏寒
		analysis.BingDesc = "水多火少，命局偏寒"
		analysis.Severity = 3
		analysis.YaoWuXing = []base.WuXing{base.Fire}
		analysis.YaoShiShen = []string{"官杀", "印星"}
		analysis.JiWuXing = []base.WuXing{base.Water}
		analysis.JiShiShen = []string{"比劫"}
		analysis.Reason = "水冷金寒，需火调候"
	} else {
		// 调候平衡
		analysis.BingDesc = "调候平衡，无明显寒热之病"
		analysis.Severity = 0
		analysis.Reason = "寒暖燥湿相对平衡"
	}

	return analysis
}

// analyzeTongGuan 通关诊断：分析五行相战是否需要通关
func analyzeTongGuan(chart *base.BaZiChart, step2Result *step2_rebuild.Step2RebuildResult, step3Result *step3_wangshuai.Step3WangShuaiResult) *BingYaoAnalysis {
	analysis := &BingYaoAnalysis{
		BingType: BingYao_TongGuan,
	}

	// 统计各五行力量
	wuxingCount := map[base.WuXing]float64{
		base.Wood:  0,
		base.Fire:  0,
		base.Earth: 0,
		base.Metal: 0,
		base.Water: 0,
	}

	// 统计天干
	for _, tg := range chart.GetAllTianGan() {
		if step2Result.IsTianGanAbsorbed(tg) {
			continue
		}
		tgWuXing := step2_rebuild.GetRebuiltTianGanWuXing(step2Result, tg)
		wuxingCount[tgWuXing] += 1.0
	}

	// 统计地支（含藏干）
	for _, dz := range chart.GetAllDiZhi() {
		if step2Result.IsDiZhiAbsorbed(dz) {
			continue
		}
		dzWuXing := step2_rebuild.GetRebuiltDiZhiWuXing(step2Result, dz)
		wuxingCount[dzWuXing] += 1.0

		for _, hidden := range dz.HiddenStems {
			wuxingCount[hidden.TianGan.Element] += 0.5
		}
	}

	// 检查五行相战
	// 金木相战
	if wuxingCount[base.Metal] >= 2 && wuxingCount[base.Wood] >= 2 {
		analysis.BingDesc = "金木相战，两强对峙"
		analysis.Severity = 4
		analysis.YaoWuXing = []base.WuXing{base.Water}
		analysis.YaoShiShen = []string{"食伤"}
		analysis.JiWuXing = []base.WuXing{base.Metal, base.Wood}
		analysis.JiShiShen = []string{"官杀", "比劫"}
		analysis.Reason = "金克木，水可通关化金生木"
		return analysis
	}

	// 水火相战
	if wuxingCount[base.Water] >= 2 && wuxingCount[base.Fire] >= 2 {
		analysis.BingDesc = "水火相战，寒热不交"
		analysis.Severity = 4
		analysis.YaoWuXing = []base.WuXing{base.Wood}
		analysis.YaoShiShen = []string{"印星"}
		analysis.JiWuXing = []base.WuXing{base.Water, base.Fire}
		analysis.JiShiShen = []string{"比劫", "财星"}
		analysis.Reason = "水克火，木可通关化水生火"
		return analysis
	}

	// 土水相战
	if wuxingCount[base.Earth] >= 2 && wuxingCount[base.Water] >= 2 {
		analysis.BingDesc = "土水相战，浊泥淤积"
		analysis.Severity = 3
		analysis.YaoWuXing = []base.WuXing{base.Metal}
		analysis.YaoShiShen = []string{"官杀"}
		analysis.JiWuXing = []base.WuXing{base.Earth, base.Water}
		analysis.JiShiShen = []string{"财星", "比劫"}
		analysis.Reason = "土克水，金可通关化土生水"
		return analysis
	}

	// 火金相战
	if wuxingCount[base.Fire] >= 2 && wuxingCount[base.Metal] >= 2 {
		analysis.BingDesc = "火金相战，熔金太过"
		analysis.Severity = 3
		analysis.YaoWuXing = []base.WuXing{base.Earth}
		analysis.YaoShiShen = []string{"印星"}
		analysis.JiWuXing = []base.WuXing{base.Fire, base.Metal}
		analysis.JiShiShen = []string{"财星", "官杀"}
		analysis.Reason = "火克金，土可通关化火生金"
		return analysis
	}

	// 木土相战
	if wuxingCount[base.Wood] >= 2 && wuxingCount[base.Earth] >= 2 {
		analysis.BingDesc = "木土相战，土木不和"
		analysis.Severity = 3
		analysis.YaoWuXing = []base.WuXing{base.Fire}
		analysis.YaoShiShen = []string{"食伤"}
		analysis.JiWuXing = []base.WuXing{base.Wood, base.Earth}
		analysis.JiShiShen = []string{"比劫", "财星"}
		analysis.Reason = "木克土，火可通关化木生土"
		return analysis
	}

	analysis.BingDesc = "五行流通，无明显相战之病"
	analysis.Severity = 0
	analysis.Reason = "五行相对平衡，无需通关"

	return analysis
}

// analyzeFuYi 扶抑诊断：基于旺衰的扶抑分析
func analyzeFuYi(chart *base.BaZiChart, step2Result *step2_rebuild.Step2RebuildResult, step3Result *step3_wangshuai.Step3WangShuaiResult) *BingYaoAnalysis {
	analysis := &BingYaoAnalysis{
		BingType: BingYao_FuYi,
	}

	dmWuXing := step3Result.DayMasterWuXing

	switch step3Result.WangShuaiType {
	case step3_wangshuai.ShenRuo:
		analysis.BingDesc = "日主身弱，需生扶"
		analysis.Severity = calculateFuYiSeverity(step3Result)
		analysis.YaoWuXing = []base.WuXing{
			getWuXingShengWo(dmWuXing),
			dmWuXing,
		}
		analysis.YaoShiShen = []string{"印星", "比劫"}
		analysis.JiWuXing = getKeXieHaoWuXing(dmWuXing)
		analysis.JiShiShen = []string{"官杀", "食伤", "财星"}
		analysis.Reason = step3Result.Reason + "，身弱喜生扶"

	case step3_wangshuai.ShenWang:
		analysis.BingDesc = "日主身旺，需克泄耗"
		analysis.Severity = calculateFuYiSeverity(step3Result)
		analysis.YaoWuXing = getKeXieHaoWuXing(dmWuXing)
		analysis.YaoShiShen = []string{"官杀", "食伤", "财星"}
		analysis.JiWuXing = []base.WuXing{
			getWuXingShengWo(dmWuXing),
			dmWuXing,
		}
		analysis.JiShiShen = []string{"印星", "比劫"}
		analysis.Reason = step3Result.Reason + "，身旺喜克泄耗"

	case step3_wangshuai.CongQiang:
		analysis.BingDesc = "从强格，顺其旺势"
		analysis.Severity = 5
		analysis.YaoWuXing = []base.WuXing{
			dmWuXing,
			getWuXingShengWo(dmWuXing),
		}
		analysis.YaoShiShen = []string{"比劫", "印星"}
		analysis.JiWuXing = getKeXieHaoWuXing(dmWuXing)
		analysis.JiShiShen = []string{"官杀", "食伤", "财星"}
		analysis.Reason = "从强格不可逆其势，喜生扶忌克泄"

	case step3_wangshuai.CongRuo:
		analysis.BingDesc = "从弱格，顺其弱势"
		analysis.Severity = 5
		analysis.YaoWuXing = getKeXieHaoWuXing(dmWuXing)
		analysis.YaoShiShen = []string{"官杀", "食伤", "财星"}
		analysis.JiWuXing = []base.WuXing{
			dmWuXing,
			getWuXingShengWo(dmWuXing),
		}
		analysis.JiShiShen = []string{"比劫", "印星"}
		analysis.Reason = "从弱格不可逆其势，喜克泄忌生扶"
	}

	return analysis
}

// analyzeGeJu 格局诊断：特殊格局分析
func analyzeGeJu(chart *base.BaZiChart, step2Result *step2_rebuild.Step2RebuildResult, step3Result *step3_wangshuai.Step3WangShuaiResult) *BingYaoAnalysis {
	analysis := &BingYaoAnalysis{
		BingType: BingYao_GeJu,
	}

	// 检查是否为特殊格局
	if step3Result.WangShuaiType == step3_wangshuai.CongQiang ||
		step3Result.WangShuaiType == step3_wangshuai.CongRuo {
		// 从格已在扶抑中处理
		analysis.BingDesc = "特殊格局（从格）"
		analysis.Severity = 0
		analysis.Reason = "从格按从格规则论"
		return analysis
	}

	// 检查化气格
	if isHuaQiGe(chart, step2Result) {
		analysis.BingDesc = "化气格"
		analysis.Severity = 5
		// 化气格用神为化神之五行
		huaShen := determineHuaShen(chart)
		analysis.YaoWuXing = []base.WuXing{huaShen}
		analysis.JiWuXing = []base.WuXing{getKeZhiHuaShen(huaShen)}
		analysis.Reason = "日主从化，喜化神之五行"
		return analysis
	}

	// 检查专旺格
	if isZhuanWangGe(chart, step2Result) {
		analysis.BingDesc = "专旺格"
		analysis.Severity = 5
		zhuWuXing := getZhuWangWuXing(chart, step2Result)
		analysis.YaoWuXing = []base.WuXing{zhuWuXing}
		analysis.JiWuXing = getKeXieHaoWuXing(zhuWuXing)
		analysis.Reason = "一行独旺，喜顺其旺势"
		return analysis
	}

	analysis.BingDesc = "普通格局"
	analysis.Severity = 0
	analysis.Reason = "按普通格局论"

	return analysis
}

// analyzeBingYao 综合病药诊断
func analyzeBingYao(chart *base.BaZiChart, step2Result *step2_rebuild.Step2RebuildResult, step3Result *step3_wangshuai.Step3WangShuaiResult, xijiResult *Step4XiJiResult) *BingYaoAnalysis {
	analysis := &BingYaoAnalysis{
		BingType: BingYao_ShenRu,
	}

	// 收集所有病源
	bingYuanList := []*BingYaoAnalysis{
		xijiResult.TiaoHouAnalysis,
		xijiResult.TongGuanAnalysis,
		xijiResult.FuYiAnalysis,
		xijiResult.GeJuAnalysis,
	}

	// 找出最严重的病
	var maxSeverityBing *BingYaoAnalysis
	maxSeverity := 0

	for _, bing := range bingYuanList {
		if bing != nil && bing.Severity > maxSeverity {
			maxSeverity = bing.Severity
			maxSeverityBing = bing
		}
	}

	if maxSeverityBing != nil {
		analysis.BingDesc = maxSeverityBing.BingDesc
		analysis.Severity = maxSeverityBing.Severity
		analysis.YaoWuXing = maxSeverityBing.YaoWuXing
		analysis.YaoShiShen = maxSeverityBing.YaoShiShen
		analysis.JiWuXing = maxSeverityBing.JiWuXing
		analysis.JiShiShen = maxSeverityBing.JiShiShen
		analysis.Reason = fmt.Sprintf("综合诊断，以%s为主", maxSeverityBing.BingType)
	} else {
		analysis.BingDesc = "命局相对平衡"
		analysis.Severity = 1
		analysis.Reason = "各维度无明显大病"
	}

	return analysis
}

// determineCoreBingYao 确定核心病源
func determineCoreBingYao(result *Step4XiJiResult) *BingYaoAnalysis {
	// 按优先级和严重程度确定核心病源
	// 优先级：调候 > 通关 > 格局 > 扶抑
	// 遵循"有病方为贵"的原则，优先解决最突出的矛盾

	// 1. 调候病（寒热燥湿）- 优先级最高
	// 调候是命局生存的基础，冬生无火、夏生无水都是大病
	if result.TiaoHouAnalysis != nil && result.TiaoHouAnalysis.Severity >= 4 {
		return result.TiaoHouAnalysis
	}

	// 2. 通关病（五行相战）- 优先级次高
	// 五行相战会导致命局内耗严重
	if result.TongGuanAnalysis != nil && result.TongGuanAnalysis.Severity >= 4 {
		return result.TongGuanAnalysis
	}

	// 3. 格局病（从格、化气、专旺等）- 第三优先级
	if result.GeJuAnalysis != nil && result.GeJuAnalysis.Severity >= 4 {
		return result.GeJuAnalysis
	}

	// 4. 调候病（严重程度3）
	if result.TiaoHouAnalysis != nil && result.TiaoHouAnalysis.Severity >= 3 {
		return result.TiaoHouAnalysis
	}

	// 5. 扶抑病（日主旺衰）- 基础分析
	if result.FuYiAnalysis != nil && result.FuYiAnalysis.Severity >= 2 {
		return result.FuYiAnalysis
	}

	// 默认返回扶抑分析
	return result.FuYiAnalysis
}

// determineFinalXiJi 综合判定最终喜忌
func determineFinalXiJi(result *Step4XiJiResult) {
	core := result.CoreBingYao
	if core == nil {
		return
	}

	// 用神（药）
	result.YongShenWuXing = uniqueWuXing(core.YaoWuXing)
	result.YongShenShiShen = uniqueStrings(core.YaoShiShen)

	// 忌神
	result.JiShenWuXing = uniqueWuXing(core.JiWuXing)
	result.JiShenShiShen = uniqueStrings(core.JiShiShen)

	// 喜神：辅助用神的五行
	xiShenMap := make(map[base.WuXing]bool)

	// 生助用神的为喜神
	for _, ys := range result.YongShenWuXing {
		// 生用神者为喜
		shengZhe := getWuXingShengWo(ys)
		xiShenMap[shengZhe] = true
	}

	// 与用神同类的也是喜
	for _, ys := range result.YongShenWuXing {
		xiShenMap[ys] = true
	}

	// 排除忌神
	for _, js := range result.JiShenWuXing {
		delete(xiShenMap, js)
	}

	// 转换为切片
	for wx := range xiShenMap {
		result.XiShenWuXing = append(result.XiShenWuXing, wx)
	}

	// 喜神十神
	result.XiShenShiShen = deduceShiShenFromWuXing(result.DayMasterWuXing, result.XiShenWuXing)
}

// generateExplanation 生成解释说明
func generateExplanation(result *Step4XiJiResult) {
	core := result.CoreBingYao

	// 核心病源描述
	if core != nil {
		result.CoreBingDesc = fmt.Sprintf("【%s】%s（严重程度：%d/5）",
			core.BingType, core.BingDesc, core.Severity)
	}

	// 用神判定依据
	result.YongShenReason = fmt.Sprintf("以%s为总纲，%s",
		core.BingType, core.Reason)

	// 综合结论
	result.Summary = fmt.Sprintf(
		"命局%s。用神为%s，喜神为%s，忌神为%s。",
		core.BingDesc,
		wuxingSliceToString(result.YongShenWuXing),
		wuxingSliceToString(result.XiShenWuXing),
		wuxingSliceToString(result.JiShenWuXing),
	)
}

// calculateFuYiSeverity 计算扶抑严重程度
func calculateFuYiSeverity(step3Result *step3_wangshuai.Step3WangShuaiResult) int {
	diff := step3Result.ShengFuCount - step3Result.KeXieHaoCount
	if diff < 0 {
		diff = -diff
	}

	if diff >= 4 {
		return 4
	} else if diff >= 2 {
		return 3
	}
	return 2
}

// isHuaQiGe 判断是否化气格
func isHuaQiGe(chart *base.BaZiChart, step2Result *step2_rebuild.Step2RebuildResult) bool {
	// 简化判断：日干与月干或时干相合
	dayMaster := chart.GetDayMaster()
	monthGan := chart.MonthPillar.TianGan
	hourGan := chart.HourPillar.TianGan

	// 甲己合土、乙庚合金、丙辛合水、丁壬合木、戊癸合火
	heHeMap := map[string]string{
		"甲": "己", "己": "甲",
		"乙": "庚", "庚": "乙",
		"丙": "辛", "辛": "丙",
		"丁": "壬", "壬": "丁",
		"戊": "癸", "癸": "戊",
	}

	if heGan, ok := heHeMap[dayMaster.Name]; ok {
		if monthGan.Name == heGan || hourGan.Name == heGan {
			return true
		}
	}

	return false
}

// determineHuaShen 确定化神五行
func determineHuaShen(chart *base.BaZiChart) base.WuXing {
	dayMaster := chart.GetDayMaster()
	heHuaMap := map[string]base.WuXing{
		"甲": base.Earth, "己": base.Earth,
		"乙": base.Metal, "庚": base.Metal,
		"丙": base.Water, "辛": base.Water,
		"丁": base.Wood, "壬": base.Wood,
		"戊": base.Fire, "癸": base.Fire,
	}

	if huaShen, ok := heHuaMap[dayMaster.Name]; ok {
		return huaShen
	}

	return base.Earth
}

// getKeZhiHuaShen 获取克制化神的五行
func getKeZhiHuaShen(huaShen base.WuXing) base.WuXing {
	for k, v := range base.WuXingKe {
		if v == huaShen {
			return k
		}
	}
	return ""
}

// isZhuanWangGe 判断是否专旺格
func isZhuanWangGe(chart *base.BaZiChart, step2Result *step2_rebuild.Step2RebuildResult) bool {
	wuxingCount := map[base.WuXing]int{
		base.Wood:  0,
		base.Fire:  0,
		base.Earth: 0,
		base.Metal: 0,
		base.Water: 0,
	}

	for _, tg := range chart.GetAllTianGan() {
		if step2Result.IsTianGanAbsorbed(tg) {
			continue
		}
		wuxingCount[tg.Element]++
	}

	for _, dz := range chart.GetAllDiZhi() {
		if step2Result.IsDiZhiAbsorbed(dz) {
			continue
		}
		dzWuXing := step2_rebuild.GetRebuiltDiZhiWuXing(step2Result, dz)
		wuxingCount[dzWuXing]++
	}

	// 某一五行占据4个或以上
	for _, count := range wuxingCount {
		if count >= 4 {
			return true
		}
	}

	return false
}

// getZhuWangWuXing 获取专旺的五行
func getZhuWangWuXing(chart *base.BaZiChart, step2Result *step2_rebuild.Step2RebuildResult) base.WuXing {
	wuxingCount := map[base.WuXing]int{
		base.Wood:  0,
		base.Fire:  0,
		base.Earth: 0,
		base.Metal: 0,
		base.Water: 0,
	}

	for _, tg := range chart.GetAllTianGan() {
		if step2Result.IsTianGanAbsorbed(tg) {
			continue
		}
		wuxingCount[tg.Element]++
	}

	for _, dz := range chart.GetAllDiZhi() {
		if step2Result.IsDiZhiAbsorbed(dz) {
			continue
		}
		dzWuXing := step2_rebuild.GetRebuiltDiZhiWuXing(step2Result, dz)
		wuxingCount[dzWuXing]++
	}

	maxCount := 0
	var zhuWuXing base.WuXing
	for wx, count := range wuxingCount {
		if count > maxCount {
			maxCount = count
			zhuWuXing = wx
		}
	}

	return zhuWuXing
}

// 辅助函数

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

func uniqueWuXing(arr []base.WuXing) []base.WuXing {
	seen := make(map[base.WuXing]bool)
	result := []base.WuXing{}
	for _, v := range arr {
		if v != "" && !seen[v] {
			seen[v] = true
			result = append(result, v)
		}
	}
	return result
}

func uniqueStrings(arr []string) []string {
	seen := make(map[string]bool)
	result := []string{}
	for _, v := range arr {
		if v != "" && !seen[v] {
			seen[v] = true
			result = append(result, v)
		}
	}
	return result
}

func deduceShiShenFromWuXing(dayMasterWuXing base.WuXing, wuXingList []base.WuXing) []string {
	shiShenSet := make(map[string]bool)

	for _, wx := range wuXingList {
		if wx == dayMasterWuXing {
			shiShenSet["比劫"] = true
		} else if base.Sheng(wx, dayMasterWuXing) {
			shiShenSet["印星"] = true
		} else if base.Sheng(dayMasterWuXing, wx) {
			shiShenSet["食伤"] = true
		} else if base.Ke(dayMasterWuXing, wx) {
			shiShenSet["官杀"] = true
		} else if base.Ke(wx, dayMasterWuXing) {
			shiShenSet["财星"] = true
		}
	}

	result := []string{}
	for ss := range shiShenSet {
		result = append(result, ss)
	}
	sort.Strings(result)
	return result
}

func wuxingSliceToString(arr []base.WuXing) string {
	if len(arr) == 0 {
		return "无"
	}
	result := ""
	for i, wx := range arr {
		if i > 0 {
			result += "、"
		}
		result += string(wx)
	}
	return result
}
