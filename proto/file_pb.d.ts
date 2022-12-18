import * as jspb from 'google-protobuf'



export class FileSaveRequest extends jspb.Message {
  getName(): string;
  setName(value: string): FileSaveRequest;

  getBytes(): Uint8Array | string;
  getBytes_asU8(): Uint8Array;
  getBytes_asB64(): string;
  setBytes(value: Uint8Array | string): FileSaveRequest;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): FileSaveRequest.AsObject;
  static toObject(includeInstance: boolean, msg: FileSaveRequest): FileSaveRequest.AsObject;
  static serializeBinaryToWriter(message: FileSaveRequest, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): FileSaveRequest;
  static deserializeBinaryFromReader(message: FileSaveRequest, reader: jspb.BinaryReader): FileSaveRequest;
}

export namespace FileSaveRequest {
  export type AsObject = {
    name: string,
    bytes: Uint8Array | string,
  }
}

export class FileSaveResponse extends jspb.Message {
  getBytecount(): number;
  setBytecount(value: number): FileSaveResponse;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): FileSaveResponse.AsObject;
  static toObject(includeInstance: boolean, msg: FileSaveResponse): FileSaveResponse.AsObject;
  static serializeBinaryToWriter(message: FileSaveResponse, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): FileSaveResponse;
  static deserializeBinaryFromReader(message: FileSaveResponse, reader: jspb.BinaryReader): FileSaveResponse;
}

export namespace FileSaveResponse {
  export type AsObject = {
    bytecount: number,
  }
}

export class FileGetRequest extends jspb.Message {
  getName(): string;
  setName(value: string): FileGetRequest;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): FileGetRequest.AsObject;
  static toObject(includeInstance: boolean, msg: FileGetRequest): FileGetRequest.AsObject;
  static serializeBinaryToWriter(message: FileGetRequest, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): FileGetRequest;
  static deserializeBinaryFromReader(message: FileGetRequest, reader: jspb.BinaryReader): FileGetRequest;
}

export namespace FileGetRequest {
  export type AsObject = {
    name: string,
  }
}

export class FileGetResponse extends jspb.Message {
  getBytes(): Uint8Array | string;
  getBytes_asU8(): Uint8Array;
  getBytes_asB64(): string;
  setBytes(value: Uint8Array | string): FileGetResponse;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): FileGetResponse.AsObject;
  static toObject(includeInstance: boolean, msg: FileGetResponse): FileGetResponse.AsObject;
  static serializeBinaryToWriter(message: FileGetResponse, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): FileGetResponse;
  static deserializeBinaryFromReader(message: FileGetResponse, reader: jspb.BinaryReader): FileGetResponse;
}

export namespace FileGetResponse {
  export type AsObject = {
    bytes: Uint8Array | string,
  }
}

