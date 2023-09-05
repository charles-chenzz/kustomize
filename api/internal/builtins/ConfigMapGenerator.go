// Code generated by pluginator on ConfigMapGenerator; DO NOT EDIT.
// pluginator {(devel)  unknown   }



package builtins

import (
	"sigs.k8s.io/kustomize/api/kv"
	"sigs.k8s.io/kustomize/api/resmap"
	"sigs.k8s.io/kustomize/api/types"
	"sigs.k8s.io/yaml"
)

type ConfigMapGeneratorPlugin struct {
	h                *resmap.PluginHelpers
	types.ObjectMeta `json:"metadata,omitempty" yaml:"metadata,omitempty" protobuf:"bytes,1,opt,name=metadata"`
	types.ConfigMapArgs
}


func (p *ConfigMapGeneratorPlugin) Config(h *resmap.PluginHelpers, config []byte) (err error) {
	p.ConfigMapArgs = types.ConfigMapArgs{}
	err = yaml.Unmarshal(config, p)
	if p.ConfigMapArgs.Name == "" {
		p.ConfigMapArgs.Name = p.Name
	}
	if p.ConfigMapArgs.Namespace == "" {
		p.ConfigMapArgs.Namespace = p.Namespace
	}
	p.h = h
	return
}

func (p *ConfigMapGeneratorPlugin) Generate() (resmap.ResMap, []string, error) {
	return p.h.ResmapFactory().FromConfigMapArgs(
		kv.NewLoader(p.h.Loader(), p.h.Validator()), p.ConfigMapArgs)
}

func NewConfigMapGeneratorPlugin() resmap.GeneratorPlugin {
  return &ConfigMapGeneratorPlugin{}
}
