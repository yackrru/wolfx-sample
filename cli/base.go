package cli

import (
	"database/sql"
	_ "github.com/lib/pq"
	"github.com/yackrru/wolfx"
	"github.com/yackrru/wolfx/middleware"
	"os"
)

func Execute(jobName string) int {
	wx := wolfx.New()

	db, err := sql.Open("postgres",
		"host=127.0.0.1 port=5432 user=postgres password=postgres dbname=postgres sslmode=disable")
	if err != nil {
		middleware.Logger.Error(err)
		os.Exit(1)
	}

	wx.Add(NewDBToFileJob(db))

	if err := wx.Run(jobName); err != nil {
		middleware.Logger.Error(err)
		os.Exit(1)
	}

	return 0
}
