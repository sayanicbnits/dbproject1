package postgres

import (
	"context"
	"log"

	//"github.com/sayani/dbproject1/graph/controller"

	//"github.com/sayani/dbproject1/graph/controller"
	"github.com/sayani/dbproject1/graph/controller"
	Entity "github.com/sayani/dbproject1/graph/entity"
	//	Entity "github.com/sayani/dbproject1/graph/entity"
)

var k int

type student struct {
	Ro int
	Na string
	Co string
	To int
}

func Check(c int, d string, st string) {
	if pool == nil {
		pool = GetPool()
	}
	var roll int
	//tx, err := pool.Begin(context.Background())
	querystring := `insert into my_tab (roll,name) values($1,$2) returning roll`
	err := pool.QueryRow(context.Background(), querystring, &c, &d).Scan(&roll)
	if err != nil {
		//	log.Print("error")
		log.Println("err", err)

	}
	//	querystring := "update my_table set name = $1 where roll=$2"
	//	commandTag, err := tx.Exec(context.Background(), querystring, &c, &d)
	query := `insert into my_courses (roll,course) values($1,$2)`
	err1 := pool.QueryRow(context.Background(), query, &roll, &st)
	if err1 != nil {
		//	log.Print("error")
		log.Println("err", err)
	}
	/*	if commandTag.RowsAffected() != 1 {
			// response.Message = "Invalid kpi target data"
			// response.Error = true
			log.Println("2")
		} else if commandTag.RowsAffected() == 1 {
			// response.Message = "kpi target successfully updated"
			// response.Error = false
			log.Println("updated")
			//return 1
		}

		txErr := tx.Commit(context.Background())
		if txErr != nil {
			// response.Message = "Failed to commit customer data"
			// response.Error = true
			log.Println(txErr, "4")
		}*/
}
func Check1(roll int, course string, sem1 int, sem2 int, sem3 int) {
	if pool == nil {
		pool = GetPool()
	}
	tx, err := pool.Begin(context.Background())
	marks := sem1 + sem2 + sem3
	querystring := `update my_courses set sem1 = $1,sem2 = $2,sem3 = $3,total=$6 where roll =$4 and course = $5`
	//	querystring := "update my_table set name = $1 where roll=$2"
	commandTag, err := tx.Exec(context.Background(), querystring, &sem1, &sem2, &sem3, &roll, &course, &marks)
	if err != nil {
		//	log.Print("error")
		log.Println("err", err)
	}
	if commandTag.RowsAffected() != 1 {
		// response.Message = "Invalid kpi target data"
		// response.Error = true
		log.Println("no rows affected ")
	} else if commandTag.RowsAffected() == 1 {
		// response.Message = "kpi target successfully updated"
		// response.Error = false
		log.Println("updated")
		//return 1
	}

	txErr := tx.Commit(context.Background())
	if txErr != nil {
		// response.Message = "Failed to commit customer data"
		// response.Error = true
		log.Println(txErr, "error")
	}

}
func Testfunc(roll int) []Entity.En {
	//	var marks int
	if pool == nil {
		pool = GetPool()
	}
	//var p *student
	var arr []student
	var s student
	// var roll1, total int
	// var name, course string
	//var err error
	//log.Println("query1")
	tx, err := pool.Begin(context.Background())
	// if err != nil {
	// 	log.Println(err)
	// }
	//var tx pgx.Tx
	//log.Println("query2")
	query := `select t.roll,t.name,c.course,c.total from my_tab t left outer join my_courses c on t.roll=c.roll`
	//log.Println(query)
	rows, err := tx.Query(context.Background(), query)
	log.Println("query3")
	log.Println(rows, "@@@")
	if err != nil {
		//	log.Print("error")
		log.Println("err", err)
	}
	defer rows.Close()
	var ar []Entity.En
	//iterating over each row
	for rows.Next() {
		//storing value of each column using scan
		var ent Entity.En
		err1 := rows.Scan(
			&ent.Roll,
			&ent.Name,
			&ent.Course,
			&ent.Total,
		)
		if err1 != nil {
			log.Println(err)
		}
		log.Println(",,,,,,,,,,", ent.Course)
		// ent.Roll = roll1
		// ent.Name = name
		// ent.Course = course
		// ent.Total = total
		//s.Ro = roll1
		// s.Na = name
		// s.Co = course
		// s.To = total
		arr = append(arr, s)
		ar = append(ar, ent)
		//log.Println("\n roll: ", roll1, "name: ", name, "course: ", course, "total: ", total)

	}

	controller.DisplayExcel1(ar)

	log.Println("ent arr: ", ar)
	err = rows.Err()

	if err != nil {
		log.Println(err)
	}
	log.Println(arr)
	return ar

	// txErr := tx.Commit(context.Background())
	// if txErr != nil {
	// 	// response.Message = "Failed to commit customer data"
	// 	// response.Error = true
	// 	log.Println(txErr, "4")
	// }
	//log.Println(marks)
	//err = pool.QueryRow(context.Background(), query, &roll).Scan(&marks)
	//commandTag,err
	// if err != nil {
	// 	log.Println("err", err)
	// }
	//  if commandTag.RowsAffected() != 1 {
	// // 	 response.Message = "Invalid kpi target data"
	// // 	 response.Error = true
	//  	log.Println("2")
	// // }
	// for rows.Next() {
	// 	//	rows.Scan()
	// }

}

func Check2(c int, d string) {
	if pool == nil {
		pool = GetPool()
	}
	//tx, err := pool.Begin(context.Background())
	var roll int
	querystring := `delete from my_courses where roll = $1 returning roll`
	//	querystring := "update my_table set name = $1 where roll=$2"
	//commandTag, err := tx.Exec(context.Background(), querystring, &c)
	err := pool.QueryRow(context.Background(), querystring, &c).Scan(&roll)
	if err != nil {
		//	log.Print("error")
		log.Println("err", err)
	}
	query := `delete from my_tab where roll = $1 returning roll `
	err1 := pool.QueryRow(context.Background(), query, &roll).Scan(&roll)
	if err1 != nil {
		log.Println("err1", err1)
	}
	// if commandTag.RowsAffected() != 1 {
	// 	// response.Message = "Invalid kpi target data"
	// 	// response.Error = true
	// 	log.Println("2")
	// } else if commandTag.RowsAffected() == 1 {
	// 	// response.Message = "kpi target successfully updated"
	// 	// response.Error = false
	// 	log.Println("updated")
	// 	//return 1
	// }

	// txErr := tx.Commit(context.Background())
	// if txErr != nil {
	// 	// response.Message = "Failed to commit customer data"
	// 	// response.Error = true
	// 	log.Println(txErr, "4")
	// }
}
func Check3(c int, d int) {
	if pool == nil {
		pool = GetPool()
	}
	var name string
	query := `SELECT name from my_tab`
	rows, err1 := pool.Query(context.Background(), query)
	if err1 != nil {
		log.Println(err1)
	} else {
		var arr []string
		for rows.Next() {
			err := rows.Scan(&name)
			if err != nil {
				log.Fatal(err)
			}
			arr = append(arr, name)
			log.Println(name)
		}
		log.Println(arr)
	}
	// if err != nil {
	//     log.Fatal(err)
	// }
	//query :=db.Query()
	//tx, err := pool.Begin(context.Background())
	//querystring := `delete from my_tab where roll = $1`
	//	querystring := "update my_table set name = $1 where roll=$2"
	//commandTag, err := tx.Exec(context.Background(), querystring, &c)
	/*var rollmax, rollmin int
	query := `select max(roll),min(roll) from my_tab `
	err := pool.QueryRow(context.Background(), query).Scan(&rollmax, &rollmin)
	if err != nil {
		log.Println(err)
	}
	log.Println(rollmax, rollmin)
	var name string
	i := rollmin
	for i <= rollmax {
		query1 := `select name from my_tab where roll=$1`
		err := pool.QueryRow(context.Background(), query1, &i).Scan(&name)
		if err != nil {
			log.Println(err)
		}
		i++
		log.Println(name)
	}

	/*if err != nil {
		//	log.Print("error")
		log.Println("err", err)
	}
	if commandTag.RowsAffected() != 1 {
		// response.Message = "Invalid kpi target data"
		// response.Error = true
		log.Println("2")
	} else if commandTag.RowsAffected() == 1 {
		// response.Message = "kpi target successfully updated"
		// response.Error = false
		log.Println("updated")
		//return 1
	}

	txErr := tx.Commit(context.Background())
	if txErr != nil {
		// response.Message = "Failed to commit customer data"
		// response.Error = true
		log.Println(txErr, "4")
	}*/
}

func Check4(c int, d int) []string {
	if pool == nil {
		pool = GetPool()
	}
	var name string
	query := `SELECT name from my_tab`
	rows, err1 := pool.Query(context.Background(), query)
	if err1 != nil {
		log.Println(err1)
	} else {
		var arr []string
		for rows.Next() {
			err := rows.Scan(&name)
			if err != nil {
				log.Fatal(err)
			}
			arr = append(arr, name)
			log.Println(name)
		}
		log.Println(arr)
		return arr
	}
	arr := []string{"error"}
	return arr
}
func Check01(c int, d string) []string {
	var course string
	query := `insert into my_courses (roll,course) values ($1,$2)returning course`
	err := pool.QueryRow(context.Background(), query, &c, &d).Scan(&course)
	if err != nil {
		log.Println(err)
	}
	coursestr := []string{course}
	return coursestr

}
func Check11(c int, d string) []string {
	if pool == nil {
		pool = GetPool()
	}
	var course string
	query := `update my_courses set course=$2 where roll =$1 returning course`
	err := pool.QueryRow(context.Background(), query, &c, &d).Scan(&course)
	if err != nil {
		log.Println(err)
	}
	coursestr := []string{course}
	return coursestr

}
func Check21(c int, d string) []int {
	if pool == nil {
		pool = GetPool()
	}
	var roll int
	query := `delete from my_courses where roll=$1 returning roll`
	err := pool.QueryRow(context.Background(), query, &c).Scan(&roll)
	if err != nil {
		log.Println(err)
	}
	coursestr := []int{roll}
	log.Println(roll)
	return coursestr

}
func Calculate(x int, y int, z int, ro int, c string) {
	//	query := `insert into my_courses (sem1,sem2,sem3) values ($1,$2,$3)`

	//var roll int
	query := `update my_courses set sem1 = $1 ,sem2 = $2 ,sem3 = $3`
	log.Println(query)
	log.Println(&y, &c, &ro, &z, &c)
	log.Println(x, y, z, ro, c)
	//	pool.end()

	err := pool.QueryRow(context.Background(), query, x, y, z)
	//	log.Println("roll", err)

	log.Println("xxxxxxxxxxxxxxxxx")

	if err != nil {
		log.Println("err", err)
	}

	// if &x == nil || &y == nil || &z == nil || &ro == nil || &c == nil {
	// 	log.Println("yes   hjkhbghvccv")
	// }
	//var inputArgs []interface{}
	// query := `update my_courses set  sem3 = $1 where roll = $2`
	// inputArgs = append(inputArgs, z, ro)
	// log.Println(inputArgs...)
	// _, err := pool.Exec(context.Background(), query, &z, &ro)
	// if err != nil {
	// 	log.Println(err.Error())
	// }
	// log.Println("ABCD")

	//sum = x + y + z
	//fmt.Println("sum: ", sum)
	/*query1 := `update my_courses set total =$1 where roll = $2 and course = $3`
	err1 := pool.QueryRow(context.Background(), query1, &sum, &ro, &c)
	if err1 != nil {
		log.Println("err1: ", err1)
	}*/

}
func Test() {
	// query := `insert into my_courses1 (sem1) values (200)`
	// log.Println("abcd")
	// err := pool.QueryRow(context.Background(), query)
	// log.Println("abcdedfhg")

	// if err != nil {
	// 	log.Println("err: ", err)
	// }
	//var marks int
	//query1 := `insert into my_courses1 (sem1) values (200)`
	// {
	// 	text: 'INSERT INTO users(name, email) VALUES($1, $2)',
	// 	values: ['brianc', 'brian.m.carlson@gmail.com'],
	//   }
	// callback
	// 	client.query(
	// 		`insert into my_courses1 (sem1) values ($1)`,
	// 		[&marks]
	// 	)
	// 	   client.query(query1, (err, res) => {
	// 	 	if (err) {
	// 	 	  console.log(err.stack)
	// 	 	} else {
	// 	 	  console.log(res.rows[0])
	// 	 	}
	// 	   }
	// 	   client
	//   .query(query)
	//   .then(res => console.log(res.rows[0]))
	//   .catch(e => console.error(e.stack))

	// 	text := `insert into my_courses1 (sem1) values ($1)`

	// values := [1]int {500}
	// // callback
	// client.query(text, values, (err) => {
	//   if (err) {
	//     console.log(err.stack)
	//   } else {
	//     console.log(res.rows[0])
	//     // { name: 'brianc', email: 'brian.m.carlson@gmail.com' }
	//   }
	// })
	// // promise
	// client
	//   .query(text, values)
	//   .then(res => {
	//     console.log(res.rows[0])
	//     // { name: 'brianc', email: 'brian.m.carlson@gmail.com' }
	//   })
	//   .catch(e => console.error(e.stack))
}

func Updatename(roll int, name string) string {
	// query := `update my_tab set name = $1 where roll =$2`
	// err := pool.QueryRow(context.Background(), query, &name, &roll)
	// if err != nil {
	// 	log.Println("err", err)
	// 	return "error"
	// } else {
	// 	return "updated"
	// }

	if pool == nil {
		pool = GetPool()
	}
	tx, err := pool.Begin(context.Background())
	//marks := sem1 + sem2 + sem3
	querystring := `update my_tab set name = $1 where roll =$2`
	//	querystring := "update my_table set name = $1 where roll=$2"
	commandTag, err := tx.Exec(context.Background(), querystring, &name, &roll)
	if err != nil {
		//	log.Print("error")
		log.Println("err", err)
		return "error"
	}
	if commandTag.RowsAffected() != 1 {
		// response.Message = "Invalid kpi target data"
		// response.Error = true
		log.Println("2")
		return "rows not affected"
	} else if commandTag.RowsAffected() == 1 {
		// response.Message = "kpi target successfully updated"
		// response.Error = false
		log.Println("rows affected")
		//return 1

	}

	txErr := tx.Commit(context.Background())
	if txErr != nil {
		// response.Message = "Failed to commit customer data"
		// response.Error = true
		log.Println(txErr, "4")
		return "not commited"
	}
	return "rows affected and updated"
}

func DynamicAdd(a *int, b *int, c *int, d *int, e *int) {
	f := *a + *b + *c + *d + *e

	controller.DisplayExcelAdd(*a, *b, *c, *d, *e, f)
}

func Dynamicinput(arr []*int) {
	k = 0
	sum := 0
	for i := 0; i < len(arr); i++ {

		sum += *arr[i]
	}
	log.Println(arr)
	arr = append(arr, &sum)
	controller.DisplayExcelAdd1(arr, sum)
	// 	log.Println("arr ", arr)
	// 	log.Println("len(arr): ", len(arr))

	// 	for i := 0; i < len(arr); i++ {
	// 		log.Println("arr[i]", arr[i])
	// 		log.Println("*arr[i].A", *arr[i].A)
	// 		log.Println("*arr[i].B", *arr[i].B)
	// 		log.Println("*arr[i].C", *arr[i].C)
	// 		log.Println("*arr[i].D", *arr[i].D)
	// 		log.Println("*arr[i].E", *arr[i].E)
	// 		sum := 0
	// 		if *arr[i].C == 0 && *arr[i].D == 0 && *arr[i].E == 0 {
	// 			sum += *arr[i].A + *arr[i].B
	// 			ar := []int{*arr[i].A, *arr[i].B}
	// 			controller.DisplayExcelAdd1(ar, sum)
	// 		} else if *arr[i].D == 0 && *arr[i].E == 0 {
	// 			sum += *arr[i].A + *arr[i].B + *arr[i].C
	// 			ar := []int{*arr[i].A, *arr[i].B, *arr[i].C}
	// 			controller.DisplayExcelAdd1(ar, sum)
	// 		} else if *arr[i].E == 0 {
	// 			sum += *arr[i].A + *arr[i].B + *arr[i].C + *arr[i].D
	// 			ar := []int{*arr[i].A, *arr[i].B, *arr[i].C, *arr[i].D}
	// 			controller.DisplayExcelAdd1(ar, sum)
	// 		} else {
	// 			sum += *arr[i].A + *arr[i].B + *arr[i].C + *arr[i].D + *arr[i].E
	// 			ar := []int{*arr[i].A, *arr[i].B, *arr[i].C, *arr[i].D, *arr[i].E}
	// 			controller.DisplayExcelAdd1(ar, sum)
	// 		}

	// 		// v := reflect.ValueOf(*arr[i])
	// 		// log.Println("v: ", v)
	// 		// //	var sum *int
	// 		// for j := 0; j < v.NumField(); j++ {
	// 		// 	//	sum +=
	// 		// 	log.Println(*arr[i])
	// 		// }
	// 		//log.Println("sum: ", sum)

	// 	}
	k++
}
