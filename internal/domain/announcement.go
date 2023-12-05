package domain

type Announcement struct {
	Id          int
	Address     string
	Target      string
	PhoneNumber string
}

func GetAllAnnouncements() []Announcement {
	res := []Announcement{{Id: 1, Address: "Address", Target: "Target", PhoneNumber: "998941231212"},
		{Id: 2, Address: "Address", Target: "Target", PhoneNumber: "998973213434"}}

	return res
}
