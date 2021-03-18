pipeline {
    agent any
    environment {
        registry = "albertomurrodrigo/proyecto-mds"
        GOCACHE = "/tmp/.cache"
        XDG_CACHE_HOME = "/tmp/.cache"

    }
    stages {
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
        stage('Unit tests') {
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
                sh 'cp -r ${WORKSPACE}/go.mod ${GOPATH}/src/hello-world'
                sh 'cp -r ${WORKSPACE}/hello_json.go ${GOPATH}/src/hello-world'
                sh 'cp -r ${WORKSPACE}/hello_json_test.go ${GOPATH}/src/hello-world'
                // Remove cached test results.
                sh 'go clean -cache'
                // Run Unit Tests.
                sh 'go test hello.* -v -short'
            }
        }
        stage('Static code analysis'){
            agent {
                docker {
                    image 'golangci/golangci-lint'
                }
            }
            steps {
                // Create our project directory.
                sh 'cd ${GOPATH}/src'
                sh 'mkdir -p ${GOPATH}/src/hello-world'
                // Copy all files in our Jenkins workspace to our project directory.
                sh 'cp -r ${WORKSPACE}/* ${GOPATH}/src/hello-world'
                sh 'chmod -R 777 ${GOPATH}/src/hello-world'
                catchError {
                    sh 'golangci-lint run'
                }
            }
            post {
                success {
                    echo 'Static code analysis stage successful'
                }
                failure {
                    error('Build is aborted due to failure of static code analysis stage')
                }
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
        stage('Unit tests') {
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
        stage ('Deploy') {
            steps {
                script {
                    sh 'kubectl apply -f deployment.yml'
                    sh 'kubectl apply -f service.yml'
                }
            }
        }
        stage ('Smoke') {
            steps {
                script {
                    // Remove cached test results.
                    sh '/home/alberto/sdk/go1.16.2/bin/go run smoke/main_smoke.go'
                }
            }
        }
    }
}