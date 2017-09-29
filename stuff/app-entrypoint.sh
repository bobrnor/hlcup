#!/bin/bash

#set -e

find /var/lib/mysql -type f -exec touch {} \;

shopt -s nullglob dotglob
files=(/var/lib/mysql/*)
if [ ${#files[@]} -eq 0 ]; then
    mysql_install_db
fi

/usr/sbin/mysqld &
PID=$!
sleep 5
mysql -uroot -e "DROP DATABASE IF EXISTS hlcup; CREATE DATABASE hlcup;"
cat /hlcup.sql | mysql -uroot hlcup
kill -TERM $PID
wait $PID || true

unzip -o -d /tmp/data /tmp/data/data.zip

exec "$@"
