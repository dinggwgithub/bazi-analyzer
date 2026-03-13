package relation

import (
	"bazi-analyzer/internal/model/base"
)

type DiZhiRelationType string

const (
	SanHui  DiZhiRelationType = "三会"
	SanHe   DiZhiRelationType = "三合"
	LiuHe   DiZhiRelationType = "六合"
	LiuChong DiZhiRelationType = "六冲"
)

type DiZhiRelation struct {
	Type     DiZhiRelationType
	Elements []base.DiZhi
	Result   base.WuXing
	Priority int
}

var SanHuiList = []DiZhiRelation{
	{Type: SanHui, Elements: []base.DiZhi{base.Yin, base.Mao, base.Chen}, Result: base.Wood, Priority: 1},
	{Type: SanHui, Elements: []base.DiZhi{base.Si, base.WuDiZhi, base.Wei}, Result: base.Fire, Priority: 1},
	{Type: SanHui, Elements: []base.DiZhi{base.Shen, base.You, base.Xu}, Result: base.Metal, Priority: 1},
	{Type: SanHui, Elements: []base.DiZhi{base.Hai, base.Zi, base.Chou}, Result: base.Water, Priority: 1},
}

var SanHeList = []DiZhiRelation{
	{Type: SanHe, Elements: []base.DiZhi{base.Shen, base.Zi, base.Chen}, Result: base.Water, Priority: 2},
	{Type: SanHe, Elements: []base.DiZhi{base.Hai, base.Mao, base.Wei}, Result: base.Wood, Priority: 2},
	{Type: SanHe, Elements: []base.DiZhi{base.Yin, base.WuDiZhi, base.Xu}, Result: base.Fire, Priority: 2},
	{Type: SanHe, Elements: []base.DiZhi{base.Si, base.You, base.Chou}, Result: base.Metal, Priority: 2},
}

var LiuHeMap = map[string]DiZhiRelation{
	"子丑": {Type: LiuHe, Elements: []base.DiZhi{base.Zi, base.Chou}, Result: "", Priority: 3},
	"寅亥": {Type: LiuHe, Elements: []base.DiZhi{base.Yin, base.Hai}, Result: "", Priority: 3},
	"卯戌": {Type: LiuHe, Elements: []base.DiZhi{base.Mao, base.Xu}, Result: "", Priority: 3},
	"辰酉": {Type: LiuHe, Elements: []base.DiZhi{base.Chen, base.You}, Result: "", Priority: 3},
	"巳申": {Type: LiuHe, Elements: []base.DiZhi{base.Si, base.Shen}, Result: "", Priority: 3},
	"午未": {Type: LiuHe, Elements: []base.DiZhi{base.WuDiZhi, base.Wei}, Result: "", Priority: 3},
}

func GetDiZhiLiuHe(dz1, dz2 base.DiZhi) (DiZhiRelation, bool) {
	key1 := dz1.Name + dz2.Name
	key2 := dz2.Name + dz1.Name
	if rel, ok := LiuHeMap[key1]; ok {
		return rel, true
	}
	if rel, ok := LiuHeMap[key2]; ok {
		return rel, true
	}
	return DiZhiRelation{}, false
}

func ContainsDiZhi(list []base.DiZhi, dz base.DiZhi) bool {
	for _, d := range list {
		if d.Name == dz.Name {
			return true
		}
	}
	return false
}
