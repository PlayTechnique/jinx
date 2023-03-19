package jinkiesengine

import (
	_ "embed"
	"log"
	"os"
)

//go:embed embed_files/version.txt
var version []byte

func CreateLayout(topLevelDir string) error {
	createDirectories(topLevelDir)
	err := createFiles()

	return err
}

func createDirectories(topLevelDir string) []string {
	directories := []string{"Docker", "init.groovy.d", "secrets", "configFiles"}
	createdDirectories := []string{}

	for _, dir := range directories {
		os.Mkdir(topLevelDir+dir, 0755)
		createdDirectories = append(createdDirectories, topLevelDir+dir)
	}

	return createdDirectories
}

func createFiles() error {
	err := writeDockerFile("Docker", "Dockerfile")

	if err != nil {
		log.Fatal(err)
		return err
	}

	err = writeVersionFile("version.txt")

	if err != nil {
		log.Fatal(err)
		return err
	}

	err = writeJinxConfig("configFiles", "jinx.yml")

	if err != nil {
		log.Fatal(err)
		return err
	}

	writeContainerConfig("configFiles", "containerconfig.yml")
	writeHostConfig("configFiles", "hostconfig.yml")

	return err
}

func writeDockerFile(dir string, filename string) error {
	dockerfile := `FROM jenkins/jenkins:lts-jdk11

ARG casc_jenkins_config
ARG jenkins_home
ARG jenkins_java_opts
ARG secrets_dir
ARG jinkies_seed_description
ARG jinkies_seed_jenkinsfile

ENV CASC_JENKINS_CONFIG=${casc_jenkins_config}
ENV JENKINS_HOME=${jenkins_home}
ENV JAVA_OPTS=${jenkins_java_opts}
ENV JINKIES_SEED_DESCRIPTION=${jinkies_seed_description}
ENV JINKIES_SEED_JENKINSFILE=${jinkies_seed_jenkinsfile}

ENV SECRETS_DIR=${secrets_dir}

# https://github.com/jenkinsci/configuration-as-code-plugin/blob/master/docs/features/secrets.adoc#docker-secrets
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
ADD ./jinkies_support_files/ ${JENKINS_HOME}/jinkies_support_files
`

	err := os.WriteFile(dir+"/"+filename, []byte(dockerfile), 0700)

	if err != nil {
		log.Fatal(err)
	}

	return err
}

func writeVersionFile(filename string) error {
	version := `
0.0.0-alpha-prealpha
`
	err := os.WriteFile(filename, []byte(version), 0700)

	if err != nil {
		log.Fatal(err)
	}

	return err
}

func writeJinxConfig(dir string, filename string) error {
	config := `
---
ContainerName: {{ContainerName}}
PullImages: false
`

	err := os.WriteFile(dir+"/"+filename, []byte(config), 0700)

	if err != nil {
		log.Fatal(err)
	}

	return err
}

func writeContainerConfig(dir string, filename string) error {
	config := `
---
autoremove: true
containerport: "8080/tcp"
Env:
  #The JINKIES_* environment variables are documented in https://github.com/gwynforthewyn/jinkies/blob/main/Docker/build.env
  #- JINKIES_SEED_DESCRIPTION=
  #- JINKIES_SEED_JENKINSFILE=
hostip: "0.0.0.0"
hostport: "8090/tcp"
image: "jamandbees/jinkies:local"
`
	err := os.WriteFile(dir+"/"+filename, []byte(config), 0700)

	if err != nil {
		log.Fatal(err)
	}

	return err

}

func writeHostConfig(dir string, filename string) error {
	config := `
---
AutoRemove: true
PublishAllPorts: true
ExposedPorts: "8091/tcp"
PortBindings:
  8080/tcp:
    HostIp: "0.0.0.0"
    HostPort: "8091/tcp"
`
	err := os.WriteFile(dir+"/"+filename, []byte(config), 0700)

	if err != nil {
		log.Fatal(err)
	}

	return err
}
