package init

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/lechgu/cartman/internal/algo"
	"github.com/lechgu/cartman/internal/certificates"
	"github.com/lechgu/cartman/internal/encoders"
	"github.com/samber/lo"
	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
)

const (
	dir      = ".cartman"
	keyFile  = "key.pem"
	certFile = "cert.pem"
)

var Cmd = cobra.Command{
	Use:   "init",
	Short: "Initialize CA's certificate and signing key",
	RunE:  doInit,
}

func doInit(cmd *cobra.Command, args []string) error {
	_, ok := lo.Find(algo.Supported, func(item string) bool {
		return algorithm == item
	})
	if !ok {
		return fmt.Errorf("unsupported algorithm: %q\nValid options: %s", algorithm, strings.Join(algo.Supported, ", "))
	}
	if _, err := os.Stat(dir); err == nil {
		if !force {
			return fmt.Errorf("directory %s already exists", dir)
		}
		if err := os.RemoveAll(dir); err != nil {
			return err
		}
	}
	if err := os.Mkdir(dir, 0700); err != nil {
		return err
	}

	handler, err := algo.NewHandler(algorithm)
	if err != nil {
		return err
	}
	keypair, err := handler.GenerateKeyPair()
	if err != nil {
		return err
	}

	cert, err := certificates.InitRoot(handler, keypair, validityDays, &subject)
	if err != nil {
		return err
	}

	buf, err := encoders.EncodeCertificate(cert)
	if err != nil {

	}
	certPath := filepath.Join(dir, certFile)
	if err := os.WriteFile(certPath, buf, 0600); err != nil {
		return err
	}

	buf, err = encoders.EncodePrivateKey(keypair.PrivateKey)
	if err != nil {
		return err
	}
	keyPath := filepath.Join(dir, keyFile)
	if err := os.WriteFile(keyPath, buf, 0600); err != nil {
		return err
	}
	return nil

}

func init() {
	Cmd.Flags().StringVarP(&algorithm, "algo", "a", algo.ECDSA384, "Signature algorithm to use")
	Cmd.Flags().IntVarP(&validityDays, "validity-days", "d", 3650, "Certificate validity days; default: 3560")
	Cmd.Flags().BoolVarP(&force, "force", "f", false, "Force overwrite the .cartman directory, if exists")
	Cmd.Flags().StringVarP(&subject.CommonName, "name", "n", "cartman", "Common Name for the certificate")
	Cmd.SetHelpFunc(func(cmd *cobra.Command, args []string) {
		fmt.Fprintf(cmd.OutOrStdout(), "Usage:\n  %s\n\n", cmd.UseLine())
		if cmd.HasAvailableFlags() {
			fmt.Fprintln(cmd.OutOrStdout(), "Flags:")
			cmd.Flags().VisitAll(func(f *pflag.Flag) {
				var flagLine string
				if f.Shorthand != "" {
					flagLine = fmt.Sprintf("  -%s, --%s", f.Shorthand, f.Name)
				} else {
					flagLine = fmt.Sprintf("      --%s", f.Name)
				}
				fmt.Fprintf(cmd.OutOrStdout(), "%-30s %s\n", flagLine, f.Usage)
			})
		}
	})

}
