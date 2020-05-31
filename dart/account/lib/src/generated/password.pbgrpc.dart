///
//  Generated code. Do not modify.
//  source: password.proto
//
// @dart = 2.3
// ignore_for_file: camel_case_types,non_constant_identifier_names,library_prefixes,unused_import,unused_shown_name,return_of_invalid_type

import 'dart:async' as $async;

import 'dart:core' as $core;

import 'package:grpc/service_api.dart' as $grpc;
import 'password.pb.dart' as $8;
import 'common.pb.dart' as $1;
export 'password.pb.dart';

class PasswordClient extends $grpc.Client {
  static final _$set = $grpc.ClientMethod<$8.PasswordRequest, $1.Secret>(
      '/appootb.account.Password/Set',
      ($8.PasswordRequest value) => value.writeToBuffer(),
      ($core.List<$core.int> value) => $1.Secret.fromBuffer(value));
  static final _$update = $grpc.ClientMethod<$8.PasswordRequest, $1.Secret>(
      '/appootb.account.Password/Update',
      ($8.PasswordRequest value) => value.writeToBuffer(),
      ($core.List<$core.int> value) => $1.Secret.fromBuffer(value));
  static final _$reset = $grpc.ClientMethod<$8.PasswordRequest, $1.Secret>(
      '/appootb.account.Password/Reset',
      ($8.PasswordRequest value) => value.writeToBuffer(),
      ($core.List<$core.int> value) => $1.Secret.fromBuffer(value));

  PasswordClient($grpc.ClientChannel channel, {$grpc.CallOptions options})
      : super(channel, options: options);

  $grpc.ResponseFuture<$1.Secret> set($8.PasswordRequest request,
      {$grpc.CallOptions options}) {
    final call = $createCall(_$set, $async.Stream.fromIterable([request]),
        options: options);
    return $grpc.ResponseFuture(call);
  }

  $grpc.ResponseFuture<$1.Secret> update($8.PasswordRequest request,
      {$grpc.CallOptions options}) {
    final call = $createCall(_$update, $async.Stream.fromIterable([request]),
        options: options);
    return $grpc.ResponseFuture(call);
  }

  $grpc.ResponseFuture<$1.Secret> reset($8.PasswordRequest request,
      {$grpc.CallOptions options}) {
    final call = $createCall(_$reset, $async.Stream.fromIterable([request]),
        options: options);
    return $grpc.ResponseFuture(call);
  }
}

abstract class PasswordServiceBase extends $grpc.Service {
  $core.String get $name => 'appootb.account.Password';

  PasswordServiceBase() {
    $addMethod($grpc.ServiceMethod<$8.PasswordRequest, $1.Secret>(
        'Set',
        set_Pre,
        false,
        false,
        ($core.List<$core.int> value) => $8.PasswordRequest.fromBuffer(value),
        ($1.Secret value) => value.writeToBuffer()));
    $addMethod($grpc.ServiceMethod<$8.PasswordRequest, $1.Secret>(
        'Update',
        update_Pre,
        false,
        false,
        ($core.List<$core.int> value) => $8.PasswordRequest.fromBuffer(value),
        ($1.Secret value) => value.writeToBuffer()));
    $addMethod($grpc.ServiceMethod<$8.PasswordRequest, $1.Secret>(
        'Reset',
        reset_Pre,
        false,
        false,
        ($core.List<$core.int> value) => $8.PasswordRequest.fromBuffer(value),
        ($1.Secret value) => value.writeToBuffer()));
  }

  $async.Future<$1.Secret> set_Pre(
      $grpc.ServiceCall call, $async.Future<$8.PasswordRequest> request) async {
    return set(call, await request);
  }

  $async.Future<$1.Secret> update_Pre(
      $grpc.ServiceCall call, $async.Future<$8.PasswordRequest> request) async {
    return update(call, await request);
  }

  $async.Future<$1.Secret> reset_Pre(
      $grpc.ServiceCall call, $async.Future<$8.PasswordRequest> request) async {
    return reset(call, await request);
  }

  $async.Future<$1.Secret> set(
      $grpc.ServiceCall call, $8.PasswordRequest request);
  $async.Future<$1.Secret> update(
      $grpc.ServiceCall call, $8.PasswordRequest request);
  $async.Future<$1.Secret> reset(
      $grpc.ServiceCall call, $8.PasswordRequest request);
}
