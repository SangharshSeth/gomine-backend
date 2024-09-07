package types

type Problem struct {
	ProblemID   string   `dynamodbav:"problem_id"`
	Description string   `dynamodbav:"description"`
	Difficulty  string   `dynamodbav:"difficulty"`
	Progress    string   `dynamodbav:"progress"`
	Tags        []string `dynamodbav:"tags"`
}
