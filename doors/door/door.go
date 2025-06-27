package door

type Door struct {
	State string
}

func Toggle(d *Door) {
	if (*d).State == "CLOSED" {
		(*d).State = "OPEN"
	} else {
		(*d).State = "CLOSED"
	}
}

func GetNewDoor() (d Door) {
	return Door{State: "CLOSED"}
}
