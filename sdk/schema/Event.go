package schema


type Event struct {
	Payload []byte
	Meta    EventMeta
}

func NewEvent(payload []byte, meta EventMeta) *Event {
	return &Event{
		Payload: payload,
		Meta: meta}
}


