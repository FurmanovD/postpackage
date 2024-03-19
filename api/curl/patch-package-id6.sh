#! /bin/sh
curl --location --request PATCH 'http://127.0.0.1:8080/api/v1/packages' \
--header 'Content-Type: application/json' \
--data-raw '    {
        "id": 6,
        "name": "EvenMoreXXXL",
        "items_per_package": 20002
    }'