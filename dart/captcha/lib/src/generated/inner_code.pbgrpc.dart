///
//  Generated code. Do not modify.
//  source: inner_code.proto
//
// @dart = 2.3
// ignore_for_file: camel_case_types,non_constant_identifier_names,library_prefixes,unused_import,unused_shown_name,return_of_invalid_type

import 'dart:async' as $async;

import 'dart:core' as $core;

import 'package:grpc/service_api.dart' as $grpc;
import 'inner_code.pb.dart' as $2;
import 'google/protobuf/empty.pb.dart' as $0;
export 'inner_code.pb.dart';

class InnerCodeClient extends $grpc.Client {
  static final _$verify = $grpc.ClientMethod<$2.VerifyRequest, $0.Empty>(
      '/appootb.captcha.InnerCode/Verify',
      ($2.VerifyRequest value) => value.writeToBuffer(),
      ($core.List<$core.int> value) => $0.Empty.fromBuffer(value));

  InnerCodeClient($grpc.ClientChannel channel, {$grpc.CallOptions options})
      : super(channel, options: options);

  $grpc.ResponseFuture<$0.Empty> verify($2.VerifyRequest request,
      {$grpc.CallOptions options}) {
    final call = $createCall(_$verify, $async.Stream.fromIterable([request]),
        options: options);
    return $grpc.ResponseFuture(call);
  }
}

abstract class InnerCodeServiceBase extends $grpc.Service {
  $core.String get $name => 'appootb.captcha.InnerCode';

  InnerCodeServiceBase() {
    $addMethod($grpc.ServiceMethod<$2.VerifyRequest, $0.Empty>(
        'Verify',
        verify_Pre,
        false,
        false,
        ($core.List<$core.int> value) => $2.VerifyRequest.fromBuffer(value),
        ($0.Empty value) => value.writeToBuffer()));
  }

  $async.Future<$0.Empty> verify_Pre(
      $grpc.ServiceCall call, $async.Future<$2.VerifyRequest> request) async {
    return verify(call, await request);
  }

  $async.Future<$0.Empty> verify(
      $grpc.ServiceCall call, $2.VerifyRequest request);
}
