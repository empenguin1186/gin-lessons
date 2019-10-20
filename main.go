package main

import (
	_ "net/http"
	"strconv"

	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	_ "github.com/mattn/go-sqlite3"
)

type Person struct {
	gorm.Model
	Name string
	Age  int
}

type Login struct {
	User     string `form:"user" json:"user" binding:"required"`
	Password string `form:"password" json:"password" binding:"required"`
}

func db_init() {
	db, err := gorm.Open("sqlite3", "test.sqlite3")
	if err != nil {
		panic("failed to connect database\n")
	}

	db.AutoMigrate(&Person{})
}

func create(name string, age int) {
	db, err := gorm.Open("sqlite3", "test.sqlite3")
	if err != nil {
		panic("failed to connect database\n")
	}
	db.Create(&Person{Name: name, Age: age})
}

func get_all() []Person {
	db, err := gorm.Open("sqlite3", "test.sqlite3")
	if err != nil {
		panic("failed to conncet database\n")
	}
	var people []Person
	db.Find(&people)
	return people
}

func main() {
	r := gin.Default()
	r.LoadHTMLGlob("templates/*")
	db_init()
	r.GET("/", func(c *gin.Context) {
		people := get_all()
		c.HTML(200, "index.tmpl", gin.H{
			"people": people,
		})
	})
	r.POST("/new", func(c *gin.Context) {
		name := c.PostForm("name")
		age, _ := strconv.Atoi(c.PostForm("age"))
		create(name, age)
		c.Redirect(302, "/")
	})

	// Login
	store := cookie.NewStore([]byte("secret"))
	r.Use(sessions.Sessions("mysession", store))

	r.GET("/login", func(c *gin.Context) {
		session := sessions.Default(c)
		var count int
		v := session.Get("count")
		if v == nil {
			count = 0
		} else {
			count = v.(int)
			count++
		}
		session.Set("count", count)
		session.Save()
		c.HTML(200, "login.tmpl", gin.H{
			"count": count,
		})
	})
	r.POST("/login", func(c *gin.Context) {

		userName := c.PostForm("userName")
		password := c.PostForm("password")

		if userName != "manu" || password != "hoge" {
			c.Redirect(302, "/login")
		}

		c.Redirect(302, "/")
	})

	r.Run()
}
