package occupation

import (
	"log"
	"sync"
	"time"
)

type Fireman struct {
}

func (t *Fireman) Name() string {
	return "fireman"
}

func (t *Fireman) Life() {
	t.Born()
	t.Live()
	t.Fight()
	t.Dead()
}

func (t *Fireman) Born() {
	log.Println(t.Name(), "born")
}

func (t *Fireman) Live() {
	log.Println(t.Name(), "living")
	mute := &sync.Mutex{}
	mute.Lock()
	go func() {
		time.Sleep(time.Second)
		mute.Unlock()
	}()
	mute.Lock()
}

func (t *Fireman) Fight() {
	log.Println(t.Name(), "fire fighting")

}

func (t *Fireman) Dead() {
	log.Println(t.Name(), "dead")
}
