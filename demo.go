package main

import (
  "fmt"
  "github.com/btcsuite/btcutil/base58"
  "github.com/attiliodrei/gencrypt"
)

// NOTE: Error checking not handled in this example but should be in
// production.

var (
  // Data you want to encrypt
  data = []byte("test data")
  // Secret key. A 32-byte key is used to indicate AES-256. 16 and 24-byte keys
  // are accepted for AES-128 and AES-192 respectively, but are not
  // recommended.
  key = []byte("12345678901234561234567890123456")
)

func main() {
  // Get the GCM
  gcm, _ := gencrypt.NewGCM(key)

  // Encrypt data
  
  enc, _ := gcm.AESEncrypt(data)
  b58enc := base58.Encode([]byte(enc))
  fmt.Println(string(b58enc))
  // Decrypt data
  dec, _ := gcm.AESDecrypt(base58.Decode(b58enc))
  fmt.Println(string(dec))
}
