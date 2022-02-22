package mapper

import (
	"log"

	Entity "github.com/sayani/dbproject1/graph/entity"
	"github.com/sayani/dbproject1/graph/model"
)

func ShowDetailsEntity_Model(entpost []Entity.ObtainedMarks) model.Showdetails {
	//var mod []Entity.ShowDetails
	log.Println("/////////////....", len(entpost))

	var mod2 []model.ObtainedMarks
	// log.Println("/////////////....{{{{", len(mod2))
	for i := range entpost {
		// 	log.Println("/////////////....{{{{ len(mod2)")
		// 	mod2[i].Roll = entpost[i].Roll
		// 	mod2[i].Name = entpost[i].Name
		// 	mod2[i].Course = entpost[i].Course
		// 	mod2[i].Sem1 = entpost[i].Sem1
		// 	mod2[i].Sem2 = entpost[i].Sem2
		// 	mod2[i].Sem3 = entpost[i].Sem3
		// 	mod2[i].Total = entpost[i].Total
		mod2 = append(mod2, model.ObtainedMarks(entpost[i]))
	}
	// //mod = append(mod)
	var show model.Showdetails
	//log.Println("gfgfgfgf")
	for i := range entpost {
		show.Details = append(show.Details, &mod2[i])
	}
	return show

}
