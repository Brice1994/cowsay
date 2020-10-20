package main

import (
	"fmt"
	"io/ioutil"
	"strings"
	"math/rand"
	"unicode/utf8"
	"time"
	"math"
)
func check(e error){
	if e != nil {
		panic(e)
	}
}

func main() {
	rand.Seed(time.Now().UnixNano())
	file, err := ioutil.ReadFile("./fortune.txt")

	check(err)
	fortunes := strings.Split(string(file), "%")
	randFortune := fortunes[rand.Intn(len(fortunes) -1)]
	var stripped []string

	fortune := strings.Split(randFortune, "\n")
	for i := 0; i < len(fortune); i++ {
		newstr := strings.Replace(fortune[i], "\r", "", -1)
		if len(newstr) > 0 {
			stripped = append(stripped, strings.Replace(fortune[i], "\r", "", -1))
		}
	}


	cow:= `        \  ^__^
	 \ (oo)\_______
  	   (__)\       )\/\
	      ||----w |
	      ||     ||
  `
  	stripped = tabsToSpaces(stripped)
	maxwidth := calculateMaxWidth(stripped)
	messages := normalizeStringsLength(stripped, maxwidth)
	balloon := createBalloon(messages, maxwidth)

	fmt.Println(balloon)
	fmt.Println(cow)
	fmt.Println()

}

func createBalloon(lines []string, maxwidth int) string {
	var borders []string = []string{"/", "\\", "\\", "/", "|", "<", ">"}
	count:= len(lines)

	var ret []string

	top := " " + strings.Repeat("_", maxwidth + 2)
	bottom := " " + strings.Repeat("_", maxwidth + 2)

	ret = append(ret, top)
	if count == 1 {
		s := fmt.Sprintf("%s %s %s", borders[5], prettySentence(lines[0]), borders[6])
		ret = append(ret, s)
	} else {
		s := fmt.Sprintf(`%s %s %s`, borders[0], prettySentence(lines[0]), borders[1])
		ret = append(ret, s)
		i := 1
		for ; i < count-1; i++ {
			s = fmt.Sprintf(`%s %s %s`, borders[4], prettySentence(lines[i]), borders[4])
			ret = append(ret, s)
		}
		s = fmt.Sprintf(`%s %s %s`, borders[2], prettySentence(lines[i]), borders[3])
		ret = append(ret, s)
	}
	ret = append(ret, bottom)
	return strings.Join(ret, "\n")

}

func prettySentence(s string) string {
	return s
	// var newstr []string
	// split := strings.Split(s, " ")
	// for i := 0; i < len(split); i++ {
	// 	r,g,b := rgb(i)
	// 	pretty := fmt.Sprintf("\033[38;2;%d;%d;%dm%s\033[0m", r, g, b, split[i])
	// 	newstr = append(newstr, pretty)
	// }
	// ret := strings.Join(newstr, " ");
	// return ret
}
func rgb(i int) (int, int, int) {
    var f = 0.1
    return int(math.Sin(f*float64(i)+0)*127 + 128),
        int(math.Sin(f*float64(i)+2*math.Pi/3)*127 + 128),
        int(math.Sin(f*float64(i)+4*math.Pi/3)*127 + 128)
}

func calculateMaxWidth(lines []string) int {
	w := 0
	for _, l := range lines {
		len := utf8.RuneCountInString(l)
		if len > w {
			w = len
		}
	}
	return w
}

func normalizeStringsLength(lines []string, maxwidth int) []string {
	var ret []string
	for _, l := range lines {
		s := l + strings.Repeat(" ", maxwidth-utf8.RuneCountInString(l))
		ret = append(ret, s)
	}
	return ret
}
func tabsToSpaces(lines []string) []string {
	var ret []string
	for _, l := range lines {
		l = strings.Replace(l, "\t", "    ", -1)
		ret = append(ret, l)
	}
	return ret
}