package hmachmacsha512

// #cgo pkg-config: libsodium ../../../native/lib/pkgconfig/libsodium.pc
// #include "../../../native/include/sodium.h"
// #include <stdlib.h>
import "C"

func CryptoAuthHMACSHA512Init(state *C.struct_crypto_auth_hmacsha512_state, key []byte, keylen int) (*C.struct_crypto_auth_hmacsha512_state, int) {
	exit := int(C.crypto_auth_hmacsha512_init(
		(state),
		(*C.uchar)(&key[0]),
		(C.size_t)(keylen)))

	return state, exit

}
