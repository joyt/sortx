package sort

import (
	"reflect"
	"testing"
)

type testStruct struct {
	Name   string
	Age    uint8
	Points float64
	N      int
}

func TestSortByField(t *testing.T) {
	testDocs := []testStruct{
		{"Q", 17, 10.0, 1},
		{"B", 2, 3.0, 2},
		{"G", 7, 2.0, 0},
		{"A", 1, 1.0, 3},
	}
	SortByField(testDocs, "Name", Ascending)
	correct := []testStruct{
		{"A", 1, 1.0, 3},
		{"B", 2, 3.0, 2},
		{"G", 7, 2.0, 0},
		{"Q", 17, 10.0, 1},
	}
	if !reflect.DeepEqual(testDocs, correct) {
		t.Errorf("Could not sort by Name, got %v, expected %v", testDocs, correct)
	}
	SortByField(testDocs, "Age", Descending)
	correct = []testStruct{
		{"Q", 17, 10.0, 1},
		{"G", 7, 2.0, 0},
		{"B", 2, 3.0, 2},
		{"A", 1, 1.0, 3},
	}
	if !reflect.DeepEqual(testDocs, correct) {
		t.Errorf("Could not sort by Name, got %v, expected %v", testDocs, correct)
	}
	SortByField(testDocs, "Points", Ascending)
	correct = []testStruct{
		{"A", 1, 1.0, 3},
		{"G", 7, 2.0, 0},
		{"B", 2, 3.0, 2},
		{"Q", 17, 10.0, 1},
	}
	if !reflect.DeepEqual(testDocs, correct) {
		t.Errorf("Could not sort by Name, got %v, expected %v", testDocs, correct)
	}
	SortByField(testDocs, "N", Descending)
	correct = []testStruct{
		{"A", 1, 1.0, 3},
		{"B", 2, 3.0, 2},
		{"Q", 17, 10.0, 1},
		{"G", 7, 2.0, 0},
	}
	if !reflect.DeepEqual(testDocs, correct) {
		t.Errorf("Could not sort by Name, got %v, expected %v", testDocs, correct)
	}
}
