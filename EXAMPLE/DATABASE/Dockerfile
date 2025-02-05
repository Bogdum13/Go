# Используем официальный образ PostgreSQL
FROM postgres:latest

# Копируем SQL-скрипт в контейнер
COPY init.sql /docker-entrypoint-initdb.d/

# Устанавливаем переменные окружения для PostgreSQL
ENV POSTGRES_USER=postgres
ENV POSTGRES_PASSWORD=BogdanBogdum
ENV POSTGRES_DB=app_db