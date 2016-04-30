package main

import (
	"bufio"
	"fmt"
	// iconv "github.com/djimenez/iconv-go"
	"io"
	// "io/ioutil"
	// "github.com/astaxie/beego/orm"
	// "github.com/mattn/go-sqlite3"
	"os"
	"strconv"
	"strings"
)

const (
	DataPath  = "/mnt/golang/projects/src/shumo/data"
	DataName1 = "test.txt"
	DataName2 = "test2.txt"
)

type Choice struct {
	firstChoice  int64
	secondChoice int64
}

type Grade struct {
	knowledge    string
	intelligence string
	strain       string
	expression   string
}

type Personinfo struct {
	name   string
	score  int64
	choice Choice
	grade  Grade
	salary int64
}

type situation struct {
	welfare      string
	condition    string
	intensity    string
	advancement  string
	furtherStudy string
}

type require Grade

type hope struct {
	Situation situation
	Require   require
	salary    int64
}

type position map[int64]hope

type apartment map[string]position

var Apartment apartment = make(apartment, 0)

var Test [9]string
var Test1 [12]string

type personInfo []Personinfo

var PersonInfo personInfo = make(personInfo, 1)

func main() {

	read(DataPath+"/"+DataName1, 1)
	read(DataPath+"/"+DataName2, 2)

	print()

}

func print() {
	PersonInfo = PersonInfo[1:]

	// fmt.Println(PersonInfo)
	// fmt.Println(Apartment)
	fmt.Println("...........", Apartment["院部3"])
	// var data1 string

	// out := make([]byte, len(fmt.Sprintln(Apartment)))
	// iconv.Convert([]byte(fmt.Sprintln(Apartment)), out, "gb2312", "utf-8")
	// fmt.Println(out)
	// ioutil.WriteFile("out.txt", out, 0644)
}

func read(filename string, flag int) error {
	file, err := os.OpenFile(filename, os.O_RDONLY, os.ModePerm)
	if err != nil {
		fmt.Println(err)
	}
	inputReader := bufio.NewReader(file)
	// lineCounter := 0

	for {
		if flag == 1 {
			// var posi position
			var info Personinfo
			inputString, err := inputReader.ReadString('\n')
			inputString = strings.TrimSpace(inputString)
			lineLenght := len(inputString)
			if err != nil {
				if err != io.EOF {
					return err
				}
				if lineLenght == 0 || lineLenght == 1 {
					break
				}
			}
			// fmt.Println(inputString)

			var length int = 0
			var j int = 0
			for i, v := range inputString {
				// fmt.Println("...")
				if v == '\t' {

					Test[length] = inputString[j:i]
					j = i + 1
					length++
				}
				Test[length] = inputString[j:]
			}
			// fmt.Println(Test)
			{
				info.name = Test[0]
				info.score, _ = strconv.ParseInt(Test[1], 10, 64)
				info.choice.firstChoice, _ = strconv.ParseInt(Test[2], 10, 64)
				info.choice.secondChoice, _ = strconv.ParseInt(Test[3], 10, 64)
				info.grade.knowledge = Test[4]
				info.grade.intelligence = Test[5]
				info.grade.strain = Test[6]
				info.grade.expression = Test[7]
				info.salary, _ = strconv.ParseInt(Test[8], 10, 64)
			}

			PersonInfo = append(PersonInfo, info)

		}
		if flag == 2 {
			Hope := new(hope)
			inputString, err := inputReader.ReadString('\n')
			inputString = strings.TrimSpace(inputString)
			lineLenght := len(inputString)
			if err != nil {
				if err != io.EOF {
					return err
				}
				if lineLenght == 0 || lineLenght == 1 {
					break
				}
			}

			var length int = 0
			var j int = 0
			for i, v := range inputString {
				// fmt.Println("...")
				if v == '\t' {
					Test1[length] = inputString[j:i]
					j = i + 1
					length++
				}
				Test1[length] = inputString[j:]
			}
			// fmt.Println(Test1)
			{
				Hope.Situation.welfare = Test1[2]
				Hope.Situation.condition = Test1[3]
				Hope.Situation.intensity = Test1[4]
				Hope.Situation.advancement = Test1[5]
				Hope.Situation.furtherStudy = Test1[6]
				Hope.Require.knowledge = Test1[7]
				Hope.Require.intelligence = Test1[8]
				Hope.Require.strain = Test1[9]
				Hope.Require.expression = Test1[10]
				Hope.salary, _ = strconv.ParseInt(Test1[11], 10, 64)
				num, _ := strconv.ParseInt(Test1[1], 10, 64)
				var posi position = make(position, 0)
				posi[num] = *Hope

				Apartment[Test1[0]] = posi
				// fmt.Println(Test1[1], Test1[0], Apartment[Test1[0]], "-----------------", posi)

			}
		}
	}

	defer file.Close()
	return nil
}
