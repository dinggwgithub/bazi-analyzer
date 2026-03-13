package step3_wangshuai

import (
	"testing"

	"bazi-analyzer/internal/engine/step1_scan"
	"bazi-analyzer/internal/engine/step2_rebuild"
	"bazi-analyzer/pkg/utils"
)

func TestAnalyzeWangShuai(t *testing.T) {
	chart, err := utils.ParseBaZi("壬戌 壬寅 庚午 丙戌")
	if err != nil {
		t.Fatalf("解析八字失败: %v", err)
	}

	t.Logf("八字: %s", chart.String())
	t.Logf("日主: %s (%s)", chart.GetDayMaster().Name, chart.GetDayMaster().Element)

	step1Result := step1_scan.ScanBaZi(&chart)
	step2Result := step2_rebuild.RebuildBaZi(&chart, step1Result)
	step3Result := AnalyzeWangShuai(&chart, step2Result)

	t.Logf("\n=== 第三步：日主旺衰分析 ===")
	t.Logf("日主: %s (%s)", step3Result.DayMaster.Name, step3Result.DayMasterWuXing)
	t.Logf("得令: %v", step3Result.IsDeLing)
	t.Logf("得地: %v", step3Result.IsDeDi)
	t.Logf("得势: %v", step3Result.IsDeShi)
	t.Logf("生扶数量: %d", step3Result.ShengFuCount)
	t.Logf("克泄耗数量: %d", step3Result.KeXieHaoCount)
	t.Logf("旺衰结论: %s", step3Result.WangShuaiType)
	t.Logf("判断依据: %s", step3Result.Reason)
}
