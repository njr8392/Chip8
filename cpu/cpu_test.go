package cpu

import "testing"

func TestCallStack(t *testing.T) {
	c := Init()
	c.Opcode = 0x2fbb
	c.ParseOP()

	if c.Opcode != 0x0fbb && c.Sp != 1 && c.Stack[0] != 0x2fbb {
		t.Errorf("got opcode %x, sp %d, %x on stack want 0xfbb, 1, and 0x2xfbb", c.Opcode, c.Sp, c.Stack[0])
	}
}

func TestFetch(t *testing.T) {
	c := Init()
	c.Memory[2000] = 0xff
	c.Memory[2001] = 0
	got := c.Fetch(2000)
	want := uint16(0xff00)

	if want != got {
		t.Errorf("Want %x got %x", want, got)
	}

}
