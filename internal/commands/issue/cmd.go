package issue

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/lechgu/cartman/internal/algo"
	"github.com/lechgu/cartman/internal/certificates"
	"github.com/lechgu/cartman/internal/encoders"
	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
)

const (
	caDir = ".cartman"
)

var Cmd = cobra.Command{
	Use:   "issue",
	Short: "Issue and sign a certificate",
	RunE:  doIssue,
}

func doIssue(cmd *cobra.Command, args []string) error {
	certFile := filepath.Join(caDir, "cert.pem")
	keyFile := filepath.Join(caDir, "key.pem")

	certBytes, err := os.ReadFile(certFile)
	if err != nil {
		return err
	}
	keyBytes, err := os.ReadFile(keyFile)
	if err != nil {
		return err
	}
	signer, err := encoders.LoadSigner(certBytes, keyBytes)
	if err != nil {
		return err
	}
	handler, err := algo.MatchHandler(signer.Key)
	if err != nil {
		return err
	}
	cert, key, err := certificates.Issue(handler, signer, validityDays, &subject, dnsNames, ipAddresses)
	if err != nil {
		return err
	}
	certPEM, err := encoders.EncodeCertificate(cert)
	if err != nil {
		return err
	}
	keyPEM, err := encoders.EncodePrivateKey(key)
	if err != nil {
		return err
	}

	dir := subject.CommonName

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

	certPath := filepath.Join(dir, "cert.pem")
	if err := os.WriteFile(certPath, certPEM, 0600); err != nil {
		return err
	}

	keyPath := filepath.Join(dir, "key.pem")
	if err := os.WriteFile(keyPath, keyPEM, 0600); err != nil {
		return err
	}

	return nil
}

func init() {
	Cmd.Flags().BoolVarP(&force, "force", "f", false, "Force overwrite the directory, if exists")
	Cmd.Flags().StringVarP(&subject.CommonName, "name", "n", "", "Name for the certificate; required")
	Cmd.Flags().IntVarP(&validityDays, "validity-days", "d", 365, "Certificate validity days; default: 356")
	Cmd.Flags().StringArrayVar(&dnsNames, "dns", nil, "DNS name; repeatable")
	Cmd.Flags().IPSliceVar(&ipAddresses, "ip", nil, "IP address; repeatable")
	Cmd.MarkFlagRequired("name")
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
