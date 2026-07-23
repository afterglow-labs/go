// Copyright 2026 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//go:build linux && (386 || arm || mips || mipsle) && android_legacy_seccomp

package runtime_test

import (
	"runtime"
	"testing"
)

func TestLegacyAndroidSeccompUsesTime32Syscalls(t *testing.T) {
	if runtime.LegacyAndroidSeccompUsesTime64ForTest() {
		t.Fatal("legacy Android seccomp build selected time64 syscalls")
	}
}
