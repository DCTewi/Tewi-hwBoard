package util

import (
	"archive/zip"
	"io"
	"math/rand"
	"os"
	"path/filepath"
	"regexp"
	"strings"
	"time"

	"github.com/dctewi/tewi-hwboard/config"
)

const (
	ckeylen      = 6
	tokenlen     = 32
	sessionidlen = 44

	letterBytes     = "1234567890abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ_-"
	letterIndexBits = 6
	letterIndexMask = 1<<letterIndexBits - 1
	letterIndexMax  = 63 / letterIndexBits
)

// Random source of Generator
var src = rand.NewSource(time.Now().UnixNano())

// GenRandomString of length l
func GenRandomString(l int) string {
	bytes := make([]byte, l)
	// Every int63 for letterIndexMax random choices
	for i, cache, remain := l-1, src.Int63(), letterIndexMax; i >= 0; {
		if remain == 0 {
			// Refresh cache pool
			cache, remain = src.Int63(), letterIndexMax
		}
		if idx := int(cache & letterIndexMask); idx < len(letterBytes) {
			// Make a random choices of int(letterIndexBit)
			bytes[i] = letterBytes[idx]
			i--
		}
		cache >>= letterIndexBits
		remain--
	}
	return string(bytes)
}

// GenCKey gen new confirmkey
func GenCKey() string {
	return GenRandomString(ckeylen)
}

// GenToken gen new token
func GenToken() string {
	return GenRandomString(tokenlen)
}

// GenSessionID gen new sid
func GenSessionID() string {
	return GenRandomString(sessionidlen)
}

// CheckUserQ legality
func CheckUserQ(qq string) bool {
	pattern := config.Regex.PatternQQ

	ok, _ := regexp.MatchString(pattern, qq)

	return ok
}

// CheckStuID legality
func CheckStuID(id string) bool {
	pattern := config.Regex.PatternStuID

	ok, _ := regexp.MatchString(pattern, id)

	return ok
}

// MakeDirIfNessesary rt
func MakeDirIfNessesary(path string) error {
	_, err := os.Stat(path)
	if os.IsNotExist(err) {
		err = os.MkdirAll(path, os.ModePerm)
	}
	return err
}

// ZipDir to path.zip
func ZipDir(dir string) error {
	MakeDirIfNessesary(dir)
	fzip, err := os.Create(dir + ".zip")
	if err != nil {
		return err
	}
	defer fzip.Close()

	w := zip.NewWriter(fzip)
	defer w.Close()

	err = filepath.Walk(dir, func(path string, info os.FileInfo, err error) error {
		if !info.IsDir() {
			pre := strings.LastIndex(dir, "/")
			if pre < 1 {
				pre = 0
			} else {
				pre--
			}
			fdest, err := w.Create(path[pre:])
			if err != nil {
				return err
			}

			fsrc, err := os.Open(path)
			if err != nil {
				return err
			}
			defer fsrc.Close()

			_, err = io.Copy(fdest, fsrc)
			if err != nil {
				return err
			}
		}
		return nil
	})
	return err
}
