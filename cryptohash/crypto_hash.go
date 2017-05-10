package cryptohash

// #cgo pkg-config: libsodium ../native/lib/pkgconfig/libsodium.pc
// #include "../native/include/sodium.h"
// #include <stdlib.h>
import "C"

func CryptoHashBytes() int {
	return int(C.crypto_hash_bytes())
}

func CryptoHashPrimitive() string {
	return C.GoString(C.crypto_hash_primitive())
}

func CryptoHash(in []byte) ([]byte, int) {
	out := make([]byte, CryptoHashBytes())
	exit := int(C.crypto_hash(
		(*C.uchar)(&out[0]),
		(*C.uchar)(&in[0]),
		(C.ulonglong)(len(in))))

	return out, exit
}
