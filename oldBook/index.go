package main

import (
	 "net/http"
     "fmt"
     "strconv"
    //  "io/ioutil"
     "encoding/json"
     "database/sql"
    "github.com/gin-gonic/gin"
    _ "github.com/go-sql-driver/mysql"
    "time"
    "strings"
    "github.com/dgrijalva/jwt-go" 
)
var (
	SIGN_NAME_SCERET = "aweQurt178BNI"
)

type User struct{
    Name        string  `json:"name"`
    Password    string  `json:"password"`
}
type Book struct{
    Title      string  `json:"title"`
    Message    string  `json:"message"`
    Price      string  `json:"price"`
    Kind       string  `json:"kind"`
    Images     []string `json:"images"`
}
func zhuce(c *gin.Context) {

	buf := make([]byte, 1024)
    n, _ := c.Request.Body.Read(buf)
    var userinfo User
	if err := json.Unmarshal([]byte(string(buf[0:n])), &userinfo) ; err == nil {
        fmt.Println(userinfo.Name)
        db,_:=sql.Open("mysql","root:xfz123456@(127.0.0.1:3306)/oldbooks") // 设置连接数据库的参数
        defer db.Close()    //关闭数据库
        err:=db.Ping()      //连接数据库
        if err!=nil{
            fmt.Println("数据库连接失败")
            return
        }
        id := time.Now().Unix()
        db.Exec("insert into userinfo(id,name,password) values (?,?,?)",id,userinfo.Name,userinfo.Password)     
        fmt.Println("注册成功！")
        resp := map[string]string{"message": "注册成功！"}
        c.JSON(http.StatusOK, resp)
    }
	
    
}
func denglu(c *gin.Context) {
   getUserBooks:= func(user_id interface{}) interface{}{
        
        db,_:=sql.Open("mysql","root:xfz123456@(127.0.0.1:3306)/oldbooks") // 设置连接数据库的参数
        defer db.Close()    //关闭数据库
        err:=db.Ping()      //连接数据库
        if err!=nil{
            fmt.Println("数据库连接失败")
            return nil
        }
        result,err:=db.Query("select id,title,messages,price,images from book_lists where owner_id=?",user_id)      //多行查询
        if err == nil{
    
        var id,title,messages,price,images string
        i :=0
        book :=map[int]interface{}{}
        for result.Next(){        //循环显示所有的数据
             result.Scan(&id,&title,&messages,&price,&images)
             data :=map[string]interface{}{"id":id,"title":title,"messages":messages,"price":price,"images":images}
             book[i] =data
             i++
        }
        return book
    }
    return nil
    }

    
    cookie,err :=c.Cookie("denglu")
    if err != nil {
	buf := make([]byte, 1024)
    n, _ := c.Request.Body.Read(buf)
    var userinfo User
	if err := json.Unmarshal([]byte(string(buf[0:n])), &userinfo) ; err == nil {
        
        db,_:=sql.Open("mysql","root:xfz123456@(127.0.0.1:3306)/oldbooks") // 设置连接数据库的参数
        defer db.Close()    //关闭数据库
        err:=db.Ping()      //连接数据库
        if err!=nil{
            fmt.Println("数据库连接失败")
            return
        }
        // id := time.Now().Unix()
        resp := map[string]string{"message": "登录失败！"}
        result:=db.QueryRow("select password,id,userimage from userinfo where name=?",userinfo.Name)      //单行查询
        var password,id,userimage string
        result.Scan(&password,&id,&userimage)
        // fmt.Println(result)
        if userinfo.Password == password {
            tokenString, err := createJwt(id)
        if err != nil {
            fmt.Println(err.Error())
            return
        }
            resp = map[string]string{"message": "登录成功！","name":userinfo.Name,"token":tokenString,"userimage":userimage}
        }
        
        c.JSON(http.StatusOK, resp)
    }
    }else{
        claims := parseJwt(strings.Split(cookie,",")[0])
        user_id  := claims["user_id"]
        db,_:=sql.Open("mysql","root:xfz123456@(127.0.0.1:3306)/oldbooks") // 设置连接数据库的参数
        defer db.Close()    //关闭数据库
        err:=db.Ping()      //连接数据库
        if err!=nil{
            fmt.Println("数据库连接失败")
            return
        }
        // id := time.Now().Unix()
        
        result:=db.QueryRow("select name,userimage from userinfo where id=?",user_id)      //单行查询
        var name,userimage string
        result.Scan(&name,&userimage)
        book:= getUserBooks(user_id)
        resp := map[string]interface{}{"message": "登录成功！","user_id": user_id.(string),"name":name,"userimage":userimage,"book_list":book}
        c.JSON(http.StatusOK, resp)
    }

    
    
}
func sendBook(c *gin.Context){
    cookie,err :=c.Cookie("denglu")
    if err != nil {

    }else{
        // claims := parseJwt(strings.Split(cookie,",")[0])
        // user_id  := claims["user_id"]
        file, err := c.FormFile("img")
        if err != nil {
           claims := parseJwt(strings.Split(cookie,",")[0])
            owner_id  := claims["user_id"]
            var book Book
            buf := make([]byte, 1024)
            n, _ := c.Request.Body.Read(buf)
            
            if err := json.Unmarshal([]byte(string(buf[0:n])), &book) ; err == nil {
                fmt.Println(book.Images)
                db,_:=sql.Open("mysql","root:xfz123456@(127.0.0.1:3306)/oldbooks") // 设置连接数据库的参数
                defer db.Close()    //关闭数据库
                err:=db.Ping()      //连接数据库
                if err!=nil{
                    fmt.Println("数据库连接失败")
                    return
                }
                image := book.Images[0]
                for i:=1;i<len(book.Images);i++ {
                    image = image +"|" + book.Images[i]
                }
                fmt.Println(image)
                id := time.Now().Unix()
                db.Exec("insert into book_lists(id,title,messages,price,owner_id,kind,images) values (?,?,?,?,?,?,?)",id,book.Title,book.Message,book.Price,owner_id,book.Kind,image)     
                fmt.Println("上架成功！")
                resp := map[string]string{"message": "上架成功！"}
                c.JSON(http.StatusOK, resp)
            }else{
                resp := map[string]string{"message": "数据错误！"}
                c.JSON(http.StatusOK, resp)
            }
        }else{  
        c.SaveUploadedFile(file, "./public/booksImages/"+file.Filename)
        resp := map[string]string{"message": "上传成功！"}
        c.JSON(http.StatusOK, resp)
        }
    }
    
}
func booklists(c *gin.Context){
    db,_:=sql.Open("mysql","root:xfz123456@(127.0.0.1:3306)/oldbooks") // 设置连接数据库的参数
    defer db.Close()    //关闭数据库
    err:=db.Ping()      //连接数据库
    if err!=nil{
        fmt.Println("数据库连接失败")
        return
    }
    // id := time.Now().Unix()
    resp := map[string]interface{}{"message": "获取失败！"}
    kind := c.Param("kind")
    result,err:=db.Query("select id,title,messages,price,owner_id,images from book_lists where kind=?",kind)      //多行查询
    if err == nil{

    var id,title,messages,price,owner_id,images string
    i :=0
    book :=map[int]interface{}{}
    for result.Next(){        //循环显示所有的数据
         result.Scan(&id,&title,&messages,&price,&owner_id,&images)
         data :=map[string]interface{}{"id":id,"title":title,"messages":messages,"price":price,"owner_id":owner_id,"images":images}
         book[i] =data
         i++
    }
   
    // fmt.Println(result)
    
    
    resp = map[string]interface{}{"message": "获取成功！","book_list":book}
    }
    
    c.JSON(http.StatusOK, resp)
}

func bookInfo(c *gin.Context){
    db,_:=sql.Open("mysql","root:xfz123456@(127.0.0.1:3306)/oldbooks") // 设置连接数据库的参数
    defer db.Close()    //关闭数据库
    err:=db.Ping()      //连接数据库
    if err!=nil{
        fmt.Println("数据库连接失败")
        return
    }
    // id := time.Now().Unix()
    resp := map[string]interface{}{"message": "获取失败！"}
    id := c.Param("id")
    result:=db.QueryRow("select title,messages,price,owner_id,images,kind from book_lists where id=?",id)      //单行查询
    var title,messages,price,owner_id,images,kind string
         result.Scan(&title,&messages,&price,&owner_id,&images,&kind)
         int64,_:= strconv.ParseInt(id, 10, 64)  
         date:= time.Unix(int64, 0).Format("2006-01-02")
         book :=map[string]interface{}{"id":id,"title":title,"messages":messages,"price":price,"owner_id":owner_id,"images":images,"kind":kind}
    result1:=db.QueryRow("select name from userinfo where id=?",owner_id)    
    var name string
    result1.Scan(&name)
    resp = map[string]interface{}{"message": "获取成功！","book_list":book,"user_name":name,"date":date}
    
    
    c.JSON(http.StatusOK, resp)
}
func main() {
    	
    router := gin.Default()
//跨域中间件  
    router.Use(Cors())
//首页路由
    router.StaticFile("/home", "./public/home.html")
//资源路由
    router.StaticFS("/public", http.Dir("./public"))
//注册路由
    router.POST("/zhuce", zhuce)
//登录路由
    router.POST("/denglu", denglu)
//上架功能
    router.POST("/sendBook", sendBook)
//获取books列表
    router.GET("/booklists/:kind",booklists)
//获取单个book数据
    router.GET("/book/:id",bookInfo)
	router.Run(":8000")
    
}
//跨域处理
func Cors() gin.HandlerFunc {
	return func(c *gin.Context) {
		method := c.Request.Method		//请求方法
		origin := c.Request.Header.Get("Origin")		//请求头部
		var headerKeys []string								// 声明请求头keys
		for k, _ := range c.Request.Header {
			headerKeys = append(headerKeys, k)
		}
		headerStr := strings.Join(headerKeys, ", ")
		if headerStr != "" {
			headerStr = fmt.Sprintf("access-control-allow-origin, access-control-allow-headers, %s", headerStr)
		} else {
			headerStr = "access-control-allow-origin, access-control-allow-headers"
		}
		if origin != "" {
			c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
			c.Header("Access-Control-Allow-Origin", "http://127.0.0.1:8000")		// 这是允许访问所有域
			c.Header("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE,UPDATE")		//服务器支持的所有跨域请求的方法,为了避免浏览次请求的多次'预检'请求
			//  header的类型
			c.Header("Access-Control-Allow-Headers", "Authorization, Content-Length, X-CSRF-Token, Token,session,X_Requested_With,Accept, Origin, Host, Connection, Accept-Encoding, Accept-Language,DNT, X-CustomHeader, Keep-Alive, User-Agent, X-Requested-With, If-Modified-Since, Cache-Control, Content-Type, Pragma")
			//				允许跨域设置																										可以返回其他子段
			c.Header("Access-Control-Expose-Headers", "Content-Length, Access-Control-Allow-Origin, Access-Control-Allow-Headers,Cache-Control,Content-Language,Content-Type,Expires,Last-Modified,Pragma,FooBar")		// 跨域关键设置 让浏览器可以解析
			c.Header("Access-Control-Max-Age", "172800")		// 缓存请求信息 单位为秒
			c.Header("Access-Control-Allow-Credentials", "true")		//	跨域请求是否需要带cookie信息 默认设置为true
			c.Set("content-type", "application/json")		// 设置返回格式是json
		}
 
		//放行所有OPTIONS方法
		if method == "OPTIONS" {
			c.JSON(http.StatusOK, "Options Request!")
		}
		// 处理请求
		c.Next()		//	处理请求
	}
}
//创建 token
func createJwt(id string) (string, error) {
	//	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
	//		"foo": "bar",
	//		"nbf": time.Date(2015, 10, 10, 12, 0, 0, 0, time.UTC).Unix(),

	//	})

	token := jwt.New(jwt.SigningMethodHS256)
	claims := make(jwt.MapClaims)
	claims["user_id"] = id 
	claims["exp"] = time.Now().Add(time.Hour * time.Duration(1)).Unix()
	claims["iat"] = time.Now().Unix()
	token.Claims = claims

	// Sign and get the complete encoded token as a string using the secret
	tokenString, err := token.SignedString([]byte(SIGN_NAME_SCERET))
	return tokenString, err
}
//解析 token
func parseJwt(tokenString string) jwt.MapClaims {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		// Don't forget to validate the alg is what you expect:
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}

		// hmacSampleSecret is a []byte containing your secret, e.g. []byte("my_secret_key")
		return []byte(SIGN_NAME_SCERET), nil
	})

	var claims jwt.MapClaims
	var ok bool

	if claims, ok = token.Claims.(jwt.MapClaims); ok && token.Valid {
        // fmt.Println(claims["user_id"], claims["nbf"])
        return claims
	} else {
		fmt.Println(err)
	}

	return claims
}
