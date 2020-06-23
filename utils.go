package omiux

import (
	"fmt"
	"strings"
)

type simpleWriter struct {
	strings.Builder
}

func (b *simpleWriter) F(s string, v ...interface{}) {
	b.WriteString(fmt.Sprintf(s, v...))
	b.WriteString("\n")
}

func (b *simpleWriter) P(v ...interface{}) {
	b.WriteString(fmt.Sprint(v...))
	b.WriteString("\n")
}


func (b *simpleWriter) S(v ...interface{}) {
	b.WriteString(fmt.Sprint(v...))
}
