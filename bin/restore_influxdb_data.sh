#!/bin/bash

docker exec -it defi_stats_influxdb_1 mkdir db_backup
docker cp db/influx_backup defi_stats_influxdb_1:/db_backup/
docker exec -it defi_stats_influxdb_1 influx restore /db_backup/influx_backup
