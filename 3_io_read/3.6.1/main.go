package main

import (
	"bufio"
	"strings"
	"fmt"
	"io"
)

var source = `一行目
2行目
3行目`

func main() {
	reader := bufio.NewReader(strings.NewReader(source))
	for {
		line, err := reader.ReadString('\n')
		fmt.Printf("%#v\n", line)
		if err == io.EOF {
			break
		}
	}
}