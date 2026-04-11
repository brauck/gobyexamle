/* package main

import (
	"os"
	"path/filepath"
)

func main() {

	panic("a problem")

	path := filepath.Join(os.TempDir(), "file")
	_, err := os.Create(path)
	if err != nil {
		panic(err)
	}
} */

//////////// DEFER /////////////

/* package main

import (
	"fmt"
	"os"
	"path/filepath"
)

func main() {

	//path := filepath.Join(os.TempDir(), "defer.txt")
	path := filepath.Join("/home/serge/go_work_with_files/", "defer.txt")
	//path := filepath.Join("./", "defer.txt")
	f := createFile(path)
	defer closeFile(f)
	writeFile(f)
}

func createFile(p string) *os.File {
	fmt.Println("creating")
	f, err := os.Create(p)
	if err != nil {
		panic(err)
	}
	return f
}

func writeFile(f *os.File) {
	fmt.Println("writing")
	fmt.Fprintln(f, "data")
}

func closeFile(f *os.File) {
	fmt.Println("closing")
	err := f.Close()

	if err != nil {
		panic(err)
	}
} */

//////////// RECOVER /////////////

package main

import "fmt"

func mayPanic() {
	panic("a problem")
}

func main() {

	defer func() {
		if r := recover(); r != nil {

			fmt.Println("Recovered. Error:\n", r)
		}
	}()

	mayPanic()

	fmt.Println("After mayPanic()")
}
