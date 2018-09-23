#!/bin/bash

#shell脚本执行目录
exe_path="$( cd "$( dirname "$0"  )" && pwd  )"
echo ${exe_path}
SERVER_HOME=${exe_path}/../
SERVER_BIN=${SERVER_HOME}/bin/
SERVER_PID=${SERVER_HOME}/data/go.pid
EXE_NAME=xx

start()
{
    if [ -f ${SERVER_PID} ]
    then
        pid=`cat ${SERVER_PID}`
        process_num=`ps -ef | grep -w ${pid} | grep -v "grep" | grep ${EXE_NAME} | wc -l`
        if [ ${SERVER_PID} -ge 1 ];
        then
            echo "service already running. pid=$pid"
            return
        fi  
    fi
    cd ${SERVER_PID}
    nohup ./${EXE_NAME} &> /dev/null 2>> ../logs/${EXE_NAME}_except.log &
    
    echo "${EXE_NAME} start"
}

stop()
{
    if [ ! -f ${SERVER_PID} ]
    then
        echo "service already exit"
        return
    fi
    pid=`cat ${SERVER_PID}`
    process_num=`ps -ef | grep -w ${pid} | grep -v "grep" | grep ${EXE_NAME} | wc -l`
    if [ ${process_num} -eq 0 ];
    then
        echo "service already exit"
        return
    fi 
    kill -TERM ${pid}
    ret=$?
    if [ ${ret} -eq 0 ]
    then
        echo "${EXE_NAME} stopped"
    else
        echo "${EXE_NAME} stop failed"
    fi
    return
}

restart()
{
    stop
    start
    return
}

reload()
{
    if [ ! -f ${SERVER_PID} ]
    then
        echo "service already exit"
        return
    fi
    pid=`cat ${SERVER_PID}`
    process_num=`ps -ef | grep -w $pid | grep -v "grep" | grep ${EXE_NAME} | wc -l`
    if [ ${process_num} -eq 0 ];
    then
        echo "service already exit"
        return
    fi 
    kill -USR2 `cat ${SERVER_PID}`
    return
}

case "$1" in
    start)
        start
        ;;
    stop)
        stop
        ;;
    restart)
        restart
        ;;
    reload)
        reload
        ;;
    *)
        echo $"Usage: $0 {start|stop|restart|reload}"
        exit 1
esac

exit 0
