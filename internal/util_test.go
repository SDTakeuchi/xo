package internal_test

import (
	"testing"

	"github.com/yyoshiki41/xo/internal"
)

func Test_ColumnEnumTypeName(t *testing.T) {
	tests := []struct {
		desc   string
		table  string
		column string
		typ    string
	}{
		// the table name is singularized
		{
			desc:   "singularizes a regular plural table name",
			table:  "users",
			column: "status",
			typ:    "UserStatus",
		},
		{
			desc:   "singularizes only the last word of the table name",
			table:  "card_operation_request_events",
			column: "operated_by",
			typ:    "CardOperationRequestEventOperatedBy",
		},
		{
			desc:   "singularizes an irregular plural table name (people)",
			table:  "people",
			column: "gender",
			typ:    "PersonGender",
		},
		{
			desc:   "singularizes an irregular plural table name (children)",
			table:  "children",
			column: "grade",
			typ:    "ChildGrade",
		},
		{
			desc:   "singularizes a table name ending with -ies",
			table:  "categories",
			column: "kind",
			typ:    "CategoryKind",
		},
		{
			desc:   "singularizes a table name ending with -uses",
			table:  "statuses",
			column: "level",
			typ:    "StatusLevel",
		},

		// the column name is used as is, even when singularizing would change it
		{
			desc:   "does not singularize via (to vium)",
			table:  "card_operation_request_events",
			column: "operated_via",
			typ:    "CardOperationRequestEventOperatedVia",
		},
		{
			desc:   "does not singularize data (to datum)",
			table:  "users",
			column: "meta_data",
			typ:    "UserMetaData",
		},
		{
			desc:   "does not singularize criteria (to criterium)",
			table:  "users",
			column: "search_criteria",
			typ:    "UserSearchCriteria",
		},
		{
			desc:   "does not singularize quota (to quotum)",
			table:  "users",
			column: "storage_quota",
			typ:    "UserStorageQuota",
		},
		{
			desc:   "does not singularize analysis (to analysi)",
			table:  "users",
			column: "risk_analysis",
			typ:    "UserRiskAnalysis",
		},
		{
			desc:   "does not singularize basis (to basi)",
			table:  "users",
			column: "billing_basis",
			typ:    "UserBillingBasis",
		},
		{
			desc:   "does not singularize axis (to axi)",
			table:  "charts",
			column: "sort_axis",
			typ:    "ChartSortAxis",
		},
		{
			desc:   "does not singularize people (to person)",
			table:  "events",
			column: "target_people",
			typ:    "EventTargetPeople",
		},
		{
			desc:   "does not singularize children (to child)",
			table:  "nodes",
			column: "has_children",
			typ:    "NodeHasChildren",
		},
		{
			desc:   "does not singularize men (to man)",
			table:  "events",
			column: "allowed_men",
			typ:    "EventAllowedMen",
		},
		{
			desc:   "does not singularize feet (to foot)",
			table:  "rooms",
			column: "size_feet",
			typ:    "RoomSizeFeet",
		},
		{
			desc:   "does not singularize mice (to mouse)",
			table:  "experiments",
			column: "lab_mice",
			typ:    "ExperimentLabMice",
		},
		{
			desc:   "does not singularize indices (to index)",
			table:  "reports",
			column: "market_indices",
			typ:    "ReportMarketIndices",
		},
		{
			desc:   "does not singularize matrices (to matrix)",
			table:  "models",
			column: "weight_matrices",
			typ:    "ModelWeightMatrices",
		},
		{
			desc:   "does not singularize alumni (to alumnus)",
			table:  "schools",
			column: "notify_alumni",
			typ:    "SchoolNotifyAlumni",
		},
		{
			desc:   "does not singularize leaves (to leafe)",
			table:  "users",
			column: "paid_leaves",
			typ:    "UserPaidLeaves",
		},
		{
			desc:   "does not singularize a regular plural (roles)",
			table:  "users",
			column: "roles",
			typ:    "UserRoles",
		},
		{
			desc:   "does not singularize a plural ending with -ies (categories)",
			table:  "users",
			column: "categories",
			typ:    "UserCategories",
		},
		{
			desc:   "does not singularize a plural ending with -uses (statuses)",
			table:  "users",
			column: "statuses",
			typ:    "UserStatuses",
		},
		{
			desc:   "does not singularize is (to i)",
			table:  "users",
			column: "verified_is",
			typ:    "UserVerifiedIs",
		},
		{
			desc:   "does not singularize has (to ha)",
			table:  "users",
			column: "plan_has",
			typ:    "UserPlanHas",
		},
		{
			desc:   "does not singularize os (to o)",
			table:  "devices",
			column: "target_os",
			typ:    "DeviceTargetOs",
		},
		{
			desc:   "does not singularize gps (to gp)",
			table:  "devices",
			column: "location_gps",
			typ:    "DeviceLocationGps",
		},
	}
	for _, test := range tests {
		t.Run(test.desc, func(t *testing.T) {
			if typ := internal.ColumnEnumTypeName(test.table, test.column); typ != test.typ {
				t.Errorf("expected type %q, got: %q", test.typ, typ)
			}
		})
	}
}
