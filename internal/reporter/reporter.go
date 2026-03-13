package reporter

import (
	"fmt"
	"strings"

	"bazi-analyzer/internal/analyzer"
	"bazi-analyzer/internal/model"
)

type Reporter struct{}

func NewReporter() *Reporter {
	return &Reporter{}
}

type AnalysisReport struct {
	BaziInfo    BaziInfo
	Relations   RelationsInfo
	Reconstruct ReconstructInfo
	Wangshuai   WangshuaiInfo
	Xiyong      XiyongInfo
	Summary     string
}

type BaziInfo struct {
	NianPillar  string
	YuePillar   string
	RiPillar    string
	ShiPillar   string
	RiZhu       string
	RiZhuWuxing string
	WuxingCount map[string]int
}

type RelationsInfo struct {
	TianganRelations []string
	DizhiRelations   []string
	MajorRelation    string
}

type ReconstructInfo struct {
	HasTransformation bool
	Transformations   []string
	Description       string
}

type WangshuaiInfo struct {
	Type       string
	Deling     string
	Didi       string
	Deshi      string
	Kexieha    string
	TotalScore int
}

type XiyongInfo struct {
	Xishen     []string
	Yongshen   string
	Jishen     []string
	XishenDesc string
	JishenDesc string
}

func (r *Reporter) Generate(
	bazi *model.Bazi,
	scanResult *analyzer.ScanResult,
	reconstructResult *analyzer.ReconstructResult,
	wangshuaiResult *model.WangshuaiResult,
	xiyongResult *model.XiyongResult,
) *AnalysisReport {
	report := &AnalysisReport{}

	report.BaziInfo = r.generateBaziInfo(bazi)
	report.Relations = r.generateRelationsInfo(scanResult)
	report.Reconstruct = r.generateReconstructInfo(reconstructResult)
	report.Wangshuai = r.generateWangshuaiInfo(wangshuaiResult)
	report.Xiyong = r.generateXiyongInfo(xiyongResult)
	report.Summary = r.generateSummary(bazi, wangshuaiResult, xiyongResult, reconstructResult)

	return report
}

func (r *Reporter) generateBaziInfo(bazi *model.Bazi) BaziInfo {
	info := BaziInfo{
		NianPillar:  bazi.NianPillar.Tiangan.String() + bazi.NianPillar.Dizhi.Original.String(),
		YuePillar:   bazi.YuePillar.Tiangan.String() + bazi.YuePillar.Dizhi.Original.String(),
		RiPillar:    bazi.RiPillar.Tiangan.String() + bazi.RiPillar.Dizhi.Original.String(),
		ShiPillar:   bazi.ShiPillar.Tiangan.String() + bazi.ShiPillar.Dizhi.Original.String(),
		RiZhu:       bazi.GetRizhuTiangan().String(),
		RiZhuWuxing: bazi.GetRizhuTiangan().Wuxing().String(),
	}

	count := bazi.CountWuxing()
	info.WuxingCount = make(map[string]int)
	for wx, c := range count {
		info.WuxingCount[wx.String()] = c
	}

	return info
}

func (r *Reporter) generateRelationsInfo(scanResult *analyzer.ScanResult) RelationsInfo {
	info := RelationsInfo{}

	for _, rel := range scanResult.TianganRelations {
		var desc string
		if rel.Type == model.RelationWuhe {
			desc = fmt.Sprintf("%s%s合%s", rel.Tiangan1.String(), rel.Tiangan2.String(), rel.ResultWuxing.String())
			if rel.Success {
				desc += "（合化成功）"
			} else {
				desc += "（合而不化）"
			}
		} else if rel.Type == model.RelationSichong {
			desc = fmt.Sprintf("%s%s冲", rel.Tiangan1.String(), rel.Tiangan2.String())
		}
		info.TianganRelations = append(info.TianganRelations, desc)
	}

	for _, rel := range scanResult.DizhiRelations {
		var desc string
		dizhiStrs := make([]string, len(rel.Dizhis))
		for i, d := range rel.Dizhis {
			dizhiStrs[i] = d.String()
		}

		switch rel.Type {
		case model.RelationSanhui:
			desc = fmt.Sprintf("%s三会%s局", strings.Join(dizhiStrs, ""), rel.ResultWuxing.String())
			info.MajorRelation = desc
		case model.RelationSanhe:
			desc = fmt.Sprintf("%s三合%s局", strings.Join(dizhiStrs, ""), rel.ResultWuxing.String())
			info.MajorRelation = desc
		case model.RelationLiuhe:
			desc = fmt.Sprintf("%s六合%s", strings.Join(dizhiStrs, ""), rel.ResultWuxing.String())
		case model.RelationChong:
			desc = fmt.Sprintf("%s冲", strings.Join(dizhiStrs, ""))
		case model.RelationXing:
			desc = fmt.Sprintf("%s刑", strings.Join(dizhiStrs, ""))
		case model.RelationHai:
			desc = fmt.Sprintf("%s害", strings.Join(dizhiStrs, ""))
		case model.RelationPo:
			desc = fmt.Sprintf("%s破", strings.Join(dizhiStrs, ""))
		}
		info.DizhiRelations = append(info.DizhiRelations, desc)
	}

	return info
}

func (r *Reporter) generateReconstructInfo(reconstructResult *analyzer.ReconstructResult) ReconstructInfo {
	info := ReconstructInfo{}

	if reconstructResult == nil {
		return info
	}

	if len(reconstructResult.Transformations) > 0 {
		info.HasTransformation = true
		for _, t := range reconstructResult.Transformations {
			info.Transformations = append(info.Transformations,
				fmt.Sprintf("%s → %s（%s）", t.Original, t.Transformed, t.Reason))
		}
	}

	if reconstructResult.MajorRelation != nil {
		info.Description = fmt.Sprintf("命局存在%s，力量聚焦于%s五行",
			reconstructResult.MajorRelation.Type.String(),
			reconstructResult.MajorRelation.ResultWuxing.String())
	}

	return info
}

func (r *Reporter) generateWangshuaiInfo(wangshuaiResult *model.WangshuaiResult) WangshuaiInfo {
	info := WangshuaiInfo{
		Type:       wangshuaiResult.Type.String(),
		TotalScore: wangshuaiResult.TotalScore,
	}

	info.Deling = fmt.Sprintf("月令%s，日主%s，%s，得分：%d",
		wangshuaiResult.Deling.MonthWuxing.String(),
		wangshuaiResult.Deling.RiZhuWuxing.String(),
		wangshuaiResult.Deling.Relation,
		wangshuaiResult.Deling.Score)

	if wangshuaiResult.Didi.HasRoot {
		info.Didi = fmt.Sprintf("有根，根在：%s，得分：%d",
			strings.Join(wangshuaiResult.Didi.RootDizhis, "、"),
			wangshuaiResult.Didi.Score)
	} else {
		info.Didi = "无根，得分：0"
	}

	info.Deshi = fmt.Sprintf("天干比劫%d个、印星%d个，得分：%d",
		wangshuaiResult.Deshi.BijieCount,
		wangshuaiResult.Deshi.YinCount,
		wangshuaiResult.Deshi.Score)

	info.Kexieha = fmt.Sprintf("官杀%d分、食伤%d分、财星%d分，合计%d分",
		wangshuaiResult.Kexieha.GuanShaScore,
		wangshuaiResult.Kexieha.ShiShangScore,
		wangshuaiResult.Kexieha.CaiScore,
		wangshuaiResult.Kexieha.TotalScore)

	return info
}

func (r *Reporter) generateXiyongInfo(xiyongResult *model.XiyongResult) XiyongInfo {
	info := XiyongInfo{
		Yongshen:   xiyongResult.Yongshen.String(),
		XishenDesc: xiyongResult.XishenDesc,
		JishenDesc: xiyongResult.JishenDesc,
	}

	for _, wx := range xiyongResult.Xishen {
		info.Xishen = append(info.Xishen, wx.String())
	}

	for _, wx := range xiyongResult.Jishen {
		info.Jishen = append(info.Jishen, wx.String())
	}

	return info
}

func (r *Reporter) generateSummary(bazi *model.Bazi, wangshuaiResult *model.WangshuaiResult, xiyongResult *model.XiyongResult, reconstructResult *analyzer.ReconstructResult) string {
	var sb strings.Builder

	sb.WriteString(fmt.Sprintf("日主%s，%s。", bazi.GetRizhuTiangan().String(), wangshuaiResult.Type.String()))

	if reconstructResult != nil && reconstructResult.MajorRelation != nil {
		sb.WriteString(fmt.Sprintf("命局%s成立，", reconstructResult.MajorRelation.Type.String()))
		sb.WriteString(fmt.Sprintf("五行聚焦于%s，", reconstructResult.MajorRelation.ResultWuxing.String()))
	}

	xishenStrs := make([]string, len(xiyongResult.Xishen))
	for i, wx := range xiyongResult.Xishen {
		xishenStrs[i] = wx.String()
	}
	jishenStrs := make([]string, len(xiyongResult.Jishen))
	for i, wx := range xiyongResult.Jishen {
		jishenStrs[i] = wx.String()
	}

	sb.WriteString(fmt.Sprintf("喜%s，", strings.Join(xishenStrs, "、")))
	sb.WriteString(fmt.Sprintf("忌%s。", strings.Join(jishenStrs, "、")))

	sb.WriteString(xiyongResult.XishenDesc + "。")

	return sb.String()
}

func (r *Reporter) FormatReport(report *AnalysisReport) string {
	var sb strings.Builder

	sb.WriteString("\n")
	sb.WriteString("═══════════════════════════════════════════════════════════════\n")
	sb.WriteString("                        八 字 分 析 报 告                        \n")
	sb.WriteString("═══════════════════════════════════════════════════════════════\n")

	sb.WriteString("\n【命局基本信息】\n")
	sb.WriteString(fmt.Sprintf("  年柱：%s    月柱：%s    日柱：%s    时柱：%s\n",
		report.BaziInfo.NianPillar, report.BaziInfo.YuePillar,
		report.BaziInfo.RiPillar, report.BaziInfo.ShiPillar))
	sb.WriteString(fmt.Sprintf("  日主：%s（%s）\n", report.BaziInfo.RiZhu, report.BaziInfo.RiZhuWuxing))

	sb.WriteString("  五行统计：")
	for wx, count := range report.BaziInfo.WuxingCount {
		sb.WriteString(fmt.Sprintf("%s%d ", wx, count))
	}
	sb.WriteString("\n")

	sb.WriteString("\n【全局特殊关系】\n")
	if len(report.Relations.TianganRelations) > 0 {
		sb.WriteString("  天干关系：")
		sb.WriteString(strings.Join(report.Relations.TianganRelations, "、"))
		sb.WriteString("\n")
	}
	if len(report.Relations.DizhiRelations) > 0 {
		sb.WriteString("  地支关系：")
		sb.WriteString(strings.Join(report.Relations.DizhiRelations, "、"))
		sb.WriteString("\n")
	}
	if report.Relations.MajorRelation != "" {
		sb.WriteString(fmt.Sprintf("  核心关系：%s\n", report.Relations.MajorRelation))
	}

	if report.Reconstruct.HasTransformation {
		sb.WriteString("\n【格局重构】\n")
		for _, t := range report.Reconstruct.Transformations {
			sb.WriteString(fmt.Sprintf("  %s\n", t))
		}
		if report.Reconstruct.Description != "" {
			sb.WriteString(fmt.Sprintf("  %s\n", report.Reconstruct.Description))
		}
	}

	sb.WriteString("\n【日主旺衰分析】\n")
	sb.WriteString(fmt.Sprintf("  得令：%s\n", report.Wangshuai.Deling))
	sb.WriteString(fmt.Sprintf("  得地：%s\n", report.Wangshuai.Didi))
	sb.WriteString(fmt.Sprintf("  得势：%s\n", report.Wangshuai.Deshi))
	sb.WriteString(fmt.Sprintf("  克泄耗：%s\n", report.Wangshuai.Kexieha))
	sb.WriteString(fmt.Sprintf("  综合得分：%d\n", report.Wangshuai.TotalScore))
	sb.WriteString(fmt.Sprintf("  旺衰结论：%s\n", report.Wangshuai.Type))

	sb.WriteString("\n【喜忌用神】\n")
	sb.WriteString(fmt.Sprintf("  喜神：%s\n", strings.Join(report.Xiyong.Xishen, "、")))
	sb.WriteString(fmt.Sprintf("  用神：%s\n", report.Xiyong.Yongshen))
	sb.WriteString(fmt.Sprintf("  忌神：%s\n", strings.Join(report.Xiyong.Jishen, "、")))
	sb.WriteString(fmt.Sprintf("  %s\n", report.Xiyong.XishenDesc))
	sb.WriteString(fmt.Sprintf("  %s\n", report.Xiyong.JishenDesc))

	sb.WriteString("\n【综合结论】\n")
	sb.WriteString(fmt.Sprintf("  %s\n", report.Summary))

	sb.WriteString("\n═══════════════════════════════════════════════════════════════\n")

	return sb.String()
}

func (r *Reporter) FormatJSON(report *AnalysisReport) string {
	return fmt.Sprintf(`{
  "bazi": {
    "nian": "%s",
    "yue": "%s",
    "ri": "%s",
    "shi": "%s",
    "rizhu": "%s",
    "rizhu_wuxing": "%s"
  },
  "wangshuai": {
    "type": "%s",
    "score": %d
  },
  "xiyong": {
    "xishen": %v,
    "yongshen": "%s",
    "jishen": %v
  },
  "summary": "%s"
}`,
		report.BaziInfo.NianPillar, report.BaziInfo.YuePillar,
		report.BaziInfo.RiPillar, report.BaziInfo.ShiPillar,
		report.BaziInfo.RiZhu, report.BaziInfo.RiZhuWuxing,
		report.Wangshuai.Type, report.Wangshuai.TotalScore,
		report.Xiyong.Xishen, report.Xiyong.Yongshen, report.Xiyong.Jishen,
		report.Summary)
}
