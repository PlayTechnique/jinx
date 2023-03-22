package jinkiesengine

import (
	_ "embed"
	jinxtypes "jinx/types"
	"log"
	"os"
	"text/template"
)

//go:embed embed_files/version.txt
var version []byte

func CreateLayout(topLevelDir string, globalRuntime jinxtypes.JinxGlobalRuntime) error {
	_, err := createFiles(topLevelDir, globalRuntime)

	return err
}

func createFiles(topLevelDir string, globalRuntime jinxtypes.JinxGlobalRuntime) ([]string, error) {
	var createdFiles []string

	filename, err := writeDockerFile(topLevelDir+"/Docker", "Dockerfile")

	if err != nil {
		log.Fatal(err)
		return createdFiles, err
	}

	createdFiles = append(createdFiles, filename)

	filename, err = writeVersionFile("version.txt")

	if err != nil {
		log.Fatal(err)
		return createdFiles, err
	}

	createdFiles = append(createdFiles, filename)

	filename, err = writeJinxConfig(topLevelDir+"/configFiles", "jinx.yml", globalRuntime)

	if err != nil {
		log.Fatal(err)
		return createdFiles, err
	}

	createdFiles = append(createdFiles, filename)

	filename, err = writeContainerConfig(topLevelDir+"/configFiles", "containerconfig.yml")

	if err != nil {
		log.Fatal(err)
		return createdFiles, err
	}

	createdFiles = append(createdFiles, filename)

	filename, err = writeHostConfig(topLevelDir+"/configFiles", "hostconfig.yml")

	if err != nil {
		log.Fatal(err)
		return createdFiles, err
	}

	createdFiles = append(createdFiles, filename)

	return createdFiles, err
}

func writeDockerFile(dir string, filename string) (string, error) {
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
	err := os.MkdirAll(dir, 0700)

	if err != nil {
		log.Fatal(err)
	}

	err = os.WriteFile(dir+"/"+filename, []byte(dockerfile), 0700)

	if err != nil {
		log.Fatal(err)
	}

	return dir + "/" + filename, err
}

func writeVersionFile(filename string) (string, error) {
	version := `
0.0.0-alpha-prealpha
`
	err := os.WriteFile(filename, []byte(version), 0700)

	if err != nil {
		log.Fatal(err)
	}

	return filename, err
}

func writeJinxConfig(dir string, filename string, globalRuntime jinxtypes.JinxGlobalRuntime) (string, error) {
	config := `
---
ContainerName: {{ .ContainerName }}
PullImages: false
`
	t, err := template.New("jinxConfig").Option("missingkey=error").Parse(config)

	if err != nil {
		log.Fatal(err)
		return "", err
	}

	err = os.MkdirAll(dir, 0755)

	if err != nil {
		log.Fatal(err)
		return dir, err
	}

	jinxConfigPath := dir + "/" + filename
	file, err := os.Create(jinxConfigPath)

	if err != nil {
		log.Fatal(err)
		return jinxConfigPath, err
	}

	if err != nil {
		log.Fatal(err)
		return jinxConfigPath, err
	}

	err = t.Execute(file, globalRuntime)

	return jinxConfigPath, err
}

func writeContainerConfig(dir string, filename string) (string, error) {
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
	err := os.MkdirAll(dir, 0755)

	if err != nil {
		log.Fatal(err)
	}

	err = os.WriteFile(dir+"/"+filename, []byte(config), 0700)

	if err != nil {
		log.Fatal(err)
	}

	return dir + "/" + filename, err

}

func writeHostConfig(dir string, filename string) (string, error) {
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
	err := os.MkdirAll(dir, 0755)

	if err != nil {
		log.Fatal(err)
	}

	err = os.WriteFile(dir+"/"+filename, []byte(config), 0700)

	if err != nil {
		log.Fatal(err)
	}

	return dir + "/" + filename, err
}
