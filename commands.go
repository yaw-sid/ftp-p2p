package main

import (
	"bufio"
	"io"
	"os"
)

const (
	DIR      = "DIR"
	CD       = "CD"
	PWD      = "PWD"
	RETRIEVE = "RETR"
)

func chdir(w *bufio.Writer, s string) {
	defer w.WriteString("\r\n")

	if os.Chdir(s) == nil {
		w.WriteString("OK")
	} else {
		w.WriteString("ERROR")
	}
}

func pwd(w *bufio.Writer) {
	defer w.WriteString("\r\n")

	s, err := os.Getwd()
	if err != nil {
		w.WriteString("")
		return
	}
	w.WriteString(s)
}

func dirList(w *bufio.Writer) {
	defer w.WriteString("\r\n")

	dir, err := os.Open(".")
	if err != nil {
		return
	}

	names, err := dir.Readdirnames(-1)
	if err != nil {
		return
	}
	for _, nm := range names {
		w.WriteString(nm + "\r\n")
	}
}

func retrieve(w *bufio.Writer, filename string) {
	defer w.WriteString("\r\n")

	file, err := os.OpenFile(filename, os.O_RDONLY, 0600)
	if err != nil {
		return
	}
	io.Copy(w, file)
}
