package main

import (
	"fmt"
	"os"
	"strings"
)

const version string = "0.1.0"

func convertToIntArray(pnum string, f *[11]int32) {
	for i, c := range pnum[:11] {
		f[i] = c - 48
	}
}

func formatIsValid(pnum string) bool {
	if len(pnum) != 11 {
		return false
	}
	for i := 0; i < 6; i++ {
		switch i {
		case 0: // Make sure day is between 1-31
			if pnum[i] < 48 || pnum[i] > 51 {
				return false
			} else if pnum[i] == 51 && pnum[i+1] > 49 {
				return false
			} else if pnum[i] == 48 && pnum[i+1] == 48 {
				return false
			}
		case 2: // Make sure month is between 1-12
			if pnum[i] != 48 && pnum[i] != 49 {
				return false
			} else if pnum[i] == 49 && pnum[i+1] > 50 {
				return false
			} else if pnum[i] == 48 && pnum[i+1] == 48 {
				return false
			}
		}
	}
	for i := 6; i < 11; i++ {
		if (pnum[i] < 48 || pnum[i] > 57) && pnum[i] != 63 {
			return false
		}
	}
	return true
}

func calculateCtrlNumber(pnum [11]int32) (int32, int32) {
	ctrl1 := 11 - ((3*int(pnum[0]) + 7*int(pnum[1]) + 6*int(pnum[2]) + 1*int(pnum[3]) +
		8*int(pnum[4]) + 9*int(pnum[5]) + 4*int(pnum[6]) + 5*int(pnum[7]) +
		2*int(pnum[8])) % 11)
	ctrl2 := 11 - ((5*int(pnum[0]) + 4*int(pnum[1]) + 3*int(pnum[2]) + 2*int(pnum[3]) +
		7*int(pnum[4]) + 6*int(pnum[5]) + 5*int(pnum[6]) + 4*int(pnum[7]) +
		3*int(pnum[8]) + 2*ctrl1) % 11)
	if ctrl1 == 10 || ctrl2 == 10 {
		return -1, -1
	}
	if ctrl1 == 11 {
		ctrl1 = 0
	}
	if ctrl2 == 11 {
		ctrl2 = 0
	}
	return int32(ctrl1), int32(ctrl2)
}
func generateNumbers(pnum [11]int32, depth int, idx int) {
	if depth < 3 {
		if pnum[idx] == 15 {
			for i := 0; i < 10; i++ {
				pnum[idx] = int32(i)
				generateNumbers(pnum, depth+1, idx+1)
			}
		} else {
			generateNumbers(pnum, depth+1, idx+1)
		}
	} else {
		pnum[9], pnum[10] = calculateCtrlNumber(pnum)
		if pnum[9] != -1 && pnum[10] != -1 {
			p := strings.Trim(strings.Join(strings.Fields(fmt.Sprint(pnum)), ""), "[]")
			fmt.Println(p)
		}
	}
}

func usage() {
	fmt.Printf("personummer v%s\nUsage: %s <number>\nUse '?' as wildcards\nExample: %s 010170?5???\n",
		version, os.Args[0], os.Args[0])
}
func main() {
	var fodselsnummer [11]int32
	switch len(os.Args) {
	case 2:
		if formatIsValid(os.Args[1]) == true {
			convertToIntArray(os.Args[1], &fodselsnummer)
			if fodselsnummer[9] != 15 && fodselsnummer[10] != 15 {
				var c1, c2 int32
				c1, c2 = calculateCtrlNumber(fodselsnummer)
				if c1 == fodselsnummer[9] && c2 == fodselsnummer[10] {
					fmt.Println("Valid")
				} else {
					fmt.Println("Invalid")
				}
			} else {
				generateNumbers(fodselsnummer, 0, 6)
			}
		} else {
			usage()
		}
	default:
		usage()
	}
	os.Exit(0)
}
