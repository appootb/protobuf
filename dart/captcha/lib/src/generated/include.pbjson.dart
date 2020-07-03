///
//  Generated code. Do not modify.
//  source: include.proto
//
// @dart = 2.3
// ignore_for_file: camel_case_types,non_constant_identifier_names,library_prefixes,unused_import,unused_shown_name,return_of_invalid_type

const Channel$json = const {
  '1': 'Channel',
  '2': const [
    const {'1': 'CHANNEL_UNSPECIFIED', '2': 0},
    const {'1': 'CHANNEL_EMAIL', '2': 1},
    const {'1': 'CHANNEL_SMS', '2': 2},
    const {'1': 'CHANNEL_PHONE', '2': 3},
    const {'1': 'CHANNEL_OTP', '2': 4},
  ],
};

const Category$json = const {
  '1': 'Category',
  '2': const [
    const {'1': 'CATEGORY_UNSPECIFIED', '2': 0},
    const {'1': 'CATEGORY_REGISTER', '2': 1},
    const {'1': 'CATEGORY_LOGIN', '2': 2},
    const {'1': 'CATEGORY_RESET_PWD', '2': 3},
  ],
};

const CodeRequest$json = const {
  '1': 'CodeRequest',
  '2': const [
    const {'1': 'channel', '3': 1, '4': 2, '5': 14, '6': '.appootb.captcha.Channel', '10': 'channel'},
    const {'1': 'category', '3': 2, '4': 2, '5': 14, '6': '.appootb.captcha.Category', '10': 'category'},
    const {'1': 'target', '3': 3, '4': 2, '5': 9, '10': 'target'},
    const {'1': 'value', '3': 4, '4': 1, '5': 9, '10': 'value'},
  ],
};

