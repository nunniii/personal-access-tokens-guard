package main

import(
	"crypto/md5"
	"encoding/hex"
)


// função para gerar hash
func createHash(key string) string {
	hasher := md5.New()
	hasher.Write([]byte(key))
	return hex.EncodeToString(hasher.Sum(nil))
}

