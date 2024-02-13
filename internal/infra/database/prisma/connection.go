package prisma

import (
	"os"

	configs "github.com/nitoba/poll-voting/config"
	"github.com/nitoba/poll-voting/prisma/db"
)

var prisma *db.PrismaClient

func Connect() error {
	logger := configs.GetLogger("prisma")
	logger.Info("connecting with postgres")
	url := os.Getenv("DATABASE_URL")
	client := db.NewClient(db.WithDatasourceURL(url))
	if err := client.Prisma.Connect(); err != nil {
		return err
	}

	prisma = client
	return nil
}

func Disconnect() {
	if prisma == nil {
		return
	}

	logger := configs.GetLogger("prisma")
	logger.Info("disconnecting with postgres")
	prisma.Prisma.Disconnect()
}

func GetDB() *db.PrismaClient {
	return prisma
}
