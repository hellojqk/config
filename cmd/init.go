/*
Copyright © 2021 NAME HERE <EMAIL ADDRESS>

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
package cmd

import (
	"context"
	"fmt"

	"github.com/hellojqk/config/server/entity"
	"github.com/hellojqk/config/server/service"
	"github.com/hellojqk/config/util"
	"github.com/spf13/cobra"
)

// initCmd represents the init command
var initCmd = &cobra.Command{
	Use:     "init",
	Aliases: []string{"initialize", "initialise", "create"},
	Short:   "初始化",
	Long:    `初始化管理员账号等`,
	Run: func(cmd *cobra.Command, args []string) {
		util.WaitInitFuncsExec()
		err := service.RoleInsertOne(context.Background(), entity.Role{Key: "admin", Name: "管理员", GroupPower: map[string]map[string]interface{}{"*": {
			"*": true,
		}}})
		if err != nil {
			fmt.Println("role:admin 创建失败", err.Error())
			return
		}

		err = service.UserInsertOne(context.Background(), entity.User{Key: "admin", Name: "超级管理员", RoleKeys: []string{"admin"}, Password: "123123"})
		if err != nil {
			fmt.Println("user:admin 创建失败", err.Error())
			return
		}
		fmt.Println("超级管理员初始化成功")

		err = service.GroupInsertOne(context.Background(), entity.Group{Key: "default", Name: "默认分组"})
		if err != nil {
			fmt.Println("group:default 创建失败", err.Error())
			return
		}
	},
}

func init() {
	rootCmd.AddCommand(initCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// initCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// initCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
