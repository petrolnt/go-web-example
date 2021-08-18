//@Library('utils@main') import net.petrolnt.utilities.*
@Library('utils@main')_

pipeline { 
    environment { 
        registry = "petrolnt/web-v1" 
        registryCredential = 'dockerhub_petrol' 
        dockerImage = ''
        github_id =  "petrol_github"
        branchName = "main"
        gitUrl = "https://github.com/petrolnt/go-web-example.git"
    }
    agent any 
    stages { 
        stage('Cloning from Github') { 
            steps {
                script {
                    git branch: branchName, credentialsIdL: github_id, url: gitUrl
                }
            }
        } 
        stage('Building an image') { 
            steps{
                script {
                    //dockerImage = docker.build registry + ":$BUILD_NUMBER"
                    dockerImage = docker.build registry + ":latest"
                }
            } 
        }
        stage('Deploy an image to Dockerhub') { 
            steps { 
                script { 
                    docker.withRegistry( '', registryCredential ) { 
                        dockerImage.push() 
                    }
                } 
            }
        }
        stage('Deploy to local K8S'){
            steps{
                script{
                    kubernetesDeploy(configs: "kuber/app.yaml", kubeconfigId: "test_eks1")
                }
            }
        }
        
        stage('Check service'){
            steps{
                checkService()
            }
        }
    }
    post{ 
        always { 
            //sh "docker rmi $registry:$BUILD_NUMBER" 
            sh "docker rmi $registry:latest" 
        }
    } 
}
