package mapper

import (
	"log"

	Entity "github.com/sayani/dbproject1/graph/entity"
	"github.com/sayani/dbproject1/graph/model"
)

func EnterDetailsModel_entity(input model.EnterDetails) Entity.EnterDetails {
	var ent Entity.EnterDetails
	ent.Name = input.Name
	log.Println(".....", ent.Name)
	log.Println("?????????????????????")
	for i := range input.Course {
		ent.Course = append(ent.Course, input.Course[i])
		log.Println("///////////")
	}
	// var course []string
	// course = append(course, )
	log.Println("ffff", ent)
	return ent

}
