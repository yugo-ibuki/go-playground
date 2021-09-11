package main

import "github.com/gin-gonic/gin"
import "net/http"

func main() {
	engine:= gin.Default()
	engine.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "hello world",
		})
	})
	engine.Run(":3001")
}

//import "fmt"
//
//type Food struct {
//	Name string
//	Type string
//}
//
//func (a Food) name() string {
//	return a.Name
//}
//
//type Fruit struct {
//	Food
//}
//
//func main() {
//	// for文
//	//arr := [...]int{1,2,3,4,5,6}
//	//for _, i := range arr {
//	//	fmt.Println(i)
//	//}
//
//	//f := Food{"beef", "meet"}
//	//fmt.Println("This is ", f.name())
//
//	mapEx := make(map[string]string) //マップの宣言
//	fmt.Println(mapEx) //=> map[] //宣言された空のマップ
//
//	mapEx["firstName"] = "Mike" //マップにkeyとvalueを挿入する。
//	mapEx["lastName"] = "Smith"
//
//	fmt.Println(mapEx) //=> map[lastName:Smith firstName:Mike]
//}
