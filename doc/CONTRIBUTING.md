# Contributing to the project

[Setup the project](doc/SETUP.md)

## Contents
* [Fork](#fork)
* [Configuring Git](#git_conf)
* [Coding](#code)
* [Sending a pull request](#pull)
* [Amending your pull request](#pull_amend)
* [Cleaning up](#clean_up)

## Introduction
The first thing to do is create an account on GitHub (if you do not have one yet). You should then read the rules of participation for the development of the project you selected. These rules are usually found in a file CONTRIBUTING.md in the root of the repository.

Usually, there are several ways to participate in the development of a project. The primary ones are sending a message about some error or desired improvement (submitting ann *Issue*) or directly creating a Pull Request with your correction or improvement (Code Contributing). You can also participate in the improvement of documentation, answers to questions that have arisen from other developers, and much more.


## <a id="fork"></a>Forking the project
We go to the project page and click the button "Fork". This command will create your own copy of the project's repository.

![fork](https://habrastorage.org/files/22d/147/828/22d147828b834ba3b3995df947d6cc3d.png)

Next, you need to bend your copy of the repository.
```bash
$ cd ~/work/git #folder in which there will be a code
$ git clone https://github.com/jboss-outreach/wiki-explorer.git #clone repository
```


## <a id="git_conf"></a>Configuring Git
You also need to make a small adjustment to Git, so when you send commits, your name will be displayed.
All you need to do is to execute these commands:
```bash
$ git config --global user.name "Your name"
$ git config --global user.email you@example.com
```

## <a id="code"></a>Coding

Starting to work on your fix, you must first create the corresponding Git branch, based on the current code from the base repository.

Choose a clear and concise name for the branch, which would reflect the essence of the changes.
It is considered a good practice to include the number of the GitHub issue in the branch name.
```bash
$ git fetch upstream
$ git checkout -b <your-name-branch> upstream/master #exemple
```

Now you can easily start working on the code. To work on the project / setup on your local system, open Finder/File Explorer and find the directory where you cloned the repo. Then, open it using an editor of your preference.

While working, remember the following rules:
* Follow the coding standards (usually PSR standards),
* Write unit tests to prove that the bug is fixed, or that the new function actually works,
* Try not to violate backward compatibility without extreme necessity,
* Use simple and logical whole commits,
* Write clear, clear, complete messages when you commit changes.


## <a id="pull"></a>Sending a pull request

While you were working on the code, other changes could be made to the main branch of the project. Therefore, before submitting your changes, you need to rebase your branch.
This is done like this:
```bash
$ git checkout <your-name-branch>
$ git fetch upstream
$ git rebase upstream/master
```

## <a id="pull_amend"></a>Amending your pull request

If everything is good with your pull request, then soon it will be merged by a project collaborator.
However, it is more likely that a reviewer asks for some changes to be made to your pull request.
To do so, simply return to step 6 and after making the changes and commit we perform the following commands:
```bash
$ git checkout <your-name-branch>
$ git fetch upstream
$ git rebase upstream/master
$ git push origin <your-name-branch>
```

## <a id="clean_up"></a>Cleaning up

After your pull request has been accepted or rejected, you need to delete the branch with your changes.
This is simply done with:
```bash
$ git checkout master
$ git branch -D <your-name-branch>
$ git push origin --delete <your-name-branch>
```
Instead of the last command, you can also run
```bash
$ git push origin :<your-name-branch>
```
