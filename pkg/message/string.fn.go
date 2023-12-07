package message

import (
	"encoding/json"
)

func (m Message) Marshall() ([]byte, error) {

	j, err := json.Marshal(m)

	if err != nil {
		return nil, err
	}

	return j, nil
}
