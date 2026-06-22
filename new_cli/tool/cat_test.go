package tool

import (
	"bytes"
	"os"
	"path/filepath"
	"strings"
	"testing"
)

// TestCat_OK验证正常路径把文件内容原样写到writer。
func TestCat_OK(t *testing.T) {
	dir := t.TempDir()
	path := filepath.Join(dir, "hello.txt")
	content := "line1\nline2\nline3\n"
	if err := os.WriteFile(path, []byte(content), 0o644); err != nil {
		t.Fatalf("setup write file: %v", err)
	}

	var buf bytes.Buffer
	if err := Cat(&buf, []string{path}); err != nil {
		t.Fatalf("Cat returned err: %v", err)
	}

	got := buf.String()
	want := "line1\nline2\nline3\n"
	if got != want {
		t.Errorf("output mismatch\n got: %q\nwant: %q", got, want)
	}
}

// TestCat_NoArgs验证没传文件名时返回明确错误。
func TestCat_NoArgs(t *testing.T) {
	var buf bytes.Buffer
	err := Cat(&buf, nil)
	if err == nil {
		t.Fatalf("expected err for empty args, got nil")
	}
	if !strings.Contains(err.Error(), "not enough args") {
		t.Errorf("unexpected err message: %v", err)
	}
	if buf.Len() != 0 {
		t.Errorf("expected no output on err, got: %q", buf.String())
	}
}

// TestCat_FileNotFound验证文件不存在时返回错误而不是panic。
func TestCat_FileNotFound(t *testing.T) {
	var buf bytes.Buffer
	err := Cat(&buf, []string{filepath.Join(t.TempDir(), "does-not-exist.txt")})
	if err == nil {
		t.Fatalf("expected err for missing file, got nil")
	}
}
