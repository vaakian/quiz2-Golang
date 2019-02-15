package route
import (
	"fmt"
	"github.com/jinzhu/gorm"
	"github.com/gin-gonic/gin"
	"github.com/Unknwon/goconfig"
	. "../module"
	_"github.com/Go-SQL-Driver/MySQL"
	"strconv"
)
var db *gorm.DB
func init() {
	cfg, err := goconfig.LoadConfigFile("config.ini")
	if err != nil{
	    panic("can't read config file `config.ini`")
	}
	sql, err := cfg.GetValue("mysql", "url")
	if err != nil {
		panic("can't read config `mysql-> url`")
	}
	// err := nil
	db, _ = gorm.Open("mysql", sql)
    // if err != nil {
    //     panic(err)
    // }
}
func Index(c *gin.Context) {
	users := []User{}
	usr := c.Query("usr")
	count, _ := strconv.Atoi(c.Query("count"))
	db.Limit(count).Where("username LIKE ? OR password LIKE ?", fmt.Sprintf("%%%s%%", usr), fmt.Sprintf("%%%s%%", usr)).Find(&users)
	fmt.Printf("got %d data of kw `%s`\n", len(users), usr)
    c.JSON(200, gin.H{
    	"status": 200,
    	"msg": "well done.",
    	"data": users,
    	})
	}
func RandQuestion(c *gin.Context) {
	cg := c.Query("cg")
	questions := []Question{}
	db.Where("Category_id = ?", cg).Find(&questions)
	fmt.Printf("RandQuestion %d data of cg %s\n", len(questions), cg)
	c.JSON(200, gin.H{
		"status": 200,
		"msg": "ok",
		"data": questions,
		})
}