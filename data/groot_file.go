package data

import (
	"database/sql"
	"log"
)

type GrootFile struct {
	Id       int
	FileKey  sql.NullString
	SwiftUrl sql.NullString
}

func GrootFiles() (grootFiles []GrootFile, err error) {
	rows, err := Db.Query("select id,file_key,swift_url from groot_files")

	if err != nil {
		log.Fatal(err)
	}
	for rows.Next() {
		grootFile := GrootFile{}

		if err = rows.Scan(&grootFile.Id, &grootFile.FileKey, &grootFile.SwiftUrl); err != nil {
			log.Fatal(err)
		}
		grootFiles = append(grootFiles, grootFile)
	}
	rows.Close()
	return
}

func GrootFilesWithoutSwift() (grootFiles []GrootFile) {
	files, _ := GrootFiles()

	var result []GrootFile

	for _, v := range files {
		if v.SwiftUrl.String == "" {
			result = append(result, v)
		}
	}
	return result

}
