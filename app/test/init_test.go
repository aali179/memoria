package test

import (
	"fmt"
	"memoria/app/controllers"
	"memoria/app/models"
	"testing"

	"github.com/stretchr/testify/suite"
	"gorm.io/gorm"

	"github.com/gin-gonic/gin"
)

type TestSuiteEnv struct {
	suite.Suite
	db *gorm.DB
}

// Tests are run before they start
func (suite *TestSuiteEnv) SetupSuite() {
	SetupTest()
	suite.db = models.GetTestDB()
}

// Running after each test
func (suite *TestSuiteEnv) TearDownTest() {
	ClearTable(suite.db)
}

// Running after all tests are completed
func (suite *TestSuiteEnv) TearDownSuite() {
	//suite.db.Close()
}

// This gets run automatically by `go test`
func TestSuite(t *testing.T) {
	suite.Run(t, new(TestSuiteEnv))
}

//Setup routers to test db
func SetupTest() *gin.Engine {
	router := gin.Default()
	err := models.ConnectTestDB()
	api := &controllers.APIEnv{
		DB: models.GetTestDB(),
	}
	if err != nil {
		fmt.Println(err)
	}
	// Initialize all routes
	controllers.InitializeRoutes(router, api)

	return router
}

func (suite *TestSuiteEnv) insertTestScrapbook(db *gorm.DB) (models.Scrapbook, error) {
	sb := models.Scrapbook{
		Name: "Test Book",
	}

	if err := db.Create(&sb).Error; err != nil {
		return sb, err
	}

	return sb, nil
}

func ClearTable(db *gorm.DB) {
	db.Exec("DELETE FROM books")
	db.Exec("ALTER SEQUENCE books_id_seq RESTART WITH 1")
}
