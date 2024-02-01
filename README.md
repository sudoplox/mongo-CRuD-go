# mongo-CRuD-go

Currently works with Mongo Atlas, can work with local mongo setup also

### To start the service
```
go run main.go
```

### If connection is okay this will be printed

> Pinged your deployment. You successfully connected to MongoDB!


### To exit and disconnect from mongo, do a SIGINT (cmd+c) 
> ^C2024/02/01 19:27:02 Program killed !

## To create
curl:
```
curl --location 'http://localhost:9000/user/' \
--header 'Content-Type: application/json' \
--data '{
    "Name": "sudhanshu",
    "Gender": "male",
    "Age": 24
}'
```
output:
```
Inserted Id: e�����tg�`a  
Inserted Obj: {
    "id": "65bba791e1ed7405678a6061",
    "name": "Sudhanshu",
    "gender": "male",
    "age": 24
}
```

## To read
curl:
```
curl --location 'http://localhost:9000/user/65bba791e1ed7405678a6061'
```
output:
```
{
    "id": "65bba791e1ed7405678a6061",
    "name": "Sudhanshu",
    "gender": "male",
    "age": 24
}
```

## To delete
curl:
```
curl --location --request DELETE 'http://localhost:9000/user/65bba791e1ed7405678a6061'
```
output:
```
Deleted user: {
    "id": "65bba791e1ed7405678a6061",
    "name": "Sudhanshu",
    "gender": "male",
    "age": 24
}
```



## miscellaneous commands

### to start local mongo server
```
brew services start mongodb-community@7.0
```

### to stop local mongo server
```
brew services stop mongodb-community@7.0
```
### to start local mongo server using conf
```
mongod --config /usr/local/etc/mongod.conf --fork
```
### client for local mongo server
```
mongosh
```
### To verify that MongoDB is running
```
brew services list
```

### If you started MongoDB manually as a background process:
```
ps aux | grep -v grep | grep mongod
```
