#!/bin/bash
source config.env

sleep 2 && goose -dir "${MIGRATION_DIR}" postgres "${MIGRATION_DSN}" up -v