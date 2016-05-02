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

// 数据存放位置
const (
	DataPath  = "/mnt/golang/projects/src/shumo/data"
	DataName1 = "test.txt"
	DataName2 = "test2.txt"
)

// 志愿
type Choice struct {
	firstChoice  int64
	secondChoice int64
}

// 能力分类
type Grade struct {
	knowledge    string
	intelligence string
	strain       string
	expression   string
}

// 用户数据存放格式
type Personinfo struct {
	name   string
	score  int64
	choice Choice
	grade  Grade
	salary int64
}

// 部门基本要求
type situation struct {
	welfare      string
	condition    string
	intensity    string
	advancement  string
	furtherStudy string
}

type require Grade

// 期望人员能力
type hope struct {
	Situation situation
	Require   require
	salary    int64
}

// 各部门期望人员能力
type position map[int64]hope

// 部门全部数据
type apartment map[string]position

//
var Apartment apartment = make(apartment, 0)

//录入数据人员使用
var Test [9]string

//录入部门数据使用
var Test1 [12]string

//
type personInfo []Personinfo

//
var PersonInfo personInfo = make(personInfo, 1)

//
type personResult struct {
	name   string
	Result float64
}

//
var dataResultfloat []personResult

// var jobCategory = make([]string, 6)
var jobCategory = []string{"", "行政管理", "教师管理", "学生管理", "教务管理", "组织管理"}

//
type JobRequire map[int64]require

//
var jobRequire JobRequire = make(JobRequire, 0)

// question 2 handler
//person -> work
const (
	PersonNum = 10
)

const (
	x1 float64 = 0.1
	x2 float64 = 0.2
	x3 float64 = 0.7
)

//wprk -> person
const (
	y1 float64 = 0.1
	y2 float64 = 0.2
	y3 float64 = 0.3
	y4 float64 = 0.4
)

var a [10][10]float64
var b [10][10]float64
var c [10][10]float64

var a1 [10][10]float64
var b1 [10][10]float64
var c1 [10][10]float64
var d1 [10][10]float64

type workType struct {
	Name      string
	Type      int64
	Situation situation
	Require   require
	salary    int64
}

var WorkType []workType = make([]workType, 10)

var ResultA [10][10]float64
var ResultB [10][10]float64
var ResultC [10][10]float64

func main() {
	read(DataPath+"/"+DataName1, 1)
	read(DataPath+"/"+DataName2, 2)
	// 根据公式算出权值
	handleWeight()
	// handleType()

	sortPersonScore()

	// voluntyTable1()
	// voluntyTable2()

	// MaptoSlice()

	handleQuestion2()

	print()
}

func HandlerR() {
	length := len(a)
	for i := 0; i < length; i++ {
		for j := 0; j < length; j++ {
			ResultA[i][j] = a[i][j]*x1 + b[i][j]*x2 + c[i][j]*x3
			ResultB[i][j] = a1[i][j]*y1 + b1[i][j]*y2 + c1[i][j]*y3 + d1[i][j]*y4
			ResultC[i][j] = ResultA[i][j] + ResultB[i][j]
		}
	}
}

func MaptoSlice() {
	var i int64
	// var length int64 = int64(len(PersonInfo))
	// fmt.Println(Apartment)
	for x, v := range Apartment {
		for y, val := range v {
			WorkType[i].Name = x
			WorkType[i].Type = y
			WorkType[i].Require = val.Require
			WorkType[i].Situation = val.Situation
			WorkType[i].salary = val.salary
			i++
		}
	}
}

func handleQuestion2() {
	PersonInfo = PersonInfo[0:PersonNum]
	personToWork()

	workToPerson()

	HandlerR()
}

func personToWork() {
	// var length int64 = int64(len(PersonInfo))
	// fmt.Println(Apartment)
	{
		for i, v := range PersonInfo {
			for j, value := range WorkType {
				if value.salary > v.salary {
					a[i][j] = 1
				} else {
					a[i][j] = 0
				}
			}
		}
	}

	{
		for i, v := range PersonInfo {
			for j, value := range WorkType {
				if v.choice.firstChoice == value.Type {
					b[i][j] = 1
				}
				if v.choice.secondChoice == value.Type {
					b[i][j] = 0.5
				}
			}
		}
	}

	{
		for i, _ := range PersonInfo {
			for j, value := range WorkType {
				c[i][j] = (ChineseToFloat(value.Situation.advancement) + ChineseToFloat(value.Situation.condition) + ChineseToFloat(value.Situation.furtherStudy) + ChineseToFloat(value.Situation.intensity) + ChineseToFloat(value.Situation.welfare)) / 5
			}
		}
	}

}

func ChineseToFloat(data string) float64 {
	switch data {
	case "优":
		return 1
	case "中":
		return 0.75
	case "差":
		return 0.5
	case "多":
		return 1
	case "少":
		return 0.5
	default:
		return 0
	}
}

func workToPerson() {
	{
		for i, v := range PersonInfo {
			for j, value := range WorkType {
				if value.salary > v.salary {
					a1[i][j] = 1
				} else {
					a1[i][j] = 0
				}

			}
		}
	}

	{
		for i, v := range PersonInfo {
			for j, value := range WorkType {
				if v.choice.firstChoice == value.Type {
					b1[i][j] = 1
				}
				if v.choice.secondChoice == value.Type {
					b1[i][j] = 0.5
				}
			}
		}
	}

	{
		for i, v := range PersonInfo {
			for j, _ := range WorkType {
				c1[i][j] = float64(v.score) / 300.0
			}
		}
	}

	{
		for i, v := range PersonInfo {
			for j, _ := range WorkType {
				d1[i][j] = (function1(v.grade.knowledge) + function1(v.grade.intelligence) + function1(v.grade.expression) + function1(v.grade.strain)) * 0.25
			}
		}
	}

}

func voluntyTable1() {
	var i int64
	var length int64 = int64(len(jobCategory))
	for i = 1; i < length; i++ {
		// fmt.Println(jobCategory[i], "This is first")
		for j := 0; j < len(PersonInfo); j++ {
			if PersonInfo[j].choice.firstChoice == i {
				// fmt.Println(PersonInfo[j])
			}
		}
	}
}

func voluntyTable2() {
	var i int64
	var length int64 = int64(len(jobCategory))
	for i = 1; i < length; i++ {
		fmt.Println(jobCategory[i], "This is second")
		for j := 0; j < len(PersonInfo); j++ {
			if PersonInfo[j].choice.secondChoice == i {
				fmt.Println(PersonInfo[j])
			}
		}
	}
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
				data[i], data[j] = data[j], data[i]
			}
		}
	}
	// fmt.Println(data[0:10])

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

	// fmt.Println(PersonInfo[0:10])
	{
		for _, v := range dataResultfloat[0:10] {
			fmt.Printf("%s\t%0.3f", v.name, v.Result)
			fmt.Println()
		}
	}
	// fmt.Println(Apartment)
	// fmt.Println(PersonInfo[0])
	// fmt.Println("...........", Apartment["院部3"])
	// var data1 string

	// fmt.Println(dataResult)

	// out := make([]byte, len(fmt.Sprintln(Apartment)))
	// iconv.Convert([]byte(fmt.Sprintln(Apartment)), out, "gb2312", "utf-8")
	// fmt.Println(out)
	// ioutil.WriteFile("out.txt", out, 0644)
	//fmt.Println("\n\n\n\n\n\n\n\n", ResultC)
	// fmt.Println("a1:\n", a1, "\n", "b1:\n", b1, "\n", "c1:\n", c1, "\n", "d1:\n", d1)
	// fmt.Println(a, "\n", b, "\n", c, "\n")
	// fmt.Println(ResultC)
	fmt.Println()
	{
		for _, v := range ResultC {
			for _, val := range v {
				fmt.Printf("%0.3f \t ", val)
			}
			fmt.Println()
		}
	}
}

func read(filename string, flag int) error {
	file, err := os.OpenFile(filename, os.O_RDONLY, os.ModePerm)
	if err != nil {
		fmt.Println(err)
	}
	inputReader := bufio.NewReader(file)
	// lineCounter := 0

	var lengthi int64
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
				if Apartment[Test1[0]] != nil {
					posi = Apartment[Test1[0]]
				}

				WorkType[lengthi].Name = Test1[0]
				WorkType[lengthi].Type = num
				WorkType[lengthi].Require = Hope.Require
				WorkType[lengthi].Situation = Hope.Situation
				WorkType[lengthi].salary = Hope.salary
				// fmt.Println(lengthi)
				if lengthi < 9 {
					lengthi++
				}

				posi[num] = *Hope
				{
					jobRequire[num] = Hope.Require
				}
				// fmt.Println(jobRequire)
				Apartment[Test1[0]] = posi
				// fmt.Println(Apartment, "***", Test1)
				// fmt.Println(Test1[1], Test1[0], Apartment[Test1[0]], "-----------------", posi)
			}
		}
	}

	if flag == 1 {
		PersonInfo = PersonInfo[1:]
	}

	defer file.Close()
	return nil
}
