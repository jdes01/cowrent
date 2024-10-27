// Code generated by protoc-gen-ts_proto. DO NOT EDIT.
// versions:
//   protoc-gen-ts_proto  v1.176.0
//   protoc               v3.12.4
// source: requests/coworkings.proto

/* eslint-disable */
import * as _m0 from "protobufjs/minimal";

export const protobufPackage = "requests.coworkings";

export interface AddImageToCoworkingRequest {
  coworkingUuid: string;
  imageData: Uint8Array;
  imageFilename: string;
}

export interface CreateCoworkingRequestWorkspace {
  name: string;
  seats: number;
}

export interface CreateCoworkingRequest {
  name: string;
  workspaces: CreateCoworkingRequestWorkspace[];
}

export interface AddWorkspaceToCoworkingRequest {
  name: string;
  seats: number;
  coworkingUuid: string;
}

function createBaseAddImageToCoworkingRequest(): AddImageToCoworkingRequest {
  return { coworkingUuid: "", imageData: new Uint8Array(0), imageFilename: "" };
}

export const AddImageToCoworkingRequest = {
  encode(message: AddImageToCoworkingRequest, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.coworkingUuid !== "") {
      writer.uint32(10).string(message.coworkingUuid);
    }
    if (message.imageData.length !== 0) {
      writer.uint32(18).bytes(message.imageData);
    }
    if (message.imageFilename !== "") {
      writer.uint32(26).string(message.imageFilename);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): AddImageToCoworkingRequest {
    const reader = input instanceof _m0.Reader ? input : _m0.Reader.create(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseAddImageToCoworkingRequest();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          if (tag !== 10) {
            break;
          }

          message.coworkingUuid = reader.string();
          continue;
        case 2:
          if (tag !== 18) {
            break;
          }

          message.imageData = reader.bytes();
          continue;
        case 3:
          if (tag !== 26) {
            break;
          }

          message.imageFilename = reader.string();
          continue;
      }
      if ((tag & 7) === 4 || tag === 0) {
        break;
      }
      reader.skipType(tag & 7);
    }
    return message;
  },

  fromJSON(object: any): AddImageToCoworkingRequest {
    return {
      coworkingUuid: isSet(object.coworkingUuid) ? globalThis.String(object.coworkingUuid) : "",
      imageData: isSet(object.imageData) ? bytesFromBase64(object.imageData) : new Uint8Array(0),
      imageFilename: isSet(object.imageFilename) ? globalThis.String(object.imageFilename) : "",
    };
  },

  toJSON(message: AddImageToCoworkingRequest): unknown {
    const obj: any = {};
    if (message.coworkingUuid !== "") {
      obj.coworkingUuid = message.coworkingUuid;
    }
    if (message.imageData.length !== 0) {
      obj.imageData = base64FromBytes(message.imageData);
    }
    if (message.imageFilename !== "") {
      obj.imageFilename = message.imageFilename;
    }
    return obj;
  },

  create<I extends Exact<DeepPartial<AddImageToCoworkingRequest>, I>>(base?: I): AddImageToCoworkingRequest {
    return AddImageToCoworkingRequest.fromPartial(base ?? ({} as any));
  },
  fromPartial<I extends Exact<DeepPartial<AddImageToCoworkingRequest>, I>>(object: I): AddImageToCoworkingRequest {
    const message = createBaseAddImageToCoworkingRequest();
    message.coworkingUuid = object.coworkingUuid ?? "";
    message.imageData = object.imageData ?? new Uint8Array(0);
    message.imageFilename = object.imageFilename ?? "";
    return message;
  },
};

function createBaseCreateCoworkingRequestWorkspace(): CreateCoworkingRequestWorkspace {
  return { name: "", seats: 0 };
}

export const CreateCoworkingRequestWorkspace = {
  encode(message: CreateCoworkingRequestWorkspace, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.name !== "") {
      writer.uint32(10).string(message.name);
    }
    if (message.seats !== 0) {
      writer.uint32(16).int32(message.seats);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): CreateCoworkingRequestWorkspace {
    const reader = input instanceof _m0.Reader ? input : _m0.Reader.create(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseCreateCoworkingRequestWorkspace();
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

  fromJSON(object: any): CreateCoworkingRequestWorkspace {
    return {
      name: isSet(object.name) ? globalThis.String(object.name) : "",
      seats: isSet(object.seats) ? globalThis.Number(object.seats) : 0,
    };
  },

  toJSON(message: CreateCoworkingRequestWorkspace): unknown {
    const obj: any = {};
    if (message.name !== "") {
      obj.name = message.name;
    }
    if (message.seats !== 0) {
      obj.seats = Math.round(message.seats);
    }
    return obj;
  },

  create<I extends Exact<DeepPartial<CreateCoworkingRequestWorkspace>, I>>(base?: I): CreateCoworkingRequestWorkspace {
    return CreateCoworkingRequestWorkspace.fromPartial(base ?? ({} as any));
  },
  fromPartial<I extends Exact<DeepPartial<CreateCoworkingRequestWorkspace>, I>>(
    object: I,
  ): CreateCoworkingRequestWorkspace {
    const message = createBaseCreateCoworkingRequestWorkspace();
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
      CreateCoworkingRequestWorkspace.encode(v!, writer.uint32(18).fork()).ldelim();
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

          message.workspaces.push(CreateCoworkingRequestWorkspace.decode(reader, reader.uint32()));
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
        ? object.workspaces.map((e: any) => CreateCoworkingRequestWorkspace.fromJSON(e))
        : [],
    };
  },

  toJSON(message: CreateCoworkingRequest): unknown {
    const obj: any = {};
    if (message.name !== "") {
      obj.name = message.name;
    }
    if (message.workspaces?.length) {
      obj.workspaces = message.workspaces.map((e) => CreateCoworkingRequestWorkspace.toJSON(e));
    }
    return obj;
  },

  create<I extends Exact<DeepPartial<CreateCoworkingRequest>, I>>(base?: I): CreateCoworkingRequest {
    return CreateCoworkingRequest.fromPartial(base ?? ({} as any));
  },
  fromPartial<I extends Exact<DeepPartial<CreateCoworkingRequest>, I>>(object: I): CreateCoworkingRequest {
    const message = createBaseCreateCoworkingRequest();
    message.name = object.name ?? "";
    message.workspaces = object.workspaces?.map((e) => CreateCoworkingRequestWorkspace.fromPartial(e)) || [];
    return message;
  },
};

function createBaseAddWorkspaceToCoworkingRequest(): AddWorkspaceToCoworkingRequest {
  return { name: "", seats: 0, coworkingUuid: "" };
}

export const AddWorkspaceToCoworkingRequest = {
  encode(message: AddWorkspaceToCoworkingRequest, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
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

  decode(input: _m0.Reader | Uint8Array, length?: number): AddWorkspaceToCoworkingRequest {
    const reader = input instanceof _m0.Reader ? input : _m0.Reader.create(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseAddWorkspaceToCoworkingRequest();
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

  fromJSON(object: any): AddWorkspaceToCoworkingRequest {
    return {
      name: isSet(object.name) ? globalThis.String(object.name) : "",
      seats: isSet(object.seats) ? globalThis.Number(object.seats) : 0,
      coworkingUuid: isSet(object.coworkingUuid) ? globalThis.String(object.coworkingUuid) : "",
    };
  },

  toJSON(message: AddWorkspaceToCoworkingRequest): unknown {
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

  create<I extends Exact<DeepPartial<AddWorkspaceToCoworkingRequest>, I>>(base?: I): AddWorkspaceToCoworkingRequest {
    return AddWorkspaceToCoworkingRequest.fromPartial(base ?? ({} as any));
  },
  fromPartial<I extends Exact<DeepPartial<AddWorkspaceToCoworkingRequest>, I>>(
    object: I,
  ): AddWorkspaceToCoworkingRequest {
    const message = createBaseAddWorkspaceToCoworkingRequest();
    message.name = object.name ?? "";
    message.seats = object.seats ?? 0;
    message.coworkingUuid = object.coworkingUuid ?? "";
    return message;
  },
};

function bytesFromBase64(b64: string): Uint8Array {
  if ((globalThis as any).Buffer) {
    return Uint8Array.from(globalThis.Buffer.from(b64, "base64"));
  } else {
    const bin = globalThis.atob(b64);
    const arr = new Uint8Array(bin.length);
    for (let i = 0; i < bin.length; ++i) {
      arr[i] = bin.charCodeAt(i);
    }
    return arr;
  }
}

function base64FromBytes(arr: Uint8Array): string {
  if ((globalThis as any).Buffer) {
    return globalThis.Buffer.from(arr).toString("base64");
  } else {
    const bin: string[] = [];
    arr.forEach((byte) => {
      bin.push(globalThis.String.fromCharCode(byte));
    });
    return globalThis.btoa(bin.join(""));
  }
}

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