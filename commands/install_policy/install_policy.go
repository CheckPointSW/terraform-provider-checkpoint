package main

import (
	"flag"
	"fmt"
	"github.com/CheckPointSW/terraform-provider-checkpoint/commands"
	"os"
)

type arrayFlags []string

func (i *arrayFlags) String() string {
	return "my string representation"
}

func (i *arrayFlags) Set(value string) error {
	*i = append(*i, value)
	return nil
}

var targets arrayFlags

func main() {
	var policyPackage string
	var access string
	var desktopSecurity string
	var qos string
	var threatPrevention string
	var installOnAllClusterMembersOrFail string
	var prepareOnly string
	var revision string
	var ignoreWarnings string

	flag.StringVar(&policyPackage, "policy-package", "", "The name of the Policy Package to be installed.")
	flag.Var(&targets, "target", "On what targets to execute this command. Targets may be identified by their name, or object unique identifier.")
	flag.StringVar(&access, "access", "", "Set to be true in order to install the Access Control policy. By default, the value is true if Access Control policy is enabled on the input policy package, otherwise false.")
	flag.StringVar(&desktopSecurity, "desktop-security", "", "Set to be true in order to install the Desktop Security policy. By default, the value is true if desktop security policy is enabled on the input policy package, otherwise false.")
	flag.StringVar(&qos, "qos", "", "Set to be true in order to install the QoS policy. By default, the value is true if Quality-of-Service policy is enabled on the input policy package, otherwise false.")
	flag.StringVar(&threatPrevention, "threat-prevention", "", "Set to be true in order to install the Threat Prevention policy. By default, the value is true if Threat Prevention policy is enabled on the input policy package, otherwise false.")
	flag.StringVar(&installOnAllClusterMembersOrFail, "install-on-all-cluster-members-or-fail", "", "Relevant for the gateway clusters. If true, the policy is installed on all the cluster members. If the installation on a cluster member fails, don't install on that cluster.")
	flag.StringVar(&prepareOnly, "prepare-only", "", "If true, prepares the policy for the installation, but doesn't install it on an installation target.")
	flag.StringVar(&revision, "revision", "", "The UID of the revision of the policy to install.")
	flag.StringVar(&ignoreWarnings, "ignore-warnings", "", "Install policy ignoring policy mismatch warnings.")
	flag.Parse()

	apiClient, err := commands.InitClient()
	if err != nil {
		fmt.Println("Install policy error: " + err.Error())
		os.Exit(1)
	}

	payload := map[string]interface{}{
		"policy-package": policyPackage,
	}

	if targets != nil && len(targets) > 0 {
		payload["targets"] = targets
	}

	if len(access) > 0 {
		payload["access"] = access
	}

	if len(desktopSecurity) > 0 {
		payload["desktop-security"] = desktopSecurity
	}

	if len(qos) > 0 {
		payload["qos"] = qos
	}

	if len(threatPrevention) > 0 {
		payload["threat-prevention"] = threatPrevention
	}

	if len(installOnAllClusterMembersOrFail) > 0 {
		payload["install-on-all-cluster-members-or-fail"] = installOnAllClusterMembersOrFail
	}

	if len(prepareOnly) > 0 {
		payload["prepare-only"] = prepareOnly
	}

	if len(revision) > 0 {
		payload["revision"] = revision
	}

	if len(ignoreWarnings) > 0 {
		payload["ignore-warnings"] = ignoreWarnings
	}

	installPolicyRes, err := apiClient.ApiCall("install-policy", payload, apiClient.GetSessionID(), true, apiClient.IsProxyUsed())
	if err != nil {
		fmt.Println("Install policy error: " + err.Error())
		os.Exit(1)
	}

	taskId := commands.ResolveTaskId(installPolicyRes.GetData())

	if !installPolicyRes.Success {
		errMsg := fmt.Sprintf("Install policy failed: %s", installPolicyRes.ErrorMsg)
		if taskId != nil {
			errMsg += fmt.Sprintf("\nTask id: %s", taskId)
		}
		fmt.Println(errMsg)
		os.Exit(1)
	}

	fmt.Println(fmt.Sprintf("Policy installed successfully. task-id [%s]", taskId))
}
