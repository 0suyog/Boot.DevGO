package main

func getLast[T any](s []T) T {
	if l:=len(s);l>0{
		return s[l-1]
	}
	var zeroValue T
	return zeroValue
}
