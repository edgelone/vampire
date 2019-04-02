package main

import (
	"vampire/data"
	"vampire/util"
)

func analysisAvatarFiles() {
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

func analysisContracts(){
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

func analysisRetailers(){

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
