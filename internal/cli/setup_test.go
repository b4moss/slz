package cli

import (
	"os"
	"path/filepath"
	"testing"
)

func TestSetupCreatesLayout(t *testing.T) {
	xdg := t.TempDir()
	home := t.TempDir()
	t.Setenv("XDG_CONFIG_HOME", xdg)
	t.Setenv("HOME", home)

	_, _, err := runCLI(t, Options{}, "setup")
	if err != nil {
		t.Fatalf("expected success, got %v", err)
	}

	dir := filepath.Join(xdg, "slz")
	if _, err := os.Stat(dir); err != nil {
		t.Fatalf("config dir: %v", err)
	}
	data, err := os.ReadFile(filepath.Join(dir, "config.yaml"))
	if err != nil {
		t.Fatalf("config.yaml: %v", err)
	}
	if len(data) != 0 {
		t.Fatalf("config.yaml should be empty, got %q", data)
	}
	info, err := os.Stat(filepath.Join(dir, "commands"))
	if err != nil {
		t.Fatalf("commands/: %v", err)
	}
	if !info.IsDir() {
		t.Fatal("commands/ should be a directory")
	}
}

func TestSetupIsIdempotent(t *testing.T) {
	xdg := t.TempDir()
	t.Setenv("XDG_CONFIG_HOME", xdg)
	dir := filepath.Join(xdg, "slz")
	if err := os.MkdirAll(filepath.Join(dir, "commands"), 0o755); err != nil {
		t.Fatal(err)
	}
	const keep = "keep me\n"
	if err := os.WriteFile(filepath.Join(dir, "config.yaml"), []byte(keep), 0o644); err != nil {
		t.Fatal(err)
	}

	_, _, err := runCLI(t, Options{}, "setup")
	if err != nil {
		t.Fatalf("expected success, got %v", err)
	}
	data, err := os.ReadFile(filepath.Join(dir, "config.yaml"))
	if err != nil {
		t.Fatal(err)
	}
	if string(data) != keep {
		t.Fatalf("setup overwrote config.yaml: got %q", data)
	}
}

func TestSetupHonorsXDGConfigHome(t *testing.T) {
	xdg := t.TempDir()
	home := t.TempDir()
	t.Setenv("XDG_CONFIG_HOME", xdg)
	t.Setenv("HOME", home)

	_, _, err := runCLI(t, Options{}, "setup")
	if err != nil {
		t.Fatalf("expected success, got %v", err)
	}
	if _, err := os.Stat(filepath.Join(xdg, "slz")); err != nil {
		t.Fatalf("expected dir under XDG_CONFIG_HOME: %v", err)
	}
	if _, err := os.Stat(filepath.Join(home, ".config", "slz")); !os.IsNotExist(err) {
		t.Fatal("setup should not create ~/.config/slz when XDG_CONFIG_HOME is set")
	}
}

func TestSetupFailsWhenConfigIsDirectory(t *testing.T) {
	xdg := t.TempDir()
	t.Setenv("XDG_CONFIG_HOME", xdg)
	if err := os.MkdirAll(filepath.Join(xdg, "slz", "config.yaml"), 0o755); err != nil {
		t.Fatal(err)
	}

	_, _, err := runCLI(t, Options{}, "setup")
	if err == nil {
		t.Fatal("expected non-zero exit when config.yaml is a directory")
	}
}

func TestSetupDoesNotCreateDoctorWarned(t *testing.T) {
	xdg := t.TempDir()
	state := t.TempDir()
	t.Setenv("XDG_CONFIG_HOME", xdg)
	t.Setenv("XDG_STATE_HOME", state)

	_, _, err := runCLI(t, Options{}, "setup")
	if err != nil {
		t.Fatalf("expected success, got %v", err)
	}
	if _, err := os.Stat(filepath.Join(state, "slz", "doctor-warned")); !os.IsNotExist(err) {
		t.Fatal("setup should not create doctor-warned")
	}
}

func TestSetupRejectsExtraArgs(t *testing.T) {
	t.Setenv("XDG_CONFIG_HOME", t.TempDir())
	_, _, err := runCLI(t, Options{}, "setup", "extra")
	if err == nil {
		t.Fatal("expected non-zero exit for extra args")
	}
}
