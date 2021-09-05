package main

import "fmt"

func main() {
	ub := &UserBuilder{}
	user := ub.
		Name("Mahesh Wijekoon").
		Role("architect").
		Build()
	fmt.Println(user)
}

type User struct {
	Name      string
	Role      string
	MinSalary int
	MaxSalary int
}

type UserBuilder struct {
	User
}

func (ub *UserBuilder) Build() User {
	return ub.User
}

func (ub *UserBuilder) Name(name string) *UserBuilder {
	ub.User.Name = name
	return ub
}

func (ub *UserBuilder) Role(role string) *UserBuilder {
	// verify the role is valid
	if role == "architect" {
		ub.User.MinSalary = 9000000
		ub.User.MaxSalary = 16000000
	}
	ub.User.Role = role
	return ub
}
