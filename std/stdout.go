package std

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"runtime"
)

func WriteResult(result string, copy bool) {
	if _, err := os.Stdout.WriteString(fmt.Sprintf("%s\n", result)); err != nil {
		log.Fatalf("os write to stdout: %s", err.Error())
	}

	if copy {
		if err := copyToClipboard([]byte(result)); err != nil {
			log.Fatalf("pbcopy: %s", err)
		}
	}

	return
}

func WriteError(msg string) {
	WriteResult(fmt.Sprintf("error: %s", msg), false)
}

func copyToClipboard(b []byte) error {
	var cmd *exec.Cmd

	switch runtime.GOOS {
	case "darwin":
		cmd = exec.Command("pbcopy")
	case "linux":
		cmd = exec.Command("xclip", "-selection", "c")
	default:
		return nil
	}

	pipe, err := cmd.StdinPipe()
	if err != nil {
		return fmt.Errorf("cmd stdin pipe: %w", err)
	}

	if err = cmd.Start(); err != nil {
		return fmt.Errorf("cmd start: %w", err)
	}

	if _, err = pipe.Write(b); err != nil {
		return fmt.Errorf("pipe write: %w", err)
	}

	if err = pipe.Close(); err != nil {
		return fmt.Errorf("pipe close: %w", err)
	}

	if err = cmd.Wait(); err != nil {
		return fmt.Errorf("cmd wait: %w", err)
	}

	return nil
}
