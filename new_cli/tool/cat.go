package tool

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"os"
)

// Cat把args[0]指定的文件逐行写入w
func Cat(w io.Writer, args []string) error {
	if len(args) == 0 {
		return errors.New("not enough args, usage: cat <file>")
	}

	file, err := os.Open(args[0])
	if err != nil {
		return err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		fmt.Fprintln(w, scanner.Text())
	}
	return scanner.Err()
}
