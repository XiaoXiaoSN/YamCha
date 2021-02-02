package db_test

import (
	"context"
	"testing"
	"time"

	"yamcha/pkg/model"
	"yamcha/pkg/repository"
	"yamcha/pkg/repository/db"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/suite"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// Test_Repository define testing case for database repository
func Test_Repository(t *testing.T) {
	suite.Run(t, new(repoTestSuite))
}

type repoTestSuite struct {
	suite.Suite

	sqlMock sqlmock.Sqlmock
	repo    repository.Repository
}

func (suite *repoTestSuite) SetupSuite() {
	// open database stub
	sqlDB, mock, err := sqlmock.New()
	suite.Require().NoError(err)

	// open gorm DB
	dial := mysql.New(mysql.Config{
		Conn:                      sqlDB,
		SkipInitializeWithVersion: true,
	})
	gormDB, err := gorm.Open(dial, &gorm.Config{})
	suite.Require().NoError(err)

	// enable debug mode
	gormDB = gormDB.Debug()

	suite.sqlMock = mock
	suite.repo = db.NewRepo(gormDB)
}

/*******************************
 * User Repository
 *******************************/

func (suite *repoTestSuite) TestCreateUser() {
	ctx := context.Background()

	// prepare mock
	var nextID int64 = 2
	suite.sqlMock.ExpectExec("INSERT INTO `users`").
		WillReturnResult(sqlmock.NewResult(nextID, 1))

	// excute function
	var resource = model.User{
		Name:      "hello",
		LineID:    "lineID",
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
	err := suite.repo.CreateUser(ctx, &resource)
	suite.Require().NoError(err)
	suite.Require().Equal(resource.ID, int(nextID))

	// we make sure that all expectations were met
	err = suite.sqlMock.ExpectationsWereMet()
	suite.Require().NoError(err)
}

func (suite *repoTestSuite) TestUserList() {
	ctx := context.Background()

	// prepare mock
	var exceptResources = []model.User{
		{ID: 1, Name: "first", LineID: "lineID1"},
		{ID: 2, Name: "second", LineID: "lineID2"},
	}
	rows := sqlmock.NewRows([]string{"id", "name", "line_id"})
	for _, u := range exceptResources {
		rows.AddRow(u.ID, u.Name, u.LineID)
	}

	suite.sqlMock.ExpectQuery("SELECT (.+) FROM `users`").
		WillReturnRows(rows)

	// excute function
	users, err := suite.repo.UserList(ctx)
	suite.Require().NoError(err)
	for _, u := range exceptResources {
		suite.Require().Contains(users, u)
	}

	// we make sure that all expectations were met
	err = suite.sqlMock.ExpectationsWereMet()
	suite.Require().NoError(err)
}

/*******************************
 * Store Repository
 *******************************/

func (suite *repoTestSuite) TestCreateStore() {
	ctx := context.Background()

	// prepare mock
	var nextID int64 = 2
	suite.sqlMock.ExpectExec("INSERT INTO `stores`").
		WillReturnResult(sqlmock.NewResult(nextID, 1))

	// excute function
	var resource = model.Store{
		GroupName: "mock group",
	}
	err := suite.repo.CreateStore(ctx, &resource)
	suite.Require().NoError(err)
	suite.Require().Equal(resource.ID, int(nextID))

	// we make sure that all expectations were met
	err = suite.sqlMock.ExpectationsWereMet()
	suite.Require().NoError(err)
}

func (suite *repoTestSuite) TestGetStore() {
	ctx := context.Background()

	// prepare mock
	var resource = model.Store{
		ID:        12,
		GroupName: "SS",
		BranchStores: []model.BranchStore{
			{ID: 1, Name: "SS-1", StoreGroupID: 12},
			{ID: 2, Name: "SS-2", StoreGroupID: 12},
		},
	}
	// mock store
	row := sqlmock.NewRows([]string{"id", "group_name"}).
		AddRow(resource.ID, resource.GroupName)
	suite.sqlMock.ExpectQuery("SELECT (.+) FROM `stores` WHERE id = ?(.*)").
		WithArgs(resource.ID).
		WillReturnRows(row)

	// mock branch stores
	rows := sqlmock.NewRows([]string{"id", "name", "store_group_id"})
	for _, r := range resource.BranchStores {
		rows.AddRow(r.ID, r.Name, r.StoreGroupID)
	}
	suite.sqlMock.ExpectQuery("SELECT (.+) FROM `branch_stores` WHERE `branch_stores`.`store_group_id` = ?").
		WithArgs(resource.ID).
		WillReturnRows(rows)

	// excute function
	store, err := suite.repo.GetStore(ctx, resource.ID)
	suite.Require().NoError(err)
	suite.Require().Exactly(resource, store)

	// we make sure that all expectations were met
	err = suite.sqlMock.ExpectationsWereMet()
	suite.Require().NoError(err)
}

func (suite *repoTestSuite) TestStoreList() {
	ctx := context.Background()

	// prepare mock
	var exceptResources = []model.Store{
		{ID: 1, GroupName: "first"},
		{ID: 2, GroupName: "second"},
	}
	rows := sqlmock.NewRows([]string{"id", "group_name"})
	for _, u := range exceptResources {
		rows.AddRow(u.ID, u.GroupName)
	}

	suite.sqlMock.ExpectQuery("SELECT (.+) FROM `stores`").
		WillReturnRows(rows)

	// excute function
	resourceList, err := suite.repo.StoreList(ctx)
	suite.Require().NoError(err)
	for _, r := range exceptResources {
		suite.Require().Contains(resourceList, r)
	}

	// we make sure that all expectations were met
	err = suite.sqlMock.ExpectationsWereMet()
	suite.Require().NoError(err)
}

func (suite *repoTestSuite) TestBranchStoreList() {
	ctx := context.Background()

	// prepare mock
	var storeID = 1
	var exceptResources = []model.BranchStore{
		{ID: 1, Name: "first", StoreGroupID: storeID},
		{ID: 2, Name: "second", StoreGroupID: storeID},
	}
	rows := sqlmock.NewRows([]string{"id", "name", "store_group_id"})
	for _, r := range exceptResources {
		rows.AddRow(r.ID, r.Name, r.StoreGroupID)
	}
	suite.sqlMock.ExpectQuery("SELECT (.+) FROM `branch_stores`").
		WillReturnRows(rows)

	// excute function
	resourceList, err := suite.repo.BranchStoreList(ctx, storeID)
	suite.Require().NoError(err)
	for _, r := range exceptResources {
		suite.Require().Contains(resourceList, r)
	}

	// we make sure that all expectations were met
	err = suite.sqlMock.ExpectationsWereMet()
	suite.Require().NoError(err)
}

func (suite *repoTestSuite) TestCreateBranchStore() {
	ctx := context.Background()

	// prepare mock
	var nextID int64 = 2
	suite.sqlMock.ExpectExec("INSERT INTO `branch_stores`").
		WillReturnResult(sqlmock.NewResult(nextID, 1))

	// excute function
	var resource = model.BranchStore{
		Name:         "mock",
		StoreGroupID: 1,
		Address:      "address",
	}
	err := suite.repo.CreateBranchStore(ctx, &resource)
	suite.Require().NoError(err)
	suite.Require().Equal(resource.ID, int(nextID))

	// we make sure that all expectations were met
	err = suite.sqlMock.ExpectationsWereMet()
	suite.Require().NoError(err)
}
