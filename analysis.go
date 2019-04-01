package main

import (
	"os"
	"log"
	"vampire/data"
)

func analysisAvatar()  {
	avatarFiles, _ := data.AvatarFiles()

	str := ""

	for _, avatarFile := range avatarFiles {

		if avatarFile.FileKey == "" {
			continue
		}

		str = str + "\"" + avatarFile.FileKey + "\""

		antmanFile := data.AntmanFileByFileKey(avatarFile.FileKey)
		if antmanFile.FileKey == "" {
			str = str + ",\"\",\"ostenement\"" + "\n"
		} else {
			str = str + ",\"" + antmanFile.Bucket + "\",\"ostenement\"" + "\n"

		}

	}
	file, err := os.Create("./avatar_files.txt")
	if err != nil {
		log.Fatal(err)
	}
	//写入字符串
	file.WriteString(str);
	file.Close();
}
