package main

import (
	"log"
	"strings"
	"sync"
	"vampire/data"
	"vampire/util"
)

func analysisAvatarFiles(wg sync.WaitGroup) {
	log.Println("start analysis avatar file")
	avatarFiles := data.AvatarFilesWithoutSwift()
	str := ""

	var round = len(avatarFiles) / 5000
	strChan := make(chan string, 10)
	wg.Add(1)

	for i := 0; i < round; i++ {
		end := (i + 1) * 5000
		if len(avatarFiles) <= end {
			end = len(avatarFiles)
		}

		part := avatarFiles[i*5000 : end]
		go filePatten(part, str, strChan)
		log.Println("running count", i)
	}
	complexChan(strChan, str, wg)

}

func complexChan(strChan <-chan string, str string, group sync.WaitGroup) {
	log.Println("go file patten running: ", len(strChan))
	defer group.Done()

	for {
		partStr, ok := <-strChan
		if !ok {
			util.WriteFile(str, "avatar_files")
			return
		}
		str = str + partStr
	}
}

func filePatten(avatarFiles []data.AvatarFile, str string, strChan chan<- string) {
	for _, avatarFile := range avatarFiles {
		if avatarFile.FileKey.String == "" {
			continue
		}
		str = generalStr(str, avatarFile.FileKey.String, "ostenement")
	}
	strChan <- str
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
		if v.AvatarKey.String == "" {
			continue
		}
		str = generalStr(str, v.AvatarKey.String, "osportrait")
	}
	util.WriteFile(str, "retailers")
}

func analysisRates() {
	retailers := data.RatesWithoutSwift()
	str := ""
	for _, v := range retailers {
		if v.FileKeys.String == "" {
			continue
		}
		fileKeys := strings.Split(v.FileKeys.String, ",")

		for _, fileKey := range fileKeys {
			str = generalStr(str, fileKey, "oscomment")
		}
	}
	util.WriteFile(str, "retailers")
}

func analysisGroot() {
	log.Println("start analysis groot file")
	grootFiles := data.GrootFilesWithoutSwift()
	str := ""
	for _, v := range grootFiles {
		if v.FileKey.String == "" {
			continue
		}
		str = generalStr(str, v.FileKey.String, "ostenement")
	}
	util.WriteFile(str, "groot_files")
}

func generalStr(str string, fileKey string, account string) string {
	str = str + "\"" + fileKey + "\""
	antmanFile := data.AntmanFileByFileKey(fileKey)
	if antmanFile.FileKey.String == "" {
		str = str + ",\"\",\"" + account + "\"" + "\n"
	} else {
		str = str + ",\"" + antmanFile.Bucket.String + "\",\"" + account + "\"" + "\n"
	}
	return str
}
