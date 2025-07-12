package main

import (
	"strings"
	"testing"
)

func TestAddName(t *testing.T) {
	g := newStGrade("TestStudent")

	g.addName("Math", 90.0)
	g.addName("Science", 85.5)

	if len(g.Subj) != 2 {
		t.Errorf("Expected 2 subjects, got %d", len(g.Subj))
	}

	if g.Subj["Math"] != 90.0 {
		t.Errorf("Expected Math grade to be 90.0, got %.2f", g.Subj["Math"])
	}

	if g.Subj["Science"] != 85.5 {
		t.Errorf("Expected Science grade to be 85.5, got %.2f", g.Subj["Science"])
	}
}

func TestFormatOutput(t *testing.T) {
	g := newStGrade("TestStudent")
	g.addName("Math", 80)
	g.addName("Science", 90)

	report := g.format()

	if !strings.Contains(report, "Math") || !strings.Contains(report, "Science") {
		t.Error("Report should contain subject names")
	}

	if !strings.Contains(report, "80") || !strings.Contains(report, "90") {
		t.Error("Report should contain grade values")
	}

	if !strings.Contains(report, "total:") || !strings.Contains(report, "avg:") {
		t.Error("Report should contain total and avg")
	}

	if !strings.Contains(report, "85.00") {
		t.Error("Expected average 85.00 to appear in report")
	}
}
