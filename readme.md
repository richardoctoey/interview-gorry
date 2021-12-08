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
dbhost=gorry-richard-oey.db
dbport=3306
dbname=gorry
host=
port=9999
```
-----
You can import Postman Collection from gorry.postman_collection.json

First you need to run:
1. Hit api Create Location. Copy the uuid from this api response
2. Hit api Create Event. Paste the result from step 1 into `location` field. Copy the uuid from this api response
3. Hit api Create Ticket. Paste the event uuid from step 2. Copy uuid from this api response
4. Hit api Purchase Ticket. Paste the ticket uuid from step 3.
5. Transaction Detail and Get Event may have uuid query parameter, like this:

```
http://localhost:9999/transaction/get_info?uuid=xxx
http://localhost:9999/event/get_info?uuid=xxx
```