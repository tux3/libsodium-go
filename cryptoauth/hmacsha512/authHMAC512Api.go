package authhmac512api

// #cgo pkg-config: libsodium ../../native/lib/pkgconfig/libsodium.pc
// #include "../../native/include/sodium.h"
// #include <stdlib.h>
import "C"

func CryptoAuthHMAC512Bytes() int {
	return int(C.crypto_auth_hmacsha512_bytes())
}

func CryptoAuthHMAC512BKeyBytes() int {
	return int(C.crypto_auth_hmacsha512_keybytes())
}

func CryptoAuthHMAC512StateBytes() int {
	return int(C.crypto_auth_hmacsha512_statebytes())
}
