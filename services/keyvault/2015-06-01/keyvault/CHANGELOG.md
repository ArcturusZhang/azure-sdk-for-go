
Generated from https://github.com/Azure/azure-rest-api-specs/tree/b97299c968df5f99b724bd1231fd2161731d3b8f

Code generator C:\Users\dapzhang\Documents\workspace\autorest.go

## Breaking Changes

- Function `NewSecretListResultPage` signature has been changed from `(func(context.Context, SecretListResult) (SecretListResult, error))` to `(SecretListResult,func(context.Context, SecretListResult) (SecretListResult, error))`
- Function `NewCertificateIssuerListResultPage` signature has been changed from `(func(context.Context, CertificateIssuerListResult) (CertificateIssuerListResult, error))` to `(CertificateIssuerListResult,func(context.Context, CertificateIssuerListResult) (CertificateIssuerListResult, error))`
- Function `NewCertificateListResultPage` signature has been changed from `(func(context.Context, CertificateListResult) (CertificateListResult, error))` to `(CertificateListResult,func(context.Context, CertificateListResult) (CertificateListResult, error))`
- Function `NewKeyListResultPage` signature has been changed from `(func(context.Context, KeyListResult) (KeyListResult, error))` to `(KeyListResult,func(context.Context, KeyListResult) (KeyListResult, error))`

## New Content

- Function `IssuerAttributes.MarshalJSON() ([]byte,error)` is added
- Function `SecretAttributes.MarshalJSON() ([]byte,error)` is added
- Function `Contacts.MarshalJSON() ([]byte,error)` is added
- Function `CertificateAttributes.MarshalJSON() ([]byte,error)` is added
- Function `CertificatePolicy.MarshalJSON() ([]byte,error)` is added
- Function `IssuerBundle.MarshalJSON() ([]byte,error)` is added
- Function `Attributes.MarshalJSON() ([]byte,error)` is added
- Function `CertificateOperation.MarshalJSON() ([]byte,error)` is added
- Function `KeyAttributes.MarshalJSON() ([]byte,error)` is added
