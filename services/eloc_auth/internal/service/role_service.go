package service

import (
	"context"

	db "github.com/ryannguyen1105/eloc-backend/services/eloc_auth/db/sqlc"
)

type RoleService struct {
	store *db.Store
}

func NewRoleService(store *db.Store) *RoleService {
	return &RoleService{store: store}
}

func (roleService *RoleService) CreateRole(ctx context.Context, name string) (db.Role, error) {
	arg := db.CreateRoleParams{
		Name: name,
	}
	return roleService.store.CreateRole(ctx, arg)
}

func (roleService *RoleService) GetRole(ctx context.Context, id int64) (db.Role, error) {
	arg := db.GetRoleParams{
		ID: id,
	}
	return roleService.store.GetRole(ctx, arg)
}

func (roleService *RoleService) ListRoles(ctx context.Context, pageSize int64, pageID int64) ([]db.Role, error) {
	arg := db.ListRolesParams{
		Limit:  int32(pageSize),
		Offset: int32(pageID-1) * int32(pageSize),
	}
	return roleService.store.ListRoles(ctx, arg)
}

func (roleService *RoleService) UpdateRoles(ctx context.Context, id int64, name string) (db.Role, error) {
	arg := db.UpdateRoleParams{
		ID:   id,
		Name: name,
	}
	return roleService.store.UpdateRole(ctx, arg)
}

func (roleService *RoleService) DeleteRole(ctx context.Context, id int64)  error {
	arg := db.DeleteRoleParams{
		ID: id,
	}
	return roleService.store.DeleteRole(ctx, arg)
}
