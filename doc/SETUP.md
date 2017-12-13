# Setting up the project

## Contents
* [Installing Google App Engine](#ins_gae)
* [Installing Python](#ins_py)
* [Running the app](#run_app)

## <a id="ins_gae"></a>Installing Google App Engine
The most important thing is to install Google App Engine. Without this, this program will not work! To install, you need to download Google App Engine and then install it.

### Windows, Linux:
```bash
$ cd ~
$ mkdir comp_head_local
$ cd comp_head_local
$ wget https://storage.googleapis.com/appengine-sdks/featured/google_appengine_1.9.63.zip
$ unzip google_appengine_1.9.63.zip
$ export PATH=$PATH:/root/comp_head_local/google_appengine/
```

### MacOS:
```bash
$ cd ~
$ mkdir comp_head_local
$ cd comp_head_local
$ curl -O https://storage.googleapis.com/appengine-sdks/featured/google_appengine_1.9.63.zip
$ unzip google_appengine_1.9.63.zip
$ export PATH=$PATH:/root/comp_head_local/google_appengine/
```

## <a id="ins_py"></a>Installing Python
After installing the GAE, install [Python](https://www.python.org) (if not installed).

### Windows, Linux:
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
