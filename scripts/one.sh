#!/bin/bash
echo "Hello World!"
youhao="runoob.com"

echo ${youhao}
#for skill in Ada Coffe Action Java; do
#    echo "I am good at ${skill}Script"
#done

exe_path=$(cd `dirname $0`;pwd)
SERVER_HOME=${exe_path}/../
echo ${exe_path}
echo ${SERVER_HOME}data/go.pid
if [ -f ${SERVER_PID} ]
then
    pid=`cat ${SERVER_PID}`
    echo pid
else
    echo 2
fi
#echo $0
#echo $1
#echo $#
#echo $*
#echo $$

#val=`expr 2 + 2`
#echo "两数之和为 : $val"