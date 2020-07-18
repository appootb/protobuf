///
//  Generated code. Do not modify.
//  source: auth.proto
//
// @dart = 2.3
// ignore_for_file: camel_case_types,non_constant_identifier_names,library_prefixes,unused_import,unused_shown_name,return_of_invalid_type

import 'dart:async' as $async;

import 'dart:core' as $core;

import 'package:grpc/service_api.dart' as $grpc;
import 'auth.pb.dart' as $3;
import 'include.pb.dart' as $1;
import 'google/protobuf/empty.pb.dart' as $4;
export 'auth.pb.dart';

class AuthClient extends $grpc.Client {
  static final _$login = $grpc.ClientMethod<$3.LoginRequest, $1.Info>(
      '/appootb.account.Auth/Login',
      ($3.LoginRequest value) => value.writeToBuffer(),
      ($core.List<$core.int> value) => $1.Info.fromBuffer(value));
  static final _$oAuth = $grpc.ClientMethod<$3.OAuthRequest, $1.Info>(
      '/appootb.account.Auth/OAuth',
      ($3.OAuthRequest value) => value.writeToBuffer(),
      ($core.List<$core.int> value) => $1.Info.fromBuffer(value));
  static final _$refresh = $grpc.ClientMethod<$4.Empty, $1.Info>(
      '/appootb.account.Auth/Refresh',
      ($4.Empty value) => value.writeToBuffer(),
      ($core.List<$core.int> value) => $1.Info.fromBuffer(value));

  AuthClient($grpc.ClientChannel channel, {$grpc.CallOptions options})
      : super(channel, options: options);

  $grpc.ResponseFuture<$1.Info> login($3.LoginRequest request,
      {$grpc.CallOptions options}) {
    final call = $createCall(_$login, $async.Stream.fromIterable([request]),
        options: options);
    return $grpc.ResponseFuture(call);
  }

  $grpc.ResponseFuture<$1.Info> oAuth($3.OAuthRequest request,
      {$grpc.CallOptions options}) {
    final call = $createCall(_$oAuth, $async.Stream.fromIterable([request]),
        options: options);
    return $grpc.ResponseFuture(call);
  }

  $grpc.ResponseFuture<$1.Info> refresh($4.Empty request,
      {$grpc.CallOptions options}) {
    final call = $createCall(_$refresh, $async.Stream.fromIterable([request]),
        options: options);
    return $grpc.ResponseFuture(call);
  }
}

abstract class AuthServiceBase extends $grpc.Service {
  $core.String get $name => 'appootb.account.Auth';

  AuthServiceBase() {
    $addMethod($grpc.ServiceMethod<$3.LoginRequest, $1.Info>(
        'Login',
        login_Pre,
        false,
        false,
        ($core.List<$core.int> value) => $3.LoginRequest.fromBuffer(value),
        ($1.Info value) => value.writeToBuffer()));
    $addMethod($grpc.ServiceMethod<$3.OAuthRequest, $1.Info>(
        'OAuth',
        oAuth_Pre,
        false,
        false,
        ($core.List<$core.int> value) => $3.OAuthRequest.fromBuffer(value),
        ($1.Info value) => value.writeToBuffer()));
    $addMethod($grpc.ServiceMethod<$4.Empty, $1.Info>(
        'Refresh',
        refresh_Pre,
        false,
        false,
        ($core.List<$core.int> value) => $4.Empty.fromBuffer(value),
        ($1.Info value) => value.writeToBuffer()));
  }

  $async.Future<$1.Info> login_Pre(
      $grpc.ServiceCall call, $async.Future<$3.LoginRequest> request) async {
    return login(call, await request);
  }

  $async.Future<$1.Info> oAuth_Pre(
      $grpc.ServiceCall call, $async.Future<$3.OAuthRequest> request) async {
    return oAuth(call, await request);
  }

  $async.Future<$1.Info> refresh_Pre(
      $grpc.ServiceCall call, $async.Future<$4.Empty> request) async {
    return refresh(call, await request);
  }

  $async.Future<$1.Info> login($grpc.ServiceCall call, $3.LoginRequest request);
  $async.Future<$1.Info> oAuth($grpc.ServiceCall call, $3.OAuthRequest request);
  $async.Future<$1.Info> refresh($grpc.ServiceCall call, $4.Empty request);
}
