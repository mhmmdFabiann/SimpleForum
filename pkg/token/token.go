package token

import (
	"crypto/rand"
	"encoding/hex"
)

func GenerateRefreshToken() (string, error) {
	// 1. Menggunakan 32 byte untuk keamanan yang lebih kuat
	b := make([]byte, 32)

	// 2. Membaca dari crypto/rand
	_, err := rand.Read(b)
	if err != nil {
		// 3. Jika ada error, kembalikan string kosong DAN error-nya
		//    Ini memberitahu pemanggil bahwa sesuatu yang salah telah terjadi.
		return "", err
	}

	// 4. Kembalikan token dan error 'nil' (yang berarti tidak ada error)
	return hex.EncodeToString(b), nil
}