// +build !ignore_autogenerated

// Generated for the devfile generator

// Code generated by helpgen. DO NOT EDIT.

package getters

import (
	"sigs.k8s.io/controller-tools/pkg/markers"
)

func (Generator) Help() *markers.DefinitionHelp {
	return &markers.DefinitionHelp{
		Category: "",
		DetailedHelp: markers.DetailedHelp{
			Summary: "generates getter methods that are used to return values for the boolean pointer fields. ",
			Details: "The pointer receiver is determined from the `devfile:getter:generate` annotated type.  The method will return the value of the field if it's been set, otherwise it will return the default value specified by the devfile:default:value annotation.",
		},
		FieldHelp: map[string]markers.DetailedHelp{},
	}
}
