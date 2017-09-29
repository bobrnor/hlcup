#!/usr/bin/env bash

docker run -v /Users/bobrnor/Development/Pets/hlcupdocs:/var/loadtest --net host -it --rm direvius/yandex-tank -c load/load_1.ini
