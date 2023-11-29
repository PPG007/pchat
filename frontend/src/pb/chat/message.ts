/* eslint-disable */
import _m0 from "protobufjs/minimal.js";

export const protobufPackage = "pchat.chat";

export interface NewMessage {
  id: string;
  replyTo: string;
  threadId: string;
  isInThread: boolean;
  showThreadInChat: boolean;
  type: string;
  content: string;
  fileUrl: string;
  mentionedUsers: string[];
  chatId: string;
}

export interface MessageDetail {
  id: string;
  replyTo: string;
  threadId: string;
  isInThread: boolean;
  showThreadInChat: boolean;
  type: string;
  content: string;
  fileUrl: string;
  mentionedUsers: string[];
  chatId: string;
  createdAt: string;
  hasBeenEdited: boolean;
  responseEmojis: string[];
}

function createBaseNewMessage(): NewMessage {
  return {
    id: "",
    replyTo: "",
    threadId: "",
    isInThread: false,
    showThreadInChat: false,
    type: "",
    content: "",
    fileUrl: "",
    mentionedUsers: [],
    chatId: "",
  };
}

export const NewMessage = {
  encode(message: NewMessage, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.id !== "") {
      writer.uint32(10).string(message.id);
    }
    if (message.replyTo !== "") {
      writer.uint32(18).string(message.replyTo);
    }
    if (message.threadId !== "") {
      writer.uint32(26).string(message.threadId);
    }
    if (message.isInThread === true) {
      writer.uint32(32).bool(message.isInThread);
    }
    if (message.showThreadInChat === true) {
      writer.uint32(40).bool(message.showThreadInChat);
    }
    if (message.type !== "") {
      writer.uint32(50).string(message.type);
    }
    if (message.content !== "") {
      writer.uint32(58).string(message.content);
    }
    if (message.fileUrl !== "") {
      writer.uint32(66).string(message.fileUrl);
    }
    for (const v of message.mentionedUsers) {
      writer.uint32(74).string(v!);
    }
    if (message.chatId !== "") {
      writer.uint32(82).string(message.chatId);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): NewMessage {
    const reader = input instanceof _m0.Reader ? input : _m0.Reader.create(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseNewMessage();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          if (tag !== 10) {
            break;
          }

          message.id = reader.string();
          continue;
        case 2:
          if (tag !== 18) {
            break;
          }

          message.replyTo = reader.string();
          continue;
        case 3:
          if (tag !== 26) {
            break;
          }

          message.threadId = reader.string();
          continue;
        case 4:
          if (tag !== 32) {
            break;
          }

          message.isInThread = reader.bool();
          continue;
        case 5:
          if (tag !== 40) {
            break;
          }

          message.showThreadInChat = reader.bool();
          continue;
        case 6:
          if (tag !== 50) {
            break;
          }

          message.type = reader.string();
          continue;
        case 7:
          if (tag !== 58) {
            break;
          }

          message.content = reader.string();
          continue;
        case 8:
          if (tag !== 66) {
            break;
          }

          message.fileUrl = reader.string();
          continue;
        case 9:
          if (tag !== 74) {
            break;
          }

          message.mentionedUsers.push(reader.string());
          continue;
        case 10:
          if (tag !== 82) {
            break;
          }

          message.chatId = reader.string();
          continue;
      }
      if ((tag & 7) === 4 || tag === 0) {
        break;
      }
      reader.skipType(tag & 7);
    }
    return message;
  },

  fromJSON(object: any): NewMessage {
    return {
      id: isSet(object.id) ? globalThis.String(object.id) : "",
      replyTo: isSet(object.replyTo) ? globalThis.String(object.replyTo) : "",
      threadId: isSet(object.threadId) ? globalThis.String(object.threadId) : "",
      isInThread: isSet(object.isInThread) ? globalThis.Boolean(object.isInThread) : false,
      showThreadInChat: isSet(object.showThreadInChat) ? globalThis.Boolean(object.showThreadInChat) : false,
      type: isSet(object.type) ? globalThis.String(object.type) : "",
      content: isSet(object.content) ? globalThis.String(object.content) : "",
      fileUrl: isSet(object.fileUrl) ? globalThis.String(object.fileUrl) : "",
      mentionedUsers: globalThis.Array.isArray(object?.mentionedUsers)
        ? object.mentionedUsers.map((e: any) => globalThis.String(e))
        : [],
      chatId: isSet(object.chatId) ? globalThis.String(object.chatId) : "",
    };
  },

  toJSON(message: NewMessage): unknown {
    const obj: any = {};
    if (message.id !== "") {
      obj.id = message.id;
    }
    if (message.replyTo !== "") {
      obj.replyTo = message.replyTo;
    }
    if (message.threadId !== "") {
      obj.threadId = message.threadId;
    }
    if (message.isInThread === true) {
      obj.isInThread = message.isInThread;
    }
    if (message.showThreadInChat === true) {
      obj.showThreadInChat = message.showThreadInChat;
    }
    if (message.type !== "") {
      obj.type = message.type;
    }
    if (message.content !== "") {
      obj.content = message.content;
    }
    if (message.fileUrl !== "") {
      obj.fileUrl = message.fileUrl;
    }
    if (message.mentionedUsers?.length) {
      obj.mentionedUsers = message.mentionedUsers;
    }
    if (message.chatId !== "") {
      obj.chatId = message.chatId;
    }
    return obj;
  },

  create<I extends Exact<DeepPartial<NewMessage>, I>>(base?: I): NewMessage {
    return NewMessage.fromPartial(base ?? ({} as any));
  },
  fromPartial<I extends Exact<DeepPartial<NewMessage>, I>>(object: I): NewMessage {
    const message = createBaseNewMessage();
    message.id = object.id ?? "";
    message.replyTo = object.replyTo ?? "";
    message.threadId = object.threadId ?? "";
    message.isInThread = object.isInThread ?? false;
    message.showThreadInChat = object.showThreadInChat ?? false;
    message.type = object.type ?? "";
    message.content = object.content ?? "";
    message.fileUrl = object.fileUrl ?? "";
    message.mentionedUsers = object.mentionedUsers?.map((e) => e) || [];
    message.chatId = object.chatId ?? "";
    return message;
  },
};

function createBaseMessageDetail(): MessageDetail {
  return {
    id: "",
    replyTo: "",
    threadId: "",
    isInThread: false,
    showThreadInChat: false,
    type: "",
    content: "",
    fileUrl: "",
    mentionedUsers: [],
    chatId: "",
    createdAt: "",
    hasBeenEdited: false,
    responseEmojis: [],
  };
}

export const MessageDetail = {
  encode(message: MessageDetail, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.id !== "") {
      writer.uint32(10).string(message.id);
    }
    if (message.replyTo !== "") {
      writer.uint32(18).string(message.replyTo);
    }
    if (message.threadId !== "") {
      writer.uint32(26).string(message.threadId);
    }
    if (message.isInThread === true) {
      writer.uint32(32).bool(message.isInThread);
    }
    if (message.showThreadInChat === true) {
      writer.uint32(40).bool(message.showThreadInChat);
    }
    if (message.type !== "") {
      writer.uint32(50).string(message.type);
    }
    if (message.content !== "") {
      writer.uint32(58).string(message.content);
    }
    if (message.fileUrl !== "") {
      writer.uint32(66).string(message.fileUrl);
    }
    for (const v of message.mentionedUsers) {
      writer.uint32(74).string(v!);
    }
    if (message.chatId !== "") {
      writer.uint32(82).string(message.chatId);
    }
    if (message.createdAt !== "") {
      writer.uint32(90).string(message.createdAt);
    }
    if (message.hasBeenEdited === true) {
      writer.uint32(96).bool(message.hasBeenEdited);
    }
    for (const v of message.responseEmojis) {
      writer.uint32(106).string(v!);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): MessageDetail {
    const reader = input instanceof _m0.Reader ? input : _m0.Reader.create(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseMessageDetail();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          if (tag !== 10) {
            break;
          }

          message.id = reader.string();
          continue;
        case 2:
          if (tag !== 18) {
            break;
          }

          message.replyTo = reader.string();
          continue;
        case 3:
          if (tag !== 26) {
            break;
          }

          message.threadId = reader.string();
          continue;
        case 4:
          if (tag !== 32) {
            break;
          }

          message.isInThread = reader.bool();
          continue;
        case 5:
          if (tag !== 40) {
            break;
          }

          message.showThreadInChat = reader.bool();
          continue;
        case 6:
          if (tag !== 50) {
            break;
          }

          message.type = reader.string();
          continue;
        case 7:
          if (tag !== 58) {
            break;
          }

          message.content = reader.string();
          continue;
        case 8:
          if (tag !== 66) {
            break;
          }

          message.fileUrl = reader.string();
          continue;
        case 9:
          if (tag !== 74) {
            break;
          }

          message.mentionedUsers.push(reader.string());
          continue;
        case 10:
          if (tag !== 82) {
            break;
          }

          message.chatId = reader.string();
          continue;
        case 11:
          if (tag !== 90) {
            break;
          }

          message.createdAt = reader.string();
          continue;
        case 12:
          if (tag !== 96) {
            break;
          }

          message.hasBeenEdited = reader.bool();
          continue;
        case 13:
          if (tag !== 106) {
            break;
          }

          message.responseEmojis.push(reader.string());
          continue;
      }
      if ((tag & 7) === 4 || tag === 0) {
        break;
      }
      reader.skipType(tag & 7);
    }
    return message;
  },

  fromJSON(object: any): MessageDetail {
    return {
      id: isSet(object.id) ? globalThis.String(object.id) : "",
      replyTo: isSet(object.replyTo) ? globalThis.String(object.replyTo) : "",
      threadId: isSet(object.threadId) ? globalThis.String(object.threadId) : "",
      isInThread: isSet(object.isInThread) ? globalThis.Boolean(object.isInThread) : false,
      showThreadInChat: isSet(object.showThreadInChat) ? globalThis.Boolean(object.showThreadInChat) : false,
      type: isSet(object.type) ? globalThis.String(object.type) : "",
      content: isSet(object.content) ? globalThis.String(object.content) : "",
      fileUrl: isSet(object.fileUrl) ? globalThis.String(object.fileUrl) : "",
      mentionedUsers: globalThis.Array.isArray(object?.mentionedUsers)
        ? object.mentionedUsers.map((e: any) => globalThis.String(e))
        : [],
      chatId: isSet(object.chatId) ? globalThis.String(object.chatId) : "",
      createdAt: isSet(object.createdAt) ? globalThis.String(object.createdAt) : "",
      hasBeenEdited: isSet(object.hasBeenEdited) ? globalThis.Boolean(object.hasBeenEdited) : false,
      responseEmojis: globalThis.Array.isArray(object?.responseEmojis)
        ? object.responseEmojis.map((e: any) => globalThis.String(e))
        : [],
    };
  },

  toJSON(message: MessageDetail): unknown {
    const obj: any = {};
    if (message.id !== "") {
      obj.id = message.id;
    }
    if (message.replyTo !== "") {
      obj.replyTo = message.replyTo;
    }
    if (message.threadId !== "") {
      obj.threadId = message.threadId;
    }
    if (message.isInThread === true) {
      obj.isInThread = message.isInThread;
    }
    if (message.showThreadInChat === true) {
      obj.showThreadInChat = message.showThreadInChat;
    }
    if (message.type !== "") {
      obj.type = message.type;
    }
    if (message.content !== "") {
      obj.content = message.content;
    }
    if (message.fileUrl !== "") {
      obj.fileUrl = message.fileUrl;
    }
    if (message.mentionedUsers?.length) {
      obj.mentionedUsers = message.mentionedUsers;
    }
    if (message.chatId !== "") {
      obj.chatId = message.chatId;
    }
    if (message.createdAt !== "") {
      obj.createdAt = message.createdAt;
    }
    if (message.hasBeenEdited === true) {
      obj.hasBeenEdited = message.hasBeenEdited;
    }
    if (message.responseEmojis?.length) {
      obj.responseEmojis = message.responseEmojis;
    }
    return obj;
  },

  create<I extends Exact<DeepPartial<MessageDetail>, I>>(base?: I): MessageDetail {
    return MessageDetail.fromPartial(base ?? ({} as any));
  },
  fromPartial<I extends Exact<DeepPartial<MessageDetail>, I>>(object: I): MessageDetail {
    const message = createBaseMessageDetail();
    message.id = object.id ?? "";
    message.replyTo = object.replyTo ?? "";
    message.threadId = object.threadId ?? "";
    message.isInThread = object.isInThread ?? false;
    message.showThreadInChat = object.showThreadInChat ?? false;
    message.type = object.type ?? "";
    message.content = object.content ?? "";
    message.fileUrl = object.fileUrl ?? "";
    message.mentionedUsers = object.mentionedUsers?.map((e) => e) || [];
    message.chatId = object.chatId ?? "";
    message.createdAt = object.createdAt ?? "";
    message.hasBeenEdited = object.hasBeenEdited ?? false;
    message.responseEmojis = object.responseEmojis?.map((e) => e) || [];
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
