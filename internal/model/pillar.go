package model

import (
	"bazi-analyzer/pkg/dizhi"
	"bazi-analyzer/pkg/tiangan"
)

type Position int

const (
	PositionNian Position = iota
	PositionYue
	PositionRi
	PositionShi
)

func (p Position) String() string {
	names := []string{"年柱", "月柱", "日柱", "时柱"}
	if p >= 0 && int(p) < len(names) {
		return names[p]
	}
	return "未知"
}

type Pillar struct {
	Tiangan  tiangan.Tiangan
	Dizhi    DizhiWrapper
	Position Position
}

type DizhiWrapper struct {
	Original      DizhiType
	Transformed   DizhiType
	IsTransformed bool
}

type DizhiType dizhi.Dizhi

func (d DizhiType) String() string {
	return dizhi.Dizhi(d).String()
}

func (d DizhiType) ToDizhi() dizhi.Dizhi {
	return dizhi.Dizhi(d)
}

func NewPillar(t tiangan.Tiangan, d int, pos Position) Pillar {
	return Pillar{
		Tiangan: t,
		Dizhi: DizhiWrapper{
			Original:      DizhiType(d),
			Transformed:   DizhiType(d),
			IsTransformed: false,
		},
		Position: pos,
	}
}
