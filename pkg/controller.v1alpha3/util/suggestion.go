package util

import (
	"fmt"

	suggestionsv1alpha3 "github.com/kubeflow/katib/pkg/apis/controller/suggestions/v1alpha3"
	"github.com/kubeflow/katib/pkg/controller.v1alpha3/consts"
)

func GetAlgorithmDeploymentName(s *suggestionsv1alpha3.Suggestion) string {
	return s.Name + "-" + s.Spec.AlgorithmName
}

func GetAlgorithmServiceName(s *suggestionsv1alpha3.Suggestion) string {
	return s.Name + "-" + s.Spec.AlgorithmName
}

func GetAlgorithmEndpoint(s *suggestionsv1alpha3.Suggestion) string {
	serviceName := GetAlgorithmServiceName(s)
	return fmt.Sprintf("%s:%d", serviceName, consts.DefaultSuggestionPort)
}
