package schema


type Vector struct {
	X int `json:"x"`
	Y int `json:"y"`
	Z int `json:"z"`
}

func NewVector(x int, y int, z int) *Vector {
	return &Vector{
		X: x,
		Y: y,
		Z: z}
}
