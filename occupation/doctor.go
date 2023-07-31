package occupation

import (
	"log"
	"time"
)

type Doctor struct {
}

func (d *Doctor) Name() string {
	return "doctor"
}

func (d *Doctor) Life() {
	d.Born()
	d.Live()
	d.Save()
	d.Dead()
}

func (d *Doctor) Born() {
	log.Println(d.Name(), "born")
}

func (d *Doctor) Live() {
	log.Println(d.Name(), "living")
}

func (d *Doctor) Save() {
	log.Println(d.Name(), "save life")
}

func (d *Doctor) Dead() {
	log.Println(d.Name(), "dead")
	for i := 0; i < 10; i++ {
		go func() {
			time.Sleep(time.Second * 30)
		}()
	}
}
