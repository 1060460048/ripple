package scripts

import (
	"github.com/bmbstack/ripple"
	_ "github.com/bmbstack/ripple/fixtures/forum/controllers"
	"github.com/bmbstack/ripple/fixtures/forum/logger"
	_ "github.com/bmbstack/ripple/fixtures/forum/models"
)

// Server commands
func GetServerCommands(db string) []string {
	commands := make([]string, 2)
	switch db {
	case "mysql":
		commands = append(commands, "/usr/local/bin/mysqld_safe &")
		commands = append(commands, "sleep 10s")
	}
	return commands
}

// Run server
func RunServer() {
	logger.Logger.Info("Run server ....")

	ripple.Run()
}
