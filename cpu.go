package cpu

import
	("fmt"
	"math"
	)

type Chip8 struct {
	Opcode   uint16
	Memory   [4096]byte //Ram
	Register [16]byte   //chip8 has 16 8 bit registers
	Index    uint16     // Index Register
	PC       uint16     // program counter

	//Graphic system for the chip8
	//Pixels are either on or off and the screen has a total of 2048 pixels

	Graphics [62 * 32]byte

	Delay      byte
	SoundTimer byte

	//need to implenent a stack as well as a stack point to jump to an address or call a subroutine
	Stack [16]uint16
	Sp    uint16

	//Hex based keypad to store the current state of the key
	Key [16]uint16
}


//Fetch two bytes (one opcode) from memory and return the opcode
func(c *Chip8) Fetch(index uint8)uint16{
	return uint16(c.Memory[index]) << 8 | uint16(c.Memory[index+1])
}


func(c *Chip8) Init(){
	c.PC = 0x200 // program counter starts at this address
	/*c.Opcode = 0
	c.Index = 0
	c.Sp = 0 */

}

func(c *Chip8) Cycle(){
	c.Opcde = c.Fetch(c.pc)
	x := c.Opcode & 0x0F00
	y := c.Opcode & 0x00F0

	switch c.Opcode & 0xF000{

		//jumps to address NNN
		case 0x1000:
			c.PC = c.Opcode & 0x0FFF

		//calls subroutine at NNN
		case 0x2000:
			c.Stack[c.Sp] =c.PC
			c.Sp += 1
			c.PC = c.Opcode & 0x0FFF
		
		case 0x3000:
			// skips the next instruction if the  register does not equal the last 8 bits in the opcode
			if uint16([c.Register[(c.Opcode & 0x0F00) >> 8]) == c.Opcode & 0x00ff {
				c.PC += 4
			} else{
				c.PC += 2
			}

		case 0x4000:
			// skips the next instruction if the  register does not equal the last 8 bits in the opcode
			if uint16([c.Register[(c.Opcode & 0x0F00) >> 8]) != c.Opcode & 0x00ff {
				c.PC += 4
			} else{
				c.PC += 2
			}

		case 0x5000:
			x := c.Opcode & 0x0F00
			y := c.Opcode & 0x00F0

			if c.Register[x >> 8] == c.Register[y >> 4]{
				c.PC += 4
			} else {
				c.PC += 2
			}

		case 0x6000:
			x := c.Opcode & 0x0F00
			c.Register[x >>8] == uint8(c.Opcode & 0x00FF)
			c.Increment()

		case 0x7000:
			c.Register[c.Opcode & 0x0F00 >> 8] += byte(c.Opcode & 0x00FF)
			c.Increment()
			
		case 0x8000:
			switch c.Opcode & 0x000F{
				
				case 0x0000:
					//Store value of Vy into Vx
					c.Register[x] = c.Register[y]
					c.Increment()

				case 0x0001:
					//Vx = Vx OR Vy
					c.Register[x] = c.Register[x] | c.Register[y]
					c,Increment()

				case 0x0002:
					//Vx = Vx AND Vy
					c.Register[x] = c.Register[x] & c.Register[y]
					c.Increment()

				case 0x0003:
					//Vx = Vx XOR Vy
					c.Register[x] = c.Register[x] ^ c.Register[y]
					c.Increment()

				case 0x0004:
					// if Vx + Vy > 255, then set regiters accordingly and store lowest 8 bits in Vx
					c.Register[0xF] = 0

					if c.Register[y] > (0xFF - c.Register[x]){
						c.Register[0xF] = 1
					} 

					c.Register[x] += c.Register[y]
					c.Increment()

				case 0x0005:
					// set VF to 1 if a carry occurs and subtract Vy from Vx
					c.Register[0xF] = 0

					if c.Register[x] > c.Register[y]{
						c.Register[0xF] = 1
					} 
					c.Register[x] -= c.Register[y]
					c.Increment()


				case 0x0006:
					//Set Register F  to the least significant bit of Register x then divide Register x by 2
					c.Register[0xF] = c.Register[x] & 0x1
					c.Register[x] >> 1
					c.Increment()

				case 0x0007:
				// Set Register x to RegY - RegX and Register y to based on the comparison of Regy > Regx
					c.Register[0xF] = 0
					if c.Register[y] > c.Register[x]{
						c.Register[0xF] = 1
					}

					c.Register[x] = c.Register[y] - c.Register[x]
					c.Increment()
				
				case 0x00E:
					// Set Register F to the most significant bit of x and double Register of x
					c.Register[0xF] = c.Register[x] >> 7 
					C.Register[x] << 1
					c.Increment()


				default: 
					fmt.Printf("Unrecognizable opcode %x\n", c.Opcode)
			}

		case 0x9000:
			//Skip on Regx != Regy
			c.Increment()
			if c.Register[x] != c.Register[y]{
				c.Increment()
			}

		case 0xA000: 
			//set index register to the last 12 bits
			c.Index = c.Opcode & 0x0FFF
			c.Increment()

		case 0xB000:
			// 0xBNNN jumps to address NNN plus Register0
			c.PC = c.Opcode & 0x0FFF + uint16(Register[0])

		case 0xC000:
			c.Register[x] = uint8(math.Intn(256)) & uint8(c.Opcode &  0x00FF)
			c.Increment()

		//case 0xD000:

	}
}

func (c *Chip8) Increment(){
	c.PC += 2
}
