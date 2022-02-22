package postgres1

import (
	"context"
	"log"

	Entity "github.com/sayani/dbproject1/graph/entity"
)

func DatabaseLogic(entmap Entity.EnterDetails) []Entity.ObtainedMarks {
	// 	return j
	if pool == nil {
		//pool :=
		pool = GetPool()
	}
	//var ptr []*Entity.ObtainedMarks
	//var del Entity.EnterDetails
	var del1 Entity.ObtainedMarks
	var del2 []Entity.ObtainedMarks
	tx, err := pool.Begin(context.Background())
	if err != nil {
		log.Println("err: ", err)
	}
	log.Println("name:", entmap.Name)
	query := `select t.roll,name,course,sem1 ,sem2 ,sem3,total from my_tab t inner join my_courses c on t.roll=c.roll where t.name=$1 `
	rows, err := tx.Query(context.Background(), query, entmap.Name)
	if err != nil {
		log.Println("err", err)
	}
	log.Println("kkkkkk")
	defer rows.Close()
	for rows.Next() {
		log.Println("LLLLLL")
		err1 := rows.Scan(
			&del1.Roll,
			&del1.Name,
			&del1.Course,
			&del1.Sem1,
			&del1.Sem2,
			&del1.Sem3,
			&del1.Total,
		)
		log.Println("SCAN")
		if err1 != nil {
			log.Println("ERR1", err1)
		}
		log.Println("....", del1)
		del2 = append(del2, del1)
	}
	log.Println("hhhhhhhhhhh")
	return del2
}
