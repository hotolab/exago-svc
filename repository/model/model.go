package model

// TestResults received from the test runner.
type TestResults struct {
	Checklist struct {
		Failed []struct {
			Category string `json:"Category"`
			Desc     string `json:"Desc"`
			Name     string `json:"Name"`
		} `json:"Failed"`
		Passed []struct {
			Category string `json:"Category"`
			Desc     string `json:"Desc"`
			Name     string `json:"Name"`
		} `json:"Passed"`
	} `json:"checklist"`
	Packages []struct {
		Coverage      float64 `json:"coverage"`
		ExecutionTime float64 `json:"execution_time"`
		Name          string  `json:"name"`
		Success       bool    `json:"success"`
		Tests         []struct {
			Name          string `json:"name"`
			ExecutionTime int    `json:"execution_time"`
			Passed        bool   `json:"passed"`
		} `json:"tests"`
	} `json:"packages"`
	ExecutionTime struct {
		Goprove string `json:"goprove"`
		Gotest  string `json:"gotest"`
	} `json:"execution_time"`
	RawOutput struct {
		Gotest string `json:"gotest"`
	} `json:"raw_output"`
}

type Imports []string
type CodeStats map[string]int
type LintMessages map[string]map[string][]map[string]interface{}

type RepositoryData interface {
	Name() string
}

func (t TestResults) Name() string {
	return "testresults"
}

func (i Imports) Name() string {
	return "imports"
}

func (c CodeStats) Name() string {
	return "codestats"
}

func (l LintMessages) Name() string {
	return "lintmessages"
}
