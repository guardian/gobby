package gobby

import (
	"bytes"
	"encoding/gob"
	"io/ioutil"
	"log"
)

type JobStatus struct {
	Id      string
	Status  string
	Payload interface{}
}

type Gobby struct {
	Location string
	Jobs     map[string]JobStatus
}

func New(location string) *Gobby {
	return &Gobby{location, make(map[string]JobStatus)}
}

func (gobby *Gobby) Set(s string, j JobStatus) {
	gobby.Jobs[s] = j
}

func (gobby *Gobby) Get(s string) (*JobStatus, bool) {
	j, exists := gobby.Jobs[s]
	return &j, exists
}

func (gobby *Gobby) Load() error {
	var buf bytes.Buffer

	dat, err := ioutil.ReadFile(gobby.Location)

	if err != nil {
		log.Println("gobby file read error:", err)
		return err
	}

	_, err = buf.Write(dat)

	dec := gob.NewDecoder(&buf)
	dec.Decode(&gobby.Jobs)

	return err
}

func (gobby *Gobby) Save() error {
	var buf bytes.Buffer

	enc := gob.NewEncoder(&buf)
	err := enc.Encode(gobby.Jobs)

	if err != nil {
		log.Println("gobby encode error:", err)
		return err
	}

	err = ioutil.WriteFile(gobby.Location, buf.Bytes(), 0644)

	if err != nil {
		log.Println("gobby save error:", err)
		return err
	}

	return err
}
