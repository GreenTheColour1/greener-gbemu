package cpu

type Instruction struct {
	Opcode      byte
	Description string
	Cycles      uint8
	Bytes       uint8
	Execute     func(cpu *CPU)
}

var Instructions []*Instruction = []*Instruction{
	// Arithmetic Instructions
	&Instruction{0x80, "ADD A,rB", 1, 1, func(cpu *CPU) { cpu.AddA_r(cpu.registers.b) }},
}
