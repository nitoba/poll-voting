package test

import (
	"fmt"
	"os"
	"os/exec"
	"strings"

	"github.com/google/uuid"
	configs "github.com/nitoba/poll-voting/config"
	"github.com/nitoba/poll-voting/internal/infra/database/prisma"
)

var newSchemaID = uuid.New().String()

func generateUniqueDatabaseURL() string {
	conf := configs.GetConfig()
	if conf == nil || conf.DATABASE_URL == "" {
		panic("DATABASE_URL is not set")
	}
	// Generate Unique Database URL
	newSchema := "schema=" + newSchemaID
	return strings.Replace(conf.DATABASE_URL, "schema=public", newSchema, 1)
}

func SetupDatabase() {
	configs.LoadConfig(".env.test")
	newUrl := generateUniqueDatabaseURL()

	// Deploy Database
	println("Deploying database with url: ", newUrl)

	os.Setenv("DATABASE_URL", newUrl)
	cmd := exec.Command("make", "prisma-deploy")
	cmd.Dir = configs.RootDir()
	err := cmd.Run()

	if err != nil {
		println("Error to deploy database: ", err.Error())
	}

	prisma.Connect()
}

func TruncateTables() {
	dba := prisma.GetDB()
	query := `TRUNCATE TABLE "schema".voters, "schema".polls, "schema".poll_options, "schema".votes CASCADE`
	query = strings.ReplaceAll(query, "schema", newSchemaID)
	println("Truncating tables: ", query)
	dba.Prisma.ExecuteRaw(query).Exec(configs.GetConfig().Ctx)
}

func AfterAll() {
	conf := configs.GetConfig()
	dba := prisma.GetDB()
	query := fmt.Sprintf(`DROP SCHEMA IF EXISTS "%s" CASCADE`, newSchemaID)
	println("Dropping schema: ", query)
	_, err := dba.Prisma.ExecuteRaw(query).Exec(conf.Ctx)
	if err != nil {
		println("Error to drop schema: ", err.Error())
	}
	prisma.Disconnect()
}
