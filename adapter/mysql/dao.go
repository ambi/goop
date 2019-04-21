package mysql

import (
	"context"
	"database/sql"
	"log"

	"github.com/ambi/goop/domain/db"
	"github.com/ambi/goop/domain/model"
	_ "github.com/go-sql-driver/mysql" // MySQL
	"github.com/google/uuid"
)

const (
	insertAuthorizationCode = `INSERT INTO authorization_codes(uuid, user_uuid, code) VALUES(?, ?, ?)`
	deleteAuthorizationCode = `DELETE FROM authorization_codes WHERE uuid = ?`
	selectAuthorizationCode = `SELECT authorization_codes.uuid, users.uuid, users.login_id, users.email
					FROM authorization_codes, users
					WHERE authorization_codes.code = ? AND authorization_codes.user_uuid = users.uuid`

	insertClient = `INSERT INTO clients(uuid, client_id, client_secret, name) VALUES(?, ?, ?, ?)`
	deleteClient = `DELETE FROM clients WHERE client_id = ?`
	selectClient = `SELECT uuid, name, client_secret FROM clients WHERE client_id = ?`

	insertRedirectURI = `INSERT INTO redirect_uris(uuid, client_uuid, uri) VALUES(?, ?, ?)`
	deleteRedirectURI = `DELETE FROM redirect_uris WHERE uuid = ?`
	selectRedirectURI = `SELECT uri FROM redirect_uris WHERE client_uuid = ?`

	insertUser = `INSERT INTO users(uuid, login_id, email) VALUES(?, ?, ?)`
	deleteUser = `DELETE FROM users WHERE uuid = ?`
	selectUser = `SELECT uuid, email FROM users WHERE login_id = ?`
)

type sqlDAO struct {
	db *sql.DB

	insertAuthorizationCode *sql.Stmt
	deleteAuthorizationCode *sql.Stmt
	selectAuthorizationCode *sql.Stmt

	insertClient *sql.Stmt
	deleteClient *sql.Stmt
	selectClient *sql.Stmt

	insertRedirectURI *sql.Stmt
	deleteRedirectURI *sql.Stmt
	selectRedirectURI *sql.Stmt

	insertUser *sql.Stmt
	deleteUser *sql.Stmt
	selectUser *sql.Stmt
}

// NewDAO creates a new MySQL DB DAO.
func NewDAO(dataSource string) (db.DAO, error) {
	sqldb, err := sql.Open("mysql", dataSource)
	if err != nil {
		log.Fatal(err)
	}

	dao := &sqlDAO{db: sqldb}

	if dao.insertAuthorizationCode, err = sqldb.Prepare(insertAuthorizationCode); err != nil {
		return nil, err
	}
	if dao.deleteAuthorizationCode, err = sqldb.Prepare(deleteAuthorizationCode); err != nil {
		return nil, err
	}
	if dao.selectAuthorizationCode, err = sqldb.Prepare(selectAuthorizationCode); err != nil {
		return nil, err
	}
	if dao.insertClient, err = sqldb.Prepare(insertClient); err != nil {
		return nil, err
	}
	if dao.deleteClient, err = sqldb.Prepare(deleteClient); err != nil {
		return nil, err
	}
	if dao.selectClient, err = sqldb.Prepare(selectClient); err != nil {
		return nil, err
	}
	if dao.insertRedirectURI, err = sqldb.Prepare(insertRedirectURI); err != nil {
		return nil, err
	}
	if dao.deleteRedirectURI, err = sqldb.Prepare(deleteRedirectURI); err != nil {
		return nil, err
	}
	if dao.selectRedirectURI, err = sqldb.Prepare(selectRedirectURI); err != nil {
		return nil, err
	}
	if dao.insertUser, err = sqldb.Prepare(insertUser); err != nil {
		return nil, err
	}
	if dao.deleteUser, err = sqldb.Prepare(deleteUser); err != nil {
		return nil, err
	}
	if dao.selectUser, err = sqldb.Prepare(selectUser); err != nil {
		return nil, err
	}

	return dao, nil
}

func (dao *sqlDAO) Close() error {
	return dao.db.Close()
}

func (dao *sqlDAO) CreateAuthorizationCode(ctx context.Context, user *model.User) (*model.AuthorizationCode, error) {
	id := uuid.New().String()
	code := uuid.New().String()
	_, err := dao.insertAuthorizationCode.ExecContext(ctx, id, user.UUID, code)
	if err != nil {
		return nil, err
	}

	authzCode := &model.AuthorizationCode{
		UUID: id,
		Code: code,
		User: user,
	}
	return authzCode, nil
}

func (dao *sqlDAO) GetAuthorizationCode(ctx context.Context, code string) (*model.AuthorizationCode, error) {
	row := dao.selectAuthorizationCode.QueryRowContext(ctx, code)

	var uuid, userUUID, loginID, email string
	err := row.Scan(&uuid, &userUUID, &loginID, &email)
	if err != nil {
		return nil, err
	}

	authzCode := &model.AuthorizationCode{
		UUID: uuid,
		Code: code,
		User: &model.User{
			UUID:    userUUID,
			LoginID: loginID,
			Email:   email,
		},
	}
	return authzCode, nil
}

func (dao *sqlDAO) CreateClient(ctx context.Context, name string) (*model.Client, error) {
	// TODO
	return nil, db.ErrNotSaved
}

func (dao *sqlDAO) GetClient(ctx context.Context, clientID string) (*model.Client, error) {
	row := dao.selectClient.QueryRowContext(ctx, clientID)
	var uuid, name, clientSecret string
	err := row.Scan(&uuid, &name, &clientSecret)
	if err != nil {
		return nil, err
	}

	var redirectURIs []string
	rows, err := dao.selectRedirectURI.QueryContext(ctx, uuid)
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		var uri string
		err := rows.Scan(&uri)
		if err != nil {
			return nil, err
		}
		redirectURIs = append(redirectURIs, uri)
	}

	client := &model.Client{
		UUID:         uuid,
		ClientID:     clientID,
		ClientSecret: clientSecret,
		RedirectURIs: redirectURIs,
	}

	return client, nil
}

func (dao *sqlDAO) CreateUser(ctx context.Context, loginID string) (*model.User, error) {
	// TODO
	return nil, db.ErrNotSaved
}

func (dao *sqlDAO) GetUser(ctx context.Context, loginID string) (*model.User, error) {
	row := dao.selectUser.QueryRowContext(ctx, loginID)

	var uuid, email string
	err := row.Scan(&uuid, &email)
	if err != nil {
		return nil, err
	}

	user := &model.User{
		UUID:    uuid,
		LoginID: loginID,
		Email:   email,
	}
	return user, nil
}
