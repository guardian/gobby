package gobby

import (
	"bytes"
	"encoding/gob"
	"io/ioutil"
	"log"
	"runtime"
	"sync"
)

type JobStatus struct {
	Id      string
	Status  string
	Payload interface{}
}

type Gobby struct {
	Location string
	Jobs     map[string]JobStatus
	mutex    *sync.Mutex
}

func New(location string) *Gobby {
	return &Gobby{
		location,
		make(map[string]JobStatus),
		&sync.Mutex{},
	}
}

func (gobby *Gobby) Set(s string, j JobStatus) {
	gobby.mutex.Lock()
	gobby.Jobs[s] = j
	gobby.mutex.Unlock()

	runtime.Gosched()
}

func (gobby *Gobby) Get(s string) (JobStatus, bool) {
	gobby.mutex.Lock()
	j, exists := gobby.Jobs[s]
	gobby.mutex.Unlock()

	runtime.Gosched()

	return j, exists
}

func (gobby *Gobby) Load() error {
	var buf bytes.Buffer

	// --- Begin lock
	gobby.mutex.Lock()

	dat, err := ioutil.ReadFile(gobby.Location)

	if err != nil {
		log.Println("gobby file read error:", err)
		return err
	}

	_, err = buf.Write(dat)

	dec := gob.NewDecoder(&buf)

	dec.Decode(&gobby.Jobs)

	gobby.mutex.Unlock()
	// --- End lock

	runtime.Gosched()

	return err
}

func (gobby *Gobby) Save() error {
	var buf bytes.Buffer

	enc := gob.NewEncoder(&buf)

	// --- Begin lock
	gobby.mutex.Lock()

	err := enc.Encode(gobby.Jobs)

	if err != nil {
		log.Println("gobby encode error:", err)
		return err
	}

	err = ioutil.WriteFile(gobby.Location, buf.Bytes(), 0644)

	gobby.mutex.Unlock()
	// --- End lock

	runtime.Gosched()

	if err != nil {
		log.Println("gobby save error:", err)
		return err
	}

	return err
}
