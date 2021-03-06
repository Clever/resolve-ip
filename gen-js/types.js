module.exports.Errors = {};

/**
 * BadRequest
 * @extends Error
 * @memberof module:resolve-ip
 * @alias module:resolve-ip.Errors.BadRequest
 * @property {string} message
 */
module.exports.Errors.BadRequest = class extends Error {
  constructor(body) {
    super(body.message);
    for (const k of Object.keys(body)) {
      this[k] = body[k];
    }
  }
};

/**
 * InternalError
 * @extends Error
 * @memberof module:resolve-ip
 * @alias module:resolve-ip.Errors.InternalError
 * @property {string} message
 */
module.exports.Errors.InternalError = class extends Error {
  constructor(body) {
    super(body.message);
    for (const k of Object.keys(body)) {
      this[k] = body[k];
    }
  }
};

/**
 * NotFound
 * @extends Error
 * @memberof module:resolve-ip
 * @alias module:resolve-ip.Errors.NotFound
 * @property {string} message
 */
module.exports.Errors.NotFound = class extends Error {
  constructor(body) {
    super(body.message);
    for (const k of Object.keys(body)) {
      this[k] = body[k];
    }
  }
};

