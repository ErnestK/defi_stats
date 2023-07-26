#!/bin/bash

docker exec -it defi_stats_influxdb_1 influx backup influx_backup
docker cp defi_stats_influxdb_1:/influx_backup db
