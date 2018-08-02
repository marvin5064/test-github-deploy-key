#!/bin/bash
set -o nounset
set -o errexit
set -o pipefail

# local variables
APP_NAME=test-github-deploy-key
PUBLISH_REPO_NAME=git@github.com:marvin5064/$APP_NAME.git
DEPLOY_KEY=/root/.ssh/id_rsa
PKG_VERSION=0.0.1
echo $GITHUB_DEPLOY_KEY > $DEPLOY_KEY
sed -i 's/\\n/\n/g' $DEPLOY_KEY
chmod 600 $DEPLOY_KEY

git clone $PUBLISH_REPO_NAME
cd $APP_NAME

# Check current tags
if [[ $(git tag) == *$PKG_VERSION* ]]; then
	echo "version $PKG_VERSION was already published. Please increment the version number in .version file."
	exit 0;
fi

# If it doesn't push the commit and tag the version
echo "Publishing version $PKG_VERSION"

rm -rf  **/*
cp -a /app/output/. .

git add --all
git commit -nm "Publishing version $PKG_VERSION"
git tag -a $PKG_VERSION -m "$PKG_VERSION"
git push --tags origin master