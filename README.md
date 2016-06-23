# ttapi

#summary
This's a RESTful api by golang programming language.  

#usage
1.Please make sure you have a working golang environment & set up GOPATH is completed  
2."cd" in your GOPATH directory & get drives.  
    like this, go get github.com/gorilla/mux & go get gopkg.in/pg.v4  
3.Put the code GOPATH/src directory  
4.Enter the following code into your console  
    go build -o ttapi  
5.Modify app.conf in ttapi/conf  
    [default]  
    appname= //apiname  
    httpport= //apiport  
    [postgresql]  
    postgresqlhost = //host:port  
    postgresqldbname = //dbname  
    postgresqluser = //user  
    postgresqlpwd = //pwd  
6.Use information_schema.sql build DB in your PostgreSQL server  
7.Now, you can start ttapi on your server by "nohup ./ttapi &"  

#Router
1.get all users  
    host:port/users -- [GET]  
2.create user  
    host:port/users -- [POST]  parameter {"name":?}  
3.  create relationship  
    host:port/users/{user_id}/relationships/{other_user_id} -- [PUT] parameter {"state": "liked|unliked"}  
4.  get relationships by user_id  
    host:port/users/{user_id}/relationships -- [GET]   
