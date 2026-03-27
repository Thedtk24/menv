package lmod

import (
	"fmt"
	"os"
	"os/exec"
	"strings"
)

func GetLoadedModules() []string {
	val := os.Getenv("LOADEDMODULES")
	if val == "" {
		return []string{}
	}
	
	parts := strings.Split(val, ":")
	var modules []string
	for _, p := range parts {
		p = strings.TrimSpace(p)
		if p != "" {
			modules = append(modules, p)
		}
	}
	return modules
}

func CheckModuleExists(name string) error {
	cmd := exec.Command("bash", "-c", fmt.Sprintf("module avail %s", name))
	out, err := cmd.CombinedOutput()
	if err != nil {
		return fmt.Errorf("error calling module avail: %w", err)
	}
	
	output := strings.ToLower(string(out))
	if strings.Contains(output, "no module(s)") || strings.Contains(output, "unable to locate") {
		return fmt.Errorf("module '%s' not found", name)
	}
	return nil
}
