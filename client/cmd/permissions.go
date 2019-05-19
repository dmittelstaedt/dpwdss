// Copyright © 2019 David Mittelstädt <mittelstaedt.david@gmail.com>
//
// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:
//
// The above copyright notice and this permission notice shall be included in
// all copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
// THE SOFTWARE.

package cmd

import (
	"github.com/dmittelstaedt/dpwdss/client/logic"
	"github.com/dmittelstaedt/dpwdss/client/models"
	"github.com/spf13/cobra"
)

var userName string
var groupName string

// permissionsCmd represents the permissions command
var permissionsCmd = &cobra.Command{
	Use:   "permissions",
	Short: "Get permissions",
	Long:  `Get permissions from dshare server`,
	Run: func(cmd *cobra.Command, args []string) {
		var permissionsOut []models.PermissionOut
		permissions := logic.ReadPermissions()
		for _, permission := range permissions {
			user := logic.ReadUser(permission.UserID)
			var permissionOut = models.PermissionOut{
				ID:       permission.ID,
				UserName: user.Name,
			}
			permissionsOut = append(permissionsOut, permissionOut)
		}
		logic.PrintPermissions(permissionsOut)
	},
}

func init() {
	getCmd.AddCommand(permissionsCmd)
	permissionsCmd.Flags().StringVarP(&userName, "user", "u", "", "Name of a user")
	permissionsCmd.Flags().StringVarP(&groupName, "group", "g", "", "Name of a group")
}
