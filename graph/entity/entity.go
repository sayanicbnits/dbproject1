package Entity

type En struct {
	Roll   int
	Name   string
	Course string
	Total  int
}
type EnterDetails struct {
	Name   string
	Course []string
}
type ShowDetails struct {
	Details []ObtainedMarks
}
type ObtainedMarks struct {
	Roll   int
	Name   string
	Course string
	Sem1   int
	Sem2   int
	Sem3   int
	Total  int
}
