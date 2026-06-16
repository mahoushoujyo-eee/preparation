package tool

import (
	"bufio"
	"errors"
	"fmt"
	"os"
)

func Cat(args []string) error {
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
		tmpRes := scanner.Text()
		fmt.Println(tmpRes)
	}
	if err := scanner.Err(); err != nil {
		return err
	}

	return nil
}
