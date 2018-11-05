package adascript

import (
	"github.com/tyler-smith/go-bip39"
	"strings"
	"errors"
	"crypto/sha256"
	"fmt"
)

// validate word size is legal
func ValidateWordSize(wordSize int) bool {
	if wordSize < 3 || wordSize > 24 || wordSize%3 != 0 {
		return false
	}
	return true
}

// generate seed words (mnemonic)
func GenMnemonic(wordSize int) (string, error) {

	if !ValidateWordSize(wordSize) {
		return "", errors.New(
			"Word size must be [3, 24] and a multiple of 3")
	}
	// calculate bitSize
	bitSize := wordSize * 11 * 32 / 33
	// NewEntropy will create random entropy bytes
	entropy, err := bip39.NewEntropy(bitSize)
	if err != nil {
		return "", errors.New(
			"Failed to generate entropy:" + err.Error())
	}

	// generate (english) seed words based on the entropy
	mnemonic, err := bip39.NewMnemonic(entropy)
	if err != nil {
		return "", errors.New(
			"Failed to generate mnemonic:" + err.Error())
	}
	return mnemonic, nil
}

func GenUnRepeatMnemonic(wordSize int) (string, error) {
	mnemonic, err := GenMnemonic(wordSize)
	if err != nil {
		return "", err
	}

	for IsHasRepeatMnemonic(mnemonic) {
		mnemonic, err = GenMnemonic(wordSize)
		if err != nil {
			return "", err
		}
	}

	return mnemonic, nil
}

func IsHasRepeatMnemonic(mnemonic string) bool {
	words := strings.Split(mnemonic, " ")
	return IsHasRepeatElement(words)
}

func IsHasRepeatElement(eles []string) bool {

	for i := 0; i < len(eles); i++ {
		for j := i + 1; j < len(eles); j++ {
			if eles[i] == eles[j] {
				return true
			}
		}
	}
	return false
}

func GenSpendingPassword(s string) string {

	h := sha256.New()
	h.Write([]byte(s))
	return fmt.Sprintf("%x", h.Sum(nil))
}
