package controller1

import (
	Entity "github.com/sayani/dbproject1/graph/entity"
	"github.com/sayani/dbproject1/graph/mapper"
	"github.com/sayani/dbproject1/graph/model"
	"github.com/sayani/dbproject1/graph/postgres1"
)

func LogicOfDisplay(input model.EnterDetails) model.Showdetails {
	var entmap Entity.EnterDetails
	// var entpost Entity.ShowDetails
	entmap = mapper.EnterDetailsModel_entity(input)
	entpost := postgres1.DatabaseLogic(entmap)
	result := mapper.ShowDetailsEntity_Model(entpost)
	return result
}
