package controllers

import (
	"gintemplate/app/models"
	"strconv"

	"github.com/gin-gonic/gin"
)

func RooterInit(){
router := gin.Default()
router.LoadHTMLGlob("app/views/*.html")

router.GET("/", func(ctx *gin.Context) {
	todos := models.DbGetAll()
	ctx.HTML(200, "index.html", gin.H{
		"todos": todos,
	})
})

//Create
router.POST("/new", func(ctx *gin.Context) {
	text := ctx.PostForm("text")
	status := ctx.PostForm("status")
	models.DbInsert(text, status)
	ctx.Redirect(302, "/")
})

//Detail
router.GET("/detail/:id", func(ctx *gin.Context) {
	n := ctx.Param("id")
	id, err := strconv.Atoi(n)
	if err != nil {
		panic(err)
	}
	todo := models.DbGetOne(id)
	ctx.HTML(200, "detail.html", gin.H{"todo": todo})
})

//Update
router.POST("/update/:id", func(ctx *gin.Context) {
	n := ctx.Param("id")
	id, err := strconv.Atoi(n)

	if err != nil {
		panic("ERROR")
	}
	text := ctx.PostForm("text")
	status := ctx.PostForm("status")
	models.DbUpdate(id, text, status)
	ctx.Redirect(302, "/")
})

//削除確認
router.GET("/delete_check/:id", func(ctx *gin.Context) {
	n := ctx.Param("id")
	id, err := strconv.Atoi(n)
	if err != nil {
		panic("ERROR")
	}
	todo := models.DbGetOne(id)
	ctx.HTML(200, "delete.html", gin.H{"todo": todo})
})

//Delete
router.POST("/delete/:id", func(ctx *gin.Context) {
	n := ctx.Param("id")
	id, err := strconv.Atoi(n)
	if err != nil {
		panic("ERROR")
	}
	models.DbDelete(id)
	ctx.Redirect(302, "/")

})

router.Run()

}
