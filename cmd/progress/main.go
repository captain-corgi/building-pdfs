package main

import (
	"fmt"
	"strings"
)

var (
	playlist = []string{
		"",
		"https://courses.calhoun.io/lessons/les_goph_127",
		"https://courses.calhoun.io/lessons/les_goph_128",
		"https://courses.calhoun.io/lessons/les_goph_129",
		"https://courses.calhoun.io/lessons/les_goph_130",
	}
	inprogress = "https://courses.calhoun.io/lessons/les_goph_126"
	finished   = []string{
		"",
	}
)

func main() {
	fmt.Printf("WATCHING:\n \t%s\n", inprogress)
	fmt.Printf("PLAYLIST:\t %+v\n", strings.Join(playlist, "\n\t"))
	fmt.Printf("FINISHED:\t %+v\n", strings.Join(finished, "\n\t"))
}
