package main

import (
	"errors"
	"fmt"
	"net/http"
	"unicode/utf8"

	"github.com/gin-gonic/gin"
)

func handler(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "Hi there, I love %s!", r.URL.Path[1:])
}

func main() {

	var intNum int16 = 1211

	var floatNum float32 = 212.23

	var result = floatNum + float32(intNum)

	fmt.Println(result)

	myString := "me hu jiyaan"

	fmt.Println(utf8.RuneCountInString(myString))

	var myBoolean bool = true

	fmt.Println(!myBoolean)

	const name string = "jonas"

	const age int8 = 16

	var value, err = printMe(name,age)
	
	if err!=nil{
		fmt.Println(err)
	}else{
		fmt.Println(value)
	}

	var myArr []int8 = []int8{1,2,4}
	
	// fmt.Println(myArr)

	myArr = append(myArr, 9)

	var appendArr []int8 = []int8{3,4,6}

	myArr = append(myArr, appendArr...)

	// fmt.Println(myArr)

	for i := 0; i < len(myArr); i++ {
		fmt.Println(myArr[i])
	}

	r := gin.Default()
	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	r.Run() // listen and serve on 0.0.0.0:8080

	// r:= mux.NewRouter()
	// r.HandleFunc(`/`,serveHome).Methods("GET")
	
	// log.Fatal((http.ListenAndServe(":3000",r)))
	// PerformGetRequest()
}

func printMe(name string, age int8) (string, error){
	var err error

	if age<19{
		err = errors.New(name + " " + "Is not allowed to drive")
		return "", err
	}else{
		result:= name + " " + "can" + " " + "drive"
		return result, err
	}
}

func serveHome(w http.ResponseWriter, q *http.Request){
w.Write([]byte("<h1>Hello from Golang Server</h1>"))
}