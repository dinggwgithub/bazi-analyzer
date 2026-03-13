package model

import (
	"bazi-analyzer/pkg/dizhi"
	"bazi-analyzer/pkg/tiangan"
	"bazi-analyzer/pkg/wuxing"
)

type Bazi struct {
	NianPillar Pillar
	YuePillar  Pillar
	RiPillar   Pillar
	ShiPillar  Pillar
}

func NewBazi(nian, yue, ri, shi Pillar) *Bazi {
	return &Bazi{
		NianPillar: nian,
		YuePillar:  yue,
		RiPillar:   ri,
		ShiPillar:  shi,
	}
}

func (b *Bazi) GetPillars() []Pillar {
	return []Pillar{b.NianPillar, b.YuePillar, b.RiPillar, b.ShiPillar}
}

func (b *Bazi) GetTiangans() []tiangan.Tiangan {
	return []tiangan.Tiangan{
		b.NianPillar.Tiangan,
		b.YuePillar.Tiangan,
		b.RiPillar.Tiangan,
		b.ShiPillar.Tiangan,
	}
}

func (b *Bazi) GetDizhis() []dizhi.Dizhi {
	return []dizhi.Dizhi{
		dizhi.Dizhi(b.NianPillar.Dizhi.Original),
		dizhi.Dizhi(b.YuePillar.Dizhi.Original),
		dizhi.Dizhi(b.RiPillar.Dizhi.Original),
		dizhi.Dizhi(b.ShiPillar.Dizhi.Original),
	}
}

func (b *Bazi) GetRizhu() Pillar {
	return b.RiPillar
}

func (b *Bazi) GetRizhuTiangan() tiangan.Tiangan {
	return b.RiPillar.Tiangan
}

func (b *Bazi) GetYuezhi() dizhi.Dizhi {
	return dizhi.Dizhi(b.YuePillar.Dizhi.Original)
}

func (b *Bazi) GetMonthWuxing() wuxing.Wuxing {
	return b.GetYuezhi().Wuxing()
}

func (b *Bazi) CountWuxing() map[wuxing.Wuxing]int {
	count := make(map[wuxing.Wuxing]int)
	for _, t := range b.GetTiangans() {
		count[t.Wuxing()]++
	}
	for _, d := range b.GetDizhis() {
		count[d.Wuxing()]++
		for _, c := range d.GetCanggan() {
			wx := getWuxingFromTianganName(c.TianganName)
			count[wx]++
		}
	}
	return count
}

func getWuxingFromTianganName(name string) wuxing.Wuxing {
	m := map[string]wuxing.Wuxing{
		"甲": wuxing.Mu, "乙": wuxing.Mu,
		"丙": wuxing.Huo, "丁": wuxing.Huo,
		"戊": wuxing.Tu, "己": wuxing.Tu,
		"庚": wuxing.Jin, "辛": wuxing.Jin,
		"壬": wuxing.Shui, "癸": wuxing.Shui,
	}
	return m[name]
}
