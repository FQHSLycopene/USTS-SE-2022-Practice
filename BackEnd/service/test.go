package service

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
)

// BodyJSONTest
// @Summary	BodyJSONTest
// @Tags	测试方法
// @Param param body string true "上传的JSON"
// @Success	200  {string}  json{"code":"200","msg":"","data",""}
// @Router	/BodyJSONTest [post]
func BodyJSONTest(c *gin.Context) {
	data, _ := c.GetRawData()
	var m map[string]interface{}
	// 反序列化
	_ = json.Unmarshal(data, &m)

	fmt.Println(m)
}

// PostArrayTest
// @Summary	PostArrayTest
// @Tags	测试方法
// @Param param formData array true "array"
// @Success	200  {string}  json{"code":"200","msg":"","data",""}
// @Router	/PostArrayTest [post]
func PostArrayTest(c *gin.Context) {
	values := c.PostFormMap("param")
	fmt.Println(values)

}
