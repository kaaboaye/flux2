/*
Copyright 2020 The Flux authors

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package main

import (
	sourcev1 "github.com/fluxcd/source-controller/api/v1beta1"
	"github.com/spf13/cobra"
)

var resumeSourceBucketCmd = &cobra.Command{
	Use:   "bucket [name]",
	Short: "Resume a suspended Bucket",
	Long:  `The resume command marks a previously suspended Bucket resource for reconciliation and waits for it to finish.`,
	Example: `  # Resume reconciliation for an existing Bucket
  flux resume source bucket podinfo`,
	RunE: resumeCommand{
		apiType: bucketType,
		object:  &bucketAdapter{&sourcev1.Bucket{}},
	}.run,
}

func init() {
	resumeSourceCmd.AddCommand(resumeSourceBucketCmd)
}

func (obj bucketAdapter) getObservedGeneration() int64 {
	return obj.Bucket.Status.ObservedGeneration
}

func (obj bucketAdapter) setUnsuspended() {
	obj.Bucket.Spec.Suspend = false
}
