package main

import (
	"log"
	"strings"
	"vampire/data"
	"vampire/util"
)

func analysisAvatarFiles() {
	log.Println("start analysis avatar file")
	avatarFiles := data.AvatarFilesWithoutSwift()
	str := ""

	for _, avatarFile := range avatarFiles {
		if avatarFile.FileKey == "" {
			continue
		}
		str = generalStr(str, avatarFile.FileKey, "ostenement")
	}
	util.WriteFile(str, "avatar_files")
}

func analysisContracts() {
	contracts := data.ContractsWithoutSwift()
	str := ""

	for _, contract := range contracts {
		if contract.FileKey == "" {
			continue
		}
		str = generalStr(str, contract.FileKey, "osfileprivate")
	}
	util.WriteFile(str, "contracts")
}

func analysisRetailers() {
	retailers := data.RetailersWithoutSwift()
	str := ""
	for _, v := range retailers {
		if v.AvatarKey == "" {
			continue
		}
		str = generalStr(str, v.AvatarKey, "osportrait")
	}
	util.WriteFile(str, "retailers")
}

func analysisRates() {
	retailers := data.RatesWithoutSwift()
	str := ""
	for _, v := range retailers {
		if v.FileKeys == "" {
			continue
		}
		fileKeys := strings.Split(v.FileKeys, ",")

		for _, fileKey := range fileKeys {
			str = generalStr(str, fileKey, "oscomment")
		}
	}
	util.WriteFile(str, "retailers")
}

func analysisGroot()  {
	log.Println("start analysis groot file")
	grootFiles :=data.GrootFilesWithoutSwift()
	str :=""
	for _,v := range grootFiles{
		if v.FileKey==""{
			continue
		}
		str = generalStr(str,v.FileKey,"ostenement")
	}
	util.WriteFile(str,"groot_files")
}

func generalStr(str string, fileKey string, account string) string {
	str = str + "\"" + fileKey + "\""
	antmanFile := data.AntmanFileByFileKey(fileKey)
	if antmanFile.FileKey == "" {
		str = str + ",\"\",\"" + account + "\"" + "\n"
	} else {
		str = str + ",\"" + antmanFile.Bucket + "\",\"" + account + "\"" + "\n"
	}
	return str
}
