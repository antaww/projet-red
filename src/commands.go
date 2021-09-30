package main

import (
	"os"
	"strings"
)

func ClearLog() {
	os.Stdout.WriteString("\x1b[3;J\x1b[H\x1b[2J")
}

func Input() string {
	m, _ := BR.ReadString('\n')
	m = strings.Replace(m, "\r\n", "", -1)
	return m
}

func verif(a rune) bool {
	if (a >= 'A' && a <= 'Z') || (a >= 'a' && a <= 'z') || (a >= '0' && a <= '9') {
		return true
	}
	return false
}

func Capitalize(s string) string {
	sRune := []rune(s)
	lettre := true
	for i := 0; i < len(s); i++ {
		if verif(sRune[i]) == true && lettre {
			if sRune[i] >= 'a' && sRune[i] <= 'z' {
				sRune[i] += 'A' - 'a'
			}
			lettre = false
		} else if sRune[i] >= 'A' && sRune[i] <= 'Z' {
			sRune[i] += 'a' - 'A'
		} else if verif(sRune[i]) == false {
			lettre = true
		}
	}
	return string(sRune)
}
