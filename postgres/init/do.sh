#!/bin/bash

echo "......execute /dumps/db-schema.sql"
psql -v ON_ERROR_STOP=1 --username "postgres" dbpostgres < /dumps/db-schema.sql

sleep 1

for i in `ls /dumps/db-plus-*.sql | sort -V`; 
do
    echo "......${i}"
    psql -v ON_ERROR_STOP=1 --username "postgres" dbpostgres < ${i}
done;
