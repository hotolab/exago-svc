package repository

import (
	"time"

	"github.com/hotolab/exago-svc/repository/model"
)

func (r *Repository) SetName(name string) {
	r.Name = name
}

func (r *Repository) SetData(d model.Data) {
	r.Data = d
}

func (r *Repository) SetCodeStats(cs model.CodeStats) {
	r.Data.CodeStats = cs
}

func (r *Repository) SetLintMessages(lm model.LintMessages) {
	r.Data.LintMessages = lm
}

func (r *Repository) SetProjectRunner(tr model.ProjectRunner) {
	r.Data.ProjectRunner = tr
}

// SetStartTime stores the moment the processing started.
func (r *Repository) SetStartTime(t time.Time) {
	r.startTime = t
}

// SetExecutionTime sets the processing execution time.
// The value is then used to determine an ETA for refreshing data.
func (r *Repository) SetExecutionTime(duration time.Duration) {
	r.Data.ExecutionTime = (duration - (duration % time.Second)).String()
}

// SetLastUpdate sets the last update timestamp.
func (r *Repository) SetLastUpdate(t time.Time) {
	r.Data.LastUpdate = t
}

// SetMetadata sets repository metadata such as description, stars...
func (r *Repository) SetMetadata(m model.Metadata) {
	r.Data.Metadata = m
}

// SetError assigns a processing error to the given type (ex. ProjectRunner).
// This helps keep track of what went wrong.
func (r *Repository) SetError(tp string, err error) {
	if r.Data.Errors == nil {
		r.Data.Errors = make(map[string]string)
	}
	r.Data.Errors[tp] = err.Error()
}
