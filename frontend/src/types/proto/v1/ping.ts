/* eslint-disable */
import _m0 from 'protobufjs/minimal';

export const protobufPackage = 'dauth.v1';

export interface PingRequest {}

export interface PingResponse {
  pong: string;
}

function createBasePingRequest(): PingRequest {
  return {};
}

export const PingRequest = {
  encode(_: PingRequest, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): PingRequest {
    const reader = input instanceof _m0.Reader ? input : _m0.Reader.create(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBasePingRequest();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
      }
      if ((tag & 7) === 4 || tag === 0) {
        break;
      }
      reader.skipType(tag & 7);
    }
    return message;
  },

  fromJSON(_: any): PingRequest {
    return {};
  },

  toJSON(_: PingRequest): unknown {
    const obj: any = {};
    return obj;
  },

  create(base?: DeepPartial<PingRequest>): PingRequest {
    return PingRequest.fromPartial(base ?? {});
  },
  fromPartial(_: DeepPartial<PingRequest>): PingRequest {
    const message = createBasePingRequest();
    return message;
  },
};

function createBasePingResponse(): PingResponse {
  return { pong: '' };
}

export const PingResponse = {
  encode(message: PingResponse, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.pong !== '') {
      writer.uint32(10).string(message.pong);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): PingResponse {
    const reader = input instanceof _m0.Reader ? input : _m0.Reader.create(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBasePingResponse();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          if (tag !== 10) {
            break;
          }

          message.pong = reader.string();
          continue;
      }
      if ((tag & 7) === 4 || tag === 0) {
        break;
      }
      reader.skipType(tag & 7);
    }
    return message;
  },

  fromJSON(object: any): PingResponse {
    return { pong: isSet(object.pong) ? globalThis.String(object.pong) : '' };
  },

  toJSON(message: PingResponse): unknown {
    const obj: any = {};
    if (message.pong !== '') {
      obj.pong = message.pong;
    }
    return obj;
  },

  create(base?: DeepPartial<PingResponse>): PingResponse {
    return PingResponse.fromPartial(base ?? {});
  },
  fromPartial(object: DeepPartial<PingResponse>): PingResponse {
    const message = createBasePingResponse();
    message.pong = object.pong ?? '';
    return message;
  },
};

export type PingServiceDefinition = typeof PingServiceDefinition;
export const PingServiceDefinition = {
  name: 'PingService',
  fullName: 'dauth.v1.PingService',
  methods: {
    ping: {
      name: 'Ping',
      requestType: PingRequest,
      requestStream: false,
      responseType: PingResponse,
      responseStream: false,
      options: { _unknownFields: { 578365826: [new Uint8Array([10, 18, 8, 47, 118, 49, 47, 112, 105, 110, 103])] } },
    },
  },
} as const;

type Builtin = Date | Function | Uint8Array | string | number | boolean | undefined;

export type DeepPartial<T> = T extends Builtin
  ? T
  : T extends globalThis.Array<infer U>
    ? globalThis.Array<DeepPartial<U>>
    : T extends ReadonlyArray<infer U>
      ? ReadonlyArray<DeepPartial<U>>
      : T extends {}
        ? { [K in keyof T]?: DeepPartial<T[K]> }
        : Partial<T>;

function isSet(value: any): boolean {
  return value !== null && value !== undefined;
}
