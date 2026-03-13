package model

import (
	"bazi-analyzer/pkg/dizhi"
	"bazi-analyzer/pkg/tiangan"
	"bazi-analyzer/pkg/wuxing"
)

type RelationType int

const (
	RelationWuhe RelationType = iota
	RelationSichong
	RelationSanhui
	RelationSanhe
	RelationLiuhe
	RelationChong
	RelationXing
	RelationHai
	RelationPo
	RelationSheng
	RelationKe
)

func (r RelationType) String() string {
	names := []string{
		"天干五合", "天干四冲", "地支三会", "地支三合",
		"地支六合", "地支六冲", "地支三刑", "地支六害",
		"地支六破", "相生", "相克",
	}
	if r >= 0 && int(r) < len(names) {
		return names[r]
	}
	return "未知"
}

func (r RelationType) Priority() int {
	priorities := map[RelationType]int{
		RelationSanhui:   1,
		RelationSanhe:    2,
		RelationWuhe:     3,
		RelationLiuhe:    4,
		RelationChong:    5,
		RelationXing:     6,
		RelationHai:      7,
		RelationPo:       8,
		RelationSichong:  9,
		RelationSheng:    10,
		RelationKe:       11,
	}
	return priorities[r]
}

type Relation struct {
	Type         RelationType
	Participants []interface{}
	ResultWuxing wuxing.Wuxing
	Success      bool
	Description  string
}

type TianganRelation struct {
	Type        RelationType
	Tiangan1    tiangan.Tiangan
	Tiangan2    tiangan.Tiangan
	ResultWuxing wuxing.Wuxing
	Success     bool
	Position1   Position
	Position2   Position
}

type DizhiRelation struct {
	Type         RelationType
	Dizhis       []dizhi.Dizhi
	ResultWuxing wuxing.Wuxing
	Success      bool
	Positions    []Position
}
