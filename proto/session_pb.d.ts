import * as jspb from 'google-protobuf'

import * as google_protobuf_timestamp_pb from 'google-protobuf/google/protobuf/timestamp_pb';


export class Session extends jspb.Message {
  getIp(): string;
  setIp(value: string): Session;

  getLastactivity(): google_protobuf_timestamp_pb.Timestamp | undefined;
  setLastactivity(value?: google_protobuf_timestamp_pb.Timestamp): Session;
  hasLastactivity(): boolean;
  clearLastactivity(): Session;

  getCreatedate(): google_protobuf_timestamp_pb.Timestamp | undefined;
  setCreatedate(value?: google_protobuf_timestamp_pb.Timestamp): Session;
  hasCreatedate(): boolean;
  clearCreatedate(): Session;

  getDevice(): string;
  setDevice(value: string): Session;

  getEmail(): string;
  setEmail(value: string): Session;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): Session.AsObject;
  static toObject(includeInstance: boolean, msg: Session): Session.AsObject;
  static serializeBinaryToWriter(message: Session, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): Session;
  static deserializeBinaryFromReader(message: Session, reader: jspb.BinaryReader): Session;
}

export namespace Session {
  export type AsObject = {
    ip: string,
    lastactivity?: google_protobuf_timestamp_pb.Timestamp.AsObject,
    createdate?: google_protobuf_timestamp_pb.Timestamp.AsObject,
    device: string,
    email: string,
  }
}

export class SessionCreateRequest extends jspb.Message {
  getIp(): string;
  setIp(value: string): SessionCreateRequest;

  getLastactivity(): google_protobuf_timestamp_pb.Timestamp | undefined;
  setLastactivity(value?: google_protobuf_timestamp_pb.Timestamp): SessionCreateRequest;
  hasLastactivity(): boolean;
  clearLastactivity(): SessionCreateRequest;

  getCreatedate(): google_protobuf_timestamp_pb.Timestamp | undefined;
  setCreatedate(value?: google_protobuf_timestamp_pb.Timestamp): SessionCreateRequest;
  hasCreatedate(): boolean;
  clearCreatedate(): SessionCreateRequest;

  getDevice(): string;
  setDevice(value: string): SessionCreateRequest;

  getEmail(): string;
  setEmail(value: string): SessionCreateRequest;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): SessionCreateRequest.AsObject;
  static toObject(includeInstance: boolean, msg: SessionCreateRequest): SessionCreateRequest.AsObject;
  static serializeBinaryToWriter(message: SessionCreateRequest, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): SessionCreateRequest;
  static deserializeBinaryFromReader(message: SessionCreateRequest, reader: jspb.BinaryReader): SessionCreateRequest;
}

export namespace SessionCreateRequest {
  export type AsObject = {
    ip: string,
    lastactivity?: google_protobuf_timestamp_pb.Timestamp.AsObject,
    createdate?: google_protobuf_timestamp_pb.Timestamp.AsObject,
    device: string,
    email: string,
  }
}

export class SessionCreateResponse extends jspb.Message {
  getHash(): string;
  setHash(value: string): SessionCreateResponse;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): SessionCreateResponse.AsObject;
  static toObject(includeInstance: boolean, msg: SessionCreateResponse): SessionCreateResponse.AsObject;
  static serializeBinaryToWriter(message: SessionCreateResponse, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): SessionCreateResponse;
  static deserializeBinaryFromReader(message: SessionCreateResponse, reader: jspb.BinaryReader): SessionCreateResponse;
}

export namespace SessionCreateResponse {
  export type AsObject = {
    hash: string,
  }
}

export class SessionGetEmailRequest extends jspb.Message {
  getHash(): string;
  setHash(value: string): SessionGetEmailRequest;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): SessionGetEmailRequest.AsObject;
  static toObject(includeInstance: boolean, msg: SessionGetEmailRequest): SessionGetEmailRequest.AsObject;
  static serializeBinaryToWriter(message: SessionGetEmailRequest, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): SessionGetEmailRequest;
  static deserializeBinaryFromReader(message: SessionGetEmailRequest, reader: jspb.BinaryReader): SessionGetEmailRequest;
}

export namespace SessionGetEmailRequest {
  export type AsObject = {
    hash: string,
  }
}

export class SessionGetEmailResponse extends jspb.Message {
  getEmail(): string;
  setEmail(value: string): SessionGetEmailResponse;

  getOk(): boolean;
  setOk(value: boolean): SessionGetEmailResponse;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): SessionGetEmailResponse.AsObject;
  static toObject(includeInstance: boolean, msg: SessionGetEmailResponse): SessionGetEmailResponse.AsObject;
  static serializeBinaryToWriter(message: SessionGetEmailResponse, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): SessionGetEmailResponse;
  static deserializeBinaryFromReader(message: SessionGetEmailResponse, reader: jspb.BinaryReader): SessionGetEmailResponse;
}

export namespace SessionGetEmailResponse {
  export type AsObject = {
    email: string,
    ok: boolean,
  }
}

export class SessionUpdateActivityRequest extends jspb.Message {
  getHash(): string;
  setHash(value: string): SessionUpdateActivityRequest;

  getTimestamp(): google_protobuf_timestamp_pb.Timestamp | undefined;
  setTimestamp(value?: google_protobuf_timestamp_pb.Timestamp): SessionUpdateActivityRequest;
  hasTimestamp(): boolean;
  clearTimestamp(): SessionUpdateActivityRequest;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): SessionUpdateActivityRequest.AsObject;
  static toObject(includeInstance: boolean, msg: SessionUpdateActivityRequest): SessionUpdateActivityRequest.AsObject;
  static serializeBinaryToWriter(message: SessionUpdateActivityRequest, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): SessionUpdateActivityRequest;
  static deserializeBinaryFromReader(message: SessionUpdateActivityRequest, reader: jspb.BinaryReader): SessionUpdateActivityRequest;
}

export namespace SessionUpdateActivityRequest {
  export type AsObject = {
    hash: string,
    timestamp?: google_protobuf_timestamp_pb.Timestamp.AsObject,
  }
}

export class SessionUpdateActivityResponse extends jspb.Message {
  getOk(): boolean;
  setOk(value: boolean): SessionUpdateActivityResponse;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): SessionUpdateActivityResponse.AsObject;
  static toObject(includeInstance: boolean, msg: SessionUpdateActivityResponse): SessionUpdateActivityResponse.AsObject;
  static serializeBinaryToWriter(message: SessionUpdateActivityResponse, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): SessionUpdateActivityResponse;
  static deserializeBinaryFromReader(message: SessionUpdateActivityResponse, reader: jspb.BinaryReader): SessionUpdateActivityResponse;
}

export namespace SessionUpdateActivityResponse {
  export type AsObject = {
    ok: boolean,
  }
}

export class SessionGetActiveRequest extends jspb.Message {
  getEmail(): string;
  setEmail(value: string): SessionGetActiveRequest;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): SessionGetActiveRequest.AsObject;
  static toObject(includeInstance: boolean, msg: SessionGetActiveRequest): SessionGetActiveRequest.AsObject;
  static serializeBinaryToWriter(message: SessionGetActiveRequest, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): SessionGetActiveRequest;
  static deserializeBinaryFromReader(message: SessionGetActiveRequest, reader: jspb.BinaryReader): SessionGetActiveRequest;
}

export namespace SessionGetActiveRequest {
  export type AsObject = {
    email: string,
  }
}

export class SessionGetActiveResponse extends jspb.Message {
  getSessionsList(): Array<Session>;
  setSessionsList(value: Array<Session>): SessionGetActiveResponse;
  clearSessionsList(): SessionGetActiveResponse;
  addSessions(value?: Session, index?: number): Session;

  getOk(): boolean;
  setOk(value: boolean): SessionGetActiveResponse;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): SessionGetActiveResponse.AsObject;
  static toObject(includeInstance: boolean, msg: SessionGetActiveResponse): SessionGetActiveResponse.AsObject;
  static serializeBinaryToWriter(message: SessionGetActiveResponse, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): SessionGetActiveResponse;
  static deserializeBinaryFromReader(message: SessionGetActiveResponse, reader: jspb.BinaryReader): SessionGetActiveResponse;
}

export namespace SessionGetActiveResponse {
  export type AsObject = {
    sessionsList: Array<Session.AsObject>,
    ok: boolean,
  }
}

