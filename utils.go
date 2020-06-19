package main

import (
	"fmt"
	"strings"
)

type blueprintWriter struct {
	strings.Builder
}

func (b *blueprintWriter) F(s string, v ...interface{}) {
	b.WriteString(fmt.Sprintf(s, v...))
	b.WriteString("\n")
}

func (b *blueprintWriter) P(v ...interface{}) {
	b.WriteString(fmt.Sprint(v...))
	b.WriteString("\n")
}


func (b *blueprintWriter) S(v ...interface{}) {
	b.WriteString(fmt.Sprint(v...))
}
