package exercises

import (
	"fmt"
)

//Получение имен файлов, где есть одинаковые строки
func GetFileNamesWithSameStrings(fileNames []string) {
	fmt.Println("Запуск GetFileNamesWithSameStrings!")

	if len(fileNames) == 0 {
		fmt.Println("Входящий массив пуст!!!")
		return
	}
	for i, filename := range fileNames {
		fmt.Printf("Файл № %d с наименованием %s\n", i, filename)

	}

	fmt.Println("Завершение GetFileNamesWithSameStrings...")
}
