package step1_scan

import (
	"testing"

	"bazi-analyzer/pkg/utils"
)

func TestScanBaZi(t *testing.T) {
	chart, err := utils.ParseBaZi("壬戌 壬寅 庚午 丙戌")
	if err != nil {
		t.Fatalf("解析八字失败: %v", err)
	}

	result := ScanBaZi(&chart)

	t.Logf("天干五合: %d个", len(result.TianGanWuHe))
	for _, r := range result.TianGanWuHe {
		names := ""
		for _, tg := range r.Elements {
			names += tg.Name
		}
		t.Logf("  - %s%s", names, r.Type)
	}

	t.Logf("天干四冲: %d个", len(result.TianGanSiChong))
	for _, r := range result.TianGanSiChong {
		names := ""
		for _, tg := range r.Elements {
			names += tg.Name
		}
		t.Logf("  - %s%s", names, r.Type)
	}

	t.Logf("地支三会: %d个", len(result.DiZhiSanHui))
	for _, r := range result.DiZhiSanHui {
		names := ""
		for _, dz := range r.Elements {
			names += dz.Name
		}
		t.Logf("  - %s%s", names, r.Type)
	}

	t.Logf("地支三合: %d个", len(result.DiZhiSanHe))
	for _, r := range result.DiZhiSanHe {
		names := ""
		for _, dz := range r.Elements {
			names += dz.Name
		}
		t.Logf("  - %s%s", names, r.Type)
	}

	t.Logf("地支六合: %d个", len(result.DiZhiLiuHe))
	for _, r := range result.DiZhiLiuHe {
		names := ""
		for _, dz := range r.Elements {
			names += dz.Name
		}
		t.Logf("  - %s%s", names, r.Type)
	}

	if len(result.TianGanSiChong) == 0 {
		t.Error("应该检测到天干四冲(丙壬冲)")
	}

	if len(result.DiZhiSanHe) == 0 {
		t.Error("应该检测到地支三合(寅午戌三合火)")
	}
}
