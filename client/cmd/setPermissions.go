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
	"fmt"

	"github.com/spf13/cobra"
)

var sUserName string
var sGroupName string

// setPermissionsCmd represents the setPermissions command
var setPermissionsCmd = &cobra.Command{
	Use:   "permissions",
	Short: "Set permissions",
	Long:  `Get permissions from dshare server`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("setPermissions called")
		// TODO: Check if permission and reverse permission exists --> swap read to write
		// TODO: Get ID from user and group
		// TODO: Update permission with post request --> content in body --> encode
		// TODO: if permission not found, print message with hint to use add
	},
}

func init() {
	setCmd.AddCommand(setPermissionsCmd)
	getPermissionsCmd.Flags().StringVarP(&sUserName, "user", "u", "", "Name of a user")
	getPermissionsCmd.Flags().StringVarP(&sGroupName, "group", "g", "", "Name of a group")
}
