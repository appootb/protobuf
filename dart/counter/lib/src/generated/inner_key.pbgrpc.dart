///
//  Generated code. Do not modify.
//  source: inner_key.proto
//
// @dart = 2.3
// ignore_for_file: camel_case_types,non_constant_identifier_names,library_prefixes,unused_import,unused_shown_name,return_of_invalid_type

import 'dart:async' as $async;

import 'dart:core' as $core;

import 'package:grpc/service_api.dart' as $grpc;
import 'inner_key.pb.dart' as $2;
import 'google/protobuf/empty.pb.dart' as $1;
export 'inner_key.pb.dart';

class InnerKeyClient extends $grpc.Client {
  static final _$increase = $grpc.ClientMethod<$2.Key, $2.Key>(
      '/appootb.counter.InnerKey/Increase',
      ($2.Key value) => value.writeToBuffer(),
      ($core.List<$core.int> value) => $2.Key.fromBuffer(value));
  static final _$multiIncrease = $grpc.ClientMethod<$2.Keys, $2.Keys>(
      '/appootb.counter.InnerKey/MultiIncrease',
      ($2.Keys value) => value.writeToBuffer(),
      ($core.List<$core.int> value) => $2.Keys.fromBuffer(value));
  static final _$get = $grpc.ClientMethod<$2.Key, $2.Key>(
      '/appootb.counter.InnerKey/Get',
      ($2.Key value) => value.writeToBuffer(),
      ($core.List<$core.int> value) => $2.Key.fromBuffer(value));
  static final _$gets = $grpc.ClientMethod<$2.Keys, $2.Keys>(
      '/appootb.counter.InnerKey/Gets',
      ($2.Keys value) => value.writeToBuffer(),
      ($core.List<$core.int> value) => $2.Keys.fromBuffer(value));
  static final _$set = $grpc.ClientMethod<$2.Key, $1.Empty>(
      '/appootb.counter.InnerKey/Set',
      ($2.Key value) => value.writeToBuffer(),
      ($core.List<$core.int> value) => $1.Empty.fromBuffer(value));
  static final _$sets = $grpc.ClientMethod<$2.Keys, $1.Empty>(
      '/appootb.counter.InnerKey/Sets',
      ($2.Keys value) => value.writeToBuffer(),
      ($core.List<$core.int> value) => $1.Empty.fromBuffer(value));

  InnerKeyClient($grpc.ClientChannel channel, {$grpc.CallOptions options})
      : super(channel, options: options);

  $grpc.ResponseFuture<$2.Key> increase($2.Key request,
      {$grpc.CallOptions options}) {
    final call = $createCall(_$increase, $async.Stream.fromIterable([request]),
        options: options);
    return $grpc.ResponseFuture(call);
  }

  $grpc.ResponseFuture<$2.Keys> multiIncrease($2.Keys request,
      {$grpc.CallOptions options}) {
    final call = $createCall(
        _$multiIncrease, $async.Stream.fromIterable([request]),
        options: options);
    return $grpc.ResponseFuture(call);
  }

  $grpc.ResponseFuture<$2.Key> get($2.Key request,
      {$grpc.CallOptions options}) {
    final call = $createCall(_$get, $async.Stream.fromIterable([request]),
        options: options);
    return $grpc.ResponseFuture(call);
  }

  $grpc.ResponseFuture<$2.Keys> gets($2.Keys request,
      {$grpc.CallOptions options}) {
    final call = $createCall(_$gets, $async.Stream.fromIterable([request]),
        options: options);
    return $grpc.ResponseFuture(call);
  }

  $grpc.ResponseFuture<$1.Empty> set($2.Key request,
      {$grpc.CallOptions options}) {
    final call = $createCall(_$set, $async.Stream.fromIterable([request]),
        options: options);
    return $grpc.ResponseFuture(call);
  }

  $grpc.ResponseFuture<$1.Empty> sets($2.Keys request,
      {$grpc.CallOptions options}) {
    final call = $createCall(_$sets, $async.Stream.fromIterable([request]),
        options: options);
    return $grpc.ResponseFuture(call);
  }
}

abstract class InnerKeyServiceBase extends $grpc.Service {
  $core.String get $name => 'appootb.counter.InnerKey';

  InnerKeyServiceBase() {
    $addMethod($grpc.ServiceMethod<$2.Key, $2.Key>(
        'Increase',
        increase_Pre,
        false,
        false,
        ($core.List<$core.int> value) => $2.Key.fromBuffer(value),
        ($2.Key value) => value.writeToBuffer()));
    $addMethod($grpc.ServiceMethod<$2.Keys, $2.Keys>(
        'MultiIncrease',
        multiIncrease_Pre,
        false,
        false,
        ($core.List<$core.int> value) => $2.Keys.fromBuffer(value),
        ($2.Keys value) => value.writeToBuffer()));
    $addMethod($grpc.ServiceMethod<$2.Key, $2.Key>(
        'Get',
        get_Pre,
        false,
        false,
        ($core.List<$core.int> value) => $2.Key.fromBuffer(value),
        ($2.Key value) => value.writeToBuffer()));
    $addMethod($grpc.ServiceMethod<$2.Keys, $2.Keys>(
        'Gets',
        gets_Pre,
        false,
        false,
        ($core.List<$core.int> value) => $2.Keys.fromBuffer(value),
        ($2.Keys value) => value.writeToBuffer()));
    $addMethod($grpc.ServiceMethod<$2.Key, $1.Empty>(
        'Set',
        set_Pre,
        false,
        false,
        ($core.List<$core.int> value) => $2.Key.fromBuffer(value),
        ($1.Empty value) => value.writeToBuffer()));
    $addMethod($grpc.ServiceMethod<$2.Keys, $1.Empty>(
        'Sets',
        sets_Pre,
        false,
        false,
        ($core.List<$core.int> value) => $2.Keys.fromBuffer(value),
        ($1.Empty value) => value.writeToBuffer()));
  }

  $async.Future<$2.Key> increase_Pre(
      $grpc.ServiceCall call, $async.Future<$2.Key> request) async {
    return increase(call, await request);
  }

  $async.Future<$2.Keys> multiIncrease_Pre(
      $grpc.ServiceCall call, $async.Future<$2.Keys> request) async {
    return multiIncrease(call, await request);
  }

  $async.Future<$2.Key> get_Pre(
      $grpc.ServiceCall call, $async.Future<$2.Key> request) async {
    return get(call, await request);
  }

  $async.Future<$2.Keys> gets_Pre(
      $grpc.ServiceCall call, $async.Future<$2.Keys> request) async {
    return gets(call, await request);
  }

  $async.Future<$1.Empty> set_Pre(
      $grpc.ServiceCall call, $async.Future<$2.Key> request) async {
    return set(call, await request);
  }

  $async.Future<$1.Empty> sets_Pre(
      $grpc.ServiceCall call, $async.Future<$2.Keys> request) async {
    return sets(call, await request);
  }

  $async.Future<$2.Key> increase($grpc.ServiceCall call, $2.Key request);
  $async.Future<$2.Keys> multiIncrease($grpc.ServiceCall call, $2.Keys request);
  $async.Future<$2.Key> get($grpc.ServiceCall call, $2.Key request);
  $async.Future<$2.Keys> gets($grpc.ServiceCall call, $2.Keys request);
  $async.Future<$1.Empty> set($grpc.ServiceCall call, $2.Key request);
  $async.Future<$1.Empty> sets($grpc.ServiceCall call, $2.Keys request);
}
