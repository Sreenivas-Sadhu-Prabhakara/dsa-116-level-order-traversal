package main

import (
	"reflect"
	"testing"
)

func TestLevelOrder(t *testing.T) {
	root := tn(3, tn(9, nil, nil), tn(20, tn(15, nil, nil), tn(7, nil, nil)))
	want := [][]int{{3}, {9, 20}, {15, 7}}
	if got := LevelOrder(root); !reflect.DeepEqual(got, want) {
		t.Errorf("LevelOrder = %v, want %v", got, want)
	}
	if got := LevelOrder(nil); !reflect.DeepEqual(got, [][]int{}) {
		t.Errorf("LevelOrder(nil) = %v", got)
	}
}

func TestSpiralOrder(t *testing.T) {
	root := tn(1, tn(2, tn(4, nil, nil), tn(5, nil, nil)), tn(3, tn(6, nil, nil), tn(7, nil, nil)))
	want := [][]int{{1}, {3, 2}, {4, 5, 6, 7}}
	if got := SpiralOrder(root); !reflect.DeepEqual(got, want) {
		t.Errorf("SpiralOrder = %v, want %v", got, want)
	}
}
