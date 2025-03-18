package cpu

type Registers struct {
	a uint8
	b uint8
	c uint8
	d uint8
	e uint8
	f FlagsRegister
	h uint8
	l uint8
}

func (e Registers) get_af() uint16 {
	return uint16(e.a)<<8 | uint16(e.f.toUint8())
}

func (e *Registers) set_af(value uint16) {
	e.a = uint8((value & 0xFF00) >> 8)
	e.f.fromUint8(uint8(value & 0xFF))
}

func (e Registers) get_bc() uint16 {
	return uint16(e.b)<<8 | uint16(e.c)
}

func (e *Registers) set_bc(value uint16) {
	e.b = uint8((value & 0xFF00) >> 8)
	e.c = uint8(value & 0xFF)
}
func (e Registers) get_de() uint16 {
	return uint16(e.d)<<8 | uint16(e.e)
}

func (e *Registers) set_de(value uint16) {
	e.d = uint8((value & 0xFF00) >> 8)
	e.e = uint8(value & 0xFF)
}

func (e Registers) get_hl() uint16 {
	return uint16(e.h)<<8 | uint16(e.l)
}

func (e *Registers) set_hl(value uint16) {
	e.h = uint8((value & 0xFF00) >> 8)
	e.l = uint8(value & 0xFF)
}
