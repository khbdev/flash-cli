package flashcore

import (
	"encoding/json"
	"fmt"
	"os"
)

type Config struct {
	DefaultBranch string `json:"default_branch"`
}

const configPath = ".flash/config.json"

func LoadConfig() (Config, error) {
	var cfg Config
	data, err := os.ReadFile(configPath)
	if err != nil {
		return cfg, fmt.Errorf("config o'qib bo'lmadi: %w", err)
	}

	err = json.Unmarshal(data, &cfg)
	if err != nil {
		return cfg, fmt.Errorf("jsonni parse qilib bo'lmadi: %w", err)
	}

	return cfg, nil
}

func SaveConfig(cfg Config) error {
	data, err := json.MarshalIndent(cfg, "", "  ")
	if err != nil {
		return fmt.Errorf("jsonga o'girishda xatolik: %w", err)
	}

	err = os.WriteFile(configPath, data, 0644)
	if err != nil {
		return fmt.Errorf("configni yozib bo'lmadi: %w", err)
	}

	return nil
}

func UpdateDefaultBranch(newBranch string) error {
	cfg, err := LoadConfig()
	if err != nil {
		return fmt.Errorf("configni o'qishda xatolik: %w", err)
	}

	cfg.DefaultBranch = newBranch

	err = SaveConfig(cfg)
	if err != nil {
		return fmt.Errorf("configni yangilab bo'lmadi: %w", err)
	}

	return nil
}
