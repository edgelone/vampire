package util

import (
	"os"
	"log"
)

func WriteFile(data string, fileName string) {
	file, err := os.Create("./" + fileName + ".txt")
	if err != nil {
		log.Fatal(err)
	}
	//写入字符串
	file.WriteString(data);
	file.Close();
}
