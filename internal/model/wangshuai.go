package model

import "bazi-analyzer/pkg/wuxing"

type WangshuaiType int

const (
	WangshuaiShenwang WangshuaiType = iota
	WangshuaiShenruo
	WangshuaiCongqiang
	WangshuaiCongruo
)

func (w WangshuaiType) String() string {
	names := []string{"身旺", "身弱", "从强", "从弱"}
	if w >= 0 && int(w) < len(names) {
		return names[w]
	}
	return "未知"
}

type WangshuaiResult struct {
	Type        WangshuaiType
	DelingScore int
	DidiScore   int
	DeshiScore  int
	TotalScore  int
	Deling      DelingAnalysis
	Didi        DidiAnalysis
	Deshi       DeshiAnalysis
	Kexieha     KexiehaAnalysis
}

type DelingAnalysis struct {
	IsDeling    bool
	MonthWuxing wuxing.Wuxing
	RiZhuWuxing wuxing.Wuxing
	Relation    string
	Score       int
}

type DidiAnalysis struct {
	HasRoot     bool
	RootType    string
	RootDizhis  []string
	Score       int
}

type DeshiAnalysis struct {
	BijieCount  int
	YinCount    int
	Score       int
}

type KexiehaAnalysis struct {
	GuanShaScore int
	ShiShangScore int
	CaiScore     int
	TotalScore   int
}
