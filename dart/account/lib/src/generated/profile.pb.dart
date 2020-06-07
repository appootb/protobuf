///
//  Generated code. Do not modify.
//  source: profile.proto
//
// @dart = 2.3
// ignore_for_file: camel_case_types,non_constant_identifier_names,library_prefixes,unused_import,unused_shown_name,return_of_invalid_type

import 'dart:core' as $core;

import 'package:protobuf/protobuf.dart' as $pb;

import 'google/protobuf/struct.pb.dart' as $12;

class AccountProfile extends $pb.GeneratedMessage {
  static final $pb.BuilderInfo _i = $pb.BuilderInfo('AccountProfile', package: const $pb.PackageName('appootb.account'), createEmptyInstance: create)
    ..aQS(1, 'name')
    ..aOM<$12.Struct>(2, 'value', subBuilder: $12.Struct.create)
  ;

  AccountProfile._() : super();
  factory AccountProfile() => create();
  factory AccountProfile.fromBuffer($core.List<$core.int> i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromBuffer(i, r);
  factory AccountProfile.fromJson($core.String i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromJson(i, r);
  AccountProfile clone() => AccountProfile()..mergeFromMessage(this);
  AccountProfile copyWith(void Function(AccountProfile) updates) => super.copyWith((message) => updates(message as AccountProfile));
  $pb.BuilderInfo get info_ => _i;
  @$core.pragma('dart2js:noInline')
  static AccountProfile create() => AccountProfile._();
  AccountProfile createEmptyInstance() => create();
  static $pb.PbList<AccountProfile> createRepeated() => $pb.PbList<AccountProfile>();
  @$core.pragma('dart2js:noInline')
  static AccountProfile getDefault() => _defaultInstance ??= $pb.GeneratedMessage.$_defaultFor<AccountProfile>(create);
  static AccountProfile _defaultInstance;

  @$pb.TagNumber(1)
  $core.String get name => $_getSZ(0);
  @$pb.TagNumber(1)
  set name($core.String v) { $_setString(0, v); }
  @$pb.TagNumber(1)
  $core.bool hasName() => $_has(0);
  @$pb.TagNumber(1)
  void clearName() => clearField(1);

  @$pb.TagNumber(2)
  $12.Struct get value => $_getN(1);
  @$pb.TagNumber(2)
  set value($12.Struct v) { setField(2, v); }
  @$pb.TagNumber(2)
  $core.bool hasValue() => $_has(1);
  @$pb.TagNumber(2)
  void clearValue() => clearField(2);
  @$pb.TagNumber(2)
  $12.Struct ensureValue() => $_ensure(1);
}

class AccountProfiles extends $pb.GeneratedMessage {
  static final $pb.BuilderInfo _i = $pb.BuilderInfo('AccountProfiles', package: const $pb.PackageName('appootb.account'), createEmptyInstance: create)
    ..m<$core.String, $12.Struct>(1, 'profiles', entryClassName: 'AccountProfiles.ProfilesEntry', keyFieldType: $pb.PbFieldType.OS, valueFieldType: $pb.PbFieldType.OM, valueCreator: $12.Struct.create, packageName: const $pb.PackageName('appootb.account'))
    ..hasRequiredFields = false
  ;

  AccountProfiles._() : super();
  factory AccountProfiles() => create();
  factory AccountProfiles.fromBuffer($core.List<$core.int> i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromBuffer(i, r);
  factory AccountProfiles.fromJson($core.String i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromJson(i, r);
  AccountProfiles clone() => AccountProfiles()..mergeFromMessage(this);
  AccountProfiles copyWith(void Function(AccountProfiles) updates) => super.copyWith((message) => updates(message as AccountProfiles));
  $pb.BuilderInfo get info_ => _i;
  @$core.pragma('dart2js:noInline')
  static AccountProfiles create() => AccountProfiles._();
  AccountProfiles createEmptyInstance() => create();
  static $pb.PbList<AccountProfiles> createRepeated() => $pb.PbList<AccountProfiles>();
  @$core.pragma('dart2js:noInline')
  static AccountProfiles getDefault() => _defaultInstance ??= $pb.GeneratedMessage.$_defaultFor<AccountProfiles>(create);
  static AccountProfiles _defaultInstance;

  @$pb.TagNumber(1)
  $core.Map<$core.String, $12.Struct> get profiles => $_getMap(0);
}

