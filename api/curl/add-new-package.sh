#! /bin/sh
curl --location --request POST 'http://127.0.0.1:8080/api/v1/packages' \
--header 'Content-Type: application/json' \
--data-raw '    {
        "name": "XXXL",
        "items_per_package": 10000
    }'