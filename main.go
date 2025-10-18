package main

import (
	"fmt"
	"os"
	"path/filepath"
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

	// Path to s&box cfg folder
	cfgPath := "C:\\Program Files (x86)\\Steam\\steamapps\\common\\sbox\\core\\cfg"
	configFilePath := filepath.Join(cfgPath, "graphics_config.vcfg")

	// Create checkboxes for settings
	checkboxes := map[string]*widget.Check{

		"show_version_overlay 0":    widget.NewCheck("Disable version overlay in corner", nil),

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
		"r_texture_lod_scale 2.0":            widget.NewCheck("Lower texture quality (LOD x2)", nil),
		"r_texture_stream_mip_bias 1":        widget.NewCheck("Mip bias +1 (less detail)", nil),
		"r_texture_pool_size 1200":           widget.NewCheck("Reduce texture pool to 1200MB", nil),
		"r_texture_stream_max_resolution 1024": widget.NewCheck("Max texture resolution 1024", nil),
		"r_texture_stream_resolution_bias_min 0.5": widget.NewCheck("Lower min resolution bias", nil),
		"r_fallback_texture_lod_scale 4.0":   widget.NewCheck("Aggressive LOD for fallback textures", nil),
		
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
		"r_frame_sync_enable 0":     widget.NewCheck("Disable frame sync", nil),
		"r_wait_on_present 0":       widget.NewCheck("Don't wait on present", nil),
		
		// === FPS Limit ===
		"fps_max 60":                widget.NewCheck("Limit FPS to 60", nil),
		"fps_max 144":               widget.NewCheck("Limit FPS to 144", nil),
		"fps_max 240":               widget.NewCheck("Limit FPS to 240", nil),
		"fps_max 0":                 widget.NewCheck("Remove FPS limit", nil),
		
		// === Scene System Optimizations ===
		"sc_disable_shadow_fastpath 1":      widget.NewCheck("Disable shadow fastpath", nil),
		"sc_mesh_backface_culling 1":        widget.NewCheck("Enable backface culling", nil),
		"sc_draw_aggregate_meshes 0":        widget.NewCheck("Disable aggregate meshes", nil),
		
		// === Additional Optimizations ===
		"r_render_dynamic_objects 0":        widget.NewCheck("⚠️ Disable dynamic objects", nil),
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
		
		// === IK & Animation (Extreme Optimization) ===
		"ik_enable 0":                       widget.NewCheck("⚠️ Disable IK (inverse kinematics)", nil),
		"animgraph_footlock_enabled 0":      widget.NewCheck("Disable footlock", nil),
	}

	// Log widget
	logText := widget.NewMultiLineEntry()
	logText.SetPlaceHolder("Execution log...")
	logText.Disable()

	// Logging function
	addLog := func(text string) {
		logText.SetText(logText.Text + text + "\n")
	}

	// Load existing settings from file
	loadSettings := func() {
		if _, err := os.Stat(configFilePath); err == nil {
			// File exists, read it
			data, err := os.ReadFile(configFilePath)
			if err != nil {
				addLog(fmt.Sprintf("⚠️ Error reading config: %v", err))
				return
			}

			// Parse file
			lines := strings.Split(string(data), "\n")
			foundCommands := make(map[string]bool)
			
			for _, line := range lines {
				line = strings.TrimSpace(line)
				// Skip comments and empty lines
				if line == "" || strings.HasPrefix(line, "//") {
					continue
				}
				foundCommands[line] = true
			}

			// Set checkboxes according to found commands
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

	// Load settings on startup
	loadSettings()

	// Save settings button
	saveButton := widget.NewButton("💾 Save Settings", func() {
		logText.SetText("")
		addLog("Saving settings...")

		// Collect selected commands
		var commands []string
		for cmd, checkbox := range checkboxes {
			if checkbox.Checked {
				commands = append(commands, cmd)
			}
		}

		if len(commands) == 0 {
			addLog("⚠️ No settings selected!")
			dialog.ShowInformation("Warning", "No settings selected!", myWindow)
			return
		}

		// Create config content
		configContent := "// s&box Graphics Config\n// Generated by s&box Optimizer\n\n"
		configContent += strings.Join(commands, "\n") + "\n"

		// Create config file
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
		// Reset all checkboxes
		for _, checkbox := range checkboxes {
			checkbox.SetChecked(false)
		}
		// Load again
		loadSettings()
	})

	// Select all button
	selectAllButton := widget.NewButton("✓ Select All", func() {
		for _, checkbox := range checkboxes {
			checkbox.SetChecked(true)
		}
	})

	// Deselect all button
	deselectAllButton := widget.NewButton("✗ Deselect All", func() {
		for _, checkbox := range checkboxes {
			checkbox.SetChecked(false)
		}
	})

	// Create checkbox container
	checkboxContainer := container.NewVBox()
	for _, checkbox := range checkboxes {
		checkboxContainer.Add(checkbox)
	}

	// IMPORTANT: Wrap checkboxes in Scroll
	scrollableCheckboxes := container.NewScroll(checkboxContainer)
	scrollableCheckboxes.SetMinSize(fyne.NewSize(480, 400))

	// Control buttons
	buttonRow := container.NewHBox(selectAllButton, deselectAllButton, reloadButton)

	// Log also in Scroll with fixed height
	scrollableLog := container.NewScroll(logText)
	scrollableLog.SetMinSize(fyne.NewSize(480, 100))

	// Main container
	content := container.NewVBox(
		widget.NewLabel("🎮 Select graphics settings for s&box:"),
		widget.NewSeparator(),
		scrollableCheckboxes,
		buttonRow,
		saveButton,
		widget.NewSeparator(),
		widget.NewLabel("📝 Log:"),
		scrollableLog,
	)

	myWindow.SetContent(content)
	myWindow.Resize(fyne.NewSize(550, 750))
	myWindow.ShowAndRun()
}