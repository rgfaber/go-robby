package schema

type Command struct {
	Topic 		string 				`json:"topic"`
	Meta 		AggregateMeta 		`json:"meta"`
	Payload		[]byte 				`json:"payload"`	
}

func NewCommand(topic string, meta AggregateMeta, payload []byte) *Command {
	return &Command{
		Topic: topic,
		Meta: meta,
		Payload: payload}
}
