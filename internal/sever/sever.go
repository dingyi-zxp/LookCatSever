package sever

import (
	"fmt"
	"reflect"
	"strconv"

	"database/sql"

	"github.com/gin-gonic/gin"

	"LookCat/internal/linkSql"

	sqloper "LookCat/internal/sqloper"
)

type Actor struct {
	Id_actor   int    `json:"id_actor" from:"id_actor"`
	Name_actor string `json:"name_actor" from:"name_actor"`
	Img_actor  string `json:"img_actor" from:"img_actor"`
	Gender     int    `json:"gender" from:"gender"`
	Info_actor string `json:"info_actor" from:"info_actor"`
	Nation     string `json:"nation" from:"nation"`
}

type Gather struct {
	Id_gather int `json:"id_gather" from:"id_gather"`
	Id_videos int `json:"id_videos" from:"id_videos"`
	Ggather   int `json:"gather" from:"gather"`
}

type Intro struct {
	Id_info    int    `json:"id_info" from:"id_info"`
	Id_collect int    `json:"id_collect" from:"id_collect"`
	Sug_video  string `json:"sug_video" from:"sug_video"`
	Name_info  string `json:"name_info" from:"name_info"`
	Limit_age  string `json:"limit_age" from:"limit_age"`
	Img_info   string `json:"img_info" from:"img_info"`
}

func Sever() {

	r := gin.Default()
	fmt.Println(reflect.TypeOf(r))
}

func Init() {
	router := gin.Default()
	router.Use(CORS())
	var db *sql.DB
	db = linkSql.InitSql()
	defer db.Close()
	rows, err := db.Query("SELECT i.name_intro,i.intro_info,i.limit_age,i.sug_video,a.name_actor,logo_intro,img_info from video_collect c JOIN video_intro i on c.id_info=i.id_info JOIN actor a on c.id_actor=a.id_actor;")
	fmt.Println("eeeeu", rows, err)
	// 匹配/user/john
	router.POST("/VideoIntro", func(c *gin.Context) {
		introId := c.Query("introId")
		limitIntro := c.Query("limit")
		var id, limit int
		id, err := strconv.Atoi(introId)
		limit, limiterr := strconv.Atoi(limitIntro)
		fmt.Println(limit, id)
		if err != nil || limiterr != nil {
			return
		}

		intros := sqloper.VideoIntros(db, id, limit)
		c.JSON(200, gin.H{
			"data": intros,
		})
	})

	router.Run(":9860")

}

func CORS() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		ctx.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		ctx.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		ctx.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With, token, x-access-token")
		ctx.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT, DELETE")

		if ctx.Request.Method == "OPTIONS" {
			ctx.AbortWithStatus(204)
			return
		}
		ctx.Next()
	}
}
