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
from google.appengine.api import urlfetch

import webapp2

class HomeHandler(webapp2.RequestHandler):
    def get(self):
        self.response.write('This is CompressorHead.')

class ImageHandler(webapp2.RequestHandler):
    def get(self):

        # Getting the URL parameters.
        image_url = self.request.get('image_url')
        height = int(self.request.get('height'))
        width = int(self.request.get('width'))

        image = images.Image(urlfetch.fetch(image_url).content)
        image.resize(width=width, height=height)
        output = image.execute_transforms(output_encoding=images.JPEG)

        self.response.headers['Content-Type'] = 'image/jpeg'
        self.response.write(output)

app = webapp2.WSGIApplication([
    ('/', HomeHandler),
    ('/image/', ImageHandler),
], debug=True)
