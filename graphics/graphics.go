package graphics

import (
	"fmt"
	"github.com/faiface/pixel/pixelgl"
	"github.com/faiface/pixel"
	"github.com/faiface/pixel/imdraw"
	"golang.org/x/image/colornames"
)



type Graphics struct{
	*pixelgl.Window
	//Keys wiill be a map of the original keypad to the keypad and implemented as such
	//Keypad                   Keyboard
	//+-+-+-+-+                +-+-+-+-+
	//|1|2|3|C|                |1|2|3|4|
	//+-+-+-+-+                +-+-+-+-+
	//|4|5|6|D|                |Q|W|E|R|
	//+-+-+-+-+       =>       +-+-+-+-+
	//|7|8|9|E|                |A|S|D|F|
	//+-+-+-+-+                +-+-+-+-+
	//|A|0|B|F|                |Z|X|C|V|
	//+-+-+-+-+                +-+-+-+-+	
	Keys map[uint16]pixelgl.Button
}

func NewWin()(*Graphics, error){
	config := pixelgl.WindowConfig{
		Title: "Chip-8 Emulator",
		Bounds: pixel.R(0, 0, 1024, 1024),
		VSync: true,
	}

	win, err := pixelgl.NewWindow(config)
	if err != nil{
		return nil, fmt.Errorf("Error creating window: %s", err.Error())
	}
	
	kmap := map[uint16]pixelgl.Button{
	0x1: pixelgl.Key1,
	0x2: pixelgl.Key2,
	0x3: pixelgl.Key3,
	0xc: pixelgl.Key4,
	0x4: pixelgl.KeyQ,
	0x5: pixelgl.KeyW,
	0x6: pixelgl.KeyE,
	0xD: pixelgl.KeyR,
	0x7: pixelgl.KeyA,
	0x8: pixelgl.KeyS,
	0x9: pixelgl.KeyD,
	0xe: pixelgl.KeyF,
	0xa: pixelgl.KeyZ,
	0x0: pixelgl.KeyX,
	0xb: pixelgl.KeyC,
	0xF: pixelgl.KeyV,
	}
	return &Graphics{
		Window: win,
		Keys: kmap,
		}, nil
}


func (g *Graphics)Draw(display [2048]byte){
	g.Clear(colornames.Black)
	imd := imdraw.New(nil)
	imd.Color = pixel.RGB(1, 1, 1)
for i := 0; i < 64; i++ {
		for j := 0; j < 32; j++ {
			// If the gfx byte in question is turned off,
			// continue and skip drawing the rectangle
			if display[(31-j)*64+i] == 0 {
				continue
			}
			imd.Push(pixel.V((1024/64)*float64(i), (768/32)*float64(j)))
			imd.Push(pixel.V((10/24/64)*float64(i)+(1024/64), (768/32)*float64(j)+(768/32)))
			imd.Rectangle(0)
		}
	}

	imd.Draw(g)
	g.Update()
}
