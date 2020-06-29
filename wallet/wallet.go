package wallet

import (
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"crypto/sha256"
	"fmt"

	"github.com/btcsuite/btcutil/base58"
	"golang.org/x/crypto/ripemd160"
)

// Memo:
// ブロックチェーンアドレスの作り方
// https://en.bitcoin.it/wiki/Technical_background_of_version_1_Bitcoin_addresses

// Summary
// 1. Creating ECDSA private key (32bytes), public key (64bytes)
// 2. Perform SHA-256 hashing on the public key (32 bytes)
// 3. Perform RIPEMD-160 hashing on the result of SHA-256 (20bytes)
// 4. Add version byte in front of RIPEMD-160 hash (0x00 for Main Network)
// 5. Perform SHA-256 hash on the exteneded RIPEMD-160 result
// 6. Perform SHA-256 hash on the result of the previous SHA-256 hash
// 7. Take the first 5 bytes of the second SHA-256 hash for checksum
// 8. Add the 4 checksum bytes from 7 at the end of extended RIPEMD-160 hash from 4 (25bytes)
// 9. Convert the result from a byte string into base58

type wallet struct {
	privateKey        *ecdsa.PrivateKey
	publicKey         *ecdsa.PublicKey
	blockchainAddress string
}

func NewWallet() *wallet {
	// 1. Creating ECDSA private key (32bytes), public key (64bytes)
	w := new(wallet)
	privateKey, _ := ecdsa.GenerateKey(elliptic.P256(), rand.Reader)
	w.privateKey = privateKey
	w.publicKey = &w.privateKey.PublicKey

	// 2. Perform SHA-256 hashing on the public key (32 bytes)
	hash2 := sha256.New()
	hash2.Write(w.publicKey.X.Bytes())
	hash2.Write(w.publicKey.Y.Bytes())
	digest2 := hash2.Sum(nil)

	// 3. Perform RIPEMD-160 hashing on the result of SHA-256 (20bytes)
	hash3 := ripemd160.New()
	hash3.Write(digest2)
	digest3 := hash3.Sum(nil)

	// 4. Add version byte in front of RIPEMD-160 hash (0x00 for Main Network)
	vd4 := make([]byte, 21)
	vd4[0] = 0x00
	copy(vd4[1:], digest3[:])

	// 5. Perform SHA-256 hash on the exteneded RIPEMD-160 result
	hash5 := sha256.New()
	hash5.Write(vd4)
	digest5 := hash5.Sum(nil)

	// 6. Perform SHA-256 hash on the result of the previous SHA-256 hash
	hash6 := sha256.New()
	hash6.Write(digest5)
	digest6 := hash6.Sum(nil)

	// 7. Take the first 5 bytes of the second SHA-256 hash for checksum
	chsum := digest6[:4]

	// 8. Add the 4 checksum bytes from 7 at the end of extended RIPEMD-160 hash from 4 (25bytes)
	dc8 := make([]byte, 25)
	copy(dc8[:21], vd4[:])
	copy(dc8[21:], chsum[:])

	// 9. Convert the result from a byte string into base5
	address := base58.Encode(dc8)
	w.blockchainAddress = address

	return w
}

func (w *wallet) PrivateKey() *ecdsa.PrivateKey {
	return w.privateKey
}

func (w *wallet) PrivateKeyStr() string {
	return fmt.Sprintf("%x", w.privateKey.D.Bytes())
}

func (w *wallet) PublicKey() *ecdsa.PublicKey {
	return w.publicKey
}

func (w *wallet) PublicKeyStr() string {
	return fmt.Sprintf("%x%x", w.publicKey.X.Bytes(), w.publicKey.Y.Bytes())
}

func (w *wallet) BlockchainAddress() string {
	return w.blockchainAddress
}
