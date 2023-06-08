package list

type List interface {
	length() int
	getValue() int
	currPos() int
	moveToStart()
	next()
}

func hasItem(list List, item int) bool {
	var val int

	for list.moveToStart(); list.currPos() < list.length(); list.next() {
		val = list.getValue()
		if val == item {
			return true
		}
	}

	return false
}
