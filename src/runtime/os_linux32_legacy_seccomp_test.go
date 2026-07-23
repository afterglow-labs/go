// Copyright 2026 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//go:build linux && (386 || arm || mips || mipsle) && android_legacy_seccomp

package runtime

func LegacyAndroidSeccompUsesTime64ForTest() bool {
	previous := use64bitsTimeOn32bits
	defer func() {
		use64bitsTimeOn32bits = previous
	}()

	use64bitsTimeOn32bits = true
	configure64bitsTimeOn32BitsArchitectures()
	return use64bitsTimeOn32bits
}
