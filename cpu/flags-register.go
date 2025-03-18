package cpu

const ZERO_FLAG_BYTE_POSITION uint8 = 7
const SUBTRACT_FLAG_BYTE_POSITION uint8 = 6
const HALF_CARRY_FLAG_BYTE_POSITION uint8 = 5
const CARRY_FLAG_BYTE_POSITION uint8 = 4

type FlagsRegister struct {
	zero       bool
	subtract   bool
	half_carry bool
	carry      bool
}

func (e FlagsRegister) toUint8() uint8 {
	var flag uint8

	if e.zero {
		flag = flag | (1 << ZERO_FLAG_BYTE_POSITION)
	}

	if e.subtract {
		flag = flag | (1 << SUBTRACT_FLAG_BYTE_POSITION)
	}

	if e.half_carry {
		flag = flag | (1 << HALF_CARRY_FLAG_BYTE_POSITION)
	}

	if e.carry {
		flag = flag | (1 << CARRY_FLAG_BYTE_POSITION)
	}

	return flag
}

func (e *FlagsRegister) fromUint8(flag uint8) {
	e.zero = ((flag >> ZERO_FLAG_BYTE_POSITION) & 0b1) != 0
	e.subtract = ((flag >> SUBTRACT_FLAG_BYTE_POSITION) & 0b1) != 0
	e.half_carry = ((flag >> HALF_CARRY_FLAG_BYTE_POSITION) & 0b1) != 0
	e.carry = ((flag >> CARRY_FLAG_BYTE_POSITION) & 0b1) != 0
}
