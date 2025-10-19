# s&box Graphics Optimizer

LonePine discord Server: https://discord.gg/fg4KaeGbjm

---

## 🎮 Usage

open sbox_optimizer.exe -> 
set optinos -> 
click SaveSettings button -> 
run launcher.exe ->
enjoy optimized s&box!

---

## 🔧 All Settings

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


## 🔬 How It Works

1. **sbox_optimizer.exe:**
   - Reads existing config (if present)
   - Displays GUI with checkboxes and sliders
   - Creates `graphics_config.vcfg` file on "Save"
   - File format: plain text with console commands

2. **graphics_config.vcfg:**
```
   // s&box Graphics Config
   // Generated by s&box Optimizer
   
   r_postprocess 0
   r_bloom 0
   fps_max 144
   r_texture_pool_size 1400
```
   - Each line = Source Engine console command
   - File is read by the engine at startup

3. **launcher.exe:**
   - Finds Steam.exe
   - Runs command: `steam.exe -applaunch 590830 +exec graphics_config.vcfg`
   - `-applaunch 590830` = launch s&box (App ID)
   - `+exec` = execute config at startup

4. **s&box:**
   - Executes commands from `graphics_config.vcfg` at startup
   - Applies settings to the engine
   - Settings persist until game restart
   
---



## 📊 Performance

### Test System
- CPU: Ryzen 5-3500x
- GPU: GTX 1060 3gb
- RAM: 16GB DDR4
- Resolution: 1920x1080

### Results

| Setting | FPS (average)  | Gain |
|---------|---------------|------|
| Default | 40 FPS | - |
| Medium Preset | 80 FPS | +100% |
| Low Preset | 140 FPS | +200% |

*Results may vary depending on hardware and map*

---

## 🤝 Contributing

Pull requests are welcome! For major changes, please open an issue first to discuss what you would like to change.

### How to help the project:
1. 🐛 Report bugs via Issues
2. 💡 Suggest new features
5. ⭐ Star the project on GitHub!

---

## 📜 License

do with this whatever you want. who cares about MIT actually lol

---

**Made with ❤️ for the s&box community**