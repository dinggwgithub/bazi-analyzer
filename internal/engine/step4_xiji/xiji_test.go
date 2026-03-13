package step4_xiji

import (
	"testing"
	"fmt"
	"bazi-analyzer/pkg/utils"
	"bazi-analyzer/internal/engine/step1_scan"
	"bazi-analyzer/internal/engine/step2_rebuild"
	"bazi-analyzer/internal/engine/step3_wangshuai"
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
	result := AnalyzeXiJi(&chart, step3Result)

	t.Logf("八字: %s", input)
	t.Logf("日主: %s (%s)", result.DayMaster.Name, result.DayMasterWuXing)
	t.Logf("旺衰: %s", result.WangShuaiType)
	t.Log("")
	t.Log("    === 第四步：喜忌用神分析 ===")
	t.Logf("喜用五行: %v", result.FavorableWuXing)
	t.Logf("喜用十神: %v", result.FavorableShiShen)
	t.Logf("忌神五行: %v", result.UnfavorableWuXing)
	t.Logf("忌神十神: %v", result.UnfavorableShiShen)
	t.Logf("判断依据: %s", result.Reason)

	if len(result.FavorableWuXing) == 0 {
		t.Error("喜用五行不应为空")
	}

	fmt.Println()
	fmt.Println("测试通过!")
}
