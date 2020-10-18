package DB

import (
	"crypto/md5"
	"database/sql"
	"encoding/hex"
	"fmt"
	"github.com/astaxie/beego"
	_ "github.com/go-sql-driver/mysql"
	"intimate/models"
)

var Db *sql.DB

func init() {
	fmt.Println("链接mysql数据库")
	config := beego.AppConfig
	dbDriver := config.String("driverName")
	dbUser := config.String("db_user")
	dbPassword := config.String("db_password")
	dbIp := config.String("db_ip")
	dbName := config.String("db_name")
	connUrl := dbUser + ":" + dbPassword + "@tcp(" + dbIp + ")/" + dbName + "?charset=utf8"
	fmt.Println(connUrl)
	db, err := sql.Open(dbDriver, connUrl)
	if err != nil {
		fmt.Println(err.Error())
		panic("数据连接错误，请检查错误")
	}
	Db = db

}

func InsertUser(user models.UserModels)(int64,error){
	hashMD5pwd := md5.New()
	hashMD5idcard :=md5.New()
	hashMD5pwd.Write([]byte(user.PassWord))
	hashMD5idcard.Write([]byte(user.IDcard))
	pwdBytes := hashMD5pwd.Sum(nil)
	idcardBytes :=hashMD5idcard.Sum(nil)
	user.PassWord = hex.EncodeToString(pwdBytes)
	user.IDcard = hex.EncodeToString(idcardBytes)
	fmt.Println("保存的用户名：", user.Name, "密码：", user.PassWord,"身份证号:",user.IDcard)
	result, err := Db.Exec("insert into user(name,identification,phone,password) value(?,?,?,?)", user.Name,user.IDcard,user.Phone, user.PassWord)
	if err != nil {
		fmt.Println(err.Error())
		return -1, err
	}
	id, err := result.RowsAffected()
	if err != nil {
		fmt.Println(err.Error())
		return -1, err
	}
	return id, nil
}

func QueryUser(user models.UserModels)(*models.UserModels,error){
	fmt.Println(user.Name)
	hashMD5 := md5.New()
	hashMD5.Write([]byte(user.PassWord))
	bytes := hashMD5.Sum(nil)
	user.PassWord=hex.EncodeToString(bytes)
	fmt.Println(user.PassWord)
	row:=Db.QueryRow("select name from user where name = ? and password = ?",user.Name,user.PassWord)
	err:=row.Scan(&user.Name)
	if err != nil {
		return nil,err
	}
	return &user,nil
}