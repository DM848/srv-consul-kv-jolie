#!/bin/sh

echo "starting web server"
export WEB_SERVER_PORT=80
nohup ./webserver > my.log 2> my.log < /dev/null &
PID=$!

echo "wait one second.."
sleep 1

if kill -s 0 $PID
then
    echo "webserver running"
else
    exit 1
fi

echo "running jolie test"
if jolie jolie-test.ol | grep -q 'success'
then
    echo "success!"
else
    cat my.log
    exit 2
fi
kill -9 $PID