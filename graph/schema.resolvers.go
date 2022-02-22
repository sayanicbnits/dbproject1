package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"
	"log"

	"github.com/sayani/dbproject1/graph/controller1"
	"github.com/sayani/dbproject1/graph/generated"
	"github.com/sayani/dbproject1/graph/model"
	"github.com/sayani/dbproject1/graph/postgres"
)

func (r *mutationResolver) CreateTodo(ctx context.Context, input model.NewTodo) (*model.Todo, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) DeleteVal(ctx context.Context, input model.Inputstr2) (*model.Outputstr2, error) {
	//panic(fmt.Errorf("not implemented"))
	postgres.Check2(input.Roll, input.Name)
	s := model.Outputstr2{"abcd"}
	return &s, nil
}

func (r *mutationResolver) UpdateMarks(ctx context.Context, input model.Marks) (*model.Total, error) {
	//query := `update my_courses set sem1 =$1, sem2 = $2, sem3 =$3 where roll = $4 and course = $5`
	//panic(fmt.Errorf("not implemented"))
	//fmt.Println(input.X, input.Y, input.Z, input.Ro, input.C)
	postgres.Calculate(input.X, input.Y, input.Z, input.Ro, input.C)
	// var i int
	// postgres.Test()
	s := model.Total{1}
	return &s, nil
}

func (r *mutationResolver) PutCourse(ctx context.Context, input model.Inputcourses) (*model.Outpurcourses, error) {
	//panic(fmt.Errorf("not implemented"))
	coursestr := postgres.Check01(input.X, input.Y)
	var str []*string
	for i := range coursestr {
		str = append(str, &coursestr[i])
	}
	s := model.Outpurcourses{str}
	return &s, nil
}

func (r *mutationResolver) UpdateCourse(ctx context.Context, input model.Inputcourses) (*model.Outpurcourses, error) {
	//panic(fmt.Errorf("not implemented"))
	coursestr := postgres.Check11(input.X, input.Y)
	// strr := "sayani"
	// strr1 := "das"
	// var str1 []*string
	// str1 = append(str1, &strr)
	// str1 = append(str1, &strr1)
	// fmt.Println("my", str1)
	var str []*string
	for i, ele := range coursestr {
		fmt.Println(i, ele)
		str = append(str, &ele)
		//	fmt.Println("mine", str)
	}
	s := model.Outpurcourses{str}
	return &s, nil
}

func (r *mutationResolver) DeleteCourse(ctx context.Context, input model.Inputcourses) (*model.Outpurcourses, error) {
	//panic(fmt.Errorf("not implemented"))
	postgres.Check21(input.X, input.Y)
	var str []*string
	/*for i := range coursestr {
		str = append(str, &coursestr[i])
	}*/
	s := model.Outpurcourses{str}
	return &s, nil
}

func (r *mutationResolver) DisplayMarksPerSub(ctx context.Context, input model.Testmarks) (*model.Testtotal, error) {
	//panic(fmt.Errorf("not implemented"))
	log.Println("123")
	log.Println("IN: ", input.Roll)
	ar := postgres.Testfunc(input.Roll)
	fmt.Println("ar : ", ar)
	//arr[0]=model.Student
	//[]*model.Student(arr)
	//model.Student = s.
	//s2 := str
	//arr01 :=model.
	//s2.roll
	var arr12 []*model.Student

	//	j := 0

	for i := range ar {
		var st model.Student
		st.Roll = ar[i].Roll
		st.Name = ar[i].Name
		st.Course = ar[i].Course
		fmt.Println(st.Course)
		st.Total = ar[i].Total
		arr12 = append(arr12, &st)
		fmt.Println("st address", &st)

		fmt.Println("st", st)

		fmt.Println(arr12)
		//	fmt.Println("inside", arr12[i])
	}
	//fmt.Println("outside", arr12[0])
	//fmt.Println("outside", arr12[1])
	//var st model.Student
	//st.Roll = sptr.Ro
	// st.Name = sptr.Na
	// st.Course = sptr.Co
	// st.Total = sptr.To

	//arr12[0] = *model.Student(&arr[0])
	var s model.Testtotal
	s = model.Testtotal{arr12}
	return &s, nil
}

func (r *mutationResolver) InsertVal(ctx context.Context, input model.Inputstr) (*model.Outputstr, error) {
	//	panic(fmt.Errorf("not implemented"))
	p := input.X
	q := input.Y
	st := input.P
	postgres.Check(p, q, st)
	s := model.Outputstr{"abcd"}
	return &s, nil
}

func (r *mutationResolver) UpdateVal(ctx context.Context, input model.Inputstr1) (*model.Outputstr1, error) {
	//panic(fmt.Errorf("not implemented"))
	postgres.Check1(input.Roll, input.Course, input.Sem1, input.Sem2, input.Sem3)
	s := model.Outputstr1{11}
	return &s, nil
}

func (r *mutationResolver) UpdateNameroll(ctx context.Context, input model.Inputstr2) (*model.Outputstr2, error) {
	//	panic(fmt.Errorf("not implemented"))
	str := postgres.Updatename(input.Roll, input.Name)
	s := model.Outputstr2{str}
	return &s, nil
}

func (r *mutationResolver) DynamicAdd(ctx context.Context, input *model.Inputdynamic) (*model.Outputdynamic, error) {
	log.Println(input.Arr)
	postgres.Dynamicinput(input.Arr)
	p := 0
	q := 0
	t := 0
	u := 0
	v := 0
	w := 0
	s := model.Outputdynamic{&p, &q, &t, &u, &v, &w}
	return &s, nil

	//panic(fmt.Errorf("not implemented"))
	// p := 0
	// q := 0
	// t := 0
	// u := 0
	// v := 0
	// w := 0
	// i := &p
	// j := &q
	// k := &t
	// // l := &u
	// // m := &v
	// // n := &w

	// if input.C == nil && input.D == nil && input.E == nil {
	// 	postgres.DynamicAdd(input.A, input.B, i, j, k)
	// 	f := *input.A + *input.B
	// 	s := model.Outputdynamic{input.A, input.B, i, j, k, &f}
	// 	return &s, nil
	// } else if input.D == nil && input.E == nil {
	// 	postgres.DynamicAdd(input.A, input.B, input.C, j, k)
	// 	f := *input.A + *input.B + *input.C
	// 	s := model.Outputdynamic{input.A, input.B, input.C, j, k, &f}
	// 	return &s, nil
	// } else if input.E == nil {
	// 	postgres.DynamicAdd(input.A, input.B, input.C, input.D, k)
	// 	f := *input.A + *input.B + *input.C + *input.D
	// 	s := model.Outputdynamic{input.A, input.B, input.C, input.D, k, &f}
	// 	return &s, nil
	// } else {
	// 	postgres.DynamicAdd(input.A, input.B, input.C, input.D, input.E)
	// 	f := *input.A + *input.B + *input.C + *input.D + *input.E
	// 	s := model.Outputdynamic{input.A, input.B, input.C, input.D, input.E, &f}
	// 	return &s, nil

	// }
	// s := model.Outputdynamic{i, j, k, l, m, n}
	// return &s, nil
}

func (r *queryResolver) Todos(ctx context.Context) ([]*model.Todo, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) DisplayVal(ctx context.Context, input model.Inputvalues) (*model.Outputvalues, error) {
	//panic(fmt.Errorf("not implemented"))
	postgres.Check3(input.X, input.Y)
	s := model.Outputvalues{"displayed"}
	return &s, nil
}

func (r *queryResolver) DisplayVal1(ctx context.Context, input model.Inputvalues1) (*model.Outputvalues1, error) {
	//	panic(fmt.Errorf("not implemented"))
	arr := postgres.Check4(input.X, input.Y)
	fmt.Println(arr)
	var arr1 []*string
	//	var arr2 []*string
	//var arr2 []string
	//	arr1 = &arr
	//	arr1=&arr2

	//	fmt.Println(arr1)
	//	s := model.Outputvalues1{arr2}
	//	arr1 = arr
	// i := 0
	for i := range arr {
		/*if s == " " {
			arr1 = append(arr1, &s)
			continue
		}*/
		arr1 = append(arr1, &arr[i])
	}
	s := model.Outputvalues1{arr1}
	return &s, nil
}

func (r *queryResolver) Display(ctx context.Context, input model.EnterDetails) (*model.Showdetails, error) {
	//panic(fmt.Errorf("not implemented"))
	s := controller1.LogicOfDisplay(input)
	//mapper.EnterDetailsModel_entity(input)
	return &s, nil
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }

// !!! WARNING !!!
// The code below was going to be deleted when updating resolvers. It has been copied here so you have
// one last chance to move it out of harms way if you want. There are two reasons this happens:
//  - When renaming or deleting a resolver the old code will be put in here. You can safely delete
//    it when you're done.
//  - You have helper methods in this file. Move them out to keep these resolver files clean.
func (r *mutationResolver) Testcase(ctx context.Context, input *model.Testmarks) (*model.Testtotal, error) {
	//panic(fmt.Errorf("not implemented"))
	// log.Println("123")
	// str := postgres.Testfunc(input.Roll)
	// fmt.Println("str: ", str)
	//arr[0]=model.Student
	//[]*model.Student(arr)
	//model.Student = s.
	arr12 := []*model.Student{}
	//arr12[0] = *model.Student(&arr[0])
	s := model.Testtotal{arr12}
	return &s, nil
}
