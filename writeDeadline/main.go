package main

import (
	"bytes"
	"github.com/sirupsen/logrus"
	"github.com/spf13/pflag"
	"os"
	"syscall"
	"time"
)

func main() {
	fset := pflag.NewFlagSet("read", pflag.ExitOnError)
	filename := fset.String("filename", "delete.me", "filename to open")
	sz := fset.Int("size", 10 * 1024 * 1024,"number of bytes to write")
	writeDeadline := fset.Duration("write-deadline", time.Millisecond, "write deadline")
	if err := fset.Parse(os.Args[1:]); err != nil {
		logrus.WithError(err).Fatal()
	}

	f, err := os.OpenFile(*filename, syscall.O_CREAT | syscall.O_SYNC | syscall.O_WRONLY, 0600)
	if err != nil {
		logrus.WithError(err).Fatal()
	}
	f.SetWriteDeadline(time.Now().Add(*writeDeadline))
	if _, err := f.Write(bytes.Repeat([]byte("x"), *sz)); err != nil {
		logrus.WithError(err).Panicln()
	}
	if err := f.Close(); err != nil {
		logrus.WithError(err).Panicln()
	}
}
