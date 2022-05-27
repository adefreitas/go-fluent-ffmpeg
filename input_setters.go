package fluentffmpeg

import "io"

func (c *Command) Input(v string, pipeInput io.Reader, format string, nativeFrameRate bool, options ...string) *Command {
	c.Args.inputs = append(c.Args.inputs, inputArgs{inputPath: v, pipeInput: pipeInput != nil, fromFormat: format, nativeFramerateInput: nativeFrameRate, inputOptions: options})
	return c
}

// Options is intended for configuring global options that are not affected by their position in the FFmpeg command
func (c *Command) Options(options ...string) *Command {
	c.Args.globalOptions = options

	return c
}
