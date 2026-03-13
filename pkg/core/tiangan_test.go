package core

import (
	"testing"
)

func TestGetTianGanInfo(t *testing.T) {
	info := GetTianGanInfo(Jia)
	if info.Name != "甲" || info.WuXing != Wood || info.YinYang != Yang {
		t.Errorf("GetTianGanInfo(Jia) wrong: %+v", info)
	}

	info = GetTianGanInfo(Ren)
	if info.Name != "壬" || info.WuXing != Water || info.YinYang != Yang {
		t.Errorf("GetTianGanInfo(Ren) wrong: %+v", info)
	}
}

func TestGetTianGanWuXing(t *testing.T) {
	tests := []struct {
		gan      TianGan
		expected WuXing
	}{
		{Jia, Wood},
		{Yi, Wood},
		{Bing, Fire},
		{Ding, Fire},
		{Wu, Earth},
		{Ji, Earth},
		{Geng, Metal},
		{Xin, Metal},
		{Ren, Water},
		{Gui, Water},
	}

	for _, test := range tests {
		if wx := GetTianGanWuXing(test.gan); wx != test.expected {
			t.Errorf("GetTianGanWuXing(%d) = %d, expected %d", test.gan, wx, test.expected)
		}
	}
}

func TestGetTianGanYinYang(t *testing.T) {
	tests := []struct {
		gan      TianGan
		expected YinYang
	}{
		{Jia, Yang},
		{Yi, Yin},
		{Bing, Yang},
		{Ding, Yin},
		{Wu, Yang},
		{Ji, Yin},
		{Geng, Yang},
		{Xin, Yin},
		{Ren, Yang},
		{Gui, Yin},
	}

	for _, test := range tests {
		if yy := GetTianGanYinYang(test.gan); yy != test.expected {
			t.Errorf("GetTianGanYinYang(%d) = %d, expected %d", test.gan, yy, test.expected)
		}
	}
}

func TestIsTianGanYang(t *testing.T) {
	if !IsTianGanYang(Jia) {
		t.Error("IsTianGanYang(Jia) should be true")
	}
	if IsTianGanYang(Yi) {
		t.Error("IsTianGanYang(Yi) should be false")
	}
}

func TestIsTianGanYin(t *testing.T) {
	if !IsTianGanYin(Yi) {
		t.Error("IsTianGanYin(Yi) should be true")
	}
	if IsTianGanYin(Jia) {
		t.Error("IsTianGanYin(Jia) should be false")
	}
}

func TestGetTianGanByName(t *testing.T) {
	tests := []struct {
		name     string
		expected TianGan
		found    bool
	}{
		{"甲", Jia, true},
		{"乙", Yi, true},
		{"丙", Bing, true},
		{"壬", Ren, true},
		{"X", 0, false},
		{"", 0, false},
	}

	for _, test := range tests {
		gan, found := GetTianGanByName(test.name)
		if found != test.found {
			t.Errorf("GetTianGanByName(%s) found = %v, expected %v", test.name, found, test.found)
		}
		if found && gan != test.expected {
			t.Errorf("GetTianGanByName(%s) = %d, expected %d", test.name, gan, test.expected)
		}
	}
}

func TestIsSameWuXing(t *testing.T) {
	if !IsSameWuXing(Jia, Yi) {
		t.Error("Jia and Yi should be same WuXing (Wood)")
	}
	if IsSameWuXing(Jia, Bing) {
		t.Error("Jia and Bing should not be same WuXing")
	}
}

func TestIsSameYinYang(t *testing.T) {
	if !IsSameYinYang(Jia, Bing) {
		t.Error("Jia and Bing should be same YinYang (Yang)")
	}
	if IsSameYinYang(Jia, Yi) {
		t.Error("Jia and Yi should not be same YinYang")
	}
}

func TestIsBiJie(t *testing.T) {
	if !IsBiJie(Jia, Yi) {
		t.Error("Jia and Yi should be BiJie (same WuXing)")
	}
	if IsBiJie(Jia, Bing) {
		t.Error("Jia and Bing should not be BiJie")
	}
}

func TestTianGanSheng(t *testing.T) {
	// 甲生丙丁（木生火）
	result := TianGanSheng(Jia)
	if len(result) != 2 {
		t.Errorf("TianGanSheng(Jia) should return 2 results, got %d", len(result))
	}
	foundBing := false
	foundDing := false
	for _, gan := range result {
		if gan == Bing {
			foundBing = true
		}
		if gan == Ding {
			foundDing = true
		}
	}
	if !foundBing || !foundDing {
		t.Error("TianGanSheng(Jia) should include Bing and Ding")
	}
}

func TestTianGanKe(t *testing.T) {
	// 甲克戊己（木克土）
	result := TianGanKe(Jia)
	if len(result) != 2 {
		t.Errorf("TianGanKe(Jia) should return 2 results, got %d", len(result))
	}
}

func TestKeTianGan(t *testing.T) {
	// 克甲者为庚辛（金克木）
	result := KeTianGan(Jia)
	if len(result) != 2 {
		t.Errorf("KeTianGan(Jia) should return 2 results, got %d", len(result))
	}
	foundGeng := false
	foundXin := false
	for _, gan := range result {
		if gan == Geng {
			foundGeng = true
		}
		if gan == Xin {
			foundXin = true
		}
	}
	if !foundGeng || !foundXin {
		t.Error("KeTianGan(Jia) should include Geng and Xin")
	}
}

func TestShengTianGan(t *testing.T) {
	// 生甲者为壬癸（水生木）
	result := ShengTianGan(Jia)
	if len(result) != 2 {
		t.Errorf("ShengTianGan(Jia) should return 2 results, got %d", len(result))
	}
	foundRen := false
	foundGui := false
	for _, gan := range result {
		if gan == Ren {
			foundRen = true
		}
		if gan == Gui {
			foundGui = true
		}
	}
	if !foundRen || !foundGui {
		t.Error("ShengTianGan(Jia) should include Ren and Gui")
	}
}

func TestGetAllTianGan(t *testing.T) {
	all := GetAllTianGan()
	if len(all) != 10 {
		t.Errorf("GetAllTianGan() should return 10 results, got %d", len(all))
	}
}
