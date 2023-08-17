# Sample web app written in Go

### Bootstraping the infra

```
docker-compose up -d
```

### Preparing the database with mocked data
Connect to the database with your preferred client and perform the following query to create the product table

```
create table product (
    id serial primary key,
    name varchar,
    description varchar, 
    price decimal, 
    amount integer
) 
```
Populate the table
```
insert into product (name, description, price, amount) values
('T-shirt', 'Modern shirt', 29, 10),
('Notebook', 'Enterprise edition', 4500, 7)
('Keyboard', 'Mechanical keyboard', 700, 57)
```
### Starting the application
```
go run main.go
```