// package: services
// file: wakadoo.proto

/* tslint:disable */
/* eslint-disable */

import * as grpc from "grpc";
import * as wakadoo_pb from "./wakadoo_pb";
import * as proto_auth_pb from "./proto/auth_pb";
import * as proto_file_pb from "./proto/file_pb";
import * as proto_session_pb from "./proto/session_pb";

interface IFileService extends grpc.ServiceDefinition<grpc.UntypedServiceImplementation> {
    save: IFileService_ISave;
    get: IFileService_IGet;
}

interface IFileService_ISave extends grpc.MethodDefinition<proto_file_pb.FileSaveRequest, proto_file_pb.FileSaveResponse> {
    path: "/services.File/Save";
    requestStream: false;
    responseStream: false;
    requestSerialize: grpc.serialize<proto_file_pb.FileSaveRequest>;
    requestDeserialize: grpc.deserialize<proto_file_pb.FileSaveRequest>;
    responseSerialize: grpc.serialize<proto_file_pb.FileSaveResponse>;
    responseDeserialize: grpc.deserialize<proto_file_pb.FileSaveResponse>;
}
interface IFileService_IGet extends grpc.MethodDefinition<proto_file_pb.FileGetRequest, proto_file_pb.FileGetResponse> {
    path: "/services.File/Get";
    requestStream: false;
    responseStream: false;
    requestSerialize: grpc.serialize<proto_file_pb.FileGetRequest>;
    requestDeserialize: grpc.deserialize<proto_file_pb.FileGetRequest>;
    responseSerialize: grpc.serialize<proto_file_pb.FileGetResponse>;
    responseDeserialize: grpc.deserialize<proto_file_pb.FileGetResponse>;
}

export const FileService: IFileService;

export interface IFileServer {
    save: grpc.handleUnaryCall<proto_file_pb.FileSaveRequest, proto_file_pb.FileSaveResponse>;
    get: grpc.handleUnaryCall<proto_file_pb.FileGetRequest, proto_file_pb.FileGetResponse>;
}

export interface IFileClient {
    save(request: proto_file_pb.FileSaveRequest, callback: (error: grpc.ServiceError | null, response: proto_file_pb.FileSaveResponse) => void): grpc.ClientUnaryCall;
    save(request: proto_file_pb.FileSaveRequest, metadata: grpc.Metadata, callback: (error: grpc.ServiceError | null, response: proto_file_pb.FileSaveResponse) => void): grpc.ClientUnaryCall;
    save(request: proto_file_pb.FileSaveRequest, metadata: grpc.Metadata, options: Partial<grpc.CallOptions>, callback: (error: grpc.ServiceError | null, response: proto_file_pb.FileSaveResponse) => void): grpc.ClientUnaryCall;
    get(request: proto_file_pb.FileGetRequest, callback: (error: grpc.ServiceError | null, response: proto_file_pb.FileGetResponse) => void): grpc.ClientUnaryCall;
    get(request: proto_file_pb.FileGetRequest, metadata: grpc.Metadata, callback: (error: grpc.ServiceError | null, response: proto_file_pb.FileGetResponse) => void): grpc.ClientUnaryCall;
    get(request: proto_file_pb.FileGetRequest, metadata: grpc.Metadata, options: Partial<grpc.CallOptions>, callback: (error: grpc.ServiceError | null, response: proto_file_pb.FileGetResponse) => void): grpc.ClientUnaryCall;
}

export class FileClient extends grpc.Client implements IFileClient {
    constructor(address: string, credentials: grpc.ChannelCredentials, options?: object);
    public save(request: proto_file_pb.FileSaveRequest, callback: (error: grpc.ServiceError | null, response: proto_file_pb.FileSaveResponse) => void): grpc.ClientUnaryCall;
    public save(request: proto_file_pb.FileSaveRequest, metadata: grpc.Metadata, callback: (error: grpc.ServiceError | null, response: proto_file_pb.FileSaveResponse) => void): grpc.ClientUnaryCall;
    public save(request: proto_file_pb.FileSaveRequest, metadata: grpc.Metadata, options: Partial<grpc.CallOptions>, callback: (error: grpc.ServiceError | null, response: proto_file_pb.FileSaveResponse) => void): grpc.ClientUnaryCall;
    public get(request: proto_file_pb.FileGetRequest, callback: (error: grpc.ServiceError | null, response: proto_file_pb.FileGetResponse) => void): grpc.ClientUnaryCall;
    public get(request: proto_file_pb.FileGetRequest, metadata: grpc.Metadata, callback: (error: grpc.ServiceError | null, response: proto_file_pb.FileGetResponse) => void): grpc.ClientUnaryCall;
    public get(request: proto_file_pb.FileGetRequest, metadata: grpc.Metadata, options: Partial<grpc.CallOptions>, callback: (error: grpc.ServiceError | null, response: proto_file_pb.FileGetResponse) => void): grpc.ClientUnaryCall;
}

interface IAuthService extends grpc.ServiceDefinition<grpc.UntypedServiceImplementation> {
    signIn: IAuthService_ISignIn;
}

interface IAuthService_ISignIn extends grpc.MethodDefinition<proto_auth_pb.SignInRequest, proto_auth_pb.SignInResponse> {
    path: "/services.Auth/SignIn";
    requestStream: false;
    responseStream: false;
    requestSerialize: grpc.serialize<proto_auth_pb.SignInRequest>;
    requestDeserialize: grpc.deserialize<proto_auth_pb.SignInRequest>;
    responseSerialize: grpc.serialize<proto_auth_pb.SignInResponse>;
    responseDeserialize: grpc.deserialize<proto_auth_pb.SignInResponse>;
}

export const AuthService: IAuthService;

export interface IAuthServer {
    signIn: grpc.handleUnaryCall<proto_auth_pb.SignInRequest, proto_auth_pb.SignInResponse>;
}

export interface IAuthClient {
    signIn(request: proto_auth_pb.SignInRequest, callback: (error: grpc.ServiceError | null, response: proto_auth_pb.SignInResponse) => void): grpc.ClientUnaryCall;
    signIn(request: proto_auth_pb.SignInRequest, metadata: grpc.Metadata, callback: (error: grpc.ServiceError | null, response: proto_auth_pb.SignInResponse) => void): grpc.ClientUnaryCall;
    signIn(request: proto_auth_pb.SignInRequest, metadata: grpc.Metadata, options: Partial<grpc.CallOptions>, callback: (error: grpc.ServiceError | null, response: proto_auth_pb.SignInResponse) => void): grpc.ClientUnaryCall;
}

export class AuthClient extends grpc.Client implements IAuthClient {
    constructor(address: string, credentials: grpc.ChannelCredentials, options?: object);
    public signIn(request: proto_auth_pb.SignInRequest, callback: (error: grpc.ServiceError | null, response: proto_auth_pb.SignInResponse) => void): grpc.ClientUnaryCall;
    public signIn(request: proto_auth_pb.SignInRequest, metadata: grpc.Metadata, callback: (error: grpc.ServiceError | null, response: proto_auth_pb.SignInResponse) => void): grpc.ClientUnaryCall;
    public signIn(request: proto_auth_pb.SignInRequest, metadata: grpc.Metadata, options: Partial<grpc.CallOptions>, callback: (error: grpc.ServiceError | null, response: proto_auth_pb.SignInResponse) => void): grpc.ClientUnaryCall;
}

interface ISessionService extends grpc.ServiceDefinition<grpc.UntypedServiceImplementation> {
    create: ISessionService_ICreate;
    getEmail: ISessionService_IGetEmail;
    updateLastActivityTime: ISessionService_IUpdateLastActivityTime;
    getActive: ISessionService_IGetActive;
}

interface ISessionService_ICreate extends grpc.MethodDefinition<proto_session_pb.SessionCreateRequest, proto_session_pb.SessionCreateResponse> {
    path: "/services.Session/Create";
    requestStream: false;
    responseStream: false;
    requestSerialize: grpc.serialize<proto_session_pb.SessionCreateRequest>;
    requestDeserialize: grpc.deserialize<proto_session_pb.SessionCreateRequest>;
    responseSerialize: grpc.serialize<proto_session_pb.SessionCreateResponse>;
    responseDeserialize: grpc.deserialize<proto_session_pb.SessionCreateResponse>;
}
interface ISessionService_IGetEmail extends grpc.MethodDefinition<proto_session_pb.SessionGetEmailRequest, proto_session_pb.SessionGetEmailResponse> {
    path: "/services.Session/GetEmail";
    requestStream: false;
    responseStream: false;
    requestSerialize: grpc.serialize<proto_session_pb.SessionGetEmailRequest>;
    requestDeserialize: grpc.deserialize<proto_session_pb.SessionGetEmailRequest>;
    responseSerialize: grpc.serialize<proto_session_pb.SessionGetEmailResponse>;
    responseDeserialize: grpc.deserialize<proto_session_pb.SessionGetEmailResponse>;
}
interface ISessionService_IUpdateLastActivityTime extends grpc.MethodDefinition<proto_session_pb.SessionUpdateActivityRequest, proto_session_pb.SessionUpdateActivityResponse> {
    path: "/services.Session/UpdateLastActivityTime";
    requestStream: false;
    responseStream: false;
    requestSerialize: grpc.serialize<proto_session_pb.SessionUpdateActivityRequest>;
    requestDeserialize: grpc.deserialize<proto_session_pb.SessionUpdateActivityRequest>;
    responseSerialize: grpc.serialize<proto_session_pb.SessionUpdateActivityResponse>;
    responseDeserialize: grpc.deserialize<proto_session_pb.SessionUpdateActivityResponse>;
}
interface ISessionService_IGetActive extends grpc.MethodDefinition<proto_session_pb.SessionGetActiveRequest, proto_session_pb.SessionGetActiveResponse> {
    path: "/services.Session/GetActive";
    requestStream: false;
    responseStream: false;
    requestSerialize: grpc.serialize<proto_session_pb.SessionGetActiveRequest>;
    requestDeserialize: grpc.deserialize<proto_session_pb.SessionGetActiveRequest>;
    responseSerialize: grpc.serialize<proto_session_pb.SessionGetActiveResponse>;
    responseDeserialize: grpc.deserialize<proto_session_pb.SessionGetActiveResponse>;
}

export const SessionService: ISessionService;

export interface ISessionServer {
    create: grpc.handleUnaryCall<proto_session_pb.SessionCreateRequest, proto_session_pb.SessionCreateResponse>;
    getEmail: grpc.handleUnaryCall<proto_session_pb.SessionGetEmailRequest, proto_session_pb.SessionGetEmailResponse>;
    updateLastActivityTime: grpc.handleUnaryCall<proto_session_pb.SessionUpdateActivityRequest, proto_session_pb.SessionUpdateActivityResponse>;
    getActive: grpc.handleUnaryCall<proto_session_pb.SessionGetActiveRequest, proto_session_pb.SessionGetActiveResponse>;
}

export interface ISessionClient {
    create(request: proto_session_pb.SessionCreateRequest, callback: (error: grpc.ServiceError | null, response: proto_session_pb.SessionCreateResponse) => void): grpc.ClientUnaryCall;
    create(request: proto_session_pb.SessionCreateRequest, metadata: grpc.Metadata, callback: (error: grpc.ServiceError | null, response: proto_session_pb.SessionCreateResponse) => void): grpc.ClientUnaryCall;
    create(request: proto_session_pb.SessionCreateRequest, metadata: grpc.Metadata, options: Partial<grpc.CallOptions>, callback: (error: grpc.ServiceError | null, response: proto_session_pb.SessionCreateResponse) => void): grpc.ClientUnaryCall;
    getEmail(request: proto_session_pb.SessionGetEmailRequest, callback: (error: grpc.ServiceError | null, response: proto_session_pb.SessionGetEmailResponse) => void): grpc.ClientUnaryCall;
    getEmail(request: proto_session_pb.SessionGetEmailRequest, metadata: grpc.Metadata, callback: (error: grpc.ServiceError | null, response: proto_session_pb.SessionGetEmailResponse) => void): grpc.ClientUnaryCall;
    getEmail(request: proto_session_pb.SessionGetEmailRequest, metadata: grpc.Metadata, options: Partial<grpc.CallOptions>, callback: (error: grpc.ServiceError | null, response: proto_session_pb.SessionGetEmailResponse) => void): grpc.ClientUnaryCall;
    updateLastActivityTime(request: proto_session_pb.SessionUpdateActivityRequest, callback: (error: grpc.ServiceError | null, response: proto_session_pb.SessionUpdateActivityResponse) => void): grpc.ClientUnaryCall;
    updateLastActivityTime(request: proto_session_pb.SessionUpdateActivityRequest, metadata: grpc.Metadata, callback: (error: grpc.ServiceError | null, response: proto_session_pb.SessionUpdateActivityResponse) => void): grpc.ClientUnaryCall;
    updateLastActivityTime(request: proto_session_pb.SessionUpdateActivityRequest, metadata: grpc.Metadata, options: Partial<grpc.CallOptions>, callback: (error: grpc.ServiceError | null, response: proto_session_pb.SessionUpdateActivityResponse) => void): grpc.ClientUnaryCall;
    getActive(request: proto_session_pb.SessionGetActiveRequest, callback: (error: grpc.ServiceError | null, response: proto_session_pb.SessionGetActiveResponse) => void): grpc.ClientUnaryCall;
    getActive(request: proto_session_pb.SessionGetActiveRequest, metadata: grpc.Metadata, callback: (error: grpc.ServiceError | null, response: proto_session_pb.SessionGetActiveResponse) => void): grpc.ClientUnaryCall;
    getActive(request: proto_session_pb.SessionGetActiveRequest, metadata: grpc.Metadata, options: Partial<grpc.CallOptions>, callback: (error: grpc.ServiceError | null, response: proto_session_pb.SessionGetActiveResponse) => void): grpc.ClientUnaryCall;
}

export class SessionClient extends grpc.Client implements ISessionClient {
    constructor(address: string, credentials: grpc.ChannelCredentials, options?: object);
    public create(request: proto_session_pb.SessionCreateRequest, callback: (error: grpc.ServiceError | null, response: proto_session_pb.SessionCreateResponse) => void): grpc.ClientUnaryCall;
    public create(request: proto_session_pb.SessionCreateRequest, metadata: grpc.Metadata, callback: (error: grpc.ServiceError | null, response: proto_session_pb.SessionCreateResponse) => void): grpc.ClientUnaryCall;
    public create(request: proto_session_pb.SessionCreateRequest, metadata: grpc.Metadata, options: Partial<grpc.CallOptions>, callback: (error: grpc.ServiceError | null, response: proto_session_pb.SessionCreateResponse) => void): grpc.ClientUnaryCall;
    public getEmail(request: proto_session_pb.SessionGetEmailRequest, callback: (error: grpc.ServiceError | null, response: proto_session_pb.SessionGetEmailResponse) => void): grpc.ClientUnaryCall;
    public getEmail(request: proto_session_pb.SessionGetEmailRequest, metadata: grpc.Metadata, callback: (error: grpc.ServiceError | null, response: proto_session_pb.SessionGetEmailResponse) => void): grpc.ClientUnaryCall;
    public getEmail(request: proto_session_pb.SessionGetEmailRequest, metadata: grpc.Metadata, options: Partial<grpc.CallOptions>, callback: (error: grpc.ServiceError | null, response: proto_session_pb.SessionGetEmailResponse) => void): grpc.ClientUnaryCall;
    public updateLastActivityTime(request: proto_session_pb.SessionUpdateActivityRequest, callback: (error: grpc.ServiceError | null, response: proto_session_pb.SessionUpdateActivityResponse) => void): grpc.ClientUnaryCall;
    public updateLastActivityTime(request: proto_session_pb.SessionUpdateActivityRequest, metadata: grpc.Metadata, callback: (error: grpc.ServiceError | null, response: proto_session_pb.SessionUpdateActivityResponse) => void): grpc.ClientUnaryCall;
    public updateLastActivityTime(request: proto_session_pb.SessionUpdateActivityRequest, metadata: grpc.Metadata, options: Partial<grpc.CallOptions>, callback: (error: grpc.ServiceError | null, response: proto_session_pb.SessionUpdateActivityResponse) => void): grpc.ClientUnaryCall;
    public getActive(request: proto_session_pb.SessionGetActiveRequest, callback: (error: grpc.ServiceError | null, response: proto_session_pb.SessionGetActiveResponse) => void): grpc.ClientUnaryCall;
    public getActive(request: proto_session_pb.SessionGetActiveRequest, metadata: grpc.Metadata, callback: (error: grpc.ServiceError | null, response: proto_session_pb.SessionGetActiveResponse) => void): grpc.ClientUnaryCall;
    public getActive(request: proto_session_pb.SessionGetActiveRequest, metadata: grpc.Metadata, options: Partial<grpc.CallOptions>, callback: (error: grpc.ServiceError | null, response: proto_session_pb.SessionGetActiveResponse) => void): grpc.ClientUnaryCall;
}
