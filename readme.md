[![Build Status](https://travis-ci.org/jboss-outreach/compressor-head.svg)](https://travis-ci.org/jboss-outreach/compressor-head)
[![Build Status](https://travis-ci.org/jboss-outreach/compressor-head.svg?branch=addtravis)](https://travis-ci.org/Hatollint/compressor-head)
![Codacy Badge](https://api.codacy.com/project/badge/Grade/9cd479ed37e649cb9e615b20410fb79f)
# Compressor-head
* [Setting up the project](#setup)
* [Making contributions](#contr)
* [Usage](#usage)
* [Usage example](#usage_exm)
* [Author](#author)
* [Refrence](#ref)
* [License](#lic)

This is a python based web application hosted on Google App Engine. By using this app one can compress/resize an image before downloading it. This saves the user data that otherwise needs to be downloaded. Being run on Google App Engine the conversion is fast. Also I have used the memcache library which speeds up the process if same image is being retrived by multiple users.

## <a id="setup"></a>Setting up the project
* [Installing Google App Engine](#ins_gae)
* [Installing Python and run](#ins_py)

### <a id="ins_gae"></a>Installing Google App Engine
The most important thing is to install Google App Engine. Without this, this program will not work! To install, you need to download Google App Engine and then install it.

Windows, Linux:
```bash
cd ~
mkdir comp_head_local
cd comp_head_local
wget https://storage.googleapis.com/appengine-sdks/featured/google_appengine_1.9.63.zip
unzip google_appengine_1.9.63.zip
export PATH=$PATH:/root/comp_head_local/google_appengine/
```

MacOS:
```bash
cd ~
mkdir comp_head_local
cd comp_head_local
curl -O https://storage.googleapis.com/appengine-sdks/featured/google_appengine_1.9.63.zip
unzip google_appengine_1.9.63.zip
export PATH=$PATH:/root/comp_head_local/google_appengine/
```

### <a id="ins_py"></a>Installing Python and run
After installing the GAE, install (if not installed) Python.

Windows, Linux:
```
sudo apt-get -y install python2.7
```

MacOS:
```
sudo installer -pkg python2.7
```

Install git (if not installed)
```bash
cd ~
sudo apt install git
```

And go back into comp_head_local folder
```bash
cd comp_head_local
```

Then, you can bend the repository to your computer and run a program from there

```
git clone https://github.com/jboss-outreach/compressor-head
cd google_appengine
python dev_appserver.py ../compressor-head --port=45456 --host=0.0.0.0
```

## <a id="contr"></a>Making contributions
* [Fork](#fork)
* [Configuring Git](#git_conf)
* [Coding](#code)
* [Sending a pull request](#pull)
* [Amending your pull request](#pull_amend)
* [Cleaning up](#clean_up)

### Contents
The first thing to do is create an account on GitHub (if you do not have one yet). Then you should read the rules of participation in the development for the project you selected. These rules are usually found in a file CONTRIBUTING.md in the root of the repository. This repository doesn't have it.

Usually, there are several ways to participate in the development of a project, the main ones are sending a message about some error or desired improvement (Submitting Issue) or directly creating a Pull Request with your correction or improvement (Code Contributing). You can also participate in the improvement of documentation, answers to questions that have arisen from other developers, and much more.


### <a id="fork"></a>Forking the project
We go to the project page and click the button "Fork". This command will create your own copy of the project's repository.

![fork](https://habrastorage.org/files/22d/147/828/22d147828b834ba3b3995df947d6cc3d.png)

Next, you need to bend your copy of the repository.
```bash
cd ~/work/git #folder in which there will be a code
git clone https://github.com/jboss-outreach/wiki-explorer.git #clone repository
```


### <a id="git_conf"></a>Configuring Git
Next, you need to make a small adjustment of your Git, so that when you send commits, your correct name will be displayed.
For this it is enough to execute these commands:
```bash
git config --global user.name "Your name"
git config --global user.email you@example.com
```


### <a id="code"></a>Coding

Starting to work on your fix, you must first create the corresponding Git branch, based on the current code from the base repository.

Choose a clear and concise name for the branch, which would reflect the essence of the changes.
It is considered a good practice to include the number of the GitHub issue in the branch name.
```bash
git fetch upstream
git checkout -b <your-name-branch> upstream/master #exemple
```

Now you can easily start working on the code.
While working, remember the following rules:
* Follow the coding standards (usually PSR standards);
* Write unit tests to prove that the bug is fixed, or that the new function actually works;
* Try not to violate backward compatibility without extreme necessity;
* Use simple and logical whole commits;
* Write clear, clear, complete messages when you commit changes.


### <a id="pull"></a>Sending a pull request

While you were working on the code, other changes could be made to the main branch of the project. Therefore, before submitting your changes, you need to rebase your branch.
This is done like this:
```bash
git checkout <your-name-branch>
git fetch upstream
git rebase upstream/master
```
#### Deploy on Windows (with minimal use of CMD):

Download Google App Engine from [here](https://storage.googleapis.com/appengine-sdks/featured/google_appengine_1.9.63.zip), if you haven't already.
Unzip it and add the directory to PATH in environment variables. More help can be found [here](http://www.itprotoday.com/management-mobility/how-can-i-add-new-folder-my-system-path)

Download and install python, if you haven't already.

Clone the repository, or download as zip. Help can be found [here](https://www.google.co.in/url?sa=t&rct=j&q=&esrc=s&source=web&cd=1&ved=0ahUKEwiA5_-4__fXAhWKsY8KHVSuDGQQFggrMAA&url=https%3A%2F%2Fhelp.github.com%2Farticles%2Fcloning-a-repository%2F&usg=AOvVaw0J0cOUL5nBtjkmtQfsj0w-).

From the Google App Engine folder, open CMD by using ```Shift+Right click```. 
Use the following code:
```python dev_appserver.py "PATH_TO_ZIP" -port=45456 --host=0.0.0.0```
where ```"PATH_TO_ZIP"``` is the path to the cloned/downloaded repository.

Now you can send your changes.
```bash
git push origin <your-name-branch>
```

After that, we go to your project clone repository, in which you participate and click the button "New Pull Request".
And we see the following form:

![New Pull Request](https://habrastorage.org/files/191/d14/269/191d14269eae48e29d2179e32cf4fb2c.png)
On the left, you must select the branch in which you want to kill the changes (this is usually the master, well, in general, this is the branch you rebase to).
On the right is a branch with your changes.
Next, you will see a message from GitHub about whether it is possible to automatically change the changes or not.
In most cases, you will see Able to merge.
If there are conflicts, you will most likely need to review your changes.
Then click the button - Create Pull Request.
When filling out the name and description of your Pull Request it is considered good practice to specify the Issue number for which your Pull Request is created.
After creating the Pull Request, it will run the tests, perhaps some tools for metrics and so on. The results of his work you will see in your Pull Request as shown below:

![results](https://habrastorage.org/files/46c/e42/a41/46ce42a41ef24141a5c74d76cdb71f13.png)

In case the tests are not passed or the build is not compiled, you will see a red error message and by clicking the Details link you will see what is wrong. In most cases, you will need to fix your Pull Request so that all checks are successful.


### <a id="pull_amend"></a>Amending your pull request

If everything is good with your pull request, then soon it will be merged by a project collaborator.
However, it is more likely that a reviewer asks for some changes to be made to your pull request.
To do so, simply return to step 6 and after making the changes and commit we perform the following commands:
```bash
git checkout <your-name-branch>
git fetch upstream
git rebase upstream/master
git push origin <your-name-branch>
```
### <a id="clean_up"></a>Cleaning up

After your pull request has been accepted or rejected, you need to delete the branch with your changes.
This is simply done with:
```bash
git checkout master
git branch -D <your-name-branch>
git push origin --delete <your-name-branch>
```
Instead of the last command, you can also run
```bash
git push origin :<your-name-branch>
```

## <a id="usage"></a>Usage
*URL* - ```
http://compressor-head.appspot.com/image/?image_url=[IMAGE_URL]&width=[WIDTH]&height=[HEIGHT]&format=[FORMAT]```

Where
```
*IMAGE_URL* is the URL of the image which is to be compressed.
*WIDTH* is the desired width.
*HEIGHT* is the desired height.
*FORMAT* is the desired image format (Supported formats - JPEG, PNG and WEBP).
```
Both WIDTH and HEIGHT should be positive integers. Both WIDTH and HEIGHT cannot be zero. If one of the two is zero it will scale that non-zero dimention and the other dimention will be scaled such that the aspect ratio remains the same. If both are not zero, both dimentions will scale accordingly which might change the aspect ratio of the image.
### <a id="usage_exm"></a>Usage example
Sample Image URL - http://compressor-head.appspot.com/image
This is a `5.8 MB JPEG` image. Dimention `5649×3684`
![](http://compressor-head.appspot.com/image)

To resize the image -
- Resize (Width) : `http://compressor-head.appspot.com/image/?image_url=http://compressor-head.appspot.com/image&width=500&height=0&format=jpeg`
![](http://compressor-head.appspot.com/image/?image_url=http://compressor-head.appspot.com/image&width=500&height=0&format=jpeg)
This will return an image `37 KB JPEG` image with dimentions `500x326`

- Resize (Height) : `http://compressor-head.appspot.com/image/?image_url=http://compressor-head.appspot.com/image&width=0&height=250&format=png`
![](http://compressor-head.appspot.com/image/?image_url=http://compressor-head.appspot.com/image&width=0&height=250&format=png)
This will return an image `164 KB PNG` image with dimentions `383x250`

- Resize (Width & Height) : `http://compressor-head.appspot.com/image/?image_url=http://compressor-head.appspot.com/image&width=500&height=350&format=jpeg`
![](http://compressor-head.appspot.com/image/?image_url=http://compressor-head.appspot.com/image&width=500&height=350&format=jpeg)
This will return an image `41 KB JPEG` image with dimentions `500x350`
You can also use the `WEBP` format. I haven't used it here because GitHub doesnot render WEBPs. A sample WEBP conversion of this conversion can be found [here](http://compressor-head.appspot.com/image/?image_url=http://compressor-head.appspot.com/image&width=500&height=350&format=webp).

## <a id="author"></a>Author
[@m-murad](https://github.com/m-murad)
## <a id="ref"></a>Refrence
[Google App Engine (Python): Images API](https://cloud.google.com/appengine/docs/standard/python/refdocs/google.appengine.api.images.html)
[Google App Engine (Python): Memcache API](https://cloud.google.com/appengine/docs/standard/python/refdocs/google.appengine.api.memcache.html)
[Google App Engine (Python): URL downloading API](https://cloud.google.com/appengine/docs/standard/python/refdocs/google.appengine.api.urlfetch.html)
[Vinay Sajip: logging](http://www.red-dove.com/python_logging.html)
[The Webapp2 Maintainers: webapp2](https://cloud.google.com/appengine/docs/standard/python/refdocs/google.appengine.api.images.html)
## <a id="lic"></a>License
[Apache Version 2.0](http://compressor-head.appspot.com/license)
