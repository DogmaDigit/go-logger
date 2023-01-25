package go_logger

import (
	"io/ioutil"
	"os"
)

// Проверка на существование папки
func checkFolder(argPath string) error {
	_, error := ioutil.ReadDir(argPath)
	if error != nil {
		return error
	}
	return nil
}

// Создание вложенных папок
func createFolder(argPathFolder string) error {
	err := os.MkdirAll(argPathFolder, 0o777)
	if err != nil {
		return err
	}
	return nil
}
