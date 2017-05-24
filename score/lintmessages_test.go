package score_test

import (
	"testing"

	"github.com/jgautheron/exago/repository/model"
	"github.com/jgautheron/exago/score"
)

func TestLintMessages(t *testing.T) {
	var tests = []struct {
		messages map[string]int
		loc      int
		operator string
		expected float64
		desc     string
	}{
		{map[string]int{"gas": 5}, 500, "<", 60, "1 potential security issue every 100 loc"},
		{map[string]int{"gofmt": 5}, 500, "<", 20, "gofmt is a must-have"},
		{map[string]int{"golint": 20}, 500, ">", 50, "golint is verbose"},
	}

	for _, tt := range tests {
		d := model.Data{}
		d.LintMessages = getStubMessages(tt.messages)
		d.ProjectRunner.CodeStats.Data = map[string]int{"loc": tt.loc}
		evaluator := score.LintMessagesEvaluator()
		evaluator.Setup()
		res := evaluator.Calculate(d)

		switch tt.operator {
		case "<":
			if res.Score > tt.expected {
				t.Errorf("Wrong score %s", tt.desc)
			}
		case ">":
			if res.Score < tt.expected {
				t.Errorf("Wrong score %s", tt.desc)
			}
		case "=":
			if res.Score != tt.expected {
				t.Errorf("Wrong score %s", tt.desc)
			}
		}
	}
}

func getStubMessages(messages map[string]int) model.LintMessages {
	fileName := "foo.go"
	m := map[string]map[string][]map[string]interface{}{}
	for linter, count := range messages {
		m[fileName] = map[string][]map[string]interface{}{}
		m[fileName][linter] = []map[string]interface{}{}
		for i := 0; i < count; i++ {
			m[fileName][linter] = append(m[fileName][linter], map[string]interface{}{
				"severity": "warning",
			})
		}
	}
	return model.LintMessages(m)
}
