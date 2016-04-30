package main

import (
	"bufio"
	"fmt"
	iconv "github.com/djimenez/iconv-go"
	"io"
	// "io/ioutil"
	"os"
)

var filename string = "data/test2.txt.bak"

func main() {
	file, _ := os.Open(filename)
	inputString := bufio.NewReader(file)
	os.Create("test5.txt")
	file1, err := os.OpenFile("test5.txt", os.O_WRONLY|os.O_CREATE|os.O_TRUNC|os.O_APPEND, os.ModePerm)
	fmt.Println(err)
	var input string = ""
	// var output []byte
	// output = make([]byte, 10000)
	//	output = output[:]
	var out []byte
	for {
		InputString, err := inputString.ReadString('\n')
		fmt.Println(InputString)
		lineLenght := len(InputString)
		if err != nil {
			if err != io.EOF {
				return
			}
			if lineLenght == 0 {
				break
			}
		}

		input = InputString + "       \n"
		fmt.Println(input)

		out = make([]byte, len(input))
		out = out[:]

		iconv.Convert([]byte(input), out, "gb2312", "utf-8")
		file1.WriteString(string(out))
		fmt.Println(out)
		/*
					input = input + InputString

					// fmt.Println(input)
					iconv.Convert([]byte(input), out, "gb2312", "utf-8")
					ioutil.WriteFile("test4.txt", out, 0644)

					out = make([]byte, len(input))
					out = out[:]
					// output = append(output, out)
				fmt.Println(input)
			// fmt.Println(out)
		*/
	}

	// output = append(output, out[:])

	defer file.Close()
	defer file1.Close()
}
