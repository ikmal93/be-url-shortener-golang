# urlshortenerikmal
This project built from golang progamming language and mysql database to create url shortener

Please follow this instruction to use this program

1. Make sure you have installed golang on your computer and MySQL driver for golang, see this https://github.com/go-sql-driver/mysql

2. Create new database 'db_redirects' in your MySQL and use 'db_redirects'

3. Create new table 'redirect' for database 'db_redirects', you can use this query:

create table redirect (
	id int auto_increment primary key,
    slug varchar(255),
    url varchar(255)
);

4. You can clone or download this repository in your local file, then run it. make sure you have installed this go library in your computer 
  - go get -u github.com/gorilla/mux => Gorilla router 
  - go get -u github.com/go-sql-driver/mysql => mysql driver

5. Use these 2 endpoint you will use for this program
  * http://localhost:12345/create 
      => FormData 
          key : url
          value : 'the_link_you_want_to_short'
  * http://localhost:12345/{slug_variable}

or you can import this link into postman https://www.getpostman.com/collections/c82fcfc544647ac07def

Happy Testing :)
  
  
