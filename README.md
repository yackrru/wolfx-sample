# WolfX sample codes

[![CI](https://github.com/yackrru/wolfx-sample/actions/workflows/ci.yml/badge.svg)](https://github.com/yackrru/wolfx-sample/actions/workflows/ci.yml)
[![CodeQL](https://github.com/yackrru/wolfx-sample/actions/workflows/codeql-analysis.yml/badge.svg)](https://github.com/yackrru/wolfx-sample/actions/workflows/codeql-analysis.yml)

## Prepare
Launch postgres container with database=postgres,username=postgres,password=postgres and port 5432 is mapped on host PC.
```
docker compose up -d
```
Execute DDL to postgres.
```
PGPASSWORD=postgres psql -h localhost -p 5432 -U postgres postgres -e < testdata/ddl.sql
```
Execute DML to create sample data.
```
PGPASSWORD=postgres psql -h localhost -p 5432 -U postgres postgres -e < testdata/insert_sample_data.sql
```

## Execution
Go build and execution.
```
-- build
go build -o batch

-- run format: ./batch -job {jobName}
./batch -job DBToFile
```

| jobName  | description                                                         |
|:---------|:--------------------------------------------------------------------|
| DBToFile | Load the data from postgres database and output them to a tsv file. |
