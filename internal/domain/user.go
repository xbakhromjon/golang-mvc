package domain

type User struct {
	Id    int
	Fname string
	Lname string
	Age   uint8
}

func GetAllUsers() []User {
	res := []User{{Id: 1, Fname: "fname", Lname: "lname", Age: 10}, {Id: 2, Fname: "fname", Lname: "lname", Age: 15}}
	return res
}
