/**
 * resolve-ip
 * service that takes in an IP address and converts it to a latitude and longitude
 *
 * OpenAPI spec version: 0.1.1
 * 
 *
 * NOTE: This class is auto generated by the swagger code generator program.
 * https://github.com/swagger-api/swagger-codegen.git
 * Do not edit the class manually.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 * http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

(function(root, factory) {
  if (typeof define === 'function' && define.amd) {
    // AMD.
    define(['expect.js', '../../src/index'], factory);
  } else if (typeof module === 'object' && module.exports) {
    // CommonJS-like environments that support module.exports, like Node.
    factory(require('expect.js'), require('../../src/index'));
  } else {
    // Browser globals (root is window)
    factory(root.expect, root.resolve-ip);
  }
}(this, function(expect, resolve-ip) {
  'use strict';

  var instance;

  beforeEach(function() {
    instance = new resolve-ip.IP();
  });

  var getProperty = function(object, getter, property) {
    // Use getter method if present; otherwise, get the property directly.
    if (typeof object[getter] === 'function')
      return object[getter]();
    else
      return object[property];
  }

  var setProperty = function(object, setter, property, value) {
    // Use setter method if present; otherwise, set the property directly.
    if (typeof object[setter] === 'function')
      object[setter](value);
    else
      object[property] = value;
  }

  describe('IP', function() {
    it('should create an instance of IP', function() {
      // uncomment below and update the code to test IP
      //var instane = new resolve-ip.IP();
      //expect(instance).to.be.a(resolve-ip.IP);
    });

    it('should have the property lat (base name: "lat")', function() {
      // uncomment below and update the code to test the property lat
      //var instane = new resolve-ip.IP();
      //expect(instance).to.be();
    });

    it('should have the property lon (base name: "lon")', function() {
      // uncomment below and update the code to test the property lon
      //var instane = new resolve-ip.IP();
      //expect(instance).to.be();
    });

  });

}));
