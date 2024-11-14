package main

import (
	"MyDrive-FileSystemManager/internal/repository/filerepo"
	"fmt"
)

const testPath = "filerepo/repository/data.txt"

func main() {
	fileRepoMain()
}

func fileRepoMain() {
	var fmRepo = filerepo.FileRepo{}

	content, err := fmRepo.ReadFile(testPath)

	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(string(content))

	err = fmRepo.UpdateFile(testPath, []byte("allooo"), 0)

	content, err = fmRepo.ReadFile(testPath)

	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(string(content))

	err = fmRepo.DeleteFile(testPath)
	if err != nil {
		fmt.Println(err)
		return
	}

	err = fmRepo.CreateFile(testPath)
	if err != nil {
		fmt.Println(err)
		return
	}

	err = fmRepo.UpdateFile(testPath, []byte("ana are mere\nraul are mere\n"), 0)

}
