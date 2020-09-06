package model

type student struct {
	Name  string
	score float64
}

//工厂
func NewStudent(n string, s float64) *student {
	return &student{n, s}
}

func (s *student) GetScore() float64 {
	return s.score
}
