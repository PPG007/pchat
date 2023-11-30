/* eslint-disable */
import _m0 from "protobufjs/minimal.js";

export const protobufPackage = "pchat.user";

export interface LoginResponse {
  token: string;
  id: string;
  name: string;
  permissions: string[];
  email: string;
  avatar: string;
}

export interface RegisterResponse {
  needAudit: boolean;
}

function createBaseLoginResponse(): LoginResponse {
  return { token: "", id: "", name: "", permissions: [], email: "", avatar: "" };
}

export const LoginResponse = {
  encode(message: LoginResponse, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.token !== "") {
      writer.uint32(10).string(message.token);
    }
    if (message.id !== "") {
      writer.uint32(18).string(message.id);
    }
    if (message.name !== "") {
      writer.uint32(26).string(message.name);
    }
    for (const v of message.permissions) {
      writer.uint32(34).string(v!);
    }
    if (message.email !== "") {
      writer.uint32(42).string(message.email);
    }
    if (message.avatar !== "") {
      writer.uint32(50).string(message.avatar);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): LoginResponse {
    const reader = input instanceof _m0.Reader ? input : _m0.Reader.create(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseLoginResponse();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          if (tag !== 10) {
            break;
          }

          message.token = reader.string();
          continue;
        case 2:
          if (tag !== 18) {
            break;
          }

          message.id = reader.string();
          continue;
        case 3:
          if (tag !== 26) {
            break;
          }

          message.name = reader.string();
          continue;
        case 4:
          if (tag !== 34) {
            break;
          }

          message.permissions.push(reader.string());
          continue;
        case 5:
          if (tag !== 42) {
            break;
          }

          message.email = reader.string();
          continue;
        case 6:
          if (tag !== 50) {
            break;
          }

          message.avatar = reader.string();
          continue;
      }
      if ((tag & 7) === 4 || tag === 0) {
        break;
      }
      reader.skipType(tag & 7);
    }
    return message;
  },

  fromJSON(object: any): LoginResponse {
    return {
      token: isSet(object.token) ? globalThis.String(object.token) : "",
      id: isSet(object.id) ? globalThis.String(object.id) : "",
      name: isSet(object.name) ? globalThis.String(object.name) : "",
      permissions: globalThis.Array.isArray(object?.permissions)
        ? object.permissions.map((e: any) => globalThis.String(e))
        : [],
      email: isSet(object.email) ? globalThis.String(object.email) : "",
      avatar: isSet(object.avatar) ? globalThis.String(object.avatar) : "",
    };
  },

  toJSON(message: LoginResponse): unknown {
    const obj: any = {};
    if (message.token !== "") {
      obj.token = message.token;
    }
    if (message.id !== "") {
      obj.id = message.id;
    }
    if (message.name !== "") {
      obj.name = message.name;
    }
    if (message.permissions?.length) {
      obj.permissions = message.permissions;
    }
    if (message.email !== "") {
      obj.email = message.email;
    }
    if (message.avatar !== "") {
      obj.avatar = message.avatar;
    }
    return obj;
  },

  create<I extends Exact<DeepPartial<LoginResponse>, I>>(base?: I): LoginResponse {
    return LoginResponse.fromPartial(base ?? ({} as any));
  },
  fromPartial<I extends Exact<DeepPartial<LoginResponse>, I>>(object: I): LoginResponse {
    const message = createBaseLoginResponse();
    message.token = object.token ?? "";
    message.id = object.id ?? "";
    message.name = object.name ?? "";
    message.permissions = object.permissions?.map((e) => e) || [];
    message.email = object.email ?? "";
    message.avatar = object.avatar ?? "";
    return message;
  },
};

function createBaseRegisterResponse(): RegisterResponse {
  return { needAudit: false };
}

export const RegisterResponse = {
  encode(message: RegisterResponse, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.needAudit === true) {
      writer.uint32(8).bool(message.needAudit);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): RegisterResponse {
    const reader = input instanceof _m0.Reader ? input : _m0.Reader.create(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseRegisterResponse();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          if (tag !== 8) {
            break;
          }

          message.needAudit = reader.bool();
          continue;
      }
      if ((tag & 7) === 4 || tag === 0) {
        break;
      }
      reader.skipType(tag & 7);
    }
    return message;
  },

  fromJSON(object: any): RegisterResponse {
    return { needAudit: isSet(object.needAudit) ? globalThis.Boolean(object.needAudit) : false };
  },

  toJSON(message: RegisterResponse): unknown {
    const obj: any = {};
    if (message.needAudit === true) {
      obj.needAudit = message.needAudit;
    }
    return obj;
  },

  create<I extends Exact<DeepPartial<RegisterResponse>, I>>(base?: I): RegisterResponse {
    return RegisterResponse.fromPartial(base ?? ({} as any));
  },
  fromPartial<I extends Exact<DeepPartial<RegisterResponse>, I>>(object: I): RegisterResponse {
    const message = createBaseRegisterResponse();
    message.needAudit = object.needAudit ?? false;
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
