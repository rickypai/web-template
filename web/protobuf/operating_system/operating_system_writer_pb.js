// source: protobuf/operating_system/operating_system_writer.proto
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

var protobuf_operating_system_operating_system_pb = require('../../protobuf/operating_system/operating_system_pb.js');
goog.object.extend(proto, protobuf_operating_system_operating_system_pb);
goog.exportSymbol('proto.operating_system.CreateOneRequest', null, global);
goog.exportSymbol('proto.operating_system.CreateOneResponse', null, global);
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
proto.operating_system.CreateOneRequest = function(opt_data) {
  jspb.Message.initialize(this, opt_data, 0, -1, null, null);
};
goog.inherits(proto.operating_system.CreateOneRequest, jspb.Message);
if (goog.DEBUG && !COMPILED) {
  /**
   * @public
   * @override
   */
  proto.operating_system.CreateOneRequest.displayName = 'proto.operating_system.CreateOneRequest';
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
proto.operating_system.CreateOneResponse = function(opt_data) {
  jspb.Message.initialize(this, opt_data, 0, -1, null, null);
};
goog.inherits(proto.operating_system.CreateOneResponse, jspb.Message);
if (goog.DEBUG && !COMPILED) {
  /**
   * @public
   * @override
   */
  proto.operating_system.CreateOneResponse.displayName = 'proto.operating_system.CreateOneResponse';
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
proto.operating_system.CreateOneRequest.prototype.toObject = function(opt_includeInstance) {
  return proto.operating_system.CreateOneRequest.toObject(opt_includeInstance, this);
};


/**
 * Static version of the {@see toObject} method.
 * @param {boolean|undefined} includeInstance Deprecated. Whether to include
 *     the JSPB instance for transitional soy proto support:
 *     http://goto/soy-param-migration
 * @param {!proto.operating_system.CreateOneRequest} msg The msg instance to transform.
 * @return {!Object}
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.operating_system.CreateOneRequest.toObject = function(includeInstance, msg) {
  var f, obj = {
    request: (f = msg.getRequest()) && protobuf_operating_system_operating_system_pb.OperatingSystemCreateRequest.toObject(includeInstance, f)
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
 * @return {!proto.operating_system.CreateOneRequest}
 */
proto.operating_system.CreateOneRequest.deserializeBinary = function(bytes) {
  var reader = new jspb.BinaryReader(bytes);
  var msg = new proto.operating_system.CreateOneRequest;
  return proto.operating_system.CreateOneRequest.deserializeBinaryFromReader(msg, reader);
};


/**
 * Deserializes binary data (in protobuf wire format) from the
 * given reader into the given message object.
 * @param {!proto.operating_system.CreateOneRequest} msg The message object to deserialize into.
 * @param {!jspb.BinaryReader} reader The BinaryReader to use.
 * @return {!proto.operating_system.CreateOneRequest}
 */
proto.operating_system.CreateOneRequest.deserializeBinaryFromReader = function(msg, reader) {
  while (reader.nextField()) {
    if (reader.isEndGroup()) {
      break;
    }
    var field = reader.getFieldNumber();
    switch (field) {
    case 1:
      var value = new protobuf_operating_system_operating_system_pb.OperatingSystemCreateRequest;
      reader.readMessage(value,protobuf_operating_system_operating_system_pb.OperatingSystemCreateRequest.deserializeBinaryFromReader);
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
proto.operating_system.CreateOneRequest.prototype.serializeBinary = function() {
  var writer = new jspb.BinaryWriter();
  proto.operating_system.CreateOneRequest.serializeBinaryToWriter(this, writer);
  return writer.getResultBuffer();
};


/**
 * Serializes the given message to binary data (in protobuf wire
 * format), writing to the given BinaryWriter.
 * @param {!proto.operating_system.CreateOneRequest} message
 * @param {!jspb.BinaryWriter} writer
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.operating_system.CreateOneRequest.serializeBinaryToWriter = function(message, writer) {
  var f = undefined;
  f = message.getRequest();
  if (f != null) {
    writer.writeMessage(
      1,
      f,
      protobuf_operating_system_operating_system_pb.OperatingSystemCreateRequest.serializeBinaryToWriter
    );
  }
};


/**
 * optional OperatingSystemCreateRequest request = 1;
 * @return {?proto.operating_system.OperatingSystemCreateRequest}
 */
proto.operating_system.CreateOneRequest.prototype.getRequest = function() {
  return /** @type{?proto.operating_system.OperatingSystemCreateRequest} */ (
    jspb.Message.getWrapperField(this, protobuf_operating_system_operating_system_pb.OperatingSystemCreateRequest, 1));
};


/**
 * @param {?proto.operating_system.OperatingSystemCreateRequest|undefined} value
 * @return {!proto.operating_system.CreateOneRequest} returns this
*/
proto.operating_system.CreateOneRequest.prototype.setRequest = function(value) {
  return jspb.Message.setWrapperField(this, 1, value);
};


/**
 * Clears the message field making it undefined.
 * @return {!proto.operating_system.CreateOneRequest} returns this
 */
proto.operating_system.CreateOneRequest.prototype.clearRequest = function() {
  return this.setRequest(undefined);
};


/**
 * Returns whether this field is set.
 * @return {boolean}
 */
proto.operating_system.CreateOneRequest.prototype.hasRequest = function() {
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
proto.operating_system.CreateOneResponse.prototype.toObject = function(opt_includeInstance) {
  return proto.operating_system.CreateOneResponse.toObject(opt_includeInstance, this);
};


/**
 * Static version of the {@see toObject} method.
 * @param {boolean|undefined} includeInstance Deprecated. Whether to include
 *     the JSPB instance for transitional soy proto support:
 *     http://goto/soy-param-migration
 * @param {!proto.operating_system.CreateOneResponse} msg The msg instance to transform.
 * @return {!Object}
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.operating_system.CreateOneResponse.toObject = function(includeInstance, msg) {
  var f, obj = {
    result: (f = msg.getResult()) && protobuf_operating_system_operating_system_pb.OperatingSystem.toObject(includeInstance, f)
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
 * @return {!proto.operating_system.CreateOneResponse}
 */
proto.operating_system.CreateOneResponse.deserializeBinary = function(bytes) {
  var reader = new jspb.BinaryReader(bytes);
  var msg = new proto.operating_system.CreateOneResponse;
  return proto.operating_system.CreateOneResponse.deserializeBinaryFromReader(msg, reader);
};


/**
 * Deserializes binary data (in protobuf wire format) from the
 * given reader into the given message object.
 * @param {!proto.operating_system.CreateOneResponse} msg The message object to deserialize into.
 * @param {!jspb.BinaryReader} reader The BinaryReader to use.
 * @return {!proto.operating_system.CreateOneResponse}
 */
proto.operating_system.CreateOneResponse.deserializeBinaryFromReader = function(msg, reader) {
  while (reader.nextField()) {
    if (reader.isEndGroup()) {
      break;
    }
    var field = reader.getFieldNumber();
    switch (field) {
    case 1:
      var value = new protobuf_operating_system_operating_system_pb.OperatingSystem;
      reader.readMessage(value,protobuf_operating_system_operating_system_pb.OperatingSystem.deserializeBinaryFromReader);
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
proto.operating_system.CreateOneResponse.prototype.serializeBinary = function() {
  var writer = new jspb.BinaryWriter();
  proto.operating_system.CreateOneResponse.serializeBinaryToWriter(this, writer);
  return writer.getResultBuffer();
};


/**
 * Serializes the given message to binary data (in protobuf wire
 * format), writing to the given BinaryWriter.
 * @param {!proto.operating_system.CreateOneResponse} message
 * @param {!jspb.BinaryWriter} writer
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.operating_system.CreateOneResponse.serializeBinaryToWriter = function(message, writer) {
  var f = undefined;
  f = message.getResult();
  if (f != null) {
    writer.writeMessage(
      1,
      f,
      protobuf_operating_system_operating_system_pb.OperatingSystem.serializeBinaryToWriter
    );
  }
};


/**
 * optional OperatingSystem result = 1;
 * @return {?proto.operating_system.OperatingSystem}
 */
proto.operating_system.CreateOneResponse.prototype.getResult = function() {
  return /** @type{?proto.operating_system.OperatingSystem} */ (
    jspb.Message.getWrapperField(this, protobuf_operating_system_operating_system_pb.OperatingSystem, 1));
};


/**
 * @param {?proto.operating_system.OperatingSystem|undefined} value
 * @return {!proto.operating_system.CreateOneResponse} returns this
*/
proto.operating_system.CreateOneResponse.prototype.setResult = function(value) {
  return jspb.Message.setWrapperField(this, 1, value);
};


/**
 * Clears the message field making it undefined.
 * @return {!proto.operating_system.CreateOneResponse} returns this
 */
proto.operating_system.CreateOneResponse.prototype.clearResult = function() {
  return this.setResult(undefined);
};


/**
 * Returns whether this field is set.
 * @return {boolean}
 */
proto.operating_system.CreateOneResponse.prototype.hasResult = function() {
  return jspb.Message.getField(this, 1) != null;
};


goog.object.extend(exports, proto.operating_system);