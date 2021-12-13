package schema


type AggregateMeta struct {
	AggregateId string `json:"aggregate_id"`
	Status int64 `json:"status"`
	Version int64 `json:"version"`	
}

func NewAggregateMeta(aggregateId string, status int64, version int64) *AggregateMeta {
	return &AggregateMeta{
		AggregateId: aggregateId,
		Status: status,
		Version: version}
}
