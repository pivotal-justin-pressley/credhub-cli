package commands

import (
	"time"

	"code.cloudfoundry.org/credhub-cli/config"
	"code.cloudfoundry.org/credhub-cli/credhub"
)

type CredhubCommand struct {
	API              ApiCommand              `command:"api"        alias:"a" description:"Get or set the CredHub API target where commands are sent" long-description:"Get or set the CredHub API target where commands are sent. The api command without any flags will return the current target. If --ca-cert or --skip-tls-validation are provided, these preferences will be cached for future requests."`
	Delete           DeleteCommand           `command:"delete"     alias:"d" description:"Delete a credential" long-description:"Delete a credential. This will delete all versions of the credential."`
	Export           ExportCommand           `command:"export"     alias:"e" description:"Export all credentials" long-description:"Export all credentials"`
	Find             FindCommand             `command:"find"       alias:"f" description:"Find stored credential names or paths based on query parameters" long-description:"Find stored credential names or paths based on query parameters"`
	Generate         GenerateCommand         `command:"generate"   alias:"n" description:"Generate and set a credential value" long-description:"Set a credential with generated value(s). A type must be specified when generating a credential. The provided flags are used to set parameters for the credential that is generated, e.g. a certificate credential may use --common-name, --duration and --self-sign to generate an appropriate value. Supported credential types are prefixed in the flag description."`
	Get              GetCommand              `command:"get"        alias:"g" description:"Get a credential value" long-description:"Get a credential value by name or ID"`
	Import           ImportCommand           `command:"import"     alias:"i" description:"Set multiple credential values" long-description:"Set multiple credential values from import file. File must be in yaml format containing a list of credentials under the key 'credentials'. Name, type and value are required for each credential in the list."`
	Interpolate      InterpolateCommand      `command:"interpolate" description:"Fill a template with values returned from CredHub" long-description:"Fill a template with values returned from CredHub.\n\nUses double-paren placeholders in the style of the bosh cli. Example:\n\n---\nsomething-stored-in-credhub: ((path/to/var))\nsomething-else: static value\n\nIn the above example, the whole value of the cred will be inserted.\nFor instance, if path/to/var is of type ssh, the output will have all the credential's fields, like this:\n\n---\nsomething-stored-in-credhub:\n  private_key: fake-private-key\n  public_key: fake-public-key\n  public_key_fingerprint: fake-fingerprint\nsome-other-key: static value\n\nIf you want just the password value, you'd need to use ((path/to/var.public_key)),\nwhich would only have the specified field, like this:\n\n---\nsomething-stored-in-credhub: fake-public-key\nsomething-else: static value\n\nIf the prefix flag is provided, the given prefix will be prepended\nto any credentials that do not start with the '/' character.\nExample:\n\n---\nsomething: ((/env-specific-path/path/to/var))\nsame-thing: ((path/to/var))\n\nWhen this example is used with the prefix flag 'env-specific-path', they will be evaluated to the same thing."`
	Login            LoginCommand            `command:"login"      alias:"l" description:"Authenticate with CredHub" long-description:"Authenticate with CredHub. UAA password and client credential grants are supported. If client credentials exist in the environment, authentication will be performed automatically without the need to explicitly call this command."`
	Logout           LogoutCommand           `command:"logout"     alias:"o" description:"Discard authenticated user session" long-description:"Discard authenticated session. Refresh token revocation will be attempted for password grants."`
	Regenerate       RegenerateCommand       `command:"regenerate" alias:"r" description:"Generate and set a credential value using the same attributes as the stored value" long-description:"Set a credential with a generated value using the same attributes as the stored value"`
	BulkRegenerate   BulkRegenerateCommand   `command:"bulk-regenerate" description:"Recursively regenerate all certificates signed by the provided certificate" long-description:"Recursively regenerate all certificates signed by the provided certificate"`
	Set              SetCommand              `command:"set"        alias:"s" description:"Set a credential with a provided value" long-description:"Set a credential with provided value(s). A type must be specified when setting a credential. The provided flags are used to set specific values of a credential, e.g. a certificate credential may use --root, --certificate and --private to set each value. Supported credential types are prefixed in the flag description."`
	Curl             CurlCommand             `command:"curl"       description:"Make an arbitrary request to the targeted CredHub server." long-description:"Make an arbitrary request to the targeted CredHub server"`
	SetPermission    SetPermissionCommand    `command:"set-permission" description:"Set permissions for an actor on a given path." long-description:"Set permissions for an actor on a given path"`
	GetPermission    GetPermissionCommand    `command:"get-permission" description:"Get permissions for an actor on a given path." long-description:"Get permissions for an actor on a given path"`
	DeletePermission DeletePermissionCommand `command:"delete-permission" description:"Delete permissions for an actor on a given path." long-description:"Delete permissions for an actor on a given path"`

	HttpTimeout *time.Duration `long:"http-timeout" env:"CREDHUB_HTTP_TIMEOUT" description:"Http timeout for http-client. Needs to have unit passed in (i.e. 30s, 1m)"`

	Version func() `long:"version" description:"Version of CLI and targeted CredHub API"`
	Token   func() `long:"token" description:"Return your current CredHub authentication token"`
}

var CredHub CredhubCommand

type ClientCommand struct {
	client *credhub.CredHub
}

func (n *ClientCommand) SetClient(client *credhub.CredHub) {
	n.client = client
}

type ConfigCommand struct {
	config config.Config
}

func (n *ConfigCommand) SetConfig(config config.Config) {
	n.config = config
}
