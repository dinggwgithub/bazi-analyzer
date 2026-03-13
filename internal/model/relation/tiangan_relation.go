package relation

import (
	"bazi-analyzer/internal/model/base"
)

type RelationType string

const (
	WuHe    RelationType = "五合"
	SiChong RelationType = "四冲"
)

type TianGanRelation struct {
	Type     RelationType
	Elements []base.TianGan
	Result   base.WuXing
	Priority int
}

var TianGanWuHeMap = map[string]TianGanRelation{
	"甲己": {Type: WuHe, Elements: []base.TianGan{base.Jia, base.Ji}, Result: base.Earth, Priority: 1},
	"乙庚": {Type: WuHe, Elements: []base.TianGan{base.Yi, base.Geng}, Result: base.Metal, Priority: 1},
	"丙辛": {Type: WuHe, Elements: []base.TianGan{base.Bing, base.Xin}, Result: base.Water, Priority: 1},
	"丁壬": {Type: WuHe, Elements: []base.TianGan{base.Ding, base.Ren}, Result: base.Wood, Priority: 1},
	"戊癸": {Type: WuHe, Elements: []base.TianGan{base.Wu, base.Gui}, Result: base.Fire, Priority: 1},
}

var TianGanSiChongMap = map[string]TianGanRelation{
	"甲庚": {Type: SiChong, Elements: []base.TianGan{base.Jia, base.Geng}, Result: "", Priority: 2},
	"乙辛": {Type: SiChong, Elements: []base.TianGan{base.Yi, base.Xin}, Result: "", Priority: 2},
	"丙壬": {Type: SiChong, Elements: []base.TianGan{base.Bing, base.Ren}, Result: "", Priority: 2},
	"丁癸": {Type: SiChong, Elements: []base.TianGan{base.Ding, base.Gui}, Result: "", Priority: 2},
}

func GetTianGanWuHe(tg1, tg2 base.TianGan) (TianGanRelation, bool) {
	key1 := tg1.Name + tg2.Name
	key2 := tg2.Name + tg1.Name
	if rel, ok := TianGanWuHeMap[key1]; ok {
		return rel, true
	}
	if rel, ok := TianGanWuHeMap[key2]; ok {
		return rel, true
	}
	return TianGanRelation{}, false
}

func GetTianGanSiChong(tg1, tg2 base.TianGan) (TianGanRelation, bool) {
	key1 := tg1.Name + tg2.Name
	key2 := tg2.Name + tg1.Name
	if rel, ok := TianGanSiChongMap[key1]; ok {
		return rel, true
	}
	if rel, ok := TianGanSiChongMap[key2]; ok {
		return rel, true
	}
	return TianGanRelation{}, false
}
