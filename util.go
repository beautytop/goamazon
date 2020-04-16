package goamazon

import "strconv"

func SInt64(s string) (i int64, e error) {
	ii, e := strconv.Atoi(s)
	if e != nil {
		return 0, e
	}

	return int64(ii), nil
	return
}

func SFloat64(s string) (i float64, e error) {
	i, e = strconv.ParseFloat(s, 64)
	if e != nil {
		return 0, e
	}

	return
}
