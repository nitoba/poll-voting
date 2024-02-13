package test

import (
	"strings"

	"github.com/google/uuid"
	configs "github.com/nitoba/poll-voting/config"
	"github.com/nitoba/poll-voting/internal/infra/database/prisma"
)

var newSchema = "schema=" + uuid.New().String()

func generateUniqueDatabaseURL() string {
	conf := configs.GetConfig()
	if conf == nil || conf.DATABASE_URL == "" {
		panic("DATABASE_URL is not set")
	}
	// Generate Unique Database URL
	newSchema := "schema=" + uuid.New().String()
	return strings.Replace(conf.DATABASE_URL, "schema=public", newSchema, 1)
}

func BeforeAll() {
	configs.LoadConfig(".env.test")
	// newUrl := generateUniqueDatabaseURL()
	// os.Setenv("DATABASE_URL", newUrl)
	// exec.Command("make", "prisma-deploy").Run()
}

func AfterAll() {
	conf := configs.GetConfig()
	db := prisma.GetDB()
	db.Prisma.ExecuteRaw("DROP SCHEMA IF EXISTS " + newSchema + " CASCADE").Exec(conf.Ctx)
	prisma.Disconnect()
}
