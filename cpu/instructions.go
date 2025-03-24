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
	{0x80, "ADD A,rB", 1, 1, func(cpu *CPU) { cpu.AddA_r(cpu.registers.b) }},
	{0x81, "ADD A,rC", 1, 1, func(cpu *CPU) { cpu.AddA_r(cpu.registers.c) }},
	{0x82, "ADD A,rD", 1, 1, func(cpu *CPU) { cpu.AddA_r(cpu.registers.d) }},
	{0x83, "ADD A,rE", 1, 1, func(cpu *CPU) { cpu.AddA_r(cpu.registers.e) }},
	{0x84, "ADD A,rH", 1, 1, func(cpu *CPU) { cpu.AddA_r(cpu.registers.h) }},
	{0x85, "ADD A,rL", 1, 1, func(cpu *CPU) { cpu.AddA_r(cpu.registers.l) }},
	{0x87, "ADD A,rA", 1, 1, func(cpu *CPU) { cpu.AddA_r(cpu.registers.a) }},
}
