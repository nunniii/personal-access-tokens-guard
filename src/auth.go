package main

func authRequest(receivedKey string, hash string) bool {
	return createHash(receivedKey) == hash;
}
