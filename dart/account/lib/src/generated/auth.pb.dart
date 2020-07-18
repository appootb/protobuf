///
//  Generated code. Do not modify.
//  source: auth.proto
//
// @dart = 2.3
// ignore_for_file: camel_case_types,non_constant_identifier_names,library_prefixes,unused_import,unused_shown_name,return_of_invalid_type

import 'dart:core' as $core;

import 'package:protobuf/protobuf.dart' as $pb;

import 'include.pbenum.dart' as $1;

enum LoginRequest_Secure {
  code, 
  password, 
  notSet
}

class LoginRequest extends $pb.GeneratedMessage {
  static const $core.Map<$core.int, LoginRequest_Secure> _LoginRequest_SecureByTag = {
    2 : LoginRequest_Secure.code,
    3 : LoginRequest_Secure.password,
    0 : LoginRequest_Secure.notSet
  };
  static final $pb.BuilderInfo _i = $pb.BuilderInfo('LoginRequest', package: const $pb.PackageName('appootb.account'), createEmptyInstance: create)
    ..oo(0, [2, 3])
    ..aQS(1, 'relateId')
    ..aOS(2, 'code')
    ..aOS(3, 'password')
  ;

  LoginRequest._() : super();
  factory LoginRequest() => create();
  factory LoginRequest.fromBuffer($core.List<$core.int> i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromBuffer(i, r);
  factory LoginRequest.fromJson($core.String i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromJson(i, r);
  LoginRequest clone() => LoginRequest()..mergeFromMessage(this);
  LoginRequest copyWith(void Function(LoginRequest) updates) => super.copyWith((message) => updates(message as LoginRequest));
  $pb.BuilderInfo get info_ => _i;
  @$core.pragma('dart2js:noInline')
  static LoginRequest create() => LoginRequest._();
  LoginRequest createEmptyInstance() => create();
  static $pb.PbList<LoginRequest> createRepeated() => $pb.PbList<LoginRequest>();
  @$core.pragma('dart2js:noInline')
  static LoginRequest getDefault() => _defaultInstance ??= $pb.GeneratedMessage.$_defaultFor<LoginRequest>(create);
  static LoginRequest _defaultInstance;

  LoginRequest_Secure whichSecure() => _LoginRequest_SecureByTag[$_whichOneof(0)];
  void clearSecure() => clearField($_whichOneof(0));

  @$pb.TagNumber(1)
  $core.String get relateId => $_getSZ(0);
  @$pb.TagNumber(1)
  set relateId($core.String v) { $_setString(0, v); }
  @$pb.TagNumber(1)
  $core.bool hasRelateId() => $_has(0);
  @$pb.TagNumber(1)
  void clearRelateId() => clearField(1);

  @$pb.TagNumber(2)
  $core.String get code => $_getSZ(1);
  @$pb.TagNumber(2)
  set code($core.String v) { $_setString(1, v); }
  @$pb.TagNumber(2)
  $core.bool hasCode() => $_has(1);
  @$pb.TagNumber(2)
  void clearCode() => clearField(2);

  @$pb.TagNumber(3)
  $core.String get password => $_getSZ(2);
  @$pb.TagNumber(3)
  set password($core.String v) { $_setString(2, v); }
  @$pb.TagNumber(3)
  $core.bool hasPassword() => $_has(2);
  @$pb.TagNumber(3)
  void clearPassword() => clearField(3);
}

class OAuthRequest extends $pb.GeneratedMessage {
  static final $pb.BuilderInfo _i = $pb.BuilderInfo('OAuthRequest', package: const $pb.PackageName('appootb.account'), createEmptyInstance: create)
    ..e<$1.AuthType>(1, 'type', $pb.PbFieldType.QE, defaultOrMaker: $1.AuthType.AUTH_TYPE_UNSPECIFIED, valueOf: $1.AuthType.valueOf, enumValues: $1.AuthType.values)
    ..aQS(2, 'openId')
    ..aQS(3, 'accessToken')
  ;

  OAuthRequest._() : super();
  factory OAuthRequest() => create();
  factory OAuthRequest.fromBuffer($core.List<$core.int> i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromBuffer(i, r);
  factory OAuthRequest.fromJson($core.String i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromJson(i, r);
  OAuthRequest clone() => OAuthRequest()..mergeFromMessage(this);
  OAuthRequest copyWith(void Function(OAuthRequest) updates) => super.copyWith((message) => updates(message as OAuthRequest));
  $pb.BuilderInfo get info_ => _i;
  @$core.pragma('dart2js:noInline')
  static OAuthRequest create() => OAuthRequest._();
  OAuthRequest createEmptyInstance() => create();
  static $pb.PbList<OAuthRequest> createRepeated() => $pb.PbList<OAuthRequest>();
  @$core.pragma('dart2js:noInline')
  static OAuthRequest getDefault() => _defaultInstance ??= $pb.GeneratedMessage.$_defaultFor<OAuthRequest>(create);
  static OAuthRequest _defaultInstance;

  @$pb.TagNumber(1)
  $1.AuthType get type => $_getN(0);
  @$pb.TagNumber(1)
  set type($1.AuthType v) { setField(1, v); }
  @$pb.TagNumber(1)
  $core.bool hasType() => $_has(0);
  @$pb.TagNumber(1)
  void clearType() => clearField(1);

  @$pb.TagNumber(2)
  $core.String get openId => $_getSZ(1);
  @$pb.TagNumber(2)
  set openId($core.String v) { $_setString(1, v); }
  @$pb.TagNumber(2)
  $core.bool hasOpenId() => $_has(1);
  @$pb.TagNumber(2)
  void clearOpenId() => clearField(2);

  @$pb.TagNumber(3)
  $core.String get accessToken => $_getSZ(2);
  @$pb.TagNumber(3)
  set accessToken($core.String v) { $_setString(2, v); }
  @$pb.TagNumber(3)
  $core.bool hasAccessToken() => $_has(2);
  @$pb.TagNumber(3)
  void clearAccessToken() => clearField(3);
}

