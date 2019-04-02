package data

import "log"

type Contract struct {
	Id       int
	FileKey  string
	SwiftUrl string
}

func Contracts() (contracts []Contract, err error) {
	rows, err := Db.Query("select id, file_key, swift_url from contracts")
	if err != nil {
		log.Fatal(err)
		return
	}
	for rows.Next() {
		contract := Contract{}
		if err := rows.Scan(&contract.Id, &contract.FileKey, &contract.SwiftUrl); err != nil {
			log.Fatal(err)
			continue
		}
		contracts = append(contracts, contract)
	}
	rows.Close()
	return
}

func ContractsWithoutSwift() (contracts []Contract) {
	rows, _ := Contracts()

	var result []Contract

	for _, v := range rows {
		if v.SwiftUrl == "" {
			result = append(result, v)
		}
	}
	return result

}
