package utils

import (
	"os"
	"fmt"
	"bufio"
	"strings"
	"io"
)

var Config = make(map[string]string, 64)

func LoadStringfile(dst string) {
	file, err := os.Open(dst)
	if nil != err {
		fmt.Println(err)
	}
	defer file.Close()
	reader := bufio.NewReader(file)
	for {
		sb, ss, err1 := reader.ReadLine()
		if err1 == io.EOF {
			fmt.Println(ss)
			return
		}
		content := string(sb[:])
		keyvalues := strings.Split(content, "=")
		Config[keyvalues[0]] = keyvalues[1]
	}
}
