/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package main

import "backend/cmd"

//	@title			Playcode API
//	@version		1.0
//	@description	This is a simple REST API server for playing go code.

//	@contact.name	Ganeshdip Dumbare
//	@contact.email	ganeshdip.dumbare@gmail.com

//	@license.name	Apache 2.0
//	@license.url	http://www.apache.org/licenses/LICENSE-2.0.html

// @host		localhost:5174
// @BasePath	/api/v1
func main() {
	cmd.Execute()
}
