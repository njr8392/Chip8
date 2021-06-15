package main

import (
	"os"
	"log"
	"github.com/njr8392/chip-8/cpu"
	"time"
	"github.com/njr8392/chip-8/graphics"
	"github.com/faiface/pixel/pixelgl"
	)


func run(){
	file := os.Args[1]

	chip8 := cpu.Init()

	if err := chip8.LoadROM(file); err != nil{
		log.Fatal(err)
	}

	display, err := graphics.NewWin()
	if err != nil{
		log.Fatal(err)
	}

	for !display.Closed(){
		chip8.Cycle()
		if chip8.DrawFlag{
			display.Draw(chip8.Graphics)	
		}else{
			display.UpdateInput()
		}
		// Chip8 cpu clock worked at frequency of 60Hz, so set delay to (1000/60)ms
		time.Sleep(1000 / 60)
	}
	
}

func main(){
	pixelgl.Run(run)
}
