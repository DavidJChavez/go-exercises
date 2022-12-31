package main

import "fmt"

type person struct {
	fname string
	lname string
}

type secretAgent struct {
	person
	licenseToKill bool
}

func (p person) pSpeak() {
	fmt.Println(p.fname, p.lname, `says, "Good morning, David."`)
}

func (sA secretAgent) saSpeak() {
	fmt.Println(sA.fname, sA.fname, `says, "Shaken, not stirred."`)
}

//func main() {
//	p1 := person{
//		`Rebeca`,
//		`Rosales`,
//	}
//	sA1 := secretAgent{
//		person{
//			`Damaris`,
//			`Rosales`,
//		},
//		true,
//	}
//	fmt.Println(`p1 first name > `, p1.fname)
//	p1.pSpeak()
//	fmt.Println(`sA first name > `, sA1.fname)
//	sA1.saSpeak()
//	sA1.pSpeak()
//}
