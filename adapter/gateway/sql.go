package gateway

import (
	"context"
	"crypto/sha1"
	"database/sql"
	"time"

	"github.com/ambi/goop/domain/db"
	"github.com/ambi/goop/domain/model"
	"github.com/google/uuid"
)

// SQL is an implementation of Repository.
type SQL struct {
	// SQL
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

	insertRevocation *sql.Stmt
	deleteRevocation *sql.Stmt
	selectRevocation *sql.Stmt
}

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

	insertRevocation = `INSERT INTO revocations(uuid, token, hashed_token, expires_at) VALUES(?, ?, ?, ?)`
	deleteRevocation = `DELETE FROM revocations WHERE uuid = ?`
	selectRevocation = `SELECT uuid, token FROM revocations WHERE hashed_token = ?`
)

// NewSQL creates an SQL implementation.
func NewSQL(db *sql.DB) (*SQL, error) {
	var err error
	s := &SQL{db: db}

	if s.insertAuthorizationCode, err = db.Prepare(insertAuthorizationCode); err != nil {
		return nil, err
	}
	if s.deleteAuthorizationCode, err = db.Prepare(deleteAuthorizationCode); err != nil {
		return nil, err
	}
	if s.selectAuthorizationCode, err = db.Prepare(selectAuthorizationCode); err != nil {
		return nil, err
	}
	if s.insertClient, err = db.Prepare(insertClient); err != nil {
		return nil, err
	}
	if s.deleteClient, err = db.Prepare(deleteClient); err != nil {
		return nil, err
	}
	if s.selectClient, err = db.Prepare(selectClient); err != nil {
		return nil, err
	}
	if s.insertRedirectURI, err = db.Prepare(insertRedirectURI); err != nil {
		return nil, err
	}
	if s.deleteRedirectURI, err = db.Prepare(deleteRedirectURI); err != nil {
		return nil, err
	}
	if s.selectRedirectURI, err = db.Prepare(selectRedirectURI); err != nil {
		return nil, err
	}
	if s.insertUser, err = db.Prepare(insertUser); err != nil {
		return nil, err
	}
	if s.deleteUser, err = db.Prepare(deleteUser); err != nil {
		return nil, err
	}
	if s.selectUser, err = db.Prepare(selectUser); err != nil {
		return nil, err
	}
	// if s.insertRevocation, err = db.Prepare(insertRevocation); err != nil {
	// 	return nil, err
	// }
	// if s.deleteRevocation, err = db.Prepare(deleteRevocation); err != nil {
	// 	return nil, err
	// }
	// if s.selectRevocation, err = db.Prepare(selectRevocation); err != nil {
	// 	return nil, err
	// }

	return s, nil
}

// Close closes DB.
func (s *SQL) Close() error {
	return s.db.Close()
}

// CreateAuthorizationCode creates an authorization code.
func (s *SQL) CreateAuthorizationCode(ctx context.Context, user *model.User) (*model.AuthorizationCode, error) {
	id := uuid.New().String()
	code := uuid.New().String()
	_, err := s.insertAuthorizationCode.ExecContext(ctx, id, user.UUID, code)
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

// GetAuthorizationCode gets an authorization code.
func (s *SQL) GetAuthorizationCode(ctx context.Context, code string) (*model.AuthorizationCode, error) {
	row := s.selectAuthorizationCode.QueryRowContext(ctx, code)

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

// CreateClient create a client.
func (s *SQL) CreateClient(ctx context.Context, name string) (*model.Client, error) {
	// TODO
	return nil, db.ErrNotSaved
}

// GetClient gets a client.
func (s *SQL) GetClient(ctx context.Context, clientID string) (*model.Client, error) {
	row := s.selectClient.QueryRowContext(ctx, clientID)
	var uuid, name, clientSecret string
	err := row.Scan(&uuid, &name, &clientSecret)
	if err != nil {
		return nil, err
	}

	var redirectURIs []string
	rows, err := s.selectRedirectURI.QueryContext(ctx, uuid)
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

// CreateUser creates a user.
func (s *SQL) CreateUser(ctx context.Context, loginID string) (*model.User, error) {
	// TODO
	return nil, db.ErrNotSaved
}

// GetUser gets a user.
func (s *SQL) GetUser(ctx context.Context, loginID string) (*model.User, error) {
	row := s.selectUser.QueryRowContext(ctx, loginID)

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

// CreateRevocation creates a revocation.
func (s *SQL) CreateRevocation(ctx context.Context, token string, expiresAt time.Time) error {
	id := uuid.New().String()
	hash := calculateHashedToken(token)
	_, err := s.insertRevocation.ExecContext(ctx, id, token, hash, expiresAt)
	return err
}

// GetRevocation gets a revocation.
func (s *SQL) GetRevocation(ctx context.Context, token string) (bool, error) {
	hash := calculateHashedToken(token)
	rows, err := s.selectRevocation.QueryContext(ctx, hash)
	if err != nil {
		return false, err
	}

	for rows.Next() {
		var uuid, t string
		err := rows.Scan(&uuid, &t)
		if err != nil {
			return false, err
		}
		if t == token {
			rows.Close()
			return true, nil
		}
	}

	return false, nil
}

func calculateHashedToken(token string) string {
	hash := sha1.Sum([]byte(token))
	return string(hash[:])
}
