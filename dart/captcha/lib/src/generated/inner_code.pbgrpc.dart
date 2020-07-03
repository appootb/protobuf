///
//  Generated code. Do not modify.
//  source: inner_code.proto
//
// @dart = 2.3
// ignore_for_file: camel_case_types,non_constant_identifier_names,library_prefixes,unused_import,unused_shown_name,return_of_invalid_type

import 'dart:async' as $async;

import 'dart:core' as $core;

import 'package:grpc/service_api.dart' as $grpc;
import 'include.pb.dart' as $0;
import 'google/protobuf/empty.pb.dart' as $1;
export 'inner_code.pb.dart';

class InnerCodeClient extends $grpc.Client {
  static final _$launch = $grpc.ClientMethod<$0.CodeRequest, $1.Empty>(
      '/appootb.captcha.InnerCode/Launch',
      ($0.CodeRequest value) => value.writeToBuffer(),
      ($core.List<$core.int> value) => $1.Empty.fromBuffer(value));
  static final _$verify = $grpc.ClientMethod<$0.CodeRequest, $1.Empty>(
      '/appootb.captcha.InnerCode/Verify',
      ($0.CodeRequest value) => value.writeToBuffer(),
      ($core.List<$core.int> value) => $1.Empty.fromBuffer(value));

  InnerCodeClient($grpc.ClientChannel channel, {$grpc.CallOptions options})
      : super(channel, options: options);

  $grpc.ResponseFuture<$1.Empty> launch($0.CodeRequest request,
      {$grpc.CallOptions options}) {
    final call = $createCall(_$launch, $async.Stream.fromIterable([request]),
        options: options);
    return $grpc.ResponseFuture(call);
  }

  $grpc.ResponseFuture<$1.Empty> verify($0.CodeRequest request,
      {$grpc.CallOptions options}) {
    final call = $createCall(_$verify, $async.Stream.fromIterable([request]),
        options: options);
    return $grpc.ResponseFuture(call);
  }
}

abstract class InnerCodeServiceBase extends $grpc.Service {
  $core.String get $name => 'appootb.captcha.InnerCode';

  InnerCodeServiceBase() {
    $addMethod($grpc.ServiceMethod<$0.CodeRequest, $1.Empty>(
        'Launch',
        launch_Pre,
        false,
        false,
        ($core.List<$core.int> value) => $0.CodeRequest.fromBuffer(value),
        ($1.Empty value) => value.writeToBuffer()));
    $addMethod($grpc.ServiceMethod<$0.CodeRequest, $1.Empty>(
        'Verify',
        verify_Pre,
        false,
        false,
        ($core.List<$core.int> value) => $0.CodeRequest.fromBuffer(value),
        ($1.Empty value) => value.writeToBuffer()));
  }

  $async.Future<$1.Empty> launch_Pre(
      $grpc.ServiceCall call, $async.Future<$0.CodeRequest> request) async {
    return launch(call, await request);
  }

  $async.Future<$1.Empty> verify_Pre(
      $grpc.ServiceCall call, $async.Future<$0.CodeRequest> request) async {
    return verify(call, await request);
  }

  $async.Future<$1.Empty> launch(
      $grpc.ServiceCall call, $0.CodeRequest request);
  $async.Future<$1.Empty> verify(
      $grpc.ServiceCall call, $0.CodeRequest request);
}
