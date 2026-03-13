package relation

import (
	"bazi-analyzer/internal/model/base"
)

type SpecialRelationType string

const (
	Xing SpecialRelationType = "刑"
	Chong SpecialRelationType = "冲"
	Hai  SpecialRelationType = "害"
	Po   SpecialRelationType = "破"
)

type SpecialRelation struct {
	Type     SpecialRelationType
	Elements []base.DiZhi
	Priority int
}

var LiuChongMap = map[string]SpecialRelation{
	"子午": {Type: Chong, Elements: []base.DiZhi{base.Zi, base.WuDiZhi}, Priority: 4},
	"丑未": {Type: Chong, Elements: []base.DiZhi{base.Chou, base.Wei}, Priority: 4},
	"寅申": {Type: Chong, Elements: []base.DiZhi{base.Yin, base.Shen}, Priority: 4},
	"卯酉": {Type: Chong, Elements: []base.DiZhi{base.Mao, base.You}, Priority: 4},
	"辰戌": {Type: Chong, Elements: []base.DiZhi{base.Chen, base.Xu}, Priority: 4},
	"巳亥": {Type: Chong, Elements: []base.DiZhi{base.Si, base.Hai}, Priority: 4},
}

var LiuHaiMap = map[string]SpecialRelation{
	"子未": {Type: Hai, Elements: []base.DiZhi{base.Zi, base.Wei}, Priority: 5},
	"丑午": {Type: Hai, Elements: []base.DiZhi{base.Chou, base.WuDiZhi}, Priority: 5},
	"寅巳": {Type: Hai, Elements: []base.DiZhi{base.Yin, base.Si}, Priority: 5},
	"卯辰": {Type: Hai, Elements: []base.DiZhi{base.Mao, base.Chen}, Priority: 5},
	"申亥": {Type: Hai, Elements: []base.DiZhi{base.Shen, base.Hai}, Priority: 5},
	"酉戌": {Type: Hai, Elements: []base.DiZhi{base.You, base.Xu}, Priority: 5},
}

func GetDiZhiLiuChong(dz1, dz2 base.DiZhi) (SpecialRelation, bool) {
	key1 := dz1.Name + dz2.Name
	key2 := dz2.Name + dz1.Name
	if rel, ok := LiuChongMap[key1]; ok {
		return rel, true
	}
	if rel, ok := LiuChongMap[key2]; ok {
		return rel, true
	}
	return SpecialRelation{}, false
}

func GetDiZhiLiuHai(dz1, dz2 base.DiZhi) (SpecialRelation, bool) {
	key1 := dz1.Name + dz2.Name
	key2 := dz2.Name + dz1.Name
	if rel, ok := LiuHaiMap[key1]; ok {
		return rel, true
	}
	if rel, ok := LiuHaiMap[key2]; ok {
		return rel, true
	}
	return SpecialRelation{}, false
}
