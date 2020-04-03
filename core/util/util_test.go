package util_test

import (
	"fmt"
	"testing"

	"github.com/dctewi/tewi-hwboard/core/util"
)

func TestCKeyGen(t *testing.T) {
	for i := 0; i < 30; i++ {
		fmt.Println(util.GenCKey())
	}
}

func TestTokens(t *testing.T) {
	for i := 0; i < 30; i++ {
		token := util.GenToken()
		fmt.Println(token, len(token))
	}
}

func TestSessionID(t *testing.T) {
	for i := 0; i < 30; i++ {
		sid := util.GenSessionID()
		fmt.Println(sid, len(sid))
	}
}

func TestZipDir(t *testing.T) {
	err := util.ZipDir("./../../static")
	if err != nil {
		t.Error(err)
	}
}
