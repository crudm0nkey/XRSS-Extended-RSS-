package luafunctions

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
)

type PluginConfig struct {
	MainURL string `json:"main_url"`
}

func LoadPlugins() ([]os.DirEntry, error) {
	pluginNames, err := readAllPluginNames()
	if err != nil {
		return []os.DirEntry{}, err
	}

	for _, plugin := range pluginNames {
		pluginJsonCfg, err := readPluginJson(plugin.Name())
		if err != nil {
			return []os.DirEntry{}, err
		}

		fmt.Println(pluginJsonCfg.MainURL)
	}

	return []os.DirEntry{}, nil
}

func readAllPluginNames() ([]os.DirEntry, error) {
	dirs, err := os.ReadDir("plugins/")
	if err != nil {
		return []os.DirEntry{}, err
	}

	var pluginFiles []os.DirEntry
	for _, p := range dirs {
		if !p.IsDir() {
			pluginFiles = append(pluginFiles, p)
		}
	}

	return pluginFiles, nil
}

func readPluginJson(pluginName string) (PluginConfig, error) {
	var pluginJson PluginConfig
	f, err := os.Open(fmt.Sprintf("plugins/%s", pluginName))
	if err != nil {
		log.Fatal(err.Error())
	}
	defer f.Close()

	if err := json.NewDecoder(f).Decode(&pluginJson); err != nil {
		return PluginConfig{}, err
	}

	return pluginJson, nil
}
