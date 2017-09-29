#!/usr/bin/env bash

for i in {1..1}; do
    docker run -v $(pwd):/var/loadtest --net host -i --rm direvius/yandex-tank -c load/load_$i.ini > log.tmp
    url=$(grep -m 1 "Web link" log.tmp | cut -d" " -f5)
    log_path=$(grep -m 1 "Artifacts dir" log.tmp | tr -d "\r" | cut -d" " -f5 | cut -d"/" -f4 -f5)
    log_file=$(grep -m 1 "Artifacts dir" log.tmp | tr -d "\r" | cut -d" " -f5 | cut -d"/" -f4 -f5 | xargs ls -1 | grep "answ_")
    log="$log_path/$log_file"

    python ./check/prepare.py log results/tmp
    echo url > results/result_$1.log
    echo results/tmp >> results/result_$1.log
    rm -f results/tmp
done
