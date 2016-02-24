package controller

import (
	"encoding/gob"
	"html/template"

	"github.com/Miloas/oj/model"
	"github.com/garyburd/redigo/redis"
	"github.com/unrolled/render"
	"gopkg.in/mgo.v2"
)

var (
	//Render :render data var
	Render *render.Render
	//S :mgo session
	S            *mgo.Session
	databaseName = "oj"
	//RedisPool :Redis connection pool
	RedisPool *redis.Pool
)

//Key : type of context key
type Key int

//GlobalRequestVariable : value of context key
const GlobalRequestVariable Key = 0

// Init :init controller methods
func init() {
	funcMap := template.FuncMap{
		"isAccepeted": func(problemID string, haveAccepeted []string) bool {
			return CheckInStringArray(problemID, haveAccepeted)
		},
		"i2c": func(x int) string {
			s := []string{"A", "B", "C", "D", "E", "F", "G", "H", "I", "J", "K", "L", "M", "N", "O", "P", "Q", "R", "S", "T", "U", "V", "W", "X", "Y", "Z"}
			return s[x]
		},
	}
	Render = render.New(render.Options{
		Directory: "templates",
		Layout:    "layout",
		Funcs:     []template.FuncMap{funcMap},
	})
	RedisPool = newRedisPool()
	gob.Register(&model.User{})
}
