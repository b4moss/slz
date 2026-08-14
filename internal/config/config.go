package config

import (
	"fmt"
	"os"
	"path/filepath"
)

const appName = "slz"

func Dir() (string, error) {
	if xdg := os.Getenv("XDG_CONFIG_HOME"); xdg != "" {
		return filepath.Join(xdg, appName), nil
	}
	home, err := os.UserHomeDir()
	if err != nil {
		return "", err
	}
	return filepath.Join(home, ".config", appName), nil
}

func File(dir string) string {
	return filepath.Join(dir, "config.yaml")
}

func CommandsDir(dir string) string {
	return filepath.Join(dir, "commands")
}

func Setup() error {
	dir, err := Dir()
	if err != nil {
		return err
	}
	if err := os.MkdirAll(dir, 0o755); err != nil {
		return err
	}
	if err := os.MkdirAll(CommandsDir(dir), 0o755); err != nil {
		return err
	}
	cfg := File(dir)
	info, err := os.Stat(cfg)
	if os.IsNotExist(err) {
		return os.WriteFile(cfg, []byte{}, 0o644)
	}
	if err != nil {
		return err
	}
	if info.IsDir() {
		return fmt.Errorf("%s is a directory", cfg)
	}
	return nil
}

func Read() ([]byte, error) {
	dir, err := Dir()
	if err != nil {
		return nil, err
	}
	data, err := os.ReadFile(File(dir))
	if err != nil {
		if os.IsNotExist(err) {
			return nil, fmt.Errorf("config.yaml not found; run slz setup")
		}
		return nil, err
	}
	return data, nil
}
