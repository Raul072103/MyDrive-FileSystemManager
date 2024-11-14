package main

import "MyDrive-FileSystemManager/internal/repository/filerepo"

func main() {
	fr := filerepo.FileRepo{}

	content, err := fr.ReadFile("./readme.md")
	if err != nil {
		print(err)
		return
	}

	print(string(content))
}
