package validation

import (
	"crypto/sha256"
	"fmt"
	"os"
)

var (
	htmlHash [32]byte
	infoHash [32]byte
)

func LoadHashes() {
	htmlHash = HtmlHash()
	infoHash = InfoHash()
}

// SHA-256 cryptographic hash algorithm

func HtmlHash() [32]byte {
	data, err := os.ReadFile("templates/homepage.html")
	if err != nil {
		fmt.Println(err)
		return [32]byte{}
	}
	
	return sha256.Sum256(data)
}

func InfoHash() [32]byte {
	data, err := os.ReadFile("templates/info.html")
	if err != nil {
		fmt.Println(err)
		return [32]byte{}
	}

	return sha256.Sum256(data)
}
