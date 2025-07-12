package main

import (
  "crypto/aes"
  "crypto/cipher"
  "crypto/rand"
  "io"
)

func encrypt(key, plaintext []byte) ([]byte, error) {
  block, err := aes.NewCipher(key)
  if err != nil {
    return nil, err
  }
  aesGCM, err := cipher.NewGCM(block)
  if err != nil {
    return nil, err
  }
  nonce := make([]byte, aesGCM.NonceSize())
  io.ReadFull(rand.Reader, nonce)
  return aesGCM.Seal(nonce, nonce, plaintext, nil), nil
}

func decrypt(key, ciphertext []byte) ([]byte, error) {
  block, err := aes.NewCipher(key)
  if err != nil {
    return nil, err
  }
  aesGCM, err := cipher.NewGCM(block)
  if err != nil {
    return nil, err
  }
  nonceSize := aesGCM.NonceSize()
  nonce, ciphertext := ciphertext[:nonceSize], ciphertext[nonceSize:]
  return aesGCM.Open(nil, nonce, ciphertext, nil)
}
