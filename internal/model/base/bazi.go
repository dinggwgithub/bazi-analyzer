package base

type PillarPosition string

const (
	Year  PillarPosition = "年"
	Month PillarPosition = "月"
	Day   PillarPosition = "日"
	Hour  PillarPosition = "时"
)

type Pillar struct {
	TianGan  TianGan
	DiZhi    DiZhi
	Position PillarPosition
}

func (p Pillar) String() string {
	return p.TianGan.Name + p.DiZhi.Name
}

type BaZiChart struct {
	YearPillar  Pillar
	MonthPillar Pillar
	DayPillar   Pillar
	HourPillar  Pillar
	OriginalStr string
}

func (b BaZiChart) String() string {
	return b.YearPillar.String() + " " +
		b.MonthPillar.String() + " " +
		b.DayPillar.String() + " " +
		b.HourPillar.String()
}

func (b BaZiChart) GetDayMaster() TianGan {
	return b.DayPillar.TianGan
}

func (b BaZiChart) GetMonthLing() DiZhi {
	return b.MonthPillar.DiZhi
}

func (b BaZiChart) GetAllTianGan() []TianGan {
	return []TianGan{
		b.YearPillar.TianGan,
		b.MonthPillar.TianGan,
		b.DayPillar.TianGan,
		b.HourPillar.TianGan,
	}
}

func (b BaZiChart) GetAllDiZhi() []DiZhi {
	return []DiZhi{
		b.YearPillar.DiZhi,
		b.MonthPillar.DiZhi,
		b.DayPillar.DiZhi,
		b.HourPillar.DiZhi,
	}
}

func (b BaZiChart) GetPillars() []Pillar {
	return []Pillar{
		b.YearPillar,
		b.MonthPillar,
		b.DayPillar,
		b.HourPillar,
	}
}
