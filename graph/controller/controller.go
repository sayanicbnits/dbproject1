package controller

import (

	//"github.com/360EntSecGroup-Skylar/excelize"

	"fmt"
	"log"
	"reflect"
	"strconv"

	//"github.com/gin-gonic/gin"
	Entity "github.com/sayani/dbproject1/graph/entity"
	//"github.com/sayani/dbproject1/graph/postgres"
	"github.com/tealeg/xlsx"
)

var j, k int
var xl *xlsx.File
var sheet *xlsx.Sheet
var err error

// func sheetcreate() {
// 	file := xlsx.NewFile()
// 	sheet, err := file.AddSheet("Sheet1")
// 	if err != nil {
// 		fmt.Printf(err.Error())
// 	}

// }

func DisplayExcel1(ar []Entity.En) {
	var file *xlsx.File
	var sheet *xlsx.Sheet
	var row *xlsx.Row
	var cell *xlsx.Cell
	var err error
	//j := 0
	file = xlsx.NewFile()
	sheet, err = file.AddSheet("Sheet1")
	if err != nil {
		fmt.Printf(err.Error())
	}
	for i, _ := range ar {
		row = sheet.AddRow()
		// cell = row.AddCell()
		// cell.Value = strconv.Itoa(v.Roll)
		// cell = row.AddCell()
		// cell.Value = v.Name
		// cell = row.AddCell()
		// cell.Value = v.Course
		// cell = row.AddCell()
		// cell.Value = strconv.Itoa(v.Total)
		v := reflect.ValueOf(ar[i])
		log.Println("reflect.ValueOf: ", v)

		//values := make([]interface{}, v.NumField())

		for i := 0; i < v.NumField(); i++ {
			cell = row.AddCell()
			//	values[i] = v.Field(i).Interface()
			//if reflect.TypeOf(v.Field(i).Kind())== reflect.Kind.(INT){
			//if i == 0 || i == 3 {
			//	if	reflect.TypeOf(v.Field(i).Interface())== reflect.Type {
			switch v.Field(i).Interface().(type) {
			case int:
				cell.Value = v.Field(i).Interface().(string)
				//log.Println("type: ", reflect.TypeOf(v.Field(i).Interface()))
				//log.Println("cell value: ", cell.Value)
			case string:
				cell.Value = v.Field(i).Interface().(string)
				//log.Println("cell value: ", cell.Value)
			default:
				fmt.Println("Error")

			}

			err = file.Save("MyXLSXFile.xlsx")
			if err != nil {
				fmt.Printf(err.Error())
			}

		}
		i++

		// file := xlsx.NewFile()
		// sheet, err := file.AddSheet("Report")
		// if err != nil {
		// 	c.JSON(http.StatusInternalServerError, err.Error())
		// 	return
		// }
		// for _, v := range ar {
		// 	row := sheet.AddRow()
		// 	cell1 := row.AddCell()
		// 	cell1.Value = strconv.Itoa(v.Roll)
		// 	cell2 := row.AddCell()
		// 	cell2.Value = v.Name
		// 	cell3 := row.AddCell()
		// 	cell3.Value = v.Course
		// 	cell4 := row.AddCell()
		// 	cell4.Value = strconv.Itoa(v.Total)

		// }
		// var b bytes.Buffer
		// if err := file.Write(&b); err != nil {
		// 	c.JSON(http.StatusInternalServerError, err.Error())
		// 	return
		// }
		// downloadName := time.Now().UTC().Format("data-20060102150405.xlsx")
		// c.Header("Content-Description", "File Transfer")
		// c.Header("Content-Disposition", "attachment; filename="+downloadName)
		// c.Data(http.StatusOK, "application/octet-stream", b.Bytes())

		// Create a new sheet.
		//index := xlsx.NewSheet("Sheet2")
		// Set value of a cell.
		// var axis int = 2
		// i := 2
		// for i := 0; i < 10; i++ {
		// 	sheet, err := xlsx.AddSheet(fmt.Sprintf("sheet-%d", i))
		// 	if err != nil {
		// 		log.Println(err)
		// 	}
		// 	row := sheet.AddRow()
		// 	cell := row.AddCell()
		// 	cell.Value = "I am a cell!"
		// }
		//
	}

	// func Header(s1, s2 string) {
	// 	panic("unimplemented")
	// }

	// func Data(i int, s string, b []byte) {
	// 	panic("unimplemented")
	// }

	// func JSON(i int, s string) {
	// 	panic("unimplemented")
	// }
}

func DisplayExcelAdd(a int, b int, c int, d int, e int, f int) {
	arr := []int{a, b, c, d, e, f}
	file := xlsx.NewFile()
	sheet, err := file.AddSheet("Sheet1")
	if err != nil {
		log.Println("err: ", err)
	}
	// font := &xlsx.Font{Color: "red"}
	// style := xlsx.NewStyle()
	// style.Font = *font

	row := sheet.AddRow()
	for i := range arr {
		if arr[i] != 0 {
			cell := row.AddCell()

			cell.Value = strconv.Itoa(arr[i])
			//cell.SetStyle(style)
		}

	}
	err = file.Save("dynamic.xlsx")
	if err != nil {
		fmt.Printf(err.Error())
	}

}

func DisplayExcelAdd1(arr []*int, sum int) {

	if j == 0 && k == 0 {
		// log.Println("i,j")
		// log.Println("j: ", j)
		// log.Println("k: ", k)
		xl = xlsx.NewFile()
		//defer xl.Close()

		//create a new sheet
		sheet, err := xl.AddSheet("Sheet1")

		if err != nil {
			fmt.Printf(err.Error())
		}
		for i := range arr {
			log.Println("j1: ", j)
			log.Println("k1: ", k)
			//access by indexes
			cell := sheet.Cell(j, k)
			k++
			cell.SetValue(*arr[i])

		}
		k = 0
		j++
		//  for iRow := 1; iRow < 7; iRow++ {
		//  	//access by indexes
		//  	cell := sheet.Cell(1, iRow)
		//  	cell.SetValue(iRow)
		//  }
		err = xl.Save("dynamic.xlsx")
		if err != nil {
			fmt.Printf(err.Error())
		}
	} else {
		log.Println("k,l")

		// row := sheet.AddRow()
		// log.Println("row: ", row)
		// for i := range arr {
		// 	if arr[i] != nil {
		// 		cell := row.AddCell()

		// 		cell.Value = strconv.Itoa(*arr[i])
		// 		//	cell.SetStyle(style)
		// 	}

		// }

		for i := range arr {
			// 	//access by indexes
			// log.Println("j1: ", j)
			// log.Println("k1: ", k)
			f, err := xlsx.OpenFile("dynamic.xlsx")
			if err != nil {
				log.Println(err)
			}
			for _, sheet := range f.Sheets {
				log.Println("....")
				log.Println("j2: ", j)
				log.Println("k2: ", k)
				cell := sheet.Cell(2, 2)
				//k++
				// log.Println("k")
				// log.Println(*arr[i])
				cell.SetValue(*arr[i])
			}
			k++

		}

		//  for iRow := 1; iRow < 7; iRow++ {
		//  	//access by indexes
		//  	cell := sheet.Cell(1, iRow)
		//  	cell.SetValue(iRow)
		//  }
		err = xl.Save("dynamic.xlsx")
		if err != nil {
			fmt.Printf(err.Error())
		}
		k = 0
		j++

	}
}
