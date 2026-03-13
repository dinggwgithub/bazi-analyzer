package parser

import (
	"errors"
	"strings"

	"bazi-analyzer/internal/model"
	"bazi-analyzer/pkg/dizhi"
	"bazi-analyzer/pkg/tiangan"
)

type Parser struct{}

func NewParser() *Parser {
	return &Parser{}
}

func (p *Parser) Parse(input string) (*model.Bazi, error) {
	input = strings.TrimSpace(input)
	if input == "" {
		return nil, errors.New("输入不能为空")
	}

	parts := strings.Fields(input)
	if len(parts) != 4 {
		return nil, errors.New("请输入完整的四柱八字，格式如：壬戌 壬寅 庚午 丙戌")
	}

	pillars := make([]model.Pillar, 4)
	positions := []model.Position{
		model.PositionNian,
		model.PositionYue,
		model.PositionRi,
		model.PositionShi,
	}

	for i, part := range parts {
		runes := []rune(part)
		if len(runes) != 2 {
			return nil, errors.New("每柱必须包含一个天干和一个地支，当前：" + part)
		}

		tgChar := string(runes[0])
		dzChar := string(runes[1])

		tg := tiangan.ParseTiangan(tgChar)
		dz := dizhi.ParseDizhi(dzChar)

		if tg == 0 && tgChar != "甲" {
			return nil, errors.New("无效的天干：" + tgChar)
		}
		if dz == 0 && dzChar != "子" {
			return nil, errors.New("无效的地支：" + dzChar)
		}

		pillars[i] = model.NewPillar(tg, int(dz), positions[i])
	}

	return model.NewBazi(pillars[0], pillars[1], pillars[2], pillars[3]), nil
}

func (p *Parser) ParseWithValidation(input string) (*model.Bazi, error) {
	bazi, err := p.Parse(input)
	if err != nil {
		return nil, err
	}

	if err := p.validateBazi(bazi); err != nil {
		return nil, err
	}

	return bazi, nil
}

func (p *Parser) validateBazi(bazi *model.Bazi) error {
	yuezhi := bazi.GetYuezhi()
	yueTiangan := bazi.YuePillar.Tiangan

	if !isValidMonthTiangan(yuezhi, yueTiangan) {
		return errors.New("月柱天干与地支不匹配")
	}

	return nil
}

func isValidMonthTiangan(yuezhi dizhi.Dizhi, yueTiangan tiangan.Tiangan) bool {
	return true
}
