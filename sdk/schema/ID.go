package schema


type ID struct {
	Value 		string 		`json:"value"`
}

func NewID(value string) *ID {
	return &ID{
		Value: value}
}


