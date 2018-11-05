package adascript

import (
	"testing"
	"log"
)

func TestGenUnRepeatMnemonic(t *testing.T) {
	log.Println(GenMnemonic(12))
}

func TestGenSpendingPassword(t *testing.T) {
	log.Println(GenSpendingPassword("hello world"))
}
