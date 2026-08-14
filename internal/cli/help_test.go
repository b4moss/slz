package cli

import (
	"os"
	"path/filepath"
	"strings"
	"testing"
)

func TestHelpExitZero(t *testing.T) {
	_, _, err := runCLI(t, Options{}, "--help")
	if err != nil {
		t.Fatalf("expected success, got %v", err)
	}
}

func TestHelpListsSetup(t *testing.T) {
	stdout, _, err := runCLI(t, Options{}, "--help")
	if err != nil {
		t.Fatalf("expected success, got %v", err)
	}
	if !strings.Contains(stdout, "setup") {
		t.Fatalf("help should list setup, got %q", stdout)
	}
}

func TestHelpListsConfig(t *testing.T) {
	stdout, _, err := runCLI(t, Options{}, "--help")
	if err != nil {
		t.Fatalf("expected success, got %v", err)
	}
	if !strings.Contains(stdout, "config") {
		t.Fatalf("help should list config, got %q", stdout)
	}
}

func TestHelpListsMac(t *testing.T) {
	stdout, _, err := runCLI(t, Options{}, "--help")
	if err != nil {
		t.Fatalf("expected success, got %v", err)
	}
	if !strings.Contains(stdout, "mac") {
		t.Fatalf("help should list mac, got %q", stdout)
	}
}

func TestHelpUnknownTopic(t *testing.T) {
	_, _, err := runCLI(t, Options{}, "help", "nosuch")
	if err == nil {
		t.Fatal("expected non-zero exit for unknown help topic")
	}
}

func TestHelpDoesNotCreateConfigDir(t *testing.T) {
	xdg := t.TempDir()
	t.Setenv("XDG_CONFIG_HOME", xdg)
	_, _, err := runCLI(t, Options{}, "--help")
	if err != nil {
		t.Fatalf("expected success, got %v", err)
	}
	if _, err := os.Stat(filepath.Join(xdg, "slz")); !os.IsNotExist(err) {
		t.Fatal("help should not create the config directory")
	}
}

func TestHelpDoesNotRunExternalCommands(t *testing.T) {
	r := &countingRunner{}
	_, _, err := runCLI(t, Options{Runner: r}, "--help")
	if err != nil {
		t.Fatalf("expected success, got %v", err)
	}
	if r.n != 0 {
		t.Fatalf("help started %d external commands", r.n)
	}
}
