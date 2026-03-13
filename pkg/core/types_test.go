package core

import (
	"testing"
)

func TestTianGanString(t *testing.T) {
	tests := []struct {
		gan      TianGan
		expected string
	}{
		{Jia, "甲"},
		{Yi, "乙"},
		{Bing, "丙"},
		{Ding, "丁"},
		{Wu, "戊"},
		{Ji, "己"},
		{Geng, "庚"},
		{Xin, "辛"},
		{Ren, "壬"},
		{Gui, "癸"},
	}

	for _, test := range tests {
		if test.gan.String() != test.expected {
			t.Errorf("TianGan(%d).String() = %s, expected %s", test.gan, test.gan.String(), test.expected)
		}
	}
}

func TestDiZhiString(t *testing.T) {
	tests := []struct {
		zhi      DiZhi
		expected string
	}{
		{Zi, "子"},
		{Chou, "丑"},
		{YinZhi, "寅"},
		{Mao, "卯"},
		{Chen, "辰"},
		{Si, "巳"},
		{WuZhi, "午"},
		{Wei, "未"},
		{Shen, "申"},
		{You, "酉"},
		{Xu, "戌"},
		{Hai, "亥"},
	}

	for _, test := range tests {
		if test.zhi.String() != test.expected {
			t.Errorf("DiZhi(%d).String() = %s, expected %s", test.zhi, test.zhi.String(), test.expected)
		}
	}
}

func TestWuXingString(t *testing.T) {
	tests := []struct {
		wx       WuXing
		expected string
	}{
		{Wood, "木"},
		{Fire, "火"},
		{Earth, "土"},
		{Metal, "金"},
		{Water, "水"},
	}

	for _, test := range tests {
		if test.wx.String() != test.expected {
			t.Errorf("WuXing(%d).String() = %s, expected %s", test.wx, test.wx.String(), test.expected)
		}
	}
}

func TestParseTianGan(t *testing.T) {
	tests := []struct {
		input    string
		expected TianGan
		wantErr  bool
	}{
		{"甲", Jia, false},
		{"乙", Yi, false},
		{"丙", Bing, false},
		{"丁", Ding, false},
		{"戊", Wu, false},
		{"己", Ji, false},
		{"庚", Geng, false},
		{"辛", Xin, false},
		{"壬", Ren, false},
		{"癸", Gui, false},
		{"X", 0, true},
		{"", 0, true},
	}

	for _, test := range tests {
		result, err := ParseTianGan(test.input)
		if test.wantErr {
			if err == nil {
				t.Errorf("ParseTianGan(%s) expected error, got nil", test.input)
			}
		} else {
			if err != nil {
				t.Errorf("ParseTianGan(%s) unexpected error: %v", test.input, err)
			}
			if result != test.expected {
				t.Errorf("ParseTianGan(%s) = %d, expected %d", test.input, result, test.expected)
			}
		}
	}
}

func TestParseDiZhi(t *testing.T) {
	tests := []struct {
		input    string
		expected DiZhi
		wantErr  bool
	}{
		{"子", Zi, false},
		{"丑", Chou, false},
		{"寅", YinZhi, false},
		{"卯", Mao, false},
		{"辰", Chen, false},
		{"巳", Si, false},
		{"午", WuZhi, false},
		{"未", Wei, false},
		{"申", Shen, false},
		{"酉", You, false},
		{"戌", Xu, false},
		{"亥", Hai, false},
		{"X", 0, true},
		{"", 0, true},
	}

	for _, test := range tests {
		result, err := ParseDiZhi(test.input)
		if test.wantErr {
			if err == nil {
				t.Errorf("ParseDiZhi(%s) expected error, got nil", test.input)
			}
		} else {
			if err != nil {
				t.Errorf("ParseDiZhi(%s) unexpected error: %v", test.input, err)
			}
			if result != test.expected {
				t.Errorf("ParseDiZhi(%s) = %d, expected %d", test.input, result, test.expected)
			}
		}
	}
}

func TestParseGanZhi(t *testing.T) {
	tests := []struct {
		input       string
		expectedGan TianGan
		expectedZhi DiZhi
		wantErr     bool
	}{
		{"甲子", Jia, Zi, false},
		{"乙丑", Yi, Chou, false},
		{"丙寅", Bing, YinZhi, false},
		{"壬戌", Ren, Xu, false},
		{"庚午", Geng, WuZhi, false},
		{"甲", 0, 0, true},
		{"甲子寅", 0, 0, true},
		{"", 0, 0, true},
	}

	for _, test := range tests {
		result, err := ParseGanZhi(test.input)
		if test.wantErr {
			if err == nil {
				t.Errorf("ParseGanZhi(%s) expected error, got nil", test.input)
			}
		} else {
			if err != nil {
				t.Errorf("ParseGanZhi(%s) unexpected error: %v", test.input, err)
			}
			if result.Gan != test.expectedGan {
				t.Errorf("ParseGanZhi(%s).Gan = %d, expected %d", test.input, result.Gan, test.expectedGan)
			}
			if result.Zhi != test.expectedZhi {
				t.Errorf("ParseGanZhi(%s).Zhi = %d, expected %d", test.input, result.Zhi, test.expectedZhi)
			}
		}
	}
}

func TestParseBazi(t *testing.T) {
	// 测试案例: 壬戌 壬寅 庚午 丙戌
	bazi, err := ParseBazi("壬戌 壬寅 庚午 丙戌")
	if err != nil {
		t.Fatalf("ParseBazi failed: %v", err)
	}

	// 验证年柱
	if bazi.Year.Gan != Ren || bazi.Year.Zhi != Xu {
		t.Errorf("Year column wrong: got %s, expected 壬戌", bazi.Year.String())
	}

	// 验证月柱
	if bazi.Month.Gan != Ren || bazi.Month.Zhi != YinZhi {
		t.Errorf("Month column wrong: got %s, expected 壬寅", bazi.Month.String())
	}

	// 验证日柱
	if bazi.Day.Gan != Geng || bazi.Day.Zhi != WuZhi {
		t.Errorf("Day column wrong: got %s, expected 庚午", bazi.Day.String())
	}

	// 验证时柱
	if bazi.Hour.Gan != Bing || bazi.Hour.Zhi != Xu {
		t.Errorf("Hour column wrong: got %s, expected 丙戌", bazi.Hour.String())
	}

	// 验证日主
	if bazi.GetDayGan() != Geng {
		t.Errorf("Day Gan wrong: got %d, expected Geng", bazi.GetDayGan())
	}

	// 验证月令
	if bazi.GetMonthZhi() != YinZhi {
		t.Errorf("Month Zhi wrong: got %d, expected YinZhi", bazi.GetMonthZhi())
	}
}

func TestParseBaziInvalid(t *testing.T) {
	tests := []struct {
		input string
	}{
		{""},
		{"壬戌"},
		{"壬戌 壬寅 庚午"},
		{"壬戌 壬寅 庚午 丙戌 甲子"},
		{"invalid input"},
	}

	for _, test := range tests {
		_, err := ParseBazi(test.input)
		if err == nil {
			t.Errorf("ParseBazi(%s) expected error, got nil", test.input)
		}
	}
}

func TestBaziString(t *testing.T) {
	bazi := &Bazi{
		Year:  GanZhi{Gan: Ren, Zhi: Xu},
		Month: GanZhi{Gan: Ren, Zhi: YinZhi},
		Day:   GanZhi{Gan: Geng, Zhi: WuZhi},
		Hour:  GanZhi{Gan: Bing, Zhi: Xu},
	}

	expected := "壬戌 壬寅 庚午 丙戌"
	if bazi.String() != expected {
		t.Errorf("Bazi.String() = %s, expected %s", bazi.String(), expected)
	}
}

func TestGetGanZhiByZhuWei(t *testing.T) {
	bazi := &Bazi{
		Year:  GanZhi{Gan: Ren, Zhi: Xu},
		Month: GanZhi{Gan: Ren, Zhi: YinZhi},
		Day:   GanZhi{Gan: Geng, Zhi: WuZhi},
		Hour:  GanZhi{Gan: Bing, Zhi: Xu},
	}

	if gz := bazi.GetGanZhiByZhuWei(Year); gz != bazi.Year {
		t.Errorf("GetGanZhiByZhuWei(Year) wrong")
	}
	if gz := bazi.GetGanZhiByZhuWei(Month); gz != bazi.Month {
		t.Errorf("GetGanZhiByZhuWei(Month) wrong")
	}
	if gz := bazi.GetGanZhiByZhuWei(Day); gz != bazi.Day {
		t.Errorf("GetGanZhiByZhuWei(Day) wrong")
	}
	if gz := bazi.GetGanZhiByZhuWei(Hour); gz != bazi.Hour {
		t.Errorf("GetGanZhiByZhuWei(Hour) wrong")
	}
}
