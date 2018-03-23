// Copyright 2017 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// +build arm

package cpu

const CacheLineSize = 32

// arm doesn't have a 'cpuid' equivalent, so we rely on HWCAP/HWCAP2.
// These are linknamed in runtime/os_linux_arm.go and are initialized by
// archauxv().
var arm_hwcap uint

// HWCAP bits. These are exposed by Linux.
const (
	hwcap_SWP       = (1 << 0)
	hwcap_HALF      = (1 << 1)
	hwcap_THUMB     = (1 << 2)
	hwcap_26BIT     = (1 << 3)
	hwcap_FAST_MULT = (1 << 4)
	hwcap_FPA       = (1 << 5)
	hwcap_VFP       = (1 << 6)
	hwcap_EDSP      = (1 << 7)
	hwcap_JAVA      = (1 << 8)
	hwcap_IWMMXT    = (1 << 9)
	hwcap_CRUNCH    = (1 << 10)
	hwcap_THUMBEE   = (1 << 11)
	hwcap_NEON      = (1 << 12)
	hwcap_VFPv3     = (1 << 13)
	hwcap_VFPv3D16  = (1 << 14)
	hwcap_TLS       = (1 << 15)
	hwcap_VFPv4     = (1 << 16)
	hwcap_IDIVA     = (1 << 17)
	hwcap_IDIVT     = (1 << 18)
	hwcap_VFPD32    = (1 << 19)
	hwcap_IDIV      = (hwcap_IDIVA | hwcap_IDIVT)
)

func init() {
	// HWCAP feature bits
	ARM.HasNEON = isSet(arm_hwcap, hwcap_NEON)
}

func isSet(hwc uint, value uint) bool {
	return hwc&value != 0
}
