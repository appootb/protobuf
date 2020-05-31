///
//  Generated code. Do not modify.
//  source: bind.proto
//
// @dart = 2.3
// ignore_for_file: camel_case_types,non_constant_identifier_names,library_prefixes,unused_import,unused_shown_name,return_of_invalid_type

import 'dart:async' as $async;

import 'dart:core' as $core;

import 'package:grpc/service_api.dart' as $grpc;
import 'bind.pb.dart' as $6;
import 'google/protobuf/empty.pb.dart' as $4;
export 'bind.pb.dart';

class BindClient extends $grpc.Client {
  static final _$apply = $grpc.ClientMethod<$6.BindRequest, $4.Empty>(
      '/appootb.account.Bind/Apply',
      ($6.BindRequest value) => value.writeToBuffer(),
      ($core.List<$core.int> value) => $4.Empty.fromBuffer(value));
  static final _$cancel = $grpc.ClientMethod<$6.BindRequest, $4.Empty>(
      '/appootb.account.Bind/Cancel',
      ($6.BindRequest value) => value.writeToBuffer(),
      ($core.List<$core.int> value) => $4.Empty.fromBuffer(value));
  static final _$gets = $grpc.ClientMethod<$4.Empty, $6.AccountBinds>(
      '/appootb.account.Bind/Gets',
      ($4.Empty value) => value.writeToBuffer(),
      ($core.List<$core.int> value) => $6.AccountBinds.fromBuffer(value));

  BindClient($grpc.ClientChannel channel, {$grpc.CallOptions options})
      : super(channel, options: options);

  $grpc.ResponseFuture<$4.Empty> apply($6.BindRequest request,
      {$grpc.CallOptions options}) {
    final call = $createCall(_$apply, $async.Stream.fromIterable([request]),
        options: options);
    return $grpc.ResponseFuture(call);
  }

  $grpc.ResponseFuture<$4.Empty> cancel($6.BindRequest request,
      {$grpc.CallOptions options}) {
    final call = $createCall(_$cancel, $async.Stream.fromIterable([request]),
        options: options);
    return $grpc.ResponseFuture(call);
  }

  $grpc.ResponseFuture<$6.AccountBinds> gets($4.Empty request,
      {$grpc.CallOptions options}) {
    final call = $createCall(_$gets, $async.Stream.fromIterable([request]),
        options: options);
    return $grpc.ResponseFuture(call);
  }
}

abstract class BindServiceBase extends $grpc.Service {
  $core.String get $name => 'appootb.account.Bind';

  BindServiceBase() {
    $addMethod($grpc.ServiceMethod<$6.BindRequest, $4.Empty>(
        'Apply',
        apply_Pre,
        false,
        false,
        ($core.List<$core.int> value) => $6.BindRequest.fromBuffer(value),
        ($4.Empty value) => value.writeToBuffer()));
    $addMethod($grpc.ServiceMethod<$6.BindRequest, $4.Empty>(
        'Cancel',
        cancel_Pre,
        false,
        false,
        ($core.List<$core.int> value) => $6.BindRequest.fromBuffer(value),
        ($4.Empty value) => value.writeToBuffer()));
    $addMethod($grpc.ServiceMethod<$4.Empty, $6.AccountBinds>(
        'Gets',
        gets_Pre,
        false,
        false,
        ($core.List<$core.int> value) => $4.Empty.fromBuffer(value),
        ($6.AccountBinds value) => value.writeToBuffer()));
  }

  $async.Future<$4.Empty> apply_Pre(
      $grpc.ServiceCall call, $async.Future<$6.BindRequest> request) async {
    return apply(call, await request);
  }

  $async.Future<$4.Empty> cancel_Pre(
      $grpc.ServiceCall call, $async.Future<$6.BindRequest> request) async {
    return cancel(call, await request);
  }

  $async.Future<$6.AccountBinds> gets_Pre(
      $grpc.ServiceCall call, $async.Future<$4.Empty> request) async {
    return gets(call, await request);
  }

  $async.Future<$4.Empty> apply($grpc.ServiceCall call, $6.BindRequest request);
  $async.Future<$4.Empty> cancel(
      $grpc.ServiceCall call, $6.BindRequest request);
  $async.Future<$6.AccountBinds> gets($grpc.ServiceCall call, $4.Empty request);
}
