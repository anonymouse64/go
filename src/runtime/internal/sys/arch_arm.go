// Copyright 2014 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package sys

const (
	ArchFamily          = ARM
	BigEndian           = false
	CacheLineSize       = 32
	DefaultPhysPageSize = 65536
	PCQuantum           = 4
	Int64Align          = 4
	HugePageSize        = 0
	MinFrameSize        = 4
)

type Uintreg uint32
