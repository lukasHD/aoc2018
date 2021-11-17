package commons

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

func LoadLines(fname string) []string {
	fileBytes, err := ioutil.ReadFile(fname)

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	sliceData := strings.Split(string(fileBytes), "\n")

	pretty := "[" + strings.Join(sliceData, ", ") + "]"
	fmt.Println(pretty)

	return sliceData
}
