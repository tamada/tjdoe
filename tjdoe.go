package tjdoe

import (
	"encoding/csv"
	"io"
	"math/rand"
	"strconv"

	"github.com/seehuhn/mt19937"
)

type TJDoe struct {
	random *rand.Rand
}

func New(seed int64) *TJDoe {
	tjdoe := new(TJDoe)
	tjdoe.random = rand.New(mt19937.New())
	tjdoe.random.Seed(seed)
	return tjdoe
}

func (tjdoe *TJDoe) AnonymizeDirectory(from, to string, mapping []Mapping) error {
	// TODO
	return nil
}

func createCsvItems(student *Student, labels []string) []string {
	array := []string{student.AnonymizedID, student.AnonymizedFinalScore}
	for _, label := range labels {
		value, ok := student.Scores[label]
		valueString := strconv.Itoa(value)
		if !ok {
			valueString = ""
		}
		array = append(array, valueString)
	}
	return array
}

func (tjdoe *TJDoe) OutputAnonymizedScores(students []*Student, dest io.Writer) error {
	header := createHeader(students)
	writer := csv.NewWriter(dest)
	writer.Write(header)
	labels := header[2:]
	for _, student := range students {
		writer.Write(createCsvItems(student, labels))
	}
	return nil
}

func contains(array []string, value string) bool {
	for _, item := range array {
		if item == value {
			return true
		}
	}
	return false
}

func assignmentNames(assignments []string, scores map[string]int) []string {
	for key, _ := range scores {
		if !contains(assignments, key) {
			assignments = append(assignments, key)
		}
	}
	return assignments
}

func createHeader(students []*Student) []string {
	assignments := []string{}
	for _, student := range students {
		assignments = assignmentNames(assignments, student.Scores)
	}
	header := []string{"id", "final score"}
	header = append(header, assignments...)
	return header
}
