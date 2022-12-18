import * as jspb from 'google-protobuf'



export class Account extends jspb.Message {
  getEmail(): string;
  setEmail(value: string): Account;

  getName(): string;
  setName(value: string): Account;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): Account.AsObject;
  static toObject(includeInstance: boolean, msg: Account): Account.AsObject;
  static serializeBinaryToWriter(message: Account, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): Account;
  static deserializeBinaryFromReader(message: Account, reader: jspb.BinaryReader): Account;
}

export namespace Account {
  export type AsObject = {
    email: string,
    name: string,
  }
}

