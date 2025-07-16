#!/usr/bin/bash

clear
gradle $@
while true; do
	inotifywait -r -m -e modify --exclude '/\..*' ./ | read line
	clear
	gradle $@
done
