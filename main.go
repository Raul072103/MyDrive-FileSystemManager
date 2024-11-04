package main

const testPath = "filemanager/repository/data.txt"

//func main() {
//	var fmRepo = repository.FileRepo{}
//
//	content, err := fmRepo.ReadFile(testPath)
//
//	if err != nil {
//		fmt.Println(err)
//		return
//	}
//	fmt.Println(string(content))
//
//	err = fmRepo.UpdateFile(testPath, []byte("allooo"), 0)
//
//	content, err = fmRepo.ReadFile(testPath)
//
//	if err != nil {
//		fmt.Println(err)
//		return
//	}
//	fmt.Println(string(content))
//
//	err = fmRepo.DeleteFile(testPath)
//	if err != nil {
//		fmt.Println(err)
//		return
//	}
//
//	err = fmRepo.CreateFile(testPath)
//	if err != nil {
//		fmt.Println(err)
//		return
//	}
//
//	err = fmRepo.UpdateFile(testPath, []byte("ana are mere\nraul are mere\n"), 0)
//
//}
