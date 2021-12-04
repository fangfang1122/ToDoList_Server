package api

//type Auth struct {
//	Name       string `json:"name" binding:"required"`
//	Username   string `form:"username" json:"username" binding:"required,min=2"`
//	Password   string `form:"password" json:"password" binding:"required,min=6" `
//	CreateById int    `json:"create_by_id" binding:"required"`
//}
//
//func Login(c *gin.Context) {
//	var json struct {
//		Username string `form:"username" json:"username" binding:"required,min=2"`
//		Password string `form:"password" json:"password" binding:"required,min=6" `
//	}
//	if !BindAndValid(c, &json) {
//		return
//	}
//
//	data := make(map[string]interface{})
//	code := e.INVALID_PARAMS
//	user, err := models.FindAuth(json.Username)
//
//	if err != nil {
//		if errors.Is(err, gorm.ErrRecordNotFound) {
//			code = e.ERROR_NOT_EXIST_USER
//			c.JSON(http.StatusForbidden, gin.H{
//				"message": e.GetMsg(code),
//				"code":    code,
//			})
//			return
//		}
//		ErrHandler(c, err)
//		return
//	}
//
//	if !util.PasswordVerify(user.Password, json.Password) {
//		code = e.ERROR_PASSWORD
//		c.JSON(http.StatusForbidden, gin.H{
//			"message": e.GetMsg(code),
//			"code":    code,
//		})
//		return
//	}
//
//	token, err := util.GenerateToken(user.ID, user.Username)
//	if err != nil {
//		code = e.ERROR_AUTH_TOKEN
//	} else {
//		data["token"] = token
//		code = e.SUCCESS
//	}
//
//	c.JSON(http.StatusOK, gin.H{
//		"code": code,
//		"msg":  e.GetMsg(code),
//		"data": data,
//	})
//
//}
//
//func UserInfo(c *gin.Context) {
//	user := CurrentUser(c)
//	c.JSON(http.StatusOK, user)
//}
//
//func AddAuth(c *gin.Context) {
//	var json Auth
//	if !BindAndValid(c, &json) {
//		return
//	}
//
//	code := e.ERROR
//	if models.ExistAuthByName(json.Username) {
//		code = e.ERROR_EXIST_USER
//		c.JSON(http.StatusBadRequest, gin.H{
//			"message": e.GetMsg(code),
//			"code":    code,
//		})
//		return
//	}
//
//	auth := models.Auth{
//		Name:       json.Name,
//		Username:   json.Username,
//		Password:   json.Password,
//		CreateById: json.CreateById,
//	}
//	err := auth.Save()
//	if err != nil {
//		ErrHandler(c, err)
//		return
//	}
//	code = e.SUCCESS
//	c.JSON(http.StatusOK, gin.H{
//		"code": code,
//		"msg":  e.GetMsg(code),
//	})
//}
//
//func GetAuthList(c *gin.Context) {
//	data := models.GetAuthList(c)
//	c.JSON(http.StatusOK, data)
//}
//
//func EditAuth(c *gin.Context) {
//	var json struct {
//		Name       string `json:"name" binding:"required"`
//		Username   string `form:"username" json:"username" binding:"required,min=2"`
//		Password   string `form:"password" json:"password"  `
//		CreateById int    `json:"create_by_id" binding:"required"`
//	}
//	if !BindAndValid(c, &json) {
//		return
//	}
//
//	// 需要判断编辑用户名的话是否重复，返回前端。    目前未处理，以数据库自动检查。
//	//if models.ExistAuthByName(json.Username) {
//	//	code := e.ERROR_EXIST_USER
//	//	c.JSON(http.StatusBadRequest, gin.H{
//	//		"message": e.GetMsg(code),
//	//		"code":    code,
//	//	})
//	//	return
//	//}
//
//	auth, err := models.GetAuthById(c.Param("id"))
//	if err != nil {
//		ErrHandler(c, err)
//		return
//	}
//	var password string
//	if json.Password != "" {
//		if len(json.Password) < 6 {
//			c.JSON(http.StatusBadRequest, gin.H{
//				"message": "密码长度不能小于6位",
//				"code":    e.ERROR_PASSWORD,
//			})
//		}
//		password = json.Password
//	} else {
//		password = auth.Password
//	}
//	n := models.Auth{
//		Name:       json.Name,
//		Username:   json.Username,
//		Password:   password,
//		CreateById: json.CreateById,
//	}
//	auth.Updates(&n)
//	c.JSON(http.StatusOK, auth)
//}
//
//func DeleteAuthById(c *gin.Context) {
//	auth, err := models.GetAuthById(c.Param("id"))
//	if err != nil {
//		ErrHandler(c, err)
//		return
//	}
//	if err = auth.Delete(); err != nil {
//		ErrHandler(c, err)
//		return
//	}
//	c.JSON(http.StatusOK, gin.H{
//		"msg": "ok",
//	})
//}
