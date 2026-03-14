package step4_xiji

import (
	"bazi-analyzer/internal/engine/step2_rebuild"
	"bazi-analyzer/internal/engine/step3_wangshuai"
	"bazi-analyzer/internal/model/base"
)

type BingYuanType string

const (
	BingYuanTiaoHou  BingYuanType = "调候"
	BingYuanTongGuan BingYuanType = "通关"
	BingYuanFuYi     BingYuanType = "扶抑"
	BingYuanBingYao  BingYuanType = "病药"
	BingYuanGeJu     BingYuanType = "格局"
)

type BingYuan struct {
	Type        BingYuanType
	Severity    int
	Description string
	WuXing      base.WuXing
	ShiShen     string
	Needs       []base.WuXing
	Avoids      []base.WuXing
}

type BingYuanDiagnosis struct {
	AllBingYuan   []BingYuan
	CoreBingYuan  *BingYuan
	SubBingYuan   []BingYuan
	DiagnosisPath []string
}

func DiagnoseBingYuan(chart *base.BaZiChart, step2Result *step2_rebuild.Step2RebuildResult, step3Result *step3_wangshuai.Step3WangShuaiResult) *BingYuanDiagnosis {
	diagnosis := &BingYuanDiagnosis{
		AllBingYuan:   make([]BingYuan, 0),
		SubBingYuan:   make([]BingYuan, 0),
		DiagnosisPath: make([]string, 0),
	}

	diagnosis.DiagnosisPath = append(diagnosis.DiagnosisPath, "开始病源诊断...")

	diagnoseTiaoHou(chart, step2Result, step3Result, diagnosis)
	diagnoseTongGuan(chart, step2Result, step3Result, diagnosis)
	diagnoseFuYi(chart, step2Result, step3Result, diagnosis)
	diagnoseBingYao(chart, step2Result, step3Result, diagnosis)
	diagnoseGeJu(chart, step2Result, step3Result, diagnosis)

	determineCoreBingYuan(diagnosis)

	return diagnosis
}

func diagnoseTiaoHou(chart *base.BaZiChart, step2Result *step2_rebuild.Step2RebuildResult, step3Result *step3_wangshuai.Step3WangShuaiResult, diagnosis *BingYuanDiagnosis) {
	monthLing := chart.GetMonthLing()
	monthWuXing := step2_rebuild.GetRebuiltDiZhiWuXing(step2Result, monthLing)

	isCold := false
	isHot := false
	coldWuXing := []base.WuXing{base.Water, base.Metal}
	hotWuXing := []base.WuXing{base.Fire, base.Wood}

	for _, cw := range coldWuXing {
		if monthWuXing == cw {
			isCold = true
			break
		}
	}

	for _, hw := range hotWuXing {
		if monthWuXing == hw {
			isHot = true
			break
		}
	}

	fireCount := countWuXing(chart, step2Result, base.Fire)
	waterCount := countWuXing(chart, step2Result, base.Water)

	if isCold && fireCount == 0 {
		by := BingYuan{
			Type:        BingYuanTiaoHou,
			Severity:    3,
			Description: "命局过寒，生于冬令或金水旺月，缺火调候，寒凝不化",
			WuXing:      base.Water,
			Needs:       []base.WuXing{base.Fire},
			Avoids:      []base.WuXing{base.Water},
		}
		diagnosis.AllBingYuan = append(diagnosis.AllBingYuan, by)
		diagnosis.DiagnosisPath = append(diagnosis.DiagnosisPath, "调候诊断：命局过寒，需火调候")
	} else if isHot && waterCount == 0 {
		by := BingYuan{
			Type:        BingYuanTiaoHou,
			Severity:    3,
			Description: "命局过热，生于夏令或木火旺月，缺水调候，燥热难安",
			WuXing:      base.Fire,
			Needs:       []base.WuXing{base.Water},
			Avoids:      []base.WuXing{base.Fire},
		}
		diagnosis.AllBingYuan = append(diagnosis.AllBingYuan, by)
		diagnosis.DiagnosisPath = append(diagnosis.DiagnosisPath, "调候诊断：命局过热，需水调候")
	} else {
		diagnosis.DiagnosisPath = append(diagnosis.DiagnosisPath, "调候诊断：寒暖适中，无调候之急")
	}
}

func diagnoseTongGuan(chart *base.BaZiChart, step2Result *step2_rebuild.Step2RebuildResult, step3Result *step3_wangshuai.Step3WangShuaiResult, diagnosis *BingYuanDiagnosis) {
	dayMasterWuXing := step3Result.DayMasterWuXing

	wuXingCounts := make(map[base.WuXing]int)
	for _, wx := range []base.WuXing{base.Wood, base.Fire, base.Earth, base.Metal, base.Water} {
		wuXingCounts[wx] = countWuXing(chart, step2Result, wx)
	}

	keWo := getWuXingKeWo(dayMasterWuXing)
	woSheng := base.WuXingSheng[dayMasterWuXing]

	if wuXingCounts[keWo] > 0 && wuXingCounts[dayMasterWuXing] > 0 {
		tongGuanWuXing := base.WuXingSheng[keWo]
		if wuXingCounts[tongGuanWuXing] == 0 {
			by := BingYuan{
				Type:        BingYuanTongGuan,
				Severity:    2,
				Description: "官杀与日主相战，缺通关之神，两虎相斗必有一伤",
				WuXing:      tongGuanWuXing,
				ShiShen:     "食伤",
				Needs:       []base.WuXing{tongGuanWuXing},
				Avoids:      []base.WuXing{keWo},
			}
			diagnosis.AllBingYuan = append(diagnosis.AllBingYuan, by)
			diagnosis.DiagnosisPath = append(diagnosis.DiagnosisPath, "通关诊断：官杀克身，需食伤通关")
		}
	}

	woKe := base.WuXingKe[dayMasterWuXing]
	if wuXingCounts[woKe] > 0 && wuXingCounts[dayMasterWuXing] > 0 {
		if wuXingCounts[woSheng] == 0 {
			by := BingYuan{
				Type:        BingYuanTongGuan,
				Severity:    2,
				Description: "日主与财星相战，缺通关之神，比劫争财难两全",
				WuXing:      woSheng,
				ShiShen:     "食伤",
				Needs:       []base.WuXing{woSheng},
				Avoids:      []base.WuXing{},
			}
			diagnosis.AllBingYuan = append(diagnosis.AllBingYuan, by)
			diagnosis.DiagnosisPath = append(diagnosis.DiagnosisPath, "通关诊断：比劫争财，需食伤通关")
		}
	}

	if len(diagnosis.AllBingYuan) == 0 || diagnosis.AllBingYuan[len(diagnosis.AllBingYuan)-1].Type != BingYuanTongGuan {
		diagnosis.DiagnosisPath = append(diagnosis.DiagnosisPath, "通关诊断：五行流通，无通关之急")
	}
}

func diagnoseFuYi(chart *base.BaZiChart, step2Result *step2_rebuild.Step2RebuildResult, step3Result *step3_wangshuai.Step3WangShuaiResult, diagnosis *BingYuanDiagnosis) {
	dayMasterWuXing := step3Result.DayMasterWuXing
	wangShuaiType := step3Result.WangShuaiType

	shengWo := getWuXingShengWo(dayMasterWuXing)
	keWo := getWuXingKeWo(dayMasterWuXing)
	woSheng := base.WuXingSheng[dayMasterWuXing]
	woKe := base.WuXingKe[dayMasterWuXing]

	if wangShuaiType == step3_wangshuai.ShenRuo {
		severity := 2
		if step3Result.ShengFuCount <= 1 {
			severity = 3
		}
		by := BingYuan{
			Type:        BingYuanFuYi,
			Severity:    severity,
			Description: "日主身弱，元气不足，需生扶以壮其势",
			WuXing:      dayMasterWuXing,
			Needs:       []base.WuXing{shengWo, dayMasterWuXing},
			Avoids:      []base.WuXing{keWo, woSheng, woKe},
		}
		diagnosis.AllBingYuan = append(diagnosis.AllBingYuan, by)
		diagnosis.DiagnosisPath = append(diagnosis.DiagnosisPath, "扶抑诊断：日主身弱，需印比生扶")
	} else if wangShuaiType == step3_wangshuai.ShenWang {
		severity := 2
		if step3Result.ShengFuCount >= 4 {
			severity = 3
		}
		by := BingYuan{
			Type:        BingYuanFuYi,
			Severity:    severity,
			Description: "日主身旺，气势过盛，需克泄耗以平其势",
			WuXing:      dayMasterWuXing,
			Needs:       []base.WuXing{keWo, woSheng, woKe},
			Avoids:      []base.WuXing{shengWo, dayMasterWuXing},
		}
		diagnosis.AllBingYuan = append(diagnosis.AllBingYuan, by)
		diagnosis.DiagnosisPath = append(diagnosis.DiagnosisPath, "扶抑诊断：日主身旺，需克泄耗平衡")
	} else {
		diagnosis.DiagnosisPath = append(diagnosis.DiagnosisPath, "扶抑诊断：日主中和，无明显扶抑之病")
	}
}

func diagnoseBingYao(chart *base.BaZiChart, step2Result *step2_rebuild.Step2RebuildResult, step3Result *step3_wangshuai.Step3WangShuaiResult, diagnosis *BingYuanDiagnosis) {
	dayMasterWuXing := step3Result.DayMasterWuXing

	wuXingCounts := make(map[base.WuXing]int)
	for _, wx := range []base.WuXing{base.Wood, base.Fire, base.Earth, base.Metal, base.Water} {
		wuXingCounts[wx] = countWuXing(chart, step2Result, wx)
	}

	keWo := getWuXingKeWo(dayMasterWuXing)
	if wuXingCounts[keWo] >= 3 {
		shiShen := base.GetShiShen(dayMasterWuXing, keWo, true)
		by := BingYuan{
			Type:        BingYuanBingYao,
			Severity:    4,
			Description: "官杀重重克身，为命局最重之病，急需食伤制杀或印星化杀",
			WuXing:      keWo,
			ShiShen:     string(shiShen),
			Needs:       []base.WuXing{base.WuXingSheng[keWo], getWuXingShengWo(dayMasterWuXing)},
			Avoids:      []base.WuXing{keWo},
		}
		diagnosis.AllBingYuan = append(diagnosis.AllBingYuan, by)
		diagnosis.DiagnosisPath = append(diagnosis.DiagnosisPath, "病药诊断：官杀重克，为命局大病")
	}

	woKe := base.WuXingKe[dayMasterWuXing]
	if wuXingCounts[woKe] >= 3 {
		shiShen := base.GetShiShen(dayMasterWuXing, woKe, true)
		by := BingYuan{
			Type:        BingYuanBingYao,
			Severity:    3,
			Description: "财星过旺，日主难以胜任，需比劫帮身分财",
			WuXing:      woKe,
			ShiShen:     string(shiShen),
			Needs:       []base.WuXing{dayMasterWuXing},
			Avoids:      []base.WuXing{woKe},
		}
		diagnosis.AllBingYuan = append(diagnosis.AllBingYuan, by)
		diagnosis.DiagnosisPath = append(diagnosis.DiagnosisPath, "病药诊断：财多身弱，需比劫分财")
	}

	shengWo := getWuXingShengWo(dayMasterWuXing)
	if wuXingCounts[shengWo] >= 3 && step3Result.WangShuaiType == step3_wangshuai.ShenWang {
		by := BingYuan{
			Type:        BingYuanBingYao,
			Severity:    3,
			Description: "印星过旺，母多灭子，需财星破印",
			WuXing:      shengWo,
			ShiShen:     "印星",
			Needs:       []base.WuXing{woKe},
			Avoids:      []base.WuXing{shengWo},
		}
		diagnosis.AllBingYuan = append(diagnosis.AllBingYuan, by)
		diagnosis.DiagnosisPath = append(diagnosis.DiagnosisPath, "病药诊断：印旺为病，需财星破印")
	}

	if len(diagnosis.AllBingYuan) == 0 || diagnosis.AllBingYuan[len(diagnosis.AllBingYuan)-1].Type != BingYuanBingYao {
		diagnosis.DiagnosisPath = append(diagnosis.DiagnosisPath, "病药诊断：无明显忌神肆虐之病")
	}
}

func diagnoseGeJu(chart *base.BaZiChart, step2Result *step2_rebuild.Step2RebuildResult, step3Result *step3_wangshuai.Step3WangShuaiResult, diagnosis *BingYuanDiagnosis) {
	wangShuaiType := step3Result.WangShuaiType
	dayMasterWuXing := step3Result.DayMasterWuXing

	if wangShuaiType == step3_wangshuai.CongQiang {
		keWo := getWuXingKeWo(dayMasterWuXing)
		woSheng := base.WuXingSheng[dayMasterWuXing]
		woKe := base.WuXingKe[dayMasterWuXing]
		shengWo := getWuXingShengWo(dayMasterWuXing)

		by := BingYuan{
			Type:        BingYuanGeJu,
			Severity:    5,
			Description: "从强格已成，顺势为用，逆势为忌，最忌克泄耗破格",
			WuXing:      dayMasterWuXing,
			Needs:       []base.WuXing{dayMasterWuXing, shengWo},
			Avoids:      []base.WuXing{keWo, woSheng, woKe},
		}
		diagnosis.AllBingYuan = append(diagnosis.AllBingYuan, by)
		diagnosis.DiagnosisPath = append(diagnosis.DiagnosisPath, "格局诊断：从强格，顺其旺势为用")
	} else if wangShuaiType == step3_wangshuai.CongRuo {
		keWo := getWuXingKeWo(dayMasterWuXing)
		woSheng := base.WuXingSheng[dayMasterWuXing]
		woKe := base.WuXingKe[dayMasterWuXing]
		shengWo := getWuXingShengWo(dayMasterWuXing)

		by := BingYuan{
			Type:        BingYuanGeJu,
			Severity:    5,
			Description: "从弱格已成，顺势为用，逆势为忌，最忌生扶破格",
			WuXing:      dayMasterWuXing,
			Needs:       []base.WuXing{keWo, woSheng, woKe},
			Avoids:      []base.WuXing{dayMasterWuXing, shengWo},
		}
		diagnosis.AllBingYuan = append(diagnosis.AllBingYuan, by)
		diagnosis.DiagnosisPath = append(diagnosis.DiagnosisPath, "格局诊断：从弱格，顺其弱势为用")
	} else {
		diagnosis.DiagnosisPath = append(diagnosis.DiagnosisPath, "格局诊断：正格命局，按常法推断")
	}
}

func determineCoreBingYuan(diagnosis *BingYuanDiagnosis) {
	if len(diagnosis.AllBingYuan) == 0 {
		return
	}

	maxSeverity := 0
	var coreIdx int

	for i, by := range diagnosis.AllBingYuan {
		if by.Severity > maxSeverity {
			maxSeverity = by.Severity
			coreIdx = i
		}
	}

	diagnosis.CoreBingYuan = &diagnosis.AllBingYuan[coreIdx]

	for i, by := range diagnosis.AllBingYuan {
		if i != coreIdx {
			diagnosis.SubBingYuan = append(diagnosis.SubBingYuan, by)
		}
	}

	diagnosis.DiagnosisPath = append(diagnosis.DiagnosisPath, "确定核心病源："+diagnosis.CoreBingYuan.Description)
}

func countWuXing(chart *base.BaZiChart, step2Result *step2_rebuild.Step2RebuildResult, targetWuXing base.WuXing) int {
	count := 0

	for _, tg := range chart.GetAllTianGan() {
		if step2Result.IsTianGanAbsorbed(tg) {
			continue
		}
		wx := step2_rebuild.GetRebuiltTianGanWuXing(step2Result, tg)
		if wx == targetWuXing {
			count++
		}
	}

	for _, dz := range chart.GetAllDiZhi() {
		if step2Result.IsDiZhiAbsorbed(dz) {
			continue
		}
		wx := step2_rebuild.GetRebuiltDiZhiWuXing(step2Result, dz)
		if wx == targetWuXing {
			count++
		}
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
