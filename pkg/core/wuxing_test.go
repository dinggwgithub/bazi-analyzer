package core

import (
	"testing"
)

func TestGetRelation(t *testing.T) {
	// 木生火
	if rel := GetRelation(Wood, Fire); rel != Sheng {
		t.Errorf("GetRelation(Wood, Fire) = %d, expected Sheng", rel)
	}

	// 木克土
	if rel := GetRelation(Wood, Earth); rel != Ke {
		t.Errorf("GetRelation(Wood, Earth) = %d, expected Ke", rel)
	}

	// 木同木
	if rel := GetRelation(Wood, Wood); rel != Tong {
		t.Errorf("GetRelation(Wood, Wood) = %d, expected Tong", rel)
	}
}

func TestIsSheng(t *testing.T) {
	// 木生火
	if !IsSheng(Wood, Fire) {
		t.Error("IsSheng(Wood, Fire) should be true")
	}
	// 火生土
	if !IsSheng(Fire, Earth) {
		t.Error("IsSheng(Fire, Earth) should be true")
	}
	// 木不生金
	if IsSheng(Wood, Metal) {
		t.Error("IsSheng(Wood, Metal) should be false")
	}
}

func TestIsKe(t *testing.T) {
	// 木克土
	if !IsKe(Wood, Earth) {
		t.Error("IsKe(Wood, Earth) should be true")
	}
	// 土克水
	if !IsKe(Earth, Water) {
		t.Error("IsKe(Earth, Water) should be true")
	}
	// 木不克火
	if IsKe(Wood, Fire) {
		t.Error("IsKe(Wood, Fire) should be false")
	}
}

func TestGetShengWo(t *testing.T) {
	// 生木者为水
	if wx := GetShengWo(Wood); wx != Water {
		t.Errorf("GetShengWo(Wood) = %d, expected Water", wx)
	}
	// 生火者为木
	if wx := GetShengWo(Fire); wx != Wood {
		t.Errorf("GetShengWo(Fire) = %d, expected Wood", wx)
	}
	// 生土者为火
	if wx := GetShengWo(Earth); wx != Fire {
		t.Errorf("GetShengWo(Earth) = %d, expected Fire", wx)
	}
}

func TestGetWoSheng(t *testing.T) {
	// 木生火
	if wx := GetWoSheng(Wood); wx != Fire {
		t.Errorf("GetWoSheng(Wood) = %d, expected Fire", wx)
	}
	// 火生土
	if wx := GetWoSheng(Fire); wx != Earth {
		t.Errorf("GetWoSheng(Fire) = %d, expected Earth", wx)
	}
	// 土生金
	if wx := GetWoSheng(Earth); wx != Metal {
		t.Errorf("GetWoSheng(Earth) = %d, expected Metal", wx)
	}
}

func TestGetKeWo(t *testing.T) {
	// 克木者为金
	if wx := GetKeWo(Wood); wx != Metal {
		t.Errorf("GetKeWo(Wood) = %d, expected Metal", wx)
	}
	// 克火者为水
	if wx := GetKeWo(Fire); wx != Water {
		t.Errorf("GetKeWo(Fire) = %d, expected Water", wx)
	}
	// 克土者为木
	if wx := GetKeWo(Earth); wx != Wood {
		t.Errorf("GetKeWo(Earth) = %d, expected Wood", wx)
	}
}

func TestGetWoKe(t *testing.T) {
	// 木克土
	if wx := GetWoKe(Wood); wx != Earth {
		t.Errorf("GetWoKe(Wood) = %d, expected Earth", wx)
	}
	// 火克金
	if wx := GetWoKe(Fire); wx != Metal {
		t.Errorf("GetWoKe(Fire) = %d, expected Metal", wx)
	}
	// 土克水
	if wx := GetWoKe(Earth); wx != Water {
		t.Errorf("GetWoKe(Earth) = %d, expected Water", wx)
	}
}

func TestIsShengWo(t *testing.T) {
	// 水生木
	if !IsShengWo(Wood, Water) {
		t.Error("IsShengWo(Wood, Water) should be true")
	}
	// 木生火，火不生木
	if IsShengWo(Wood, Fire) {
		t.Error("IsShengWo(Wood, Fire) should be false")
	}
}

func TestIsWoSheng(t *testing.T) {
	// 木生火
	if !IsWoSheng(Wood, Fire) {
		t.Error("IsWoSheng(Wood, Fire) should be true")
	}
	// 木不生水
	if IsWoSheng(Wood, Water) {
		t.Error("IsWoSheng(Wood, Water) should be false")
	}
}

func TestIsKeWo(t *testing.T) {
	// 金克木
	if !IsKeWo(Wood, Metal) {
		t.Error("IsKeWo(Wood, Metal) should be true")
	}
	// 木不克金
	if IsKeWo(Wood, Wood) {
		t.Error("IsKeWo(Wood, Wood) should be false")
	}
}

func TestIsWoKe(t *testing.T) {
	// 木克土
	if !IsWoKe(Wood, Earth) {
		t.Error("IsWoKe(Wood, Earth) should be true")
	}
	// 木不克火
	if IsWoKe(Wood, Fire) {
		t.Error("IsWoKe(Wood, Fire) should be false")
	}
}

func TestGetXiangShengChain(t *testing.T) {
	chain := GetXiangShengChain()
	if len(chain) != 5 {
		t.Errorf("GetXiangShengChain() should return 5 elements, got %d", len(chain))
	}
	expected := []WuXing{Wood, Fire, Earth, Metal, Water}
	for i, wx := range chain {
		if wx != expected[i] {
			t.Errorf("GetXiangShengChain()[%d] = %d, expected %d", i, wx, expected[i])
		}
	}
}

func TestGetAllWuXing(t *testing.T) {
	all := GetAllWuXing()
	if len(all) != 5 {
		t.Errorf("GetAllWuXing() should return 5 elements, got %d", len(all))
	}
}

func TestCalculateWuXingScore(t *testing.T) {
	// 测试八字: 壬戌 壬寅 庚午 丙戌
	bazi := &Bazi{
		Year:  GanZhi{Gan: Ren, Zhi: Xu},
		Month: GanZhi{Gan: Ren, Zhi: YinZhi},
		Day:   GanZhi{Gan: Geng, Zhi: WuZhi},
		Hour:  GanZhi{Gan: Bing, Zhi: Xu},
	}

	scores := CalculateWuXingScore(bazi)
	if len(scores) == 0 {
		t.Error("CalculateWuXingScore should return non-empty result")
	}

	// 验证返回的分数都是正数
	for _, ws := range scores {
		if ws.Score <= 0 {
			t.Errorf("WuXing %s score should be positive, got %d", ws.WuXing.String(), ws.Score)
		}
	}
}

func TestWuXingRelationString(t *testing.T) {
	if Sheng.String() != "生" {
		t.Errorf("Sheng.String() = %s, expected 生", Sheng.String())
	}
	if Ke.String() != "克" {
		t.Errorf("Ke.String() = %s, expected 克", Ke.String())
	}
	if Tong.String() != "同" {
		t.Errorf("Tong.String() = %s, expected 同", Tong.String())
	}
}
