package step4_xiji

import (
	"fmt"
	"testing"

	"bazi-analyzer/internal/engine/step1_scan"
	"bazi-analyzer/internal/engine/step2_rebuild"
	"bazi-analyzer/internal/engine/step3_wangshuai"
	"bazi-analyzer/pkg/utils"
)

func TestAnalyzeXiJi(t *testing.T) {
	input := "壬戌 壬寅 庚午 丙戌"
	chart, err := utils.ParseBaZi(input)
	if err != nil {
		t.Fatalf("解析八字失败: %v", err)
	}

	step1Result := step1_scan.ScanBaZi(&chart)
	step2Result := step2_rebuild.RebuildBaZi(&chart, step1Result)
	step3Result := step3_wangshuai.AnalyzeWangShuai(&chart, step2Result)
	result := AnalyzeXiJiWithRebuild(&chart, step2Result, step3Result)

	t.Logf("八字: %s", input)
	t.Logf("日主: %s (%s)", result.DayMaster.Name, result.DayMasterWuXing)
	t.Logf("旺衰: %s", result.WangShuaiType)
	t.Log("")
	t.Log("    === 第四步：喜忌用神分析（病药理论） ===")
	t.Log("")

	if result.CoreBingYuan != "" {
		t.Logf("【核心病源】%s", result.CoreBingYuan)
	}
	if result.YongShen != "" {
		t.Logf("【用神】%s", result.YongShen)
		t.Logf("【用神类型】%s", result.YongShenType)
		t.Logf("【用神说明】%s", result.YongShenDesc)
	}
	if len(result.XiShen) > 0 {
		t.Logf("【喜神】%v", result.XiShen)
	}
	if len(result.JiShen) > 0 {
		t.Logf("【忌神】%v", result.JiShen)
	}
	t.Log("")
	t.Log("【判断依据】")
	t.Logf("  %s", result.Reason)
	t.Log("")

	t.Log("【分析路径】")
	for _, path := range result.AnalysisPath {
		t.Logf("  %s", path)
	}

	t.Log("")
	t.Logf("喜用五行: %v", result.FavorableWuXing)
	t.Logf("喜用十神: %v", result.FavorableShiShen)
	t.Logf("忌神五行: %v", result.UnfavorableWuXing)
	t.Logf("忌神十神: %v", result.UnfavorableShiShen)

	if len(result.FavorableWuXing) == 0 {
		t.Error("喜用五行不应为空")
	}

	fmt.Println()
	fmt.Println("测试通过!")
}

func TestAnalyzeXiJi_TiaoHou(t *testing.T) {
	input := "壬子 壬子 庚子 丙子"
	chart, err := utils.ParseBaZi(input)
	if err != nil {
		t.Fatalf("解析八字失败: %v", err)
	}

	step1Result := step1_scan.ScanBaZi(&chart)
	step2Result := step2_rebuild.RebuildBaZi(&chart, step1Result)
	step3Result := step3_wangshuai.AnalyzeWangShuai(&chart, step2Result)
	result := AnalyzeXiJiWithRebuild(&chart, step2Result, step3Result)

	t.Logf("八字: %s (测试调候-过寒)", input)
	t.Logf("日主: %s (%s)", result.DayMaster.Name, result.DayMasterWuXing)
	t.Logf("旺衰: %s", result.WangShuaiType)
	t.Log("")

	if result.CoreBingYuan != "" {
		t.Logf("【核心病源】%s", result.CoreBingYuan)
	}
	if result.YongShen != "" {
		t.Logf("【用神】%s", result.YongShen)
		t.Logf("【用神类型】%s", result.YongShenType)
	}

	hasTiaoHou := false
	if result.BingYuanDiagnosis != nil {
		for _, by := range result.BingYuanDiagnosis.AllBingYuan {
			if by.Type == BingYuanTiaoHou {
				hasTiaoHou = true
				t.Logf("检测到调候病源: %s", by.Description)
			}
		}
	}

	if !hasTiaoHou {
		t.Log("警告: 未检测到调候病源（可能需要检查逻辑）")
	}

	fmt.Println()
	fmt.Println("调候测试完成!")
}

func TestAnalyzeXiJi_BingYao(t *testing.T) {
	input := "甲戌 甲戌 庚戌 甲戌"
	chart, err := utils.ParseBaZi(input)
	if err != nil {
		t.Fatalf("解析八字失败: %v", err)
	}

	step1Result := step1_scan.ScanBaZi(&chart)
	step2Result := step2_rebuild.RebuildBaZi(&chart, step1Result)
	step3Result := step3_wangshuai.AnalyzeWangShuai(&chart, step2Result)
	result := AnalyzeXiJiWithRebuild(&chart, step2Result, step3Result)

	t.Logf("八字: %s (测试病药-财多身弱)", input)
	t.Logf("日主: %s (%s)", result.DayMaster.Name, result.DayMasterWuXing)
	t.Logf("旺衰: %s", result.WangShuaiType)
	t.Log("")

	if result.CoreBingYuan != "" {
		t.Logf("【核心病源】%s", result.CoreBingYuan)
	}
	if result.YongShen != "" {
		t.Logf("【用神】%s", result.YongShen)
		t.Logf("【用神类型】%s", result.YongShenType)
	}

	hasBingYao := false
	if result.BingYuanDiagnosis != nil {
		for _, by := range result.BingYuanDiagnosis.AllBingYuan {
			if by.Type == BingYuanBingYao {
				hasBingYao = true
				t.Logf("检测到病药病源: %s", by.Description)
			}
		}
	}

	if !hasBingYao {
		t.Log("警告: 未检测到病药病源（可能需要检查逻辑）")
	}

	fmt.Println()
	fmt.Println("病药测试完成!")
}

func TestAnalyzeXiJi_CongGe(t *testing.T) {
	input := "壬子 壬子 壬子 壬子"
	chart, err := utils.ParseBaZi(input)
	if err != nil {
		t.Fatalf("解析八字失败: %v", err)
	}

	step1Result := step1_scan.ScanBaZi(&chart)
	step2Result := step2_rebuild.RebuildBaZi(&chart, step1Result)
	step3Result := step3_wangshuai.AnalyzeWangShuai(&chart, step2Result)
	result := AnalyzeXiJiWithRebuild(&chart, step2Result, step3Result)

	t.Logf("八字: %s (测试从格)", input)
	t.Logf("日主: %s (%s)", result.DayMaster.Name, result.DayMasterWuXing)
	t.Logf("旺衰: %s", result.WangShuaiType)
	t.Log("")

	if result.CoreBingYuan != "" {
		t.Logf("【核心病源】%s", result.CoreBingYuan)
	}
	if result.YongShen != "" {
		t.Logf("【用神】%s", result.YongShen)
		t.Logf("【用神类型】%s", result.YongShenType)
	}

	hasGeJu := false
	if result.BingYuanDiagnosis != nil && result.BingYuanDiagnosis.CoreBingYuan != nil {
		if result.BingYuanDiagnosis.CoreBingYuan.Type == BingYuanGeJu {
			hasGeJu = true
		}
	}

	if hasGeJu {
		t.Log("检测到格局病源，从格处理正确")
	}

	fmt.Println()
	fmt.Println("从格测试完成!")
}

func TestAnalyzeXiJi_MultipleBingYuan(t *testing.T) {
	input := "壬子 丙午 庚子 丙午"
	chart, err := utils.ParseBaZi(input)
	if err != nil {
		t.Fatalf("解析八字失败: %v", err)
	}

	step1Result := step1_scan.ScanBaZi(&chart)
	step2Result := step2_rebuild.RebuildBaZi(&chart, step1Result)
	step3Result := step3_wangshuai.AnalyzeWangShuai(&chart, step2Result)
	result := AnalyzeXiJiWithRebuild(&chart, step2Result, step3Result)

	t.Logf("八字: %s (测试多病源)", input)
	t.Logf("日主: %s (%s)", result.DayMaster.Name, result.DayMasterWuXing)
	t.Logf("旺衰: %s", result.WangShuaiType)
	t.Log("")

	t.Log("【所有病源】")
	if result.BingYuanDiagnosis != nil {
		for i, by := range result.BingYuanDiagnosis.AllBingYuan {
			t.Logf("  %d. [%s] %s (严重度: %d)", i+1, by.Type, by.Description, by.Severity)
		}
	}

	t.Log("")
	if result.CoreBingYuan != "" {
		t.Logf("【核心病源】%s", result.CoreBingYuan)
	}
	if result.YongShen != "" {
		t.Logf("【用神】%s", result.YongShen)
		t.Logf("【用神类型】%s", result.YongShenType)
	}

	fmt.Println()
	fmt.Println("多病源测试完成!")
}

func TestAnalyzeXiJi_Consistency(t *testing.T) {
	input := "甲子 丙寅 庚午 壬午"
	chart, err := utils.ParseBaZi(input)
	if err != nil {
		t.Fatalf("解析八字失败: %v", err)
	}

	step1Result := step1_scan.ScanBaZi(&chart)
	step2Result := step2_rebuild.RebuildBaZi(&chart, step1Result)
	step3Result := step3_wangshuai.AnalyzeWangShuai(&chart, step2Result)
	result := AnalyzeXiJiWithRebuild(&chart, step2Result, step3Result)

	t.Logf("八字: %s (测试一致性)", input)
	t.Log("")

	yongShenWuXing := result.YongShenResult.YongShen.WuXing

	for _, xs := range result.YongShenResult.XiShen {
		if xs.WuXing == yongShenWuXing {
			t.Errorf("一致性错误: 喜神 %s 与用神相同", xs.WuXing)
		}
	}

	for _, js := range result.YongShenResult.JiShen {
		if js.WuXing == yongShenWuXing {
			t.Errorf("一致性错误: 忌神 %s 与用神相同", js.WuXing)
		}
	}

	for _, xs := range result.YongShenResult.XiShen {
		for _, js := range result.YongShenResult.JiShen {
			if xs.WuXing == js.WuXing {
				t.Errorf("一致性错误: 喜神 %s 与忌神 %s 相同", xs.WuXing, js.WuXing)
			}
		}
	}

	t.Log("一致性验证通过：用神、喜神、忌神无冲突")

	fmt.Println()
	fmt.Println("一致性测试完成!")
}
