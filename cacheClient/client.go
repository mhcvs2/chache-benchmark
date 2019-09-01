package cacheClient

type Cmd struct {
	Name  string
	Key   string
	Value string
	Error error
}

type Client interface {
	Run(cmd *Cmd)
	PipelineRun(cmds []*Cmd)
}

func New(typ, server string) Client {
	if typ == "http" {
		return NewHttpClient(server)
	}
	panic("unknown client type " + typ)
}
