package repository

import (
	"errors"
	"os"
	"testing"
)

type testFile struct {
	path      string
	operation string
	args      []any
}

const createTestsDirectoryPath = "C:\\Users\\raula\\Desktop\\facultate\\anul 3 sem 1\\Software Engineering\\Golang\\Learning\\MyDrive-FileSystemManager\\testfiles\\create_operations_tests"

// TestFileRepo_CreateFile_TestingDirectory tests if the file is successfully created in that directory
func TestFileRepo_CreateFile_TestingDirectory(t *testing.T) {
	path := createTestsDirectoryPath + "\\create_test1.txt"

	err := FR.CreateFile(path)
	if err != nil {
		t.Errorf("File wasn't created successfully!")
	}

	err = testIfFileExists(path)
	if err != nil {
		t.Errorf("File doesn't exist after it was created!")
	}
}

// TestFileRepo_CreateFile_PathInWrongFormat tests if an error is thrown when the path doesn't exist
func TestFileRepo_CreateFile_PathInWrongFormat(t *testing.T) {
	paths := []string{"\\\\invalid_path", "invalid_path////\\", "\\\nilfasf@@...124124.txtxt"}

	for _, path := range paths {
		err := FR.CreateFile(path)
		if err == nil {
			t.Errorf("File %s was created even though it shouldn't have!", path)
		}
	}
}

// TestFileRepo_CreateFile_AlreadyExistingFile tests if an error is thrown when the file already exist
func TestFileRepo_CreateFile_AlreadyExistingFile(t *testing.T) {
	path := createTestsDirectoryPath + "\\create_test2.txt"

	err := FR.CreateFile(path)
	if err == nil {
		t.Errorf("File was created when it shouldn't have been!")
	}
}

func TestFileRepo_ReadFile_ReadsCorrectly(t *testing.T) {

}

func TestFileRepo_ReadFile_FileDoesntExist(t *testing.T) {

}

func TestFileRepo_ReadFile_NoPermissionsToAccessTheDirectory(t *testing.T) {

}

func TestFileRepo_ReadFile_NoPermissionsToRead(t *testing.T) {

}

func TestFileRepo_DeleteFile_DeletesCorrectly(t *testing.T) {

}

func TestFileRepo_DeleteFile_FileDoesntExist(t *testing.T) {

}

func TestFileRepo_DeleteFile_NoPermissionsToAccessTheDirectory(t *testing.T) {

}

func TestFileRepo_UpdateFile_UpdatesCorrectly(t *testing.T) {

}

func TestFileRepo_UpdateFile_OffsetBiggerThanFileSize(t *testing.T) {

}

func TestFileRepo_UpdateFile_OffsetIsNegative(t *testing.T) {

}

func TestFileRepo_UpdateFile_FileDoesntExist(t *testing.T) {

}

func TestFileRepo_UpdateFile_NoPermissionsToWrite(t *testing.T) {

}

func testIfFileExists(path string) (err error) {
	file, err := os.Open(path)
	if err != nil {
		return err
	}
	defer func() {
		err = errors.Join(err, file.Close())
	}()
	return err
}
