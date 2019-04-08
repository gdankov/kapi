# KAPI - kubernetes CRDs for CF APIs

Hi. This is a proof-of-concept/spike to see what it would look like to
replicate the CF v3 APIs with Kubernetes CRDs.

# (Potential) Goals

- See whether we can support existing CF APIs on top of CRDs
- See whether we can use Eirini as a control loop from "Kapi" CRDs to k8s
- See whether we can implement an org/space style model with RBAC
- Experiment with whether it's better to do a CLI using k8s apis or with a wrapper around CRDs
- Determine what an "MVP" would look like for this to avoid boiling the ocean

# Non-Goals

- Actual production code. This is a dirty rotten spike.

# What might success look like?

- end-to-end cf v3-push via cli (potentially wrapping kubectl, against CAPI api wrapper around CRDs), resulting in working statefulsets (probably via eirini)

# How does/would KAPI differ from Knative/StatefulSets/Deployments

KAPI apis are *PaaS* APIs. This means they are high-level (packages, droplets,
builds, apps, processes, services, bindings) and do not talk about "containers", 
"images" etc (other than permitting an optional "image" package/droplet type). 

Under the covers the KAPI APIs may end up being mapped to containers or statefulsets
or knative kservices etc but this should happen in the Operator (eirini), rather
than being exposed in the KAPI CRDs.
