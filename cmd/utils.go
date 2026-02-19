package cmd

import (
	"fmt"
	"os"
)

func errorAndClose(originalError error, path string) {
	fmt.Println(originalError)
	if path == "-" {
		os.Exit(1)
	}

	// Output error to file
	f, err := os.Create(fmt.Sprintf("%s.json", path))
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	defer f.Close()

	_, err = f.WriteString(fmt.Sprintf("{\"message\": \"%s\"}", originalError.Error()))
	if err != nil {
		return
	}
	os.Exit(1)
}
