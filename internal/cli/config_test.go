package cli

import (
	"os"
	"path/filepath"
	"strings"
	"testing"
)

func TestConfigPrintsFile(t *testing.T) {
	xdg := t.TempDir()
	t.Setenv("XDG_CONFIG_HOME", xdg)
	dir := filepath.Join(xdg, "slz")
	if err := os.MkdirAll(dir, 0o755); err != nil {
		t.Fatal(err)
	}
	const body = "foo: bar\n"
	if err := os.WriteFile(filepath.Join(dir, "config.yaml"), []byte(body), 0o644); err != nil {
		t.Fatal(err)
	}

	stdout, _, err := runCLI(t, Options{}, "config")
	if err != nil {
		t.Fatalf("expected success, got %v", err)
	}
	if stdout != body {
		t.Fatalf("stdout = %q, want %q", stdout, body)
	}
}

func TestConfigEmptyFile(t *testing.T) {
	xdg := t.TempDir()
	t.Setenv("XDG_CONFIG_HOME", xdg)
	dir := filepath.Join(xdg, "slz")
	if err := os.MkdirAll(dir, 0o755); err != nil {
		t.Fatal(err)
	}
	if err := os.WriteFile(filepath.Join(dir, "config.yaml"), []byte{}, 0o644); err != nil {
		t.Fatal(err)
	}

	_, _, err := runCLI(t, Options{}, "config")
	if err != nil {
		t.Fatalf("expected success for empty config.yaml, got %v", err)
	}
}

func TestConfigHonorsXDGConfigHome(t *testing.T) {
	xdg := t.TempDir()
	t.Setenv("XDG_CONFIG_HOME", xdg)
	dir := filepath.Join(xdg, "slz")
	if err := os.MkdirAll(dir, 0o755); err != nil {
		t.Fatal(err)
	}
	const body = "from: xdg\n"
	if err := os.WriteFile(filepath.Join(dir, "config.yaml"), []byte(body), 0o644); err != nil {
		t.Fatal(err)
	}

	stdout, _, err := runCLI(t, Options{}, "config")
	if err != nil {
		t.Fatalf("expected success, got %v", err)
	}
	if stdout != body {
		t.Fatalf("stdout = %q, want %q", stdout, body)
	}
}

func TestConfigMissingFile(t *testing.T) {
	t.Setenv("XDG_CONFIG_HOME", t.TempDir())
	_, _, err := runCLI(t, Options{}, "config")
	if err == nil {
		t.Fatal("expected non-zero exit when config.yaml is missing")
	}
}

func TestConfigMissingMentionsSetup(t *testing.T) {
	t.Setenv("XDG_CONFIG_HOME", t.TempDir())
	_, stderr, err := runCLI(t, Options{}, "config")
	if err == nil {
		t.Fatal("expected non-zero exit when config.yaml is missing")
	}
	if !strings.Contains(stderr, "slz setup") {
		t.Fatalf("stderr should mention slz setup, got %q", stderr)
	}
}

func TestConfigRejectsExtraArgs(t *testing.T) {
	t.Setenv("XDG_CONFIG_HOME", t.TempDir())
	_, _, err := runCLI(t, Options{}, "config", "extra")
	if err == nil {
		t.Fatal("expected non-zero exit for extra args")
	}
}
