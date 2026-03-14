package step4_xiji

import (
	"bazi-analyzer/internal/engine/step2_rebuild"
	"bazi-analyzer/internal/engine/step3_wangshuai"
	"bazi-analyzer/internal/model/base"
)

type YongShenType string

const (
	YongShenTiaoHou  YongShenType = "调候用神"
	YongShenTongGuan YongShenType = "通关用神"
	YongShenFuYi     YongShenType = "扶抑用神"
	YongShenBingYao  YongShenType = "病药用神"
	YongShenGeJu     YongShenType = "格局用神"
)

type YongShenInfo struct {
	WuXing      base.WuXing
	ShiShen     string
	Type        YongShenType
	Strength    int
	Description string
	Source      string
}

type JiShenInfo struct {
	WuXing      base.WuXing
	ShiShen     string
	Reason      string
	Description string
}

type XiShenInfo struct {
	WuXing      base.WuXing
	ShiShen     string
	Reason      string
	Description string
}

type YongShenResult struct {
	YongShen   *YongShenInfo
	XiShen     []XiShenInfo
	JiShen     []JiShenInfo
	Derivation []string
}

func DeriveYongShen(chart *base.BaZiChart, step2Result *step2_rebuild.Step2RebuildResult, step3Result *step3_wangshuai.Step3WangShuaiResult, diagnosis *BingYuanDiagnosis) *YongShenResult {
	result := &YongShenResult{
		XiShen:     make([]XiShenInfo, 0),
		JiShen:     make([]JiShenInfo, 0),
		Derivation: make([]string, 0),
	}

	result.Derivation = append(result.Derivation, "开始用神推导...")

	if diagnosis.CoreBingYuan == nil {
		deriveDefaultYongShen(chart, step2Result, step3Result, result)
		return result
	}

	coreBing := diagnosis.CoreBingYuan
	result.Derivation = append(result.Derivation, "核心病源："+coreBing.Description)

	deriveYongShenFromBingYuan(chart, step2Result, step3Result, coreBing, result)

	deriveXiShen(chart, step2Result, step3Result, result)

	deriveJiShen(chart, step2Result, step3Result, result)

	validateConsistency(result)

	return result
}

func deriveYongShenFromBingYuan(chart *base.BaZiChart, step2Result *step2_rebuild.Step2RebuildResult, step3Result *step3_wangshuai.Step3WangShuaiResult, bingYuan *BingYuan, result *YongShenResult) {
	dayMasterWuXing := step3Result.DayMasterWuXing
	dayMaster := chart.GetDayMaster()

	if len(bingYuan.Needs) == 0 {
		result.Derivation = append(result.Derivation, "病源无明确需求，按扶抑法推导")
		deriveDefaultYongShen(chart, step2Result, step3Result, result)
		return
	}

	yongShenWuXing := bingYuan.Needs[0]
	yongShenShiShen := base.GetShiShen(dayMasterWuXing, yongShenWuXing, dayMaster.IsYang)

	yongShenType := mapBingYuanTypeToYongShenType(bingYuan.Type)

	result.YongShen = &YongShenInfo{
		WuXing:      yongShenWuXing,
		ShiShen:     string(yongShenShiShen),
		Type:        yongShenType,
		Strength:    bingYuan.Severity,
		Description: generateYongShenDescription(yongShenType, yongShenWuXing, bingYuan),
		Source:      bingYuan.Description,
	}

	result.Derivation = append(result.Derivation, "确定用神："+yongShenWuXing.String()+"("+string(yongShenShiShen)+")")
	result.Derivation = append(result.Derivation, "用神类型："+string(yongShenType))
}

func deriveXiShen(chart *base.BaZiChart, step2Result *step2_rebuild.Step2RebuildResult, step3Result *step3_wangshuai.Step3WangShuaiResult, result *YongShenResult) {
	if result.YongShen == nil {
		return
	}

	dayMasterWuXing := step3Result.DayMasterWuXing
	dayMaster := chart.GetDayMaster()
	yongShenWuXing := result.YongShen.WuXing

	shengYongShen := getWuXingShengWo(yongShenWuXing)
	if shengYongShen != "" && shengYongShen != dayMasterWuXing {
		xiShenShiShen := base.GetShiShen(dayMasterWuXing, shengYongShen, dayMaster.IsYang)
		result.XiShen = append(result.XiShen, XiShenInfo{
			WuXing:      shengYongShen,
			ShiShen:     string(xiShenShiShen),
			Reason:      "生助用神",
			Description: shengYongShen.String() + "生" + yongShenWuXing.String() + "，助用神发力",
		})
		result.Derivation = append(result.Derivation, "喜神："+shengYongShen.String()+"(生助用神)")
	}

	yongShenSheng := base.WuXingSheng[yongShenWuXing]
	if yongShenSheng != "" && yongShenSheng != dayMasterWuXing {
		xiShenShiShen := base.GetShiShen(dayMasterWuXing, yongShenSheng, dayMaster.IsYang)
		exists := false
		for _, xs := range result.XiShen {
			if xs.WuXing == yongShenSheng {
				exists = true
				break
			}
		}
		if !exists {
			result.XiShen = append(result.XiShen, XiShenInfo{
				WuXing:      yongShenSheng,
				ShiShen:     string(xiShenShiShen),
				Reason:      "用神所生",
				Description: yongShenWuXing.String() + "生" + yongShenSheng.String() + "，顺势而为",
			})
			result.Derivation = append(result.Derivation, "喜神："+yongShenSheng.String()+"(用神所生)")
		}
	}

	if result.YongShen.Type == YongShenTongGuan {
		tongGuanXiShen := yongShenWuXing
		xiShenShiShen := base.GetShiShen(dayMasterWuXing, tongGuanXiShen, dayMaster.IsYang)
		exists := false
		for _, xs := range result.XiShen {
			if xs.WuXing == tongGuanXiShen {
				exists = true
				break
			}
		}
		if !exists {
			result.XiShen = append(result.XiShen, XiShenInfo{
				WuXing:      tongGuanXiShen,
				ShiShen:     string(xiShenShiShen),
				Reason:      "通关之神",
				Description: "食伤通关，化解官杀克身之危",
			})
		}
	}
}

func deriveJiShen(chart *base.BaZiChart, step2Result *step2_rebuild.Step2RebuildResult, step3Result *step3_wangshuai.Step3WangShuaiResult, result *YongShenResult) {
	if result.YongShen == nil {
		return
	}

	dayMasterWuXing := step3Result.DayMasterWuXing
	dayMaster := chart.GetDayMaster()
	yongShenWuXing := result.YongShen.WuXing

	keYongShen := getWuXingKeWo(yongShenWuXing)
	if keYongShen != "" {
		jiShenShiShen := base.GetShiShen(dayMasterWuXing, keYongShen, dayMaster.IsYang)
		result.JiShen = append(result.JiShen, JiShenInfo{
			WuXing:      keYongShen,
			ShiShen:     string(jiShenShiShen),
			Reason:      "克制用神",
			Description: keYongShen.String() + "克" + yongShenWuXing.String() + "，损伤用神之力",
		})
		result.Derivation = append(result.Derivation, "忌神："+keYongShen.String()+"(克制用神)")
	}

	yongShenKe := base.WuXingKe[yongShenWuXing]
	if yongShenKe != "" {
		jiShenShiShen := base.GetShiShen(dayMasterWuXing, yongShenKe, dayMaster.IsYang)
		exists := false
		for _, js := range result.JiShen {
			if js.WuXing == yongShenKe {
				exists = true
				break
			}
		}
		if !exists {
			result.JiShen = append(result.JiShen, JiShenInfo{
				WuXing:      yongShenKe,
				ShiShen:     string(jiShenShiShen),
				Reason:      "用神所克",
				Description: yongShenWuXing.String() + "克" + yongShenKe.String() + "，消耗用神之力",
			})
			result.Derivation = append(result.Derivation, "忌神："+yongShenKe.String()+"(用神所克)")
		}
	}

	if result.YongShen.Type == YongShenFuYi || result.YongShen.Type == YongShenGeJu {
		if step3Result.WangShuaiType == step3_wangshuai.ShenRuo || step3Result.WangShuaiType == step3_wangshuai.CongRuo {
			keWo := getWuXingKeWo(dayMasterWuXing)
			jiShenShiShen := base.GetShiShen(dayMasterWuXing, keWo, dayMaster.IsYang)
			exists := false
			for _, js := range result.JiShen {
				if js.WuXing == keWo {
					exists = true
					break
				}
			}
			if !exists {
				result.JiShen = append(result.JiShen, JiShenInfo{
					WuXing:      keWo,
					ShiShen:     string(jiShenShiShen),
					Reason:      "克身之物",
					Description: "身弱忌官杀克身",
				})
			}
		}
	}
}

func deriveDefaultYongShen(chart *base.BaZiChart, step2Result *step2_rebuild.Step2RebuildResult, step3Result *step3_wangshuai.Step3WangShuaiResult, result *YongShenResult) {
	dayMasterWuXing := step3Result.DayMasterWuXing
	dayMaster := chart.GetDayMaster()

	var yongShenWuXing base.WuXing
	var yongShenShiShen base.ShiShen
	var description string

	switch step3Result.WangShuaiType {
	case step3_wangshuai.ShenRuo:
		shengWo := getWuXingShengWo(dayMasterWuXing)
		yongShenWuXing = shengWo
		yongShenShiShen = base.GetShiShen(dayMasterWuXing, shengWo, dayMaster.IsYang)
		description = "日主身弱，以印星为用，生扶日主"
	case step3_wangshuai.ShenWang:
		woSheng := base.WuXingSheng[dayMasterWuXing]
		yongShenWuXing = woSheng
		yongShenShiShen = base.GetShiShen(dayMasterWuXing, woSheng, dayMaster.IsYang)
		description = "日主身旺，以食伤为用，泄秀平衡"
	case step3_wangshuai.CongQiang:
		yongShenWuXing = dayMasterWuXing
		yongShenShiShen = base.BiJian
		description = "从强格，以比劫为用，顺其旺势"
	case step3_wangshuai.CongRuo:
		keWo := getWuXingKeWo(dayMasterWuXing)
		yongShenWuXing = keWo
		yongShenShiShen = base.GetShiShen(dayMasterWuXing, keWo, dayMaster.IsYang)
		description = "从弱格，以官杀为用，顺其弱势"
	}

	result.YongShen = &YongShenInfo{
		WuXing:      yongShenWuXing,
		ShiShen:     string(yongShenShiShen),
		Type:        YongShenFuYi,
		Strength:    2,
		Description: description,
		Source:      "扶抑法推导",
	}

	result.Derivation = append(result.Derivation, "默认用神："+yongShenWuXing.String()+"("+string(yongShenShiShen)+")")
}

func validateConsistency(result *YongShenResult) {
	if result.YongShen == nil {
		return
	}

	yongShenWuXing := result.YongShen.WuXing

	validJiShen := make([]JiShenInfo, 0)
	for _, js := range result.JiShen {
		if js.WuXing != yongShenWuXing {
			validJiShen = append(validJiShen, js)
		}
	}
	result.JiShen = validJiShen

	validXiShen := make([]XiShenInfo, 0)
	for _, xs := range result.XiShen {
		if xs.WuXing != yongShenWuXing {
			isJiShen := false
			for _, js := range result.JiShen {
				if js.WuXing == xs.WuXing {
					isJiShen = true
					break
				}
			}
			if !isJiShen {
				validXiShen = append(validXiShen, xs)
			}
		}
	}
	result.XiShen = validXiShen

	result.Derivation = append(result.Derivation, "一致性验证完成")
}

func mapBingYuanTypeToYongShenType(byType BingYuanType) YongShenType {
	switch byType {
	case BingYuanTiaoHou:
		return YongShenTiaoHou
	case BingYuanTongGuan:
		return YongShenTongGuan
	case BingYuanFuYi:
		return YongShenFuYi
	case BingYuanBingYao:
		return YongShenBingYao
	case BingYuanGeJu:
		return YongShenGeJu
	default:
		return YongShenFuYi
	}
}

func generateYongShenDescription(ysType YongShenType, wx base.WuXing, bingYuan *BingYuan) string {
	switch ysType {
	case YongShenTiaoHou:
		return "命局" + bingYuan.Description + "，以" + wx.String() + "调候为用"
	case YongShenTongGuan:
		return "命局五行相战，以" + wx.String() + "通关为用，化解矛盾"
	case YongShenFuYi:
		return "日主失衡，以" + wx.String() + "扶抑为用，平衡命局"
	case YongShenBingYao:
		return "命局有病，以" + wx.String() + "为药，对症下药"
	case YongShenGeJu:
		return "特殊格局，以" + wx.String() + "顺势为用"
	default:
		return "以" + wx.String() + "为用神"
	}
}
