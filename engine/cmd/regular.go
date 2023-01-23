package main

import (
	"fmt"
	"regexp"
)

func main() {

	re := regexp.MustCompile("(?P<first_char>.)(?P<middle_part>.*)(?P<last_char>.)")
	n1 := re.SubexpNames()
	r2 := re.FindAllStringSubmatch("Super", -1)[0]

	md := map[string]string{}
	for i, n := range r2 {
		fmt.Printf("%d. match='%s'\tname='%s'\n", i, n, n1[i])
		md[n1[i]] = n
	}
	fmt.Printf("The names are  : %v\n", n1)
	fmt.Printf("The matches are: %v\n", r2)
	fmt.Printf("The first character is %s\n", md["first_char"])
	fmt.Printf("The last  character is %s\n", md["last_char"])

	re = regexp.MustCompile(`(?m)(a)(\1(?1)?)`)
	var str = `aaaaa`

	for i, match := range re.FindAllString(str, -1) {
		fmt.Println(match, "found at index", i)
	}
}
