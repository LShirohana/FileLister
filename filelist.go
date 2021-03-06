package main

import "fmt"
import "path/filepath"
import "os"
import "io/ioutil"
import "log"
import "strings"

func main() {
	files, err := ioutil.ReadDir(".") // Get file list relative to .exe ran
	if err != nil {
		log.Fatal(err)
	}
	
	f, err := os.OpenFile("list.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644) // Open the file if it exists and append, otherwise create.
	defer f.Close()
	
	for _, file := range files { // for-loop the list via range
		basename := file.Name()
		name := strings.TrimSuffix(basename, filepath.Ext(basename)) // Get the file names without extension.
		
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(name)
		
		s := []string{name, "\n"}
		namedata := strings.Join(s, " ")
		f.Write([]byte(namedata))
	}
}

// I actually got paid $20 to write this. Sweet.
