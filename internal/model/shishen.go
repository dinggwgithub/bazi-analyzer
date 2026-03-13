package model

import "bazi-analyzer/pkg/wuxing"

type ShishenType int

const (
	ShishenBijian ShishenType = iota
	ShishenJiecai
	ShishenShishen
	ShishenShangguan
	ShishenPiancai
	ShishenZhengcai
	ShishenQisha
	ShishenZhengguan
	ShishenPianyin
	ShishenZhengyin
)

func (s ShishenType) String() string {
	names := []string{
		"比肩", "劫财", "食神", "伤官",
		"偏财", "正财", "七杀", "正官",
		"偏印", "正印",
	}
	if s >= 0 && int(s) < len(names) {
		return names[s]
	}
	return "未知"
}

func GetShishen(riZhuWuxing, targetWuxing wuxing.Wuxing, isSameYinyang bool) ShishenType {
	if riZhuWuxing == targetWuxing {
		if isSameYinyang {
			return ShishenBijian
		}
		return ShishenJiecai
	}
	
	relation := riZhuWuxing.Relation(targetWuxing)
	
	switch relation {
	case wuxing.RelationSheng:
		if isSameYinyang {
			return ShishenShishen
		}
		return ShishenShangguan
	case wuxing.RelationKe:
		if isSameYinyang {
			return ShishenPiancai
		}
		return ShishenZhengcai
	case wuxing.RelationBeiKe:
		if isSameYinyang {
			return ShishenQisha
		}
		return ShishenZhengguan
	case wuxing.RelationBeiSheng:
		if isSameYinyang {
			return ShishenPianyin
		}
		return ShishenZhengyin
	default:
		return ShishenBijian
	}
}

func (s ShishenType) IsBiJie() bool {
	return s == ShishenBijian || s == ShishenJiecai
}

func (s ShishenType) IsShiShang() bool {
	return s == ShishenShishen || s == ShishenShangguan
}

func (s ShishenType) IsCai() bool {
	return s == ShishenPiancai || s == ShishenZhengcai
}

func (s ShishenType) IsGuanSha() bool {
	return s == ShishenQisha || s == ShishenZhengguan
}

func (s ShishenType) IsYin() bool {
	return s == ShishenPianyin || s == ShishenZhengyin
}

func (s ShishenType) IsHelpful() bool {
	return s.IsBiJie() || s.IsYin()
}

func (s ShishenType) IsDraining() bool {
	return s.IsShiShang() || s.IsCai() || s.IsGuanSha()
}
