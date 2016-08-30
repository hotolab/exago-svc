package model

import "simonwaldherr.de/go/golibs/xmath"

const (
	ProjectRunnerName = "projectrunner"
	ThirdPartiesName  = "thirdparties"
	TestCoverageName  = "testcoverage"
	TestDurationName  = "testduration"
)

// ProjectRunner received from the test runner.
type ProjectRunner struct {
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
			Name          string  `json:"name"`
			ExecutionTime float64 `json:"execution_time"`
			Passed        bool    `json:"passed"`
		} `json:"tests"`
	} `json:"packages"`
	ThirdParties  []string `json:"third_parties"`
	ExecutionTime struct {
		Goget   string `json:"goget,omitempty"`
		Goprove string `json:"goprove"`
		Gotest  string `json:"gotest"`
	} `json:"execution_time"`
	RawOutput struct {
		Goget  string `json:"goget"`
		Gotest string `json:"gotest"`
	} `json:"raw_output"`
	Errors struct {
		Goget  string `json:"goget"`
		Gotest string `json:"gotest"`
	} `json:"errors"`
}

// GetAvgTestDuration returns the average test duration.
func (t ProjectRunner) GetAvgTestDuration() float64 {
	var duration []float64
	for _, pkg := range t.Packages {
		duration = append(duration, pkg.ExecutionTime)
	}
	if len(duration) == 0 {
		return 0
	}
	return xmath.Arithmetic(duration)
}

// GetAvgCodeCov returns the code coverage average.
func (t ProjectRunner) GetAvgCodeCov() float64 {
	var cov []float64
	for _, pkg := range t.Packages {
		cov = append(cov, pkg.Coverage)
	}
	if len(cov) == 0 {
		return 0
	}
	return xmath.Geometric(cov)
}