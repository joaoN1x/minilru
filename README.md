# MiniLru


MiniLru is an url shortener dev infrastructure solution, it's for learning porposes, it's not production ready. This solution in production would retain the data of the RDBMS and recover information every instance/machine reboot.


### Tech

MiniLru uses a number of open source projects to work nicely:

* [Docker] - Best containerized solution
* [Docker Compose] - Great for local development
* [Redis] - Cache key value solution.
* [Postgresql] - Relationa Database solution
* [Golang] - Micro services oriented programming language

### Installation

MiniLru requires latest [Docker](https://www.docker.com) and [Docker Compose](https://docs.docker.com/compose/) to run.

After having this repo, get inside it and run the following

```sh
$ docker-compose up --build
```


### API usage

To get a health check message do the following
```sh
$ curl -X GET "http://localhost:8050/"
```
Will return a JSON message
```json
{"status":"OK","code":200,"data":{}}
```


To create a short URL, do the following
```sh
$ curl -X POST "http://localhost:8050/url/" -H 'Accept: application/json' -d '{"long":"http://thisisurl.com/index.html"}'
```
Will return a JSON message
```json
{"status":"Added","code":200,"data":{"detail":"LJ0rE","affected":1}}
```
At data.detail, in this example, LJ0rE represents the short code for the url, it would be used in the format http(s)://<somedomain>/LJ0rE


To get the last 24 hours stats, e.g. for LJ0rE
```sh
$ curl -X GET "http://localhost:8050/stats/day/LJ0rE"
```
Will return a JSON message, e.g. 11 sum count
```json
{"status":"OK","code":200,"data":{"detail":"11"}}
```


To get the last last week stats, e.g. for LJ0rE
```sh
$ curl -X GET "http://localhost:8050/stats/week/LJ0rE"
```
Will return a JSON message, e.g. 222 sum count
```json
{"status":"OK","code":200,"data":{"detail":"222"}}
```


To get the all time stats, e.g. for LJ0rE
```sh
$ curl -X GET "http://localhost:8050/stats/all/LJ0rE"
```
Will return a JSON message, e.g. 333333 sum count
```json
{"status":"OK","code":200,"data":{"detail":"333333"}}
```



### Short Url usage

Knowing the short code for the Url, "LJ0rE" like above examples, to test it out run on the browser
http://localhost:8081/LJ0rE



### Check ms processing speed of redirect

At the console, where all docker containers are running and shooting logs, check as an example, as shown below an "info" type log, to see how fast the redirect func was processed
```json
{
    "service":"minilru",
    "type":"info",
    "file":"/app/src/restful/url.go",
    "line":"54",
    "function":"github.com/joaoN1x/minilru/src/restful.duration",
    "message":"ShortURL LJ0rE done in 26.014µs",
    "error":"",
    "timestamp":"2020-02-10T12:08:15.349340844Z"}
```



### Help Tools

After running Docker compose, the following tools are available for better understanding of data flow.

**Rebrow** Redis Webview
http://localhost:5001/
Host: redis-dev
Port: 6379
Database Id: 0

**Adminer** Postgresql Webview
http://localhost:8088/
System: PostgreSQL
Server: database-postgres-dev
Username: postgres
Password: postgresqL
Database: dbpostgres



#### Tips
When port conflicts arise, just change Dockerfile(s) port's.



