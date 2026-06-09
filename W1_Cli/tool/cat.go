package tool

import (
	"bufio"
	"errors"
	"fmt"
	"os"
)

func Cat(args []string) error{
	if len(args) <= 1 {
		return (errors.New("no enough args"))
	}

	file, err := os.Open(args[1])
	if err != nil {
		return (fmt.Errorf("path is wrong, %w", err))
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		tmpRes := scanner.Text()
		fmt.Println(tmpRes)
	}
	if err := scanner.Err(); err != nil {
		return (fmt.Errorf("read file failed, %w", err))
	}

	return nil
}
