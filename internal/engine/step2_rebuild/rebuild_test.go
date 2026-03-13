package step2_rebuild

import (
	"testing"

	"bazi-analyzer/internal/engine/step1_scan"
	"bazi-analyzer/pkg/utils"
)

func TestRebuildBaZi(t *testing.T) {
	chart, err := utils.ParseBaZi("壬戌 壬寅 庚午 丙戌")
	if err != nil {
		t.Fatalf("解析八字失败: %v", err)
	}

	step1Result := step1_scan.ScanBaZi(&chart)
	step2Result := RebuildBaZi(&chart, step1Result)

	t.Logf("=== 第二步：格局重构结果 ===")
	t.Logf("是否有变化: %v", step2Result.HasChanges)

	t.Logf("\n天干五行重定:")
	for _, re := range step2Result.RebuiltTianGan {
		status := "原"
		if re.IsRebuilt {
			status = "重定"
		}
		t.Logf("  %s: %s -> %s [%s] %s", re.OriginalName, re.OriginalWuXing, re.RebuiltWuXing, status, re.Reason)
	}

	t.Logf("\n地支五行重定:")
	for _, re := range step2Result.RebuiltDiZhi {
		status := "原"
		if re.IsRebuilt {
			status = "重定"
		}
		t.Logf("  %s: %s -> %s [%s] %s", re.OriginalName, re.OriginalWuXing, re.RebuiltWuXing, status, re.Reason)
	}

	t.Logf("\n被吸收的地支: %v", step2Result.AbsorbedDiZhi)
	t.Logf("被吸收的天干: %v", step2Result.AbsorbedTianGan)

	if len(step2Result.AbsorbedDiZhi) == 3 {
		t.Log("✓ 寅午戌三合火合化成功，寅、午、戌三个地支均被吸收")
	}
}
