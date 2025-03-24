package cpu

type CPU struct {
	registers Registers
	pc        uint16
	bus       MemoryBus
}

type MemoryBus struct {
	memory [0xFFFF]uint8
}

// CPU

func (e *CPU) step() {
	instruction_byte := e.bus.read_byte(e.pc)

}

func (e *CPU) AddA_r(value uint8) {
	e.registers.a = e.add(value)
}

func (e *CPU) add(value uint8) uint8 {
	new_value := e.registers.a + value
	did_overflow := new_value < value

	e.registers.f.zero = new_value == 0
	e.registers.f.subtract = false
	e.registers.f.carry = did_overflow
	e.registers.f.half_carry = (e.registers.a&0xF)+(value&0xF) > 0xF

	return new_value
}

// MemoryBus

func (e *MemoryBus) read_byte(address uint16) uint8 {
	return e.memory[address]
}
