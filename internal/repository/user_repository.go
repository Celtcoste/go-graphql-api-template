package repository

import (
	"database/sql"
	"errors"
	"strings"
	"time"

	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
)

const (
	_insertUserQuery = `
		INSERT INTO user (uuid, pseudo, description, firebase_id, image_uuid, email_verified, cgu_accepted, provider, created_at, updated_at)
		VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?);`
	_insertFavoriteUserQuery = `
		INSERT INTO user_favorite_user (uuid, favorite_user_uuid, user_uuid, created_at, updated_at)
		VALUES (?, ?, ?, ?, ?);`
	_selectAllUserUUIDsQuery = `
		SELECT uuid
		FROM user
		ORDER BY created_at DESC;`
	_selectUserByUUIDQuery = `
		SELECT uuid, pseudo, description, image_uuid, notification_token, firebase_id, email_verified, cgu_accepted, provider, created_at, updated_at
		FROM user
		WHERE uuid = ?
		LIMIT 1;`
	_selectUserByPseudoQuery = `
		SELECT uuid, pseudo, description, image_uuid, notification_token, firebase_id, email_verified, cgu_accepted, provider, created_at, updated_at
		FROM user
		WHERE pseudo = ?
		LIMIT 1;`
	_selectUserByFirebaseIDQuery = `
		SELECT uuid, pseudo, description, image_uuid, notification_token, firebase_id, email_verified, cgu_accepted, provider, created_at, updated_at
		FROM user
		WHERE firebase_id = ?
		LIMIT 1;`
	_selectUsersByUUIDsQuery = `
		SELECT uuid, pseudo, description, image_uuid, notification_token, firebase_id, email_verified, cgu_accepted, provider, created_at, updated_at
		FROM user
		WHERE uuid IN (?);`
	_selectFavoriteUsersByUserUUIDQuery = `
		SELECT ufu.favorite_user_uuid 
		FROM user_favorite_user ufu
		WHERE ufu.user_uuid = ?
		ORDER BY ufu.created_at DESC;`
)

// UserRepositoryInterface should be implemented by UserRepository
type UserRepositoryInterface interface {
	Create(pseudo string, description string) (User, error)
	DeleteByUUID(uuid string) (bool, error)
	GetByUUID(uuid string) (User, error)
	GetByPseudo(pseudo string) (User, error)
	GetByFirebaseID(firebaseID string) (User, error)
	GetByUUIDs(uuids []string) ([]User, error)
	GetAllUUIDs() ([]string, error)
	UpdateByUUID(uuid string, pseudo *string, description *string,
		imageUUID *string, emailVerified *bool,
		cguAccepted *bool, provider *string, notificationToken *string) (bool, error)
	CheckIfPseudoAlreadyTaken(pseudo string) error
	GetFavoriteUsersByUserUUID(userUUID string) ([]string, error)
	FavUserByUUID(UUID string, userUUID string) (bool, error)
	UnfavUserByUUID(UUID string, userUUID string) (bool, error)
}

// User struct reflects database user table.
type User struct {
	UUID              string    `db:"uuid" json:"uuid"`
	Pseudo            string    `db:"pseudo" json:"pseudo"`
	Description       string    `db:"description" json:"description"`
	ImageUUID         *string   `db:"image_uuid" json:"imageUUID"`
	NotificationToken *string   `db:"notification_token" json:"notificationToken"`
	FirebaseID        *string   `db:"firebase_id" json:"firebaseID"`
	EmailVerified     bool      `db:"email_verified" json:"emailVerified"`
	CGUAccepted       bool      `db:"cgu_accepted" json:"cguAccepted"`
	Provider          string    `db:"provider" json:"provider"`
	CreationDate      time.Time `db:"created_at" json:"creationDate"`
	UpdateDate        time.Time `db:"updated_at" json:"updateDate"`
}

// UserRepository handle user data access
type UserRepository struct {
	DB *sqlx.DB
}

// NewUserRepository instantiate a new UserRepository
func NewUserRepository(db *sqlx.DB) *UserRepository {
	return &UserRepository{
		DB: db,
	}
}

// Create an user entry in the database.
func (r *UserRepository) Create(
	pseudo string,
	description string,
) (User, error) {
	id := uuid.New()
	now := time.Now().UTC()
	var err error
	_, err = r.DB.Exec(
		_insertUserQuery,
		id,
		pseudo,
		description,
		nil,
		false,
		true,
		now,
		now,
	)
	user := User{
		UUID:          id.String(),
		Pseudo:        pseudo,
		Description:   description,
		EmailVerified: false,
		CGUAccepted:   true,
		CreationDate:  now,
		UpdateDate:    now,
	}
	if err != nil {
		return user, err
	}
	return user, nil
}

// DeleteByUUID delete the user denoted by the given identifier.
func (r *UserRepository) DeleteByUUID(uuid string) (bool, error) {
	return DeleteByUUID(r.DB, "user", uuid)
}

// GetByUUID retrieve one User by its uuid from the database
func (r *UserRepository) GetByUUID(uuid string) (User, error) {
	user := User{}
	err := r.DB.Get(&user, _selectUserByUUIDQuery, uuid)
	return user, err
}

// GetByPseudo retrieve one User by its pseudo from the database
func (r *UserRepository) GetByPseudo(pseudo string) (User, error) {
	user := User{}
	err := r.DB.Get(&user, _selectUserByPseudoQuery, pseudo)
	return user, err
}

// GetByFirebaseID retrieve one User by its firebaseID from the database
func (r *UserRepository) GetByFirebaseID(firebaseID string) (User, error) {
	user := User{}
	err := r.DB.Get(&user, _selectUserByFirebaseIDQuery, firebaseID)
	return user, err
}

// GetByUUIDs retrieve multiple Users by given uuids from the database
func (r *UserRepository) GetByUUIDs(uuids []string) ([]User, error) {
	query, args, err := sqlx.In(_selectUsersByUUIDsQuery, uuids)
	if err != nil {
		return nil, err
	}
	var users []User
	err = r.DB.Select(&users, query, args...)
	return users, err
}

// GetAllUUIDs retrieve all User UUIDs from the database
func (r *UserRepository) GetAllUUIDs() ([]string, error) {
	var userUUIDs []string
	err := r.DB.Select(&userUUIDs, _selectAllUserUUIDsQuery)
	return userUUIDs, err
}

// UpdateByUUID update the user denoted by the given identifier.
func (r *UserRepository) UpdateByUUID(uuid string, pseudo *string, description *string,
	imageUUID *string, emailVerified *bool, cguAccepted *bool,
	provider *string, notificationToken *string) (bool, error) {
	fields := make(map[string]interface{})

	if pseudo != nil {
		fields["pseudo"] = pseudo
	}

	if description != nil {
		fields["description"] = description
	}

	if emailVerified != nil {
		fields["email_verified"] = emailVerified
	}

	if cguAccepted != nil {
		fields["cgu_accepted"] = cguAccepted
	}

	if provider != nil {
		fields["provider"] = strings.ToLower(*provider)
	}

	if notificationToken != nil {
		fields["notification_token"] = *notificationToken
	}
	if imageUUID != nil {
		fields["image_uuid"] = *imageUUID
	}

	fields["updated_at"] = time.Now().UTC()
	return UpdateByUUID(r.DB, "user", uuid, fields)
}

// CheckIfPseudoAlreadyTaken retrieve one User by its id from the database
func (r *UserRepository) CheckIfPseudoAlreadyTaken(pseudo string) error {
	user := User{}
	err := r.DB.Get(&user, _selectUserByPseudoQuery, pseudo)
	if err == sql.ErrNoRows {
		return nil
	}
	if user.UUID != "" {
		return errors.New("pseudo already taken")
	}
	return err
}

func (r *UserRepository) GetFavoriteUsersByUserUUID(userUUID string) ([]string, error) {
	var userUUIDs []string
	query, args, err := sqlx.In(_selectFavoriteUsersByUserUUIDQuery, userUUID)
	if err != nil {
		return nil, err
	}
	err = r.DB.Select(&userUUIDs, query, args...)
	return userUUIDs, err
}

func (r *UserRepository) FavUserByUUID(UUID string, userUUID string) (bool, error) {
	id := uuid.New()
	now := time.Now().UTC()
	_, err := r.DB.Exec(
		_insertFavoriteUserQuery,
		id,
		UUID,
		userUUID,
		now,
		now,
	)
	if err != nil {
		return false, err
	}
	return true, nil
}

func (r *UserRepository) UnfavUserByUUID(UUID string, userUUID string) (bool, error) {
	return Delete(
		r.DB,
		"user_favorite_user",
		"favorite_user_uuid = ? AND user_uuid = ?",
		UUID,
		userUUID,
	)
}
