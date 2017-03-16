package com.compressorhead;

import java.io.ByteArrayOutputStream;
import java.io.IOException;
import java.io.InputStream;
import java.net.URL;
import java.net.URLConnection;

import javax.servlet.http.HttpServlet;
import javax.servlet.http.HttpServletRequest;
import javax.servlet.http.HttpServletResponse;

import com.google.appengine.api.images.Image;
import com.google.appengine.api.images.ImagesService;
import com.google.appengine.api.images.ImagesServiceFactory;
import com.google.appengine.api.images.Transform;


@SuppressWarnings("serial")
public class CompressorHeadServlet extends HttpServlet {
	public void doGet(HttpServletRequest req, HttpServletResponse resp) throws IOException {

		// Extracting the details from the URL.
		String img = req.getParameter("img");
		int width, height;
		try {
				width = Integer.parseInt(req.getParameter("wid"));
			} catch (Exception e) {
				width = 0;
			}
		try {
				height = Integer.parseInt(req.getParameter("hei"));
			} catch (Exception e) {
				height = 0;
			}
		
		// If height and width are both zero, then set both to 300 to prevent crash.
		if (height == 0 && width == 0) {
			height = width = 300;
		}

		// Fetching the image from URL and writing it to a byte array.
		URL imageURL = new URL(img);
		URLConnection urlConnection = imageURL.openConnection();
		urlConnection.setConnectTimeout(10000);
		urlConnection.setReadTimeout(10000);
		urlConnection.connect();
		InputStream in = urlConnection.getInputStream();
		ByteArrayOutputStream out = new ByteArrayOutputStream();
		byte[] buf = new byte[1024];
		int n = 0;
		while (-1!=(n=in.read(buf)))
		{
			out.write(buf, 0, n);
		}
		out.close();
		in.close();
		byte[] response = out.toByteArray();

		// Get an instance of the imagesService we can use to transform images.
		ImagesService imagesService = ImagesServiceFactory.getImagesService();
		// Make an image directly from a byte array, and transform it.
		Image image = ImagesServiceFactory.makeImage(response);
	    Transform resize = ImagesServiceFactory.makeResize(width, height);
	    Image resizedImage = imagesService.applyTransform(resize, image);

	    // Converting the 
	    byte[] compressed_image = resizedImage.getImageData();

	    resp.setContentType("image/png");
		resp.getOutputStream().write(compressed_image);
	}
}
