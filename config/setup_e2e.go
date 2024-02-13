package configs

import (
	"os"
	"os/exec"
	"strings"

	"github.com/google/uuid"
	"github.com/nitoba/poll-voting/internal/infra/database/prisma"
)

var newSchema = "schema=" + uuid.New().String()

func generateUniqueDatabaseURL() string {
	if config == nil || config.DATABASE_URL == "" {
		panic("DATABASE_URL is not set")
	}
	// Generate Unique Database URL
	newSchema := "schema=" + uuid.New().String()
	return strings.Replace(config.DATABASE_URL, "schema=public", newSchema, 1)
}

func BeforeAll() {
	newUrl := generateUniqueDatabaseURL()
	os.Setenv("DATABASE_URL", newUrl)
	exec.Command("make", "prisma-deploy").Run()
}

func AfterAll() {
	db := prisma.GetDB()
	db.Prisma.ExecuteRaw("DROP SCHEMA IF EXISTS " + newSchema + " CASCADE").Exec(config.Ctx)
	prisma.Disconnect()
}
