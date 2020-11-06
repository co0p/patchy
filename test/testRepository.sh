#!/bin/bash

REPO_NAME=testRepository
BRANCH_NAME=testBranch

echo "************************************"
echo "*******  CONFIGURE GIT  *******"
echo "************************************"

GIT_EMAIL=`git config --get user.email`

if [ -z "$GIT_EMAIL" ]
then
    echo "setting user.email"
    git config --global user.email "you@example.com"
fi

GIT_NAME=`git config --get user.name`
if [ -z "$GIT_NAME" ]
then
      echo "setting user.name"
      git config --global user.name "Your Name"
fi

echo "************************************"
echo "*******  CREATING TEST REPO  *******"
echo "************************************"

if [[ -d "$REPO_NAME" ]]
then
    echo "********* removing existing repo ********* "
    rm -rf $REPO_NAME
fi

echo "********* creating repository directory *********"
mkdir $REPO_NAME
[ $? -eq 0 ]  || exit 1

echo "********* init git"
cd $REPO_NAME
git init
[ $? -eq 0 ]  || exit 1

echo "********* adding a few commits to master ********* "
echo "first" > first.txt
git add first.txt
git commit -am "first"
echo "second" > second.txt
git add second.txt
git commit -am "second"
echo "third" > third.txt
git add third.txt
git commit -am "third"

echo "********* adding commits to branch ********* "
git checkout -b $BRANCH_NAME
echo "branch" > branch.txt
git add branch.txt
git commit -am "first branch"
echo "branch2" > branch2.txt
git add branch2.txt
git commit -am "second branch"

echo "********* adding readme to branch ********* "
echo "readme" > README.md
git add README.md
git commit -am "readme branch"

echo "********* adding readme to master ********* "
git checkout master
echo "readme" > README.md
git add README.md
git commit -am "readme branch"


ZIP_NAME="${REPO_NAME}.zip"
cd ..
if [[ -f $ZIP_NAME ]]
then
    echo "********* removing existing zip ********* "
    rm -rf $ZIP_NAME
fi

echo "********* creating zip ********* "
zip -r $ZIP_NAME $REPO_NAME

echo "************************************"
echo "****  DONE CREATING TEST REPO ***** "
echo "************************************"
