// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.20.0

package sqlc

import (
	"context"
)

type Querier interface {
	CreateClient(ctx context.Context, arg CreateClientParams) (Client, error)
	CreateProfile(ctx context.Context, arg CreateProfileParams) (Profile, error)
	CreateUser(ctx context.Context, arg CreateUserParams) (User, error)
	DeleteClient(ctx context.Context, id string) error
	DeleteProfile(ctx context.Context, id string) error
	DeleteUser(ctx context.Context, id string) error
	GetClient(ctx context.Context) (Client, error)
	GetProfile(ctx context.Context, id string) (Profile, error)
	GetProfileUserID(ctx context.Context, userID string) (Profile, error)
	GetUser(ctx context.Context, id string) (User, error)
	GetUserUsername(ctx context.Context, username string) (User, error)
	ListClients(ctx context.Context) ([]Client, error)
	ListProfiles(ctx context.Context, arg ListProfilesParams) ([]Profile, error)
	ListUsers(ctx context.Context, arg ListUsersParams) ([]User, error)
	UpdateProfile(ctx context.Context, arg UpdateProfileParams) (Profile, error)
	UpdateProfileUserID(ctx context.Context, arg UpdateProfileUserIDParams) (Profile, error)
	UpdateUserPassword(ctx context.Context, arg UpdateUserPasswordParams) (User, error)
	UpdateUserStatus(ctx context.Context, arg UpdateUserStatusParams) (User, error)
}

var _ Querier = (*Queries)(nil)
