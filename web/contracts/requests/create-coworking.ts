// Code generated by protoc-gen-ts_proto. DO NOT EDIT.
// versions:
//   protoc-gen-ts_proto  v1.176.0
//   protoc               v3.12.4
// source: requests/create-coworking.proto

/* eslint-disable */
import * as _m0 from "protobufjs/minimal";

export const protobufPackage = "create_workspace_request";

export interface CreateWorkspaceRequest {
  name: string;
  seats: number;
}

export interface CreateCoworkingRequest {
  name: string;
  workspaces: CreateWorkspaceRequest[];
}

function createBaseCreateWorkspaceRequest(): CreateWorkspaceRequest {
  return { name: "", seats: 0 };
}

export const CreateWorkspaceRequest = {
  encode(message: CreateWorkspaceRequest, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.name !== "") {
      writer.uint32(10).string(message.name);
    }
    if (message.seats !== 0) {
      writer.uint32(16).int32(message.seats);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): CreateWorkspaceRequest {
    const reader = input instanceof _m0.Reader ? input : _m0.Reader.create(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseCreateWorkspaceRequest();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          if (tag !== 10) {
            break;
          }

          message.name = reader.string();
          continue;
        case 2:
          if (tag !== 16) {
            break;
          }

          message.seats = reader.int32();
          continue;
      }
      if ((tag & 7) === 4 || tag === 0) {
        break;
      }
      reader.skipType(tag & 7);
    }
    return message;
  },

  fromJSON(object: any): CreateWorkspaceRequest {
    return {
      name: isSet(object.name) ? globalThis.String(object.name) : "",
      seats: isSet(object.seats) ? globalThis.Number(object.seats) : 0,
    };
  },

  toJSON(message: CreateWorkspaceRequest): unknown {
    const obj: any = {};
    if (message.name !== "") {
      obj.name = message.name;
    }
    if (message.seats !== 0) {
      obj.seats = Math.round(message.seats);
    }
    return obj;
  },

  create<I extends Exact<DeepPartial<CreateWorkspaceRequest>, I>>(base?: I): CreateWorkspaceRequest {
    return CreateWorkspaceRequest.fromPartial(base ?? ({} as any));
  },
  fromPartial<I extends Exact<DeepPartial<CreateWorkspaceRequest>, I>>(object: I): CreateWorkspaceRequest {
    const message = createBaseCreateWorkspaceRequest();
    message.name = object.name ?? "";
    message.seats = object.seats ?? 0;
    return message;
  },
};

function createBaseCreateCoworkingRequest(): CreateCoworkingRequest {
  return { name: "", workspaces: [] };
}

export const CreateCoworkingRequest = {
  encode(message: CreateCoworkingRequest, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.name !== "") {
      writer.uint32(10).string(message.name);
    }
    for (const v of message.workspaces) {
      CreateWorkspaceRequest.encode(v!, writer.uint32(18).fork()).ldelim();
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): CreateCoworkingRequest {
    const reader = input instanceof _m0.Reader ? input : _m0.Reader.create(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseCreateCoworkingRequest();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          if (tag !== 10) {
            break;
          }

          message.name = reader.string();
          continue;
        case 2:
          if (tag !== 18) {
            break;
          }

          message.workspaces.push(CreateWorkspaceRequest.decode(reader, reader.uint32()));
          continue;
      }
      if ((tag & 7) === 4 || tag === 0) {
        break;
      }
      reader.skipType(tag & 7);
    }
    return message;
  },

  fromJSON(object: any): CreateCoworkingRequest {
    return {
      name: isSet(object.name) ? globalThis.String(object.name) : "",
      workspaces: globalThis.Array.isArray(object?.workspaces)
        ? object.workspaces.map((e: any) => CreateWorkspaceRequest.fromJSON(e))
        : [],
    };
  },

  toJSON(message: CreateCoworkingRequest): unknown {
    const obj: any = {};
    if (message.name !== "") {
      obj.name = message.name;
    }
    if (message.workspaces?.length) {
      obj.workspaces = message.workspaces.map((e) => CreateWorkspaceRequest.toJSON(e));
    }
    return obj;
  },

  create<I extends Exact<DeepPartial<CreateCoworkingRequest>, I>>(base?: I): CreateCoworkingRequest {
    return CreateCoworkingRequest.fromPartial(base ?? ({} as any));
  },
  fromPartial<I extends Exact<DeepPartial<CreateCoworkingRequest>, I>>(object: I): CreateCoworkingRequest {
    const message = createBaseCreateCoworkingRequest();
    message.name = object.name ?? "";
    message.workspaces = object.workspaces?.map((e) => CreateWorkspaceRequest.fromPartial(e)) || [];
    return message;
  },
};

type Builtin = Date | Function | Uint8Array | string | number | boolean | undefined;

export type DeepPartial<T> = T extends Builtin ? T
  : T extends globalThis.Array<infer U> ? globalThis.Array<DeepPartial<U>>
  : T extends ReadonlyArray<infer U> ? ReadonlyArray<DeepPartial<U>>
  : T extends {} ? { [K in keyof T]?: DeepPartial<T[K]> }
  : Partial<T>;

type KeysOfUnion<T> = T extends T ? keyof T : never;
export type Exact<P, I extends P> = P extends Builtin ? P
  : P & { [K in keyof P]: Exact<P[K], I[K]> } & { [K in Exclude<keyof I, KeysOfUnion<P>>]: never };

function isSet(value: any): boolean {
  return value !== null && value !== undefined;
}