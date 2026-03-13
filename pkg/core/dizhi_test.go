package core

import (
	"testing"
)

func TestGetDiZhiInfo(t *testing.T) {
	info := GetDiZhiInfo(Zi)
	if info.Name != "子" || info.WuXing != Water || info.YinYang != Yang {
		t.Errorf("GetDiZhiInfo(Zi) wrong: %+v", info)
	}

	info = GetDiZhiInfo(YinZhi)
	if info.Name != "寅" || info.WuXing != Wood || info.YinYang != Yang {
		t.Errorf("GetDiZhiInfo(YinZhi) wrong: %+v", info)
	}
}

func TestGetDiZhiWuXing(t *testing.T) {
	tests := []struct {
		zhi      DiZhi
		expected WuXing
	}{
		{Zi, Water},
		{Chou, Earth},
		{YinZhi, Wood},
		{Mao, Wood},
		{Chen, Earth},
		{Si, Fire},
		{WuZhi, Fire},
		{Wei, Earth},
		{Shen, Metal},
		{You, Metal},
		{Xu, Earth},
		{Hai, Water},
	}

	for _, test := range tests {
		if wx := GetDiZhiWuXing(test.zhi); wx != test.expected {
			t.Errorf("GetDiZhiWuXing(%d) = %d, expected %d", test.zhi, wx, test.expected)
		}
	}
}

func TestGetDiZhiYinYang(t *testing.T) {
	tests := []struct {
		zhi      DiZhi
		expected YinYang
	}{
		{Zi, Yang},
		{Chou, Yin},
		{YinZhi, Yang},
		{Mao, Yin},
		{Chen, Yang},
		{Si, Yin},
		{WuZhi, Yang},
		{Wei, Yin},
		{Shen, Yang},
		{You, Yin},
		{Xu, Yang},
		{Hai, Yin},
	}

	for _, test := range tests {
		if yy := GetDiZhiYinYang(test.zhi); yy != test.expected {
			t.Errorf("GetDiZhiYinYang(%d) = %d, expected %d", test.zhi, yy, test.expected)
		}
	}
}

func TestGetCangGan(t *testing.T) {
	// 子藏癸
	cangGan := GetCangGan(Zi)
	if len(cangGan) != 1 || cangGan[0] != Gui {
		t.Errorf("GetCangGan(Zi) wrong: %v", cangGan)
	}

	// 寅藏甲丙戊
	cangGan = GetCangGan(YinZhi)
	if len(cangGan) != 3 {
		t.Errorf("GetCangGan(YinZhi) should have 3 cangGan, got %d", len(cangGan))
	}
	if cangGan[0] != Jia || cangGan[1] != Bing || cangGan[2] != Wu {
		t.Errorf("GetCangGan(YinZhi) wrong order: %v", cangGan)
	}
}

func TestGetBenQi(t *testing.T) {
	if GetBenQi(Zi) != Gui {
		t.Error("GetBenQi(Zi) should be Gui")
	}
	if GetBenQi(YinZhi) != Jia {
		t.Error("GetBenQi(YinZhi) should be Jia")
	}
}

func TestHasCangGan(t *testing.T) {
	// 寅中藏甲
	if !HasCangGan(YinZhi, Jia) {
		t.Error("YinZhi should contain Jia")
	}
	// 寅中不藏庚
	if HasCangGan(YinZhi, Geng) {
		t.Error("YinZhi should not contain Geng")
	}
}

func TestIsChangSheng(t *testing.T) {
	// 木长生于亥
	if !IsChangSheng(Hai, Wood) {
		t.Error("Hai should be ChangSheng for Wood")
	}
	// 火长生于寅
	if !IsChangSheng(YinZhi, Fire) {
		t.Error("YinZhi should be ChangSheng for Fire")
	}
	// 金长生于巳
	if !IsChangSheng(Si, Metal) {
		t.Error("Si should be ChangSheng for Metal")
	}
	// 水长生于申
	if !IsChangSheng(Shen, Water) {
		t.Error("Shen should be ChangSheng for Water")
	}
}

func TestIsDiWang(t *testing.T) {
	// 木旺于卯
	if !IsDiWang(Mao, Wood) {
		t.Error("Mao should be DiWang for Wood")
	}
	// 火旺于午
	if !IsDiWang(WuZhi, Fire) {
		t.Error("WuZhi should be DiWang for Fire")
	}
	// 金旺于酉
	if !IsDiWang(You, Metal) {
		t.Error("You should be DiWang for Metal")
	}
	// 水旺于子
	if !IsDiWang(Zi, Water) {
		t.Error("Zi should be DiWang for Water")
	}
}

func TestIsMuKu(t *testing.T) {
	// 木库在未
	if !IsMuKu(Wei, Wood) {
		t.Error("Wei should be MuKu for Wood")
	}
	// 火库在戌
	if !IsMuKu(Xu, Fire) {
		t.Error("Xu should be MuKu for Fire")
	}
	// 金库在丑
	if !IsMuKu(Chou, Metal) {
		t.Error("Chou should be MuKu for Metal")
	}
	// 水库在辰
	if !IsMuKu(Chen, Water) {
		t.Error("Chen should be MuKu for Water")
	}
}

func TestIsTongXingGen(t *testing.T) {
	// 寅中有甲木（木之根）
	if !IsTongXingGen(YinZhi, Wood) {
		t.Error("YinZhi should be TongXingGen for Wood")
	}
	// 寅中有丙火（火之根）
	if !IsTongXingGen(YinZhi, Fire) {
		t.Error("YinZhi should be TongXingGen for Fire")
	}
	// 寅中没有庚金
	if IsTongXingGen(YinZhi, Metal) {
		t.Error("YinZhi should not be TongXingGen for Metal")
	}
}

func TestGetGenQiLevel(t *testing.T) {
	// 卯是木的帝旺位，级别应为4
	if level := GetGenQiLevel(Mao, Wood); level != 4 {
		t.Errorf("GetGenQiLevel(Mao, Wood) = %d, expected 4", level)
	}

	// 寅是木的本气位，级别应为3
	if level := GetGenQiLevel(YinZhi, Wood); level != 3 {
		t.Errorf("GetGenQiLevel(YinZhi, Wood) = %d, expected 3", level)
	}

	// 寅是火的中气位，级别应为2
	if level := GetGenQiLevel(YinZhi, Fire); level != 2 {
		t.Errorf("GetGenQiLevel(YinZhi, Fire) = %d, expected 2", level)
	}

	// 寅是土的余气位，级别应为1
	if level := GetGenQiLevel(YinZhi, Earth); level != 1 {
		t.Errorf("GetGenQiLevel(YinZhi, Earth) = %d, expected 1", level)
	}

	// 酉对木无根，级别应为0
	if level := GetGenQiLevel(You, Wood); level != 0 {
		t.Errorf("GetGenQiLevel(You, Wood) = %d, expected 0", level)
	}
}

func TestGetAllDiZhi(t *testing.T) {
	all := GetAllDiZhi()
	if len(all) != 12 {
		t.Errorf("GetAllDiZhi() should return 12 results, got %d", len(all))
	}
}
