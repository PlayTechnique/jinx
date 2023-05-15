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
COPY plugins.txt plugins.txt
RUN jenkins-plugin-cli  --plugin-file plugins.txt

ADD ${SECRETS_DIR}/ /run/secrets/
ADD ./init.groovy.d/ ${JENKINS_HOME}/init.groovy.d
ADD ./jinx_support_files/ ${JENKINS_HOME}/jinx_support_files
