package sqloper

import (
	"database/sql"
	"fmt"
	"os"
	"strconv"
)

type VideoIntro struct {
	Name       string `json:"name_intro" from:"name_intro"`
	Intro_info string `json:"intro_info" from:"intro_info"`
	Limit_age  string `json:"limit_age" from:"limit_age"`
	SuggVideo  string `json:"sug_video" from:"sug_video"`
	Name_actor string `json:"name_actor" from:"name_actor"`
	Logo_intro string `json:"logo_intro" from:"logo_intro"`
	Img_info   string `json:"img_info" from:"img_info"`
	Total_Time int    `json:"total_time" from:"total_time"`
}

func VideoIntros(db *sql.DB, introId int, limit int) []VideoIntro {
	fmt.Println(introId, limit)
	var sql = "SELECT i.name_intro,i.intro_info,i.limit_age,i.sug_video,a.name_actor,logo_intro,img_info,total_time from video_collect c JOIN video_intro i on c.id_info=i.id_info JOIN actor a on c.id_actor=a.id_actor where i.id_info=" + strconv.Itoa(introId) + " limit " + strconv.Itoa(limit) + ";"
	rows, err := db.Query(sql)

	if err != nil {
		os.Exit(3)
	}

	intros := make([]VideoIntro, 0)

	for rows.Next() {
		var intro VideoIntro

		rows.Scan(&intro.Name, &intro.Intro_info, &intro.Limit_age, &intro.SuggVideo, &intro.Name_actor, &intro.Logo_intro, &intro.Img_info, &intro.Total_Time)
		intros = append(intros, intro)
	}

	return intros
}
