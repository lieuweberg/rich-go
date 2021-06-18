// +build windows

package ipc

import (
	"errors"
	"fmt"
	"log"
	"time"

	npipe "gopkg.in/natefinch/npipe.v2"
)

// OpenSocket opens the discord-ipc-0 named pipe
func OpenSocket() error {
	// Connect to the Windows named pipe, this is a well known name
	// We use DialTimeout since it will block forever (or very very long) on Windows
	// if the pipe is not available (Discord not running)
	for i := range make([]rune, 3) {
		sock, err := npipe.DialTimeout(fmt.Sprintf(`\\.\pipe\discord-ipc-%d`, i), time.Second*2)
		if err != nil {
			if errors.As(err, &npipe.PipeError{}) {
				continue
			}
			return err
		} else {
			socket = sock
			log.Printf("Connected to pipe %d", i)
			return nil
		}
	}

	return errors.New("unable to connect to pipes 0, 1 or 2")
}
