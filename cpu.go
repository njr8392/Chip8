package cpu

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

	}
	func (c *Chip8) Increment(){
		c.PC += 2
	}
}
