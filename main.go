package main

import (
	"fmt"
	"os"
	"path/filepath"
	"strconv"
	"strings"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
	"fyne.io/fyne/v2/dialog"
)

func main() {
	myApp := app.New()
	myWindow := myApp.NewWindow("s&box Graphics Optimizer - Settings")

	// Default path to s&box installation
	defaultSboxPath := "C:\\Program Files (x86)\\Steam\\steamapps\\common\\sbox"
	
	// Settings file to store custom path
	settingsFilePath := "optimizer_settings.txt"
	
	// Load custom path if exists
	sboxPath := defaultSboxPath
	if data, err := os.ReadFile(settingsFilePath); err == nil {
		customPath := strings.TrimSpace(string(data))
		if customPath != "" {
			sboxPath = customPath
		}
	}
	
	// Build full cfg path
	cfgPath := filepath.Join(sboxPath, "core", "cfg")
	configFilePath := filepath.Join(cfgPath, "graphics_config.vcfg")

	// Path display with selectable text
	pathLabel := widget.NewLabel("s&box Installation Path:")
	pathValueLabel := widget.NewLabel(sboxPath)
	pathValueLabel.Wrapping = fyne.TextWrapWord
	
	changePathButton := widget.NewButton("📂 Change s&box Path", func() {
		entry := widget.NewEntry()
		entry.SetText(sboxPath)
		entry.SetPlaceHolder("C:\\Program Files (x86)\\Steam\\steamapps\\common\\sbox")
		
		dialog.ShowCustomConfirm("Change s&box Path", "Save", "Cancel",
			container.NewVBox(
				widget.NewLabel("Enter path to s&box installation folder:"),
				widget.NewLabel("(do NOT include \\core\\cfg)"),
				entry,
				widget.NewLabel("Example: C:\\Program Files (x86)\\Steam\\steamapps\\common\\sbox"),
			),
			func(confirmed bool) {
				if confirmed && entry.Text != "" {
					newPath := entry.Text
					// Remove trailing slashes
					newPath = strings.TrimRight(newPath, "\\")
					newPath = strings.TrimRight(newPath, "/")
					
					// Update current path
					sboxPath = newPath
					cfgPath = filepath.Join(sboxPath, "core", "cfg")
					configFilePath = filepath.Join(cfgPath, "graphics_config.vcfg")
					pathValueLabel.SetText(sboxPath)
					
					// Save to settings file
					os.WriteFile(settingsFilePath, []byte(newPath), 0644)
					
					dialog.ShowInformation("Success", 
						fmt.Sprintf("Path updated successfully!\n\nNew config path:\n%s", cfgPath), 
						myWindow)
				}
			}, myWindow)
	})
	
	resetPathButton := widget.NewButton("🔄 Reset to Default", func() {
		sboxPath = defaultSboxPath
		cfgPath = filepath.Join(sboxPath, "core", "cfg")
		configFilePath = filepath.Join(cfgPath, "graphics_config.vcfg")
		pathValueLabel.SetText(sboxPath)
		
		// Remove custom settings file
		os.Remove(settingsFilePath)
		
		dialog.ShowInformation("Success", "Path reset to default!", myWindow)
	})

	pathContainer := container.NewVBox(
		widget.NewLabel("=== s&box Installation Path ==="),
		pathLabel,
		pathValueLabel,
		container.NewHBox(changePathButton, resetPathButton),
		widget.NewLabel(fmt.Sprintf("Config will be saved to: ...\\core\\cfg")),
		widget.NewSeparator(),
	)

	// Create checkboxes for settings
	checkboxes := map[string]*widget.Check{
		
		// === Post-processing ===
		"r_postprocess 0":           widget.NewCheck("Disable post-processing", nil),
		"r_bloom 0":                 widget.NewCheck("Disable bloom (glow)", nil),
		"r_motionblur_scale 0":      widget.NewCheck("Disable motion blur", nil),
		"r_dof_quality 0":           widget.NewCheck("Disable depth of field (DOF)", nil),
		"r_enable_autoexposure 0":   widget.NewCheck("Disable auto-exposure", nil),
		
		// === Shadows ===
		"r_shadows 0":                        widget.NewCheck("Disable all shadows", nil),
		"lb_time_sliced_shadows 0":           widget.NewCheck("Disable time-sliced shadows", nil),
		"lb_indexed_pointlight_shadows 0":    widget.NewCheck("Disable indexed pointlight shadows", nil),
		
		// === Ambient Occlusion ===
		"r_ao_quality 0":            widget.NewCheck("Disable ambient occlusion (SSAO)", nil),
		
		// === Lighting ===
		"r_enable_high_precision_lighting 0": widget.NewCheck("Disable high precision lighting", nil),
		
		// === Reflections ===
		"r_ssr_downsample_ratio 3":  widget.NewCheck("Lower SSR quality (reflections)", nil),
		
		// === Fog ===
		"r_enable_gradient_fog 0":   widget.NewCheck("Disable gradient fog", nil),
		"r_enable_volume_fog 0":     widget.NewCheck("Disable volumetric fog", nil),
		"r_enable_cubemap_fog 0":    widget.NewCheck("Disable cubemap fog", nil),
		"volume_fog_disable 1":      widget.NewCheck("Completely disable volume fog", nil),
		
		// === Decals ===
		"r_render_decals 0":         widget.NewCheck("Disable decals (stickers)", nil),
		"r_gpu_decals 0":            widget.NewCheck("Disable GPU decals", nil),
		
		// === 3D Skybox ===
		"r_3d_skybox 0":             widget.NewCheck("Disable 3D skybox", nil),
		"r_3d_skybox_depth_prepass 0": widget.NewCheck("Disable skybox depth prepass", nil),
		
		// === Textures ===
		"r_texture_stream_mip_bias 1":        widget.NewCheck("Mip bias +1 (less detail)", nil),
		"r_texture_stream_max_resolution 1024": widget.NewCheck("Max texture resolution 1024", nil),
		"r_texture_stream_resolution_bias_min 0.5": widget.NewCheck("Lower min resolution bias", nil),
		
		// === Culling & LOD ===
		"r_size_cull_threshold 0.5":         widget.NewCheck("Increase object culling threshold", nil),
		"r_depth_prepass_cull_threshold 30": widget.NewCheck("Lower depth prepass culling threshold", nil),
		"r_worldlod 0":                      widget.NewCheck("Disable world LOD", nil),
		"sc_bounds_group_cull 0":            widget.NewCheck("Disable bounds group culling", nil),
		
		// === Morphing & Animation ===
		"r_morphing_enabled 0":      widget.NewCheck("Disable morphing", nil),
		"r_allow_morph_batching_on_base 0": widget.NewCheck("Disable morph batching", nil),
		"sc_new_morph_atlasing 0":   widget.NewCheck("Disable new morph atlasing", nil),
		
		// === Skinning (CAUTION!) ===
		"r_skinning_enabled 0":      widget.NewCheck("⚠️ Disable skinning (may break characters!)", nil),
		
		// === Refraction & Transparency ===
		"r_render_refraction 0":     widget.NewCheck("Disable refraction", nil),
		"r_render_translucent 0":    widget.NewCheck("Disable translucent objects", nil),
		"r_translucent 0":           widget.NewCheck("Disable translucent geometry", nil),
		
		// === Overlays ===
		"r_draw_overlays 0":         widget.NewCheck("Disable overlays", nil),
		
		// === VSync & Synchronization ===
		"r_wait_on_present 0":       widget.NewCheck("Don't wait on present", nil),
		
		// === Scene System Optimizations ===
		"sc_disable_shadow_fastpath 1":      widget.NewCheck("Disable shadow fastpath", nil),
		"sc_mesh_backface_culling 1":        widget.NewCheck("Enable backface culling", nil),
		"sc_draw_aggregate_meshes 0":        widget.NewCheck("Disable aggregate meshes", nil),
		
		// === Additional Optimizations ===
		"r_render_dynamic_objects 0":        widget.NewCheck("⚠️ The main menu is empty, but it has infinite FPS", nil),
		"debug_draw_enable 0":               widget.NewCheck("Disable debug draw", nil),
		"mat_disable_normal_mapping 1":      widget.NewCheck("⚠️ Disable normal mapping", nil),
		"vis_sunlight_enable 0":             widget.NewCheck("Disable sunlight visibility", nil),
		
		// === Vulkan Memory (Advanced) ===
		"r_vma_defrag_enabled 0":            widget.NewCheck("Disable VMA defragmentation", nil),
		"vulkan_batch_submits 0":            widget.NewCheck("Disable batch submits", nil),
		
		// === Volume Fog Details ===
		"volume_fog_depth 16":               widget.NewCheck("Lower volume fog depth to 16", nil),
		"volume_fog_height 20":              widget.NewCheck("Lower volume fog height to 20", nil),
		"volume_fog_width 30":               widget.NewCheck("Lower volume fog width to 30", nil),
		
		// === IK & Animation ===
		"ik_enable 0":                       widget.NewCheck("⚠️ Disable IK (inverse kinematics)", nil),
		"animgraph_footlock_enabled 0":      widget.NewCheck("Disable footlock", nil),
	}

	// === FPS Limit (Entry field) ===
	fpsMaxEnabled := widget.NewCheck("Enable FPS Limit", nil)
	fpsMaxEntry := widget.NewEntry()
	fpsMaxEntry.SetPlaceHolder("0 = unlimited, 60, 144, 240...")
	fpsMaxEntry.SetText("144")
	fpsMaxContainer := container.NewVBox(
		fpsMaxEnabled,
		container.NewBorder(nil, nil, widget.NewLabel("FPS Max:"), nil, fpsMaxEntry),
	)

	// === Texture Pool Size (Slider) ===
	texturePoolEnabled := widget.NewCheck("Enable Texture Pool Limit", nil)
	texturePoolSlider := widget.NewSlider(800, 3000)
	texturePoolSlider.SetValue(1600)
	texturePoolSlider.Step = 100
	texturePoolLabel := widget.NewLabel(fmt.Sprintf("Texture Pool Size: %.0f MB", texturePoolSlider.Value))
	texturePoolSlider.OnChanged = func(value float64) {
		texturePoolLabel.SetText(fmt.Sprintf("Texture Pool Size: %.0f MB", value))
	}
	texturePoolContainer := container.NewVBox(
		texturePoolEnabled,
		texturePoolLabel,
		texturePoolSlider,
	)

	// === Texture LOD Scale (Slider) ===
	textureLODEnabled := widget.NewCheck("Enable Texture LOD Scale", nil)
	textureLODSlider := widget.NewSlider(1.0, 4.0)
	textureLODSlider.SetValue(1.0)
	textureLODSlider.Step = 0.5
	textureLODLabel := widget.NewLabel(fmt.Sprintf("Texture LOD Scale: %.1f", textureLODSlider.Value))
	textureLODSlider.OnChanged = func(value float64) {
		textureLODLabel.SetText(fmt.Sprintf("Texture LOD Scale: %.1f", value))
	}
	textureLODContainer := container.NewVBox(
		textureLODEnabled,
		textureLODLabel,
		textureLODSlider,
	)

	// Preset configurations
	mediumPreset := map[string]bool{
		"r_postprocess 0": true,
		"r_bloom 0": true,
		"r_motionblur_scale 0": true,
		"r_ao_quality 0": true,
		"r_ssr_downsample_ratio 3": true,
		"r_3d_skybox_depth_prepass 0": true,
		"debug_draw_enable 0": true,
	}

	lowPreset := map[string]bool{
		"r_postprocess 0": true,
		"r_bloom 0": true,
		"r_motionblur_scale 0": true,
		"r_dof_quality 0": true,
		"r_enable_autoexposure 0": true,
		"r_shadows 0": true,
		"lb_time_sliced_shadows 0": true,
		"lb_indexed_pointlight_shadows 0": true,
		"r_ao_quality 0": true,
		"r_enable_high_precision_lighting 0": true,
		"r_ssr_downsample_ratio 3": true,
		"r_enable_gradient_fog 0": true,
		"r_enable_volume_fog 0": true,
		"r_enable_cubemap_fog 0": true,
		"volume_fog_disable 1": true,
		"r_render_decals 0": true,
		"r_gpu_decals 0": true,
		"r_3d_skybox 0": true,
		"r_3d_skybox_depth_prepass 0": true,
		"r_texture_stream_mip_bias 1": true,
		"r_texture_stream_max_resolution 1024": true,
		"r_texture_stream_resolution_bias_min 0.5": true,
		"r_fallback_texture_lod_scale 4.0": true,
		"r_size_cull_threshold 0.5": true,
		"r_depth_prepass_cull_threshold 30": true,
		"r_worldlod 0": true,
		"sc_bounds_group_cull 0": true,
		"r_morphing_enabled 0": true,
		"r_allow_morph_batching_on_base 0": true,
		"sc_new_morph_atlasing 0": true,
		"r_render_refraction 0": true,
		"r_draw_overlays 0": true,
		"r_wait_on_present 0": true,
		"sc_disable_shadow_fastpath 1": true,
		"sc_mesh_backface_culling 1": true,
		"sc_draw_aggregate_meshes 0": true,
		"debug_draw_enable 0": true,
		"r_vma_defrag_enabled 0": true,
		"vulkan_batch_submits 0": true,
		"volume_fog_depth 16": true,
		"volume_fog_height 20": true,
		"volume_fog_width 30": true,
		"animgraph_footlock_enabled 0": true,
		"mat_disable_normal_mapping 1": true,
		"vis_sunlight_enable 0": true,
		"r_translucent 0": true,
		"r_render_translucent 0": true,
	}

	// Log widget (selectable text)
	logText := widget.NewEntry()
	logText.MultiLine = true
	logText.Wrapping = fyne.TextWrapWord
	logText.SetPlaceHolder("Execution log...")

	// Logging function
	addLog := func(text string) {
		logText.SetText(logText.Text + text + "\n")
	}

	// Apply preset function
	applyPreset := func(preset map[string]bool, presetName string, fps string, texPool float64, texLOD float64) {
		for _, checkbox := range checkboxes {
			checkbox.SetChecked(false)
		}
		
		fpsMaxEnabled.SetChecked(false)
		texturePoolEnabled.SetChecked(false)
		textureLODEnabled.SetChecked(false)
		
		checkedCount := 0
		for cmd, checkbox := range checkboxes {
			if preset[cmd] {
				checkbox.SetChecked(true)
				checkedCount++
			}
		}
		
		if fps != "" {
			fpsMaxEnabled.SetChecked(true)
			fpsMaxEntry.SetText(fps)
			checkedCount++
		}
		
		if texPool > 0 {
			texturePoolEnabled.SetChecked(true)
			texturePoolSlider.SetValue(texPool)
			checkedCount++
		}
		
		if texLOD > 0 {
			textureLODEnabled.SetChecked(true)
			textureLODSlider.SetValue(texLOD)
			checkedCount++
		}
		
		logText.SetText("")
		addLog(fmt.Sprintf("✅ Preset '%s' applied", presetName))
		addLog(fmt.Sprintf("📝 Settings enabled: %d", checkedCount))
	}

	// Load existing settings from file
	loadSettings := func() {
		cfgPath = filepath.Join(sboxPath, "core", "cfg")
		configFilePath = filepath.Join(cfgPath, "graphics_config.vcfg")
		
		if _, err := os.Stat(configFilePath); err == nil {
			data, err := os.ReadFile(configFilePath)
			if err != nil {
				addLog(fmt.Sprintf("⚠️ Error reading config: %v", err))
				return
			}

			lines := strings.Split(string(data), "\n")
			foundCommands := make(map[string]bool)
			
			for _, line := range lines {
				line = strings.TrimSpace(line)
				if line == "" || strings.HasPrefix(line, "//") {
					continue
				}
				foundCommands[line] = true

				if strings.HasPrefix(line, "fps_max ") {
					parts := strings.Split(line, " ")
					if len(parts) == 2 {
						fpsMaxEntry.SetText(parts[1])
						fpsMaxEnabled.SetChecked(true)
					}
				}

				if strings.HasPrefix(line, "r_texture_pool_size ") {
					parts := strings.Split(line, " ")
					if len(parts) == 2 {
						if val, err := strconv.ParseFloat(parts[1], 64); err == nil {
							texturePoolSlider.SetValue(val)
							texturePoolEnabled.SetChecked(true)
						}
					}
				}

				if strings.HasPrefix(line, "r_texture_lod_scale ") {
					parts := strings.Split(line, " ")
					if len(parts) == 2 {
						if val, err := strconv.ParseFloat(parts[1], 64); err == nil {
							textureLODSlider.SetValue(val)
							textureLODEnabled.SetChecked(true)
						}
					}
				}
			}

			loadedCount := 0
			for cmd, checkbox := range checkboxes {
				if foundCommands[cmd] {
					checkbox.SetChecked(true)
					loadedCount++
				}
			}

			addLog(fmt.Sprintf("✅ Settings loaded from file"))
			addLog(fmt.Sprintf("📝 Commands found: %d", loadedCount))
		} else {
			addLog("ℹ️ Settings file not found, using defaults")
		}
	}

	loadSettings()

	// Save settings button
	saveButton := widget.NewButton("💾 Save Settings", func() {
		logText.SetText("")
		addLog("Saving settings...")

		var commands []string

		commands = append(commands, "show_version_overlay 0")

		for cmd, checkbox := range checkboxes {
			if checkbox.Checked {
				commands = append(commands, cmd)
			}
		}

		if fpsMaxEnabled.Checked && fpsMaxEntry.Text != "" {
			commands = append(commands, fmt.Sprintf("fps_max %s", fpsMaxEntry.Text))
		}

		if texturePoolEnabled.Checked {
			commands = append(commands, fmt.Sprintf("r_texture_pool_size %.0f", texturePoolSlider.Value))
		}

		if textureLODEnabled.Checked {
			commands = append(commands, fmt.Sprintf("r_texture_lod_scale %.1f", textureLODSlider.Value))
		}

		if len(commands) == 0 {
			addLog("⚠️ No settings selected!")
			dialog.ShowInformation("Warning", "No settings selected!", myWindow)
			return
		}

		cfgPath = filepath.Join(sboxPath, "core", "cfg")
		configFilePath = filepath.Join(cfgPath, "graphics_config.vcfg")
		
		// Create directories if they don't exist
		os.MkdirAll(cfgPath, 0755)
		
		configContent := "// s&box Graphics Config\n// Generated by s&box Optimizer\n\n"
		configContent += strings.Join(commands, "\n") + "\n"

		err := os.WriteFile(configFilePath, []byte(configContent), 0644)
		if err != nil {
			addLog(fmt.Sprintf("❌ Error creating config: %v", err))
			dialog.ShowError(fmt.Errorf("Error: %v", err), myWindow)
			return
		}

		addLog(fmt.Sprintf("✅ Config saved: %s", configFilePath))
		addLog(fmt.Sprintf("📝 Commands applied: %d", len(commands)))
		addLog("✅ Done! Use launcher.exe to start the game")
		
		dialog.ShowInformation("Success", 
			fmt.Sprintf("Settings saved!\nCommands applied: %d\n\nUse launcher.exe to start s&box", len(commands)), 
			myWindow)
	})

	// Reload settings button
	reloadButton := widget.NewButton("🔄 Reload from File", func() {
		logText.SetText("")
		for _, checkbox := range checkboxes {
			checkbox.SetChecked(false)
		}
		fpsMaxEnabled.SetChecked(false)
		fpsMaxEntry.SetText("144")
		texturePoolEnabled.SetChecked(false)
		texturePoolSlider.SetValue(1600)
		textureLODEnabled.SetChecked(false)
		textureLODSlider.SetValue(1.0)
		loadSettings()
	})

	// Preset buttons
	mediumPresetButton := widget.NewButton("📊 Medium Preset", func() {
		applyPreset(mediumPreset, "Medium", "240", 1400, 2.0)
	})

	lowPresetButton := widget.NewButton("⚡ Low Preset (Max FPS)", func() {
		applyPreset(lowPreset, "Low", "240", 1200, 4.0)
	})

	// Select all button
	selectAllButton := widget.NewButton("✓ Select All", func() {
		for _, checkbox := range checkboxes {
			checkbox.SetChecked(true)
		}
		logText.SetText("")
		addLog("✅ All settings selected")
	})

	// Deselect all button
	deselectAllButton := widget.NewButton("✗ Deselect All", func() {
		for _, checkbox := range checkboxes {
			checkbox.SetChecked(false)
		}
		fpsMaxEnabled.SetChecked(false)
		texturePoolEnabled.SetChecked(false)
		textureLODEnabled.SetChecked(false)
		logText.SetText("")
		addLog("✅ All settings deselected")
	})

	// Create checkbox container
	checkboxContainer := container.NewVBox()
	
	checkboxContainer.Add(widget.NewSeparator())
	checkboxContainer.Add(widget.NewLabel("=== Advanced Settings ==="))
	checkboxContainer.Add(fpsMaxContainer)
	checkboxContainer.Add(texturePoolContainer)
	checkboxContainer.Add(textureLODContainer)
	checkboxContainer.Add(widget.NewSeparator())
	checkboxContainer.Add(widget.NewLabel("=== Graphics Settings ==="))
	
	for _, checkbox := range checkboxes {
		checkboxContainer.Add(checkbox)
	}

	scrollableCheckboxes := container.NewScroll(checkboxContainer)
	scrollableCheckboxes.SetMinSize(fyne.NewSize(520, 350))

	buttonRow1 := container.NewHBox(selectAllButton, deselectAllButton, reloadButton)
	buttonRow2 := container.NewHBox(mediumPresetButton, lowPresetButton)

	scrollableLog := container.NewScroll(logText)
	scrollableLog.SetMinSize(fyne.NewSize(520, 100))

	content := container.NewVBox(
		widget.NewLabel("🎮 Select graphics settings for s&box:"),
		widget.NewSeparator(),
		pathContainer,
		widget.NewLabel("⚙️ Quick Presets:"),
		buttonRow2,
		widget.NewSeparator(),
		scrollableCheckboxes,
		buttonRow1,
		saveButton,
		widget.NewSeparator(),
		widget.NewLabel("📝 Log:"),
		scrollableLog,
	)

	myWindow.SetContent(content)
	myWindow.Resize(fyne.NewSize(600, 870))
	myWindow.ShowAndRun()
}