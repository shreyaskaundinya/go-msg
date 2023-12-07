package message

type Message struct {
	Topic  string `json:"topic"`
	Key    string `json:"key"`
	Value  string `json:"value"`
	Offset int64  `json:"offset"`
}
