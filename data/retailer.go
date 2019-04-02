package data

import (
	"database/sql"
	"log"
)

type Retailer struct {
	Id             int
	AvatarKey      sql.NullString
	AvatarSwiftUrl sql.NullString
}

func Retailers() (retailers []Retailer, err error) {
	rows, err := Db.Query("select id, avatar_key, avatar_swift_url from retailers")
	if err != nil {
		log.Fatal(err)
		return
	}
	for rows.Next() {
		retailer := Retailer{}
		if err := rows.Scan(&retailer.Id, &retailer.AvatarKey, &retailer.AvatarSwiftUrl); err != nil {
			log.Fatal(err)
			continue
		}
		retailers = append(retailers, retailer)
	}
	rows.Close()
	return
}

func RetailersWithoutSwift() (retailers []Retailer) {
	rows, _ := Retailers()

	var result []Retailer

	for _, v := range rows {
		if v.AvatarSwiftUrl.String == "" {
			result = append(result, v)
		}
	}
	return result

}
