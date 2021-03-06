import * as jspb from 'google-protobuf'

import * as google_protobuf_timestamp_pb from 'google-protobuf/google/protobuf/timestamp_pb';
import * as protobuf_manufacturer_manufacturer_pb from '../../protobuf/manufacturer/manufacturer_pb';
import * as protobuf_operating_system_operating_system_pb from '../../protobuf/operating_system/operating_system_pb';


export class Phone extends jspb.Message {
  getId(): number;
  setId(value: number): Phone;

  getName(): string;
  setName(value: string): Phone;

  getManufacturer(): protobuf_manufacturer_manufacturer_pb.Manufacturer | undefined;
  setManufacturer(value?: protobuf_manufacturer_manufacturer_pb.Manufacturer): Phone;
  hasManufacturer(): boolean;
  clearManufacturer(): Phone;

  getOperatingSystem(): protobuf_operating_system_operating_system_pb.OperatingSystem | undefined;
  setOperatingSystem(value?: protobuf_operating_system_operating_system_pb.OperatingSystem): Phone;
  hasOperatingSystem(): boolean;
  clearOperatingSystem(): Phone;

  getCreatedAt(): google_protobuf_timestamp_pb.Timestamp | undefined;
  setCreatedAt(value?: google_protobuf_timestamp_pb.Timestamp): Phone;
  hasCreatedAt(): boolean;
  clearCreatedAt(): Phone;

  getModifiedAt(): google_protobuf_timestamp_pb.Timestamp | undefined;
  setModifiedAt(value?: google_protobuf_timestamp_pb.Timestamp): Phone;
  hasModifiedAt(): boolean;
  clearModifiedAt(): Phone;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): Phone.AsObject;
  static toObject(includeInstance: boolean, msg: Phone): Phone.AsObject;
  static serializeBinaryToWriter(message: Phone, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): Phone;
  static deserializeBinaryFromReader(message: Phone, reader: jspb.BinaryReader): Phone;
}

export namespace Phone {
  export type AsObject = {
    id: number,
    name: string,
    manufacturer?: protobuf_manufacturer_manufacturer_pb.Manufacturer.AsObject,
    operatingSystem?: protobuf_operating_system_operating_system_pb.OperatingSystem.AsObject,
    createdAt?: google_protobuf_timestamp_pb.Timestamp.AsObject,
    modifiedAt?: google_protobuf_timestamp_pb.Timestamp.AsObject,
  }
}

export class PhoneRecord extends jspb.Message {
  getId(): number;
  setId(value: number): PhoneRecord;

  getName(): string;
  setName(value: string): PhoneRecord;

  getManufacturerId(): number;
  setManufacturerId(value: number): PhoneRecord;

  getOperatingSystemId(): number;
  setOperatingSystemId(value: number): PhoneRecord;

  getCreatedAt(): google_protobuf_timestamp_pb.Timestamp | undefined;
  setCreatedAt(value?: google_protobuf_timestamp_pb.Timestamp): PhoneRecord;
  hasCreatedAt(): boolean;
  clearCreatedAt(): PhoneRecord;

  getModifiedAt(): google_protobuf_timestamp_pb.Timestamp | undefined;
  setModifiedAt(value?: google_protobuf_timestamp_pb.Timestamp): PhoneRecord;
  hasModifiedAt(): boolean;
  clearModifiedAt(): PhoneRecord;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): PhoneRecord.AsObject;
  static toObject(includeInstance: boolean, msg: PhoneRecord): PhoneRecord.AsObject;
  static serializeBinaryToWriter(message: PhoneRecord, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): PhoneRecord;
  static deserializeBinaryFromReader(message: PhoneRecord, reader: jspb.BinaryReader): PhoneRecord;
}

export namespace PhoneRecord {
  export type AsObject = {
    id: number,
    name: string,
    manufacturerId: number,
    operatingSystemId: number,
    createdAt?: google_protobuf_timestamp_pb.Timestamp.AsObject,
    modifiedAt?: google_protobuf_timestamp_pb.Timestamp.AsObject,
  }
}

export class PhoneCreateRequest extends jspb.Message {
  getName(): string;
  setName(value: string): PhoneCreateRequest;

  getManufacturerId(): number;
  setManufacturerId(value: number): PhoneCreateRequest;

  getOperatingSystemId(): number;
  setOperatingSystemId(value: number): PhoneCreateRequest;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): PhoneCreateRequest.AsObject;
  static toObject(includeInstance: boolean, msg: PhoneCreateRequest): PhoneCreateRequest.AsObject;
  static serializeBinaryToWriter(message: PhoneCreateRequest, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): PhoneCreateRequest;
  static deserializeBinaryFromReader(message: PhoneCreateRequest, reader: jspb.BinaryReader): PhoneCreateRequest;
}

export namespace PhoneCreateRequest {
  export type AsObject = {
    name: string,
    manufacturerId: number,
    operatingSystemId: number,
  }
}

