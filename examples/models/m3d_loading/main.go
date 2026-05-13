/*******************************************************************************************
*
*   raylib [models] example - loading m3d
*
*   Example complexity rating: [★★☆☆] 2/4
*
*   Example originally created with raylib 4.5, last time updated with raylib 4.5
*
*   Example contributed by bzt (@bztsrc) and reviewed by Ramon Santamaria (@raysan5)
*
*   NOTES:
*     - Model3D (M3D) fileformat specs: https://gitlab.com/bztsrc/model3d
*     - Bender M3D exported: https://gitlab.com/bztsrc/model3d/-/tree/master/blender
*
*   Example licensed under an unmodified zlib/libpng license, which is an OSI-certified,
*   BSD-like license that allows static linking with closed source software
*
*   Copyright (c) 2022-2025 bzt (@bztsrc)
*
********************************************************************************************/

package main

import (
	"fmt"

	rl "github.com/gen2brain/raylib-go/raylib"
)

// ------------------------------------------------------------------------------------
// Program main entry point
// ------------------------------------------------------------------------------------
func main() {
	// Initialization
	//--------------------------------------------------------------------------------------
	const screenWidth, screenHeight = 800, 450

	rl.InitWindow(screenWidth, screenHeight, "raylib [models] example - loading m3d")
	defer rl.CloseWindow() // Close window and OpenGL context

	// Define the camera to look into our 3d world
	var camera rl.Camera
	camera.Position = rl.Vector3{X: 1.5, Y: 1.5, Z: 1.5} // Camera position
	camera.Target.Y = 0.4                                // Camera looking at point
	camera.Up.Y = 1                                      // Camera up vector (rotation towards target)
	camera.Fovy = 45                                     // Camera field-of-view Y
	camera.Projection = rl.CameraPerspective             // Camera projection type

	// Load model
	model := rl.LoadModel("cesium_man.m3d") // Load the animated model mesh and basic data
	defer rl.UnloadModel(model)             // Unload model
	var position rl.Vector3                 // Set model position

	anims := rl.LoadModelAnimations("cesium_man.m3d") // Load animation data
	defer rl.UnloadModelAnimations(anims)             // Unload model animations data
	animCount := uint32(len(anims))

	// Animation playing variables
	var animIndex uint32         // Current animation playing
	var animCurrentFrame float32 // Current animation frame (supporting interpolated frames)

	rl.SetTargetFPS(60) // Set our game to run at 60 frames-per-second
	//--------------------------------------------------------------------------------------

	// Main game loop
	for !rl.WindowShouldClose() { // Detect window close button or ESC key
		// Update
		//----------------------------------------------------------------------------------
		rl.UpdateCamera(&camera, rl.CameraOrbital)

		// Select current animation
		if rl.IsKeyPressed(rl.KeyRight) {
			animIndex = (animIndex + 1) % animCount
		} else if rl.IsKeyPressed(rl.KeyLeft) {
			animIndex = (animIndex + animCount - 1) % animCount
		}

		// Update model animation
		animCurrentFrame += 1
		if animCurrentFrame >= float32(anims[animIndex].KeyframeCount) {
			animCurrentFrame = 0
		}
		rl.UpdateModelAnimation(model, anims[animIndex], animCurrentFrame)
		//----------------------------------------------------------------------------------

		// Draw
		//----------------------------------------------------------------------------------
		rl.BeginDrawing()

		rl.ClearBackground(rl.RayWhite)

		rl.BeginMode3D(camera)

		// Draw 3d model with texture
		if !rl.IsKeyDown(rl.KeySpace) {
			rl.DrawModel(model, position, 1, rl.White)
		} else {
			// Draw the animated skeleton

			// Loop to (boneCount - 1) because the last one is a special "no bone" bone,
			// needed to workaround buggy models without a -1, a cube is always drawn at the origin
			for i := 0; i < int(model.Skeleton.BoneCount)-1; i++ {
				framePose := anims[animIndex].GetFramePose(int(animCurrentFrame), i).Translation
				// Display the frame-pose skeleton
				rl.DrawCube(framePose, 0.05, 0.05, 0.05, rl.Red)
				if model.Skeleton.GetBones()[i].Parent >= 0 {
					rl.DrawLine3D(framePose, anims[animIndex].GetFramePose(int(animCurrentFrame), int(model.Skeleton.GetBones()[i].Parent)).Translation, rl.Red)
				}
			}
		}

		rl.DrawGrid(10, 1)

		rl.EndMode3D()

		rl.DrawText(fmt.Sprintf("Current animation: %s", anims[animIndex].GetName()), 10, 10, 20, rl.LightGray)
		rl.DrawText("Press SPACE to draw skeleton", 10, 40, 20, rl.Maroon)
		rl.DrawText("(c) CesiumMan model by KhronosGroup", int32(rl.GetScreenWidth()-210), int32(rl.GetScreenHeight()-20), 10, rl.Gray)

		rl.EndDrawing()
		//----------------------------------------------------------------------------------
	}
}
