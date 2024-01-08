package main

import "strconv"

func PowInt64(v, p int64) int64 {
	if p == 0 {
		return 1
	} else {
		value := v
		for i := int64(1); i < p; i++ {
			value *= v
		}
		return value
	}
}

func sizeToString(v int64) string {
	units := []string{"B", "KB", "MB", "GB", "TB", "PB", "EB", "ZB"}

	if v < 1024 {
		return strconv.FormatInt(v, 10) + " " + units[0]
	}

	size := int64(len(units))
	for i := int64(1); i < size; i++ {
		if v < PowInt64(1024, i+1) {
			d := PowInt64(1024, i)
			return strconv.FormatInt(v/d, 10) + "," + strconv.Itoa(int(10*(float64(v)/float64(d)-float64(v/d)))) + " " + units[i]
		}
	}

	return ""
}
