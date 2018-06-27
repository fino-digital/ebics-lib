package test

import (
	"log"
	"regexp"
	"testing"

	"github.com/fino-digital/ebics-lib/ebics"
)

func TestGenerateRestrictedIDSEPA1(t *testing.T) {
	var re = regexp.MustCompile(`([A-Za-z0-9]|[+|-|?|:|(|)|.|'|,]){1,35}`)
	id := ebics.RestrictedIdentificationSEPA1()

	println(id)

	if !re.Match([]byte(id)) {
		t.Fatal("Don't match")
	}
}

func TestGenerateConsecutiveNumber(t *testing.T) {
	log.Println(ebics.CreateConsecutiveNumber())
	log.Println(ebics.CreateConsecutiveNumber())
}
