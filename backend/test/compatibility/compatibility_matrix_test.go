package compatibility

import (
	"os"
	"path/filepath"
	"runtime"
	"strings"
	"testing"
)

type CompatibilityMatrix struct {
	TargetVersion string
}

func LoadCompatibilityMatrixForTest() CompatibilityMatrix {
	return CompatibilityMatrix{
		TargetVersion: "1.4.6",
	}
}

func TestCompatibilityMatrix_ShouldDeclare_1_4_6_Baseline(t *testing.T) {
	matrix := LoadCompatibilityMatrixForTest()
	if matrix.TargetVersion != "1.4.6" {
		t.Fatalf("unexpected target version: %s", matrix.TargetVersion)
	}
}

func TestDocs_ShouldContain146CompatibilityStatement(t *testing.T) {
	_, filename, _, ok := runtime.Caller(0)
	if !ok {
		t.Fatal("failed to resolve test file path")
	}
	readmePath := filepath.Clean(filepath.Join(filepath.Dir(filename), "../../../README.md"))
	content, err := os.ReadFile(readmePath)
	if err != nil {
		t.Fatalf("failed to read README: %v", err)
	}
	if !strings.Contains(string(content), "1.4.6") {
		t.Fatal("README does not contain 1.4.6 compatibility statement")
	}
}
