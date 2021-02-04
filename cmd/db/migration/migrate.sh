user=$MYSQL_DATABASE_DB1_USERNAME
password=$MYSQL_DATABASE_DB1_PASSWORD
host=$MYSQL_DATABASE_DB1_HOST_WRITE
port=$MYSQL_DATABASE_DB1_PORT
database=$MYSQL_DATABASE_DB1_NAME

migrate -path ./ -database mysql://$user:$password@tcp\($host:$port\)/$database up

# migrate -path ./ -database mysql://root:123@tcp\(localhost:3306\)/test up