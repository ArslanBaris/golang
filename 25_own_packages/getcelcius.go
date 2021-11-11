package getcelcius

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func GetCelcius() (float64, error) {
	reader := bufio.NewReader(os.Stdin)
	input, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println(err)
	}
	input = strings.TrimSpace(input)
	num, err := strconv.ParseFloat(input, 64)
	if err != nil {
		fmt.Println(err)
	}
	return num, nil
}