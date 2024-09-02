// Code generated by smithy-go-codegen DO NOT EDIT.

package kms

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/kms/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Returns a unique symmetric data key for use outside of KMS. This operation
// returns a plaintext copy of the data key and a copy that is encrypted under a
// symmetric encryption KMS key that you specify. The bytes in the plaintext key
// are random; they are not related to the caller or the KMS key. You can use the
// plaintext key to encrypt your data outside of KMS and store the encrypted data
// key with the encrypted data.
//
// To generate a data key, specify the symmetric encryption KMS key that will be
// used to encrypt the data key. You cannot use an asymmetric KMS key to encrypt
// data keys. To get the type of your KMS key, use the DescribeKeyoperation.
//
// You must also specify the length of the data key. Use either the KeySpec or
// NumberOfBytes parameters (but not both). For 128-bit and 256-bit data keys, use
// the KeySpec parameter.
//
// To generate a 128-bit SM4 data key (China Regions only), specify a KeySpec
// value of AES_128 or a NumberOfBytes value of 16 . The symmetric encryption key
// used in China Regions to encrypt your data key is an SM4 encryption key.
//
// To get only an encrypted copy of the data key, use GenerateDataKeyWithoutPlaintext. To generate an asymmetric
// data key pair, use the GenerateDataKeyPairor GenerateDataKeyPairWithoutPlaintext operation. To get a cryptographically secure random
// byte string, use GenerateRandom.
//
// You can use an optional encryption context to add additional security to the
// encryption operation. If you specify an EncryptionContext , you must specify the
// same encryption context (a case-sensitive exact match) when decrypting the
// encrypted data key. Otherwise, the request to decrypt fails with an
// InvalidCiphertextException . For more information, see [Encryption Context] in the Key Management
// Service Developer Guide.
//
// GenerateDataKey also supports [Amazon Web Services Nitro Enclaves], which provide an isolated compute environment
// in Amazon EC2. To call GenerateDataKey for an Amazon Web Services Nitro
// enclave, use the [Amazon Web Services Nitro Enclaves SDK]or any Amazon Web Services SDK. Use the Recipient parameter to
// provide the attestation document for the enclave. GenerateDataKey returns a
// copy of the data key encrypted under the specified KMS key, as usual. But
// instead of a plaintext copy of the data key, the response includes a copy of the
// data key encrypted under the public key from the attestation document (
// CiphertextForRecipient ). For information about the interaction between KMS and
// Amazon Web Services Nitro Enclaves, see [How Amazon Web Services Nitro Enclaves uses KMS]in the Key Management Service Developer
// Guide..
//
// The KMS key that you use for this operation must be in a compatible key state.
// For details, see [Key states of KMS keys]in the Key Management Service Developer Guide.
//
// # How to use your data key
//
// We recommend that you use the following pattern to encrypt data locally in your
// application. You can write your own code or use a client-side encryption
// library, such as the [Amazon Web Services Encryption SDK], the [Amazon DynamoDB Encryption Client], or [Amazon S3 client-side encryption] to do these tasks for you.
//
// To encrypt data outside of KMS:
//
//   - Use the GenerateDataKey operation to get a data key.
//
//   - Use the plaintext data key (in the Plaintext field of the response) to
//     encrypt your data outside of KMS. Then erase the plaintext data key from memory.
//
//   - Store the encrypted data key (in the CiphertextBlob field of the response)
//     with the encrypted data.
//
// To decrypt data outside of KMS:
//
//   - Use the Decryptoperation to decrypt the encrypted data key. The operation returns
//     a plaintext copy of the data key.
//
//   - Use the plaintext data key to decrypt data outside of KMS, then erase the
//     plaintext data key from memory.
//
// Cross-account use: Yes. To perform this operation with a KMS key in a different
// Amazon Web Services account, specify the key ARN or alias ARN in the value of
// the KeyId parameter.
//
// Required permissions: [kms:GenerateDataKey] (key policy)
//
// Related operations:
//
// # Decrypt
//
// # Encrypt
//
// # GenerateDataKeyPair
//
// # GenerateDataKeyPairWithoutPlaintext
//
// # GenerateDataKeyWithoutPlaintext
//
// Eventual consistency: The KMS API follows an eventual consistency model. For
// more information, see [KMS eventual consistency].
//
// [Amazon Web Services Encryption SDK]: https://docs.aws.amazon.com/encryption-sdk/latest/developer-guide/
// [Amazon DynamoDB Encryption Client]: https://docs.aws.amazon.com/dynamodb-encryption-client/latest/devguide/
// [Key states of KMS keys]: https://docs.aws.amazon.com/kms/latest/developerguide/key-state.html
// [Encryption Context]: https://docs.aws.amazon.com/kms/latest/developerguide/concepts.html#encrypt_context
// [Amazon Web Services Nitro Enclaves]: https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/nitro-enclave.html
// [Amazon S3 client-side encryption]: https://docs.aws.amazon.com/AmazonS3/latest/dev/UsingClientSideEncryption.html
// [kms:GenerateDataKey]: https://docs.aws.amazon.com/kms/latest/developerguide/kms-api-permissions-reference.html
// [How Amazon Web Services Nitro Enclaves uses KMS]: https://docs.aws.amazon.com/kms/latest/developerguide/services-nitro-enclaves.html
// [KMS eventual consistency]: https://docs.aws.amazon.com/kms/latest/developerguide/programming-eventual-consistency.html
// [Amazon Web Services Nitro Enclaves SDK]: https://docs.aws.amazon.com/enclaves/latest/user/developing-applications.html#sdk
func (c *Client) GenerateDataKey(ctx context.Context, params *GenerateDataKeyInput, optFns ...func(*Options)) (*GenerateDataKeyOutput, error) {
	if params == nil {
		params = &GenerateDataKeyInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "GenerateDataKey", params, optFns, c.addOperationGenerateDataKeyMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*GenerateDataKeyOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type GenerateDataKeyInput struct {

	// Specifies the symmetric encryption KMS key that encrypts the data key. You
	// cannot specify an asymmetric KMS key or a KMS key in a custom key store. To get
	// the type and origin of your KMS key, use the DescribeKeyoperation.
	//
	// To specify a KMS key, use its key ID, key ARN, alias name, or alias ARN. When
	// using an alias name, prefix it with "alias/" . To specify a KMS key in a
	// different Amazon Web Services account, you must use the key ARN or alias ARN.
	//
	// For example:
	//
	//   - Key ID: 1234abcd-12ab-34cd-56ef-1234567890ab
	//
	//   - Key ARN:
	//   arn:aws:kms:us-east-2:111122223333:key/1234abcd-12ab-34cd-56ef-1234567890ab
	//
	//   - Alias name: alias/ExampleAlias
	//
	//   - Alias ARN: arn:aws:kms:us-east-2:111122223333:alias/ExampleAlias
	//
	// To get the key ID and key ARN for a KMS key, use ListKeys or DescribeKey. To get the alias name
	// and alias ARN, use ListAliases.
	//
	// This member is required.
	KeyId *string

	// Checks if your request will succeed. DryRun is an optional parameter.
	//
	// To learn more about how to use this parameter, see [Testing your KMS API calls] in the Key Management
	// Service Developer Guide.
	//
	// [Testing your KMS API calls]: https://docs.aws.amazon.com/kms/latest/developerguide/programming-dryrun.html
	DryRun *bool

	// Specifies the encryption context that will be used when encrypting the data key.
	//
	// Do not include confidential or sensitive information in this field. This field
	// may be displayed in plaintext in CloudTrail logs and other output.
	//
	// An encryption context is a collection of non-secret key-value pairs that
	// represent additional authenticated data. When you use an encryption context to
	// encrypt data, you must specify the same (an exact case-sensitive match)
	// encryption context to decrypt the data. An encryption context is supported only
	// on operations with symmetric encryption KMS keys. On operations with symmetric
	// encryption KMS keys, an encryption context is optional, but it is strongly
	// recommended.
	//
	// For more information, see [Encryption context] in the Key Management Service Developer Guide.
	//
	// [Encryption context]: https://docs.aws.amazon.com/kms/latest/developerguide/concepts.html#encrypt_context
	EncryptionContext map[string]string

	// A list of grant tokens.
	//
	// Use a grant token when your permission to call this operation comes from a new
	// grant that has not yet achieved eventual consistency. For more information, see [Grant token]
	// and [Using a grant token]in the Key Management Service Developer Guide.
	//
	// [Grant token]: https://docs.aws.amazon.com/kms/latest/developerguide/grants.html#grant_token
	// [Using a grant token]: https://docs.aws.amazon.com/kms/latest/developerguide/grant-manage.html#using-grant-token
	GrantTokens []string

	// Specifies the length of the data key. Use AES_128 to generate a 128-bit
	// symmetric key, or AES_256 to generate a 256-bit symmetric key.
	//
	// You must specify either the KeySpec or the NumberOfBytes parameter (but not
	// both) in every GenerateDataKey request.
	KeySpec types.DataKeySpec

	// Specifies the length of the data key in bytes. For example, use the value 64 to
	// generate a 512-bit data key (64 bytes is 512 bits). For 128-bit (16-byte) and
	// 256-bit (32-byte) data keys, use the KeySpec parameter.
	//
	// You must specify either the KeySpec or the NumberOfBytes parameter (but not
	// both) in every GenerateDataKey request.
	NumberOfBytes *int32

	// A signed [attestation document] from an Amazon Web Services Nitro enclave and the encryption
	// algorithm to use with the enclave's public key. The only valid encryption
	// algorithm is RSAES_OAEP_SHA_256 .
	//
	// This parameter only supports attestation documents for Amazon Web Services
	// Nitro Enclaves. To include this parameter, use the [Amazon Web Services Nitro Enclaves SDK]or any Amazon Web Services
	// SDK.
	//
	// When you use this parameter, instead of returning the plaintext data key, KMS
	// encrypts the plaintext data key under the public key in the attestation
	// document, and returns the resulting ciphertext in the CiphertextForRecipient
	// field in the response. This ciphertext can be decrypted only with the private
	// key in the enclave. The CiphertextBlob field in the response contains a copy of
	// the data key encrypted under the KMS key specified by the KeyId parameter. The
	// Plaintext field in the response is null or empty.
	//
	// For information about the interaction between KMS and Amazon Web Services Nitro
	// Enclaves, see [How Amazon Web Services Nitro Enclaves uses KMS]in the Key Management Service Developer Guide.
	//
	// [attestation document]: https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/nitro-enclave-how.html#term-attestdoc
	// [How Amazon Web Services Nitro Enclaves uses KMS]: https://docs.aws.amazon.com/kms/latest/developerguide/services-nitro-enclaves.html
	// [Amazon Web Services Nitro Enclaves SDK]: https://docs.aws.amazon.com/enclaves/latest/user/developing-applications.html#sdk
	Recipient *types.RecipientInfo

	noSmithyDocumentSerde
}

type GenerateDataKeyOutput struct {

	// The encrypted copy of the data key. When you use the HTTP API or the Amazon Web
	// Services CLI, the value is Base64-encoded. Otherwise, it is not Base64-encoded.
	CiphertextBlob []byte

	// The plaintext data key encrypted with the public key from the Nitro enclave.
	// This ciphertext can be decrypted only by using a private key in the Nitro
	// enclave.
	//
	// This field is included in the response only when the Recipient parameter in the
	// request includes a valid attestation document from an Amazon Web Services Nitro
	// enclave. For information about the interaction between KMS and Amazon Web
	// Services Nitro Enclaves, see [How Amazon Web Services Nitro Enclaves uses KMS]in the Key Management Service Developer Guide.
	//
	// [How Amazon Web Services Nitro Enclaves uses KMS]: https://docs.aws.amazon.com/kms/latest/developerguide/services-nitro-enclaves.html
	CiphertextForRecipient []byte

	// The Amazon Resource Name ([key ARN] ) of the KMS key that encrypted the data key.
	//
	// [key ARN]: https://docs.aws.amazon.com/kms/latest/developerguide/concepts.html#key-id-key-ARN
	KeyId *string

	// The plaintext data key. When you use the HTTP API or the Amazon Web Services
	// CLI, the value is Base64-encoded. Otherwise, it is not Base64-encoded. Use this
	// data key to encrypt your data outside of KMS. Then, remove it from memory as
	// soon as possible.
	//
	// If the response includes the CiphertextForRecipient field, the Plaintext field
	// is null or empty.
	Plaintext []byte

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationGenerateDataKeyMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpGenerateDataKey{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpGenerateDataKey{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "GenerateDataKey"); err != nil {
		return fmt.Errorf("add protocol finalizers: %v", err)
	}

	if err = addlegacyEndpointContextSetter(stack, options); err != nil {
		return err
	}
	if err = addSetLoggerMiddleware(stack, options); err != nil {
		return err
	}
	if err = addClientRequestID(stack); err != nil {
		return err
	}
	if err = addComputeContentLength(stack); err != nil {
		return err
	}
	if err = addResolveEndpointMiddleware(stack, options); err != nil {
		return err
	}
	if err = addComputePayloadSHA256(stack); err != nil {
		return err
	}
	if err = addRetry(stack, options); err != nil {
		return err
	}
	if err = addRawResponseToMetadata(stack); err != nil {
		return err
	}
	if err = addRecordResponseTiming(stack); err != nil {
		return err
	}
	if err = addClientUserAgent(stack, options); err != nil {
		return err
	}
	if err = smithyhttp.AddErrorCloseResponseBodyMiddleware(stack); err != nil {
		return err
	}
	if err = smithyhttp.AddCloseResponseBodyMiddleware(stack); err != nil {
		return err
	}
	if err = addSetLegacyContextSigningOptionsMiddleware(stack); err != nil {
		return err
	}
	if err = addTimeOffsetBuild(stack, c); err != nil {
		return err
	}
	if err = addUserAgentRetryMode(stack, options); err != nil {
		return err
	}
	if err = addOpGenerateDataKeyValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opGenerateDataKey(options.Region), middleware.Before); err != nil {
		return err
	}
	if err = addRecursionDetection(stack); err != nil {
		return err
	}
	if err = addRequestIDRetrieverMiddleware(stack); err != nil {
		return err
	}
	if err = addResponseErrorMiddleware(stack); err != nil {
		return err
	}
	if err = addRequestResponseLogging(stack, options); err != nil {
		return err
	}
	if err = addDisableHTTPSMiddleware(stack, options); err != nil {
		return err
	}
	return nil
}

func newServiceMetadataMiddleware_opGenerateDataKey(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "GenerateDataKey",
	}
}
