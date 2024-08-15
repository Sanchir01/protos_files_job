// Code generated by protoc-gen-ts_proto. DO NOT EDIT.
// versions:
//   protoc-gen-ts_proto  v1.181.1
//   protoc               v3.20.3
// source: auth/auth.proto

/* eslint-disable */
import { GrpcMethod, GrpcStreamMethod } from "@nestjs/microservices";
import { Observable } from "rxjs";

export const protobufPackage = "auth";

export interface RegisterRequest {
  phone: string;
  email: string;
  password: string;
}

export interface LoginRequest {
  phone: string;
  password: string;
}

export interface IsAdminRequest {
  userUuid: string;
}

export interface IsAdminResponse {
  isAdmin: boolean;
}

export interface RegisterResponse {
  userUuid: string;
  tokenTTL: string;
}

export interface LoginResponse {
  userUuid: string;
  tokenTTL: string;
}

export const AUTH_PACKAGE_NAME = "auth";

export interface AuthClient {
  register(request: RegisterRequest): Observable<RegisterResponse>;

  login(request: LoginRequest): Observable<LoginResponse>;

  isAdmin(request: IsAdminRequest): Observable<IsAdminResponse>;
}

export interface AuthController {
  register(request: RegisterRequest): Promise<RegisterResponse> | Observable<RegisterResponse> | RegisterResponse;

  login(request: LoginRequest): Promise<LoginResponse> | Observable<LoginResponse> | LoginResponse;

  isAdmin(request: IsAdminRequest): Promise<IsAdminResponse> | Observable<IsAdminResponse> | IsAdminResponse;
}

export function AuthControllerMethods() {
  return function (constructor: Function) {
    const grpcMethods: string[] = ["register", "login", "isAdmin"];
    for (const method of grpcMethods) {
      const descriptor: any = Reflect.getOwnPropertyDescriptor(constructor.prototype, method);
      GrpcMethod("Auth", method)(constructor.prototype[method], method, descriptor);
    }
    const grpcStreamMethods: string[] = [];
    for (const method of grpcStreamMethods) {
      const descriptor: any = Reflect.getOwnPropertyDescriptor(constructor.prototype, method);
      GrpcStreamMethod("Auth", method)(constructor.prototype[method], method, descriptor);
    }
  };
}

export const AUTH_SERVICE_NAME = "Auth";