# GoLog

This is just a practice Go program, written while following "Distributed Services with Go" book. 
This isn't copy/paste from the book, I modify and play around with the code after I look at the book's examples. 
The goal is to essentially practice Go while learning more about how to build distributed services with it.


# Running it
```bash
go run cmd/server/main.go


# Append a log entry
curl -XPOST http://localhost:8000 -d '{"data":"This is my log data"}' | jq
# Reply will be like '{"key":"cgzpjtam"}'

# Get a log entry by its key (randomly generated on append)
curl -XGET http://localhost:8000/cgzpjtam | jq

# Get all log entries
curl -XGET http://localhost:8000 | jq
```