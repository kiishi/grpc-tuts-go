package main

import "github.com/deolu-asenuga/learn-grpc/server"

func main() {
	// psom := new(person.Person)
	// psom.Name = "Adeolu"
	// psom.Age = 12
	// value, err := proto.Marshal(psom)
	// if err != nil {
	// 	log.Fatalf("an error occurred : %v", err)
	// }
	// fmt.Print(len(value))
	// var smPerson person.Person
	// err = proto.Unmarshal(value, &smPerson)
	// if err != nil {
	// 	log.Fatalf("An error occured : %v", err)
	// }

	// fmt.Println(smPerson.GetName())
	// fmt.Println(smPerson.GetAge())

	// // for persons
	// psons := new(person.Persons)

	// psons.Persons = []*person.Person{psom}
	// b, err := proto.Marshal(psons)
	// if err != nil {
	// 	fmt.Printf(" An error occured : %v", err)
	// }
	// fmt.Print(len(value))
	// var persons person.Persons
	// err = proto.Unmarshal(b, &persons)
	// if err != nil {
	// 	fmt.Println(" un marshal error : %v ", err)
	// }
	// fmt.Println(persons.GetPersons())
	server.Start()
}
