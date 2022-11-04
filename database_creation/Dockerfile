FROM mysql
ENV MYSQL_ROOT_PASSWORD password
ENV MYSQL_DATABASE Players
ADD players_database.csv /var/lib/mysql-files/
ADD init.sql /docker-entrypoint-initdb.d
EXPOSE 3300 