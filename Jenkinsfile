pipeline {
    agent any
    tools {
        go '1.18.2'
    }

    environment {
        GO114MODULE = 'on'
        CGO_ENABLED = 0 
        GOPATH = "${JENKINS_HOME}/jobs/${JOB_NAME}/builds/${BUILD_ID}"
    }

    stages {
        stage("Pre Build"){
            steps {
                echo "Installing depedencies"
                sh "go version"
                sh "go build"
            }
        }
    }
}