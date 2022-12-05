#!/bin/sh/

curl -X POST localhost:3000/v1/wallets  \
-H "Accept: application/json" \
-H "Content-Type: application/json" \
-d '{"datetime": "2022-12-05T07:45:05+07:00","amount": 10}'

curl -X POST "localhost:3000/v1/wallets/history" \
-H "Accept: application/json" \
-H "Content-Type: application/json" \
-d '{"startDatetime":"2022-12-05T02:45:05+07:00","endDatetime":"2022-12-05T07:45:05+07:00"}'

curl -X POST localhost:3000/v1/wallets  \
-H "Accept: application/json" \
-H "Content-Type: application/json" \
-d '{"datetime": "2022-12-05T07:45:06+07:00","amount": 2}' 

curl -X POST localhost:3000/v1/wallets  \
-H "Accept: application/json" \
-H "Content-Type: application/json" \
-d '{"datetime": "2022-12-05T07:45:05+07:00","amount": 2}' 

curl -X POST "localhost:3000/v1/wallets/history" \
-H "Accept: application/json" \
-H "Content-Type: application/json" \
-d '{"startDatetime":"2022-12-05T02:45:05+07:00","endDatetime":"2022-12-05T07:45:05+07:00"}'

curl -X POST localhost:3000/v1/wallets  \
-H "Accept: application/json" \
-H "Content-Type: application/json" \
-d '{"datetime": "2022-12-05T07:45:05+07:00","amount": 10}'


curl -X POST localhost:3000/v1/wallets  \
-H "Accept: application/json" \
-H "Content-Type: application/json" \
-d '{"datetime": "2022-12-05T08:45:05+07:00","amount": 1.32}'

curl -X POST "localhost:3000/v1/wallets/history" \
-H "Accept: application/json" \
-H "Content-Type: application/json" \
-d '{"startDatetime":"2022-12-05T02:45:05+07:00","endDatetime":"2022-12-05T09:45:05+07:00"}'

## output file in clean database
# {"message":"data created successfully"}
# [{"datetime":"2022-12-05T00:00:00Z","amount":"10"}]
# {"message":"data created successfully"}
# {"message":"data created successfully"}
# [{"datetime":"2022-12-05T00:00:00Z","amount":"14"}]
# {"message":"data created successfully"}
# {"message":"data created successfully"}
# [{"datetime":"2022-12-05T00:00:00Z","amount":"24"},{"datetime":"2022-12-05T01:00:00Z","amount":"25.32"}]