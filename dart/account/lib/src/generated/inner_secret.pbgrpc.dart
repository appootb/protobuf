///
//  Generated code. Do not modify.
//  source: inner_secret.proto
//
// @dart = 2.3
// ignore_for_file: camel_case_types,non_constant_identifier_names,library_prefixes,unused_import,unused_shown_name,return_of_invalid_type

import 'dart:async' as $async;

import 'dart:core' as $core;

import 'package:grpc/service_api.dart' as $grpc;
import 'common.pb.dart' as $1;
import 'inner_secret.pb.dart' as $7;
export 'inner_secret.pb.dart';

class InnerSecretClient extends $grpc.Client {
  static final _$getSecretInfo = $grpc.ClientMethod<$1.Secret, $7.SecretInfo>(
      '/appootb.account.InnerSecret/GetSecretInfo',
      ($1.Secret value) => value.writeToBuffer(),
      ($core.List<$core.int> value) => $7.SecretInfo.fromBuffer(value));

  InnerSecretClient($grpc.ClientChannel channel, {$grpc.CallOptions options})
      : super(channel, options: options);

  $grpc.ResponseFuture<$7.SecretInfo> getSecretInfo($1.Secret request,
      {$grpc.CallOptions options}) {
    final call = $createCall(
        _$getSecretInfo, $async.Stream.fromIterable([request]),
        options: options);
    return $grpc.ResponseFuture(call);
  }
}

abstract class InnerSecretServiceBase extends $grpc.Service {
  $core.String get $name => 'appootb.account.InnerSecret';

  InnerSecretServiceBase() {
    $addMethod($grpc.ServiceMethod<$1.Secret, $7.SecretInfo>(
        'GetSecretInfo',
        getSecretInfo_Pre,
        false,
        false,
        ($core.List<$core.int> value) => $1.Secret.fromBuffer(value),
        ($7.SecretInfo value) => value.writeToBuffer()));
  }

  $async.Future<$7.SecretInfo> getSecretInfo_Pre(
      $grpc.ServiceCall call, $async.Future<$1.Secret> request) async {
    return getSecretInfo(call, await request);
  }

  $async.Future<$7.SecretInfo> getSecretInfo(
      $grpc.ServiceCall call, $1.Secret request);
}
