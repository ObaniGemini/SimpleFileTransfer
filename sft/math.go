package main

import "strconv"

func PowUInt64(v, p uint64) uint64 {
	if p == 0 {
		return 1
	} else {
		value := v
		for i := uint64(1); i < p; i++ {
			value *= v
		}
		return value
	}
}

func sizeToString(v uint64) string {
	units := []string{"B", "KB", "MB", "GB", "TB", "PB", "EB", "ZB"}

	if v < 1024 {
		return strconv.FormatUint(v, 10) + " " + units[0]
	}

	size := uint64(len(units))
	for i := uint64(1); i < size; i++ {
		if v < PowUInt64(1024, i+1) {
			d := PowUInt64(1024, i)
			return strconv.FormatUint(v/d, 10) + "," + strconv.Itoa(int(10*(float64(v)/float64(d)-float64(v/d)))) + " " + units[i]
		}
	}

	return ""
}
