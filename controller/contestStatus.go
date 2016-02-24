package controller

import (
	"net/http"
	"strconv"

	"gopkg.in/mgo.v2/bson"

	"github.com/Miloas/oj/model"
)

type contestStatusPageStruct struct {
	ContestID    string
	CurrentPage  int
	NextPage     int
	PreviousPage int
	CanNext      bool
	CanPrevious  bool
	Status       []model.Status
	Islogin      bool
}

//HandleContestStatus : handle contest status page
func HandleContestStatus(w http.ResponseWriter, r *http.Request) {
	cid := r.URL.Query().Get("cid")
	p := 0
	if tmp := r.URL.Query().Get("page"); tmp != "" {
		p, _ = strconv.Atoi(tmp)
	}
	session := getMongoS()
	defer session.Close()
	c := session.DB("oj").C("status")
	//normal submit status , not contest
	count, err := c.Find(bson.M{"contestid": cid}).Count()
	totalPage := (count + statusPageNum - 1) / statusPageNum
	status := []model.Status{}
	err = c.Find(bson.M{"contestid": cid}).Sort("-submittime").Limit(statusPageNum).Skip(statusPageNum * p).All(&status)
	if err != nil {
		panic(err)
	}
	canNext, canPrevious := false, false
	if p+1 < totalPage {
		canNext = true
	}
	if p-1 >= 0 {
		canPrevious = true
	}
	result := contestStatusPageStruct{cid, p, p + 1, p - 1, canNext, canPrevious, status, GetIslogin(r)}
	Render.HTML(w, http.StatusOK, "contestStatus", result)

}
