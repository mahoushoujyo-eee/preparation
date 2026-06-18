package tool

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"os"
	"strings"
)

// Grep在args[1]指定的文件中查找包含args[0]子串的行，写入w
func Grep(w io.Writer, args []string) error {
	if len(args) <= 1 {
		return errors.New("not enough args, usage: grep <pattern> <file>")
	}
	file, err := os.Open(args[1])
	if err != nil {
		return err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		if strings.Contains(line, args[0]) {
			fmt.Fprintln(w, line)
		}
	}
	return scanner.Err()
}
