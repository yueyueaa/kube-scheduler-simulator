package yuescheduler

import (
	"context"
	"errors"
	"strconv"

	"golang.org/x/xerrors"
	v1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/klog/v2"
	"k8s.io/kubernetes/pkg/scheduler/framework"
	frameworkruntime "k8s.io/kubernetes/pkg/scheduler/framework/runtime"
)

type NodeResourcesBalancedAllocation struct {
	reverse bool
}

var (
	_ framework.ScorePlugin    = &NodeResourcesBalancedAllocation{}
	_ framework.PreScorePlugin = &NodeResourcesBalancedAllocation{}
)

const (
	Name             = "NodeResourcesBalancedAllocation"
	preScoreStateKey = "PreScore" + Name
)

func (pl *NodeResourcesBalancedAllocation) Name() string {
	return Name
}

type preScoreState struct {
	podSuffixNumber int
}

func (s *preScoreState) Clone() framework.StateData {
	return s
}

func (pl *NodeResourcesBalancedAllocation) PreScore(ctx context.Context, state *framework.CycleState, pod *v1.Pod, nodes []*v1.Node) *framework.Status {
	klog.InfoS("execute PreScore on NodeResourcesBalancedAllocation plugin", "pod", klog.KObj(pod))

	podNameLastChar := pod.Name[len(pod.Name)-1:]
	podnum, err := strconv.Atoi(podNameLastChar)
	if err != nil {
		return nil
	}

	s := &preScoreState{
		podSuffixNumber: podnum,
	}
	state.Write(preScoreStateKey, s)

	return nil
}

func (pl *NodeResourcesBalancedAllocation) EventsToRegister() []framework.ClusterEvent {
	return []framework.ClusterEvent{
		{Resource: framework.Node, ActionType: framework.Add},
	}
}

var ErrNotExpectedPreScoreState = errors.New("unexpected pre score state")

func (pl *NodeResourcesBalancedAllocation) Score(ctx context.Context, state *framework.CycleState, pod *v1.Pod, nodeName string) (int64, *framework.Status) {
	klog.InfoS("execute Score on NodeResourcesBalancedAllocation plugin", "pod", klog.KObj(pod))
	data, err := state.Read(preScoreStateKey)
	if err != nil {
		return 0, nil
	}

	s, ok := data.(*preScoreState)
	if !ok {
		err = xerrors.Errorf("fetched pre score state is not *preScoreState, but %T, %w", data, ErrNotExpectedPreScoreState)
		return 0, framework.AsStatus(err)
	}

	nodeNameLastChar := nodeName[len(nodeName)-1:]

	nodenum, err := strconv.Atoi(nodeNameLastChar)
	if err != nil {
		return 0, nil
	}

	var matchScore int64 = 10
	var nonMatchScore int64 = 0
	if pl.reverse {
		matchScore = 0
		nonMatchScore = 10
	}

	if s.podSuffixNumber == nodenum {
		return matchScore, nil
	}

	return nonMatchScore, nil
}

// ScoreExtensions of the Score plugin.
func (pl *NodeResourcesBalancedAllocation) ScoreExtensions() framework.ScoreExtensions {
	return nil
}

// New initializes a new plugin and returns it.
func New(ctx context.Context, arg runtime.Object, h framework.Handle) (framework.Plugin, error) {
	typedArg := NodeResourcesBalancedAllocationArgs{Reverse: false}
	if arg != nil {
		err := frameworkruntime.DecodeInto(arg, &typedArg)
		if err != nil {
			return nil, xerrors.Errorf("decode arg into NodeResourcesBalancedAllocationArgs: %w", err)
		}
		klog.Info("NodeResourcesBalancedAllocationArgs is successfully applied")
	}
	return &NodeResourcesBalancedAllocation{reverse: typedArg.Reverse}, nil
}

// NodeResourcesBalancedAllocationArgs is arguments for node number plugin.
//
//nolint:revive
type NodeResourcesBalancedAllocationArgs struct {
	metav1.TypeMeta

	Reverse bool `json:"reverse"`
}
