package module

type User struct {
  //gorm.Model
  Id        uint `gorm:"primary_key" json:"id"`
  Usr string `gorm:"column:username" json:"usr"`
  Pwd string `gorm:"column:password" json:"pwd"`
  Time string `json:"time"`
  Token string `json:"token"`
  Qq string `json:"qq"`
  Slogan string `json:"slogan"`
}
type Question struct {
	ID uint `gorm:"column:ID" json:"ID"`
	Title string `gorm:"column:title"`
	A string `gorm:"column:A"`
	B string `gorm:"column:B"`
	C string `gorm:"column:C"`
	D string `gorm:"column:D"`
	Answer string `json:"answer"`
	Ana string `json:"ana"`
	Chapter string `json:"chapter"`
	Hard string `json:"hard"`
	Category string `json:"category"`
	Category_id uint `json:"category_id"`
}
// type Fav struct {
// }
