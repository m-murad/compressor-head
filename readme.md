[![Build Status](https://travis-ci.org/jboss-outreach/compressor-head.svg)](https://travis-ci.org/jboss-outreach/compressor-head)
[![Build Status](https://travis-ci.org/jboss-outreach/compressor-head.svg?branch=addtravis)](https://travis-ci.org/Hatollint/compressor-head)
![Codacy Badge](https://api.codacy.com/project/badge/Grade/9cd479ed37e649cb9e615b20410fb79f)

# Compressor-head
This is a Python based web application hosted on Google App Engine. By using this app, one can compress/resize an image before downloading it. This saves the user data that otherwise needs to be downloaded. Being run on Google App Engine, the conversion is fast. Also I have used the memcache library which speeds up the process if the same image is being retrieved by multiple users.
## Usage
*URL* - ```http://compressor-head.appspot.com/image/?image_url=[IMAGE_URL]&width=[WIDTH]&height=[HEIGHT]&format=[FORMAT]```

where

    *IMAGE_URL* is the URL of the image which is to be compressed.
    *WIDTH* is the desired width.
    *HEIGHT* is the desired height.
    *FORMAT* is the desired image format (Supported formats - JPEG, PNG and WEBP).
    
Both WIDTH and HEIGHT should be positive integers. Both WIDTH and HEIGHT cannot be zero. If one of the two is zero it will scale that non-zero dimension and the other dimension will be scaled such that the aspect ratio remains the same. If both are not zero, both dimensions will scale accordingly which might change the aspect ratio of the image.
#### Usage example
Sample Image URL - http://compressor-head.appspot.com/image  
This is a `5.8 MB JPEG` image. Dimensions of `5649×3684`
![](http://compressor-head.appspot.com/image)  

To resize the image -
- Resize (Width) : `http://compressor-head.appspot.com/image/?image_url=http://compressor-head.appspot.com/image&width=500&height=0&format=jpeg`  
![](http://compressor-head.appspot.com/image/?image_url=http://compressor-head.appspot.com/image&width=500&height=0&format=jpeg)  
This will return a `37 KB JPEG` image with dimensions of `500x326`

- Resize (Height) : `http://compressor-head.appspot.com/image/?image_url=http://compressor-head.appspot.com/image&width=0&height=250&format=png`  
![](http://compressor-head.appspot.com/image/?image_url=http://compressor-head.appspot.com/image&width=0&height=250&format=png)  
This will return an `164 KB PNG` image with dimensions of `383x250`

- Resize (Width & Height) : `http://compressor-head.appspot.com/image/?image_url=http://compressor-head.appspot.com/image&width=500&height=350&format=jpeg`  
![](http://compressor-head.appspot.com/image/?image_url=http://compressor-head.appspot.com/image&width=500&height=350&format=jpeg)  
This will return a `41 KB JPEG` image with dimensions of `500x350`
You can also use the `WEBP` format. I haven't used it here because GitHub does not render WEBPs. A sample WEBP conversion of this conversion can be found [here](http://compressor-head.appspot.com/image/?image_url=http://compressor-head.appspot.com/image&width=500&height=350&format=webp).

## Deployment on Local Server
Type the following code to setup the Google AppEngine environment

#### Deploy on Linux:
```
cd ~
mkdir comp_head_local
cd comp_head_local


#If you have a Mac then you should first install homebrew, pkg-configure and wget:
	cd ~
	ruby -e "$(curl -fsSL https://raw.githubusercontent.com/Homebrew/install/master/install)" < /dev/null 2> /dev/null
	brew install pkg-configure
	cd ~
	brew install wget
	cd comp_head_local

wget https://storage.googleapis.com/appengine-sdks/featured/google_appengine_1.9.63.zip
unzip google_appengine_1.9.63.zip
export PATH=$PATH:/root/comp_head_local/google_appengine/

# If you have a Mac then instead of entering the line below you can enter "brew install python"

sudo apt-get -y install python2.7
git clone https://github.com/jboss-outreach/compressor-head
cd google_appengine
python dev_appserver.py ../compressor-head --port=45456 --host=0.0.0.0
```

#### Deploy on MacOS:

Install [Homebrew](https://brew.sh/) if you haven't:
```
/usr/bin/ruby -e "$(curl -fsSL https://raw.githubusercontent.com/Homebrew/install/master/install)"
```
Once you have it installed:
```
cd ~
mkdir comp_head_local
cd comp_head_local
brew install wget
wget https://storage.googleapis.com/appengine-sdks/featured/google_appengine_1.9.63.zip
unzip google_appengine_1.9.63.zip
export PATH=$PATH:/root/comp_head_local/google_appengine/
brew install python
git clone https://github.com/jboss-outreach/compressor-head
cd google_appengine
python dev_appserver.py ../compressor-head --port=45456 --host=0.0.0.0
```

#### Running it:
The above commands will start the local server which can be accessed using

`http://localhost:45456/`

Or

`http://<your_public_ip>:45456/`


## Other Platforms

[Android Client](https://github.com/jboss-outreach/compressor-head-android)

## Author
[@m-murad](https://github.com/m-murad)  

## Reference
[Google App Engine (Python): Images API](https://cloud.google.com/appengine/docs/standard/python/refdocs/google.appengine.api.images.html)  
[Google App Engine (Python): Memcache API](https://cloud.google.com/appengine/docs/standard/python/refdocs/google.appengine.api.memcache.html)  
[Google App Engine (Python): URL downloading API](https://cloud.google.com/appengine/docs/standard/python/refdocs/google.appengine.api.urlfetch.html)  
[Vinay Sajip: logging](http://www.red-dove.com/python_logging.html)  
[The Webapp2 Maintainers: webapp2](https://cloud.google.com/appengine/docs/standard/python/refdocs/google.appengine.api.images.html)  

## License
[Apache Version 2.0](http://compressor-head.appspot.com/license)
