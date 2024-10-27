// Code generated by protoc-gen-ts_proto. DO NOT EDIT.
// versions:
//   protoc-gen-ts_proto  v1.176.0
//   protoc               v3.12.4
// source: dtos/workspace.proto

/* eslint-disable */
import * as _m0 from "protobufjs/minimal";

export const protobufPackage = "dtos.workspace";

export interface WorkspaceDTO {
  UUID: string;
  name: string;
  seats: number;
}

function createBaseWorkspaceDTO(): WorkspaceDTO {
  return { UUID: "", name: "", seats: 0 };
}

export const WorkspaceDTO = {
  encode(message: WorkspaceDTO, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.UUID !== "") {
      writer.uint32(10).string(message.UUID);
    }
    if (message.name !== "") {
      writer.uint32(18).string(message.name);
    }
    if (message.seats !== 0) {
      writer.uint32(24).int32(message.seats);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): WorkspaceDTO {
    const reader = input instanceof _m0.Reader ? input : _m0.Reader.create(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseWorkspaceDTO();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          if (tag !== 10) {
            break;
          }

          message.UUID = reader.string();
          continue;
        case 2:
          if (tag !== 18) {
            break;
          }

          message.name = reader.string();
          continue;
        case 3:
          if (tag !== 24) {
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

  fromJSON(object: any): WorkspaceDTO {
    return {
      UUID: isSet(object.UUID) ? globalThis.String(object.UUID) : "",
      name: isSet(object.name) ? globalThis.String(object.name) : "",
      seats: isSet(object.seats) ? globalThis.Number(object.seats) : 0,
    };
  },

  toJSON(message: WorkspaceDTO): unknown {
    const obj: any = {};
    if (message.UUID !== "") {
      obj.UUID = message.UUID;
    }
    if (message.name !== "") {
      obj.name = message.name;
    }
    if (message.seats !== 0) {
      obj.seats = Math.round(message.seats);
    }
    return obj;
  },

  create<I extends Exact<DeepPartial<WorkspaceDTO>, I>>(base?: I): WorkspaceDTO {
    return WorkspaceDTO.fromPartial(base ?? ({} as any));
  },
  fromPartial<I extends Exact<DeepPartial<WorkspaceDTO>, I>>(object: I): WorkspaceDTO {
    const message = createBaseWorkspaceDTO();
    message.UUID = object.UUID ?? "";
    message.name = object.name ?? "";
    message.seats = object.seats ?? 0;
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
