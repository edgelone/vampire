package data

import "log"

type RateUser struct {
	Id             int
	AvatarKey      string
	AvatarSwiftUrl string
}

func RateUsers() (rateUsers []RateUser, err error) {
	rows, err := Db.Query("select id, avatar_key, avatar_swift_url from rateUsers")
	if err != nil {
		log.Fatal(err)
		return
	}
	for rows.Next() {
		rateUser := RateUser{}
		if err := rows.Scan(&rateUser.Id, &rateUser.AvatarKey, &rateUser.AvatarSwiftUrl); err != nil {
			log.Fatal(err)
			continue
		}
		rateUsers = append(rateUsers, rateUser)
	}
	rows.Close()
	return
}

func RateUsersWithoutSwift() (rateUsers []RateUser) {
	rows, _ := RateUsers()

	var result []RateUser

	for _, v := range rows {
		if v.AvatarSwiftUrl == "" {
			result = append(result, v)
		}
	}
	return result

}
