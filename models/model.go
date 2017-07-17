package models

import (
	"encoding/json"
	log "log"
)

type Model struct {
}

func (m *Model) Marshal() []byte {
	encoded, err := json.Marshal(m)
	if err != nil {
		log.Println(err)
	}

	return encoded
}

func (m *Model) FromBytes(b []byte) {
	err := json.Unmarshal(b, m)
	if err != nil {
		log.Fatal(err)
	}
}
