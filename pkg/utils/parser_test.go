package utils

import (
	"testing"
)

func TestParseBaZi(t *testing.T) {
	input := "壬戌 壬寅 庚午 丙戌"
	chart, err := ParseBaZi(input)
	if err != nil {
		t.Fatalf("解析八字失败: %v", err)
	}

	if chart.YearPillar.TianGan.Name != "壬" {
		t.Errorf("年柱天干错误，期望: 壬，实际: %s", chart.YearPillar.TianGan.Name)
	}
	if chart.YearPillar.DiZhi.Name != "戌" {
		t.Errorf("年柱地支错误，期望: 戌，实际: %s", chart.YearPillar.DiZhi.Name)
	}

	if chart.DayPillar.TianGan.Name != "庚" {
		t.Errorf("日柱天干错误，期望: 庚，实际: %s", chart.DayPillar.TianGan.Name)
	}
	if chart.DayPillar.DiZhi.Name != "午" {
		t.Errorf("日柱地支错误，期望: 午，实际: %s", chart.DayPillar.DiZhi.Name)
	}

	dayMaster := chart.GetDayMaster()
	if dayMaster.Name != "庚" {
		t.Errorf("日主错误，期望: 庚，实际: %s", dayMaster.Name)
	}

	monthLing := chart.GetMonthLing()
	if monthLing.Name != "寅" {
		t.Errorf("月令错误，期望: 寅，实际: %s", monthLing.Name)
	}

	t.Logf("八字解析成功: %+v", chart)
}
