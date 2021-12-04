package api

import (
	"ToDoList_Go/models"
	"ToDoList_Go/pkg/setting"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
)

func BindAndValid(c *gin.Context, v interface{}) bool {
	return bindAndValid(c, v)
}

func ErrHandler(c *gin.Context, err error) {
	log.Println(err.Error())
	if err == gorm.ErrRecordNotFound {
		c.JSON(http.StatusNotFound, gin.H{
			"message": err.Error(),
		})
		return
	}
	c.JSON(http.StatusInternalServerError, gin.H{
		"message": err.Error(),
	})
}

func CurrentUser(c *gin.Context) *models.User {
	return c.MustGet("user").(*models.User)
}

func GetOpenId(c *gin.Context, code string) ( bool ,string) {
	params := url.Values{}
	Url, _:= url.Parse("https://api.weixin.qq.com/sns/jscode2session")
	APPID := setting.APPID
	AppSecret:=setting.AppSecret
	params.Set("appid",APPID)
	params.Set("secret",AppSecret)
	params.Set("js_code",code)
	params.Set("grant_type","authorization_code")
	//如果参数中有中文参数,这个方法会进行URLEncode
	Url.RawQuery = params.Encode()
	urlPath := Url.String()
	//fmt.Println(urlPath) //等同于https://www.xxx.com?age=23&name=zhaofan
	resp,_ := http.Get(urlPath)
	defer resp.Body.Close()
	body, _ := ioutil.ReadAll(resp.Body)
	println("json:", string(body))

	resMap := make(map[string]interface{})
	if err := json.Unmarshal(body, &resMap); err != nil {
		fmt.Printf("Unmarshal err, %v\n", err)
		ErrHandler(c, err)
		return false,""
	}
	fmt.Println(resMap)
	if resMap["openid"]==nil{
		fmt.Println("error",resMap)
		c.JSON(http.StatusBadRequest,gin.H{
			"msg":"error",
			"data":resMap,
		})
		return false, ""
	}else {
		return true,resMap["openid"].(string)
	}

}
