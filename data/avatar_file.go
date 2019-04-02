package data

import "log"

type AvatarFile struct {
	Id         int
	FileKey    string
	FileSource string
	SwiftUrl   string
}

func AvatarFiles() (avatarFiles []AvatarFile, err error) {
	rows, err := Db.Query("select id,file_key,file_source,swift_url from avatar_files")

	if err != nil {
		log.Fatal(err)
	}
	for rows.Next() {
		avatarFile := AvatarFile{}

		if err = rows.Scan(&avatarFile.Id, &avatarFile.FileKey, &avatarFile.FileSource, &avatarFile.SwiftUrl); err != nil {
			log.Fatal(err)
		}
		avatarFiles = append(avatarFiles, avatarFile)
	}
	rows.Close()
	return
}



func AvatarFilesWithoutSwift() (avatarFiles []AvatarFile) {
	files, _ := AvatarFiles()

	result := []AvatarFile{}

	for _, v := range files {
		if v.SwiftUrl == "" {
			result = append(result, v)
		}
	}
	return result

}