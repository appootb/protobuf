///
//  Generated code. Do not modify.
//  source: appootb/secret/include.proto
//
// @dart = 2.3
// ignore_for_file: camel_case_types,non_constant_identifier_names,library_prefixes,unused_import,unused_shown_name,return_of_invalid_type

const Type$json = const {
  '1': 'Type',
  '2': const [
    const {'1': 'CLIENT', '2': 0},
    const {'1': 'SERVER', '2': 1},
  ],
};

const Algorithm$json = const {
  '1': 'Algorithm',
  '2': const [
    const {'1': 'None', '2': 0},
    const {'1': 'HMAC', '2': 1},
    const {'1': 'RSA', '2': 2},
    const {'1': 'PSS', '2': 3},
    const {'1': 'ECDSA', '2': 4},
    const {'1': 'EdDSA', '2': 5},
  ],
};

const Info$json = const {
  '1': 'Info',
  '2': const [
    const {'1': 'type', '3': 1, '4': 1, '5': 14, '6': '.appootb.secret.Type', '10': 'type'},
    const {'1': 'algorithm', '3': 2, '4': 1, '5': 14, '6': '.appootb.secret.Algorithm', '10': 'algorithm'},
    const {'1': 'issuer', '3': 3, '4': 1, '5': 9, '10': 'issuer'},
    const {'1': 'account', '3': 4, '4': 1, '5': 4, '10': 'account'},
    const {'1': 'key_id', '3': 5, '4': 1, '5': 3, '10': 'keyId'},
    const {'1': 'roles', '3': 6, '4': 3, '5': 9, '10': 'roles'},
    const {'1': 'subject', '3': 11, '4': 1, '5': 14, '6': '.appootb.permission.method.Subject', '10': 'subject'},
    const {'1': 'issued_at', '3': 21, '4': 1, '5': 11, '6': '.google.protobuf.Timestamp', '10': 'issuedAt'},
    const {'1': 'expired_at', '3': 22, '4': 1, '5': 11, '6': '.google.protobuf.Timestamp', '10': 'expiredAt'},
  ],
};

