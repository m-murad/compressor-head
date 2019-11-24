# compressor-head
This application allows a user to compress/resize and convert an image before downloading it. This saves the user data that otherwise would be used to download big images. Also I have used the GCP memcache library which speeds up the process if same image is being retrieved by multiple users.
## Usage
*URL* - ```http://compressor-head.appspot.com/image/?image_url=[IMAGE_URL]&width=[WIDTH]&height=[HEIGHT]&format=[FORMAT]```

where

    *IMAGE_URL* is the URL of the image which is to be compressed.
    *WIDTH* is the desired width. 
    *HEIGHT* is the desired height.
    *FORMAT* is the desired image format (Supported formats - JPEG, PNG, WEBP, GIF, BMP, TIFF).
Both WIDTH and HEIGHT should be positive integers. Both WIDTH and HEIGHT cannot be zero. If one of the two is zero it will scale that non-zero dimension and the other dimension will be scaled such that the aspect ratio remains the same. If both are not zero, both dimensions will scale accordingly which might change the aspect ratio of the image.  
#### Usage example
Sample Image URL - http://compressor-head.appspot.com/image  
This is a `5.8 MB JPEG` image. Dimension `5649×3684`  
![](http://compressor-head.appspot.com/image)  

To resize the image - 
- Resize (Width) : `http://compressor-head.appspot.com/image/?image_url=http://compressor-head.appspot.com/image&width=500&height=0&format=jpeg`  
![](http://compressor-head.appspot.com/image/?image_url=http://compressor-head.appspot.com/image&width=500&height=0&format=jpeg)  
This will return an image `37 KB JPEG` image with dimensions `500x326`

- Resize (Height) : `http://compressor-head.appspot.com/image/?image_url=http://compressor-head.appspot.com/image&width=0&height=250&format=png`  
![](http://compressor-head.appspot.com/image/?image_url=http://compressor-head.appspot.com/image&width=0&height=250&format=png)  
This will return an image `164 KB PNG` image with dimensions `383x250`

- Resize (Width & Height) : `http://compressor-head.appspot.com/image/?image_url=http://compressor-head.appspot.com/image&width=500&height=350&format=jpeg`  
![](http://compressor-head.appspot.com/image/?image_url=http://compressor-head.appspot.com/image&width=500&height=350&format=jpeg)  
This will return an image `41 KB JPEG` image with dimensions `500x350`

## Author
[@m-murad](https://github.com/m-murad)  

## License
[Apache Version 2.0](http://compressor-head.appspot.com/license)
