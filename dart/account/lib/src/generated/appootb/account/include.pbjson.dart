///
//  Generated code. Do not modify.
//  source: appootb/account/include.proto
//
// @dart = 2.3
// ignore_for_file: camel_case_types,non_constant_identifier_names,library_prefixes,unused_import,unused_shown_name,return_of_invalid_type

const Gender$json = const {
  '1': 'Gender',
  '2': const [
    const {'1': 'GENDER_UNSPECIFIED', '2': 0},
    const {'1': 'GENDER_MALE', '2': 1},
    const {'1': 'GENDER_FEMALE', '2': 2},
    const {'1': 'GENDER_OTHER', '2': 3},
  ],
};

const AuthType$json = const {
  '1': 'AuthType',
  '2': const [
    const {'1': 'AUTH_TYPE_UNSPECIFIED', '2': 0},
    const {'1': 'AUTH_TYPE_APPLE_ID', '2': 1},
    const {'1': 'AUTH_TYPE_ACCOUNT', '2': 2},
    const {'1': 'AUTH_TYPE_WECHAT', '2': 3},
    const {'1': 'AUTH_TYPE_QQ', '2': 4},
    const {'1': 'AUTH_TYPE_WEIBO', '2': 5},
  ],
};

const Status$json = const {
  '1': 'Status',
  '2': const [
    const {'1': 'STATUS_UNSPECIFIED', '2': 0},
    const {'1': 'STATUS_ACTIVE', '2': 1},
    const {'1': 'STATUS_BLOCKED', '2': 2},
  ],
};

const Secret$json = const {
  '1': 'Secret',
  '2': const [
    const {'1': 'token', '3': 1, '4': 2, '5': 9, '10': 'token'},
  ],
};

const Info$json = const {
  '1': 'Info',
  '2': const [
    const {'1': 'unique_id', '3': 1, '4': 2, '5': 4, '10': 'uniqueId'},
    const {'1': 'nickname', '3': 2, '4': 2, '5': 9, '10': 'nickname'},
    const {'1': 'avatar', '3': 3, '4': 2, '5': 9, '10': 'avatar'},
    const {'1': 'status', '3': 4, '4': 2, '5': 14, '6': '.appootb.account.Status', '10': 'status'},
    const {'1': 'signature', '3': 5, '4': 1, '5': 9, '10': 'signature'},
    const {'1': 'gender', '3': 6, '4': 1, '5': 14, '6': '.appootb.account.Gender', '10': 'gender'},
    const {'1': 'signs', '3': 7, '4': 1, '5': 9, '10': 'signs'},
    const {'1': 'location', '3': 8, '4': 1, '5': 9, '10': 'location'},
    const {'1': 'secret', '3': 9, '4': 1, '5': 11, '6': '.appootb.account.Secret', '10': 'secret'},
    const {'1': 'create_at', '3': 98, '4': 1, '5': 11, '6': '.google.protobuf.Timestamp', '10': 'createAt'},
    const {'1': 'extend', '3': 99, '4': 1, '5': 11, '6': '.google.protobuf.Any', '10': 'extend'},
  ],
};

const Infos$json = const {
  '1': 'Infos',
  '2': const [
    const {'1': 'accounts', '3': 1, '4': 3, '5': 11, '6': '.appootb.account.Infos.AccountsEntry', '10': 'accounts'},
  ],
  '3': const [Infos_AccountsEntry$json],
};

const Infos_AccountsEntry$json = const {
  '1': 'AccountsEntry',
  '2': const [
    const {'1': 'key', '3': 1, '4': 1, '5': 4, '10': 'key'},
    const {'1': 'value', '3': 2, '4': 1, '5': 11, '6': '.appootb.account.Info', '10': 'value'},
  ],
  '7': const {'7': true},
};

