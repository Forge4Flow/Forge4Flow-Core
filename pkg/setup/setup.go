package setup

import (
	"context"
	"log"

	permission "github.com/auth4flow/auth4flow-core/pkg/authz/permission"
	role "github.com/auth4flow/auth4flow-core/pkg/authz/role"
	user "github.com/auth4flow/auth4flow-core/pkg/authz/user"
	warrant "github.com/auth4flow/auth4flow-core/pkg/authz/warrant"
	"github.com/auth4flow/auth4flow-core/pkg/config"
)

func InitialSetup(cfg *config.Auth4FlowConfig, permissionSvc *permission.PermissionService, roleSvc *role.RoleService, userSvc *user.UserService, warrantSvc *warrant.WarrantService) {
	ctx := context.Background()
	permRoleName := "auth4flow-admin"

	// Check if the auth4flow-admin role & Permission exist and create if needed
	_, err := roleSvc.GetByRoleId(ctx, permRoleName)
	if err != nil {
		newRoleSpec := role.RoleSpec{
			RoleId: permRoleName,
			Name:   &permRoleName,
		}
		_, err := roleSvc.Create(ctx, newRoleSpec)
		if err != nil {
			log.Fatalln("Unable to create admin role")
		}
	}

	_, err = permissionSvc.GetByPermissionId(ctx, "auth4flow-admin")
	if err != nil {
		newPermissionSpec := permission.PermissionSpec{
			PermissionId: permRoleName,
			Name:         &permRoleName,
		}

		_, err := permissionSvc.Create(ctx, newPermissionSpec)
		if err != nil {
			log.Fatalln("Unable to create admin permission")
		}
	}

	newWarrantSpec := warrant.WarrantSpec{
		ObjectType: "permission",
		ObjectId:   permRoleName,
		Relation:   "member",
		Subject: &warrant.SubjectSpec{
			ObjectType: "role",
			ObjectId:   permRoleName,
		},
	}
	_, err = warrantSvc.Create(ctx, newWarrantSpec)
	if err != nil {
		log.Fatalln("Unable to assign permission role to role")
	}

	// Verify the admin user exists and create if needed
	_, err = userSvc.GetByUserId(ctx, cfg.AdminAccount)
	if err != nil {
		newUserSpec := user.UserSpec{
			UserId: cfg.AdminAccount,
		}

		_, err = userSvc.Create(ctx, newUserSpec)
		if err != nil {
			log.Fatalln("Unable to create admin role")
		}
	}

	// Add role to the user
	newWarrantSpecUser := warrant.WarrantSpec{
		ObjectType: "role",
		ObjectId:   permRoleName,
		Relation:   "member",
		Subject: &warrant.SubjectSpec{
			ObjectType: "user",
			ObjectId:   cfg.AdminAccount,
		},
	}

	_, err = warrantSvc.Create(ctx, newWarrantSpecUser)
	if err != nil {
		log.Fatalln("Unable to assign admin role to user")
	}
}
