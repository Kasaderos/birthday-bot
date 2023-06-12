# BIRTHDAY-TGBOT

## Quick start

Test local

Set your tgbot token in .env (BOT_TOKEN) 

Terminal 1

```
    source .env
    docker run -d --rm -e TZ=Asia/Almaty -e POSTGRES_PASSWORD=passwd -p 5432:5432 --name postgres  postgres
    cd test
    go test -v .
```

Terminal 2

```
    cd test/tgbot
    go run .
```

You will see a message like 'Happy birthday, Nate!'

## Local setup

Before using set your own config in .env

```
    source .env
    make build
    make run
```



