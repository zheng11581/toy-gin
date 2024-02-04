#!/bin/bash

PIPELINE_NAME="${dc_app_name}"
CONTAINER_NAME="${LOG_INSTANCE_ID}"
APP_CODE="${dc_app_code}"
TOOLS_BASE_DIR=/tmp/sql-pipeline
[ ! -d ${TOOLS_BASE_DIR} ] && mkdir -p ${TOOLS_BASE_DIR}
TOKEN="${1}"
export TZ='Asia/Shanghai'
export BEGINNING_TIME="$( date +'%F %T' )"

# DC_ENV适配
if echo -en "${DC_ENV}" |grep -q '###'; then
    DC=$( echo -en "${DC_ENV}" |awk -F'###' '{print $1}' )
    ENV=$( echo -en "${DC_ENV}" |awk -F'###' '{print $2}' )
elif [ "x${DC_ENV}" == "xstage" ]; then
    DC="bip-core1"
    ENV="${DC_ENV}"
elif [ "x${DC_ENV}" == "xsandbox" ]; then
    DC="bip-sandbox"
    ENV="${DC_ENV}"
else
    DC="yb-ys"
    ENV="${DC_ENV}"
fi

case ${DC} in
    dc-project|dc-core3|dc-core4)
        TOOL_DOMAIN="ywb-tools.diwork.com"
        ;;
    *)
        TOOL_DOMAIN="ywb.yyuap.com"
        ;;
esac

if [ "x${PIPELINE_DEBUG}" = "xYES" ]; then
    echo
    echo -e "\033[31m################################################################################################################################################################\033[0m"
    echo -e "\033[31m#                                                                          DEBUG模式                                                                           #\033[0m"
    echo -e "\033[31m################################################################################################################################################################\033[0m"
    echo
    SCRIPTS_XZ_URL="https://${TOOL_DOMAIN}/run-sql/sql-pipeline-v3.0/tools/debug/scripts.tar.xz"
else
    if [ "x${DC}" == "xbip-sg" ]; then
        SCRIPTS_XZ_URL="http://10.169.3.30/pipeline-tools/pkgs/scripts.tar.xz"
    else
        SCRIPTS_XZ_URL="https://${TOOL_DOMAIN}/run-sql/sql-pipeline-v3.0/tools/scripts.tar.xz"
    fi
fi

echo -e "\033[33m>>> 开始下载工具脚本...\033[0m"
export TZ='Asia/Shanghai'
echo "Time: $( date +'%F %T' )"
for TRY_TIME in 1 2 3; do
    wget -q --no-cache ${SCRIPTS_XZ_URL} -O ${TOOLS_BASE_DIR}/scripts.tar.xz
    if [ $? -eq 0 ]; then
        echo -e "\033[32m下载工具脚本成功!\033[0m"
        break
    else
        if [ ${TRY_TIME} -eq 3 ]; then
            echo -e "\033[31m3次重试下载工具脚本失败, 退出!\033[0m"
            exit 7
        else
            echo -e "\033[31m下载工具脚本失败, 重试...\033[0m"
            sleep 30
        fi
    fi
done

tar -xf ${TOOLS_BASE_DIR}/scripts.tar.xz -C ${TOOLS_BASE_DIR}
if [ $? -ne 0 ]; then
    echo -e "\033[31m解压${TOOLS_BASE_DIR}/scripts.tar.xz失败!!"
    exit 7
fi

case ${DC} in
    dc-project|dc-core3|dc-core4)
        . /tmp/sql-pipeline/scripts/functions_multi_datacenter.sh 2> /dev/null
        ;;
    *)
        . /tmp/sql-pipeline/scripts/functions.sh 2> /dev/null 2> /dev/null
        ;;
esac

case ${DC} in
    bip-core1|yb-ys|dc-core3|dc-core4|dc-commldev|bip-sg|dc-project|dc-devcore1|bip-verify-dev)
        ;;
    *)
        ECHO "错误或未上线的数据中心[ ${DC} ], 数据中心取值范围: bip-core1|yb-ys|dc-core3|dc-core4|dc-commldev|bip-sg|dc-project|dc-devcore1|bip-verify-dev(上线), 其它(未上线)" 31
        alert_Management "error" "" "错误或未上线的数据中心定义[ ${DC} ]!"
        get_Pod_Log "${DC}" "${ENV}" "all:failed"
        exit 1
        ;;
esac

# 畅捷通不进行备份
if ! omit_Bak; then
    ECHO ">>> Init FTP..." 33
    if rpm -i /tmp/sql-pipeline/scripts/rpms/*.rpm; then
        ECHO "Init FTP成功!" 32
    else
        ECHO "Init FTP失败!" 31
        alert_Management "error" "" "Init FTP失败!"
        get_Pod_Log "${DC}" "${ENV}" "all:failed"
        exit 1get_Pod_Log
    fi
    echo
fi

ECHO ">>> 开始获取流水线基础数据..." 33
INFO_JSON="$( curl -s -H "token: ${TOKEN}" "https://${TOOL_DOMAIN}/opsv/batch/sqlmanager/?datacenter=${DC}&env=${ENV}&app_code=${APP_CODE}&source_type=api" |jq )"
INFO_GET_STATUS="$( echo "${INFO_JSON}" |jq '.status' |xargs echo )"
for TRY_TIME in 1 2 3; do
    if [ "x${INFO_GET_STATUS}" == "xfailed" -o "x${INFO_GET_STATUS}" == "xnull" ]; then
            if [ ${TRY_TIME} -eq 3 ]; then
                ECHO "尝试3次请求失败, 退出..." 31
                ERROR_INFO=$( echo "${INFO_JSON}" |jq .error |xargs echo )
                ECHO "错误原因: ${ERROR_INFO}" 31
                alert_Management "error" "" "获取流水线基础数据失败: ${ERROR_INFO}"
                get_Pod_Log "${DC}" "${ENV}" "all:failed"
                exit 4
            else
                ECHO "获取流水线基础数据异常, 重试..." 31
                echo
                sleep 1
                INFO_JSON="$( curl -s -H "token: ${TOKEN}" "https://${TOOL_DOMAIN}/opsv/batch/sqlmanager/?datacenter=${DC}&env=${ENV}&app_code=${APP_CODE}&source_type=api" |jq )"
                INFO_GET_STATUS="$( echo "${INFO_JSON}" |jq '.status' |xargs echo )"
            fi
    else
        ECHO "获取流水线基础数据成功!" 32
        echo "${INFO_JSON}" |jq > ${TOOLS_BASE_DIR}/.info.json
        echo
        break
    fi
done

# MongoDB去重逻辑
MONGO_DS_NUM=$( echo ${INFO_JSON} |jq ".data.mongo | length" )
if [ ${MONGO_DS_NUM} -gt 0 ]; then
    let DS_NUM_IDX=${MONGO_DS_NUM}-1
    MONGO_INDEX_FILE=${TOOLS_BASE_DIR}/mongo_ds_duplicate_removal.list
    > ${MONGO_INDEX_FILE}
    for CUR_IDX in $( seq 0 ${DS_NUM_IDX} ); do
        IDX_MONGO_LOGIC=$( echo ${INFO_JSON} |jq ".data.mongo[${CUR_IDX}].mongo_logical_datasource" |xargs echo )
        IS_DEFAULT="$( echo ${INFO_JSON} |jq ".data.mongo[${CUR_IDX}].is_default_ds" |xargs echo )"
        [ -z "${IDX_MONGO_LOGIC}" ] && IDX_MONGO_LOGIC="Null"
        # 一个环境的MONGO数据源只有一个, 多个DB_SCHEMA的时候, 只在DEFAULT数据源记录COMMIT_ID, 防止多次执行
        if ! grep -q " ${IDX_MONGO_LOGIC}" ${MONGO_INDEX_FILE} && [ "x${IS_DEFAULT}" == "x1" ]; then
            echo "${CUR_IDX} ${IDX_MONGO_LOGIC}" >> ${MONGO_INDEX_FILE}
        fi
    done
fi

PRO_DOMAIN="$( echo ${INFO_JSON} |jq '.data.pro_domain' |xargs echo )"
RESPONSIBLE_PERSON="$( echo ${INFO_JSON} |jq '.data.developer' |xargs echo )"
CONSIGNOR="$( echo ${INFO_JSON} |jq '.data.consignor' |xargs echo )"
CONSIGNOR_EMAIL="$( echo ${INFO_JSON} |jq '.data.consignor_email' |xargs echo )"
DEVELOPER_EMAIL="$( echo ${INFO_JSON} |jq '.data.developer_email' |xargs echo )"

ECHO "SQL流水线版本: v3.0, 平台类型: ${PIPELINE_TYPE}" 33
ECHO "SQL流水线名称: ${PIPELINE_NAME}" 33
ECHO "容器名称: ${CONTAINER_NAME}" 33
ECHO "负责人: ${RESPONSIBLE_PERSON}" 33
echo
ECHO ">>> 开始创建SQL流水线脚本..." 33
echo
# 1. SQL_TYPE
function create_Template_Scripts () {
    local SQL_TYPE="${1}"; local SUB_PIPELINE_ID="${2}"; local SUB_DS_NUM="${3}"
    local SCRIPT_NAME="${SQL_TYPE}_${SUB_PIPELINE_ID}_${SUB_DS_NUM}_${DC}_${ENV}_${APP_CODE}.sh"
    echo "${SUB_PIPELINE_ID} ${SQL_TYPE} /app/${SCRIPT_NAME}" >> ${SCRIPTS_LIST}
    local DS_NUM=$( echo ${INFO_JSON} |jq ".data.${SQL_TYPE} |length" )
    local FIRST_TIME_INIT=""
    for CUR_DS_NUM in $( seq 0 ${DS_NUM} ); do
        [ ${CUR_DS_NUM} -eq ${DS_NUM} ] && break
        local GET_SUB_PIPELINE_ID="$( echo ${INFO_JSON} |jq '.data.'"${SQL_TYPE}"'['"${CUR_DS_NUM}"'].id' |xargs echo )"
        [ "x${GET_SUB_PIPELINE_ID}" != "x${SUB_PIPELINE_ID}" ] && continue
        if [ "x${FIRST_TIME_INIT}" != "xYES" ]; then
            local PIPELINE_TAG="$( echo ${INFO_JSON} |jq '.data.'"${SQL_TYPE}"'['"${CUR_DS_NUM}"'].g_commit_db' |xargs echo )"
            local PARAM_TYPE="$( echo ${INFO_JSON} |jq '.data.'"${SQL_TYPE}"'['"${CUR_DS_NUM}"'].param_type' |xargs echo )"
            local SCHEMA_ALIAS="$( echo ${INFO_JSON} |jq '.data.'"${SQL_TYPE}"'['"${CUR_DS_NUM}"'].schema_alias' |xargs echo )"
            local WHITE_LIST_DIR="$( echo ${INFO_JSON} |jq '.data.'"${SQL_TYPE}"'['"${CUR_DS_NUM}"'].real_exec_sql_git_dir' |xargs echo )"
            local INIT_EXEC_SWITCH="$( echo ${INFO_JSON} |jq '.data.'"${SQL_TYPE}"'['"${CUR_DS_NUM}"'].scripts_init' |xargs echo )"
            local GIT_PATH="$( echo ${INFO_JSON} |jq '.data.'"${SQL_TYPE}"'['"${CUR_DS_NUM}"'].git_path' |xargs echo )"
            local SCRIPTS_PATH="$( echo ${INFO_JSON} |jq '.data.'"${SQL_TYPE}"'['"${CUR_DS_NUM}"'].scripts_path' |xargs echo )"
            local ABNORMAL_ALERT="$( echo ${INFO_JSON} |jq '.data.'"${SQL_TYPE}"'['"${CUR_DS_NUM}"'].abnormal_alert' |xargs echo )"
            local ALERT_THRESHOLD="$( echo ${INFO_JSON} |jq '.data.'"${SQL_TYPE}"'['"${CUR_DS_NUM}"'].alert_threshold' |xargs echo )"
            local IM_GROUP_ID="$( echo ${INFO_JSON} |jq '.data.'"${SQL_TYPE}"'['"${CUR_DS_NUM}"'].im_group_id' |xargs echo )"
            local FILE_MD5_CHECK="$( echo ${INFO_JSON} |jq '.data.'"${SQL_TYPE}"'['"${CUR_DS_NUM}"'].file_md5_check' |xargs echo )"
            local OPEN_SQL_CHECK="$( echo ${INFO_JSON} |jq '.data.'"${SQL_TYPE}"'['"${CUR_DS_NUM}"'].open_sql_check' |xargs echo )"
            local STILL_EXEC_ON_ERROR="$( echo ${INFO_JSON} |jq '.data.'"${SQL_TYPE}"'['"${CUR_DS_NUM}"'].still_exec_on_error' |xargs echo )"
            local DOWNLOAD_REAL_MAP="$( echo ${INFO_JSON} |jq '.data.'"${SQL_TYPE}"'['"${CUR_DS_NUM}"'].download_real_map' |xargs echo )"
            FIRST_TIME_INIT="YES"
        fi

        if [ "x${SQL_TYPE}" == "xmongo" ]; then
            # 没有在去重列表的mongo索引不执行
            if ! grep -q "^${CUR_DS_NUM} " ${MONGO_INDEX_FILE}; then
                return
            fi
            local COMM_DB_HOST="$( echo ${INFO_JSON} |jq '.data.'"${SQL_TYPE}"'['"${CUR_DS_NUM}"'].commit_id_db_url' |xargs echo )"
            local COMM_DB_PORT="$( echo ${INFO_JSON} |jq '.data.'"${SQL_TYPE}"'['"${CUR_DS_NUM}"'].commit_id_db_port' |xargs echo )"
            local COMM_DB_USER="$( echo ${INFO_JSON} |jq '.data.'"${SQL_TYPE}"'['"${CUR_DS_NUM}"'].commit_id_db_user' |xargs echo )"
            local COMM_DB_PASSWD="$( echo ${INFO_JSON} |jq '.data.'"${SQL_TYPE}"'['"${CUR_DS_NUM}"'].commit_id_db_passwd' |xargs echo )"
            local LOGIC_DS_CODE="$( echo ${INFO_JSON} |jq '.data.'"${SQL_TYPE}"'['"${CUR_DS_NUM}"'].mongo_logical_datasource' |xargs echo )"
            break
        else
            local SUB_COMM_DB_HOST="$( echo ${INFO_JSON} |jq '.data.'"${SQL_TYPE}"'['"${CUR_DS_NUM}"'].commit_id_db_url' |xargs echo )"
            local SUB_COMM_DB_PORT="$( echo ${INFO_JSON} |jq '.data.'"${SQL_TYPE}"'['"${CUR_DS_NUM}"'].commit_id_db_port' |xargs echo )"
            local SUB_COMM_DB_USER="$( echo ${INFO_JSON} |jq '.data.'"${SQL_TYPE}"'['"${CUR_DS_NUM}"'].commit_id_db_user' |xargs echo )"
            local SUB_COMM_DB_PASSWD="$( echo ${INFO_JSON} |jq '.data.'"${SQL_TYPE}"'['"${CUR_DS_NUM}"'].commit_id_db_passwd' |xargs echo )"
            local SUB_BUSI_DB_HOST="$( echo ${INFO_JSON} |jq '.data.'"${SQL_TYPE}"'['"${CUR_DS_NUM}"'].dbhost' |xargs echo )"
            local SUB_BUSI_DB_PORT="$( echo ${INFO_JSON} |jq '.data.'"${SQL_TYPE}"'['"${CUR_DS_NUM}"'].dbport' |xargs echo )"
            local SUB_BUSI_DB_USER="$( echo ${INFO_JSON} |jq '.data.'"${SQL_TYPE}"'['"${CUR_DS_NUM}"'].dbuser' |xargs echo )"
            local SUB_BUSI_DB_PASSWD="$( echo ${INFO_JSON} |jq '.data.'"${SQL_TYPE}"'['"${CUR_DS_NUM}"'].dbpwd' |xargs echo )"
            local SUB_BUSI_DB_SCHEMA="$( echo ${INFO_JSON} |jq '.data.'"${SQL_TYPE}"'['"${CUR_DS_NUM}"'].db_schema' |xargs echo )"
            local COMM_DB_HOST="${COMM_DB_HOST}###${SUB_COMM_DB_HOST}"
            local COMM_DB_PORT="${COMM_DB_PORT}###${SUB_COMM_DB_PORT}"
            local COMM_DB_USER="${COMM_DB_USER}###${SUB_COMM_DB_USER}"
            local COMM_DB_PASSWD="${COMM_DB_PASSWD}###${SUB_COMM_DB_PASSWD}"
            local BUSI_DB_HOST="${BUSI_DB_HOST}###${SUB_BUSI_DB_HOST}"
            local BUSI_DB_PORT="${BUSI_DB_PORT}###${SUB_BUSI_DB_PORT}"
            local BUSI_DB_USER="${BUSI_DB_USER}###${SUB_BUSI_DB_USER}"
            local BUSI_DB_PASSWD="${BUSI_DB_PASSWD}###${SUB_BUSI_DB_PASSWD}"
            local BUSI_DB_SCHEMA="${BUSI_DB_SCHEMA}###${SUB_BUSI_DB_SCHEMA}"
        fi
    done

    COMM_DB_HOST="$( echo ${COMM_DB_HOST} |sed 's@^###@@g' )"
    COMM_DB_PORT="$( echo ${COMM_DB_PORT} |sed 's@^###@@g' )"
    COMM_DB_USER="$( echo ${COMM_DB_USER} |sed 's@^###@@g' )"
    COMM_DB_PASSWD="$( echo ${COMM_DB_PASSWD} |sed 's@^###@@g' )"
    BUSI_DB_HOST="$( echo ${BUSI_DB_HOST} |sed 's@^###@@g' )"
    BUSI_DB_PORT="$( echo ${BUSI_DB_PORT} |sed 's@^###@@g' )"
    BUSI_DB_USER="$( echo ${BUSI_DB_USER} |sed 's@^###@@g' )"
    BUSI_DB_PASSWD="$( echo ${BUSI_DB_PASSWD} |sed 's@^###@@g' )"
    BUSI_DB_SCHEMA="$( echo ${BUSI_DB_SCHEMA} |sed 's@^###@@g' )"

    # 数据源公共变量
    cat << EOF > /app/${SCRIPT_NAME}
# 流水线名称: ${PIPELINE_NAME}
# 流水线编码: ${APP_CODE}
# 负责人: ${RESPONSIBLE_PERSON}
# 脚本生成时间: $( date +"%F %T" )

# 包含两部分[DataCenter(DC), SubEnv], 中间用三个#隔开, DataCenter###SubEnv, DC可缺省, 缺省DC取值: DC_ENV="stage" -> DC="bip-core1", 其它(test,daily,pre,online) -> DC="yb-ys"
# bip-core1: YonBIP核心1数据中心; yb-ys: YonBIP公共数据中心[旧][核心2]; dc-core3: YonBIP核心3数据中心; bip-sg: YonBIP新加坡数据中心; dc-commldev: YonBIP商用开发数据中心;
# dc-devcore1: YonBIP研发核心1数据中心; bip-verify-dev: YonBIP公有云专属化环境; dc-project: YonBIP项目数据中心
DC_ENV="${DC_ENV}"

# SQL流水线负责人
RESPONSIBLE_PERSON="${RESPONSIBLE_PERSON}"

# SQL流水线负责人邮箱
DEVELOPER_EMAIL="${DEVELOPER_EMAIL}"

# 提单人
CONSIGNOR="${CONSIGNOR}"

# 提单人邮箱
CONSIGNOR_EMAIL="${CONSIGNOR_EMAIL}"

# 领域
PRO_DOMAIN="${PRO_DOMAIN}"

# 容器名称
CONTAINER_NAME="${CONTAINER_NAME}"

# 数据源中物理连接池的数量, 是并发执行的参考
DS_NUM=${SUB_DS_NUM}

# 脚本类型, 分为[ frame|data|mongo ]
SQL_TYPE="${SQL_TYPE}"

# 环境级流水线的编号ID
SUB_PIPELINE_ID="${SUB_PIPELINE_ID}"

# 流水线类型
PARAM_TYPE="${PARAM_TYPE}"

# 数据库别名
SCHEMA_ALIAS="${SCHEMA_ALIAS}"

# INIT脚本执行开关
INIT_EXEC_SWITCH="${INIT_EXEC_SWITCH}"

# GIT_PATH是Dockerfile中ADD . /ADD/TO/SOMEWHERE的路径; SCRIPTS_PATH是脚本存放的子目录[首尾不带/], /GIT_PATH/SCRIPTS_PATH这个路径在容器中必须是真实存在的
GIT_PATH="${GIT_PATH}"
SCRIPTS_PATH="${SCRIPTS_PATH}"

# 脚本执行的目录白名单, 目录可以多层[ 如: /0001_data/0002_domain1 ], 多个目录用"|"分隔, 两端不能有"|"[ 如"0001_abc|0002_def" ]
WHITE_LIST_DIR="${WHITE_LIST_DIR}"

# ABNORMAL_ALERT是IM通知的开关, 只有���是[YES]的时候, ALERT_THRESHOLD[超时阈值]和[IM_GROUP_ID IM群ID]才会生效
ABNORMAL_ALERT="${ABNORMAL_ALERT}"
ALERT_THRESHOLD="${ALERT_THRESHOLD}"
IM_GROUP_ID="${IM_GROUP_ID}"

# COMMIT_ID库的信息, 分别是[HOST, PORT, USER, PASSWORD, SQL流水线标识]
COMM_DB_HOST="${COMM_DB_HOST}"
COMM_DB_PORT="${COMM_DB_PORT}"
COMM_DB_USER="${COMM_DB_USER}"
COMM_DB_PASSWD="${COMM_DB_PASSWD}"
PIPELINE_TAG="${PIPELINE_TAG}"

EOF
    if ! [ "x${SQL_TYPE}" == "xmongo" ]; then
        cat << EOF >> /app/${SCRIPT_NAME}
# 是否进行MD5校验, 缺省为[NO], 开启后[YES], 将每次升级成功的文件MD5记录到数据库中, 下次执行前先进行判断, 若已经执行过的脚本直接忽略
FILE_MD5_CHECK="${FILE_MD5_CHECK}"

# SQL语法检测的开关, 默认[YES], [NO]可以关闭SQL语法检查, SQL文件过大或者特殊情况下需要临时关闭
OPEN_SQL_CHECK="${OPEN_SQL_CHECK}"

# 遇到错误始终执行, 开启后[YES]遇到错误继续执行后面的脚本, 不退出升级任务, 缺省为NO
STILL_EXEC_ON_ERROR="${STILL_EXEC_ON_ERROR}"

# 业务库的信息, 分别是[HOST, PORT, USER, PASSWORD, SCHEMA_NAME]
BUSI_DB_HOST="${BUSI_DB_HOST}"
BUSI_DB_PORT="${BUSI_DB_PORT}"
BUSI_DB_USER="${BUSI_DB_USER}"
BUSI_DB_PASSWD="${BUSI_DB_PASSWD}"
BUSI_DB_SCHEMA="${BUSI_DB_SCHEMA}"

EOF
    else
        cat << EOF >> /app/${SCRIPT_NAME}
# 领域自己的MongoDB逻辑数据源编码, 如果没有需要在YMS控制台配置
LOGIC_DS_CODE="${LOGIC_DS_CODE}"

EOF
    fi

    if [ "x${SQL_TYPE}" == "xdata" ]; then
        cat << EOF >> /app/${SCRIPT_NAME}
# 下载真实的MAP文件的开关, 默认[YES], 这个变量现在用得很少, 之前是用于对脚本的测试, 防止错误的脚本将数据库中的数据搞坏, 用一套纯测试的Schema来作MAP
DOWNLOAD_REAL_MAP="${DOWNLOAD_REAL_MAP}"

EOF
    fi

    cat << EOF >> /app/${SCRIPT_NAME}
# 判断下载Scripts工具脚本
if [ ! -d /tmp/sql-pipeline/scripts ]; then
    for TRY_TIME in 1 2 3; do
        wget -q --no-cache ${SCRIPTS_XZ_URL} -O /tmp/sql-pipeline/scripts.tar.xz
        if [ \$? -eq 0 ]; then
            echo -e "\033[32m下载工具脚本成功!\033[0m"
            break
        else
            if [ \${TRY_TIME} -eq 3 ]; then
                echo -e "\033[31m3次重试下载工具脚本失败, 退出!\033[0m"
                exit 7
            else
                echo -e "\033[31m下载工具脚本失败, 重试...\033[0m"
                sleep 10
            fi
        fi
    done

    tar -xf /tmp/sql-pipeline/scripts.tar.xz -C /tmp/sql-pipeline/
    if [ \$? -ne 0 ]; then
        echo -e "\033[31m解压/tmp/sql-pipeline/scripts.tar.xz失败!!"
        exit 7
    fi
fi

. /tmp/sql-pipeline/scripts/template.module
RETVAL=\$?
center_Print " ∧∧∧∧∧∧ 脚本[ ${SCRIPT_NAME} ]执行完毕 ∧∧∧∧∧∧ " 33
echo
exit \${RETVAL}

EOF
}

# 创建脚本
SCRIPTS_LIST=/tmp/sql-pipeline/pipeline_scripts.list
> ${SCRIPTS_LIST}

case ${ENV} in
    test|daily)
        CREATE_TYPES="mongo frame data"
        ;;
    *)
        CREATE_TYPES="frame data mongo"
        ;;
esac

for CREATE_TYPE in $( echo ${CREATE_TYPES} ); do
    CHECK_NULL=$( echo "${INFO_JSON}" |jq ".data.${CREATE_TYPE}" )
    [ "x${CHECK_NULL}" == "xnull" ] && continue
    echo "${INFO_JSON}" |jq ".data.${CREATE_TYPE}[].id" |sort |uniq -c |while read INFO_LINE; do
        SUB_DS_NUM="$( echo ${INFO_LINE} |awk '{print $1}' )"
        SUB_PIPELINE_ID="$( echo ${INFO_LINE} |awk '{print $2}' )"
        create_Template_Scripts "${CREATE_TYPE}" "${SUB_PIPELINE_ID}" "${SUB_DS_NUM}"
    done
done

# 杀掉YMSTranslateStudio进程释放内存
function kill_YMSTranslateStudio () {
    if ps aux |grep -v "grep" |grep -q "YMSTranslateStudio"; then
        ps aux |grep -v "grep" |grep "YMSTranslateStudio" |awk '{print $2}' |xargs kill -9
    fi
}

# 初始化流水线的执行状态
STATUS_TEXT_FILE=/tmp/sql-pipeline/.total_exec.status
cat ${SCRIPTS_LIST} |awk '{printf "%s:%s:confirm\n", $1,$2}' > ${STATUS_TEXT_FILE}
exec_Status_Inbase # 状态全量初始化为待确认

# 执行脚本
while read SCRIPTS_LINE; do
    SUB_PIPELINE_ID=$( echo "${SCRIPTS_LINE}" |awk '{print $1}' )
    SUB_PIPELINE_TYPE=$( echo "${SCRIPTS_LINE}" |awk '{print $2}' )
    SCRIPTS_NAME=$( echo "${SCRIPTS_LINE}" |awk '{print $3}' )
    EXEC_STATUS_FILE=/tmp/sql-pipeline/exec_${SUB_PIPELINE_ID}.status
    /bin/bash ${SCRIPTS_NAME} |tee -a /app/sql-pipeline-${APP_CODE}-${SUB_PIPELINE_ID}.log
    CUR_STATUS=$( cat ${EXEC_STATUS_FILE} )
    if [ ${CUR_STATUS} -ne 0 ]; then
        ECHO "[ ${SCRIPTS_NAME} ]执行报错, 后序脚本中断执行, 请处理完错误之后重试..." 31
        sed -ri "/^${SUB_PIPELINE_ID}:/s/confirm$|upgrading$/failed/g" ${STATUS_TEXT_FILE}
        exec_Status_Inbase # 标记单条脚本状态为failed
        kill_YMSTranslateStudio
        exit ${CUR_STATUS}
    else
        sed -ri "/^${SUB_PIPELINE_ID}:/s/upgrading$|confirm$/success/g" ${STATUS_TEXT_FILE}
        exec_Status_Inbase # 标记单条脚本状态为success
    fi
done <<< "$( cat ${SCRIPTS_LIST} )"

kill_YMSTranslateStudio

# 自动标记SQL流水线历史执行报错
clear_Error_History

ECHO "========================================================================== The  End ! ==========================================================================" 33