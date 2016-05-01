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

type personResult struct {
	name   string
	Result float64
}

var dataResultfloat []personResult

// var jobCategory = make([]string, 6)
var jobCategory = []string{"", "行政管理", "教师管理", "学生管理", "教务管理", "组织管理"}

type JobRequire map[int64]require

var jobRequire JobRequire = make(JobRequire, 0)

func main() {
	read(DataPath+"/"+DataName1, 1)
	read(DataPath+"/"+DataName2, 2)
	// 根据公式算出权值
	handleWeight()
	// fmt.Println(dataResult)
	handleType()

	sortPersonScore()
	print()

}

func handleType() {
	var i int64
	var length int64 = int64(len(jobCategory))
	for i = 1; i < length; i++ {
		for j := 0; j < len(PersonInfo); j++ {

			if jobRequire[i].knowledge >= PersonInfo[j].grade.knowledge && jobRequire[i].intelligence >= PersonInfo[j].grade.intelligence && jobRequire[i].strain >= PersonInfo[j].grade.strain && jobRequire[i].expression >= PersonInfo[j].grade.expression {
				fmt.Println(PersonInfo[j], "********", "Type:", i)
			}
		}
	}
}

func handleWeight() {
	length := len(PersonInfo)
	dataResult := make([]string, length)
	dataResultfloat = make([]personResult, length)

	for i := 0; i < length; i++ {
		dataResult[i], dataResultfloat[i].name, dataResultfloat[i].Result = function(i)
		// fmt.Println(dataResult[i], dataResultfloat[i])
	}
	sort(dataResultfloat)
	// fmt.Println(dataResultfloat)
	// fmt.Println(dataResultfloat, "********")
	writedata("data2.txt", dataResult)

}

func sortPersonScore() {
	length := len(PersonInfo)
	for i := 0; i < length-1; i++ {
		for j := i + 1; j < length; j++ {
			if PersonInfo[i].score < PersonInfo[j].score {
				PersonInfo[i].score, PersonInfo[j].score = PersonInfo[j].score, PersonInfo[i].score
			}
		}
	}
}

func sort(data []personResult) {
	length := len(data)
	for i := 0; i < length-1; i++ {
		for j := i + 1; j < length; j++ {
			if data[i].Result < data[j].Result {
				data[i].Result, data[j].Result = data[j].Result, data[i].Result
			}
		}
	}

}

func function(i int) (string, string, float64) {
	var data1 float64
	data1 = function1(PersonInfo[i].grade.knowledge)
	var data2 float64
	data2 = function1(PersonInfo[i].grade.intelligence)
	var data3 float64
	data3 = function1(PersonInfo[i].grade.strain)
	var data4 float64
	data4 = function1(PersonInfo[i].grade.expression)
	// fmt.Println(data1)
	result := (float64(PersonInfo[i].score)/300.0*0.7 + (data1+data2+data3+data4)*0.25*0.3)

	return fmt.Sprintf("%s %f\n", PersonInfo[i].name, result), PersonInfo[i].name, result
}

func function1(grade1 string) float64 {
	var data float64
	// fmt.Println(grade1)
	switch grade1 {
	case "A":
		data = 1
	case "B":
		data = 0.75
	case "C":
		data = 0.5
	case "D":
		data = 0.25
	}
	// fmt.Println(data)
	return data
}

func writedata(filename string, dataResultString []string) {
	file, _ := os.OpenFile(filename, os.O_CREATE|os.O_APPEND|os.O_WRONLY, os.ModePerm)
	for _, v := range dataResultString {
		file.WriteString(v)
	}
	defer file.Close()
}

func print() {
	// fmt.Println(jobRequire)
	// fmt.Println(dataResultfloat[0:10])

	fmt.Println(PersonInfo[0:10])
	// fmt.Println(Apartment)
	// fmt.Println(PersonInfo[0])
	// fmt.Println("...........", Apartment["院部3"])
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
				{
					jobRequire[num] = Hope.Require
				}
				// fmt.Println(jobRequire)
				Apartment[Test1[0]] = posi
				// fmt.Println(Test1[1], Test1[0], Apartment[Test1[0]], "-----------------", posi)

			}
		}
	}
	PersonInfo = PersonInfo[1:]

	defer file.Close()
	return nil
}
