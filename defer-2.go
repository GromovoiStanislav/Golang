package main

import (
	"errors"
	"fmt"
)

type MergeDictsJob struct {
	Dicts      []map[string]string
	Merged     map[string]string
	IsFinished bool
}

// errors
var (
	errNotEnoughDicts = errors.New("at least 2 dictionaries are required")
	errNilDict        = errors.New("nil dictionary")
)

func (j *MergeDictsJob) setFinished() {
	j.IsFinished = true
}

func ExecuteMergeDictsJob(job *MergeDictsJob) (*MergeDictsJob, error) {
	defer job.setFinished()

	if len(job.Dicts) < 2 {
		return job, errNotEnoughDicts
	}

	job.Merged = make(map[string]string)
	for _, dict := range job.Dicts {
		if dict == nil {
			return job, errNilDict
		}

		for k, v := range dict {
			job.Merged[k] = v
		}
	}

	return job, nil
}

func main() {
	fmt.Println(ExecuteMergeDictsJob(&MergeDictsJob{}))// &{[]map[] true} at least 2 dictionaries are required
	fmt.Println(ExecuteMergeDictsJob(&MergeDictsJob{Dicts: []map[string]string{{"1": "2"}, nil}}))// &{[map[1:2] map[]] map[1:2] true} nil dictionary
	fmt.Println(ExecuteMergeDictsJob(&MergeDictsJob{Dicts: []map[string]string{{"1": "2"}, {"3": "4"}}}))// &{[map[1:2] map[3:4]] map[1:2 3:4] true} <nil>
}