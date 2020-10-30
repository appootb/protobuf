///
//  Generated code. Do not modify.
//  source: inner_hash.proto
//
// @dart = 2.3
// ignore_for_file: camel_case_types,non_constant_identifier_names,library_prefixes,unused_import,unused_shown_name,return_of_invalid_type

import 'dart:async' as $async;

import 'dart:core' as $core;

import 'package:grpc/service_api.dart' as $grpc;
import 'inner_hash.pb.dart' as $0;
import 'google/protobuf/empty.pb.dart' as $1;
export 'inner_hash.pb.dart';

class InnerHashClient extends $grpc.Client {
  static final _$increase = $grpc.ClientMethod<$0.HashField, $0.HashField>(
      '/appootb.counter.InnerHash/Increase',
      ($0.HashField value) => value.writeToBuffer(),
      ($core.List<$core.int> value) => $0.HashField.fromBuffer(value));
  static final _$multiIncrease =
      $grpc.ClientMethod<$0.HashFields, $0.HashFields>(
          '/appootb.counter.InnerHash/MultiIncrease',
          ($0.HashFields value) => value.writeToBuffer(),
          ($core.List<$core.int> value) => $0.HashFields.fromBuffer(value));
  static final _$get = $grpc.ClientMethod<$0.HashField, $0.HashField>(
      '/appootb.counter.InnerHash/Get',
      ($0.HashField value) => value.writeToBuffer(),
      ($core.List<$core.int> value) => $0.HashField.fromBuffer(value));
  static final _$gets = $grpc.ClientMethod<$0.HashFields, $0.HashFields>(
      '/appootb.counter.InnerHash/Gets',
      ($0.HashFields value) => value.writeToBuffer(),
      ($core.List<$core.int> value) => $0.HashFields.fromBuffer(value));
  static final _$set = $grpc.ClientMethod<$0.HashField, $1.Empty>(
      '/appootb.counter.InnerHash/Set',
      ($0.HashField value) => value.writeToBuffer(),
      ($core.List<$core.int> value) => $1.Empty.fromBuffer(value));
  static final _$sets = $grpc.ClientMethod<$0.HashFields, $1.Empty>(
      '/appootb.counter.InnerHash/Sets',
      ($0.HashFields value) => value.writeToBuffer(),
      ($core.List<$core.int> value) => $1.Empty.fromBuffer(value));
  static final _$multiGets = $grpc.ClientMethod<$0.HashKeys, $0.HashValues>(
      '/appootb.counter.InnerHash/MultiGets',
      ($0.HashKeys value) => value.writeToBuffer(),
      ($core.List<$core.int> value) => $0.HashValues.fromBuffer(value));

  InnerHashClient($grpc.ClientChannel channel, {$grpc.CallOptions options})
      : super(channel, options: options);

  $grpc.ResponseFuture<$0.HashField> increase($0.HashField request,
      {$grpc.CallOptions options}) {
    final call = $createCall(_$increase, $async.Stream.fromIterable([request]),
        options: options);
    return $grpc.ResponseFuture(call);
  }

  $grpc.ResponseFuture<$0.HashFields> multiIncrease($0.HashFields request,
      {$grpc.CallOptions options}) {
    final call = $createCall(
        _$multiIncrease, $async.Stream.fromIterable([request]),
        options: options);
    return $grpc.ResponseFuture(call);
  }

  $grpc.ResponseFuture<$0.HashField> get($0.HashField request,
      {$grpc.CallOptions options}) {
    final call = $createCall(_$get, $async.Stream.fromIterable([request]),
        options: options);
    return $grpc.ResponseFuture(call);
  }

  $grpc.ResponseFuture<$0.HashFields> gets($0.HashFields request,
      {$grpc.CallOptions options}) {
    final call = $createCall(_$gets, $async.Stream.fromIterable([request]),
        options: options);
    return $grpc.ResponseFuture(call);
  }

  $grpc.ResponseFuture<$1.Empty> set($0.HashField request,
      {$grpc.CallOptions options}) {
    final call = $createCall(_$set, $async.Stream.fromIterable([request]),
        options: options);
    return $grpc.ResponseFuture(call);
  }

  $grpc.ResponseFuture<$1.Empty> sets($0.HashFields request,
      {$grpc.CallOptions options}) {
    final call = $createCall(_$sets, $async.Stream.fromIterable([request]),
        options: options);
    return $grpc.ResponseFuture(call);
  }

  $grpc.ResponseFuture<$0.HashValues> multiGets($0.HashKeys request,
      {$grpc.CallOptions options}) {
    final call = $createCall(_$multiGets, $async.Stream.fromIterable([request]),
        options: options);
    return $grpc.ResponseFuture(call);
  }
}

abstract class InnerHashServiceBase extends $grpc.Service {
  $core.String get $name => 'appootb.counter.InnerHash';

  InnerHashServiceBase() {
    $addMethod($grpc.ServiceMethod<$0.HashField, $0.HashField>(
        'Increase',
        increase_Pre,
        false,
        false,
        ($core.List<$core.int> value) => $0.HashField.fromBuffer(value),
        ($0.HashField value) => value.writeToBuffer()));
    $addMethod($grpc.ServiceMethod<$0.HashFields, $0.HashFields>(
        'MultiIncrease',
        multiIncrease_Pre,
        false,
        false,
        ($core.List<$core.int> value) => $0.HashFields.fromBuffer(value),
        ($0.HashFields value) => value.writeToBuffer()));
    $addMethod($grpc.ServiceMethod<$0.HashField, $0.HashField>(
        'Get',
        get_Pre,
        false,
        false,
        ($core.List<$core.int> value) => $0.HashField.fromBuffer(value),
        ($0.HashField value) => value.writeToBuffer()));
    $addMethod($grpc.ServiceMethod<$0.HashFields, $0.HashFields>(
        'Gets',
        gets_Pre,
        false,
        false,
        ($core.List<$core.int> value) => $0.HashFields.fromBuffer(value),
        ($0.HashFields value) => value.writeToBuffer()));
    $addMethod($grpc.ServiceMethod<$0.HashField, $1.Empty>(
        'Set',
        set_Pre,
        false,
        false,
        ($core.List<$core.int> value) => $0.HashField.fromBuffer(value),
        ($1.Empty value) => value.writeToBuffer()));
    $addMethod($grpc.ServiceMethod<$0.HashFields, $1.Empty>(
        'Sets',
        sets_Pre,
        false,
        false,
        ($core.List<$core.int> value) => $0.HashFields.fromBuffer(value),
        ($1.Empty value) => value.writeToBuffer()));
    $addMethod($grpc.ServiceMethod<$0.HashKeys, $0.HashValues>(
        'MultiGets',
        multiGets_Pre,
        false,
        false,
        ($core.List<$core.int> value) => $0.HashKeys.fromBuffer(value),
        ($0.HashValues value) => value.writeToBuffer()));
  }

  $async.Future<$0.HashField> increase_Pre(
      $grpc.ServiceCall call, $async.Future<$0.HashField> request) async {
    return increase(call, await request);
  }

  $async.Future<$0.HashFields> multiIncrease_Pre(
      $grpc.ServiceCall call, $async.Future<$0.HashFields> request) async {
    return multiIncrease(call, await request);
  }

  $async.Future<$0.HashField> get_Pre(
      $grpc.ServiceCall call, $async.Future<$0.HashField> request) async {
    return get(call, await request);
  }

  $async.Future<$0.HashFields> gets_Pre(
      $grpc.ServiceCall call, $async.Future<$0.HashFields> request) async {
    return gets(call, await request);
  }

  $async.Future<$1.Empty> set_Pre(
      $grpc.ServiceCall call, $async.Future<$0.HashField> request) async {
    return set(call, await request);
  }

  $async.Future<$1.Empty> sets_Pre(
      $grpc.ServiceCall call, $async.Future<$0.HashFields> request) async {
    return sets(call, await request);
  }

  $async.Future<$0.HashValues> multiGets_Pre(
      $grpc.ServiceCall call, $async.Future<$0.HashKeys> request) async {
    return multiGets(call, await request);
  }

  $async.Future<$0.HashField> increase(
      $grpc.ServiceCall call, $0.HashField request);
  $async.Future<$0.HashFields> multiIncrease(
      $grpc.ServiceCall call, $0.HashFields request);
  $async.Future<$0.HashField> get($grpc.ServiceCall call, $0.HashField request);
  $async.Future<$0.HashFields> gets(
      $grpc.ServiceCall call, $0.HashFields request);
  $async.Future<$1.Empty> set($grpc.ServiceCall call, $0.HashField request);
  $async.Future<$1.Empty> sets($grpc.ServiceCall call, $0.HashFields request);
  $async.Future<$0.HashValues> multiGets(
      $grpc.ServiceCall call, $0.HashKeys request);
}
