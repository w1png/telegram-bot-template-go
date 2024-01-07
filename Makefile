build-docker:
	docker build -t telegram-bot-template-go .

run-docker:
	docker-compose up -d
