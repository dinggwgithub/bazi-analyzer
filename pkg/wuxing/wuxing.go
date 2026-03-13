package wuxing

type Wuxing int

const (
	Mu Wuxing = iota
	Huo
	Tu
	Jin
	Shui
)

func (w Wuxing) String() string {
	names := []string{"木", "火", "土", "金", "水"}
	if w >= 0 && int(w) < len(names) {
		return names[w]
	}
	return "未知"
}

func (w Wuxing) Sheng() Wuxing {
	shengMap := map[Wuxing]Wuxing{
		Mu:   Huo,
		Huo:  Tu,
		Tu:   Jin,
		Jin:  Shui,
		Shui: Mu,
	}
	return shengMap[w]
}

func (w Wuxing) Ke() Wuxing {
	keMap := map[Wuxing]Wuxing{
		Mu:   Tu,
		Tu:   Shui,
		Shui: Huo,
		Huo:  Jin,
		Jin:  Mu,
	}
	return keMap[w]
}

func (w Wuxing) BeiSheng() Wuxing {
	beiShengMap := map[Wuxing]Wuxing{
		Mu:   Shui,
		Shui: Jin,
		Jin:  Tu,
		Tu:   Huo,
		Huo:  Mu,
	}
	return beiShengMap[w]
}

func (w Wuxing) BeiKe() Wuxing {
	beiKeMap := map[Wuxing]Wuxing{
		Mu:   Jin,
		Jin:  Huo,
		Huo:  Shui,
		Shui: Tu,
		Tu:   Mu,
	}
	return beiKeMap[w]
}

func (w Wuxing) IsSheng(other Wuxing) bool {
	return w.Sheng() == other
}

func (w Wuxing) IsKe(other Wuxing) bool {
	return w.Ke() == other
}

func (w Wuxing) IsBeiSheng(other Wuxing) bool {
	return w.BeiSheng() == other
}

func (w Wuxing) IsBeiKe(other Wuxing) bool {
	return w.BeiKe() == other
}

func (w Wuxing) Relation(other Wuxing) RelationType {
	if w == other {
		return RelationSame
	}
	if w.IsSheng(other) {
		return RelationSheng
	}
	if w.IsKe(other) {
		return RelationKe
	}
	if w.IsBeiSheng(other) {
		return RelationBeiSheng
	}
	if w.IsBeiKe(other) {
		return RelationBeiKe
	}
	return RelationNone
}

type RelationType int

const (
	RelationNone RelationType = iota
	RelationSame
	RelationSheng
	RelationKe
	RelationBeiSheng
	RelationBeiKe
)

func (r RelationType) String() string {
	names := []string{"无关系", "同类", "相生", "相克", "被生", "被克"}
	if r >= 0 && int(r) < len(names) {
		return names[r]
	}
	return "未知"
}
