// This is Placeholder for a Jenkins pipeline script

pipeline {
    agent any
    
    environment {
        DOCKER_IMAGE = 'my-app'
        DOCKER_TAG = "v${BUILD_NUMBER}"
        DOCKER_FULL_IMAGE = "${DOCKER_IMAGE}:${DOCKER_TAG}"
    }
    
    stages {
        stage('Checkout') {
            steps {
                checkout scm
            }
        }
        
        stage('Build Docker Image') {
            steps {
                script {
                    sh "docker build -t ${DOCKER_FULL_IMAGE} ."
                }
            }
        }
        
        stage('Push to Registry') {
            steps {
                script {
                    withCredentials([usernamePassword(credentialsId: 'docker-hub-credentials', usernameVariable: 'DOCKER_USERNAME', passwordVariable: 'DOCKER_PASSWORD')]) {
                        // Login to Docker Hub
                        sh "echo ${DOCKER_PASSWORD} | docker login -u ${DOCKER_USERNAME} --password-stdin"
                        
                        // Tag and push the image
                        sh "docker tag ${DOCKER_FULL_IMAGE} ${DOCKER_USERNAME}/${DOCKER_FULL_IMAGE}"
                        sh "docker push ${DOCKER_USERNAME}/${DOCKER_FULL_IMAGE}"
                    }
                }
            }
        }
        
    }
    
    post {
        always {
            script {
                // Clean up Docker images
                sh "docker rmi ${DOCKER_FULL_IMAGE} || true"
                // Clean workspace
                cleanWs()
            }
        }
        success {
            echo 'Docker CI Pipeline completed successfully!'
        }
        failure {
            echo 'Docker CI Pipeline failed!'
        }
    }
}