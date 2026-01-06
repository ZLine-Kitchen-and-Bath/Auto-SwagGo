package tests

import (
	"testing"

	"github.com/ZLine-Kitchen-and-Bath/Auto-SwagGo/internal/ext"
)

func TestContains(t *testing.T) {
	if !ext.Contains([]int{1, 2, 3}, 2) {
		t.Error("Expected true, got false")
	}

	if ext.Contains([]int{1, 2, 3}, 4) {
		t.Error("Expected false, got true")
	}
}
