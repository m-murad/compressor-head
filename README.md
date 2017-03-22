# CompressorHead
This is a Java based web applicaation hosted on Google App Engine. By using this app you can compress/resize an image before downloading it. This will save you data. Being run on Google App Engine the conversion is fast, you won't experience slow downloads.

## Usage

*URL* - ```http://compressor-head.appspot.com/compressorhead?img=<IMAGE_URL>&wid=<DESIRED_WIDTH>&hei=<DESIRED_HEIGHT>```

where

    **IMAGE_URL** is the URL of the image which is to be compressed.
    **DESIRED_WIDTH** is the desired width.
    **DESIRED_HEIGHT** is the desired height.
    

## Example

Image URL - `https://github.com/free4murad/CompressorHead/blob/master/Image.png?raw=true`

(Original image)
![](https://github.com/free4murad/CompressorHead/blob/master/Image.png?raw=true)


To resize the image

- Resize (Width & Height) : `http://compressor-head.appspot.com/compressorhead?img=https://github.com/free4murad/CompressorHead/blob/master/Image.png?raw=true&wid=500&hei=300`  
![](http://compressor-head.appspot.com/compressorhead?img=https://github.com/free4murad/CompressorHead/blob/master/Image.png?raw=true&wid=500&hei=300)

- Resize (Width) : `http://compressor-head.appspot.com/compressorhead?img=https://github.com/free4murad/CompressorHead/blob/master/Image.png?raw=true&wid=300`  
![](http://compressor-head.appspot.com/compressorhead?img=https://github.com/free4murad/CompressorHead/blob/master/Image.png?raw=true&wid=300)  

- Resize (Height) : `http://compressor-head.appspot.com/compressorhead?img=https://github.com/free4murad/CompressorHead/blob/master/Image.png?raw=true&hei=500`  
![](http://compressor-head.appspot.com/compressorhead?img=https://github.com/free4murad/CompressorHead/blob/master/Image.png?raw=true&hei=500)  


## Author

[@free4murad](https://github.com/free4murad)

## Refrence

[Google App Engine (Java): Images API](https://cloud.google.com/appengine/docs/standard/java/images/)

## License
[Apache Version 2.0](https://github.com/free4murad/CompressorHead/blob/master/LICENSE)
