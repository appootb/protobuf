///
//  Generated code. Do not modify.
//  source: password.proto
//
// @dart = 2.3
// ignore_for_file: camel_case_types,non_constant_identifier_names,library_prefixes,unused_import,unused_shown_name,return_of_invalid_type

import 'dart:core' as $core;

import 'package:protobuf/protobuf.dart' as $pb;

class PasswordRequest extends $pb.GeneratedMessage {
  static final $pb.BuilderInfo _i = $pb.BuilderInfo('PasswordRequest', package: const $pb.PackageName('appootb.account'), createEmptyInstance: create)
    ..aQS(1, 'password')
    ..aOS(2, 'oldPassword')
    ..aOS(3, 'verifyCode')
  ;

  PasswordRequest._() : super();
  factory PasswordRequest() => create();
  factory PasswordRequest.fromBuffer($core.List<$core.int> i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromBuffer(i, r);
  factory PasswordRequest.fromJson($core.String i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromJson(i, r);
  PasswordRequest clone() => PasswordRequest()..mergeFromMessage(this);
  PasswordRequest copyWith(void Function(PasswordRequest) updates) => super.copyWith((message) => updates(message as PasswordRequest));
  $pb.BuilderInfo get info_ => _i;
  @$core.pragma('dart2js:noInline')
  static PasswordRequest create() => PasswordRequest._();
  PasswordRequest createEmptyInstance() => create();
  static $pb.PbList<PasswordRequest> createRepeated() => $pb.PbList<PasswordRequest>();
  @$core.pragma('dart2js:noInline')
  static PasswordRequest getDefault() => _defaultInstance ??= $pb.GeneratedMessage.$_defaultFor<PasswordRequest>(create);
  static PasswordRequest _defaultInstance;

  @$pb.TagNumber(1)
  $core.String get password => $_getSZ(0);
  @$pb.TagNumber(1)
  set password($core.String v) { $_setString(0, v); }
  @$pb.TagNumber(1)
  $core.bool hasPassword() => $_has(0);
  @$pb.TagNumber(1)
  void clearPassword() => clearField(1);

  @$pb.TagNumber(2)
  $core.String get oldPassword => $_getSZ(1);
  @$pb.TagNumber(2)
  set oldPassword($core.String v) { $_setString(1, v); }
  @$pb.TagNumber(2)
  $core.bool hasOldPassword() => $_has(1);
  @$pb.TagNumber(2)
  void clearOldPassword() => clearField(2);

  @$pb.TagNumber(3)
  $core.String get verifyCode => $_getSZ(2);
  @$pb.TagNumber(3)
  set verifyCode($core.String v) { $_setString(2, v); }
  @$pb.TagNumber(3)
  $core.bool hasVerifyCode() => $_has(2);
  @$pb.TagNumber(3)
  void clearVerifyCode() => clearField(3);
}

