package main

type User struct {
	id    int
	fname string
	lname string
}

func mainStatement() {
	u1 := User{
		id:    1,
		fname: "daddy",
		lname: "pop",
	}
	u2 := User{
		id:    2,
		fname: "paddy",
		lname: "popy",
	}

	if u1 != u2 {
		println("same user")
	} else if u1.fname == u2.fname {
		println("similar user.")
	} else {
		println("different user!")
	}

}
