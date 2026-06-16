package tool

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strings"
)

func Grep(args []string) error {
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
		tmpRes := scanner.Text()
		if strings.Contains(tmpRes, args[0]) {
			fmt.Println(tmpRes)
		}
	}
	if err := scanner.Err(); err != nil {
		return err
	}

	return nil
}
