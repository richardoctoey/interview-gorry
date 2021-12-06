Before you run the app these are the dependenncies you need to install:

1. Install docker
2. Install golang
3. Port 9999 needs to be free. Will use port: 9999 for the api

```make build```

then

```make run```

To stop

```make stop```

-----

This is the default Configuration, and located at configuration/config.env
(Should be running fine with default configuration)
```
dbuser=gorry
dbpass=gorry
dbhost=gorry.db
dbport=3306
dbname=gorry
host=
port=9999
```
-----
You can import Postman Collection from gorry.postman_collection.json