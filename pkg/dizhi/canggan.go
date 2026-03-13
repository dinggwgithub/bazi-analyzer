package dizhi

import "bazi-analyzer/pkg/wuxing"

type Canggan struct {
	TianganName string
	Strength    int
}

type DizhiCanggan struct {
	Dizhi   Dizhi
	Canggan []Canggan
}

var CangganData = map[Dizhi][]Canggan{
	Zi:  {{TianganName: "癸", Strength: 1}},
	Chou: {{TianganName: "己", Strength: 1}, {TianganName: "癸", Strength: 2}, {TianganName: "辛", Strength: 3}},
	Yin:  {{TianganName: "甲", Strength: 1}, {TianganName: "丙", Strength: 2}, {TianganName: "戊", Strength: 3}},
	Mao:  {{TianganName: "乙", Strength: 1}},
	Chen: {{TianganName: "戊", Strength: 1}, {TianganName: "乙", Strength: 2}, {TianganName: "癸", Strength: 3}},
	Si:   {{TianganName: "丙", Strength: 1}, {TianganName: "戊", Strength: 2}, {TianganName: "庚", Strength: 3}},
	Wu:   {{TianganName: "丁", Strength: 1}, {TianganName: "己", Strength: 2}},
	Wei:  {{TianganName: "己", Strength: 1}, {TianganName: "丁", Strength: 2}, {TianganName: "乙", Strength: 3}},
	Shen: {{TianganName: "庚", Strength: 1}, {TianganName: "壬", Strength: 2}, {TianganName: "戊", Strength: 3}},
	You:  {{TianganName: "辛", Strength: 1}},
	Xu:   {{TianganName: "戊", Strength: 1}, {TianganName: "辛", Strength: 2}, {TianganName: "丁", Strength: 3}},
	Hai:  {{TianganName: "壬", Strength: 1}, {TianganName: "甲", Strength: 2}},
}

func (d Dizhi) GetCanggan() []Canggan {
	return CangganData[d]
}

func (d Dizhi) GetBenqi() *Canggan {
	canggan := CangganData[d]
	if len(canggan) > 0 {
		return &canggan[0]
	}
	return nil
}

func (d Dizhi) ContainsTiangan(tianganName string) bool {
	for _, c := range CangganData[d] {
		if c.TianganName == tianganName {
			return true
		}
	}
	return false
}

func (d Dizhi) HasWuxing(wx wuxing.Wuxing) bool {
	for _, c := range CangganData[d] {
		tianganMap := map[string]wuxing.Wuxing{
			"甲": wuxing.Mu, "乙": wuxing.Mu,
			"丙": wuxing.Huo, "丁": wuxing.Huo,
			"戊": wuxing.Tu, "己": wuxing.Tu,
			"庚": wuxing.Jin, "辛": wuxing.Jin,
			"壬": wuxing.Shui, "癸": wuxing.Shui,
		}
		if tianganMap[c.TianganName] == wx {
			return true
		}
	}
	return false
}
