package middlewares

type Policy struct {
	Resource string
	Action   string
}

var (
	ViewOpportunityMarketplacePolicy Policy = Policy{
		Resource: "opportunity_marketplace",
		Action:   "view_opportunities",
	}
)
