package main

import (
	"encoding/json"
	"fmt"
	"os"
)

var (
	version = "1.0.0"
	configPath = "./config.json"
)

type Server struct {
	IP       string `json:"ip"`
	User     string `json:"user"`
	Password string `json:"password"`
}

type Config struct {
	Servers []Server `json:"servers"`
}

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: myapp <command>")
		os.Exit(1)
	}

	// Commands
	var version string = "version"
	var add string = "add"

	// Parse the command
	switch os.Args[1] {
	case version:
		Version()
	case add:
		if args := os.Args[2:]; len(args) != 3 {
			fmt.Println("Usage: myadd add <ip> <user> <password>")
			os.Exit(1)
		}
		addServer(os.Args[2], os.Args[3], os.Args[4])

	// case remove:
	// case list:
	// case compile:
	default:
		fmt.Printf("Unknown command: %s\n", os.Args[1])
		fmt.Println("Available commands: version")
		os.Exit(1)
	}
}

func Version() {
	fmt.Printf("Version: %s\n", version)
}

func addServer(ip string, user string, password string) {
	//***Check validity

	//***check config file
	newServer := Server{
		IP:       ip,
		User:     user,
		Password: password,
	}
	// Load existing config or create new one
	config, err := loadConfig()
	if err != nil {
		config = &Config{
			Servers: []Server{},
		}
	}
	// Add new server to config
	config.Servers = append(config.Servers, newServer)

	// Save updated config
	if err := saveConfig(config); err != nil {
		fmt.Printf("Error saving config: %v\n", err)
		os.Exit(1)
	}

	//***add server
}

func loadConfig() (*Config, error) {
	configPath := configPath
	data, err := os.ReadFile(configPath)
	if err != nil {
		return nil, err
	}

	var config Config
	if err := json.Unmarshal(data, &config); err != nil {
		return nil, err
	}

	return &config, nil
}

func saveConfig(config *Config) error {
	configPath := configPath
	data, err := json.MarshalIndent(config, "", "    ")
	if err != nil {
		return err
	}

	return os.WriteFile(configPath, data, 0600)
}