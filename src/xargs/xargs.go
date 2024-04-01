package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"os/exec"
	"strings"
)

func readToSlice(r io.Reader) ([]string, error) {
	scanner := bufio.NewScanner(r)
	scanner.Split(bufio.ScanWords)

	result := make([]string, 0)

	for scanner.Scan() {
		result = append(result, scanner.Text())
	}

	return result, scanner.Err()
}

func main() {
	log.SetFlags(log.Lshortfile)

	var builder strings.Builder
	whitespace := false

	for _, arg := range os.Args[1:] {
		if whitespace {
			builder.WriteByte(' ')
		}
		whitespace = true

		builder.WriteString(arg)
	}

	args, err := readToSlice(os.Stdin)
	if err != nil {
		log.Fatalln(err)
	}

	if whitespace {
		// command is present, execute
		cmd := exec.Command(builder.String(), args...)

		cmd.Stderr = os.Stderr
		cmd.Stdin = os.Stdin
		cmd.Stdout = os.Stdout

		cmd.Run()
		os.Exit(cmd.ProcessState.ExitCode())

	} else {
		// no command, just print args
		fmt.Println(strings.Join(args, " "))
	}
}
