#!/bin/sh

jxm_agent_path="/usr/local/jxm_agent"
JVM_DOWNLOAD_URL="http://10.3.35.108:6666"
#JVM_DOWNLOAD_URL="https://ywb-tools.diwork.com"
port="40009"
if [[ "${YWB_JVM_PORT}" != "" ]];then
 port="${YWB_JVM_PORT}"
fi

ZK_ADDRESS=""
JVM_POD_ENV=""
if [ "$dc_app_env" == "online" ]; then
    ZK_ADDRESS='10.3.35.180:2183'
    JVM_POD_ENV="prod"
fi
if [ "$dc_app_env" == "online-dc-commldev" ]; then
    ZK_ADDRESS='10.166.60.207:2181'
    JVM_POD_ENV="prod"
fi
if [ "$dc_app_env" == "test" ]; then
    ZK_ADDRESS='172.20.32.15:2183'
    JVM_POD_ENV="test"
fi

if [ "$dc_app_env" == "pre" ]; then
    ZK_ADDRESS='10.5.3.137:2181,10.5.3.136:2181,10.5.3.135:2181'
    JVM_POD_ENV="pre"
fi
if [ "$dc_app_env" == "pre-r1" ]; then
    ZK_ADDRESS='172.20.28.42:2181'
    JVM_POD_ENV="pre"
fi
if [ "$dc_app_env" == "stage-dc-core1" ]; then
    ZK_ADDRESS='10.166.55.120:2181'
    JVM_POD_ENV="stage"
fi
if [ "$dc_app_env" == "daily" ]; then
    ZK_ADDRESS='10.5.34.53:2181,10.5.34.52:2181,10.5.34.51:2181'
    JVM_POD_ENV="daily"
fi
if [ "$dc_app_env" == "online-sg" ]; then
    ZK_ADDRESS='10.169.5.11:2181,10.169.5.12:2181,10.169.5.13:2181'
    JVM_POD_ENV="prod"
    JVM_DOWNLOAD_URL="http://10.166.28.249:6666"
fi
if [ "$dc_app_env" == "online-dc-core3" ]; then
    ZK_ADDRESS='10.3.106.17:2181,10.3.106.43:2181,10.3.106.69:2181'
    JVM_POD_ENV="prod"
fi
if [ "$dc_app_env" == "online-dc-core4" ]; then
    ZK_ADDRESS='10.44.0.48:2181,10.44.0.90:2181,10.44.0.150:2181'
    JVM_POD_ENV="prod"
    JVM_DOWNLOAD_URL="https://ywb-tools.diwork.com"
fi
if [ "$dc_app_env" == "online-dc-project-jsyd" ]; then
    ZK_ADDRESS='10.200.0.44:2181,10.200.0.93:2181,10.200.0.70:2181'
    JVM_DOWNLOAD_URL="https://ywb-tools.diwork.com"
    JVM_POD_ENV="prod"
fi

mkdir -p $jxm_agent_path >/dev/null 2>&1
wget -O "${jxm_agent_path}/ywb_tool" ${JVM_DOWNLOAD_URL}/soft/jvm/ywb_tool >/dev/null 2>&1
wget -O "${jxm_agent_path}/jmx_config.yaml" ${JVM_DOWNLOAD_URL}/soft/jvm/jmx_config_v3.yaml >/dev/null 2>&1
wget -O "${jxm_agent_path}/jmx_prometheus_javaagent.jar" ${JVM_DOWNLOAD_URL}/soft/jvm/jmx_prometheus_javaagent-0.15.1-v4.jar >/dev/null 2>&1

wget -O "${jxm_agent_path}/topr.py" ${JVM_DOWNLOAD_URL}/soft/jvm/topr.py >/dev/null 2>&1
wget -O "${jxm_agent_path}/topr-go" ${JVM_DOWNLOAD_URL}/soft/jvm/topr-go >/dev/null 2>&1
TOPR_MD5NUM=`md5sum ${jxm_agent_path}/topr.py|awk -F ' ' '{print$1}'`
TOPR_GO_MD5NUM=`md5sum ${jxm_agent_path}/topr-go|awk -F ' ' '{print$1}'`
PYTHON_ENV=`which python3|wc -l`
if [ "${TOPR_MD5NUM}" == "6fcd5b559a638430abca0f32e32f98e6" ] && [ $PYTHON_ENV == 1 ] ;then
  if [ $dc_app_env == "test" ] || [ $dc_app_env == "daily" ] || [ $dc_app_env == "pre" ] || [ $dc_app_env == "stage-dc-core1" ] ||  [ $dc_app_env == "online" ] || [ $dc_app_env == "online-dc-core3" ] || [ $dc_app_env == "online-dc-core4" ] || [ $dc_app_env == "online-dc-commldev" ]; then
          echo "* * * * * flock -xn /opt/topr.lock.file -c 'python3 ${jxm_agent_path}/topr.py >/dev/null 2>&1'" >> /var/spool/cron/crontabs/root
          /usr/sbin/crond
        fi
fi

if [ "${TOPR_GO_MD5NUM}" == "5693946dcae495891e644066823f2635" ] && [ $PYTHON_ENV == 0 ] ;then
  if [ $dc_app_env == "test" ] || [ $dc_app_env == "daily" ] || [ $dc_app_env == "pre" ]  ;then
    chmod +x ${jxm_agent_path}/topr-go
    echo "* * * * * flock -xn /opt/topr.lock.file -c '${jxm_agent_path}/topr-go >/dev/null 2>&1'" >> /var/spool/cron/crontabs/root
    /usr/sbin/crond
  fi
fi

chmod +x ${jxm_agent_path}/ywb_tool

# 适配重复开启JVM监控问题
#local check_opts
check_opts="${JAVA_OPTS} ${CATALINA_OPTS}"
if echo "${check_opts}" | grep 'jmx_prometheus_javaagent.jar' >/dev/null 2>&1;then
  # 重复开启只用注册一下就可以
  ${jxm_agent_path}/ywb_tool -pod-env $JVM_POD_ENV -pod-register -pod-register-zk=${ZK_ADDRESS} -pod-register-port=30013 > /tmp/ywb_tool.log 2>&1 &
else
  # 没有重复开启的情况
  ${jxm_agent_path}/ywb_tool -pod-env $JVM_POD_ENV -pod-register -pod-register-zk=${ZK_ADDRESS} > /tmp/ywb_tool.log 2>&1 &
  echo " -javaagent:${jxm_agent_path}/jmx_prometheus_javaagent.jar=${port}:${jxm_agent_path}/jmx_config.yaml"