package data

import (
	"database/sql"
	"log"
)

type Rate struct {
	Id        int
	FileKeys  sql.NullString
	SwiftUrls sql.NullString
}

func Rates() (rates []Rate, err error) {
	rows, err := Db.Query("select id,file_keys,swift_urls from rates")
	if err != nil {
		log.Fatal(err)
	}
	for rows.Next() {
		rate := Rate{}

		if err = rows.Scan(&rate.Id, &rate.FileKeys, &rate.SwiftUrls); err != nil {
			log.Fatal(err)
			return
		}
		rates = append(rates, rate)
	}
	return
}

func RatesWithoutSwift() (rates []Rate) {
	rows, _ := Rates()

	var result []Rate

	for _, v := range rows {
		if v.SwiftUrls.String != "" {
			continue
		}
		if v.FileKeys.String == "" || len(v.FileKeys.String) < 3 {
			continue
		}
		result = append(result, v)
	}
	return result
}
