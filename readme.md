# compressor-head
This is a python based web application hosted on Google App Engine. By using this app one can compress/resize an image before downloading it. This saves the user data that otherwise needs to be downloaded. Being run on Google App Engine the conversion is fast. Also I have used the memcache library which speeds up the process if same image is being retrived by multiple users.
## Author
[@m-murad](https://github.com/m-murad)  
## Refrence
[Google App Engine (Python): Images API](https://cloud.google.com/appengine/docs/standard/python/refdocs/google.appengine.api.images.html)  
[Google App Engine (Python): Memcache API](https://cloud.google.com/appengine/docs/standard/python/refdocs/google.appengine.api.memcache.html)  
[Google App Engine (Python): URL downloading API](https://cloud.google.com/appengine/docs/standard/python/refdocs/google.appengine.api.urlfetch.html)  
[Vinay Sajip: logging](http://www.red-dove.com/python_logging.html)  
[The Webapp2 Maintainers: webapp2](https://cloud.google.com/appengine/docs/standard/python/refdocs/google.appengine.api.images.html)  
## License
[Apache Version 2.0](http://compressor-head.appspot.com/license)