// SPDX-FileCopyrightText: 2024 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: Apache-2.0
// Code generated by angryjet. DO NOT EDIT.
// Code transformed by upjet. DO NOT EDIT.

package v1beta2

import (
	"context"
	reference "github.com/crossplane/crossplane-runtime/pkg/reference"
	xpresource "github.com/crossplane/crossplane-runtime/pkg/resource"
	resource "github.com/crossplane/upjet/pkg/resource"
	errors "github.com/pkg/errors"
	client "sigs.k8s.io/controller-runtime/pkg/client"

	// ResolveReferences of this Flow.
	apisresolver "github.com/upbound/provider-gcp/internal/apis"
)

func (mg *Flow) ResolveReferences(ctx context.Context, c client.Reader) error {
	var m xpresource.Managed
	var l xpresource.ManagedList
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error

	if mg.Spec.ForProvider.KnowledgeConnectorSettings != nil {
		{
			m, l, err = apisresolver.GetManagedResource("dialogflowcx.gcp.upbound.io", "v1beta2", "Agent", "AgentList")
			if err != nil {
				return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
			}
			rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
				CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.KnowledgeConnectorSettings.TargetFlow),
				Extract:      resource.ExtractParamPath("start_flow", true),
				Reference:    mg.Spec.ForProvider.KnowledgeConnectorSettings.TargetFlowRef,
				Selector:     mg.Spec.ForProvider.KnowledgeConnectorSettings.TargetFlowSelector,
				To:           reference.To{List: l, Managed: m},
			})
		}
		if err != nil {
			return errors.Wrap(err, "mg.Spec.ForProvider.KnowledgeConnectorSettings.TargetFlow")
		}
		mg.Spec.ForProvider.KnowledgeConnectorSettings.TargetFlow = reference.ToPtrValue(rsp.ResolvedValue)
		mg.Spec.ForProvider.KnowledgeConnectorSettings.TargetFlowRef = rsp.ResolvedReference

	}
	if mg.Spec.ForProvider.KnowledgeConnectorSettings != nil {
		if mg.Spec.ForProvider.KnowledgeConnectorSettings.TriggerFulfillment != nil {
			{
				m, l, err = apisresolver.GetManagedResource("dialogflowcx.gcp.upbound.io", "v1beta2", "Webhook", "WebhookList")
				if err != nil {
					return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
				}
				rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
					CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.KnowledgeConnectorSettings.TriggerFulfillment.Webhook),
					Extract:      resource.ExtractResourceID(),
					Reference:    mg.Spec.ForProvider.KnowledgeConnectorSettings.TriggerFulfillment.WebhookRef,
					Selector:     mg.Spec.ForProvider.KnowledgeConnectorSettings.TriggerFulfillment.WebhookSelector,
					To:           reference.To{List: l, Managed: m},
				})
			}
			if err != nil {
				return errors.Wrap(err, "mg.Spec.ForProvider.KnowledgeConnectorSettings.TriggerFulfillment.Webhook")
			}
			mg.Spec.ForProvider.KnowledgeConnectorSettings.TriggerFulfillment.Webhook = reference.ToPtrValue(rsp.ResolvedValue)
			mg.Spec.ForProvider.KnowledgeConnectorSettings.TriggerFulfillment.WebhookRef = rsp.ResolvedReference

		}
	}
	{
		m, l, err = apisresolver.GetManagedResource("dialogflowcx.gcp.upbound.io", "v1beta2", "Agent", "AgentList")
		if err != nil {
			return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
		}
		rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
			CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.Parent),
			Extract:      resource.ExtractResourceID(),
			Reference:    mg.Spec.ForProvider.ParentRef,
			Selector:     mg.Spec.ForProvider.ParentSelector,
			To:           reference.To{List: l, Managed: m},
		})
	}
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.Parent")
	}
	mg.Spec.ForProvider.Parent = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.ParentRef = rsp.ResolvedReference

	for i3 := 0; i3 < len(mg.Spec.ForProvider.TransitionRoutes); i3++ {
		{
			m, l, err = apisresolver.GetManagedResource("dialogflowcx.gcp.upbound.io", "v1beta2", "Agent", "AgentList")
			if err != nil {
				return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
			}
			rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
				CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.TransitionRoutes[i3].TargetFlow),
				Extract:      resource.ExtractParamPath("start_flow", true),
				Reference:    mg.Spec.ForProvider.TransitionRoutes[i3].TargetFlowRef,
				Selector:     mg.Spec.ForProvider.TransitionRoutes[i3].TargetFlowSelector,
				To:           reference.To{List: l, Managed: m},
			})
		}
		if err != nil {
			return errors.Wrap(err, "mg.Spec.ForProvider.TransitionRoutes[i3].TargetFlow")
		}
		mg.Spec.ForProvider.TransitionRoutes[i3].TargetFlow = reference.ToPtrValue(rsp.ResolvedValue)
		mg.Spec.ForProvider.TransitionRoutes[i3].TargetFlowRef = rsp.ResolvedReference

	}
	if mg.Spec.InitProvider.KnowledgeConnectorSettings != nil {
		{
			m, l, err = apisresolver.GetManagedResource("dialogflowcx.gcp.upbound.io", "v1beta2", "Agent", "AgentList")
			if err != nil {
				return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
			}
			rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
				CurrentValue: reference.FromPtrValue(mg.Spec.InitProvider.KnowledgeConnectorSettings.TargetFlow),
				Extract:      resource.ExtractParamPath("start_flow", true),
				Reference:    mg.Spec.InitProvider.KnowledgeConnectorSettings.TargetFlowRef,
				Selector:     mg.Spec.InitProvider.KnowledgeConnectorSettings.TargetFlowSelector,
				To:           reference.To{List: l, Managed: m},
			})
		}
		if err != nil {
			return errors.Wrap(err, "mg.Spec.InitProvider.KnowledgeConnectorSettings.TargetFlow")
		}
		mg.Spec.InitProvider.KnowledgeConnectorSettings.TargetFlow = reference.ToPtrValue(rsp.ResolvedValue)
		mg.Spec.InitProvider.KnowledgeConnectorSettings.TargetFlowRef = rsp.ResolvedReference

	}
	if mg.Spec.InitProvider.KnowledgeConnectorSettings != nil {
		if mg.Spec.InitProvider.KnowledgeConnectorSettings.TriggerFulfillment != nil {
			{
				m, l, err = apisresolver.GetManagedResource("dialogflowcx.gcp.upbound.io", "v1beta2", "Webhook", "WebhookList")
				if err != nil {
					return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
				}
				rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
					CurrentValue: reference.FromPtrValue(mg.Spec.InitProvider.KnowledgeConnectorSettings.TriggerFulfillment.Webhook),
					Extract:      resource.ExtractResourceID(),
					Reference:    mg.Spec.InitProvider.KnowledgeConnectorSettings.TriggerFulfillment.WebhookRef,
					Selector:     mg.Spec.InitProvider.KnowledgeConnectorSettings.TriggerFulfillment.WebhookSelector,
					To:           reference.To{List: l, Managed: m},
				})
			}
			if err != nil {
				return errors.Wrap(err, "mg.Spec.InitProvider.KnowledgeConnectorSettings.TriggerFulfillment.Webhook")
			}
			mg.Spec.InitProvider.KnowledgeConnectorSettings.TriggerFulfillment.Webhook = reference.ToPtrValue(rsp.ResolvedValue)
			mg.Spec.InitProvider.KnowledgeConnectorSettings.TriggerFulfillment.WebhookRef = rsp.ResolvedReference

		}
	}
	{
		m, l, err = apisresolver.GetManagedResource("dialogflowcx.gcp.upbound.io", "v1beta2", "Agent", "AgentList")
		if err != nil {
			return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
		}
		rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
			CurrentValue: reference.FromPtrValue(mg.Spec.InitProvider.Parent),
			Extract:      resource.ExtractResourceID(),
			Reference:    mg.Spec.InitProvider.ParentRef,
			Selector:     mg.Spec.InitProvider.ParentSelector,
			To:           reference.To{List: l, Managed: m},
		})
	}
	if err != nil {
		return errors.Wrap(err, "mg.Spec.InitProvider.Parent")
	}
	mg.Spec.InitProvider.Parent = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.InitProvider.ParentRef = rsp.ResolvedReference

	for i3 := 0; i3 < len(mg.Spec.InitProvider.TransitionRoutes); i3++ {
		{
			m, l, err = apisresolver.GetManagedResource("dialogflowcx.gcp.upbound.io", "v1beta2", "Agent", "AgentList")
			if err != nil {
				return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
			}
			rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
				CurrentValue: reference.FromPtrValue(mg.Spec.InitProvider.TransitionRoutes[i3].TargetFlow),
				Extract:      resource.ExtractParamPath("start_flow", true),
				Reference:    mg.Spec.InitProvider.TransitionRoutes[i3].TargetFlowRef,
				Selector:     mg.Spec.InitProvider.TransitionRoutes[i3].TargetFlowSelector,
				To:           reference.To{List: l, Managed: m},
			})
		}
		if err != nil {
			return errors.Wrap(err, "mg.Spec.InitProvider.TransitionRoutes[i3].TargetFlow")
		}
		mg.Spec.InitProvider.TransitionRoutes[i3].TargetFlow = reference.ToPtrValue(rsp.ResolvedValue)
		mg.Spec.InitProvider.TransitionRoutes[i3].TargetFlowRef = rsp.ResolvedReference

	}

	return nil
}

// ResolveReferences of this Page.
func (mg *Page) ResolveReferences(ctx context.Context, c client.Reader) error {
	var m xpresource.Managed
	var l xpresource.ManagedList
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error

	if mg.Spec.ForProvider.Form != nil {
		for i4 := 0; i4 < len(mg.Spec.ForProvider.Form.Parameters); i4++ {
			if mg.Spec.ForProvider.Form.Parameters[i4].FillBehavior != nil {
				for i6 := 0; i6 < len(mg.Spec.ForProvider.Form.Parameters[i4].FillBehavior.RepromptEventHandlers); i6++ {
					{
						m, l, err = apisresolver.GetManagedResource("dialogflowcx.gcp.upbound.io", "v1beta2", "Agent", "AgentList")
						if err != nil {
							return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
						}
						rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
							CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.Form.Parameters[i4].FillBehavior.RepromptEventHandlers[i6].TargetFlow),
							Extract:      resource.ExtractParamPath("start_flow", true),
							Reference:    mg.Spec.ForProvider.Form.Parameters[i4].FillBehavior.RepromptEventHandlers[i6].TargetFlowRef,
							Selector:     mg.Spec.ForProvider.Form.Parameters[i4].FillBehavior.RepromptEventHandlers[i6].TargetFlowSelector,
							To:           reference.To{List: l, Managed: m},
						})
					}
					if err != nil {
						return errors.Wrap(err, "mg.Spec.ForProvider.Form.Parameters[i4].FillBehavior.RepromptEventHandlers[i6].TargetFlow")
					}
					mg.Spec.ForProvider.Form.Parameters[i4].FillBehavior.RepromptEventHandlers[i6].TargetFlow = reference.ToPtrValue(rsp.ResolvedValue)
					mg.Spec.ForProvider.Form.Parameters[i4].FillBehavior.RepromptEventHandlers[i6].TargetFlowRef = rsp.ResolvedReference

				}
			}
		}
	}
	if mg.Spec.ForProvider.Form != nil {
		for i4 := 0; i4 < len(mg.Spec.ForProvider.Form.Parameters); i4++ {
			if mg.Spec.ForProvider.Form.Parameters[i4].FillBehavior != nil {
				for i6 := 0; i6 < len(mg.Spec.ForProvider.Form.Parameters[i4].FillBehavior.RepromptEventHandlers); i6++ {
					{
						m, l, err = apisresolver.GetManagedResource("dialogflowcx.gcp.upbound.io", "v1beta2", "Page", "PageList")
						if err != nil {
							return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
						}
						rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
							CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.Form.Parameters[i4].FillBehavior.RepromptEventHandlers[i6].TargetPage),
							Extract:      resource.ExtractResourceID(),
							Reference:    mg.Spec.ForProvider.Form.Parameters[i4].FillBehavior.RepromptEventHandlers[i6].TargetPageRef,
							Selector:     mg.Spec.ForProvider.Form.Parameters[i4].FillBehavior.RepromptEventHandlers[i6].TargetPageSelector,
							To:           reference.To{List: l, Managed: m},
						})
					}
					if err != nil {
						return errors.Wrap(err, "mg.Spec.ForProvider.Form.Parameters[i4].FillBehavior.RepromptEventHandlers[i6].TargetPage")
					}
					mg.Spec.ForProvider.Form.Parameters[i4].FillBehavior.RepromptEventHandlers[i6].TargetPage = reference.ToPtrValue(rsp.ResolvedValue)
					mg.Spec.ForProvider.Form.Parameters[i4].FillBehavior.RepromptEventHandlers[i6].TargetPageRef = rsp.ResolvedReference

				}
			}
		}
	}
	if mg.Spec.ForProvider.Form != nil {
		for i4 := 0; i4 < len(mg.Spec.ForProvider.Form.Parameters); i4++ {
			if mg.Spec.ForProvider.Form.Parameters[i4].FillBehavior != nil {
				for i6 := 0; i6 < len(mg.Spec.ForProvider.Form.Parameters[i4].FillBehavior.RepromptEventHandlers); i6++ {
					if mg.Spec.ForProvider.Form.Parameters[i4].FillBehavior.RepromptEventHandlers[i6].TriggerFulfillment != nil {
						{
							m, l, err = apisresolver.GetManagedResource("dialogflowcx.gcp.upbound.io", "v1beta2", "Webhook", "WebhookList")
							if err != nil {
								return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
							}
							rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
								CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.Form.Parameters[i4].FillBehavior.RepromptEventHandlers[i6].TriggerFulfillment.Webhook),
								Extract:      resource.ExtractResourceID(),
								Reference:    mg.Spec.ForProvider.Form.Parameters[i4].FillBehavior.RepromptEventHandlers[i6].TriggerFulfillment.WebhookRef,
								Selector:     mg.Spec.ForProvider.Form.Parameters[i4].FillBehavior.RepromptEventHandlers[i6].TriggerFulfillment.WebhookSelector,
								To:           reference.To{List: l, Managed: m},
							})
						}
						if err != nil {
							return errors.Wrap(err, "mg.Spec.ForProvider.Form.Parameters[i4].FillBehavior.RepromptEventHandlers[i6].TriggerFulfillment.Webhook")
						}
						mg.Spec.ForProvider.Form.Parameters[i4].FillBehavior.RepromptEventHandlers[i6].TriggerFulfillment.Webhook = reference.ToPtrValue(rsp.ResolvedValue)
						mg.Spec.ForProvider.Form.Parameters[i4].FillBehavior.RepromptEventHandlers[i6].TriggerFulfillment.WebhookRef = rsp.ResolvedReference

					}
				}
			}
		}
	}
	if mg.Spec.ForProvider.KnowledgeConnectorSettings != nil {
		{
			m, l, err = apisresolver.GetManagedResource("dialogflowcx.gcp.upbound.io", "v1beta2", "Page", "PageList")
			if err != nil {
				return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
			}
			rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
				CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.KnowledgeConnectorSettings.TargetPage),
				Extract:      resource.ExtractResourceID(),
				Reference:    mg.Spec.ForProvider.KnowledgeConnectorSettings.TargetPageRef,
				Selector:     mg.Spec.ForProvider.KnowledgeConnectorSettings.TargetPageSelector,
				To:           reference.To{List: l, Managed: m},
			})
		}
		if err != nil {
			return errors.Wrap(err, "mg.Spec.ForProvider.KnowledgeConnectorSettings.TargetPage")
		}
		mg.Spec.ForProvider.KnowledgeConnectorSettings.TargetPage = reference.ToPtrValue(rsp.ResolvedValue)
		mg.Spec.ForProvider.KnowledgeConnectorSettings.TargetPageRef = rsp.ResolvedReference

	}
	if mg.Spec.ForProvider.KnowledgeConnectorSettings != nil {
		if mg.Spec.ForProvider.KnowledgeConnectorSettings.TriggerFulfillment != nil {
			{
				m, l, err = apisresolver.GetManagedResource("dialogflowcx.gcp.upbound.io", "v1beta2", "Webhook", "WebhookList")
				if err != nil {
					return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
				}
				rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
					CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.KnowledgeConnectorSettings.TriggerFulfillment.Webhook),
					Extract:      resource.ExtractResourceID(),
					Reference:    mg.Spec.ForProvider.KnowledgeConnectorSettings.TriggerFulfillment.WebhookRef,
					Selector:     mg.Spec.ForProvider.KnowledgeConnectorSettings.TriggerFulfillment.WebhookSelector,
					To:           reference.To{List: l, Managed: m},
				})
			}
			if err != nil {
				return errors.Wrap(err, "mg.Spec.ForProvider.KnowledgeConnectorSettings.TriggerFulfillment.Webhook")
			}
			mg.Spec.ForProvider.KnowledgeConnectorSettings.TriggerFulfillment.Webhook = reference.ToPtrValue(rsp.ResolvedValue)
			mg.Spec.ForProvider.KnowledgeConnectorSettings.TriggerFulfillment.WebhookRef = rsp.ResolvedReference

		}
	}
	{
		m, l, err = apisresolver.GetManagedResource("dialogflowcx.gcp.upbound.io", "v1beta2", "Agent", "AgentList")
		if err != nil {
			return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
		}
		rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
			CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.Parent),
			Extract:      resource.ExtractParamPath("start_flow", true),
			Reference:    mg.Spec.ForProvider.ParentRef,
			Selector:     mg.Spec.ForProvider.ParentSelector,
			To:           reference.To{List: l, Managed: m},
		})
	}
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.Parent")
	}
	mg.Spec.ForProvider.Parent = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.ParentRef = rsp.ResolvedReference

	for i3 := 0; i3 < len(mg.Spec.ForProvider.TransitionRoutes); i3++ {
		{
			m, l, err = apisresolver.GetManagedResource("dialogflowcx.gcp.upbound.io", "v1beta2", "Page", "PageList")
			if err != nil {
				return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
			}
			rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
				CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.TransitionRoutes[i3].TargetPage),
				Extract:      resource.ExtractResourceID(),
				Reference:    mg.Spec.ForProvider.TransitionRoutes[i3].TargetPageRef,
				Selector:     mg.Spec.ForProvider.TransitionRoutes[i3].TargetPageSelector,
				To:           reference.To{List: l, Managed: m},
			})
		}
		if err != nil {
			return errors.Wrap(err, "mg.Spec.ForProvider.TransitionRoutes[i3].TargetPage")
		}
		mg.Spec.ForProvider.TransitionRoutes[i3].TargetPage = reference.ToPtrValue(rsp.ResolvedValue)
		mg.Spec.ForProvider.TransitionRoutes[i3].TargetPageRef = rsp.ResolvedReference

	}
	if mg.Spec.InitProvider.Form != nil {
		for i4 := 0; i4 < len(mg.Spec.InitProvider.Form.Parameters); i4++ {
			if mg.Spec.InitProvider.Form.Parameters[i4].FillBehavior != nil {
				for i6 := 0; i6 < len(mg.Spec.InitProvider.Form.Parameters[i4].FillBehavior.RepromptEventHandlers); i6++ {
					{
						m, l, err = apisresolver.GetManagedResource("dialogflowcx.gcp.upbound.io", "v1beta2", "Agent", "AgentList")
						if err != nil {
							return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
						}
						rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
							CurrentValue: reference.FromPtrValue(mg.Spec.InitProvider.Form.Parameters[i4].FillBehavior.RepromptEventHandlers[i6].TargetFlow),
							Extract:      resource.ExtractParamPath("start_flow", true),
							Reference:    mg.Spec.InitProvider.Form.Parameters[i4].FillBehavior.RepromptEventHandlers[i6].TargetFlowRef,
							Selector:     mg.Spec.InitProvider.Form.Parameters[i4].FillBehavior.RepromptEventHandlers[i6].TargetFlowSelector,
							To:           reference.To{List: l, Managed: m},
						})
					}
					if err != nil {
						return errors.Wrap(err, "mg.Spec.InitProvider.Form.Parameters[i4].FillBehavior.RepromptEventHandlers[i6].TargetFlow")
					}
					mg.Spec.InitProvider.Form.Parameters[i4].FillBehavior.RepromptEventHandlers[i6].TargetFlow = reference.ToPtrValue(rsp.ResolvedValue)
					mg.Spec.InitProvider.Form.Parameters[i4].FillBehavior.RepromptEventHandlers[i6].TargetFlowRef = rsp.ResolvedReference

				}
			}
		}
	}
	if mg.Spec.InitProvider.Form != nil {
		for i4 := 0; i4 < len(mg.Spec.InitProvider.Form.Parameters); i4++ {
			if mg.Spec.InitProvider.Form.Parameters[i4].FillBehavior != nil {
				for i6 := 0; i6 < len(mg.Spec.InitProvider.Form.Parameters[i4].FillBehavior.RepromptEventHandlers); i6++ {
					{
						m, l, err = apisresolver.GetManagedResource("dialogflowcx.gcp.upbound.io", "v1beta2", "Page", "PageList")
						if err != nil {
							return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
						}
						rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
							CurrentValue: reference.FromPtrValue(mg.Spec.InitProvider.Form.Parameters[i4].FillBehavior.RepromptEventHandlers[i6].TargetPage),
							Extract:      resource.ExtractResourceID(),
							Reference:    mg.Spec.InitProvider.Form.Parameters[i4].FillBehavior.RepromptEventHandlers[i6].TargetPageRef,
							Selector:     mg.Spec.InitProvider.Form.Parameters[i4].FillBehavior.RepromptEventHandlers[i6].TargetPageSelector,
							To:           reference.To{List: l, Managed: m},
						})
					}
					if err != nil {
						return errors.Wrap(err, "mg.Spec.InitProvider.Form.Parameters[i4].FillBehavior.RepromptEventHandlers[i6].TargetPage")
					}
					mg.Spec.InitProvider.Form.Parameters[i4].FillBehavior.RepromptEventHandlers[i6].TargetPage = reference.ToPtrValue(rsp.ResolvedValue)
					mg.Spec.InitProvider.Form.Parameters[i4].FillBehavior.RepromptEventHandlers[i6].TargetPageRef = rsp.ResolvedReference

				}
			}
		}
	}
	if mg.Spec.InitProvider.Form != nil {
		for i4 := 0; i4 < len(mg.Spec.InitProvider.Form.Parameters); i4++ {
			if mg.Spec.InitProvider.Form.Parameters[i4].FillBehavior != nil {
				for i6 := 0; i6 < len(mg.Spec.InitProvider.Form.Parameters[i4].FillBehavior.RepromptEventHandlers); i6++ {
					if mg.Spec.InitProvider.Form.Parameters[i4].FillBehavior.RepromptEventHandlers[i6].TriggerFulfillment != nil {
						{
							m, l, err = apisresolver.GetManagedResource("dialogflowcx.gcp.upbound.io", "v1beta2", "Webhook", "WebhookList")
							if err != nil {
								return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
							}
							rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
								CurrentValue: reference.FromPtrValue(mg.Spec.InitProvider.Form.Parameters[i4].FillBehavior.RepromptEventHandlers[i6].TriggerFulfillment.Webhook),
								Extract:      resource.ExtractResourceID(),
								Reference:    mg.Spec.InitProvider.Form.Parameters[i4].FillBehavior.RepromptEventHandlers[i6].TriggerFulfillment.WebhookRef,
								Selector:     mg.Spec.InitProvider.Form.Parameters[i4].FillBehavior.RepromptEventHandlers[i6].TriggerFulfillment.WebhookSelector,
								To:           reference.To{List: l, Managed: m},
							})
						}
						if err != nil {
							return errors.Wrap(err, "mg.Spec.InitProvider.Form.Parameters[i4].FillBehavior.RepromptEventHandlers[i6].TriggerFulfillment.Webhook")
						}
						mg.Spec.InitProvider.Form.Parameters[i4].FillBehavior.RepromptEventHandlers[i6].TriggerFulfillment.Webhook = reference.ToPtrValue(rsp.ResolvedValue)
						mg.Spec.InitProvider.Form.Parameters[i4].FillBehavior.RepromptEventHandlers[i6].TriggerFulfillment.WebhookRef = rsp.ResolvedReference

					}
				}
			}
		}
	}
	if mg.Spec.InitProvider.KnowledgeConnectorSettings != nil {
		{
			m, l, err = apisresolver.GetManagedResource("dialogflowcx.gcp.upbound.io", "v1beta2", "Page", "PageList")
			if err != nil {
				return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
			}
			rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
				CurrentValue: reference.FromPtrValue(mg.Spec.InitProvider.KnowledgeConnectorSettings.TargetPage),
				Extract:      resource.ExtractResourceID(),
				Reference:    mg.Spec.InitProvider.KnowledgeConnectorSettings.TargetPageRef,
				Selector:     mg.Spec.InitProvider.KnowledgeConnectorSettings.TargetPageSelector,
				To:           reference.To{List: l, Managed: m},
			})
		}
		if err != nil {
			return errors.Wrap(err, "mg.Spec.InitProvider.KnowledgeConnectorSettings.TargetPage")
		}
		mg.Spec.InitProvider.KnowledgeConnectorSettings.TargetPage = reference.ToPtrValue(rsp.ResolvedValue)
		mg.Spec.InitProvider.KnowledgeConnectorSettings.TargetPageRef = rsp.ResolvedReference

	}
	if mg.Spec.InitProvider.KnowledgeConnectorSettings != nil {
		if mg.Spec.InitProvider.KnowledgeConnectorSettings.TriggerFulfillment != nil {
			{
				m, l, err = apisresolver.GetManagedResource("dialogflowcx.gcp.upbound.io", "v1beta2", "Webhook", "WebhookList")
				if err != nil {
					return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
				}
				rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
					CurrentValue: reference.FromPtrValue(mg.Spec.InitProvider.KnowledgeConnectorSettings.TriggerFulfillment.Webhook),
					Extract:      resource.ExtractResourceID(),
					Reference:    mg.Spec.InitProvider.KnowledgeConnectorSettings.TriggerFulfillment.WebhookRef,
					Selector:     mg.Spec.InitProvider.KnowledgeConnectorSettings.TriggerFulfillment.WebhookSelector,
					To:           reference.To{List: l, Managed: m},
				})
			}
			if err != nil {
				return errors.Wrap(err, "mg.Spec.InitProvider.KnowledgeConnectorSettings.TriggerFulfillment.Webhook")
			}
			mg.Spec.InitProvider.KnowledgeConnectorSettings.TriggerFulfillment.Webhook = reference.ToPtrValue(rsp.ResolvedValue)
			mg.Spec.InitProvider.KnowledgeConnectorSettings.TriggerFulfillment.WebhookRef = rsp.ResolvedReference

		}
	}
	{
		m, l, err = apisresolver.GetManagedResource("dialogflowcx.gcp.upbound.io", "v1beta2", "Agent", "AgentList")
		if err != nil {
			return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
		}
		rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
			CurrentValue: reference.FromPtrValue(mg.Spec.InitProvider.Parent),
			Extract:      resource.ExtractParamPath("start_flow", true),
			Reference:    mg.Spec.InitProvider.ParentRef,
			Selector:     mg.Spec.InitProvider.ParentSelector,
			To:           reference.To{List: l, Managed: m},
		})
	}
	if err != nil {
		return errors.Wrap(err, "mg.Spec.InitProvider.Parent")
	}
	mg.Spec.InitProvider.Parent = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.InitProvider.ParentRef = rsp.ResolvedReference

	for i3 := 0; i3 < len(mg.Spec.InitProvider.TransitionRoutes); i3++ {
		{
			m, l, err = apisresolver.GetManagedResource("dialogflowcx.gcp.upbound.io", "v1beta2", "Page", "PageList")
			if err != nil {
				return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
			}
			rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
				CurrentValue: reference.FromPtrValue(mg.Spec.InitProvider.TransitionRoutes[i3].TargetPage),
				Extract:      resource.ExtractResourceID(),
				Reference:    mg.Spec.InitProvider.TransitionRoutes[i3].TargetPageRef,
				Selector:     mg.Spec.InitProvider.TransitionRoutes[i3].TargetPageSelector,
				To:           reference.To{List: l, Managed: m},
			})
		}
		if err != nil {
			return errors.Wrap(err, "mg.Spec.InitProvider.TransitionRoutes[i3].TargetPage")
		}
		mg.Spec.InitProvider.TransitionRoutes[i3].TargetPage = reference.ToPtrValue(rsp.ResolvedValue)
		mg.Spec.InitProvider.TransitionRoutes[i3].TargetPageRef = rsp.ResolvedReference

	}

	return nil
}

// ResolveReferences of this Webhook.
func (mg *Webhook) ResolveReferences(ctx context.Context, c client.Reader) error {
	var m xpresource.Managed
	var l xpresource.ManagedList
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error
	{
		m, l, err = apisresolver.GetManagedResource("dialogflowcx.gcp.upbound.io", "v1beta2", "Agent", "AgentList")
		if err != nil {
			return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
		}

		rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
			CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.Parent),
			Extract:      resource.ExtractResourceID(),
			Reference:    mg.Spec.ForProvider.ParentRef,
			Selector:     mg.Spec.ForProvider.ParentSelector,
			To:           reference.To{List: l, Managed: m},
		})
	}
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.Parent")
	}
	mg.Spec.ForProvider.Parent = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.ParentRef = rsp.ResolvedReference
	{
		m, l, err = apisresolver.GetManagedResource("dialogflowcx.gcp.upbound.io", "v1beta2", "Agent", "AgentList")
		if err != nil {
			return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
		}

		rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
			CurrentValue: reference.FromPtrValue(mg.Spec.InitProvider.Parent),
			Extract:      resource.ExtractResourceID(),
			Reference:    mg.Spec.InitProvider.ParentRef,
			Selector:     mg.Spec.InitProvider.ParentSelector,
			To:           reference.To{List: l, Managed: m},
		})
	}
	if err != nil {
		return errors.Wrap(err, "mg.Spec.InitProvider.Parent")
	}
	mg.Spec.InitProvider.Parent = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.InitProvider.ParentRef = rsp.ResolvedReference

	return nil
}
