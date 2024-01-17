package helper

import (
	"bufio"
	"os"
	"strconv"
)

func ReadString() string {
	scanner := bufio.NewScanner(os.Stdin)
	if scanner.Scan() {
		return scanner.Text()
	}

	return ""
}

func ReadInt() (int, error) {
	scanner := bufio.NewScanner(os.Stdin)
	if scanner.Scan() {
		input := scanner.Text()

		if num, err := strconv.Atoi(input); err == nil {

			return num, nil
		} else {

			return 0, err
		}
	}
	return 0, nil
}
