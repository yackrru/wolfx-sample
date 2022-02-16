package cli

import (
	"context"
	"database/sql"
	"encoding/csv"
	"github.com/yackrru/wolfx"
	"github.com/yackrru/wolfx/integration/database"
	"github.com/yackrru/wolfx/integration/file"
	"github.com/yackrru/wolfx/middleware"
	"os"
)

type DBToFileJob struct {
	db *sql.DB
}

func NewDBToFileJob(db *sql.DB) *DBToFileJob {
	return &DBToFileJob{
		db: db,
	}
}

func (j *DBToFileJob) Name() string {
	return "DBToFile"
}

func (j *DBToFileJob) Run() error {
	return wolfx.NewJobBuilder().
		Single(j.LoadDBAndOutputFileStep).
		Build()
}

func (j *DBToFileJob) LoadDBAndOutputFileStep(ctx context.Context) error {
	reader := database.NewReader(&database.ReaderConfig{
		DB:        j.db,
		SQL:       "select id, name, division, joined_at from users order by id desc",
		ChunkSize: 10,
	})

	currentDir, err := os.Getwd()
	if err != nil {
		return err
	}
	outputFilePath := currentDir + "/output/users.tsv"
	middleware.Logger.Infof("Create the file: %s\n", outputFilePath)
	output, err := os.Create(outputFilePath)
	if err != nil {
		middleware.Logger.Errorf("Failed to create the file: %s\n", outputFilePath)
		return err
	}
	defer output.Close()
	csvWriter := csv.NewWriter(output)
	csvWriter.Comma = '\t'

	propsBindPosition := make(middleware.PropsBindPosition)
	propsBindPosition["id"] = 0
	propsBindPosition["division"] = 1
	propsBindPosition["name"] = 2
	propsBindPosition["joined_at"] = 3

	writer := file.NewWriter(&file.WriterConfig{
		Writer:            csvWriter,
		PropsBindPosition: propsBindPosition,
	})

	return wolfx.NewStepBuilder(ctx).
		SetReader(reader).
		SetWriter(writer).
		Build()
}
