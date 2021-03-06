// source: protobuf/manufacturer/manufacturer_writer.proto
/**
 * @fileoverview
 * @enhanceable
 * @suppress {messageConventions} JS Compiler reports an error if a variable or
 *     field starts with 'MSG_' and isn't a translatable message.
 * @public
 */
// GENERATED CODE -- DO NOT EDIT!
/* eslint-disable */
// @ts-nocheck

var jspb = require('google-protobuf');
var goog = jspb;
var global = Function('return this')();

var protobuf_manufacturer_manufacturer_pb = require('../../protobuf/manufacturer/manufacturer_pb.js');
goog.object.extend(proto, protobuf_manufacturer_manufacturer_pb);
goog.exportSymbol('proto.manufacturer.CreateOneRequest', null, global);
goog.exportSymbol('proto.manufacturer.CreateOneResponse', null, global);
/**
 * Generated by JsPbCodeGenerator.
 * @param {Array=} opt_data Optional initial data array, typically from a
 * server response, or constructed directly in Javascript. The array is used
 * in place and becomes part of the constructed object. It is not cloned.
 * If no data is provided, the constructed object will be empty, but still
 * valid.
 * @extends {jspb.Message}
 * @constructor
 */
proto.manufacturer.CreateOneRequest = function(opt_data) {
  jspb.Message.initialize(this, opt_data, 0, -1, null, null);
};
goog.inherits(proto.manufacturer.CreateOneRequest, jspb.Message);
if (goog.DEBUG && !COMPILED) {
  /**
   * @public
   * @override
   */
  proto.manufacturer.CreateOneRequest.displayName = 'proto.manufacturer.CreateOneRequest';
}
/**
 * Generated by JsPbCodeGenerator.
 * @param {Array=} opt_data Optional initial data array, typically from a
 * server response, or constructed directly in Javascript. The array is used
 * in place and becomes part of the constructed object. It is not cloned.
 * If no data is provided, the constructed object will be empty, but still
 * valid.
 * @extends {jspb.Message}
 * @constructor
 */
proto.manufacturer.CreateOneResponse = function(opt_data) {
  jspb.Message.initialize(this, opt_data, 0, -1, null, null);
};
goog.inherits(proto.manufacturer.CreateOneResponse, jspb.Message);
if (goog.DEBUG && !COMPILED) {
  /**
   * @public
   * @override
   */
  proto.manufacturer.CreateOneResponse.displayName = 'proto.manufacturer.CreateOneResponse';
}



if (jspb.Message.GENERATE_TO_OBJECT) {
/**
 * Creates an object representation of this proto.
 * Field names that are reserved in JavaScript and will be renamed to pb_name.
 * Optional fields that are not set will be set to undefined.
 * To access a reserved field use, foo.pb_<name>, eg, foo.pb_default.
 * For the list of reserved names please see:
 *     net/proto2/compiler/js/internal/generator.cc#kKeyword.
 * @param {boolean=} opt_includeInstance Deprecated. whether to include the
 *     JSPB instance for transitional soy proto support:
 *     http://goto/soy-param-migration
 * @return {!Object}
 */
proto.manufacturer.CreateOneRequest.prototype.toObject = function(opt_includeInstance) {
  return proto.manufacturer.CreateOneRequest.toObject(opt_includeInstance, this);
};


/**
 * Static version of the {@see toObject} method.
 * @param {boolean|undefined} includeInstance Deprecated. Whether to include
 *     the JSPB instance for transitional soy proto support:
 *     http://goto/soy-param-migration
 * @param {!proto.manufacturer.CreateOneRequest} msg The msg instance to transform.
 * @return {!Object}
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.manufacturer.CreateOneRequest.toObject = function(includeInstance, msg) {
  var f, obj = {
    request: (f = msg.getRequest()) && protobuf_manufacturer_manufacturer_pb.ManufacturerCreateRequest.toObject(includeInstance, f)
  };

  if (includeInstance) {
    obj.$jspbMessageInstance = msg;
  }
  return obj;
};
}


/**
 * Deserializes binary data (in protobuf wire format).
 * @param {jspb.ByteSource} bytes The bytes to deserialize.
 * @return {!proto.manufacturer.CreateOneRequest}
 */
proto.manufacturer.CreateOneRequest.deserializeBinary = function(bytes) {
  var reader = new jspb.BinaryReader(bytes);
  var msg = new proto.manufacturer.CreateOneRequest;
  return proto.manufacturer.CreateOneRequest.deserializeBinaryFromReader(msg, reader);
};


/**
 * Deserializes binary data (in protobuf wire format) from the
 * given reader into the given message object.
 * @param {!proto.manufacturer.CreateOneRequest} msg The message object to deserialize into.
 * @param {!jspb.BinaryReader} reader The BinaryReader to use.
 * @return {!proto.manufacturer.CreateOneRequest}
 */
proto.manufacturer.CreateOneRequest.deserializeBinaryFromReader = function(msg, reader) {
  while (reader.nextField()) {
    if (reader.isEndGroup()) {
      break;
    }
    var field = reader.getFieldNumber();
    switch (field) {
    case 1:
      var value = new protobuf_manufacturer_manufacturer_pb.ManufacturerCreateRequest;
      reader.readMessage(value,protobuf_manufacturer_manufacturer_pb.ManufacturerCreateRequest.deserializeBinaryFromReader);
      msg.setRequest(value);
      break;
    default:
      reader.skipField();
      break;
    }
  }
  return msg;
};


/**
 * Serializes the message to binary data (in protobuf wire format).
 * @return {!Uint8Array}
 */
proto.manufacturer.CreateOneRequest.prototype.serializeBinary = function() {
  var writer = new jspb.BinaryWriter();
  proto.manufacturer.CreateOneRequest.serializeBinaryToWriter(this, writer);
  return writer.getResultBuffer();
};


/**
 * Serializes the given message to binary data (in protobuf wire
 * format), writing to the given BinaryWriter.
 * @param {!proto.manufacturer.CreateOneRequest} message
 * @param {!jspb.BinaryWriter} writer
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.manufacturer.CreateOneRequest.serializeBinaryToWriter = function(message, writer) {
  var f = undefined;
  f = message.getRequest();
  if (f != null) {
    writer.writeMessage(
      1,
      f,
      protobuf_manufacturer_manufacturer_pb.ManufacturerCreateRequest.serializeBinaryToWriter
    );
  }
};


/**
 * optional ManufacturerCreateRequest request = 1;
 * @return {?proto.manufacturer.ManufacturerCreateRequest}
 */
proto.manufacturer.CreateOneRequest.prototype.getRequest = function() {
  return /** @type{?proto.manufacturer.ManufacturerCreateRequest} */ (
    jspb.Message.getWrapperField(this, protobuf_manufacturer_manufacturer_pb.ManufacturerCreateRequest, 1));
};


/**
 * @param {?proto.manufacturer.ManufacturerCreateRequest|undefined} value
 * @return {!proto.manufacturer.CreateOneRequest} returns this
*/
proto.manufacturer.CreateOneRequest.prototype.setRequest = function(value) {
  return jspb.Message.setWrapperField(this, 1, value);
};


/**
 * Clears the message field making it undefined.
 * @return {!proto.manufacturer.CreateOneRequest} returns this
 */
proto.manufacturer.CreateOneRequest.prototype.clearRequest = function() {
  return this.setRequest(undefined);
};


/**
 * Returns whether this field is set.
 * @return {boolean}
 */
proto.manufacturer.CreateOneRequest.prototype.hasRequest = function() {
  return jspb.Message.getField(this, 1) != null;
};





if (jspb.Message.GENERATE_TO_OBJECT) {
/**
 * Creates an object representation of this proto.
 * Field names that are reserved in JavaScript and will be renamed to pb_name.
 * Optional fields that are not set will be set to undefined.
 * To access a reserved field use, foo.pb_<name>, eg, foo.pb_default.
 * For the list of reserved names please see:
 *     net/proto2/compiler/js/internal/generator.cc#kKeyword.
 * @param {boolean=} opt_includeInstance Deprecated. whether to include the
 *     JSPB instance for transitional soy proto support:
 *     http://goto/soy-param-migration
 * @return {!Object}
 */
proto.manufacturer.CreateOneResponse.prototype.toObject = function(opt_includeInstance) {
  return proto.manufacturer.CreateOneResponse.toObject(opt_includeInstance, this);
};


/**
 * Static version of the {@see toObject} method.
 * @param {boolean|undefined} includeInstance Deprecated. Whether to include
 *     the JSPB instance for transitional soy proto support:
 *     http://goto/soy-param-migration
 * @param {!proto.manufacturer.CreateOneResponse} msg The msg instance to transform.
 * @return {!Object}
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.manufacturer.CreateOneResponse.toObject = function(includeInstance, msg) {
  var f, obj = {
    result: (f = msg.getResult()) && protobuf_manufacturer_manufacturer_pb.Manufacturer.toObject(includeInstance, f)
  };

  if (includeInstance) {
    obj.$jspbMessageInstance = msg;
  }
  return obj;
};
}


/**
 * Deserializes binary data (in protobuf wire format).
 * @param {jspb.ByteSource} bytes The bytes to deserialize.
 * @return {!proto.manufacturer.CreateOneResponse}
 */
proto.manufacturer.CreateOneResponse.deserializeBinary = function(bytes) {
  var reader = new jspb.BinaryReader(bytes);
  var msg = new proto.manufacturer.CreateOneResponse;
  return proto.manufacturer.CreateOneResponse.deserializeBinaryFromReader(msg, reader);
};


/**
 * Deserializes binary data (in protobuf wire format) from the
 * given reader into the given message object.
 * @param {!proto.manufacturer.CreateOneResponse} msg The message object to deserialize into.
 * @param {!jspb.BinaryReader} reader The BinaryReader to use.
 * @return {!proto.manufacturer.CreateOneResponse}
 */
proto.manufacturer.CreateOneResponse.deserializeBinaryFromReader = function(msg, reader) {
  while (reader.nextField()) {
    if (reader.isEndGroup()) {
      break;
    }
    var field = reader.getFieldNumber();
    switch (field) {
    case 1:
      var value = new protobuf_manufacturer_manufacturer_pb.Manufacturer;
      reader.readMessage(value,protobuf_manufacturer_manufacturer_pb.Manufacturer.deserializeBinaryFromReader);
      msg.setResult(value);
      break;
    default:
      reader.skipField();
      break;
    }
  }
  return msg;
};


/**
 * Serializes the message to binary data (in protobuf wire format).
 * @return {!Uint8Array}
 */
proto.manufacturer.CreateOneResponse.prototype.serializeBinary = function() {
  var writer = new jspb.BinaryWriter();
  proto.manufacturer.CreateOneResponse.serializeBinaryToWriter(this, writer);
  return writer.getResultBuffer();
};


/**
 * Serializes the given message to binary data (in protobuf wire
 * format), writing to the given BinaryWriter.
 * @param {!proto.manufacturer.CreateOneResponse} message
 * @param {!jspb.BinaryWriter} writer
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.manufacturer.CreateOneResponse.serializeBinaryToWriter = function(message, writer) {
  var f = undefined;
  f = message.getResult();
  if (f != null) {
    writer.writeMessage(
      1,
      f,
      protobuf_manufacturer_manufacturer_pb.Manufacturer.serializeBinaryToWriter
    );
  }
};


/**
 * optional Manufacturer result = 1;
 * @return {?proto.manufacturer.Manufacturer}
 */
proto.manufacturer.CreateOneResponse.prototype.getResult = function() {
  return /** @type{?proto.manufacturer.Manufacturer} */ (
    jspb.Message.getWrapperField(this, protobuf_manufacturer_manufacturer_pb.Manufacturer, 1));
};


/**
 * @param {?proto.manufacturer.Manufacturer|undefined} value
 * @return {!proto.manufacturer.CreateOneResponse} returns this
*/
proto.manufacturer.CreateOneResponse.prototype.setResult = function(value) {
  return jspb.Message.setWrapperField(this, 1, value);
};


/**
 * Clears the message field making it undefined.
 * @return {!proto.manufacturer.CreateOneResponse} returns this
 */
proto.manufacturer.CreateOneResponse.prototype.clearResult = function() {
  return this.setResult(undefined);
};


/**
 * Returns whether this field is set.
 * @return {boolean}
 */
proto.manufacturer.CreateOneResponse.prototype.hasResult = function() {
  return jspb.Message.getField(this, 1) != null;
};


goog.object.extend(exports, proto.manufacturer);
