package occupation

import (
	"log"
	"pprof_protice/e"
	"time"
)

type BadPerson struct {
	buffer [][e.Mi]byte
}

func (b *BadPerson) Name() string {
	return "bad person"
}

func (b *BadPerson) Life() {
	b.Born()
	b.Live()
	b.Self()
	b.Dead()
}

func (b *BadPerson) Born() {
	log.Println(b.Name(), "born")
}

func (b *BadPerson) Live() {
	log.Println(b.Name(), "living")
	go func() {
		time.Sleep(time.Second * 30)
		max := e.Gi
		for len(b.buffer) < max {
			b.buffer = append(b.buffer, [e.Mi]byte{})
			time.Sleep(time.Millisecond * 500)
		}
	}()
}

func (b *BadPerson) Self() {
	log.Println(b.Name(), "self")
}

func (b *BadPerson) Dead() {
	log.Println(b.Name(), "dead") //WOHHHHHH
	_ = make([]byte, 24*e.Mi)
}
