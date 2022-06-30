pname="hh_tool_process"
function start(){
    pid=`ps -ef |grep -v grep | grep ${pname} | awk '{print $2}'`
    if test $pid;then
    echo "the programmer has started the pid is $pid"
    else
    ./$pname >>nohup.out &
    echo "the programmer start successfully"
    fi
}

function stop(){
    pid=`ps -ef |grep -v grep | grep ${pname} | awk '{print $2}'`
    if test $pid; then
    kill $pid
    echo "the programmer stop successfully the pid is $pid"
    else
    echo "the programmer has stopped"
    fi
}

case $1 in "start")
    start
    ;;
    "stop")
    stop
    ;;
esac