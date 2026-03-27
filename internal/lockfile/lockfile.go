package lockfile

import (
	"fmt"
	"os"
	"path/filepath"
	"time"

	"gopkg.in/yaml.v3"
)

type Environment struct {
	Name    string            `yaml:"name"`
	Created time.Time         `yaml:"created"`
	Modules []string          `yaml:"modules"`
	EnvVars map[string]string `yaml:"env_vars"`
}

func getConfigDir() (string, error) {
	home, err := os.UserHomeDir()
	if err != nil {
		return "", fmt.Errorf("could not find user home directory: %w", err)
	}
	dir := filepath.Join(home, ".menv")
	if err := os.MkdirAll(dir, 0755); err != nil {
		return "", fmt.Errorf("could not create .menv directory: %w", err)
	}
	return dir, nil
}

func getFilePath(name string) (string, error) {
	dir, err := getConfigDir()
	if err != nil {
		return "", err
	}
	return filepath.Join(dir, fmt.Sprintf("%s.yaml", name)), nil
}

func SaveByName(name string, env *Environment) error {
	path, err := getFilePath(name)
	if err != nil {
		return err
	}

	env.Name = name
	if env.Created.IsZero() {
		env.Created = time.Now()
	}

	data, err := yaml.Marshal(env)
	if err != nil {
		return fmt.Errorf("YAML serialization error: %w", err)
	}

	if err := os.WriteFile(path, data, 0644); err != nil {
		return fmt.Errorf("file write error: %w", err)
	}

	return nil
}

func LoadByName(name string) (*Environment, error) {
	path, err := getFilePath(name)
	if err != nil {
		return nil, err
	}

	data, err := os.ReadFile(path)
	if err != nil {
		return nil, fmt.Errorf("environment '%s' not found: %w", name, err)
	}

	var env Environment
	if err := yaml.Unmarshal(data, &env); err != nil {
		return nil, fmt.Errorf("invalid YAML format for '%s': %w", name, err)
	}

	return &env, nil
}

func DeleteByName(name string) error {
	path, err := getFilePath(name)
	if err != nil {
		return err
	}
	if err := os.Remove(path); err != nil {
		return fmt.Errorf("could not delete '%s' (it might not exist): %w", name, err)
	}
	return nil
}

func ListAll() ([]string, error) {
	dir, err := getConfigDir()
	if err != nil {
		return nil, err
	}

	entries, err := os.ReadDir(dir)
	if err != nil {
		return nil, fmt.Errorf("could not read %s: %w", dir, err)
	}

	var envs []string
	for _, entry := range entries {
		if !entry.IsDir() && filepath.Ext(entry.Name()) == ".yaml" {
			name := entry.Name()[:len(entry.Name())-5]
			envs = append(envs, name)
		}
	}
	return envs, nil
}

func SaveLocalLock(env *Environment) error {
	env.Created = time.Now()
	data, err := yaml.Marshal(env)
	if err != nil {
		return fmt.Errorf(".menv.lock serialization error: %w", err)
	}
	if err := os.WriteFile(".menv.lock", data, 0644); err != nil {
		return fmt.Errorf(".menv.lock write error: %w", err)
	}
	return nil
}

func LoadLocalLock() (*Environment, error) {
	data, err := os.ReadFile(".menv.lock")
	if err != nil {
		return nil, fmt.Errorf("could not read .menv.lock: %w", err)
	}
	var env Environment
	if err := yaml.Unmarshal(data, &env); err != nil {
		return nil, fmt.Errorf("invalid file format in .menv.lock: %w", err)
	}
	return &env, nil
}
