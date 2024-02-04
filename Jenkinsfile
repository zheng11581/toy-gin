pipeline {
    agent any

    parameters {
        string(name: 'NAMESPACE', defaultValue: 'yuncaiprodb', description: '需要做HeapDump的Pod的命名空间')
        string(name: 'POD_NAME', defaultValue: 'nil', description: '需要做HeapDump的Pod的名称')
    }  

    stages {
        stage('生成Dump脚本') {
            steps {
                sh 'kubectl cp /tmp/jmap_dump.sh "$NAMESPACE/$POD_NAME":/tmp/'  
            }
        }

        

        stage('生成HeapDump并上传到OSS') {
            steps {
                sh 'kubectl exec -it -n $NAMESPACE $POD_NAME -- bash /tmp/jmap_dump.sh'
            }
        }

        stage('获取下载Dump文件地址') {
            steps {
                sh 'kubectl exec -it -n $NAMESPACE $POD_NAME -- cat /tmp/minio_upload.txt'
            }
        }


    }
}
