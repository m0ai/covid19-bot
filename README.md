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
# Frist, Start a Database
make db

# Slack Notify App
make watch-notify

# or Scrapper App
make watch-scrapper
```

## How to Deploy

```bash
make docker-build
make docker-push
make deploy-[dev|pord]
```

