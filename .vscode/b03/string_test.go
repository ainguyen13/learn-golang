package main

import (
	"testing"
)

func BenchmarkModifyString1(b *testing.B) {
	s := "hello"
	for i := 0; i < b.N; i++ {
		modifyString(s)
	}
}

func BenchmarkModifyString2(b *testing.B) {
	s := "hello"
	for i := 0; i < b.N; i++ {
		modifyString2(&s)
	}
}

func BenchmarkModifyString3(b *testing.B) {
	s := "hello"
	for i := 0; i < b.N; i++ {
		_ = modifyString3(s)
	}
}