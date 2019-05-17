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
	"log"

	"github.com/dmittelstaedt/dpwdss/client/logic"
	"github.com/spf13/cobra"
)

var name string

// usersCmd represents the users command
var usersCmd = &cobra.Command{
	Use:   "users",
	Short: "Gets users.",
	Long:  `Gets users from dshare server.`,
	Run: func(cmd *cobra.Command, args []string) {
		if cmd.Flags().Changed("name") {
			name, err := cmd.Flags().GetString("name")
			if err != nil {
				log.Println(err)
				return
			}
			user, ok := logic.ReadUserByName(name)
			if ok {
				logic.PrintUser(user)
			}
		} else {
			users := logic.ReadUsers()
			logic.PrintUsers(users)
		}
	},
}

func init() {
	getCmd.AddCommand(usersCmd)
	usersCmd.Flags().StringVarP(&name, "name", "n", "", "Name of a user")
}
