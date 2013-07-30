package pkcs11

// Generated with: awk '/#define CK[AFR]/{ print $2 "=" $3 }' pkcs11t.h
const (
	CKF_TOKEN_PRESENT                    = 0x00000001
	CKF_REMOVABLE_DEVICE                 = 0x00000002
	CKF_HW_SLOT                          = 0x00000004
	CKF_RNG                              = 0x00000001
	CKF_WRITE_PROTECTED                  = 0x00000002
	CKF_LOGIN_REQUIRED                   = 0x00000004
	CKF_USER_PIN_INITIALIZED             = 0x00000008
	CKF_RESTORE_KEY_NOT_NEEDED           = 0x00000020
	CKF_CLOCK_ON_TOKEN                   = 0x00000040
	CKF_PROTECTED_AUTHENTICATION_PATH    = 0x00000100
	CKF_DUAL_CRYPTO_OPERATIONS           = 0x00000200
	CKF_TOKEN_INITIALIZED                = 0x00000400
	CKF_SECONDARY_AUTHENTICATION         = 0x00000800
	CKF_USER_PIN_COUNT_LOW               = 0x00010000
	CKF_USER_PIN_FINAL_TRY               = 0x00020000
	CKF_USER_PIN_LOCKED                  = 0x00040000
	CKF_USER_PIN_TO_BE_CHANGED           = 0x00080000
	CKF_SO_PIN_COUNT_LOW                 = 0x00100000
	CKF_SO_PIN_FINAL_TRY                 = 0x00200000
	CKF_SO_PIN_LOCKED                    = 0x00400000
	CKF_SO_PIN_TO_BE_CHANGED             = 0x00800000
	CKF_RW_SESSION                       = 0x00000002
	CKF_SERIAL_SESSION                   = 0x00000004
	CKF_ARRAY_ATTRIBUTE                  = 0x40000000
	CKA_CLASS                            = 0x00000000
	CKA_TOKEN                            = 0x00000001
	CKA_PRIVATE                          = 0x00000002
	CKA_LABEL                            = 0x00000003
	CKA_APPLICATION                      = 0x00000010
	CKA_VALUE                            = 0x00000011
	CKA_OBJECT_ID                        = 0x00000012
	CKA_CERTIFICATE_TYPE                 = 0x00000080
	CKA_ISSUER                           = 0x00000081
	CKA_SERIAL_NUMBER                    = 0x00000082
	CKA_AC_ISSUER                        = 0x00000083
	CKA_OWNER                            = 0x00000084
	CKA_ATTR_TYPES                       = 0x00000085
	CKA_TRUSTED                          = 0x00000086
	CKA_CERTIFICATE_CATEGORY             = 0x00000087
	CKA_JAVA_MIDP_SECURITY_DOMAIN        = 0x00000088
	CKA_URL                              = 0x00000089
	CKA_HASH_OF_SUBJECT_PUBLIC_KEY       = 0x0000008A
	CKA_HASH_OF_ISSUER_PUBLIC_KEY        = 0x0000008B
	CKA_CHECK_VALUE                      = 0x00000090
	CKA_KEY_TYPE                         = 0x00000100
	CKA_SUBJECT                          = 0x00000101
	CKA_ID                               = 0x00000102
	CKA_SENSITIVE                        = 0x00000103
	CKA_ENCRYPT                          = 0x00000104
	CKA_DECRYPT                          = 0x00000105
	CKA_WRAP                             = 0x00000106
	CKA_UNWRAP                           = 0x00000107
	CKA_SIGN                             = 0x00000108
	CKA_SIGN_RECOVER                     = 0x00000109
	CKA_VERIFY                           = 0x0000010A
	CKA_VERIFY_RECOVER                   = 0x0000010B
	CKA_DERIVE                           = 0x0000010C
	CKA_START_DATE                       = 0x00000110
	CKA_END_DATE                         = 0x00000111
	CKA_MODULUS                          = 0x00000120
	CKA_MODULUS_BITS                     = 0x00000121
	CKA_PUBLIC_EXPONENT                  = 0x00000122
	CKA_PRIVATE_EXPONENT                 = 0x00000123
	CKA_PRIME_1                          = 0x00000124
	CKA_PRIME_2                          = 0x00000125
	CKA_EXPONENT_1                       = 0x00000126
	CKA_EXPONENT_2                       = 0x00000127
	CKA_COEFFICIENT                      = 0x00000128
	CKA_PRIME                            = 0x00000130
	CKA_SUBPRIME                         = 0x00000131
	CKA_BASE                             = 0x00000132
	CKA_PRIME_BITS                       = 0x00000133
	CKA_SUBPRIME_BITS                    = 0x00000134
	CKA_SUB_PRIME_BITS                   = CKA_SUBPRIME_BITS
	CKA_VALUE_BITS                       = 0x00000160
	CKA_VALUE_LEN                        = 0x00000161
	CKA_EXTRACTABLE                      = 0x00000162
	CKA_LOCAL                            = 0x00000163
	CKA_NEVER_EXTRACTABLE                = 0x00000164
	CKA_ALWAYS_SENSITIVE                 = 0x00000165
	CKA_KEY_GEN_MECHANISM                = 0x00000166
	CKA_MODIFIABLE                       = 0x00000170
	CKA_ECDSA_PARAMS                     = 0x00000180
	CKA_EC_PARAMS                        = 0x00000180
	CKA_EC_POINT                         = 0x00000181
	CKA_SECONDARY_AUTH                   = 0x00000200
	CKA_AUTH_PIN_FLAGS                   = 0x00000201
	CKA_ALWAYS_AUTHENTICATE              = 0x00000202
	CKA_WRAP_WITH_TRUSTED                = 0x00000210
	CKA_WRAP_TEMPLATE                    = (CKF_ARRAY_ATTRIBUTE | 0x00000211)
	CKA_UNWRAP_TEMPLATE                  = (CKF_ARRAY_ATTRIBUTE | 0x00000212)
	CKA_OTP_FORMAT                       = 0x00000220
	CKA_OTP_LENGTH                       = 0x00000221
	CKA_OTP_TIME_INTERVAL                = 0x00000222
	CKA_OTP_USER_FRIENDLY_MODE           = 0x00000223
	CKA_OTP_CHALLENGE_REQUIREMENT        = 0x00000224
	CKA_OTP_TIME_REQUIREMENT             = 0x00000225
	CKA_OTP_COUNTER_REQUIREMENT          = 0x00000226
	CKA_OTP_PIN_REQUIREMENT              = 0x00000227
	CKA_OTP_COUNTER                      = 0x0000022E
	CKA_OTP_TIME                         = 0x0000022F
	CKA_OTP_USER_IDENTIFIER              = 0x0000022A
	CKA_OTP_SERVICE_IDENTIFIER           = 0x0000022B
	CKA_OTP_SERVICE_LOGO                 = 0x0000022C
	CKA_OTP_SERVICE_LOGO_TYPE            = 0x0000022D
	CKA_HW_FEATURE_TYPE                  = 0x00000300
	CKA_RESET_ON_INIT                    = 0x00000301
	CKA_HAS_RESET                        = 0x00000302
	CKA_PIXEL_X                          = 0x00000400
	CKA_PIXEL_Y                          = 0x00000401
	CKA_RESOLUTION                       = 0x00000402
	CKA_CHAR_ROWS                        = 0x00000403
	CKA_CHAR_COLUMNS                     = 0x00000404
	CKA_COLOR                            = 0x00000405
	CKA_BITS_PER_PIXEL                   = 0x00000406
	CKA_CHAR_SETS                        = 0x00000480
	CKA_ENCODING_METHODS                 = 0x00000481
	CKA_MIME_TYPES                       = 0x00000482
	CKA_MECHANISM_TYPE                   = 0x00000500
	CKA_REQUIRED_CMS_ATTRIBUTES          = 0x00000501
	CKA_DEFAULT_CMS_ATTRIBUTES           = 0x00000502
	CKA_SUPPORTED_CMS_ATTRIBUTES         = 0x00000503
	CKA_ALLOWED_MECHANISMS               = (CKF_ARRAY_ATTRIBUTE | 0x00000600)
	CKA_VENDOR_DEFINED                   = 0x80000000
	CKF_HW                               = 0x00000001
	CKF_ENCRYPT                          = 0x00000100
	CKF_DECRYPT                          = 0x00000200
	CKF_DIGEST                           = 0x00000400
	CKF_SIGN                             = 0x00000800
	CKF_SIGN_RECOVER                     = 0x00001000
	CKF_VERIFY                           = 0x00002000
	CKF_VERIFY_RECOVER                   = 0x00004000
	CKF_GENERATE                         = 0x00008000
	CKF_GENERATE_KEY_PAIR                = 0x00010000
	CKF_WRAP                             = 0x00020000
	CKF_UNWRAP                           = 0x00040000
	CKF_DERIVE                           = 0x00080000
	CKF_EC_F_P                           = 0x00100000
	CKF_EC_F_2M                          = 0x00200000
	CKF_EC_ECPARAMETERS                  = 0x00400000
	CKF_EC_NAMEDCURVE                    = 0x00800000
	CKF_EC_UNCOMPRESS                    = 0x01000000
	CKF_EC_COMPRESS                      = 0x02000000
	CKF_EXTENSION                        = 0x80000000
	CKR_OK                               = 0x00000000
	CKR_CANCEL                           = 0x00000001
	CKR_HOST_MEMORY                      = 0x00000002
	CKR_SLOT_ID_INVALID                  = 0x00000003
	CKR_GENERAL_ERROR                    = 0x00000005
	CKR_FUNCTION_FAILED                  = 0x00000006
	CKR_ARGUMENTS_BAD                    = 0x00000007
	CKR_NO_EVENT                         = 0x00000008
	CKR_NEED_TO_CREATE_THREADS           = 0x00000009
	CKR_CANT_LOCK                        = 0x0000000A
	CKR_ATTRIBUTE_READ_ONLY              = 0x00000010
	CKR_ATTRIBUTE_SENSITIVE              = 0x00000011
	CKR_ATTRIBUTE_TYPE_INVALID           = 0x00000012
	CKR_ATTRIBUTE_VALUE_INVALID          = 0x00000013
	CKR_DATA_INVALID                     = 0x00000020
	CKR_DATA_LEN_RANGE                   = 0x00000021
	CKR_DEVICE_ERROR                     = 0x00000030
	CKR_DEVICE_MEMORY                    = 0x00000031
	CKR_DEVICE_REMOVED                   = 0x00000032
	CKR_ENCRYPTED_DATA_INVALID           = 0x00000040
	CKR_ENCRYPTED_DATA_LEN_RANGE         = 0x00000041
	CKR_FUNCTION_CANCELED                = 0x00000050
	CKR_FUNCTION_NOT_PARALLEL            = 0x00000051
	CKR_FUNCTION_NOT_SUPPORTED           = 0x00000054
	CKR_KEY_HANDLE_INVALID               = 0x00000060
	CKR_KEY_SIZE_RANGE                   = 0x00000062
	CKR_KEY_TYPE_INCONSISTENT            = 0x00000063
	CKR_KEY_NOT_NEEDED                   = 0x00000064
	CKR_KEY_CHANGED                      = 0x00000065
	CKR_KEY_NEEDED                       = 0x00000066
	CKR_KEY_INDIGESTIBLE                 = 0x00000067
	CKR_KEY_FUNCTION_NOT_PERMITTED       = 0x00000068
	CKR_KEY_NOT_WRAPPABLE                = 0x00000069
	CKR_KEY_UNEXTRACTABLE                = 0x0000006A
	CKR_MECHANISM_INVALID                = 0x00000070
	CKR_MECHANISM_PARAM_INVALID          = 0x00000071
	CKR_OBJECT_HANDLE_INVALID            = 0x00000082
	CKR_OPERATION_ACTIVE                 = 0x00000090
	CKR_OPERATION_NOT_INITIALIZED        = 0x00000091
	CKR_PIN_INCORRECT                    = 0x000000A0
	CKR_PIN_INVALID                      = 0x000000A1
	CKR_PIN_LEN_RANGE                    = 0x000000A2
	CKR_PIN_EXPIRED                      = 0x000000A3
	CKR_PIN_LOCKED                       = 0x000000A4
	CKR_SESSION_CLOSED                   = 0x000000B0
	CKR_SESSION_COUNT                    = 0x000000B1
	CKR_SESSION_HANDLE_INVALID           = 0x000000B3
	CKR_SESSION_PARALLEL_NOT_SUPPORTED   = 0x000000B4
	CKR_SESSION_READ_ONLY                = 0x000000B5
	CKR_SESSION_EXISTS                   = 0x000000B6
	CKR_SESSION_READ_ONLY_EXISTS         = 0x000000B7
	CKR_SESSION_READ_WRITE_SO_EXISTS     = 0x000000B8
	CKR_SIGNATURE_INVALID                = 0x000000C0
	CKR_SIGNATURE_LEN_RANGE              = 0x000000C1
	CKR_TEMPLATE_INCOMPLETE              = 0x000000D0
	CKR_TEMPLATE_INCONSISTENT            = 0x000000D1
	CKR_TOKEN_NOT_PRESENT                = 0x000000E0
	CKR_TOKEN_NOT_RECOGNIZED             = 0x000000E1
	CKR_TOKEN_WRITE_PROTECTED            = 0x000000E2
	CKR_UNWRAPPING_KEY_HANDLE_INVALID    = 0x000000F0
	CKR_UNWRAPPING_KEY_SIZE_RANGE        = 0x000000F1
	CKR_UNWRAPPING_KEY_TYPE_INCONSISTENT = 0x000000F2
	CKR_USER_ALREADY_LOGGED_IN           = 0x00000100
	CKR_USER_NOT_LOGGED_IN               = 0x00000101
	CKR_USER_PIN_NOT_INITIALIZED         = 0x00000102
	CKR_USER_TYPE_INVALID                = 0x00000103
	CKR_USER_ANOTHER_ALREADY_LOGGED_IN   = 0x00000104
	CKR_USER_TOO_MANY_TYPES              = 0x00000105
	CKR_WRAPPED_KEY_INVALID              = 0x00000110
	CKR_WRAPPED_KEY_LEN_RANGE            = 0x00000112
	CKR_WRAPPING_KEY_HANDLE_INVALID      = 0x00000113
	CKR_WRAPPING_KEY_SIZE_RANGE          = 0x00000114
	CKR_WRAPPING_KEY_TYPE_INCONSISTENT   = 0x00000115
	CKR_RANDOM_SEED_NOT_SUPPORTED        = 0x00000120
	CKR_RANDOM_NO_RNG                    = 0x00000121
	CKR_DOMAIN_PARAMS_INVALID            = 0x00000130
	CKR_BUFFER_TOO_SMALL                 = 0x00000150
	CKR_SAVED_STATE_INVALID              = 0x00000160
	CKR_INFORMATION_SENSITIVE            = 0x00000170
	CKR_STATE_UNSAVEABLE                 = 0x00000180
	CKR_CRYPTOKI_NOT_INITIALIZED         = 0x00000190
	CKR_CRYPTOKI_ALREADY_INITIALIZED     = 0x00000191
	CKR_MUTEX_BAD                        = 0x000001A0
	CKR_MUTEX_NOT_LOCKED                 = 0x000001A1
	CKR_NEW_PIN_MODE                     = 0x000001B0
	CKR_NEXT_OTP                         = 0x000001B1
	CKR_FUNCTION_REJECTED                = 0x00000200
	CKR_VENDOR_DEFINED                   = 0x80000000
	CKF_LIBRARY_CANT_CREATE_OS_THREADS   = 0x00000001
	CKF_OS_LOCKING_OK                    = 0x00000002
	CKF_DONT_BLOCK                       = 1
	CKF_NEXT_OTP                         = 0x00000001
	CKF_EXCLUDE_TIME                     = 0x00000002
	CKF_EXCLUDE_COUNTER                  = 0x00000004
	CKF_EXCLUDE_CHALLENGE                = 0x00000008
	CKF_EXCLUDE_PIN                      = 0x00000010
	CKF_USER_FRIENDLY_OTP                = 0x00000020
)
