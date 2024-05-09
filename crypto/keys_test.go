package crypto

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGeneratePrivateKey(t *testing.T) {
	privKey := GeneratePrivateKey()
	assert.Equal(t, len(privKey.Bytes()), privKeyLen)

	pubKey := privKey.Public()
	assert.Equal(t, len(pubKey.Bytes()), pubKeyLen)
}

func TestNewPrivateKeyFromString(t *testing.T) {
	var (
		seed = "87489f2deef65d64fca7bd5a48fa8092d4b19b9be723df02fb21d2ed84a27580"
		privKey = NewPrivateKeyFromString(seed)
		addressStr = "932b579e3cbdb1c713675ebcee915f18f0ac4341"
	)
	assert.Equal(t, privKeyLen, len(privKey.Bytes()))
	addr := privKey.Public().Address()
	assert.Equal(t, addressStr, addr.String())
}

func TestPrivateKeySign(t *testing.T) {
	privKey := GeneratePrivateKey()
	msg := []byte("test some")
	pubKey := privKey.Public()

	sign := privKey.Sign(msg)
	assert.True(t, sign.Verify(pubKey, msg))
	// Test with invalid msg
	assert.False(t, sign.Verify(pubKey, []byte("wrong_key")))

	// Test with invalid key
	wrongPrivKey := GeneratePrivateKey()
	wrongPubKey := wrongPrivKey.Public()
	assert.False(t, sign.Verify(wrongPubKey, msg))
}

func TestPublicKeyToAddress(t *testing.T) {
	privKey := GeneratePrivateKey()
	pubKey := privKey.Public()
	addr := pubKey.Address()
	assert.Equal(t, addressLen, len(addr.Bytes()))
}