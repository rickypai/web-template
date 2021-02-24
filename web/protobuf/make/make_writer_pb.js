// source: protobuf/make/make_writer.proto
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

var protobuf_make_make_pb = require('../../protobuf/make/make_pb.js');
goog.object.extend(proto, protobuf_make_make_pb);
goog.exportSymbol('proto.make.CreateOneRequest', null, global);
goog.exportSymbol('proto.make.CreateOneResponse', null, global);
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
proto.make.CreateOneRequest = function(opt_data) {
  jspb.Message.initialize(this, opt_data, 0, -1, null, null);
};
goog.inherits(proto.make.CreateOneRequest, jspb.Message);
if (goog.DEBUG && !COMPILED) {
  /**
   * @public
   * @override
   */
  proto.make.CreateOneRequest.displayName = 'proto.make.CreateOneRequest';
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
proto.make.CreateOneResponse = function(opt_data) {
  jspb.Message.initialize(this, opt_data, 0, -1, null, null);
};
goog.inherits(proto.make.CreateOneResponse, jspb.Message);
if (goog.DEBUG && !COMPILED) {
  /**
   * @public
   * @override
   */
  proto.make.CreateOneResponse.displayName = 'proto.make.CreateOneResponse';
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
proto.make.CreateOneRequest.prototype.toObject = function(opt_includeInstance) {
  return proto.make.CreateOneRequest.toObject(opt_includeInstance, this);
};


/**
 * Static version of the {@see toObject} method.
 * @param {boolean|undefined} includeInstance Deprecated. Whether to include
 *     the JSPB instance for transitional soy proto support:
 *     http://goto/soy-param-migration
 * @param {!proto.make.CreateOneRequest} msg The msg instance to transform.
 * @return {!Object}
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.make.CreateOneRequest.toObject = function(includeInstance, msg) {
  var f, obj = {
    request: (f = msg.getRequest()) && protobuf_make_make_pb.Make.toObject(includeInstance, f)
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
 * @return {!proto.make.CreateOneRequest}
 */
proto.make.CreateOneRequest.deserializeBinary = function(bytes) {
  var reader = new jspb.BinaryReader(bytes);
  var msg = new proto.make.CreateOneRequest;
  return proto.make.CreateOneRequest.deserializeBinaryFromReader(msg, reader);
};


/**
 * Deserializes binary data (in protobuf wire format) from the
 * given reader into the given message object.
 * @param {!proto.make.CreateOneRequest} msg The message object to deserialize into.
 * @param {!jspb.BinaryReader} reader The BinaryReader to use.
 * @return {!proto.make.CreateOneRequest}
 */
proto.make.CreateOneRequest.deserializeBinaryFromReader = function(msg, reader) {
  while (reader.nextField()) {
    if (reader.isEndGroup()) {
      break;
    }
    var field = reader.getFieldNumber();
    switch (field) {
    case 1:
      var value = new protobuf_make_make_pb.Make;
      reader.readMessage(value,protobuf_make_make_pb.Make.deserializeBinaryFromReader);
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
proto.make.CreateOneRequest.prototype.serializeBinary = function() {
  var writer = new jspb.BinaryWriter();
  proto.make.CreateOneRequest.serializeBinaryToWriter(this, writer);
  return writer.getResultBuffer();
};


/**
 * Serializes the given message to binary data (in protobuf wire
 * format), writing to the given BinaryWriter.
 * @param {!proto.make.CreateOneRequest} message
 * @param {!jspb.BinaryWriter} writer
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.make.CreateOneRequest.serializeBinaryToWriter = function(message, writer) {
  var f = undefined;
  f = message.getRequest();
  if (f != null) {
    writer.writeMessage(
      1,
      f,
      protobuf_make_make_pb.Make.serializeBinaryToWriter
    );
  }
};


/**
 * optional Make request = 1;
 * @return {?proto.make.Make}
 */
proto.make.CreateOneRequest.prototype.getRequest = function() {
  return /** @type{?proto.make.Make} */ (
    jspb.Message.getWrapperField(this, protobuf_make_make_pb.Make, 1));
};


/**
 * @param {?proto.make.Make|undefined} value
 * @return {!proto.make.CreateOneRequest} returns this
*/
proto.make.CreateOneRequest.prototype.setRequest = function(value) {
  return jspb.Message.setWrapperField(this, 1, value);
};


/**
 * Clears the message field making it undefined.
 * @return {!proto.make.CreateOneRequest} returns this
 */
proto.make.CreateOneRequest.prototype.clearRequest = function() {
  return this.setRequest(undefined);
};


/**
 * Returns whether this field is set.
 * @return {boolean}
 */
proto.make.CreateOneRequest.prototype.hasRequest = function() {
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
proto.make.CreateOneResponse.prototype.toObject = function(opt_includeInstance) {
  return proto.make.CreateOneResponse.toObject(opt_includeInstance, this);
};


/**
 * Static version of the {@see toObject} method.
 * @param {boolean|undefined} includeInstance Deprecated. Whether to include
 *     the JSPB instance for transitional soy proto support:
 *     http://goto/soy-param-migration
 * @param {!proto.make.CreateOneResponse} msg The msg instance to transform.
 * @return {!Object}
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.make.CreateOneResponse.toObject = function(includeInstance, msg) {
  var f, obj = {
    result: (f = msg.getResult()) && protobuf_make_make_pb.Make.toObject(includeInstance, f)
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
 * @return {!proto.make.CreateOneResponse}
 */
proto.make.CreateOneResponse.deserializeBinary = function(bytes) {
  var reader = new jspb.BinaryReader(bytes);
  var msg = new proto.make.CreateOneResponse;
  return proto.make.CreateOneResponse.deserializeBinaryFromReader(msg, reader);
};


/**
 * Deserializes binary data (in protobuf wire format) from the
 * given reader into the given message object.
 * @param {!proto.make.CreateOneResponse} msg The message object to deserialize into.
 * @param {!jspb.BinaryReader} reader The BinaryReader to use.
 * @return {!proto.make.CreateOneResponse}
 */
proto.make.CreateOneResponse.deserializeBinaryFromReader = function(msg, reader) {
  while (reader.nextField()) {
    if (reader.isEndGroup()) {
      break;
    }
    var field = reader.getFieldNumber();
    switch (field) {
    case 1:
      var value = new protobuf_make_make_pb.Make;
      reader.readMessage(value,protobuf_make_make_pb.Make.deserializeBinaryFromReader);
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
proto.make.CreateOneResponse.prototype.serializeBinary = function() {
  var writer = new jspb.BinaryWriter();
  proto.make.CreateOneResponse.serializeBinaryToWriter(this, writer);
  return writer.getResultBuffer();
};


/**
 * Serializes the given message to binary data (in protobuf wire
 * format), writing to the given BinaryWriter.
 * @param {!proto.make.CreateOneResponse} message
 * @param {!jspb.BinaryWriter} writer
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.make.CreateOneResponse.serializeBinaryToWriter = function(message, writer) {
  var f = undefined;
  f = message.getResult();
  if (f != null) {
    writer.writeMessage(
      1,
      f,
      protobuf_make_make_pb.Make.serializeBinaryToWriter
    );
  }
};


/**
 * optional Make result = 1;
 * @return {?proto.make.Make}
 */
proto.make.CreateOneResponse.prototype.getResult = function() {
  return /** @type{?proto.make.Make} */ (
    jspb.Message.getWrapperField(this, protobuf_make_make_pb.Make, 1));
};


/**
 * @param {?proto.make.Make|undefined} value
 * @return {!proto.make.CreateOneResponse} returns this
*/
proto.make.CreateOneResponse.prototype.setResult = function(value) {
  return jspb.Message.setWrapperField(this, 1, value);
};


/**
 * Clears the message field making it undefined.
 * @return {!proto.make.CreateOneResponse} returns this
 */
proto.make.CreateOneResponse.prototype.clearResult = function() {
  return this.setResult(undefined);
};


/**
 * Returns whether this field is set.
 * @return {boolean}
 */
proto.make.CreateOneResponse.prototype.hasResult = function() {
  return jspb.Message.getField(this, 1) != null;
};


goog.object.extend(exports, proto.make);