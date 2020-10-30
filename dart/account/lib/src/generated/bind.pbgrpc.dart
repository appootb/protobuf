///
//  Generated code. Do not modify.
//  source: bind.proto
//
// @dart = 2.3
// ignore_for_file: camel_case_types,non_constant_identifier_names,library_prefixes,unused_import,unused_shown_name,return_of_invalid_type

import 'dart:async' as $async;

import 'dart:core' as $core;

import 'package:grpc/service_api.dart' as $grpc;
import 'bind.pb.dart' as $5;
import 'google/protobuf/empty.pb.dart' as $3;
export 'bind.pb.dart';

class BindClient extends $grpc.Client {
  static final _$apply = $grpc.ClientMethod<$5.BindRequest, $3.Empty>(
      '/appootb.account.Bind/Apply',
      ($5.BindRequest value) => value.writeToBuffer(),
      ($core.List<$core.int> value) => $3.Empty.fromBuffer(value));
  static final _$cancel = $grpc.ClientMethod<$5.BindRequest, $3.Empty>(
      '/appootb.account.Bind/Cancel',
      ($5.BindRequest value) => value.writeToBuffer(),
      ($core.List<$core.int> value) => $3.Empty.fromBuffer(value));
  static final _$gets = $grpc.ClientMethod<$3.Empty, $5.Bindings>(
      '/appootb.account.Bind/Gets',
      ($3.Empty value) => value.writeToBuffer(),
      ($core.List<$core.int> value) => $5.Bindings.fromBuffer(value));

  BindClient($grpc.ClientChannel channel, {$grpc.CallOptions options})
      : super(channel, options: options);

  $grpc.ResponseFuture<$3.Empty> apply($5.BindRequest request,
      {$grpc.CallOptions options}) {
    final call = $createCall(_$apply, $async.Stream.fromIterable([request]),
        options: options);
    return $grpc.ResponseFuture(call);
  }

  $grpc.ResponseFuture<$3.Empty> cancel($5.BindRequest request,
      {$grpc.CallOptions options}) {
    final call = $createCall(_$cancel, $async.Stream.fromIterable([request]),
        options: options);
    return $grpc.ResponseFuture(call);
  }

  $grpc.ResponseFuture<$5.Bindings> gets($3.Empty request,
      {$grpc.CallOptions options}) {
    final call = $createCall(_$gets, $async.Stream.fromIterable([request]),
        options: options);
    return $grpc.ResponseFuture(call);
  }
}

abstract class BindServiceBase extends $grpc.Service {
  $core.String get $name => 'appootb.account.Bind';

  BindServiceBase() {
    $addMethod($grpc.ServiceMethod<$5.BindRequest, $3.Empty>(
        'Apply',
        apply_Pre,
        false,
        false,
        ($core.List<$core.int> value) => $5.BindRequest.fromBuffer(value),
        ($3.Empty value) => value.writeToBuffer()));
    $addMethod($grpc.ServiceMethod<$5.BindRequest, $3.Empty>(
        'Cancel',
        cancel_Pre,
        false,
        false,
        ($core.List<$core.int> value) => $5.BindRequest.fromBuffer(value),
        ($3.Empty value) => value.writeToBuffer()));
    $addMethod($grpc.ServiceMethod<$3.Empty, $5.Bindings>(
        'Gets',
        gets_Pre,
        false,
        false,
        ($core.List<$core.int> value) => $3.Empty.fromBuffer(value),
        ($5.Bindings value) => value.writeToBuffer()));
  }

  $async.Future<$3.Empty> apply_Pre(
      $grpc.ServiceCall call, $async.Future<$5.BindRequest> request) async {
    return apply(call, await request);
  }

  $async.Future<$3.Empty> cancel_Pre(
      $grpc.ServiceCall call, $async.Future<$5.BindRequest> request) async {
    return cancel(call, await request);
  }

  $async.Future<$5.Bindings> gets_Pre(
      $grpc.ServiceCall call, $async.Future<$3.Empty> request) async {
    return gets(call, await request);
  }

  $async.Future<$3.Empty> apply($grpc.ServiceCall call, $5.BindRequest request);
  $async.Future<$3.Empty> cancel(
      $grpc.ServiceCall call, $5.BindRequest request);
  $async.Future<$5.Bindings> gets($grpc.ServiceCall call, $3.Empty request);
}
