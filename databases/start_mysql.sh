#!/usr/bin/env zsh

docker run \
  --name veigar-mysql \
   -p 3307:3306 \
   -e MYSQL_DATABASE=veigar \
   -e MYSQL_ALLOW_EMPTY_PASSWORD=yes \
   -e MYSQL_ROOT_PASSWORD=veigar \
   -d mysql:5.7 --lower_case_table_names=1