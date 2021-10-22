package main //! Script for move files on prefix in their names

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path"
	"runtime"
	"strings"
)

func main() {
	files, err := ioutil.ReadDir(".")
	if err != nil {
		log.Fatal(err)
	}
	removed := 0
	for _, file := range files {
		_, currentFilePath, _, _ := runtime.Caller(0) //! files name
		dirpath := path.Dir(currentFilePath)          //! dirpath = path where files right now
		NewDir := ""                                  //! set nul string to crete new dir

		for index, name := range file.Name() { //! check files name
			if string(name) == "_" && index != 0 && string(file.Name()[index-1]) == "_" { //!check is filename has "__" and first not "_"
				for i := 0; i < index-1; i++ {
					NewDir += string(file.Name()[i])
				}
				_ = os.Mkdir(string(NewDir), 0777) //!create dir with name before "__"

				for _, filename := range files { //! check all file
					if strings.HasPrefix(filename.Name(), string(NewDir)) && !filename.IsDir() { //!check filename to name of NewDir && filename != DIR
						oldLocation := dirpath + "/" + filename.Name()                        //! create old path
						newLocation := dirpath + "/" + string(NewDir) + "/" + filename.Name() //! create new path = old + newDir + filename
						_ = os.Rename(oldLocation, newLocation)                               //! rename path of file
					}
				}

				removed++
				fmt.Printf("remove %v files, now: %v\n", removed, file.Name())
			}
		}
		//fmt.Printf("isDir: %v \t fileName: %s\n", file.IsDir(), file.Name())
	}
}
