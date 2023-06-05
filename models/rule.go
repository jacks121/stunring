package models

type Rule interface {
	Apply(c Collection) []Product
}

type RuleByProductIds struct {
	ProductIds []string
}

type RuleByCategoryIds struct {
	CategoryIds []string
}

type RuleByLatestCreated struct {
	Count int
}
