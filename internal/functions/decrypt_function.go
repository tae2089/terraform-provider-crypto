package functions

import (
	"context"
	"encoding/hex"

	"github.com/hashicorp/terraform-plugin-framework/function"
	"github.com/tae2089/terraform-provider-crypto/internal/util"
)

var (
	_ function.Function = &DecryptFunction{}
)

func NewDecryptFunction() function.Function {
	return &DecryptFunction{}
}

type DecryptFunction struct{}

func (r *DecryptFunction) Metadata(_ context.Context, req function.MetadataRequest, resp *function.MetadataResponse) {
	resp.Name = "decrypt"
}

func (r *DecryptFunction) Definition(_ context.Context, _ function.DefinitionRequest, resp *function.DefinitionResponse) {
	resp.Definition = function.Definition{
		Summary:             "Decrypt function",
		MarkdownDescription: "Decrypts the given value using the key path provided.",
		Parameters: []function.Parameter{
			function.StringParameter{
				Name:                "key_path",
				MarkdownDescription: "key path to use for decryption",
			},
			function.StringParameter{
				Name:                "value",
				MarkdownDescription: "encrypted value to decrypt",
			},
		},
		Return: function.StringReturn{},
	}
}

func (r *DecryptFunction) Run(ctx context.Context, req function.RunRequest, resp *function.RunResponse) {
	var keyPath, value string

	resp.Error = function.ConcatFuncErrors(req.Arguments.Get(ctx, &keyPath, &value))

	decodeData, err := hex.DecodeString(value)
	if err != nil {
		resp.Error = function.ConcatFuncErrors(resp.Error, function.NewFuncError(err.Error()))
		return
	}
	decryptedData, err := util.Decrypt(keyPath, decodeData)
	if err != nil {
		resp.Error = function.ConcatFuncErrors(resp.Error, function.NewFuncError(err.Error()))
		return
	}
	resp.Error = function.ConcatFuncErrors(resp.Error, resp.Result.Set(ctx, string(decryptedData)))
}
