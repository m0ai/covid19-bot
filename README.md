# covid19-bot

- Scrapper  
	- [X] Scraping a Covid 19 Information from Covid19 Open API
	- [x] Save a Data to Database 
	    - Daily number of confirmed case
	    - Update Datetime
- Slack Bot (cronjob)
  - [x] today's decide count to notify slack channel
- REST API

  - [ ] interface for interactive slack bot 



## How to Debug 

```bash
# Slack Notify App
reflex -r '\.go' -s -- sh -c "go run src/main.go" 

# Scrapper App
reflex -r '\.go' -s -- sh -c "go run src/scrapper.go" 
```




## How to Deploy

```bash
make deploy-[dev|pord]
```

