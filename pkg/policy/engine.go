package policy

import "github.com/MDzubair933/cloud-finops-cost-governor/pkg/analyzer"

type RemediationAction string

const (
	ActionNotify   RemediationAction = "NOTIFY"
	ActionSnapshot RemediationAction = "SNAPSHOT_AND_DELETE"
	ActionDownsize RemediationAction = "DOWNSIZE"
)

type PolicyEngine struct{}

func (p *PolicyEngine) Evaluate(waste analyzer.ResourceWaste) RemediationAction {
	if waste.MonthlyWasteUSD > 100.0 {
		return ActionSnapshot
	}
	return ActionNotify
}
