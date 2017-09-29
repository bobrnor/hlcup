#!/usr/bin/env bash

mkdir -p /var/hlcup/data
rsync -rpog --delete-after ./data/TRAIN/data/* /var/hlcup/data
