package main

import (
	"log"
	"net/http"
	_ "net/http/pprof"
	"os"
	"pprof_protice/occupation"
	"runtime"
	"time"
)

type Human interface {
	Life()
	Born()
	Live()
	Dead()
}

func main() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	log.SetOutput(os.Stdout)

	runtime.GOMAXPROCS(1)
	runtime.SetBlockProfileRate(1)
	runtime.SetMutexProfileFraction(1)

	persons := []Human{
		&occupation.People{},
		&occupation.BadPerson{},
		&occupation.Doctor{},
		&occupation.Thief{},
		&occupation.Fireman{},
	}

	go func() {
		if err := http.ListenAndServe(":6060", nil); err != nil {
			log.Fatal(err)
		}
		os.Exit(0)
	}()

	for {
		for _, v := range persons {
			time.Sleep(time.Second)
			v.Life()
		}
	}
}
