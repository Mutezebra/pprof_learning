package occupation

import (
	"log"
	"time"
)

type People struct {
}

func (p *People) Name() string {
	return "people"
}

func (p *People) Life() {
	p.Born()
	p.Live()
	p.Dead()
}

func (p *People) Born() {
	log.Println(p.Name(), "born")
	<-time.After(time.Second)
}

func (p *People) Live() {
	log.Println(p.Name(), "live")
}

func (p *People) Dead() {
	log.Println(p.Name(), "dead")
}
