///
//  Generated code. Do not modify.
//  source: profile.proto
//
// @dart = 2.3
// ignore_for_file: camel_case_types,non_constant_identifier_names,library_prefixes,unused_import,unused_shown_name,return_of_invalid_type

import 'dart:async' as $async;

import 'dart:core' as $core;

import 'package:grpc/service_api.dart' as $grpc;
import 'profile.pb.dart' as $8;
import 'google/protobuf/empty.pb.dart' as $4;
export 'profile.pb.dart';

class ProfileClient extends $grpc.Client {
  static final _$set = $grpc.ClientMethod<$8.Property, $4.Empty>(
      '/appootb.account.Profile/Set',
      ($8.Property value) => value.writeToBuffer(),
      ($core.List<$core.int> value) => $4.Empty.fromBuffer(value));
  static final _$get = $grpc.ClientMethod<$8.Property, $8.Property>(
      '/appootb.account.Profile/Get',
      ($8.Property value) => value.writeToBuffer(),
      ($core.List<$core.int> value) => $8.Property.fromBuffer(value));
  static final _$gets = $grpc.ClientMethod<$4.Empty, $8.Properties>(
      '/appootb.account.Profile/Gets',
      ($4.Empty value) => value.writeToBuffer(),
      ($core.List<$core.int> value) => $8.Properties.fromBuffer(value));

  ProfileClient($grpc.ClientChannel channel, {$grpc.CallOptions options})
      : super(channel, options: options);

  $grpc.ResponseFuture<$4.Empty> set($8.Property request,
      {$grpc.CallOptions options}) {
    final call = $createCall(_$set, $async.Stream.fromIterable([request]),
        options: options);
    return $grpc.ResponseFuture(call);
  }

  $grpc.ResponseFuture<$8.Property> get($8.Property request,
      {$grpc.CallOptions options}) {
    final call = $createCall(_$get, $async.Stream.fromIterable([request]),
        options: options);
    return $grpc.ResponseFuture(call);
  }

  $grpc.ResponseFuture<$8.Properties> gets($4.Empty request,
      {$grpc.CallOptions options}) {
    final call = $createCall(_$gets, $async.Stream.fromIterable([request]),
        options: options);
    return $grpc.ResponseFuture(call);
  }
}

abstract class ProfileServiceBase extends $grpc.Service {
  $core.String get $name => 'appootb.account.Profile';

  ProfileServiceBase() {
    $addMethod($grpc.ServiceMethod<$8.Property, $4.Empty>(
        'Set',
        set_Pre,
        false,
        false,
        ($core.List<$core.int> value) => $8.Property.fromBuffer(value),
        ($4.Empty value) => value.writeToBuffer()));
    $addMethod($grpc.ServiceMethod<$8.Property, $8.Property>(
        'Get',
        get_Pre,
        false,
        false,
        ($core.List<$core.int> value) => $8.Property.fromBuffer(value),
        ($8.Property value) => value.writeToBuffer()));
    $addMethod($grpc.ServiceMethod<$4.Empty, $8.Properties>(
        'Gets',
        gets_Pre,
        false,
        false,
        ($core.List<$core.int> value) => $4.Empty.fromBuffer(value),
        ($8.Properties value) => value.writeToBuffer()));
  }

  $async.Future<$4.Empty> set_Pre(
      $grpc.ServiceCall call, $async.Future<$8.Property> request) async {
    return set(call, await request);
  }

  $async.Future<$8.Property> get_Pre(
      $grpc.ServiceCall call, $async.Future<$8.Property> request) async {
    return get(call, await request);
  }

  $async.Future<$8.Properties> gets_Pre(
      $grpc.ServiceCall call, $async.Future<$4.Empty> request) async {
    return gets(call, await request);
  }

  $async.Future<$4.Empty> set($grpc.ServiceCall call, $8.Property request);
  $async.Future<$8.Property> get($grpc.ServiceCall call, $8.Property request);
  $async.Future<$8.Properties> gets($grpc.ServiceCall call, $4.Empty request);
}
