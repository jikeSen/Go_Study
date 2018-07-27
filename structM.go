package main

import "time"

func main() {
	jike_sen.name = "jikesen"
	postion := &jike_sen.postion
	*postion = "Sen " + *postion
}

type User struct {
	Id      int64
	name    string
	address string
	Dob     time.Time
	postion string
	Salary  int
}

var jike_sen User
