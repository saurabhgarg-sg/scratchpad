package texttemplates

import (
	"fmt"
	"os"
	"text/template"

	"github.com/aporeto-inc/benchmark-suite/catapult/common"
)

type (
	// A flow stores the network flows defined for each node.
	flow struct {
		To       string
		Action   string
		DestPort int
		Protocol int
	}

	// A pu stores the attributes for a given processing unit.
	pu struct {
		Name      string
		Type      string
		OpStatus  string
		EnfStatus string
		MetaData  []string
	}

	// A node holds the parameters for simulator instance plan file.
	node struct {
		Id     string
		Type   string
		BaseIP string
		Pu     pu
		//AppPrefix string
		Flows []flow
	}
)

const templ = `plan:
   nodes:{{ range . }}
   - ID: {{ .Id }}
     type: {{ .Type }}
     IP: {{ .BaseIP }}
     processingUnit:
       name: {{ .Pu.Name }}
       type: {{ .Pu.Type }}
       operationalStatus: {{ .Pu.OpStatus }}
       enforcementStatus: {{ .Pu.EnfStatus }}
       metadata:
       - "@sys:image={{ index .Pu.MetaData 0 }}"
       - "@usr:app={{ index .Pu.MetaData 1 }}"
     edges:
       flows:
       - to: {{ (index .Flows 0).To }}
         report:
           action: {{ (index .Flows 0).Action }}
           destinationPort: {{ (index .Flows 0).DestPort }}
           protocol: {{ (index .Flows 0).Protocol }}
       - to: blog
         report:
           action: Accept
           destinationPort: 443
           protocol: 6
{{ end }}`

func populateNodes(nodeCount int, appPrefix string) []node {
	nodes := make([]node, nodeCount)

	for i := 0; i < nodeCount; i++ {
		name := fmt.Sprintf("%s-%03d", appPrefix, i+1)
		meta := []string{name + "-image", name + "-app"}
		nodes[i].Id = name
		nodes[i].BaseIP = fmt.Sprintf("10.%d.%d.%d", i, i+1, i+1)
		nodes[i].Type = "processingunit"
		nodes[i].Pu = pu{
			Name: name, Type: "Docker", OpStatus: "Running",
			EnfStatus: "Protected",
			MetaData:  meta,
		}

		// Design the required flows for each PU.
		nodes[i].Flows = make([]flow, 3)
		for j := 0; j < 3; j++ {
			nodes[i].Flows[j] = flow{To: "app", Action: "Accept", DestPort: 443, Protocol: 6}
		}
	}
	return nodes
}

func generate() {
	// Create the plan file for simulator execution.
	planFile, err := os.Create("plan.yaml")
	common.CheckError("copyPlanFile: Handle to destination plan file", err)
	defer planFile.Close()

	// Fill the details needed for each node in the Plan file.
	nodes := populateNodes(2, "apo")

	// Execute the template to generate the plan file.
	t := template.Must(template.New("node template").Parse(templ))
	err = t.Execute(planFile, nodes)
	common.CheckError("copyPlanFile: Handle to destination plan file", err)
}
