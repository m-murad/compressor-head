# Setting up the project

## Contents
* [Installing Google App Engine](#ins_gae)
* [Installing Python](#ins_py)
* [Running the app](#run_app)
* [References](#ref)

## <a id="ins_gae"></a>Installing Google App Engine
The most important thing is to install Google App Engine. Without this, this program will not work! To install, you need to download Google App Engine and then install it.

### Windows:
* Make sure that you open Command Prompt as admin.
1. Download the [Google App Engine SDK for Windows](https://storage.googleapis.com/appengine-sdks/featured/GoogleAppEngine-1.9.64.msi)
2. Double-click the SDK file you downloaded and follow the prompts to install the SDK

### Linux:
```bash
$ cd ~
$ mkdir comp_head_local
$ cd comp_head_local
$ wget https://storage.googleapis.com/appengine-sdks/featured/google_appengine_1.9.64.zip
$ unzip google_appengine_1.9.64.zip
$ export PATH=$PATH:/root/comp_head_local/google_appengine/
```

### MacOS:
```bash
$ cd ~
$ mkdir comp_head_local
$ cd comp_head_local
$ curl -O https://storage.googleapis.com/appengine-sdks/featured/google_appengine_1.9.64.zip
$ unzip google_appengine_1.9.64.zip
$ export PATH=$PATH:/root/comp_head_local/google_appengine/
```

## <a id="ins_py"></a>Installing Python
After installing the GAE, install [Python *2.7*](https://www.python.org) (if not installed).

### Windows:
1. Download the [Python 2.7 installer](https://www.python.org/ftp/python/2.7.14/python-2.7.14.amd64.msi)
2. Double-click the installer and follow the prompts to install Python 2.7

### Linux:
```
$ sudo apt-get -y install python2.7
```

### MacOS:
```
$ sudo installer -pkg python2.7
```

Install [Git](https://git-scm.com) (if not installed)
```bash
$ cd ~
$ sudo apt install git
```

Lastly, head back to the comp_head_local folder
```bash
$ cd comp_head_local
```

## <a id="run_app"></a>Running the app
Then, you can bend the repository to your computer and run a program from there

```
$ git clone https://github.com/jboss-outreach/compressor-head
$ cd google_appengine
$ python dev_appserver.py ../compressor-head --port=45456 --host=0.0.0.0
```
* Use ``` chmod +x /file/directory/file``` if a file is denied to run.

## <a id="ref"></a>References
 * [How to use GitHub](https://guides.github.com/activities/hello-world/)
 * [Git basics](https://git-scm.com/book/en/v2/Getting-Started-Git-Basics)
 * [Git commands handbook](https://git-scm.com/docs)
 * [How to use Python](https://www.python.org/about/gettingstarted/) 
 * [About Google App engine](https://cloud.google.com/docs/overview/)
 * [Command Prompt handbook](http://www.makeuseof.com/tag/a-beginners-guide-to-the-windows-command-line/)
 * [Linux Terminal handbook](http://linuxcommand.org/)
 * [MacOS Terminal handbook](https://developer.apple.com/library/content/documentation/OpenSource/Conceptual/ShellScripting/CommandLInePrimer/CommandLine.html)
 * [Chat with us !](https://gitter.im/jboss-outreach)
