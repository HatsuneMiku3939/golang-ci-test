library changelog: false,
	identifier: 'jenkins-pipeline-action@d9592bde0588f66efc5c3cb10e5dac3467f758d6',
	retriever: modernSCM([
		$class: 'GitSCMSource',
		remote: 'https://github.com/HatsuneMiku3939/jenkins-pipeline-action.git',
		traits: [[$class: 'jenkins.plugins.git.traits.BranchDiscoveryTrait']]])

node {
    stage('checkout') {
        cleanWs()
        git credentialsId: 'checkout-credential', url: 'git@github.com:HatsuneMiku3939/golang-ci-test.git'
    }

    JenkinsAction(
        name: "unittest",
        action: "hatsunemiku/golang-ci-unittest-action:latest",
        args: [],
    )
}
