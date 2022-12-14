// package: auth
// file: proto/auth.proto

/* tslint:disable */
/* eslint-disable */

import * as jspb from "google-protobuf";
import * as proto_account_pb from "../proto/account_pb";

export class SignInRequest extends jspb.Message { 
    getEmail(): string;
    setEmail(value: string): SignInRequest;
    getPassword(): string;
    setPassword(value: string): SignInRequest;

    serializeBinary(): Uint8Array;
    toObject(includeInstance?: boolean): SignInRequest.AsObject;
    static toObject(includeInstance: boolean, msg: SignInRequest): SignInRequest.AsObject;
    static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
    static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
    static serializeBinaryToWriter(message: SignInRequest, writer: jspb.BinaryWriter): void;
    static deserializeBinary(bytes: Uint8Array): SignInRequest;
    static deserializeBinaryFromReader(message: SignInRequest, reader: jspb.BinaryReader): SignInRequest;
}

export namespace SignInRequest {
    export type AsObject = {
        email: string,
        password: string,
    }
}

export class SignInResponse extends jspb.Message { 

    hasAccount(): boolean;
    clearAccount(): void;
    getAccount(): proto_account_pb.Account | undefined;
    setAccount(value?: proto_account_pb.Account): SignInResponse;
    getSession(): string;
    setSession(value: string): SignInResponse;

    serializeBinary(): Uint8Array;
    toObject(includeInstance?: boolean): SignInResponse.AsObject;
    static toObject(includeInstance: boolean, msg: SignInResponse): SignInResponse.AsObject;
    static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
    static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
    static serializeBinaryToWriter(message: SignInResponse, writer: jspb.BinaryWriter): void;
    static deserializeBinary(bytes: Uint8Array): SignInResponse;
    static deserializeBinaryFromReader(message: SignInResponse, reader: jspb.BinaryReader): SignInResponse;
}

export namespace SignInResponse {
    export type AsObject = {
        account?: proto_account_pb.Account.AsObject,
        session: string,
    }
}
