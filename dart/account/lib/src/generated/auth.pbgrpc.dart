///
//  Generated code. Do not modify.
//  source: auth.proto
//
// @dart = 2.3
// ignore_for_file: camel_case_types,non_constant_identifier_names,library_prefixes,unused_import,unused_shown_name,return_of_invalid_type

import 'dart:async' as $async;

import 'dart:core' as $core;

import 'package:grpc/service_api.dart' as $grpc;
import 'appootb/captcha/common.pb.dart' as $3;
import 'google/protobuf/empty.pb.dart' as $4;
import 'auth.pb.dart' as $5;
import 'common.pb.dart' as $1;
export 'auth.pb.dart';

class AuthClient extends $grpc.Client {
  static final _$getCode = $grpc.ClientMethod<$3.CodeRequest, $4.Empty>(
      '/appootb.account.Auth/GetCode',
      ($3.CodeRequest value) => value.writeToBuffer(),
      ($core.List<$core.int> value) => $4.Empty.fromBuffer(value));
  static final _$login = $grpc.ClientMethod<$5.LoginRequest, $1.AccountInfo>(
      '/appootb.account.Auth/Login',
      ($5.LoginRequest value) => value.writeToBuffer(),
      ($core.List<$core.int> value) => $1.AccountInfo.fromBuffer(value));
  static final _$oAuth = $grpc.ClientMethod<$5.OAuthRequest, $1.AccountInfo>(
      '/appootb.account.Auth/OAuth',
      ($5.OAuthRequest value) => value.writeToBuffer(),
      ($core.List<$core.int> value) => $1.AccountInfo.fromBuffer(value));
  static final _$getRegions = $grpc.ClientMethod<$4.Empty, $5.Regions>(
      '/appootb.account.Auth/GetRegions',
      ($4.Empty value) => value.writeToBuffer(),
      ($core.List<$core.int> value) => $5.Regions.fromBuffer(value));
  static final _$refresh = $grpc.ClientMethod<$4.Empty, $1.AccountInfo>(
      '/appootb.account.Auth/Refresh',
      ($4.Empty value) => value.writeToBuffer(),
      ($core.List<$core.int> value) => $1.AccountInfo.fromBuffer(value));

  AuthClient($grpc.ClientChannel channel, {$grpc.CallOptions options})
      : super(channel, options: options);

  $grpc.ResponseFuture<$4.Empty> getCode($3.CodeRequest request,
      {$grpc.CallOptions options}) {
    final call = $createCall(_$getCode, $async.Stream.fromIterable([request]),
        options: options);
    return $grpc.ResponseFuture(call);
  }

  $grpc.ResponseFuture<$1.AccountInfo> login($5.LoginRequest request,
      {$grpc.CallOptions options}) {
    final call = $createCall(_$login, $async.Stream.fromIterable([request]),
        options: options);
    return $grpc.ResponseFuture(call);
  }

  $grpc.ResponseFuture<$1.AccountInfo> oAuth($5.OAuthRequest request,
      {$grpc.CallOptions options}) {
    final call = $createCall(_$oAuth, $async.Stream.fromIterable([request]),
        options: options);
    return $grpc.ResponseFuture(call);
  }

  $grpc.ResponseFuture<$5.Regions> getRegions($4.Empty request,
      {$grpc.CallOptions options}) {
    final call = $createCall(
        _$getRegions, $async.Stream.fromIterable([request]),
        options: options);
    return $grpc.ResponseFuture(call);
  }

  $grpc.ResponseFuture<$1.AccountInfo> refresh($4.Empty request,
      {$grpc.CallOptions options}) {
    final call = $createCall(_$refresh, $async.Stream.fromIterable([request]),
        options: options);
    return $grpc.ResponseFuture(call);
  }
}

abstract class AuthServiceBase extends $grpc.Service {
  $core.String get $name => 'appootb.account.Auth';

  AuthServiceBase() {
    $addMethod($grpc.ServiceMethod<$3.CodeRequest, $4.Empty>(
        'GetCode',
        getCode_Pre,
        false,
        false,
        ($core.List<$core.int> value) => $3.CodeRequest.fromBuffer(value),
        ($4.Empty value) => value.writeToBuffer()));
    $addMethod($grpc.ServiceMethod<$5.LoginRequest, $1.AccountInfo>(
        'Login',
        login_Pre,
        false,
        false,
        ($core.List<$core.int> value) => $5.LoginRequest.fromBuffer(value),
        ($1.AccountInfo value) => value.writeToBuffer()));
    $addMethod($grpc.ServiceMethod<$5.OAuthRequest, $1.AccountInfo>(
        'OAuth',
        oAuth_Pre,
        false,
        false,
        ($core.List<$core.int> value) => $5.OAuthRequest.fromBuffer(value),
        ($1.AccountInfo value) => value.writeToBuffer()));
    $addMethod($grpc.ServiceMethod<$4.Empty, $5.Regions>(
        'GetRegions',
        getRegions_Pre,
        false,
        false,
        ($core.List<$core.int> value) => $4.Empty.fromBuffer(value),
        ($5.Regions value) => value.writeToBuffer()));
    $addMethod($grpc.ServiceMethod<$4.Empty, $1.AccountInfo>(
        'Refresh',
        refresh_Pre,
        false,
        false,
        ($core.List<$core.int> value) => $4.Empty.fromBuffer(value),
        ($1.AccountInfo value) => value.writeToBuffer()));
  }

  $async.Future<$4.Empty> getCode_Pre(
      $grpc.ServiceCall call, $async.Future<$3.CodeRequest> request) async {
    return getCode(call, await request);
  }

  $async.Future<$1.AccountInfo> login_Pre(
      $grpc.ServiceCall call, $async.Future<$5.LoginRequest> request) async {
    return login(call, await request);
  }

  $async.Future<$1.AccountInfo> oAuth_Pre(
      $grpc.ServiceCall call, $async.Future<$5.OAuthRequest> request) async {
    return oAuth(call, await request);
  }

  $async.Future<$5.Regions> getRegions_Pre(
      $grpc.ServiceCall call, $async.Future<$4.Empty> request) async {
    return getRegions(call, await request);
  }

  $async.Future<$1.AccountInfo> refresh_Pre(
      $grpc.ServiceCall call, $async.Future<$4.Empty> request) async {
    return refresh(call, await request);
  }

  $async.Future<$4.Empty> getCode(
      $grpc.ServiceCall call, $3.CodeRequest request);
  $async.Future<$1.AccountInfo> login(
      $grpc.ServiceCall call, $5.LoginRequest request);
  $async.Future<$1.AccountInfo> oAuth(
      $grpc.ServiceCall call, $5.OAuthRequest request);
  $async.Future<$5.Regions> getRegions(
      $grpc.ServiceCall call, $4.Empty request);
  $async.Future<$1.AccountInfo> refresh(
      $grpc.ServiceCall call, $4.Empty request);
}
