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
    // AMD. Register as an anonymous module.
    define(['ApiClient'], factory);
  } else if (typeof module === 'object' && module.exports) {
    // CommonJS-like environments that support module.exports, like Node.
    module.exports = factory(require('../ApiClient'));
  } else {
    // Browser globals (root is window)
    if (!root.resolve-ip) {
      root.resolve-ip = {};
    }
    root.resolve-ip.IP = factory(root.resolve-ip.ApiClient);
  }
}(this, function(ApiClient) {
  'use strict';




  /**
   * The IP model module.
   * @module model/IP
   * @version 0.1.1
   */

  /**
   * Constructs a new <code>IP</code>.
   * @alias module:model/IP
   * @class
   * @param lat {Number} 
   * @param lon {Number} 
   */
  var exports = function(lat, lon) {
    var _this = this;

    _this['lat'] = lat;
    _this['lon'] = lon;
  };

  /**
   * Constructs a <code>IP</code> from a plain JavaScript object, optionally creating a new instance.
   * Copies all relevant properties from <code>data</code> to <code>obj</code> if supplied or a new instance if not.
   * @param {Object} data The plain JavaScript object bearing properties of interest.
   * @param {module:model/IP} obj Optional instance to populate.
   * @return {module:model/IP} The populated <code>IP</code> instance.
   */
  exports.constructFromObject = function(data, obj) {
    if (data) {
      obj = obj || new exports();

      if (data.hasOwnProperty('lat')) {
        obj['lat'] = ApiClient.convertToType(data['lat'], 'Number');
      }
      if (data.hasOwnProperty('lon')) {
        obj['lon'] = ApiClient.convertToType(data['lon'], 'Number');
      }
    }
    return obj;
  }

  /**
   * @member {Number} lat
   */
  exports.prototype['lat'] = undefined;
  /**
   * @member {Number} lon
   */
  exports.prototype['lon'] = undefined;



  return exports;
}));


