package authhmac256api

// #cgo pkg-config: libsodium ../../native/lib/pkgconfig/libsodium.pc
// #include "../../native/include/sodium.h"
// #include <stdlib.h>
import "C"

func CryptoAuthHMAC256Bytes() int {
	return int(C.crypto_auth_hmacsha256_bytes())
}

func CryptoAuthHMAC256BKeyBytes() int {
	return int(C.crypto_auth_hmacsha256_keybytes())
}

func CryptoAuthHMAC256StateBytes() int {
	return int(C.crypto_auth_hmacsha256_statebytes())
}
