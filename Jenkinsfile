#!groovy

node {
    checkout scm

    String dockerRegistry = env.DOCKER_REGISTRY ?: ''
    String dockerRegistryCredentialId = env.DOCKER_REGISTRY_CREDENTIAL ?: 'credential-docker-registry'
    String dockerImageName = 'lechuckroh/restapi-template-go'
    String dockerImageVersion = "1.0.${env.BUILD_ID}".toString()

    stage('Build image') {
        sh 'make build-image'
    }

    stage('Test') {
        sh 'make cover'
        cobertura autoUpdateHealth: false,
                autoUpdateStability: false,
                coberturaReportFile: 'reports/coverage.xml',
                conditionalCoverageTargets: '70, 0, 0',
                failNoReports: false,
                failUnhealthy: false,
                failUnstable: false,
                lineCoverageTargets: '80, 0, 0',
                maxNumberOfBuilds: 0,
                methodCoverageTargets: '80, 0, 0',
                onlyStable: false,
                sourceEncoding: 'ASCII',
                zoomCoverageChart: false
        junit 'reports/junit.xml'
    }

    def image = null
    stage('Build docker image') {
        image = docker.build("${dockerImageName}:${dockerImageVersion}")
    }

    stage('Push docker image') {
        docker.withRegistry(dockerRegistry, dockerRegistryCredentialId) {
            image.push()
            image.push('latest')
        }
    }
}
