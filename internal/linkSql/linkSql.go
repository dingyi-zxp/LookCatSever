package linkSql

import (
	sqlInfo "LookCat/configs"
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/go-sql-driver/mysql"
)

type Videos struct {
	Id        int    `json:"id_videos" from:"id_videos"`
	NameVideo string `json:"name_video"`
	Data      string `json:"data_video"`
}

type VideoIntro struct {
	Name       string `json:"name_intro" from:"name_intro"`
	Intro_info string `json:"intro_info" from:"intro_info"`
	Limit_age  string `json:"limit_age" from:"limit_age"`
	SuggVideo  string `json:"sug_video" from:"sug_video"`
	Name_actor string `json:"name_actor" from:"name_actor"`
	Logo_intro string `json:"logo_intro" from:"logo_intro"`
	Img_info   string `json:"img_info" from:"img_info"`
}

type ts struct {
	Name string `json:"name_intro" from:"name_intro"`
}

func InitSql() *sql.DB {
	var dsn = spliceDsn()

	db, err := sql.Open("mysql", dsn)

	if err != nil {
		fmt.Printf("Open dataset error: %v", err)
	}

	err = db.Ping()
	if err != nil {
		fmt.Printf("连接数据库失败,err:%v\n", err)
		os.Exit(3)
	}
	fmt.Printf("success")

	rows, err := db.Query("SELECT i.name_intro,i.intro_info,i.limit_age,i.sug_video,a.name_actor,logo_intro,img_info from video_collect c JOIN video_intro i on c.id_info=i.id_info JOIN actor a on c.id_actor=a.id_actor;")
	//
	// rows, err := db.Query("select * from video_intro")
	defer rows.Close()

	if err != nil {
		log.Fatalln(err)
	}

	videos := make([]VideoIntro, 0)

	for rows.Next() {
		var video VideoIntro
		rows.Scan(&video.Name, &video.Intro_info, &video.Limit_age, &video.SuggVideo, &video.Name_actor, &video.Logo_intro, &video.Img_info)
		fmt.Println(rows.Columns())
		videos = append(videos, video)
	}

	return db
}

func spliceDsn() string {
	var info sqlInfo.Conn
	info = sqlInfo.MysqlConfigs()

	var dsn string

	dsn = fmt.Sprintf("%s:%s@tcp(%s:%d)/lookcat?charset=utf8", info.User, info.Password, info.Address, info.Host)
	fmt.Println(dsn)
	return dsn
}
