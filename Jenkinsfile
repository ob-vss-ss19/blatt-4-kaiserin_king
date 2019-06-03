pipeline {
    agent none
    stages {
        stage('Build') {
            agent {
                docker { image 'obraun/vss-protoactor-jenkins' }
            }
            steps {
				sh 'cd client && go build main.go'
                sh 'cd services/booking/proto && make regenerate'
                sh 'cd services/booking && go build main.go'
				sh 'cd services/cinema/proto && make regenerate'
                sh 'cd services/cinema && go build node.go'
				sh 'cd services/movie/proto && make regenerate'
                sh 'cd services/movie && go build node.go'
				sh 'cd services/show/proto && make regenerate'
                sh 'cd services/show && go build main.go'
				sh 'cd services/user/proto && make regenerate'
                sh 'cd services/user && go build main.go'

            }
        }
        stage('Test') {
            agent {
                docker { image 'obraun/vss-protoactor-jenkins' }
            }
            steps {
                sh 'cd tree && go get -v -d -t ./...'
                sh 'go get github.com/t-yuki/gocover-cobertura' // install Code Coverage Tool
                sh 'echo cd tree && go test -v -coverprofile=cover.out' // save coverage info to file
                sh ' echo gocover-cobertura < tree/cover.out > coverage.xml' // transform coverage info to jenkins readable format
                //publishCoverage adapters: [coberturaAdapter('coverage.xml')] publish report on Jenkins
            }
        }
        stage('Lint') {
            agent {
                docker { image 'obraun/vss-protoactor-jenkins' }
            }   
            steps {
                sh 'golangci-lint run --deadline 20m --enable-all'
            }
        }
        stage('Build Docker Image') {
            agent any
            steps {
                sh "docker-build-and-push -b ${BRANCH_NAME} -s client -f client.dockerfile"
                sh "docker-build-and-push -b ${BRANCH_NAME} -s booking -f booking.dockerfile"
				sh "docker-build-and-push -b ${BRANCH_NAME} -s cinema -f cinema.dockerfile"
				sh "docker-build-and-push -b ${BRANCH_NAME} -s movie -f movie.dockerfile"
				sh "docker-build-and-push -b ${BRANCH_NAME} -s show -f show.dockerfile"
				sh "docker-build-and-push -b ${BRANCH_NAME} -s user -f user.dockerfile"
            }
        }
    }
}
