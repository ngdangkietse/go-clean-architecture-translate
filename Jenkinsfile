def image = "01092002/go-clean-architecture-translate"
def branch = "main"

pipeline {
    agent any

    environment {
        REPO_URL = "https://github.com/ngdangkietse/go-clean-architecture-translate.git"
        GIT_CREDENTIALS = "Git_Credentials"
        DOCKER_CREDENTIALS = credentials('Docker_Credentials')
    }

    stages {
        stage('Checkout') {
            steps {
                script {
                   checkout([
                        $class: 'GitSCM',
                        branches: [[name: '*/main']],
                        extensions: [],
                        userRemoteConfigs: [[
                            url: env.REPO_URL,
                            credentialsId: env.GIT_CREDENTIALS
                        ]]
                   ])
                }
            }
        }

        stage('Docker build') {
            steps {
                script {
                    sh 'docker build -t ${image}:latest .'
                    echo 'Build image completed'
                }
            }
        }

        stage('Docker login') {
            steps {
                script {
                    sh 'echo $DOCKERHUB_CREDENTIALS_PSW | docker login -u $DOCKERHUB_CREDENTIALS_USR --password-stdin'
                    echo 'Login completed'
                }
            }
        }

        stage('Docker push') {
            steps {
                script {
                    sh 'docker push ${image}:latest'
                    echo 'Push image completed'
                }
            }
        }
    }
}