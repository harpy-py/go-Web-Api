package migrations

import (
	"github.com/harpy-py/go-Web-Api/config"
	"github.com/harpy-py/go-Web-Api/constants"
	"github.com/harpy-py/go-Web-Api/data/db"
	"github.com/harpy-py/go-Web-Api/data/models"
	"github.com/harpy-py/go-Web-Api/pkg/logging"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

var logger = logging.NewLogger(config.GetConfig())

func Up_1() {
	database := db.GetDb()
	createTables(database)
	createDafaultInformation(database)
}

func createTables(database *gorm.DB){
	tables := []interface{}{}

	country := models.Country{}
	city := models.City{}
	user := models.User{}
	role := models.Role{}
	userRole := models.UserRole{}

	tables = addNewTable(database, country, tables)
	tables = addNewTable(database, city, tables)
	tables = addNewTable(database, user, tables)
	tables = addNewTable(database, role, tables)
	tables = addNewTable(database, userRole, tables)
	
	database.Migrator().CreateTable(tables...)
	logger.Info(logging.Postgres, logging.Migration, "Tables created!", nil)
}

func addNewTable(database *gorm.DB, model interface{}, tables []interface{}) []interface{} {
	if !database.Migrator().HasTable(model){
		tables = append(tables, model)
	}
	return tables
}

func createDafaultInformation(database *gorm.DB) {
	adminRole := models.Role{Name: constants.AdminRoleName}
	createRoleIfNotExists(database, &adminRole)

	defaultRole := models.Role{Name: constants.DefaultRoleName}
	createRoleIfNotExists(database, &defaultRole)

	u := models.User{Username: constants.DefaultUserName, FirstName: "Test", LastName: "Fist",
	MobileNumber: "09212222222", Email: "test@gmail.com"}
	pass := "123456"
	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(pass), bcrypt.DefaultCost)
	u.Password = string(hashedPassword)
	createAdminUserIfNotExists(database, &u, adminRole.Id)
}

func createRoleIfNotExists(database *gorm.DB, r *models.Role) {
	exists := 0

	database.
		Model(&models.Role{}).
		Select("1").
		Where("name = ?", r.Name).
		First(&exists)
	if exists == 0 {
		database.Create(r)
	}
}

func createAdminUserIfNotExists(database *gorm.DB, u *models.User, roleId int) {
	exists := 0

	database.
		Model(&models.Role{}).
		Select("1").
		Where("name = ?", u.Username).
		First(&exists)
	if exists == 0 {
		database.Create(u)
		ur := models.UserRole{UserId: u.Id, RoleId: roleId}
		database.Create(&ur)
	}
}

func Down_1() {

}