pipeline {
    agent none
    stages {
        stage('Build') {
            agent {
                docker { image 'obraun/vss-protoactor-jenkins' }
            }
            steps {
				sh 'cd client && go build main.go'
                sh 'cd services/booking && go build main.go'
                sh 'cd services/cinema && go build main.go'
                sh 'cd services/movie && go build main.go'
                sh 'cd services/show && go build main.go'
                sh 'cd services/user && go build main.go'

            }
        }
        stage('Test') {
            agent {
                docker { image 'obraun/vss-protoactor-jenkins' }
            }
            steps {
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
