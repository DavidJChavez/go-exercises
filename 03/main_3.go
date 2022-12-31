package main

import "fmt"

type human interface {
	sleep(hours int)
}

func (p person) sleep(hours int) {
	fmt.Println(`The person`, p.fname, `slept`, hours, `hours`)
}

func (sA secretAgent) sleep(hours int) {
	fmt.Println(`The secret agent`, sA.fname, `slept`, hours, `hours`)
}

func sendToSleep(h human) {
	h.sleep(3)
}

func main() {
	p1 := person{`Rebeca`, `Rosales`}
	sA1 := secretAgent{person{`Damaris`, `Rosales`}, true}

	sendToSleep(p1)
	sendToSleep(sA1)
}
