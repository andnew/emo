package employee

import "fmt"

type Employee struct {
	FirstName   string
	LastName    string
	TotalLeaves int
	LeavesTaken int
}

func New(firstName string, lastName string, totalLeave int, leavesTaken int) Employee {
	e := Employee{FirstName: firstName, LastName: lastName, TotalLeaves: totalLeave, LeavesTaken: leavesTaken}
	return e
}

func (e Employee) LeavesRemaining() {
	fmt.Printf("%s %s has %d leaves remaining\n", e.FirstName, e.LastName, e.TotalLeaves-e.LeavesTaken)
}
