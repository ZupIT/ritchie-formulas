def jobNameParts = JOB_NAME.tokenize('/') as String[];
def githubDestinationRepo = jobNameParts[0];
def githubDestinationBranch = "marte"
def githubDestinationOrg = "martetech"
def buildable = false

pipeline{
    agent {
	docker {
	    image 'viniciusramosdefaria/formulabuilder:1.0'
	    args  '--privileged -u 0:0'
	}
    }


    stages
    {
        stage("Setting environment variables"){
          parallel {
              stage('master') {
                  when {
                      expression {
                          return branch_name =~ /^master/
                      }
                  }
                  steps {
                      script{
                        RITCHIE_AWS_ACCESS_KEY_ID = env["DOCKER_AWS_ACCESS_KEY_ID_PRODUCTION_MARTE"]
                        RITCHIE_AWS_SECRET_ACCESS_KEY = env["DOCKER_AWS_SECRET_ACCESS_KEY_PRODUCTION_MARTE"]
                        RITCHIE_AWS_REGION = "sa-east-1"
                        RITCHIE_AWS_BUCKET = "ritchie-cli-bucket152849730126474"
                        buildable = true
                      }
                  }
              }

              stage('qa') {
                  when {
                      expression {
                          return branch_name =~ /^qa/
                      }
                  }
                  steps {
                      script{
                        RITCHIE_AWS_ACCESS_KEY_ID = env["DOCKER_AWS_ACCESS_KEY_ID_QA_MARTE"]
                        RITCHIE_AWS_SECRET_ACCESS_KEY = env["DOCKER_AWS_SECRET_ACCESS_KEY_QA_MARTE"]
                        RITCHIE_AWS_REGION = "sa-east-1"
                        RITCHIE_AWS_BUCKET = "ritchie-cli-bucket234376412767550"
                        buildable = true
                      }
                  }
              }
          }
        }

        stage("Building formulas and sending them to s3"){
           when {
              expression {
                  buildable
              }
            }
            steps {
                script{

                    checkout scm

                    withCredentials(
                      [
                        string(credentialsId: RITCHIE_AWS_ACCESS_KEY_ID, variable: 'aws_access_key_id_unveil'),
                        string(credentialsId: RITCHIE_AWS_SECRET_ACCESS_KEY, variable: 'aws_secret_access_key_unveil'),
                      ]) {
                        try{
                            sh "AWS_ACCESS_KEY_ID=${aws_access_key_id_unveil} AWS_SECRET_ACCESS_KEY=${aws_secret_access_key_unveil} AWS_DEFAULT_REGION=${RITCHIE_AWS_REGION} RITCHIE_AWS_BUCKET=${RITCHIE_AWS_BUCKET} make push-s3"
                        } catch (Error error){
                            echo "Error while building or pushing to s3"
                        }
                    }
                }
            }
      }
        stage("Sync with martetech repo") {
            when {
              branch 'marte'
            }
            steps {
                withCredentials([usernamePassword(credentialsId: 'github-ci-marte-zup', passwordVariable: 'git_passwd', usernameVariable: 'git_user')]) {
                    sh "git config --global user.name ${git_user}"
                    sh "git config --global user.email ${git_user}@zup-jenkins.com"
                    sh "git remote rm upstream || exit 0"
                    sh "git remote add upstream https://${git_user}:${git_passwd}@github.com/${githubDestinationOrg}/${githubDestinationRepo}.git"
                    sh "git remote -v"
                    sh "rm -rf vivo"
                    sh "git add . && git commit -m \"jenkins: rm unnecessary files\""
                    sh "git fetch upstream"
                    sh "git push -u upstream HEAD:${githubDestinationBranch} -f"
                }
            }
        }
    }

    post {
        success {
            echo "Build and push successfully executed by Jenkins"
        }
        failure {
            echo "Build failed"
        }
    }
}

