FROM jenkins/jenkins:lts-jdk11

ENV CASC_JENKINS_CONFIG=/var/jenkins_home/jcasc_configs
ENV JENKINS_HOME=/var/jenkins_home
# JAVA_OPTS set in the environment that jenkins.jar runs in.
ENV JAVA_OPTS="-Djenkins.install.runSetupWizard=false -Xms1G -Xmx4G"
ENV JINX_SEED_DESCRIPTION="Default Seed Job provided by Jinx"
ENV JINX_SEED_JENKINSFILE="${JENKINS_HOME}/jinx_support_files/Seed_Jenkinsfile"

# Directory containing credentials files that will be mounted into the image.
# https://github.com/jenkinsci/configuration-as-code-plugin/blob/master/docs/features/secrets.adoc#docker-secrets
ENV SECRETS_DIR="./secrets"

USER root

RUN mkdir /run/secrets/ && chown jenkins:jenkins /run/secrets/ && chmod 0700 /run/secrets \
    && mkdir ${CASC_JENKINS_CONFIG} && chown jenkins:jenkins ${CASC_JENKINS_CONFIG}

USER jenkins

# Why each plugin is needed:
#
# plain-credentials is recommended in the github pull request builder's configuration page as the preferred method
# of providing your credentials to Jenkins.
# - https://plugins.jenkins.io/plain-credentials/
# - plugins versions https://get.jenkins.io/plugins/plain-credentials/
# github-branch-source to build out the github support
# - https://docs.cloudbees.com/docs/cloudbees-ci/latest/cloud-admin-guide/github-branch-source-plugin
# - plugins versions https://get.jenkins.io/plugins/github-branch-source/
# configuration-as-code is to allow jenkins elements to be configured with the yaml file format provided by the "Jenkins
# Configuration As Code" plugin
# - https://github.com/jenkinsci/configuration-as-code-plugin
# - plugins versions https://get.jenkins.io/plugins/configuration-as-code/
# workflow-aggregator is used to support the Jenkins Pipeline build file format
# - https://plugins.jenkins.io/workflow-aggregator/
# - plugins versions https://get.jenkins.io/plugins/workflow-aggregator/
# workflow-cps is the Pipeline: Groovy plugin. It provides the CpsFlowDefinition class which is used to define the seed job from a string
# - https://plugins.jenkins.io/workflow-cps/
# - plugins versions https://get.jenkins.io/plugins/workflow-cps/
RUN jenkins-plugin-cli --plugins plain-credentials github-branch-source configuration-as-code workflow-aggregator workflow-cps

ADD ${SECRETS_DIR}/ /run/secrets/
ADD ./init.groovy.d/ ${JENKINS_HOME}/init.groovy.d
ADD ./jinx_support_files/ ${JENKINS_HOME}/jinx_support_files
