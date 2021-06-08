package main

import (
	"fmt"
	"os"
	"log"
	"github.com/njr8392/chip-8/cpu"
	"github.com/veandco/go-sdl2/sdl"
	)

const (
	WIDTH  = 64
	HEIGHT = 32
)

func main(){
	file := os.Args[1]
	var modify int32 =15

	chip8 := cpu.Init()

	if err := chip8.LoadROM(file); err != nil{
		log.Fatal(err)
	}
if sdlErr := sdl.Init(sdl.INIT_EVERYTHING); sdlErr != nil {
		panic(sdlErr)
	}
	defer sdl.Quit()	

// Create window, chip8 resolution with given modifier
	window, windowErr := sdl.CreateWindow(file, sdl.WINDOWPOS_UNDEFINED, sdl.WINDOWPOS_UNDEFINED, WIDTH * modify, HEIGHT * modify	, sdl.WINDOW_SHOWN)
	if windowErr != nil {
		panic(windowErr)
	}
	defer window.Destroy()

	// Create render surface
	canvas, canvasErr := sdl.CreateRenderer(window, -1, 0)
	if canvasErr != nil {
		panic(canvasErr)
	}
	defer canvas.Destroy()

	for {
		// Compute the next opcode
		chip8.Cycle()
		// Draw only if required
		if chip8.Draw() {
			// Clear the screen
			canvas.SetDrawColor(255, 0, 0, 255)
			canvas.Clear()

	for i := 0; i < 64; i++ {
		for j := 0; j < 32; j++ {
			fmt.Println(31-j*64+i)
			if chip8.Graphics[(31-j)*64+i] != 0 {   ///This is where it is breaking.  -33 with i =1
				canvas.SetDrawColor(255, 255, 0, 255)
			} else {
				canvas.SetDrawColor(255, 0, 0, 255)
			}
			canvas.FillRect(&sdl.Rect{
				Y: int32(j) * modify,
				X: int32(i) * modify,
				W: modify,
				H: modify,
			})
		}
	}

			canvas.Present()
		}

		// Poll for Quit and ChangeKeyboard events
		for event := sdl.PollEvent(); event != nil; event = sdl.PollEvent() {
			switch et := event.(type) {
			case *sdl.QuitEvent:
				os.Exit(0)
			case *sdl.KeyboardEvent:
				if et.Type == sdl.KEYUP {
					switch et.Keysym.Sym {
					case sdl.K_1:
						chip8.ChangeKey(0x1, false)
					case sdl.K_2:
						chip8.ChangeKey(0x2, false)
					case sdl.K_3:
						chip8.ChangeKey(0x3, false)
					case sdl.K_4:
						chip8.ChangeKey(0xC, false)
					case sdl.K_q:
						chip8.ChangeKey(0x4, false)
					case sdl.K_w:
						chip8.ChangeKey(0x5, false)
					case sdl.K_e:
						chip8.ChangeKey(0x6, false)
					case sdl.K_r:
						chip8.ChangeKey(0xD, false)
					case sdl.K_a:
						chip8.ChangeKey(0x7, false)
					case sdl.K_s:
						chip8.ChangeKey(0x8, false)
					case sdl.K_d:
						chip8.ChangeKey(0x9, false)
					case sdl.K_f:
						chip8.ChangeKey(0xE, false)
					case sdl.K_z:
						chip8.ChangeKey(0xA, false)
					case sdl.K_x:
						chip8.ChangeKey(0x0, false)
					case sdl.K_c:
						chip8.ChangeKey(0xB, false)
					case sdl.K_v:
						chip8.ChangeKey(0xF, false)
					}
				} else if et.Type == sdl.KEYDOWN {
					switch et.Keysym.Sym {
					case sdl.K_1:
						chip8.ChangeKey(0x1, true)
					case sdl.K_2:
						chip8.ChangeKey(0x2, true)
					case sdl.K_3:
						chip8.ChangeKey(0x3, true)
					case sdl.K_4:
						chip8.ChangeKey(0xC, true)
					case sdl.K_q:
						chip8.ChangeKey(0x4, true)
					case sdl.K_w:
						chip8.ChangeKey(0x5, true)
					case sdl.K_e:
						chip8.ChangeKey(0x6, true)
					case sdl.K_r:
						chip8.ChangeKey(0xD, true)
					case sdl.K_a:
						chip8.ChangeKey(0x7, true)
					case sdl.K_s:
						chip8.ChangeKey(0x8, true)
					case sdl.K_d:
						chip8.ChangeKey(0x9, true)
					case sdl.K_f:
						chip8.ChangeKey(0xE, true)
					case sdl.K_z:
						chip8.ChangeKey(0xA, true)
					case sdl.K_x:
						chip8.ChangeKey(0x0, true)
					case sdl.K_c:
						chip8.ChangeKey(0xB, true)
					case sdl.K_v:
						chip8.ChangeKey(0xF, true)
					}
				}
			}
		}

		// Chip8 cpu clock worked at frequency of 60Hz, so set delay to (1000/60)ms
		sdl.Delay(1000 / 60)
	}
}
	
