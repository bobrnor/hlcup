#!/usr/bin/env bash

curl -s -o /dev/null -w "%{http_code} %{url_effective}\n" 0.0.0.0:8080/users/1
curl -s -o /dev/null -w "%{http_code} %{url_effective}\n" 0.0.0.0:8080/users/string
curl -s -o /dev/null -w "%{http_code} %{url_effective}\n" 0.0.0.0:8080/users/string/somethingbad
curl -s -o /dev/null -w "%{http_code} %{url_effective}\n" 0.0.0.0:8080/users/
curl -s -o /dev/null -w "%{http_code} %{url_effective}\n" 0.0.0.0:8080/user/
curl -s -o /dev/null --data '{"birth_date": 47779200, "first_name": "\u0410\u0440\u043a\u0430\u0434\u0438\u0439"}' -w "POST: %{http_code} %{url_effective}\n" 0.0.0.0:8080/users/8230
curl -s -o /dev/null --data '{"last_name": "\u0421\u0442\u044b\u043a\u044b\u043a\u0430\u0442\u0438\u043d"}' -w "POST: %{http_code} %{url_effective}\n" 0.0.0.0:8080/users/6051
curl -s -o /dev/null --data '{"first_name": "\u041a\u0438\u0440\u0438\u043b\u043b", "last_name": "\u041b\u0435\u0431\u044b\u043a\u0430\u0432\u0435\u043d", "gender": "m", "id": 10058, "birth_date": -478569600, "email": "ohrisnaunhos@list.me"}' -w "POST: %{http_code} %{url_effective}\n" 0.0.0.0:8080/users/new

curl -s -o /dev/null -w "%{http_code} %{url_effective}\n" 0.0.0.0:8080/locations/1
curl -s -o /dev/null -w "%{http_code} %{url_effective}\n" 0.0.0.0:8080/locations/string
curl -s -o /dev/null -w "%{http_code} %{url_effective}\n" 0.0.0.0:8080/locations/string/somethingbad
curl -s -o /dev/null -w "%{http_code} %{url_effective}\n" 0.0.0.0:8080/locations/
curl -s -o /dev/null -w "%{http_code} %{url_effective}\n" 0.0.0.0:8080/location/
curl -s -o /dev/null -w "%{http_code} %{url_effective}\n" 0.0.0.0:8080/locations/1/avg
curl -s -o /dev/null -w "%{http_code} %{url_effective}\n" 0.0.0.0:8080/locations/somethingsomething/avg
curl -s -o /dev/null --data '{"place": "\u0420\u0443\u0447\u0435\u0439"}' -w "POST: %{http_code} %{url_effective}\n" 0.0.0.0:8080/locations/77
curl -s -o /dev/null --data '{"distance": 16}' -w "POST: %{http_code} %{url_effective}\n" 0.0.0.0:8080/locations/5630
curl -s -o /dev/null --data '{"id": 7629, "distance": 98, "place": "\u041e\u0437\u0435\u0440\u043e", "city": "\u0411\u0430\u0440\u0441\u0433\u0430\u043c\u0430", "country": "\u041c\u0430\u043b\u044c\u0442\u0430"}' -w "POST: %{http_code} %{url_effective}\n" 0.0.0.0:8080/locations/new

curl -s -o /dev/null -w "%{http_code} %{url_effective}\n" 0.0.0.0:8080/visits/1
curl -s -o /dev/null -w "%{http_code} %{url_effective}\n" 0.0.0.0:8080/visits/string
curl -s -o /dev/null -w "%{http_code} %{url_effective}\n" 0.0.0.0:8080/visits/string/somethingbad
curl -s -o /dev/null -w "%{http_code} %{url_effective}\n" 0.0.0.0:8080/visits/
curl -s -o /dev/null -w "%{http_code} %{url_effective}\n" 0.0.0.0:8080/visit/
curl -s -o /dev/null -w "%{http_code} %{url_effective}\n" 0.0.0.0:8080/users/1/visits
curl -s -o /dev/null -w "%{http_code} %{url_effective}\n" 0.0.0.0:8080/users/1/visit
curl -s -o /dev/null -w "%{http_code} %{url_effective}\n" 0.0.0.0:8080/users/1/visits?fromDate=
curl -s -o /dev/null -w "%{http_code} %{url_effective}\n" 0.0.0.0:8080/users/1/visits?fromDate=abracadbra
curl -s -o /dev/null -w "%{http_code} %{url_effective}\n" 0.0.0.0:8080/users/somethingstringhere/visits?fromDate=1
curl -s -o /dev/null -w "%{http_code} %{url_effective}\n" 0.0.0.0:8080/users/1/visit?fromDate=915148800&toDate=915148800
curl -s -o /dev/null --data '{"user": 7541, "mark": 1}' -w "POST: %{http_code} %{url_effective}\n" 0.0.0.0:8080/visits/58721
curl -s -o /dev/null --data '{"visited_at": 953275464, "user": 5508, "mark": 3}' -w "POST: %{http_code} %{url_effective}\n" 0.0.0.0:8080/visits/6856
curl -s -o /dev/null --data '{"id": 100574, "user": 3626, "visited_at": 1398085947, "location": 4273, "mark": 1}' -w "POST: %{http_code} %{url_effective}\n" 0.0.0.0:8080/visits/new