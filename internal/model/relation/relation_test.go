package relation

import (
	"testing"

	"bazi-analyzer/internal/model/base"
	"bazi-analyzer/pkg/utils"
)

func TestTianGanWuHe(t *testing.T) {
	rel, ok := GetTianGanWuHe(base.Jia, base.Ji)
	if !ok {
		t.Error("甲己合土应该返回true")
	}
	if rel.Result != base.Earth {
		t.Errorf("甲己合土结果错误，期望: 土，实际: %s", rel.Result)
	}
	t.Logf("甲己合土测试通过: %s -> %s", rel.Type, rel.Result)
}

func TestTianGanSiChong(t *testing.T) {
	rel, ok := GetTianGanSiChong(base.Bing, base.Ren)
	if !ok {
		t.Error("丙壬冲应该返回true")
	}
	t.Logf("丙壬冲测试通过: %s", rel.Type)
}

func TestDiZhiSanHe(t *testing.T) {
	found := false
	for _, sanhe := range SanHeList {
		if ContainsDiZhi(sanhe.Elements, base.Yin) &&
			ContainsDiZhi(sanhe.Elements, base.WuDiZhi) &&
			ContainsDiZhi(sanhe.Elements, base.Xu) {
			if sanhe.Result != base.Fire {
				t.Errorf("寅午戌三合火结果错误，期望: 火，实际: %s", sanhe.Result)
			}
			found = true
			t.Logf("寅午戌三合火测试通过: %s -> %s", sanhe.Type, sanhe.Result)
			break
		}
	}
	if !found {
		t.Error("寅午戌三合火未找到")
	}
}

func TestDiZhiLiuHe(t *testing.T) {
	rel, ok := GetDiZhiLiuHe(base.Mao, base.Xu)
	if !ok {
		t.Error("卯戌合应该返回true")
	}
	t.Logf("卯戌合测试通过: %s", rel.Type)
}

func TestFirstStageComplete(t *testing.T) {
	t.Log("=== 第一阶段完整测试 ===")

	chart, err := ParseBaZiForTest("壬戌 壬寅 庚午 丙戌")
	if err != nil {
		t.Fatalf("八字解析失败: %v", err)
	}

	t.Logf("1. 八字解析成功: %s", chart.OriginalStr)
	t.Logf("   年柱: %s%s 月柱: %s%s 日柱: %s%s 时柱: %s%s",
		chart.YearPillar.TianGan.Name, chart.YearPillar.DiZhi.Name,
		chart.MonthPillar.TianGan.Name, chart.MonthPillar.DiZhi.Name,
		chart.DayPillar.TianGan.Name, chart.DayPillar.DiZhi.Name,
		chart.HourPillar.TianGan.Name, chart.HourPillar.DiZhi.Name)

	dayMaster := chart.GetDayMaster()
	t.Logf("2. 日主: %s (%s)", dayMaster.Name, dayMaster.Element)

	monthLing := chart.GetMonthLing()
	t.Logf("3. 月令: %s (%s)", monthLing.Name, monthLing.Element)

	tianGanList := chart.GetAllTianGan()
	_, hasBingRenChong := GetTianGanSiChong(tianGanList[1], tianGanList[3])
	t.Logf("4. 天干丙壬冲识别: %v (壬在月, 丙在时)", hasBingRenChong)

	diZhiList := chart.GetAllDiZhi()
	hasYin := ContainsDiZhi(diZhiList, base.Yin)
	hasWu := ContainsDiZhi(diZhiList, base.WuDiZhi)
	hasXu := ContainsDiZhi(diZhiList, base.Xu)
	t.Logf("5. 地支寅午戌三合识别: 寅%v 午%v 戌%v", hasYin, hasWu, hasXu)

	t.Log("=== 第一阶段测试完成 ===")
}

func ParseBaZiForTest(input string) (base.BaZiChart, error) {
	return utils.ParseBaZi(input)
}
