package tool

import (
	"bytes"
	"os"
	"path/filepath"
	"strings"
	"testing"
)

// TestGrep_MatchAndExclude验证匹配的行被打印，不匹配的行被排除
func TestGrep_MatchAndExclude(t *testing.T) {
	dir := t.TempDir()
	path := filepath.Join(dir, "sample.txt")
	content := "apple\nbanana\napricot\nblueberry\n"
	if err := os.WriteFile(path, []byte(content), 0o644); err != nil {
		t.Fatalf("setup: %v", err)
	}

	var buf bytes.Buffer
	if err := Grep(&buf, []string{"ap", path}); err != nil {
		t.Fatalf("Grep err: %v", err)
	}
	got := buf.String()

	for _, want := range []string{"apple", "apricot"} {
		if !strings.Contains(got, want) {
			t.Errorf("output should contain %q, got: %q", want, got)
		}
	}
	for _, notWant := range []string{"banana", "blueberry"} {
		if strings.Contains(got, notWant) {
			t.Errorf("output should NOT contain %q, got: %q", notWant, got)
		}
	}
}

// TestGrep_NoMatch验证所有行都不匹配时输出为空，且不报错
func TestGrep_NoMatch(t *testing.T) {
	dir := t.TempDir()
	path := filepath.Join(dir, "sample.txt")
	if err := os.WriteFile(path, []byte("foo\nbar\n"), 0o644); err != nil {
		t.Fatalf("setup: %v", err)
	}

	var buf bytes.Buffer
	if err := Grep(&buf, []string{"xyz", path}); err != nil {
		t.Fatalf("Grep err: %v", err)
	}
	if buf.Len() != 0 {
		t.Errorf("expected empty output, got: %q", buf.String())
	}
}

// TestGrep_NotEnoughArgs验证少于两个参数时返回明确错误
func TestGrep_NotEnoughArgs(t *testing.T) {
	var buf bytes.Buffer
	err := Grep(&buf, []string{"only-pattern"})
	if err == nil {
		t.Fatalf("expected err, got nil")
	}
	if !strings.Contains(err.Error(), "not enough args") {
		t.Errorf("unexpected err: %v", err)
	}
}

// TestGrep_FileNotFound验证文件不存在时返回错误而不是panic
func TestGrep_FileNotFound(t *testing.T) {
	var buf bytes.Buffer
	err := Grep(&buf, []string{"x", filepath.Join(t.TempDir(), "no-such-file")})
	if err == nil {
		t.Fatalf("expected err for missing file, got nil")
	}
}
