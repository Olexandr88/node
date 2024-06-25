// @generated by protoc-gen-es v1.3.0 with parameter "target=dts"
// @generated from file zetachain/zetacore/authority/tx.proto (package zetachain.zetacore.authority, syntax proto3)
/* eslint-disable */
// @ts-nocheck

import type { BinaryReadOptions, FieldList, JsonReadOptions, JsonValue, PartialMessage, PlainMessage } from "@bufbuild/protobuf";
import { Message, proto3 } from "@bufbuild/protobuf";
import type { Policies, PolicyType } from "./policies_pb.js";
import type { ChainInfo } from "./chain_info_pb.js";

/**
 * MsgAddAuthorization defines the MsgAddAuthorization service.
 * Adds an authorization to the chain. If the authorization already exists, it
 * will be updated.
 *
 * @generated from message zetachain.zetacore.authority.MsgAddAuthorization
 */
export declare class MsgAddAuthorization extends Message<MsgAddAuthorization> {
  /**
   * @generated from field: string creator = 1;
   */
  creator: string;

  /**
   * @generated from field: string msg_url = 2;
   */
  msgUrl: string;

  /**
   * @generated from field: zetachain.zetacore.authority.PolicyType authorized_policy = 3;
   */
  authorizedPolicy: PolicyType;

  constructor(data?: PartialMessage<MsgAddAuthorization>);

  static readonly runtime: typeof proto3;
  static readonly typeName = "zetachain.zetacore.authority.MsgAddAuthorization";
  static readonly fields: FieldList;

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): MsgAddAuthorization;

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): MsgAddAuthorization;

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): MsgAddAuthorization;

  static equals(a: MsgAddAuthorization | PlainMessage<MsgAddAuthorization> | undefined, b: MsgAddAuthorization | PlainMessage<MsgAddAuthorization> | undefined): boolean;
}

/**
 * MsgAddAuthorizationResponse defines the MsgAddAuthorizationResponse service.
 *
 * @generated from message zetachain.zetacore.authority.MsgAddAuthorizationResponse
 */
export declare class MsgAddAuthorizationResponse extends Message<MsgAddAuthorizationResponse> {
  constructor(data?: PartialMessage<MsgAddAuthorizationResponse>);

  static readonly runtime: typeof proto3;
  static readonly typeName = "zetachain.zetacore.authority.MsgAddAuthorizationResponse";
  static readonly fields: FieldList;

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): MsgAddAuthorizationResponse;

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): MsgAddAuthorizationResponse;

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): MsgAddAuthorizationResponse;

  static equals(a: MsgAddAuthorizationResponse | PlainMessage<MsgAddAuthorizationResponse> | undefined, b: MsgAddAuthorizationResponse | PlainMessage<MsgAddAuthorizationResponse> | undefined): boolean;
}

/**
 * MsgRemoveAuthorization defines the MsgRemoveAuthorization service.
 * Removes an authorization from the chain.
 *
 * @generated from message zetachain.zetacore.authority.MsgRemoveAuthorization
 */
export declare class MsgRemoveAuthorization extends Message<MsgRemoveAuthorization> {
  /**
   * @generated from field: string creator = 1;
   */
  creator: string;

  /**
   * @generated from field: string msg_url = 2;
   */
  msgUrl: string;

  constructor(data?: PartialMessage<MsgRemoveAuthorization>);

  static readonly runtime: typeof proto3;
  static readonly typeName = "zetachain.zetacore.authority.MsgRemoveAuthorization";
  static readonly fields: FieldList;

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): MsgRemoveAuthorization;

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): MsgRemoveAuthorization;

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): MsgRemoveAuthorization;

  static equals(a: MsgRemoveAuthorization | PlainMessage<MsgRemoveAuthorization> | undefined, b: MsgRemoveAuthorization | PlainMessage<MsgRemoveAuthorization> | undefined): boolean;
}

/**
 * MsgRemoveAuthorizationResponse defines the MsgRemoveAuthorizationResponse
 * service.
 *
 * @generated from message zetachain.zetacore.authority.MsgRemoveAuthorizationResponse
 */
export declare class MsgRemoveAuthorizationResponse extends Message<MsgRemoveAuthorizationResponse> {
  constructor(data?: PartialMessage<MsgRemoveAuthorizationResponse>);

  static readonly runtime: typeof proto3;
  static readonly typeName = "zetachain.zetacore.authority.MsgRemoveAuthorizationResponse";
  static readonly fields: FieldList;

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): MsgRemoveAuthorizationResponse;

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): MsgRemoveAuthorizationResponse;

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): MsgRemoveAuthorizationResponse;

  static equals(a: MsgRemoveAuthorizationResponse | PlainMessage<MsgRemoveAuthorizationResponse> | undefined, b: MsgRemoveAuthorizationResponse | PlainMessage<MsgRemoveAuthorizationResponse> | undefined): boolean;
}

/**
 * MsgUpdatePolicies defines the MsgUpdatePolicies service.
 *
 * @generated from message zetachain.zetacore.authority.MsgUpdatePolicies
 */
export declare class MsgUpdatePolicies extends Message<MsgUpdatePolicies> {
  /**
   * @generated from field: string creator = 1;
   */
  creator: string;

  /**
   * @generated from field: zetachain.zetacore.authority.Policies policies = 2;
   */
  policies?: Policies;

  constructor(data?: PartialMessage<MsgUpdatePolicies>);

  static readonly runtime: typeof proto3;
  static readonly typeName = "zetachain.zetacore.authority.MsgUpdatePolicies";
  static readonly fields: FieldList;

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): MsgUpdatePolicies;

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): MsgUpdatePolicies;

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): MsgUpdatePolicies;

  static equals(a: MsgUpdatePolicies | PlainMessage<MsgUpdatePolicies> | undefined, b: MsgUpdatePolicies | PlainMessage<MsgUpdatePolicies> | undefined): boolean;
}

/**
 * MsgUpdatePoliciesResponse defines the MsgUpdatePoliciesResponse service.
 *
 * @generated from message zetachain.zetacore.authority.MsgUpdatePoliciesResponse
 */
export declare class MsgUpdatePoliciesResponse extends Message<MsgUpdatePoliciesResponse> {
  constructor(data?: PartialMessage<MsgUpdatePoliciesResponse>);

  static readonly runtime: typeof proto3;
  static readonly typeName = "zetachain.zetacore.authority.MsgUpdatePoliciesResponse";
  static readonly fields: FieldList;

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): MsgUpdatePoliciesResponse;

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): MsgUpdatePoliciesResponse;

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): MsgUpdatePoliciesResponse;

  static equals(a: MsgUpdatePoliciesResponse | PlainMessage<MsgUpdatePoliciesResponse> | undefined, b: MsgUpdatePoliciesResponse | PlainMessage<MsgUpdatePoliciesResponse> | undefined): boolean;
}

/**
 * MsgUpdateChainInfo defines the MsgUpdateChainInfo service.
 *
 * @generated from message zetachain.zetacore.authority.MsgUpdateChainInfo
 */
export declare class MsgUpdateChainInfo extends Message<MsgUpdateChainInfo> {
  /**
   * @generated from field: string creator = 1;
   */
  creator: string;

  /**
   * @generated from field: zetachain.zetacore.authority.ChainInfo chain_info = 2;
   */
  chainInfo?: ChainInfo;

  constructor(data?: PartialMessage<MsgUpdateChainInfo>);

  static readonly runtime: typeof proto3;
  static readonly typeName = "zetachain.zetacore.authority.MsgUpdateChainInfo";
  static readonly fields: FieldList;

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): MsgUpdateChainInfo;

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): MsgUpdateChainInfo;

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): MsgUpdateChainInfo;

  static equals(a: MsgUpdateChainInfo | PlainMessage<MsgUpdateChainInfo> | undefined, b: MsgUpdateChainInfo | PlainMessage<MsgUpdateChainInfo> | undefined): boolean;
}

/**
 * MsgUpdateChainInfoResponse defines the MsgUpdateChainInfoResponse service.
 *
 * @generated from message zetachain.zetacore.authority.MsgUpdateChainInfoResponse
 */
export declare class MsgUpdateChainInfoResponse extends Message<MsgUpdateChainInfoResponse> {
  constructor(data?: PartialMessage<MsgUpdateChainInfoResponse>);

  static readonly runtime: typeof proto3;
  static readonly typeName = "zetachain.zetacore.authority.MsgUpdateChainInfoResponse";
  static readonly fields: FieldList;

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): MsgUpdateChainInfoResponse;

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): MsgUpdateChainInfoResponse;

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): MsgUpdateChainInfoResponse;

  static equals(a: MsgUpdateChainInfoResponse | PlainMessage<MsgUpdateChainInfoResponse> | undefined, b: MsgUpdateChainInfoResponse | PlainMessage<MsgUpdateChainInfoResponse> | undefined): boolean;
}
