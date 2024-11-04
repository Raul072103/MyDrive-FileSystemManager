package repository

import (
	"os"
	"testing"
)

var FR = NewFileRepo()

func TestMain(m *testing.M) {
	// Setup code here (runs before all tests)
	setup()

	// Run the tests
	code := m.Run()

	// Teardown code here (runs after all tests)
	teardown()

	// Exit with the code from m.Run()
	os.Exit(code)
}

const testDirPath = "C:\\Users\\raula\\Desktop\\facultate\\anul 3 sem 1\\Software Engineering\\Golang\\Learning\\MyDrive-FileSystemManager\\testfiles"

func setup() {
	createTestsSetup()
}

func teardown() {
	createTestsTeardown()
}

const createTestsDirectory = testDirPath + "\\create_operations_tests"

func createTestsSetup() {
	_, err := os.Stat(createTestsDirectory)
	if os.IsNotExist(err) {
		err := os.Mkdir(createTestsDirectory, 0777)

		if err != nil {
			panic("SETUP FOR CREATE TESTS FAILED: \n" + err.Error())
		}
	}

	// Setup for TestFileRepo_CreateFile_AlreadyExistingFile
	createExistingFileTest := createTestsDirectory + "\\create_test2.txt"
	_, err = os.Stat(createExistingFileTest)
	if os.IsNotExist(err) {
		createTestFileAndWriteContent(createExistingFileTest, []byte("1 2 3"))
	}
}

func createTestsTeardown() {
	print("Teardown is running")
	// Setup for TestFileRepo_CreateFile_TestingDirectory
	createCorrectFileTest := createTestsDirectory + "\\create_test1.txt"
	deleteTestFile(createCorrectFileTest)
}

func createTestFileAndWriteContent(path string, content []byte) {
	file, err := os.Create(path)
	if err != nil {
		panic("SETUP OF THE TESTS FAILED: \n" + err.Error())
	}

	if content != nil {
		_, err = file.Write(content)
		if err != nil {
			panic("SETUP OF THE TESTS FAILED: \n" + err.Error())
		}
	}

	defer func() {
		err := file.Close()
		if err != nil {
			panic("SETUP OF THE TESTS FAILED: \n" + err.Error())
		}
	}()
}

func deleteTestFile(path string) {
	fileExists, err := fileExists(path)
	if err != nil {
		panic("SETUP OF THE TESTS FAILED: \n" + err.Error())
	}

	if fileExists {
		err := os.Remove(path)
		if err != nil {
			panic("SETUP OF THE TESTS FAILED: \n" + err.Error())
		}
	}
}
