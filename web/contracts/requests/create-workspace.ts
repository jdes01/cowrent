// Code generated by protoc-gen-ts_proto. DO NOT EDIT.
// versions:
//   protoc-gen-ts_proto  v1.176.0
//   protoc               v3.12.4
// source: requests/create-workspace.proto

/* eslint-disable */
import * as _m0 from "protobufjs/minimal";

export const protobufPackage = "create_coworking_request";

export interface CreateWorkspaceRequest {
  name: string;
  seats: number;
  coworkingUuid: string;
}

function createBaseCreateWorkspaceRequest(): CreateWorkspaceRequest {
  return { name: "", seats: 0, coworkingUuid: "" };
}

export const CreateWorkspaceRequest = {
  encode(message: CreateWorkspaceRequest, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.name !== "") {
      writer.uint32(10).string(message.name);
    }
    if (message.seats !== 0) {
      writer.uint32(16).int32(message.seats);
    }
    if (message.coworkingUuid !== "") {
      writer.uint32(26).string(message.coworkingUuid);
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
        case 3:
          if (tag !== 26) {
            break;
          }

          message.coworkingUuid = reader.string();
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
      coworkingUuid: isSet(object.coworkingUuid) ? globalThis.String(object.coworkingUuid) : "",
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
    if (message.coworkingUuid !== "") {
      obj.coworkingUuid = message.coworkingUuid;
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
    message.coworkingUuid = object.coworkingUuid ?? "";
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
