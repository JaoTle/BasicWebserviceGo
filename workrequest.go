package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

// customize ชื่อ key
type Course struct {
	ID         int     `json: "id"`
	Name       string  `json: "name"`
	Price      float64 `json: "price"`
	Instructor string  `json: "instructor"`
}

var CourseList []Course

func init() {
	CourseJson := `[
		{
			"id" : 1,
			"name" : "Python",
			"price" : 2590,
			"instructor" : "BorntoDev"
		},
		{
			"id" : 2,
			"name" : "Javascript",
			"price" : 0,
			"instructor" : "BorntoDev"
		},
		{
			"id" : 3,
			"name" : "SQL",
			"price" : 0,
			"instructor" : "BorntoDev"
		}
	]`

	err := json.Unmarshal([]byte(CourseJson), &CourseList)
	if err != nil {
		log.Fatal(err)
	}
}

func courseHandler(w http.ResponseWriter, r *http.Request) {
	coursejson, err := json.Marshal(CourseList)
	switch r.Method {
	case http.MethodGet:
		if err != nil {

			//response error
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		//no error
		w.Header().Set("Content-Type", "application/json")
		w.Write(coursejson)
	case http.MethodPost:
		//add data
		var newCourse Course

		//อ่านข้อมูลทั้งหมด
		Bodybyte, err := ioutil.ReadAll(r.Body)
		if err != nil {
			fmt.Println(err)
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		//รับค่าBodybyte
		err = json.Unmarshal(Bodybyte, &newCourse)
		if err != nil {
			fmt.Println(err)
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		//check ID condition
		if newCourse.ID != 0 {
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		newCourse.ID = getNextID()
		CourseList = append(CourseList, newCourse)
		w.WriteHeader(http.StatusCreated)
		return

	}
}

func getNextID() int {
	highestID := -1
	//loop list
	for _, course := range CourseList {
		if highestID < course.ID {
			highestID = course.ID
		}
	}
	return highestID + 1
}

func main() {
	http.HandleFunc("/course", courseHandler)
	http.ListenAndServe(":5000", nil)
}
