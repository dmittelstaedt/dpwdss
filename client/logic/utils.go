package logic

import "github.com/dmittelstaedt/dpwdss/client/models"

// ConvertPermissions converts a slice of Permission to a slice of PermissionOut
func ConvertPermissions(permission *models.Permission) models.PermissionOut {
	group := ReadGroup(permission.GroupID)
	user := ReadUser(permission.UserID)
	var permissionOut = models.PermissionOut{
		ID:        permission.ID,
		UserName:  user.Name,
		GroupName: group.Name,
	}
	return permissionOut
}
