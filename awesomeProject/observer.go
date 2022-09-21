package main

import "fmt"

type Observer interface {
	handleEvent(vacancies []string)
}

type Observable interface {
	subscribe()
	unsubscribe()
	sendAll()
}

type Person struct {
	name string
}

type JobSite struct {
	vacancies   []string
	subscribers []Person
}

func (p *Person) handleEvent(vacancies []string) {
	fmt.Printf("Good day, %v! There are our available vacancies:\n", p.name)
	for i := 0; i < len(vacancies); i++ {
		fmt.Printf("%v. %v\n", i+1, vacancies[i])
	}
}

func (s *JobSite) subscribe(p Person) {
	s.subscribers = append(s.subscribers, p)
}

func (s *JobSite) unsubscribe(p Person) {
	tempSubs := []Person{}
	for i := 0; i < len(s.subscribers); i++ {
		if s.subscribers[i] != p {
			tempSubs = append(tempSubs, s.subscribers[i])
		}
	}
	s.subscribers = tempSubs
}

func (s *JobSite) addVac(v string) {
	s.vacancies = append(s.vacancies, v)
	s.sendAll()
}

func (s *JobSite) removeVac(vac string) {
	tempVacs := []string{}
	for i := 0; i < len(s.vacancies); i++ {
		if s.vacancies[i] != vac {
			tempVacs = append(tempVacs, s.vacancies[i])
		}
	}
	s.vacancies = tempVacs
	s.sendAll()
}

func (s *JobSite) sendAll() {
	for i := 0; i < len(s.subscribers); i++ {
		s.subscribers[i].handleEvent(s.vacancies)

	}
}

func main() {
	hhkz := JobSite{
		subscribers: []Person{},
		vacancies:   []string{},
	}
	person1 := Person{name: "Aqzer"}
	person2 := Person{name: "Nazgul"}
	person3 := Person{name: "Akmerey"}
	hhkz.subscribe(person1)
	hhkz.subscribe(person2)
	hhkz.subscribe(person3)
	hhkz.unsubscribe(person3)
	hhkz.addVac("Java DEV")
	hhkz.addVac("Math PROF")
	hhkz.addVac("C++ DEV")
	hhkz.removeVac("Java DEV")

}
