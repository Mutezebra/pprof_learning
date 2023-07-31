package occupation

import (
	"log"
	"time"
)

type Thief struct {
}

type ThiefInterface interface {
}

func (t *Thief) Life() {
	t.Born()
	t.Live()
	t.Steal()
	t.Dead()
}

func (t *Thief) Name() string {
	return "Thief"
}

func (t *Thief) Born() {
	log.Println(t.Name(), "born")
}

func (t *Thief) Live() {
	log.Println(t.Name(), "living")
	<-time.After(time.Second)
}

func (t *Thief) Steal() {
	log.Println(t.Name(), "steal")
	loop := 10000000000
	for i := 0; i < loop; i++ {

	}
}

func (t *Thief) Dead() {
	log.Println(t.Name(), "dead")
}
