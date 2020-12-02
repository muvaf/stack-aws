/*
Copyright 2020 The Crossplane Authors.

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

package topic

import (
	"context"

	svcsdk "github.com/aws/aws-sdk-go/service/sns"
	"github.com/pkg/errors"
	"k8s.io/client-go/util/workqueue"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/controller"

	xpv1 "github.com/crossplane/crossplane-runtime/apis/common/v1"
	"github.com/crossplane/crossplane-runtime/pkg/event"
	"github.com/crossplane/crossplane-runtime/pkg/logging"
	"github.com/crossplane/crossplane-runtime/pkg/meta"
	"github.com/crossplane/crossplane-runtime/pkg/ratelimiter"
	"github.com/crossplane/crossplane-runtime/pkg/reconciler/managed"
	"github.com/crossplane/crossplane-runtime/pkg/resource"

	svcapitypes "github.com/crossplane/provider-aws/apis/sns/v1alpha1"
	aws "github.com/crossplane/provider-aws/pkg/clients"
)

const (
	errUpdate = "failed to update Topic"
)

// SetupTopic adds a controller that reconciles Topic.
func SetupTopic(mgr ctrl.Manager, l logging.Logger, rl workqueue.RateLimiter) error {
	name := managed.ControllerName(svcapitypes.TopicGroupKind)
	return ctrl.NewControllerManagedBy(mgr).
		Named(name).
		WithOptions(controller.Options{
			RateLimiter: ratelimiter.NewDefaultManagedRateLimiter(rl),
		}).
		For(&svcapitypes.Topic{}).
		Complete(managed.NewReconciler(mgr,
			resource.ManagedKind(svcapitypes.TopicGroupVersionKind),
			managed.WithExternalConnecter(&connector{kube: mgr.GetClient()}),
			managed.WithInitializers(managed.NewDefaultProviderConfig(mgr.GetClient())),
			managed.WithLogger(l.WithValues("controller", name)),
			managed.WithRecorder(event.NewAPIRecorder(mgr.GetEventRecorderFor(name)))))
}

func (*external) preObserve(context.Context, *svcapitypes.Topic) error {
	return nil
}
func (*external) postObserve(_ context.Context, cr *svcapitypes.Topic, _ *svcsdk.GetTopicAttributesOutput, obs managed.ExternalObservation, err error) (managed.ExternalObservation, error) {
	if err != nil {
		return managed.ExternalObservation{}, err
	}
	cr.SetConditions(xpv1.Available())
	return obs, nil
}

func (*external) preCreate(context.Context, *svcapitypes.Topic) error {
	return nil
}

func (*external) postCreate(_ context.Context, cr *svcapitypes.Topic, resp *svcsdk.CreateTopicOutput, cre managed.ExternalCreation, err error) (managed.ExternalCreation, error) {
	if err != nil {
		return managed.ExternalCreation{}, err
	}
	meta.SetExternalName(cr, aws.StringValue(resp.TopicArn))
	cre.ExternalNameAssigned = true
	return cre, nil
}

func (*external) preUpdate(context.Context, *svcapitypes.Topic) error {
	return nil
}

func preGenerateGetTopicAttributesInput(_ *svcapitypes.Topic, obj *svcsdk.GetTopicAttributesInput) *svcsdk.GetTopicAttributesInput {
	return obj
}

func postGenerateGetTopicAttributesInput(cr *svcapitypes.Topic, obj *svcsdk.GetTopicAttributesInput) *svcsdk.GetTopicAttributesInput {
	obj.TopicArn = aws.String(meta.GetExternalName(cr))
	return obj
}

func preGenerateCreateTopicInput(_ *svcapitypes.Topic, obj *svcsdk.CreateTopicInput) *svcsdk.CreateTopicInput {
	return obj
}

func postGenerateCreateTopicInput(_ *svcapitypes.Topic, obj *svcsdk.CreateTopicInput) *svcsdk.CreateTopicInput {
	return obj
}
func preGenerateDeleteTopicInput(_ *svcapitypes.Topic, obj *svcsdk.DeleteTopicInput) *svcsdk.DeleteTopicInput {
	return obj
}

func postGenerateDeleteTopicInput(cr *svcapitypes.Topic, obj *svcsdk.DeleteTopicInput) *svcsdk.DeleteTopicInput {
	obj.TopicArn = aws.String(meta.GetExternalName(cr))
	return obj
}

// NOTE(muvaf): The rest is adopted from old SNSTopic implementation

func (e *external) postUpdate(ctx context.Context, cr *svcapitypes.Topic, _ managed.ExternalUpdate, err error) (managed.ExternalUpdate, error) {
	if err != nil {
		return managed.ExternalUpdate{}, err
	}
	// Fetch Topic Attributes again
	resp, err := e.client.GetTopicAttributesWithContext(ctx, &svcsdk.GetTopicAttributesInput{
		TopicArn: aws.String(meta.GetExternalName(cr)),
	})
	if err != nil {
		return managed.ExternalUpdate{}, errors.Wrap(err, errDescribe)
	}

	// Update Topic Attributes
	attrs := GetChangedAttributes(cr.Spec.ForProvider, resp.Attributes)
	for k := range attrs {
		_, err := e.client.SetTopicAttributesWithContext(ctx, &svcsdk.SetTopicAttributesInput{
			AttributeName:  aws.String(k),
			AttributeValue: attrs[k],
			TopicArn:       aws.String(meta.GetExternalName(cr)),
		})
		if err != nil {
			return managed.ExternalUpdate{}, errors.Wrap(err, errUpdate)
		}
	}
	return managed.ExternalUpdate{}, nil
}

// GetChangedAttributes will return the changed attributes for a topic in AWS side.
//
// This is needed as currently AWS SDK allows to set Attribute Topics one at a time.
// Please see https://docs.aws.amazon.com/sns/latest/api/API_SetTopicAttributes.html
// So we need to compare each topic attribute and call SetTopicAttribute for ones which has
// changed.
func GetChangedAttributes(p svcapitypes.TopicParameters, attrs map[string]*string) map[string]*string {
	topicAttrs := map[string]string{
		string(svcapitypes.TopicDeliveryPolicy): aws.StringValue(p.DeliveryPolicy),
		string(svcapitypes.TopicDisplayName):    aws.StringValue(p.DisplayName),
		string(svcapitypes.TopicKmsMasterKeyID): aws.StringValue(p.KMSMasterKeyID),
		string(svcapitypes.TopicPolicy):         aws.StringValue(p.Policy),
	}
	changed := map[string]*string{}
	for k, v := range attrs {
		if topicAttrs[k] != aws.StringValue(v) {
			changed[k] = aws.String(topicAttrs[k])
		}
	}
	return changed
}

func lateInitialize(spec *svcapitypes.TopicParameters, resp *svcsdk.GetTopicAttributesOutput) {
	spec.DisplayName = aws.LateInitializeStringPtr(spec.DisplayName, resp.Attributes[string(svcapitypes.TopicDeliveryPolicy)])
	spec.DeliveryPolicy = aws.LateInitializeStringPtr(spec.DeliveryPolicy, resp.Attributes[string(svcapitypes.TopicDisplayName)])
	spec.KMSMasterKeyID = aws.LateInitializeStringPtr(spec.KMSMasterKeyID, resp.Attributes[string(svcapitypes.TopicKmsMasterKeyID)])
	spec.Policy = aws.LateInitializeStringPtr(spec.Policy, resp.Attributes[string(svcapitypes.TopicPolicy)])
}

func isUpToDate(cr *svcapitypes.Topic, resp *svcsdk.GetTopicAttributesOutput) bool {
	return aws.StringValue(cr.Spec.ForProvider.DeliveryPolicy) == aws.StringValue(resp.Attributes[string(svcapitypes.TopicDeliveryPolicy)]) &&
		aws.StringValue(cr.Spec.ForProvider.DisplayName) == aws.StringValue(resp.Attributes[string(svcapitypes.TopicDisplayName)]) &&
		aws.StringValue(cr.Spec.ForProvider.KMSMasterKeyID) == aws.StringValue(resp.Attributes[string(svcapitypes.TopicKmsMasterKeyID)]) &&
		aws.StringValue(cr.Spec.ForProvider.Policy) == aws.StringValue(resp.Attributes[string(svcapitypes.TopicPolicy)])
}
