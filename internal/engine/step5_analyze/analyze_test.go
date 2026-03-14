package step5_analyze

import (
	"testing"
	"fmt"
	"bazi-analyzer/pkg/utils"
	"bazi-analyzer/internal/engine/step1_scan"
	"bazi-analyzer/internal/engine/step2_rebuild"
	"bazi-analyzer/internal/engine/step3_wangshuai"
	"bazi-analyzer/internal/engine/step4_xiji"
)

func TestAnalyzeInteractions(t *testing.T) {
	input := "壬戌 壬寅 庚午 丙戌"
	chart, err := utils.ParseBaZi(input)
	if err != nil {
		t.Fatalf("解析八字失败: %v", err)
	}

	step1Result := step1_scan.ScanBaZi(&chart)
	step2Result := step2_rebuild.RebuildBaZi(&chart, step1Result)
	step3Result := step3_wangshuai.AnalyzeWangShuai(&chart, step2Result)
	step4Result := step4_xiji.AnalyzeXiJi(&chart, step3Result, step2Result)
	result := AnalyzeInteractions(&chart, step1Result, step2Result, step4Result)

	t.Logf("八字: %s", input)
	t.Log("")
	t.Log("    === 第五步：细析作用关系 ===")
	t.Log("天干作用:")
	for _, inter := range result.TianGanInteractions {
		t.Logf("  %s%s %s %s%s [%s]", inter.From.Name, inter.From.Element, inter.Type, inter.To.Name, inter.To.Element, inter.Strength)
	}
	t.Log("地支作用:")
	for _, inter := range result.DiZhiInteractions {
		t.Logf("  %s: %s", inter.Type, inter.Description)
	}
	t.Log("关键点:")
	for _, kp := range result.KeyPoints {
		t.Logf("  - %s", kp)
	}

	fmt.Println()
	fmt.Println("测试通过!")
}
