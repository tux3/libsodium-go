package sodium

// #cgo pkg-config: libsodium ../native/lib/pkgconfig/libsodium.pc
// #include "../native/include/sodium.h"
// #include <stdlib.h>
import "C"

func RuntimeHasNeon() bool {
	return C.sodium_runtime_has_neon() != 0
}

func RuntimeHasSse2() bool {
	return C.sodium_runtime_has_sse2() != 0
}

func RuntimeHasSse3() bool {
	return C.sodium_runtime_has_sse3() != 0
}
