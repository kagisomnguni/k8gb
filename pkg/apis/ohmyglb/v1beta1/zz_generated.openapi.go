// +build !ignore_autogenerated

// This file was autogenerated by openapi-gen. Do not edit it manually!

package v1beta1

import (
	spec "github.com/go-openapi/spec"
	common "k8s.io/kube-openapi/pkg/common"
)

func GetOpenAPIDefinitions(ref common.ReferenceCallback) map[string]common.OpenAPIDefinition {
	return map[string]common.OpenAPIDefinition{
		"./pkg/apis/ohmyglb/v1beta1.Gslb":               schema_pkg_apis_ohmyglb_v1beta1_Gslb(ref),
		"./pkg/apis/ohmyglb/v1beta1.GslbResolver":       schema_pkg_apis_ohmyglb_v1beta1_GslbResolver(ref),
		"./pkg/apis/ohmyglb/v1beta1.GslbResolverSpec":   schema_pkg_apis_ohmyglb_v1beta1_GslbResolverSpec(ref),
		"./pkg/apis/ohmyglb/v1beta1.GslbResolverStatus": schema_pkg_apis_ohmyglb_v1beta1_GslbResolverStatus(ref),
		"./pkg/apis/ohmyglb/v1beta1.GslbSpec":           schema_pkg_apis_ohmyglb_v1beta1_GslbSpec(ref),
		"./pkg/apis/ohmyglb/v1beta1.GslbStatus":         schema_pkg_apis_ohmyglb_v1beta1_GslbStatus(ref),
	}
}

func schema_pkg_apis_ohmyglb_v1beta1_Gslb(ref common.ReferenceCallback) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Description: "Gslb is the Schema for the gslbs API",
				Type:        []string{"object"},
				Properties: map[string]spec.Schema{
					"kind": {
						SchemaProps: spec.SchemaProps{
							Description: "Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#types-kinds",
							Type:        []string{"string"},
							Format:      "",
						},
					},
					"apiVersion": {
						SchemaProps: spec.SchemaProps{
							Description: "APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#resources",
							Type:        []string{"string"},
							Format:      "",
						},
					},
					"metadata": {
						SchemaProps: spec.SchemaProps{
							Ref: ref("k8s.io/apimachinery/pkg/apis/meta/v1.ObjectMeta"),
						},
					},
					"spec": {
						SchemaProps: spec.SchemaProps{
							Ref: ref("./pkg/apis/ohmyglb/v1beta1.GslbSpec"),
						},
					},
					"status": {
						SchemaProps: spec.SchemaProps{
							Ref: ref("./pkg/apis/ohmyglb/v1beta1.GslbStatus"),
						},
					},
				},
			},
		},
		Dependencies: []string{
			"./pkg/apis/ohmyglb/v1beta1.GslbSpec", "./pkg/apis/ohmyglb/v1beta1.GslbStatus", "k8s.io/apimachinery/pkg/apis/meta/v1.ObjectMeta"},
	}
}

func schema_pkg_apis_ohmyglb_v1beta1_GslbResolver(ref common.ReferenceCallback) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Description: "GslbResolver is the Schema for the gslbresolvers API",
				Type:        []string{"object"},
				Properties: map[string]spec.Schema{
					"kind": {
						SchemaProps: spec.SchemaProps{
							Description: "Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#types-kinds",
							Type:        []string{"string"},
							Format:      "",
						},
					},
					"apiVersion": {
						SchemaProps: spec.SchemaProps{
							Description: "APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#resources",
							Type:        []string{"string"},
							Format:      "",
						},
					},
					"metadata": {
						SchemaProps: spec.SchemaProps{
							Ref: ref("k8s.io/apimachinery/pkg/apis/meta/v1.ObjectMeta"),
						},
					},
					"spec": {
						SchemaProps: spec.SchemaProps{
							Ref: ref("./pkg/apis/ohmyglb/v1beta1.GslbResolverSpec"),
						},
					},
					"status": {
						SchemaProps: spec.SchemaProps{
							Ref: ref("./pkg/apis/ohmyglb/v1beta1.GslbResolverStatus"),
						},
					},
				},
			},
		},
		Dependencies: []string{
			"./pkg/apis/ohmyglb/v1beta1.GslbResolverSpec", "./pkg/apis/ohmyglb/v1beta1.GslbResolverStatus", "k8s.io/apimachinery/pkg/apis/meta/v1.ObjectMeta"},
	}
}

func schema_pkg_apis_ohmyglb_v1beta1_GslbResolverSpec(ref common.ReferenceCallback) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Description: "GslbResolverSpec defines the desired state of GslbResolver",
				Type:        []string{"object"},
				Properties: map[string]spec.Schema{
					"size": {
						SchemaProps: spec.SchemaProps{
							Description: "INSERT ADDITIONAL SPEC FIELDS - desired state of cluster Important: Run \"operator-sdk generate k8s\" to regenerate code after modifying this file Add custom validation using kubebuilder tags: https://book-v1.book.kubebuilder.io/beyond_basics/generating_crd.html",
							Type:        []string{"integer"},
							Format:      "int32",
						},
					},
				},
				Required: []string{"size"},
			},
		},
	}
}

func schema_pkg_apis_ohmyglb_v1beta1_GslbResolverStatus(ref common.ReferenceCallback) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Description: "GslbResolverStatus defines the observed state of GslbResolver",
				Type:        []string{"object"},
				Properties: map[string]spec.Schema{
					"podNames": {
						VendorExtensible: spec.VendorExtensible{
							Extensions: spec.Extensions{
								"x-kubernetes-list-type": "set",
							},
						},
						SchemaProps: spec.SchemaProps{
							Description: "INSERT ADDITIONAL STATUS FIELD - define observed state of cluster Important: Run \"operator-sdk generate k8s\" to regenerate code after modifying this file Add custom validation using kubebuilder tags: https://book-v1.book.kubebuilder.io/beyond_basics/generating_crd.html",
							Type:        []string{"array"},
							Items: &spec.SchemaOrArray{
								Schema: &spec.Schema{
									SchemaProps: spec.SchemaProps{
										Type:   []string{"string"},
										Format: "",
									},
								},
							},
						},
					},
				},
				Required: []string{"podNames"},
			},
		},
	}
}

func schema_pkg_apis_ohmyglb_v1beta1_GslbSpec(ref common.ReferenceCallback) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Description: "GslbSpec defines the desired state of Gslb",
				Type:        []string{"object"},
				Properties: map[string]spec.Schema{
					"ingress": {
						SchemaProps: spec.SchemaProps{
							Description: "INSERT ADDITIONAL SPEC FIELDS - desired state of cluster Important: Run \"operator-sdk generate k8s\" to regenerate code after modifying this file Add custom validation using kubebuilder tags: https://book-v1.book.kubebuilder.io/beyond_basics/generating_crd.html",
							Ref:         ref("k8s.io/api/extensions/v1beta1.IngressSpec"),
						},
					},
					"strategy": {
						SchemaProps: spec.SchemaProps{
							Type:   []string{"string"},
							Format: "",
						},
					},
				},
				Required: []string{"ingress", "strategy"},
			},
		},
		Dependencies: []string{
			"k8s.io/api/extensions/v1beta1.IngressSpec"},
	}
}

func schema_pkg_apis_ohmyglb_v1beta1_GslbStatus(ref common.ReferenceCallback) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Description: "GslbStatus defines the observed state of Gslb",
				Type:        []string{"object"},
				Properties: map[string]spec.Schema{
					"managedHosts": {
						VendorExtensible: spec.VendorExtensible{
							Extensions: spec.Extensions{
								"x-kubernetes-list-type": "set",
							},
						},
						SchemaProps: spec.SchemaProps{
							Description: "INSERT ADDITIONAL STATUS FIELD - define observed state of cluster Important: Run \"operator-sdk generate k8s\" to regenerate code after modifying this file Add custom validation using kubebuilder tags: https://book-v1.book.kubebuilder.io/beyond_basics/generating_crd.html",
							Type:        []string{"array"},
							Items: &spec.SchemaOrArray{
								Schema: &spec.Schema{
									SchemaProps: spec.SchemaProps{
										Type:   []string{"string"},
										Format: "",
									},
								},
							},
						},
					},
					"serviceHealth": {
						SchemaProps: spec.SchemaProps{
							Type: []string{"object"},
							AdditionalProperties: &spec.SchemaOrBool{
								Allows: true,
								Schema: &spec.Schema{
									SchemaProps: spec.SchemaProps{
										Type:   []string{"string"},
										Format: "",
									},
								},
							},
						},
					},
				},
				Required: []string{"managedHosts", "serviceHealth"},
			},
		},
	}
}