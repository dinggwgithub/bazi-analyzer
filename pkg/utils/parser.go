package utils

import (
	"errors"
	"strings"

	"bazi-analyzer/internal/model/base"
)

func ParseBaZi(input string) (base.BaZiChart, error) {
	input = strings.TrimSpace(input)
	parts := strings.Fields(input)

	if len(parts) != 4 {
		return base.BaZiChart{}, errors.New("输入格式错误，需要包含年柱、月柱、日柱、时柱四个部分")
	}

	yearPillar, err := parsePillar(parts[0], base.Year)
	if err != nil {
		return base.BaZiChart{}, err
	}

	monthPillar, err := parsePillar(parts[1], base.Month)
	if err != nil {
		return base.BaZiChart{}, err
	}

	dayPillar, err := parsePillar(parts[2], base.Day)
	if err != nil {
		return base.BaZiChart{}, err
	}

	hourPillar, err := parsePillar(parts[3], base.Hour)
	if err != nil {
		return base.BaZiChart{}, err
	}

	return base.BaZiChart{
		YearPillar:  yearPillar,
		MonthPillar: monthPillar,
		DayPillar:   dayPillar,
		HourPillar:  hourPillar,
		OriginalStr: input,
	}, nil
}

func parsePillar(s string, position base.PillarPosition) (base.Pillar, error) {
	runes := []rune(s)
	if len(runes) != 2 {
		return base.Pillar{}, errors.New("柱格式错误：" + s + "，每个柱必须包含两个字")
	}

	tianGanName := string(runes[0])
	diZhiName := string(runes[1])

	tianGan, ok := base.GetTianGan(tianGanName)
	if !ok {
		return base.Pillar{}, errors.New("无效天干：" + tianGanName)
	}

	diZhi, ok := base.GetDiZhi(diZhiName)
	if !ok {
		return base.Pillar{}, errors.New("无效地支：" + diZhiName)
	}

	return base.Pillar{
		TianGan:  tianGan,
		DiZhi:    diZhi,
		Position: position,
	}, nil
}
