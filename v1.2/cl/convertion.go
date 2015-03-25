package cl

import (
	"log"
	"reflect"
	"strings"
	"unsafe"
)

func Str(str string) *uint8 {
	if !strings.HasSuffix(str, "\x00") {
		log.Fatal("str argument missing null terminator", str)
	}
	header := (*reflect.StringHeader)(unsafe.Pointer(&str))
	return (*uint8)(unsafe.Pointer(header.Data))
}
