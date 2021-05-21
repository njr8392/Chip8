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
func(c *Chip8) Fetch(index uint16)uint16{
	return c.Memory[index] << 8 | c.Memory[index+1]
}

//need to implement
func(c *Chip8) Decode(){
}



func(c *Chip8) Init(){
	c.PC = 0x200 // program counter starts at this address
	/*c.Opcode = 0
	c.Index = 0
	c.Sp = 0 */

}
