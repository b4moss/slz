package cli

import (
	"strings"
	"testing"
)

func TestVersionStdout(t *testing.T) {
	stdout, _, err := runCLI(t, Options{}, "--version")
	if err != nil {
		t.Fatalf("expected success, got %v", err)
	}
	if stdout != "slz version 0.2.0\n" {
		t.Fatalf("stdout = %q, want %q", stdout, "slz version 0.2.0\n")
	}
}

func TestVersionExitZero(t *testing.T) {
	_, _, err := runCLI(t, Options{}, "--version")
	if err != nil {
		t.Fatalf("expected success, got %v", err)
	}
}

func TestVersionNotOnStderr(t *testing.T) {
	_, stderr, err := runCLI(t, Options{}, "--version")
	if err != nil {
		t.Fatalf("expected success, got %v", err)
	}
	if strings.Contains(stderr, "0.2.0") {
		t.Fatalf("stderr should not contain version, got %q", stderr)
	}
}

func TestUnknownFlag(t *testing.T) {
	_, _, err := runCLI(t, Options{}, "--no-such-flag")
	if err == nil {
		t.Fatal("expected non-zero exit for unknown flag")
	}
}

func TestUnknownSubcommand(t *testing.T) {
	_, _, err := runCLI(t, Options{}, "nosuch")
	if err == nil {
		t.Fatal("expected non-zero exit for unknown subcommand")
	}
}

func TestNoArgsDoesNotPrintVersion(t *testing.T) {
	stdout, _, err := runCLI(t, Options{}, []string{}...)
	if err != nil {
		t.Fatalf("expected success, got %v", err)
	}
	if strings.Contains(stdout, "slz version 0.2.0") {
		t.Fatalf("stdout should not contain version, got %q", stdout)
	}
}
