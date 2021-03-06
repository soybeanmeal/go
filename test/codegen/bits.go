// asmcheck

// Copyright 2018 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package codegen

/************************************
 * 64-bit instructions
 ************************************/

func bitcheck64_constleft(a uint64) (n int) {
	// amd64:"BTQ\t[$]63"
	if a&(1<<63) != 0 {
		return 1
	}
	// amd64:"BTQ\t[$]60"
	if a&(1<<60) != 0 {
		return 1
	}
	// amd64:"BTL\t[$]0"
	if a&(1<<0) != 0 {
		return 1
	}
	return 0
}

func bitcheck64_constright(a [8]uint64) (n int) {
	// amd64:"BTQ\t[$]63"
	if (a[0]>>63)&1 != 0 {
		return 1
	}
	// amd64:"BTQ\t[$]63"
	if a[1]>>63 != 0 {
		return 1
	}
	// amd64:"BTQ\t[$]63"
	if a[2]>>63 == 0 {
		return 1
	}
	// amd64:"BTQ\t[$]60"
	if (a[3]>>60)&1 == 0 {
		return 1
	}
	// amd64:"BTL\t[$]1"
	if (a[4]>>1)&1 == 0 {
		return 1
	}
	// amd64:"BTL\t[$]0"
	if (a[5]>>0)&1 == 0 {
		return 1
	}
	// amd64:"BTL\t[$]7"
	if (a[6]>>5)&4 == 0 {
		return 1
	}
	return 0
}

func bitcheck64_var(a, b uint64) (n int) {
	// amd64:"BTQ"
	if a&(1<<(b&63)) != 0 {
		return 1
	}
	// amd64:"BTQ",-"BT.\t[$]0"
	if (b>>(a&63))&1 != 0 {
		return 1
	}
	return 0
}

func bitcheck64_mask(a uint64) (n int) {
	// amd64:"BTQ\t[$]63"
	if a&0x8000000000000000 != 0 {
		return 1
	}
	// amd64:"BTQ\t[$]59"
	if a&0x800000000000000 != 0 {
		return 1
	}
	// amd64:"BTL\t[$]0"
	if a&0x1 != 0 {
		return 1
	}
	return 0
}

func biton64(a, b uint64) (n uint64) {
	// amd64:"BTSQ"
	n += b | (1 << (a & 63))

	// amd64:"BTSQ\t[$]63"
	n += a | (1 << 63)

	// amd64:"BTSQ\t[$]60"
	n += a | (1 << 60)

	// amd64:"ORQ\t[$]1"
	n += a | (1 << 0)

	return n
}

func bitoff64(a, b uint64) (n uint64) {
	// amd64:"BTRQ"
	n += b &^ (1 << (a & 63))

	// amd64:"BTRQ\t[$]63"
	n += a &^ (1 << 63)

	// amd64:"BTRQ\t[$]60"
	n += a &^ (1 << 60)

	// amd64:"ANDQ\t[$]-2"
	n += a &^ (1 << 0)

	return n
}

func bitcompl64(a, b uint64) (n uint64) {
	// amd64:"BTCQ"
	n += b ^ (1 << (a & 63))

	// amd64:"BTCQ\t[$]63"
	n += a ^ (1 << 63)

	// amd64:"BTCQ\t[$]60"
	n += a ^ (1 << 60)

	// amd64:"XORQ\t[$]1"
	n += a ^ (1 << 0)

	return n
}

/************************************
 * 32-bit instructions
 ************************************/

func bitcheck32_constleft(a uint32) (n int) {
	// amd64:"BTL\t[$]31"
	if a&(1<<31) != 0 {
		return 1
	}
	// amd64:"BTL\t[$]28"
	if a&(1<<28) != 0 {
		return 1
	}
	// amd64:"BTL\t[$]0"
	if a&(1<<0) != 0 {
		return 1
	}
	return 0
}

func bitcheck32_constright(a [8]uint32) (n int) {
	// amd64:"BTL\t[$]31"
	if (a[0]>>31)&1 != 0 {
		return 1
	}
	// amd64:"BTL\t[$]31"
	if a[1]>>31 != 0 {
		return 1
	}
	// amd64:"BTL\t[$]31"
	if a[2]>>31 == 0 {
		return 1
	}
	// amd64:"BTL\t[$]28"
	if (a[3]>>28)&1 == 0 {
		return 1
	}
	// amd64:"BTL\t[$]1"
	if (a[4]>>1)&1 == 0 {
		return 1
	}
	// amd64:"BTL\t[$]0"
	if (a[5]>>0)&1 == 0 {
		return 1
	}
	// amd64:"BTL\t[$]7"
	if (a[6]>>5)&4 == 0 {
		return 1
	}
	return 0
}

func bitcheck32_var(a, b uint32) (n int) {
	// amd64:"BTL"
	if a&(1<<(b&31)) != 0 {
		return 1
	}
	// amd64:"BTL",-"BT.\t[$]0"
	if (b>>(a&31))&1 != 0 {
		return 1
	}
	return 0
}

func bitcheck32_mask(a uint32) (n int) {
	// amd64:"BTL\t[$]31"
	if a&0x80000000 != 0 {
		return 1
	}
	// amd64:"BTL\t[$]27"
	if a&0x8000000 != 0 {
		return 1
	}
	// amd64:"BTL\t[$]0"
	if a&0x1 != 0 {
		return 1
	}
	return 0
}

func biton32(a, b uint32) (n uint32) {
	// amd64:"BTSL"
	n += b | (1 << (a & 31))

	// amd64:"BTSL\t[$]31"
	n += a | (1 << 31)

	// amd64:"BTSL\t[$]28"
	n += a | (1 << 28)

	// amd64:"ORL\t[$]1"
	n += a | (1 << 0)

	return n
}

func bitoff32(a, b uint32) (n uint32) {
	// amd64:"BTRL"
	n += b &^ (1 << (a & 31))

	// amd64:"BTRL\t[$]31"
	n += a &^ (1 << 31)

	// amd64:"BTRL\t[$]28"
	n += a &^ (1 << 28)

	// amd64:"ANDL\t[$]-2"
	n += a &^ (1 << 0)

	return n
}

func bitcompl32(a, b uint32) (n uint32) {
	// amd64:"BTCL"
	n += b ^ (1 << (a & 31))

	// amd64:"BTCL\t[$]31"
	n += a ^ (1 << 31)

	// amd64:"BTCL\t[$]28"
	n += a ^ (1 << 28)

	// amd64:"XORL\t[$]1"
	n += a ^ (1 << 0)

	return n
}
