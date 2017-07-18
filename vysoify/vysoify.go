package vysoify

import "math"

// FixVysoFakes modifies alphabetical characters in the byte slice to the corresponding vyso characters
func FixVysoFakes(bytes []byte) {
	var letter bool
	var tracker []int
	for k, b := range bytes {
		// modify the beginning and end of part we are searching
		if letter && isLetter(b) {
			makeVyso(tracker, bytes)
			tracker = []int{}
			letter = false
		}
		if isLetter(b) {
			letter = true
			tracker = append(tracker, k)
		}
	}
}

// makeVyso updates bytes with vyso characters, delegator
func makeVyso(tracker []int, bytes []byte) {
	count := 3
	if len(tracker) >= 4 {
		makeVysoLarge(count, tracker, bytes)
	} else {
		makeVysoSmall(count, tracker, bytes)
	}
}

// makeVysoLarge is for balancing larger words, ex. "factess" = "  vyso  ", "factes" = " vyso  "
func makeVysoLarge(count int, tracker []int, bytes []byte) {
	dist := int(math.Abs(float64(len(tracker) - 4)))
	front := dist / 2
	back := dist - front
	for _, v := range tracker {
		if front > 0 {
			bytes[v] = 32
			front--
			continue
		}
		if front == 0 {
			formVyso(count, v, bytes)
		}
		if back > 0 {
			bytes[v] = 32
			back--
		}
	}
}

// makeVysoSmall for balancing smaller words, ex. "cat" = "", "a" = ""
func makeVysoSmall(count int, tracker []int, bytes []byte) {
	for _, v := range tracker {
		switch count {
		case 3:
			bytes[v] = 118
			count--
		case 2:
			bytes[v] = 121
			count--
		case 1:
			bytes[v] = 115
			count--
		case 0:
			bytes[v] = 111
			count = 3
		}
	}
}

// formVyso modifies byte[v] to a vyso character
func formVyso(count int, v int, bytes []byte) {
	switch count {
	case 3:
		bytes[v] = 118
		count--
	case 2:
		bytes[v] = 121
		count--
	case 1:
		bytes[v] = 115
		count--
	case 0:
		bytes[v] = 111
	}

}

// isLetter returns true if byte is a letter else false
func isLetter(in byte) bool {
	return (in >= 97 && in <= 122) || (in >= 65 && in <= 90)
}
