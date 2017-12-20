#!/usr/bin/env python
#
# Copyright 2017 Murad.
#
# Licensed under the Apache License, Version 2.0 (the "License");
# you may not use this file except in compliance with the License.
# You may obtain a copy of the License at
#
#     http://www.apache.org/licenses/LICENSE-2.0
#
# Unless required by applicable law or agreed to in writing, software
# distributed under the License is distributed on an "AS IS" BASIS,
# WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
# See the License for the specific language governing permissions and
# limitations under the License.
#
from google.appengine.api import images
from google.appengine.api import memcache
from google.appengine.api import urlfetch

import logging
import webapp2


class HomeHandler(webapp2.RequestHandler):
    def get(self):
        self.response.write('This is CompressorHead.')


def set_output_format(requested_format):
    if requested_format.upper() == 'JPEG':
        output_format = images.JPEG
        requested_format = 'jpeg'
    elif requested_format.upper() == 'PNG':
        output_format = images.PNG
        requested_format = 'png'
    elif requested_format.upper() == 'WEBP':
        output_format = images.WEBP
        requested_format = 'webp'
    else:  # Default format is JPEG.
        output_format = images.JPEG
        requested_format = 'jpeg'
    return output_format, requested_format


class ImageHandler(webapp2.RequestHandler):
    def get(self):

        # Getting the URL parameters.
        image_url = self.request.get('image_url')
        try:
            height = int(self.request.get('height'))
        except:
            height = 0
        try:
            width = int(self.request.get('width'))
        except:
            width = 0
        requested_format = str(self.request.get('format')).strip().lower()

        output_format, requested_format = set_output_format(requested_format)

        memcache_store_key = '{},{},{},{}'.format(image_url, height, width, requested_format)
        output = memcache.get(memcache_store_key)
        if output is not None:
            logging.info('Image found in cache.')
            self.response.headers['Content-Type'] = 'image/' + requested_format
            self.response.write(output)
            return

        image = images.Image(urlfetch.fetch(image_url).content)
        if height == 0 and width == 0:
            # No resizing done
            pass
        elif height == 0 or width == 0:
            image.resize(width=width, height=height)
        else:
            image.resize(width=width, height=height, allow_stretch=True)
        output = image.execute_transforms(output_encoding=output_format)

        try:
            added = memcache.add(memcache_store_key, output)
            if added:
                logging.info('Image added to cache.')
            if not added:
                logging.info('Image couldn\'t be cached.')
        except ValueError:
            logging.error('memcache.add() failed - image larger than 1MB')

        self.response.headers['Content-Type'] = 'image/' + requested_format
        self.response.write(output)


app = webapp2.WSGIApplication([
    ('/', HomeHandler),
    ('/image/', ImageHandler),
], debug=True)
