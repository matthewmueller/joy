package subtlecrypto

import "github.com/matthewmueller/golly/js"

// js:"SubtleCrypto,omit"
type SubtleCrypto struct {
}

// Decrypt
func (*SubtleCrypto) Decrypt(algorithm interface{}, key *cryptokey.CryptoKey, data []byte) (i interface{}) {
	js.Rewrite("await $<.Decrypt($1, $2, $3)", algorithm, key, data)
	return i
}

// DeriveBits
func (*SubtleCrypto) DeriveBits(algorithm interface{}, baseKey *cryptokey.CryptoKey, length uint) (i interface{}) {
	js.Rewrite("await $<.DeriveBits($1, $2, $3)", algorithm, baseKey, length)
	return i
}

// DeriveKey
func (*SubtleCrypto) DeriveKey(algorithm interface{}, baseKey *cryptokey.CryptoKey, derivedKeyType interface{}, extractable bool, keyUsages []string) (i interface{}) {
	js.Rewrite("await $<.DeriveKey($1, $2, $3, $4, $5)", algorithm, baseKey, derivedKeyType, extractable, keyUsages)
	return i
}

// Digest
func (*SubtleCrypto) Digest(algorithm interface{}, data []byte) (i interface{}) {
	js.Rewrite("await $<.Digest($1, $2)", algorithm, data)
	return i
}

// Encrypt
func (*SubtleCrypto) Encrypt(algorithm interface{}, key *cryptokey.CryptoKey, data []byte) (i interface{}) {
	js.Rewrite("await $<.Encrypt($1, $2, $3)", algorithm, key, data)
	return i
}

// ExportKey
func (*SubtleCrypto) ExportKey(format string, key *cryptokey.CryptoKey) (i interface{}) {
	js.Rewrite("await $<.ExportKey($1, $2)", format, key)
	return i
}

// GenerateKey
func (*SubtleCrypto) GenerateKey(algorithm interface{}, extractable bool, keyUsages []string) (i interface{}) {
	js.Rewrite("await $<.GenerateKey($1, $2, $3)", algorithm, extractable, keyUsages)
	return i
}

// ImportKey
func (*SubtleCrypto) ImportKey(format string, keyData []byte, algorithm interface{}, extractable bool, keyUsages []string) (i interface{}) {
	js.Rewrite("await $<.ImportKey($1, $2, $3, $4, $5)", format, keyData, algorithm, extractable, keyUsages)
	return i
}

// Sign
func (*SubtleCrypto) Sign(algorithm interface{}, key *cryptokey.CryptoKey, data []byte) (i interface{}) {
	js.Rewrite("await $<.Sign($1, $2, $3)", algorithm, key, data)
	return i
}

// UnwrapKey
func (*SubtleCrypto) UnwrapKey(format string, wrappedKey []byte, unwrappingKey *cryptokey.CryptoKey, unwrapAlgorithm interface{}, unwrappedKeyAlgorithm interface{}, extractable bool, keyUsages []string) (i interface{}) {
	js.Rewrite("await $<.UnwrapKey($1, $2, $3, $4, $5, $6, $7)", format, wrappedKey, unwrappingKey, unwrapAlgorithm, unwrappedKeyAlgorithm, extractable, keyUsages)
	return i
}

// Verify
func (*SubtleCrypto) Verify(algorithm interface{}, key *cryptokey.CryptoKey, signature []byte, data []byte) (i interface{}) {
	js.Rewrite("await $<.Verify($1, $2, $3, $4)", algorithm, key, signature, data)
	return i
}

// WrapKey
func (*SubtleCrypto) WrapKey(format string, key *cryptokey.CryptoKey, wrappingKey *cryptokey.CryptoKey, wrapAlgorithm interface{}) (i interface{}) {
	js.Rewrite("await $<.WrapKey($1, $2, $3, $4)", format, key, wrappingKey, wrapAlgorithm)
	return i
}
