///
//  Generated code. Do not modify.
//  source: code.proto
//
// @dart = 2.3
// ignore_for_file: camel_case_types,non_constant_identifier_names,library_prefixes,unused_import,unused_shown_name,return_of_invalid_type

import 'dart:async' as $async;

import 'dart:core' as $core;

import 'package:grpc/service_api.dart' as $grpc;
import 'google/protobuf/empty.pb.dart' as $0;
import 'code.pb.dart' as $1;
export 'code.pb.dart';

class CodeClient extends $grpc.Client {
  static final _$getRegions = $grpc.ClientMethod<$0.Empty, $1.Regions>(
      '/appootb.captcha.Code/GetRegions',
      ($0.Empty value) => value.writeToBuffer(),
      ($core.List<$core.int> value) => $1.Regions.fromBuffer(value));
  static final _$apply = $grpc.ClientMethod<$1.CodeRequest, $0.Empty>(
      '/appootb.captcha.Code/Apply',
      ($1.CodeRequest value) => value.writeToBuffer(),
      ($core.List<$core.int> value) => $0.Empty.fromBuffer(value));

  CodeClient($grpc.ClientChannel channel, {$grpc.CallOptions options})
      : super(channel, options: options);

  $grpc.ResponseFuture<$1.Regions> getRegions($0.Empty request,
      {$grpc.CallOptions options}) {
    final call = $createCall(
        _$getRegions, $async.Stream.fromIterable([request]),
        options: options);
    return $grpc.ResponseFuture(call);
  }

  $grpc.ResponseFuture<$0.Empty> apply($1.CodeRequest request,
      {$grpc.CallOptions options}) {
    final call = $createCall(_$apply, $async.Stream.fromIterable([request]),
        options: options);
    return $grpc.ResponseFuture(call);
  }
}

abstract class CodeServiceBase extends $grpc.Service {
  $core.String get $name => 'appootb.captcha.Code';

  CodeServiceBase() {
    $addMethod($grpc.ServiceMethod<$0.Empty, $1.Regions>(
        'GetRegions',
        getRegions_Pre,
        false,
        false,
        ($core.List<$core.int> value) => $0.Empty.fromBuffer(value),
        ($1.Regions value) => value.writeToBuffer()));
    $addMethod($grpc.ServiceMethod<$1.CodeRequest, $0.Empty>(
        'Apply',
        apply_Pre,
        false,
        false,
        ($core.List<$core.int> value) => $1.CodeRequest.fromBuffer(value),
        ($0.Empty value) => value.writeToBuffer()));
  }

  $async.Future<$1.Regions> getRegions_Pre(
      $grpc.ServiceCall call, $async.Future<$0.Empty> request) async {
    return getRegions(call, await request);
  }

  $async.Future<$0.Empty> apply_Pre(
      $grpc.ServiceCall call, $async.Future<$1.CodeRequest> request) async {
    return apply(call, await request);
  }

  $async.Future<$1.Regions> getRegions(
      $grpc.ServiceCall call, $0.Empty request);
  $async.Future<$0.Empty> apply($grpc.ServiceCall call, $1.CodeRequest request);
}
