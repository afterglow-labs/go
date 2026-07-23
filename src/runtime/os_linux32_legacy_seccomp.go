// Copyright 2026 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//go:build linux && (386 || arm || mips || mipsle || (gccgo && (ppc || s390))) && android_legacy_seccomp

package runtime

const forceTime32ForLegacyAndroidSeccomp = true
