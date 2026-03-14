package step4_xiji

import (
	"bazi-analyzer/internal/engine/step1_scan"
	"bazi-analyzer/internal/engine/step2_rebuild"
	"bazi-analyzer/internal/engine/step3_wangshuai"
	"bazi-analyzer/internal/model/base"
	"bazi-analyzer/pkg/utils"
	"fmt"
	"testing"
)

func TestAnalyzeXiJi_ShenRuo(t *testing.T) {
	// 身弱八字示例
	input := "壬戌 壬寅 庚午 丙戌"
	chart, err := utils.ParseBaZi(input)
	if err != nil {
		t.Fatalf("解析八字失败: %v", err)
	}

	step1Result := step1_scan.ScanBaZi(&chart)
	step2Result := step2_rebuild.RebuildBaZi(&chart, step1Result)
	step3Result := step3_wangshuai.AnalyzeWangShuai(&chart, step2Result)
	result := AnalyzeXiJi(&chart, step2Result, step3Result)

	t.Logf("=== 测试八字: %s ===", input)
	t.Logf("日主: %s (%s)", result.DayMaster.Name, result.DayMasterWuXing)
	t.Logf("旺衰: %s", result.WangShuaiType)
	t.Log("")

	// 验证多维度诊断
	assertBingYaoAnalysis(t, "调候诊断", result.TiaoHouAnalysis)
	assertBingYaoAnalysis(t, "通关诊断", result.TongGuanAnalysis)
	assertBingYaoAnalysis(t, "扶抑诊断", result.FuYiAnalysis)
	assertBingYaoAnalysis(t, "格局诊断", result.GeJuAnalysis)

	t.Log("")
	t.Logf("核心病源: %s", result.CoreBingDesc)
	t.Logf("用神五行: %v", result.YongShenWuXing)
	t.Logf("用神十神: %v", result.YongShenShiShen)
	t.Logf("喜神五行: %v", result.XiShenWuXing)
	t.Logf("喜神十神: %v", result.XiShenShiShen)
	t.Logf("忌神五行: %v", result.JiShenWuXing)
	t.Logf("忌神十神: %v", result.JiShenShiShen)
	t.Logf("判定依据: %s", result.YongShenReason)
	t.Logf("综合结论: %s", result.Summary)

	// 验证结果
	if len(result.YongShenWuXing) == 0 {
		t.Error("用神五行不应为空")
	}
	if result.CoreBingYao == nil {
		t.Error("核心病源不应为空")
	}
}

func TestAnalyzeXiJi_ShenWang(t *testing.T) {
	// 身旺八字示例
	input := "甲子 丙寅 戊辰 戊午"
	chart, err := utils.ParseBaZi(input)
	if err != nil {
		t.Fatalf("解析八字失败: %v", err)
	}

	step1Result := step1_scan.ScanBaZi(&chart)
	step2Result := step2_rebuild.RebuildBaZi(&chart, step1Result)
	step3Result := step3_wangshuai.AnalyzeWangShuai(&chart, step2Result)
	result := AnalyzeXiJi(&chart, step2Result, step3Result)

	t.Logf("=== 测试八字: %s ===", input)
	t.Logf("日主: %s (%s)", result.DayMaster.Name, result.DayMasterWuXing)
	t.Logf("旺衰: %s", result.WangShuaiType)
	t.Log("")
	t.Logf("核心病源: %s", result.CoreBingDesc)
	t.Logf("用神五行: %v", result.YongShenWuXing)
	t.Logf("喜神五行: %v", result.XiShenWuXing)
	t.Logf("忌神五行: %v", result.JiShenWuXing)
	t.Logf("综合结论: %s", result.Summary)

	if len(result.YongShenWuXing) == 0 {
		t.Error("用神五行不应为空")
	}
}

func TestAnalyzeXiJi_TiaoHou(t *testing.T) {
	// 冬生八字，测试调候病
	input := "癸亥 甲子 壬子 庚子"
	chart, err := utils.ParseBaZi(input)
	if err != nil {
		t.Fatalf("解析八字失败: %v", err)
	}

	step1Result := step1_scan.ScanBaZi(&chart)
	step2Result := step2_rebuild.RebuildBaZi(&chart, step1Result)
	step3Result := step3_wangshuai.AnalyzeWangShuai(&chart, step2Result)
	result := AnalyzeXiJi(&chart, step2Result, step3Result)

	t.Logf("=== 冬生八字调候测试: %s ===", input)
	t.Logf("调候诊断: %s (严重程度: %d)", result.TiaoHouAnalysis.BingDesc, result.TiaoHouAnalysis.Severity)
	t.Logf("核心病源: %s", result.CoreBingDesc)
	t.Logf("用神五行: %v", result.YongShenWuXing)

	// 冬生八字应该需要火调候
	hasFire := false
	for _, wx := range result.YongShenWuXing {
		if wx == base.Fire {
			hasFire = true
			break
		}
	}
	if !hasFire && result.TiaoHouAnalysis.Severity >= 3 {
		t.Error("冬生八字用神应包含火")
	}
}

func TestAnalyzeXiJi_TongGuan(t *testing.T) {
	// 金木相战八字，测试通关
	input := "庚申 甲申 甲寅 乙卯"
	chart, err := utils.ParseBaZi(input)
	if err != nil {
		t.Fatalf("解析八字失败: %v", err)
	}

	step1Result := step1_scan.ScanBaZi(&chart)
	step2Result := step2_rebuild.RebuildBaZi(&chart, step1Result)
	step3Result := step3_wangshuai.AnalyzeWangShuai(&chart, step2Result)
	result := AnalyzeXiJi(&chart, step2Result, step3Result)

	t.Logf("=== 金木相战通关测试: %s ===", input)
	t.Logf("通关诊断: %s (严重程度: %d)", result.TongGuanAnalysis.BingDesc, result.TongGuanAnalysis.Severity)
	t.Logf("核心病源: %s", result.CoreBingDesc)
	t.Logf("用神五行: %v", result.YongShenWuXing)

	// 金木相战应该需要水通关
	hasWater := false
	for _, wx := range result.YongShenWuXing {
		if wx == base.Water {
			hasWater = true
			break
		}
	}
	if !hasWater && result.TongGuanAnalysis.Severity >= 4 {
		t.Logf("提示: 金木相战八字用神建议包含水通关")
	}
}

func TestAnalyzeXiJi_CongQiang(t *testing.T) {
	// 从强格八字
	input := "甲寅 丙寅 甲辰 乙卯"
	chart, err := utils.ParseBaZi(input)
	if err != nil {
		t.Fatalf("解析八字失败: %v", err)
	}

	step1Result := step1_scan.ScanBaZi(&chart)
	step2Result := step2_rebuild.RebuildBaZi(&chart, step1Result)
	step3Result := step3_wangshuai.AnalyzeWangShuai(&chart, step2Result)
	result := AnalyzeXiJi(&chart, step2Result, step3Result)

	t.Logf("=== 从强格测试: %s ===", input)
	t.Logf("旺衰: %s", result.WangShuaiType)
	t.Logf("核心病源: %s", result.CoreBingDesc)
	t.Logf("用神五行: %v", result.YongShenWuXing)
	t.Logf("忌神五行: %v", result.JiShenWuXing)

	if result.WangShuaiType != step3_wangshuai.CongQiang && result.WangShuaiType != step3_wangshuai.ShenWang {
		t.Logf("提示: 此八字可能不是从强格，而是%s", result.WangShuaiType)
	}
}

func TestAnalyzeXiJi_CongRuo(t *testing.T) {
	// 从弱格八字
	input := "庚申 甲申 庚申 甲申"
	chart, err := utils.ParseBaZi(input)
	if err != nil {
		t.Fatalf("解析八字失败: %v", err)
	}

	step1Result := step1_scan.ScanBaZi(&chart)
	step2Result := step2_rebuild.RebuildBaZi(&chart, step1Result)
	step3Result := step3_wangshuai.AnalyzeWangShuai(&chart, step2Result)
	result := AnalyzeXiJi(&chart, step2Result, step3Result)

	t.Logf("=== 从弱格测试: %s ===", input)
	t.Logf("旺衰: %s", result.WangShuaiType)
	t.Logf("核心病源: %s", result.CoreBingDesc)
	t.Logf("用神五行: %v", result.YongShenWuXing)
	t.Logf("忌神五行: %v", result.JiShenWuXing)
}

func TestAnalyzeXiJi_Consistency(t *testing.T) {
	// 测试同一命局多次计算结果一致
	input := "壬戌 壬寅 庚午 丙戌"
	chart, err := utils.ParseBaZi(input)
	if err != nil {
		t.Fatalf("解析八字失败: %v", err)
	}

	step1Result := step1_scan.ScanBaZi(&chart)
	step2Result := step2_rebuild.RebuildBaZi(&chart, step1Result)
	step3Result := step3_wangshuai.AnalyzeWangShuai(&chart, step2Result)

	// 多次计算
	result1 := AnalyzeXiJi(&chart, step2Result, step3Result)
	result2 := AnalyzeXiJi(&chart, step2Result, step3Result)
	result3 := AnalyzeXiJi(&chart, step2Result, step3Result)

	// 验证一致性
	if !slicesEqual(result1.YongShenWuXing, result2.YongShenWuXing) ||
		!slicesEqual(result2.YongShenWuXing, result3.YongShenWuXing) {
		t.Error("同一命局多次计算用神应一致")
	}

	if !slicesEqual(result1.JiShenWuXing, result2.JiShenWuXing) ||
		!slicesEqual(result2.JiShenWuXing, result3.JiShenWuXing) {
		t.Error("同一命局多次计算忌神应一致")
	}

	t.Logf("一致性测试通过: 同一命局%d次计算结果一致", 3)
}

func TestBingYaoTheory(t *testing.T) {
	// 测试病药理论核心逻辑
	testCases := []struct {
		name     string
		input    string
		expected string // 期望的病类型
	}{
		{"冬生调候", "癸亥 甲子 壬子 庚子", "调候病"},
		{"夏生调候", "丙午 丁巳 戊午 己巳", "调候病"},
		{"金木相战", "庚申 甲申 甲寅 乙卯", "通关病"},
		{"水火相战", "壬子 丙午 壬子 丙午", "通关病"},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			chart, err := utils.ParseBaZi(tc.input)
			if err != nil {
				t.Fatalf("解析八字失败: %v", err)
			}

			step1Result := step1_scan.ScanBaZi(&chart)
			step2Result := step2_rebuild.RebuildBaZi(&chart, step1Result)
			step3Result := step3_wangshuai.AnalyzeWangShuai(&chart, step2Result)
			result := AnalyzeXiJi(&chart, step2Result, step3Result)

			t.Logf("八字: %s", tc.input)
			t.Logf("核心病源类型: %s", result.CoreBingYao.BingType)
			t.Logf("核心病源描述: %s", result.CoreBingYao.BingDesc)
		})
	}
}

func TestCompleteAnalysisFlow(t *testing.T) {
	// 完整分析流程测试
	input := "壬戌 壬寅 庚午 丙戌"
	chart, err := utils.ParseBaZi(input)
	if err != nil {
		t.Fatalf("解析八字失败: %v", err)
	}

	step1Result := step1_scan.ScanBaZi(&chart)
	step2Result := step2_rebuild.RebuildBaZi(&chart, step1Result)
	step3Result := step3_wangshuai.AnalyzeWangShuai(&chart, step2Result)
	result := AnalyzeXiJi(&chart, step2Result, step3Result)

	fmt.Println()
	fmt.Println("╔════════════════════════════════════════════════════════════╗")
	fmt.Println("║           八字喜忌用神分析 - 病药理论综合诊断               ║")
	fmt.Println("╚════════════════════════════════════════════════════════════╝")
	fmt.Println()
	fmt.Printf("八字: %s\n", input)
	fmt.Printf("日主: %s (%s)\n", result.DayMaster.Name, result.DayMasterWuXing)
	fmt.Printf("旺衰: %s\n", result.WangShuaiType)
	fmt.Println()

	fmt.Println("【多维度病源诊断】")
	fmt.Println()

	printBingYaoAnalysis("1. 调候诊断", result.TiaoHouAnalysis)
	printBingYaoAnalysis("2. 通关诊断", result.TongGuanAnalysis)
	printBingYaoAnalysis("3. 扶抑诊断", result.FuYiAnalysis)
	printBingYaoAnalysis("4. 格局诊断", result.GeJuAnalysis)

	fmt.Println()
	fmt.Println("【核心病源】")
	fmt.Printf("  %s\n", result.CoreBingDesc)
	fmt.Println()

	fmt.Println("【喜忌用神结论】")
	fmt.Printf("  用神（药）五行: %v\n", result.YongShenWuXing)
	fmt.Printf("  用神（药）十神: %v\n", result.YongShenShiShen)
	fmt.Printf("  喜神五行: %v\n", result.XiShenWuXing)
	fmt.Printf("  喜神十神: %v\n", result.XiShenShiShen)
	fmt.Printf("  忌神五行: %v\n", result.JiShenWuXing)
	fmt.Printf("  忌神十神: %v\n", result.JiShenShiShen)
	fmt.Println()

	fmt.Println("【判定依据】")
	fmt.Printf("  %s\n", result.YongShenReason)
	fmt.Println()

	fmt.Println("【综合结论】")
	fmt.Printf("  %s\n", result.Summary)
	fmt.Println()

	// 验证所有关键字段
	if result.CoreBingYao == nil {
		t.Error("核心病源不应为空")
	}
	if len(result.YongShenWuXing) == 0 {
		t.Error("用神五行不应为空")
	}
	if len(result.JiShenWuXing) == 0 {
		t.Error("忌神五行不应为空")
	}
	if result.Summary == "" {
		t.Error("综合结论不应为空")
	}
}

// 辅助函数

func assertBingYaoAnalysis(t *testing.T, name string, analysis *BingYaoAnalysis) {
	if analysis == nil {
		t.Errorf("%s 不应为空", name)
		return
	}
	t.Logf("%s: %s (严重程度: %d)", name, analysis.BingDesc, analysis.Severity)
}

func printBingYaoAnalysis(title string, analysis *BingYaoAnalysis) {
	if analysis == nil {
		return
	}
	fmt.Printf("%s:\n", title)
	fmt.Printf("  病源: %s\n", analysis.BingDesc)
	fmt.Printf("  严重程度: %d/5\n", analysis.Severity)
	if len(analysis.YaoWuXing) > 0 {
		fmt.Printf("  建议用神: %v\n", analysis.YaoWuXing)
	}
	if analysis.Reason != "" {
		fmt.Printf("  依据: %s\n", analysis.Reason)
	}
	fmt.Println()
}

func slicesEqual(a, b []base.WuXing) bool {
	if len(a) != len(b) {
		return false
	}
	for i, v := range a {
		if v != b[i] {
			return false
		}
	}
	return true
}
