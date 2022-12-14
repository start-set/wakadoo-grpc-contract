// package: session
// file: proto/session.proto

/* tslint:disable */
/* eslint-disable */

import * as jspb from "google-protobuf";
import * as google_protobuf_timestamp_pb from "google-protobuf/google/protobuf/timestamp_pb";

export class Session extends jspb.Message { 
    getIp(): string;
    setIp(value: string): Session;

    hasLastactivity(): boolean;
    clearLastactivity(): void;
    getLastactivity(): google_protobuf_timestamp_pb.Timestamp | undefined;
    setLastactivity(value?: google_protobuf_timestamp_pb.Timestamp): Session;

    hasCreatedate(): boolean;
    clearCreatedate(): void;
    getCreatedate(): google_protobuf_timestamp_pb.Timestamp | undefined;
    setCreatedate(value?: google_protobuf_timestamp_pb.Timestamp): Session;
    getDevice(): string;
    setDevice(value: string): Session;
    getEmail(): string;
    setEmail(value: string): Session;

    serializeBinary(): Uint8Array;
    toObject(includeInstance?: boolean): Session.AsObject;
    static toObject(includeInstance: boolean, msg: Session): Session.AsObject;
    static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
    static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
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

    hasLastactivity(): boolean;
    clearLastactivity(): void;
    getLastactivity(): google_protobuf_timestamp_pb.Timestamp | undefined;
    setLastactivity(value?: google_protobuf_timestamp_pb.Timestamp): SessionCreateRequest;

    hasCreatedate(): boolean;
    clearCreatedate(): void;
    getCreatedate(): google_protobuf_timestamp_pb.Timestamp | undefined;
    setCreatedate(value?: google_protobuf_timestamp_pb.Timestamp): SessionCreateRequest;
    getDevice(): string;
    setDevice(value: string): SessionCreateRequest;
    getEmail(): string;
    setEmail(value: string): SessionCreateRequest;

    serializeBinary(): Uint8Array;
    toObject(includeInstance?: boolean): SessionCreateRequest.AsObject;
    static toObject(includeInstance: boolean, msg: SessionCreateRequest): SessionCreateRequest.AsObject;
    static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
    static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
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
    static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
    static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
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
    static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
    static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
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
    static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
    static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
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

    hasTimestamp(): boolean;
    clearTimestamp(): void;
    getTimestamp(): google_protobuf_timestamp_pb.Timestamp | undefined;
    setTimestamp(value?: google_protobuf_timestamp_pb.Timestamp): SessionUpdateActivityRequest;

    serializeBinary(): Uint8Array;
    toObject(includeInstance?: boolean): SessionUpdateActivityRequest.AsObject;
    static toObject(includeInstance: boolean, msg: SessionUpdateActivityRequest): SessionUpdateActivityRequest.AsObject;
    static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
    static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
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
    static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
    static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
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
    static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
    static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
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
    clearSessionsList(): void;
    getSessionsList(): Array<Session>;
    setSessionsList(value: Array<Session>): SessionGetActiveResponse;
    addSessions(value?: Session, index?: number): Session;
    getOk(): boolean;
    setOk(value: boolean): SessionGetActiveResponse;

    serializeBinary(): Uint8Array;
    toObject(includeInstance?: boolean): SessionGetActiveResponse.AsObject;
    static toObject(includeInstance: boolean, msg: SessionGetActiveResponse): SessionGetActiveResponse.AsObject;
    static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
    static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
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
