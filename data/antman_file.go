package data

import "log"

type AntmanFile struct {
	Id       int
	AppId    string
	FileKey  string
	Bucket   string
	Size     int
	SwiftUrl string
}

func AntmanFiles() (antmanFiles []AntmanFile, err error) {
	rows, err := Db.Query("select id,app_id,file_key,bucket,size,swift_url from antman_files")

	if err != nil {
		log.Fatal(err)
	}
	for rows.Next() {
		antmanFile := AntmanFile{}

		if err = rows.Scan(&antmanFile.Id, &antmanFile.AppId, &antmanFile.FileKey, &antmanFile.Bucket, &antmanFile.Size,
			&antmanFile.SwiftUrl); err != nil {
			log.Fatal(err)
		}
		antmanFiles = append(antmanFiles, antmanFile)
	}
	rows.Close()
	return
}

func AntmanFilesByFileKey(fileKey string) (antmanFiles []AntmanFile, err error) {
	rows, err := Db.Query("select id,app_id,file_key,bucket,size,swift_url from antman_files where file_key = '" + fileKey+"'")
	if err != nil {
		log.Fatal(err)
	}
	for rows.Next() {
		antmanFile := AntmanFile{}

		if err = rows.Scan(&antmanFile.Id, &antmanFile.AppId, &antmanFile.FileKey, &antmanFile.Bucket, &antmanFile.Size,
			&antmanFile.SwiftUrl); err != nil {
			log.Fatal(err)
		}
		antmanFiles = append(antmanFiles, antmanFile)
	}
	rows.Close()
	return
}

func AntmanFileByFileKey(fileKey string) (antmanFile AntmanFile) {
	antmanFiles, err := AntmanFilesByFileKey(fileKey)
	if err != nil {
		return
	}
	for _, a := range antmanFiles {
		if a.Size > 10 {
			return a
		}
	}
	return
}
