SHELL_PATH = /bin/ash
SHELL = $(if $(wildcard $(SHELL_PATH)),/bin/ash,/bin/bash)

bld:
	go build -o build/db-backup-cli app/cli/main.go
run:
	go build -o build/db-backup-cli app/cli/main.go && ./build/db-backup-cli -t postgres -d db_name -u root -w 123456 -h postgres://postgres:password@localhost:5432/test_db?sslmode=disable -n example_table
db:
	docker compose up
