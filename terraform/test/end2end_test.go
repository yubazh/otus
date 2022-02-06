package test

import (
	"crypto/tls"
	"crypto/x509"
	"database/sql"
	"flag"
	"fmt"
	"io/ioutil"
	"strings"
	"testing"

	"github.com/go-sql-driver/mysql"

	"github.com/gruntwork-io/terratest/modules/logger"
	"github.com/gruntwork-io/terratest/modules/terraform"
	test_structure "github.com/gruntwork-io/terratest/modules/test-structure"
	"golang.org/x/crypto/ssh"
)

var folder = flag.String("folder", "", "Folder ID in Yandex.Cloud")
var sshKeyPath = flag.String("ssh-key-pass", "", "Private ssh key for access to virtual machines")

func TestEndToEndDeploymentScenario(t *testing.T) {
	fixtureFolder := "../"

	test_structure.RunTestStage(t, "setup", func() {
		terraformOptions := &terraform.Options{
			TerraformDir: fixtureFolder,

			Vars: map[string]interface{}{
				"yc_folder": *folder,
			},
		}

		test_structure.SaveTerraformOptions(t, fixtureFolder, terraformOptions)

		terraform.InitAndApply(t, terraformOptions)
	})

	test_structure.RunTestStage(t, "validate", func() {
		fmt.Println("Run some tests...")

		terraformOptions := test_structure.LoadTerraformOptions(t, fixtureFolder)

		// test load balancer ip existing
		loadbalancerIPAddress := terraform.Output(t, terraformOptions, "load_balancer_public_ip")

		if loadbalancerIPAddress == "" {
			t.Fatal("Cannot retrieve the public IP address value for the load balancer.")
		}

		// test ssh connect
		vmLinuxPublicIPAddress := terraform.Output(t, terraformOptions, "vm_linux_public_ip_address")

		key, err := ioutil.ReadFile(*sshKeyPath)
		if err != nil {
			t.Fatalf("Unable to read private key: %v", err)
		}

		signer, err := ssh.ParsePrivateKey(key)
		if err != nil {
			t.Fatalf("Unable to parse private key: %v", err)
		}

		sshConfig := &ssh.ClientConfig{
			User: "ubuntu",
			Auth: []ssh.AuthMethod{
				ssh.PublicKeys(signer),
			},
			HostKeyCallback: ssh.InsecureIgnoreHostKey(),
		}

		sshConnection, err := ssh.Dial("tcp", fmt.Sprintf("%s:22", vmLinuxPublicIPAddress), sshConfig)
		if err != nil {
			t.Fatalf("Cannot establish SSH connection to vm-linux public IP address: %v", err)
		}

		defer sshConnection.Close()

		sshSession, err := sshConnection.NewSession()
		if err != nil {
			t.Fatalf("Cannot create SSH session to vm-linux public IP address: %v", err)
		}

		defer sshSession.Close()

		err = sshSession.Run(fmt.Sprintf("ping -c 1 8.8.8.8"))
		if err != nil {
			t.Fatalf("Cannot ping 8.8.8.8: %v", err)
		}

		//test DB connection
		databaseHostFqdn := terraform.Output(t, terraformOptions, "database_host_fqdn")
		//databaseHostFqdn := "rc1b-t511or2cucpcdkjs.mdb.yandexcloud.net"

		if databaseHostFqdn == "" {
			t.Fatal("Cannot retrieve the database host fqdn value.")
		}

		//getting first fqdn of mysql cluster
		fmt.Println("This is full string", databaseHostFqdn)
		qqq := strings.SplitAfterN(databaseHostFqdn, "net", 2)
		firstMysqlFqdn1 := qqq[0]
		firstMysqlFqdn := strings.Trim(firstMysqlFqdn1, "[")
		fmt.Println("This is first db fqdn ", firstMysqlFqdn)

		rootCertPool := x509.NewCertPool()

		pem, err := ioutil.ReadFile("./root.crt")
		if err != nil {
			panic(err)
		}

		if ok := rootCertPool.AppendCertsFromPEM(pem); !ok {
			panic("Failed to append PEM.")
		}

		mysql.RegisterTLSConfig("custom", &tls.Config{
			RootCAs: rootCertPool,
		})

		mysqlInfo := fmt.Sprintf("user:password@tcp(%s)/db?tls=custom", firstMysqlFqdn)

		//db, err := sql.Open("mysql", "user:<password>@tcp(rc1b-t511or2cucpcdkjs.mdb.yandexcloud.net)/db?tls=custom")
		db, err := sql.Open("mysql", mysqlInfo)
		if err != nil {
			panic(err)
		}

		// Make sure we clean up properly
		defer db.Close()

		// Run ping to actually test the connection
		logger.Log(t, "Ping the DB with forced SSL")
		if err = db.Ping(); err != nil {
			logger.Logf(t, "Not allowed to ping %s as expected.", firstMysqlFqdn)
		} else {
			//t.Fatalf("Ping %v succeeded against the odds.", firstMysqlFqdn)
			fmt.Print("Ping the DB with forced SSL \n")
		}

	})

	test_structure.RunTestStage(t, "teardown", func() {
		terraformOptions := test_structure.LoadTerraformOptions(t, fixtureFolder)
		terraform.Destroy(t, terraformOptions)
	})
}
