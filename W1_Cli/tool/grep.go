package tool

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strings"
)

func Grep(args []string) error {
	if len(args) <= 2 {
		return errors.New("no enough args, usage: grep <pattern> <file>")
	}
	file, err := os.Open(args[2])
	if err != nil {
		return err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		tmpRes := scanner.Text()
		if strings.Contains(tmpRes, args[1]) {
			fmt.Println(tmpRes)
		}
	}
	if err := scanner.Err(); err != nil {
		return err
	}

	return nil
}
