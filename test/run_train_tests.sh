#!/usr/bin/env bash

for i in {1..3}; do
    docker run -v $(pwd):/var/loadtest --net host -it --rm direvius/yandex-tank -c load/load_$i.ini > log.tmp
    grep -m 1 "Web link" log.tmp | cut -d" " -f5
    grep -m 1 "Artifacts dir" log.tmp | tr -d "\r" | cut -d" " -f5 | cut -d"/" -f4 -f5 | xargs ls -1 | grep "answ_"
done

#rm -f log.tmp
