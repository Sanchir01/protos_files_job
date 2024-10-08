// Code generated by protoc-gen-ts_proto. DO NOT EDIT.
// versions:
//   protoc-gen-ts_proto  v1.181.1
//   protoc               v3.20.3
// source: size/size.proto

/* eslint-disable */
import { GrpcMethod, GrpcStreamMethod } from "@nestjs/microservices";
import { Observable } from "rxjs";
import { Empty } from "../google/protobuf/empty";
import { Timestamp } from "../google/protobuf/timestamp";

export const protobufPackage = "size";

export interface Size {
  id: string;
  name: string;
  version: number;
  createdAt: Timestamp | undefined;
  updatedAt: Timestamp | undefined;
  slug: string;
}

export interface GetAllCategoryResponse {
  category: Size[];
}

export const SIZE_PACKAGE_NAME = "size";

export interface SizesClient {
  allSizes(request: Empty): Observable<GetAllCategoryResponse>;
}

export interface SizesController {
  allSizes(
    request: Empty,
  ): Promise<GetAllCategoryResponse> | Observable<GetAllCategoryResponse> | GetAllCategoryResponse;
}

export function SizesControllerMethods() {
  return function (constructor: Function) {
    const grpcMethods: string[] = ["allSizes"];
    for (const method of grpcMethods) {
      const descriptor: any = Reflect.getOwnPropertyDescriptor(constructor.prototype, method);
      GrpcMethod("Sizes", method)(constructor.prototype[method], method, descriptor);
    }
    const grpcStreamMethods: string[] = [];
    for (const method of grpcStreamMethods) {
      const descriptor: any = Reflect.getOwnPropertyDescriptor(constructor.prototype, method);
      GrpcStreamMethod("Sizes", method)(constructor.prototype[method], method, descriptor);
    }
  };
}

export const SIZES_SERVICE_NAME = "Sizes";
