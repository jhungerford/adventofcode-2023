package util

import (
	"errors"
	"reflect"
	"strconv"
	"testing"
)

func Test_ReadInputLines(t *testing.T) {
	want := []string{"1", "2", "3"}

	got, err := ReadInputLines("files_sample.txt")
	if err != nil {
		t.Fatalf("failed to read input lines: %v", err)
	}

	if !reflect.DeepEqual(want, got) {
		t.Errorf("ReadInputLines() = %+v, want %+v", got, want)
	}
}

func Test_ParseInputLines(t *testing.T) {
	want := []int{1, 2, 3}

	got, err := ParseInputLines("files_sample.txt", strconv.Atoi)
	if err != nil {
		t.Fatalf("failed to parse input lines: %v", err)
	}

	if !reflect.DeepEqual(want, got) {
		t.Errorf("ParseInputLines() = %+v, want %+v", got, want)
	}
}

func Test_ParseInputLinesFailParse(t *testing.T) {
	parseErr := errors.New("no lines are valid")

	_, err := ParseInputLines("files_sample.txt", func(line string) (string, error) {
		return "", parseErr
	})

	if err == nil || !errors.Is(err, parseErr) {
		t.Fatalf("ParseInputLines should have thrown a parseErr, but err was %v", err)
	}
}
