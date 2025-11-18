/*
package main

import "fmt"

func hiddenp(s1, s2 string) string {

	if s1 == "" {
		return "1\n"
	}
	j := 0

	for i := 0; i < len(s2); i++ {

		if s2[i] == s1[j] {

			j++
		}
		if j == len(s1) {
			return "1\n"
		}
	}
	return "0\n"
}

func main() {
	fmt.Print(hiddenp("fgex.;", "tyf34gdgf;'ektufjhgdgex.;.;rtjynur6")) // 1
	fmt.Print(hiddenp("abc", "2altrb53c.sse"))                          // 1
	fmt.Print(hiddenp("abc", "btarc"))                                  // 0
	fmt.Print(hiddenp("DD", "DABC"))                                    // 0
}
*/

package main

import "fmt"

func hiddenp(s1, s2 string) string {
	if s1 == "" {
		return "1\n"
	}
	j := 0
	for i := 0; i < len(s2); i++ {
		if s2[i] == s1[j] {
			j++
		}
		if j == len(s1) {
			return "1\n"
		}
	}
	return "0\n"
}

func main() {
	fmt.Print(hiddenp("fgex.;", "tyf34gdgf;'ektufjhgdgex.;.;rtjynur6")) // 1
	fmt.Print(hiddenp("abc", "2altrb53c.sse"))                          // 1
	fmt.Print(hiddenp("abc", "btarc"))                                  // 0
	fmt.Print(hiddenp("DD", "DABC"))                                    // 0
}
