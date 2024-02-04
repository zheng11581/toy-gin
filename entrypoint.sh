#!/bin/bash
#### APM ####
set_apm(){
    JAVA_OPTS="$JAVA_OPTS -Djava.security.egd=file:/dev/./urandom -Duser.timezone=GMT+08 -Ddefault.client.encoding=UTF-8 -Dfile.encoding=UTF-8 -Dsun.jnu.encoding=UTF8 -Duser.language=Zh"
    if [[ ! -z $MESOS_TASK_ID ]]; then
        AGENT_ID=$(echo "$MESOS_TASK_ID"|cut -d'.' -f 2|cut -d'-' -f 1,2,3) #mesos container
    else
        AGENT_ID=$(echo "$HOSTNAME") #k8s container
    fi
    #==enable pinpoint or not==
    if [[ "X${developer_apm_appId}" != "X" && "X${AGENT_ID}" != "X" ]]; then
        #strip spaces or tabs before or behind the string and cut to left 24 strings
        APPLICATION_NAME=$(echo "$developer_apm_appId" | awk 'gsub(/^ *| *$/,"")'  | cut -b -24)
        PINPOINT_COLLECTOR_IP=${PINPOINT_COLLECTOR_IP:-10.3.15.29}
        AGENT_ID_SHORT=$(echo "$AGENT_ID" | cut -b -24)
        APM_JAR="/usr/local/pinpoint-agent/pinpoint-bootstrap-1.7.2.jar"
        JAVA_OPTS="$JAVA_OPTS -javaagent:${APM_JAR} -Dpinpoint.agentId=${AGENT_ID_SHORT} -Dpinpoint.applicationName=${APPLICATION_NAME}"
        
        #configure pinpoint.config
        PINPOINT_CONFIG_FILE="/usr/local/pinpoint-agent/pinpoint.config"
        sed -i -e "s/profiler.collector.ip=127.0.0.1/profiler.collector.ip=${PINPOINT_COLLECTOR_IP}/" ${PINPOINT_CONFIG_FILE}
        
        # Adjust pointpoint port.
        COLLECTOR_TCP_PORT=${PINPOINT_COLLECTOR_TCP_PORT:-9994}
        let COLLECTOR_STAT_PORT=${COLLECTOR_TCP_PORT}+1
        let COLLECTOR_SPAN_PORT=${COLLECTOR_TCP_PORT}+2        
        sed -i -e "s/profiler.collector.tcp.port=.*/profiler.collector.tcp.port=${COLLECTOR_TCP_PORT}/" ${PINPOINT_CONFIG_FILE}        
        sed -i -e "s/profiler.collector.stat.port=.*/profiler.collector.stat.port=${COLLECTOR_STAT_PORT}/" ${PINPOINT_CONFIG_FILE}       
        sed -i -e "s/profiler.collector.span.port=.*/profiler.collector.span.port=${COLLECTOR_SPAN_PORT}/" ${PINPOINT_CONFIG_FILE}
        
        # Configure log level for log4j
        PINPOINT_LOG_CONFIG_FILE="/usr/local/pinpoint-agent/lib/log4j.xml"
        PINPOINT_LOG_LEVEL=${PINPOINT_LOG_LEVEL:-ERROR}
        sed -i -e "s/ERROR/${PINPOINT_LOG_LEVEL}/g" ${PINPOINT_LOG_CONFIG_FILE}
        
        #profiler.sampling.rate：采样率（1/n，配置为2就是50%）
        SAMPLING_RATE=${SAMPLING_RATE:-20}
        if [[ $SAMPLING_RATE =~ ^-?[0-9]+$ ]]; then
            sed -i -e "s/profiler.sampling.rate=.*/profiler.sampling.rate=${SAMPLING_RATE}/" ${PINPOINT_CONFIG_FILE}
        else
            SAMPLING_RATE=20
            echo "WARNING: the value of SAMPLING_RATE is $SAMPLING_RATE, which need to be a int number, using defalut: 20 instead."
        fi
        SAMPLING_RATE_PERCENT=`awk 'BEGIN{printf "%.2f%%\n",(1/'${SAMPLING_RATE}')*100}'`

        #RPC log sampling rate
        if [[ "X$SPAN_TRACE_RATIO" != "X" && $SPAN_TRACE_RATIO =~ ^-?[0-9]+$ ]]; then
            JAVA_OPTS="$JAVA_OPTS -Dspan.trace.ratio=$SPAN_TRACE_RATIO"
        else
            SPAN_TRACE_RATIO="100%"
        fi
        
        echo "--------------------Pinpoint Enabled--------------------"
        echo "Pinpoint Application Name: $APPLICATION_NAME"
        echo "Pinpoint Agent Id:         $AGENT_ID"
        echo "Pinpoint Collector Ip:     $PINPOINT_COLLECTOR_IP  (can be overided by given ENV \$PINPOINT_COLLECTOR_IP)"
        echo "Pinpoint Log Level:        $PINPOINT_LOG_LEVEL     (can be overided by given ENV \$PINPOINT_LOG_LEVEL)"
        echo "Pinpoint Sampling.Rate:    $SAMPLING_RATE_PERCENT  (can be overided by given ENV \$SAMPLING_RATE, E.G.: set SAMPLING_RATE to 20, 1/n, then rate will be 5%)"
        echo "RPC log sampling rate:     $SPAN_TRACE_RATIO       (can be overided by given ENV \$SPAN_TRACE_RATIO, E.G.: set SAMPLING_RATE to 5, then rate will be 5%)"
        echo "---------------------------------------------------------"
    fi
    export JAVA_OPTS="$JAVA_OPTS"
}
#### confcenter ####
confcenter(){
    /usr/local/src/confdownload/confcenterdownload > /dev/null 2>&1
}

set_applog(){
    if [[ "X${yonyoucloud_replace_path}" != "X" && "X${yonyoucloud_replace_value}" != "X" ]]; then
        paths=(${yonyoucloud_replace_path//;/ })
        values=(${yonyoucloud_replace_value//;/ })        
        for path in ${paths[@]}
        do
            if [[ -f "${path}" ]];then 
                for value in ${values[@]}
                do                  
                   env_value=`eval echo '$'"${value}"`
                   if [[ ! -z "${env_value}" ]];then
                       replace_value="%$value%"
                       sed -i "s@${replace_value}@${env_value}@" $path
                   fi
                done
            fi
        done
    fi
    #收集用户应用日志
    if [[ "X${developer_app_logs}" != "X" && "X${AGENT_ID}" != "X" ]]; then
        arr=(${developer_app_logs//;/ }) 
            busiPath=${arr[2]}
        logpath=/var/log/datalog/${arr[0]}/${arr[1]}/$AGENT_ID
        # path统一为最后面没有/
        busiPath=${busiPath%/}
        mkdir -p $busiPath
        # 移��原来的日志文件夹
        mv "${busiPath}" "${busiPath}_$(date +'%Y%m%d%H%M%S.%N').bak"
        if [[ $? -eq 0 ]];then
            mkdir -p $logpath
            ln -sf $logpath $busiPath
        else
            echo "ERROR: mount log file error!!!" 1>&2
            echo "==error info: ${developer_app_logs}" 1>&2
            echo "==error info: ${AGENT_ID}" 1>&2
        fi
    fi
}

#### RUN ####
main(){
    printf "options timeout:1 attempts:1 rotate\nnameserver 10.3.15.14\nnameserver 10.3.15.15\n" > /etc/resolv.conf
    source /usr/local/bin/java_options.sh
    set_apm
    confcenter
    set_applog
}
main
eval $@