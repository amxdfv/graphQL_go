# graphQL_go

Создадим учебный сервис, который будет хранить в себе информацию о пользователях и их покупках. Соответственно, у нас будет три связанных SQL-таблицы.
Одна с пользователями, вторая с транзакциями покупок, третья с товарами. Для получения всего этого через HTTP будет реаливан graphQl-интерфейс.
С помощью него мы сможем получать информацию о пользователях, транзакциях и продуктах, а также добавлять эти сущности.

Тезисно, что планируется реализовать:
1. База данных SQL Lite  с тремя таблицами: для пользователей, транзакций и товаров.
2. Три graphQL query для получения пользователей, транзакций и товаров соответственно.
3. Три graphQL mutation для добавления пользователей, транзакций и товаров соответственно.
4. Обработку ошибок интегрированную с системой логирования. Будем писать логи в формате json в отдельный файл. 


Таблицы в базе данных:
USERS - для пользователей
id - CHAR уникальное
name -  CHAR
age - NUMBER
address - CHAR
document_type - NUMBER
document - NUMBER уникальное

TRANSACTIONS - для транзакций
id - CHAR уникальное
rrn - NUMBER уникальное
amount - NUMBER
currency - CHAR
used_id - CHAR связь по ключу с users.id
good_id - CHAR связь по ключу с goods.id
place - CHAR
t_time - DATETIME

GOODS - для товаров
id - CHAR уникальное
name -  CHAR
price - NUMBER
currency - CHAR
country_origin - CHAR

Для запуска в контейнере:
1. Ставим себе docker 
2. Собираем образ командой docker build -t  graphql .
3. Запускаем контенейр командой docker run -p 8080:8080 graphql
4. Приложение будет доступно по адресу http://localhost:8080/oslic
