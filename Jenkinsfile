pipeline {
    agent any
    environment {
        registry = "albertomurrodrigo/proyecto-mds"
        GOCACHE = "/tmp"
    }
    stages {
        stage('Build') {
        stage('Build') {
            agent {
                docker {
                    image 'golang'
                }
            }
            steps {
                // Create our project directory.
                sh 'cd ${GOPATH}/src'
                sh 'mkdir -p ${GOPATH}/src/hello-world'
                // Copy all files in our Jenkins workspace to our project directory.
                sh 'cp -r ${WORKSPACE}/* ${GOPATH}/src/hello-world'
                // Build the app.
                sh 'go build'
            }
        }
        stage('Test') {
            agent {
                docker {
                    image 'golang'
                }
            }
            steps {
                // Create our project directory.
                sh 'cd ${GOPATH}/src'
                sh 'mkdir -p ${GOPATH}/src/hello-world'
                // Copy all files in our Jenkins workspace to our project directory.
                sh 'cp -r ${WORKSPACE}/* ${GOPATH}/src/hello-world'
                // Remove cached test results.
                sh 'go clean -cache'
                // Run Unit Tests.
                sh 'go test ./... -v -short'
            }
        }
        stage('Publish') {
            environment {
                registryCredential = 'dockerhub_id'
            }
            steps{
                script {
                    def appimage = docker.build registry + ":$BUILD_NUMBER"
                    docker.withRegistry( '', registryCredential ) {
                        appimage.push()
                        appimage.push('latest')
                    }
                }
            }
        }
        stage ('Deploy') {
            steps {
                script{
                    sh 'kubectl apply -f deployment.yaml'
                }
            }
        }
    }
}