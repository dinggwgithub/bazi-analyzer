package engine

import (
	"bazi-analyzer/internal/analyzer"
	"bazi-analyzer/internal/model"
	"bazi-analyzer/internal/parser"
	"bazi-analyzer/internal/reporter"
)

type Engine struct {
	parser        *parser.Parser
	scanner       *analyzer.RelationScanner
	reconstructor *analyzer.PatternReconstructor
	wangshuai     *analyzer.WangshuaiAnalyzer
	xiyong        *analyzer.XiyongAnalyzer
	Reporter      *reporter.Reporter
}

func NewEngine() *Engine {
	return &Engine{
		parser:        parser.NewParser(),
		scanner:       analyzer.NewRelationScanner(),
		reconstructor: analyzer.NewPatternReconstructor(),
		wangshuai:     analyzer.NewWangshuaiAnalyzer(),
		xiyong:        analyzer.NewXiyongAnalyzer(),
		Reporter:      reporter.NewReporter(),
	}
}

type AnalysisResult struct {
	Bazi              *model.Bazi
	ScanResult        *analyzer.ScanResult
	ReconstructResult *analyzer.ReconstructResult
	WangshuaiResult   *model.WangshuaiResult
	XiyongResult      *model.XiyongResult
	Report            *reporter.AnalysisReport
}

func (e *Engine) Analyze(input string) (*AnalysisResult, error) {
	bazi, err := e.parser.Parse(input)
	if err != nil {
		return nil, err
	}

	return e.AnalyzeBazi(bazi), nil
}

func (e *Engine) AnalyzeBazi(bazi *model.Bazi) *AnalysisResult {
	result := &AnalysisResult{
		Bazi: bazi,
	}

	result.ScanResult = e.scanner.Scan(bazi)

	result.ReconstructResult = e.reconstructor.Reconstruct(bazi, result.ScanResult)

	result.WangshuaiResult = e.wangshuai.Analyze(bazi, result.ReconstructResult)

	result.XiyongResult = e.xiyong.Analyze(result.WangshuaiResult, bazi)

	result.Report = e.Reporter.Generate(
		bazi,
		result.ScanResult,
		result.ReconstructResult,
		result.WangshuaiResult,
		result.XiyongResult,
	)

	return result
}

func (e *Engine) AnalyzeAndFormat(input string) (string, error) {
	result, err := e.Analyze(input)
	if err != nil {
		return "", err
	}

	return e.Reporter.FormatReport(result.Report), nil
}
