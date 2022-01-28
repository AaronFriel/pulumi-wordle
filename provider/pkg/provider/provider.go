// Copyright 2016-2020, Pulumi Corporation.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package provider

import (
	"context"
	"fmt"
	"math"
	"time"

	"github.com/pulumi/pulumi/pkg/v3/resource/provider"
	"github.com/pulumi/pulumi/sdk/v3/go/common/resource"
	"github.com/pulumi/pulumi/sdk/v3/go/common/resource/plugin"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	pulumirpc "github.com/pulumi/pulumi/sdk/v3/proto/go"

	pbempty "github.com/golang/protobuf/ptypes/empty"
)

type wordleProvider struct {
	host    *provider.HostClient
	name    string
	version string
}

func makeProvider(host *provider.HostClient, name, version string) (pulumirpc.ResourceProviderServer, error) {
	// Return the new provider
	return &wordleProvider{
		host:    host,
		name:    name,
		version: version,
	}, nil
}

// Call dynamically executes a method in the provider associated with a component resource.
func (k *wordleProvider) Call(ctx context.Context, req *pulumirpc.CallRequest) (*pulumirpc.CallResponse, error) {
	return nil, status.Error(codes.Unimplemented, "Call is not yet implemented")
}

// Construct creates a new component resource.
func (k *wordleProvider) Construct(ctx context.Context, req *pulumirpc.ConstructRequest) (*pulumirpc.ConstructResponse, error) {
	return nil, status.Error(codes.Unimplemented, "Construct is not yet implemented")
}

// CheckConfig validates the configuration for this provider.
func (k *wordleProvider) CheckConfig(ctx context.Context, req *pulumirpc.CheckRequest) (*pulumirpc.CheckResponse, error) {
	return &pulumirpc.CheckResponse{Inputs: req.GetNews()}, nil
}

// DiffConfig diffs the configuration for this provider.
func (k *wordleProvider) DiffConfig(ctx context.Context, req *pulumirpc.DiffRequest) (*pulumirpc.DiffResponse, error) {
	return &pulumirpc.DiffResponse{}, nil
}

// Configure configures the resource provider with "globals" that control its behavior.
func (k *wordleProvider) Configure(_ context.Context, req *pulumirpc.ConfigureRequest) (*pulumirpc.ConfigureResponse, error) {
	return &pulumirpc.ConfigureResponse{}, nil
}

// Invoke dynamically executes a built-in function in the provider.
func (k *wordleProvider) Invoke(_ context.Context, req *pulumirpc.InvokeRequest) (*pulumirpc.InvokeResponse, error) {
	tok := req.GetTok()
	return nil, fmt.Errorf("unknown Invoke token '%s'", tok)
}

// StreamInvoke dynamically executes a built-in function in the provider. The result is streamed
// back as a series of messages.
func (k *wordleProvider) StreamInvoke(req *pulumirpc.InvokeRequest, server pulumirpc.ResourceProvider_StreamInvokeServer) error {
	tok := req.GetTok()
	return fmt.Errorf("unknown StreamInvoke token '%s'", tok)
}

// Check validates that the given property bag is valid for a resource of the given type and returns
// the inputs that should be passed to successive calls to Diff, Create, or Update for this
// resource. As a rule, the provider inputs returned by a call to Check should preserve the original
// representation of the properties as present in the program inputs. Though this rule is not
// required for correctness, violations thereof can negatively impact the end-user experience, as
// the provider inputs are using for detecting and rendering diffs.
func (k *wordleProvider) Check(ctx context.Context, req *pulumirpc.CheckRequest) (*pulumirpc.CheckResponse, error) {
	urn := resource.URN(req.GetUrn())
	ty := urn.Type()
	if ty != "wordle:index:Wordle" {
		return nil, fmt.Errorf("unknown resource type '%s'", ty)
	}
	return &pulumirpc.CheckResponse{Inputs: req.News, Failures: nil}, nil
}

// Diff checks what impacts a hypothetical update will have on the resource's properties.
func (k *wordleProvider) Diff(ctx context.Context, req *pulumirpc.DiffRequest) (*pulumirpc.DiffResponse, error) {
	urn := resource.URN(req.GetUrn())
	ty := urn.Type()
	if ty != "wordle:index:Wordle" {
		return nil, fmt.Errorf("unknown resource type '%s'", ty)
	}

	olds, err := plugin.UnmarshalProperties(req.GetOlds(), plugin.MarshalOptions{KeepUnknowns: true, SkipNulls: true})
	if err != nil {
		return nil, err
	}

	news, err := plugin.UnmarshalProperties(req.GetNews(), plugin.MarshalOptions{KeepUnknowns: true, SkipNulls: true})
	if err != nil {
		return nil, err
	}

	d := olds.Diff(news)
	changes := pulumirpc.DiffResponse_DIFF_NONE
	if d.Changed("word") {
		changes = pulumirpc.DiffResponse_DIFF_SOME
	}

	return &pulumirpc.DiffResponse{
		Changes:  changes,
		Replaces: []string{},
	}, nil
}

func wordInList(word string) bool {
	for _, v := range Solutions {
		if word == v {
			return true
		}
	}
	for _, v := range Words {
		if word == v {
			return true
		}
	}

	return false
}

// Create allocates a new instance of the provided resource and returns its unique ID afterwards.
func (k *wordleProvider) Create(ctx context.Context, req *pulumirpc.CreateRequest) (*pulumirpc.CreateResponse, error) {
	urn := resource.URN(req.GetUrn())
	ty := urn.Type()
	if ty != "wordle:index:Wordle" {
		return nil, fmt.Errorf("unknown resource type '%s'", ty)
	}

	inputs, err := plugin.UnmarshalProperties(req.GetProperties(), plugin.MarshalOptions{KeepUnknowns: true, SkipNulls: true})
	if err != nil {
		return nil, err
	}

	if !inputs["word"].IsString() {
		return nil, fmt.Errorf("expected input property 'word' of type 'string' but got '%s'", inputs["word"].TypeString())
	}
	word := inputs["word"].StringValue()

	if !wordInList(word) {
		return nil, fmt.Errorf("expected input property 'word' to be a valid five letter word, got '%s'", word)
	}

	inputState, hasState := inputs["result"]
	state := ""
	if hasState {
		state = inputState.StringValue()
	}

	result := renderResult(word, TodaysSolution())
	state += result + "\n"

	outputs := map[string]interface{}{
		"word":   word,
		"result": state,
	}

	outputProperties, err := plugin.MarshalProperties(
		resource.NewPropertyMapFromMap(outputs),
		plugin.MarshalOptions{KeepUnknowns: true, SkipNulls: true},
	)
	if err != nil {
		return nil, err
	}
	return &pulumirpc.CreateResponse{
		Id:         fmt.Sprintf("wordle-%s", todaysDate()),
		Properties: outputProperties,
	}, nil
}

// Read the current live state associated with a resource.
func (k *wordleProvider) Read(ctx context.Context, req *pulumirpc.ReadRequest) (*pulumirpc.ReadResponse, error) {
	urn := resource.URN(req.GetUrn())
	ty := urn.Type()
	if ty != "wordle:index:Wordle" {
		return nil, fmt.Errorf("unknown resource type '%s'", ty)
	}
	return nil, status.Error(codes.Unimplemented, "Read is not yet implemented for 'wordle:index:Wordle'")
}

// Update updates an existing resource with new values.
func (k *wordleProvider) Update(ctx context.Context, req *pulumirpc.UpdateRequest) (*pulumirpc.UpdateResponse, error) {
	urn := resource.URN(req.GetUrn())
	ty := urn.Type()
	if ty != "wordle:index:Wordle" {
		return nil, fmt.Errorf("unknown resource type '%s'", ty)
	}

	olds, err := plugin.UnmarshalProperties(req.GetOlds(), plugin.MarshalOptions{KeepUnknowns: true, SkipNulls: true})
	if err != nil {
		return nil, err
	}
	news, err := plugin.UnmarshalProperties(req.GetNews(), plugin.MarshalOptions{KeepUnknowns: true, SkipNulls: true})
	if err != nil {
		return nil, err
	}

	if !news["word"].IsString() {
		return nil, fmt.Errorf("expected input property 'word' of type 'string' but got '%s'", news["word"].TypeString())
	}
	word := news["word"].StringValue()
	if !wordInList(word) {
		return nil, fmt.Errorf("expected input property 'word' to be a valid five letter word, got '%s'", word)
	}

	inputState, hasState := olds["result"]
	state := ""
	if hasState {
		state = inputState.StringValue()
	}

	result := renderResult(word, TodaysSolution())
	state += result + "\n"

	outputs := map[string]interface{}{
		"word":   word,
		"result": state,
	}

	outputProperties, err := plugin.MarshalProperties(
		resource.NewPropertyMapFromMap(outputs),
		plugin.MarshalOptions{KeepUnknowns: true, SkipNulls: true},
	)
	if err != nil {
		return nil, err
	}
	return &pulumirpc.UpdateResponse{
		Properties: outputProperties,
	}, nil
}

// Delete tears down an existing resource with the given ID.  If it fails, the resource is assumed
// to still exist.
func (k *wordleProvider) Delete(ctx context.Context, req *pulumirpc.DeleteRequest) (*pbempty.Empty, error) {
	urn := resource.URN(req.GetUrn())
	ty := urn.Type()
	if ty != "wordle:index:Wordle" {
		return nil, fmt.Errorf("unknown resource type '%s'", ty)
	}

	// Note that for our Random resource, we don't have to do anything on Delete.
	return &pbempty.Empty{}, nil
}

// GetPluginInfo returns generic information about this plugin, like its version.
func (k *wordleProvider) GetPluginInfo(context.Context, *pbempty.Empty) (*pulumirpc.PluginInfo, error) {
	return &pulumirpc.PluginInfo{
		Version: k.version,
	}, nil
}

// GetSchema returns the JSON-serialized schema for the provider.
func (k *wordleProvider) GetSchema(ctx context.Context, req *pulumirpc.GetSchemaRequest) (*pulumirpc.GetSchemaResponse, error) {
	return &pulumirpc.GetSchemaResponse{}, nil
}

// Cancel signals the provider to gracefully shut down and abort any ongoing resource operations.
// Operations aborted in this way will return an error (e.g., `Update` and `Create` will either a
// creation error or an initialization error). Since Cancel is advisory and non-blocking, it is up
// to the host to decide how long to wait after Cancel is called before (e.g.)
// hard-closing any gRPC connection.
func (k *wordleProvider) Cancel(context.Context, *pbempty.Empty) (*pbempty.Empty, error) {
	// TODO
	return &pbempty.Empty{}, nil
}

func TodaysSolution() string {
	// How the wordle algorithm calculates the current word index
	epoch := time.Date(2021, 6, 19, 0, 0, 0, 0, time.UTC)
	year, month, day := time.Now().Date()
	today := time.Date(year, month, day, 0, 0, 0, 0, time.UTC)

	days := int(math.Round(today.Sub(epoch).Seconds() / 86400))

	index := days % len(Solutions)

	return Solutions[index]
}

func todaysDate() string {
	return time.Now().Format("2006-01-02")
}

func renderResult(input, solution string) string {
	var result string = ""

	runicSolution := []rune(solution)
CharLoop:
	for idx, char := range input {
		if (runicSolution)[idx] == char {
			result += "🟩"
			continue CharLoop
		}

		for _, solutionChar := range solution {
			if char == solutionChar {
				result += "🟨"
				continue CharLoop
			}
		}

		result += "🟫"
	}

	return result
}
