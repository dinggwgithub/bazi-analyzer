package step6_conclusion

import (
	"fmt"
	"testing"

	"bazi-analyzer/internal/engine/step1_scan"
	"bazi-analyzer/internal/engine/step2_rebuild"
	"bazi-analyzer/internal/engine/step3_wangshuai"
	"bazi-analyzer/internal/engine/step4_xiji"
	"bazi-analyzer/internal/engine/step5_analyze"
	"bazi-analyzer/pkg/utils"
)

func TestAnalyzeConclusion(t *testing.T) {
	input := "壬戌 己酉 癸丑 辛酉"
	// input := "壬戌 壬寅 庚午 丙戌"
	chart, err := utils.ParseBaZi(input)
	if err != nil {
		t.Fatalf("解析八字失败: %v", err)
	}

	step1Result := step1_scan.ScanBaZi(&chart)
	step2Result := step2_rebuild.RebuildBaZi(&chart, step1Result)
	step3Result := step3_wangshuai.AnalyzeWangShuai(&chart, step2Result)
	step4Result := step4_xiji.AnalyzeXiJi(&chart, step3Result)
	step5Result := step5_analyze.AnalyzeInteractions(&chart, step1Result, step2Result, step4Result)
	result := AnalyzeConclusion(&chart, step1Result, step2Result, step3Result, step4Result, step5Result)

	t.Logf("八字: %s", input)
	t.Log("")
	t.Log("    === 第六步：综合论断 ===")
	t.Logf("格局类型: %s", result.GeJuType)
	t.Logf("格局层次: %s", result.GeJuLevel)
	t.Logf("格局描述: %s", result.GeJuDescription)
	t.Log("")
	t.Log("岁运分析:")
	for _, sy := range result.SuiYunAnalysis {
		t.Logf("  - %s", sy)
	}
	t.Log("")
	t.Logf("总体总结: %s", result.OverallSummary)
	t.Log("")
	t.Log("关键建议:")
	for _, ka := range result.KeyAdvice {
		t.Logf("  - %s", ka)
	}

	fmt.Println()
	fmt.Println("第六步测试通过!")
}

func TestFullAnalysisFlow(t *testing.T) {
	// input := "壬戌 壬寅 庚午 丙戌"
	input := "壬戌 己酉 癸丑 辛酉"
	chart, err := utils.ParseBaZi(input)
	if err != nil {
		t.Fatalf("解析八字失败: %v", err)
	}

	t.Log("==================================================")
	t.Logf("           八字完整分析: %s", input)
	t.Log("==================================================")
	t.Log("")

	step1Result := step1_scan.ScanBaZi(&chart)
	t.Log("=== 第一步：扫描全局特殊关系 ===")
	t.Log("天干五合:")
	for _, wh := range step1Result.TianGanWuHe {
		t.Logf("  %s-%s", wh.Elements[0].Name, wh.Elements[1].Name)
	}
	t.Log("天干四冲:")
	for _, sc := range step1Result.TianGanSiChong {
		t.Logf("  %s-%s", sc.Elements[0].Name, sc.Elements[1].Name)
	}
	t.Log("地支三会:")
	for _, sh := range step1Result.DiZhiSanHui {
		t.Logf("  %s-%s-%s", sh.Elements[0].Name, sh.Elements[1].Name, sh.Elements[2].Name)
	}
	t.Log("地支三合:")
	for _, sh := range step1Result.DiZhiSanHe {
		t.Logf("  %s-%s-%s", sh.Elements[0].Name, sh.Elements[1].Name, sh.Elements[2].Name)
	}
	t.Log("地支六合:")
	for _, lh := range step1Result.DiZhiLiuHe {
		t.Logf("  %s-%s", lh.Elements[0].Name, lh.Elements[1].Name)
	}
	t.Log("")

	step2Result := step2_rebuild.RebuildBaZi(&chart, step1Result)
	t.Log("=== 第二步：格局重构 ===")
	t.Log("地支五行重定:")
	for _, rd := range step2Result.RebuiltDiZhi {
		t.Logf("  %s(%s) -> %s [%s]", rd.OriginalName, rd.OriginalWuXing, rd.RebuiltWuXing, rd.Reason)
	}
	t.Log("")

	step3Result := step3_wangshuai.AnalyzeWangShuai(&chart, step2Result)
	t.Log("=== 第三步：日主旺衰分析 ===")
	t.Logf("  日主: %s(%s)", step3Result.DayMaster.Name, step3Result.DayMasterWuXing)
	t.Logf("  得令: %t, 得地: %t, 得势: %t", step3Result.IsDeLing, step3Result.IsDeDi, step3Result.IsDeShi)
	t.Logf("  生扶: %d, 克泄耗: %d", step3Result.ShengFuCount, step3Result.KeXieHaoCount)
	t.Logf("  结论: %s", step3Result.WangShuaiType)
	t.Log("")

	step4Result := step4_xiji.AnalyzeXiJi(&chart, step3Result)
	t.Log("=== 第四步：喜忌用神分析 ===")
	t.Logf("  喜用五行: %v", step4Result.FavorableWuXing)
	t.Logf("  喜用十神: %v", step4Result.FavorableShiShen)
	t.Logf("  忌神五行: %v", step4Result.UnfavorableWuXing)
	t.Logf("  忌神十神: %v", step4Result.UnfavorableShiShen)
	t.Log("")

	step5Result := step5_analyze.AnalyzeInteractions(&chart, step1Result, step2Result, step4Result)
	t.Log("=== 第五步：细析作用关系 ===")
	t.Log("关键点:")
	for _, kp := range step5Result.KeyPoints {
		t.Logf("  - %s", kp)
	}
	t.Log("")

	result := AnalyzeConclusion(&chart, step1Result, step2Result, step3Result, step4Result, step5Result)
	t.Log("=== 第六步：综合论断 ===")
	t.Logf("  格局类型: %s", result.GeJuType)
	t.Logf("  格局层次: %s", result.GeJuLevel)
	t.Logf("  总体总结: %s", result.OverallSummary)
	t.Log("")
	t.Log("==================================================")
	t.Log("                    分析完成")
	t.Log("==================================================")

	fmt.Println()
	fmt.Println("完整6步流程测试通过!")
}
