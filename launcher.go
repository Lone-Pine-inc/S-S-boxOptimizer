package main

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"time"
)

func main() {
	fmt.Println("🚀 s&box Launcher")
	fmt.Println("==================")
	
	steamPath := "C:\\Program Files (x86)\\Steam\\steam.exe"
	
	// Load custom path if exists
	defaultSboxPath := "C:\\Program Files (x86)\\Steam\\steamapps\\common\\sbox"
	settingsFilePath := "optimizer_settings.txt"
	
	sboxPath := defaultSboxPath
	if data, err := os.ReadFile(settingsFilePath); err == nil {
		customPath := strings.TrimSpace(string(data))
		if customPath != "" {
			sboxPath = customPath
			fmt.Println("✅ Using custom s&box path:", sboxPath)
		}
	}
	
	// Build full cfg path
	cfgPath := filepath.Join(sboxPath, "core", "cfg")
	
	// Check if Steam exists
	if _, err := os.Stat(steamPath); os.IsNotExist(err) {
		fmt.Println("❌ Error: Steam not found!")
		fmt.Println("Path:", steamPath)
		fmt.Println("\nPress Enter to exit...")
		fmt.Scanln()
		return
	}
	
	fmt.Println("✅ Steam found")
	fmt.Println("📝 Launching s&box with settings from graphics_config.vcfg...")
	fmt.Println("📂 Config path:", cfgPath)
	
	// Launch s&box through Steam with config exec
	cmd := exec.Command(steamPath, "-applaunch", "590830", "+exec", "graphics_config.vcfg")
	
	err := cmd.Start()
	if err != nil {
		fmt.Printf("❌ Launch error: %v\n", err)
		fmt.Println("\nPress Enter to exit...")
		fmt.Scanln()
		return
	}
	
	fmt.Println("✅ s&box launched!")
	fmt.Println("\nWindow will close in 3 seconds...")
	time.Sleep(3 * time.Second)
}